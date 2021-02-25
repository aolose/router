package anysrv

type rawNode struct {
	parent  *rawNode
	right   *rawNode
	nodes   []*rawNode
	deep    int
	path    string
	handler Handler
	skip    bool
	params  []*param
}

func (r *rawNode) toNode(n, p *node) *node {
	if r.skip {
		if len(r.nodes) > 0 {
			return r.nodes[0].toNode(n, p)
		} else if r.handler != nil && p != nil {
			p.handler = r.handler
			p.params = r.params
		}
		return nil
	}
	if n == nil {
		n = &node{}
	}
	n.path = r.path
	n.handler = r.handler
	n.deep = r.deep
	n.params = r.params
	if len(r.nodes) > 0 {
		r = r.nodes[0]
		n.next = r.toNode(nil, n)
	}
	if r.right != nil {
		n.right = r.right.toNode(nil, nil)
	}
	return n
}

func (r *rawNode) new(path string) *rawNode {
	return addRawNode(r, &r.nodes, path)
}

func (r rawNode) ready() {
	ns := r.nodes
	sortRawNode(ns)
	for _, n := range ns {
		n.ready()
	}
}
