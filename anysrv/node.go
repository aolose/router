package anysrv

type node struct {
	handler Handler
	next    *node
	right   *node
	deep    int
	path    string
	params  []*param
}

func (n *node) lookup(path *string, rq *reqPath) (Handler, []*param) {
	s := rq.start[n.deep]
	e := rq.end[n.deep]
	if n.path == (*path)[s:e] {
		if n.next != nil {
			h, d := n.next.lookup(path, rq)
			if h != nil {
				return h, d
			}
		}
		if n.handler != nil {
			return n.handler, n.params
		}
	}

	if n.right != nil {
		return n.right.lookup(path, rq)
	}
	return nil, nil
}
