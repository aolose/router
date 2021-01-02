package anysrv

import "fmt"

type node struct {
	path   string
	parent *node
	handle Handler
	start  int
}

func (n *node) match(ps []string) (bool, int) {
	p := n
	l := len(ps)
	d := -1
	var p0, p1 string
	for i := l - 1; i > -1; i-- {
		p0 = p.path
		p1 = ps[i]
		if match(p0, p1) {
			p = p.parent
			d++
		} else {
			return false, p.start
		}
	}
	return true, -1
}

func (n *node) String() string {
	return fmt.Sprintf("(%s) - %v, ", n.path, n.start)
}
