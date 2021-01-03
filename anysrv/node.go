package anysrv

type node struct {
	path   string
	deep   int
	parent *node
	right  *node
	next   *node
	handle Handler
	cate   int
}

func (n *node) match(ps []string) *node {
	for n != nil {
		if match(n, ps[n.deep]) {
			if n.next == nil {
				return n
			}
			n = n.next
		} else {
			n = n.right
		}
	}
	return nil
}

func (n *node) String() string {
	return n.path
}
