package anysrv

type node struct {
	parent  *node
	right   *node
	next    *node
	nodes   []*node
	deep    int
	path    string
	handler Handler
	skip    bool
	params  *[]*param
}

func (r *node) new(path string) *node {
	return addRawNode(r, &r.nodes, path)
}

func readNs(ns *[]*node, p *node) {
	for _, n := range *ns {
		readNs(&n.nodes, n)
	}
	l := len(*ns)
	for i := 0; i < l; i++ {
		n := (*ns)[i]
		if n.skip {
			if len((*n).nodes) == 0 {
				if n.handler != nil && p != nil {
					p.params = n.params
					p.handler = n.handler
					*ns = append((*ns)[0:i], (*ns)[i+1:]...)
					l--
					i--
				}
			} else {
				*ns = append(append((*ns)[0:i], (*ns)[i+1:]...), n.nodes...)
				i--
				l--
			}
		}
	}
	sortRawNode(*ns)
}

func (n *node) lookup(path *string) (Handler, *[]*param) {
	s := share[0][n.deep]
	e := share[1][n.deep]
	if n.path == (*path)[s:e] {
		if n.next != nil {
			h, d := n.next.lookup(path)
			if h != nil {
				return h, d
			}
		}
		if n.handler != nil {
			return n.handler, n.params
		}
	}

	if n.right != nil {
		return n.right.lookup(path)
	}
	return nil, nil
}
