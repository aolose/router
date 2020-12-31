package anysrv

import "fmt"

type node struct {
	path   string
	parent *node
	handle [7]Handler
	start  []int
}

func (n *node) match(ps []string) (bool, int) {
	p := n
	l := len(ps)
	d := -1
	var p0, p1 string
	var n0, n1 int
	for i := l - 1; i > -1 && p != nil; i-- {
		p0 = p.path
		p1 = ps[i]
		n0 = len(p0)
		n1 = len(p1)
		if p0 == p1 || len(p0) > 0 && (p0[0] == ':' || p0 == "*" ||
			(n1 >= n0 &&
				(p0[0] == '*' && (p0[1:] == p1[n1-n0+1:]) ||
					(p0[n0-1] == '*' && p0[:n0-1] == p1[:n0-1])))) {
			p = p.parent
			d++
		} else {
			if d == -1 || len(p.start) == 0 {
				return false, -1
			}
			return false, p.start[d]
		}
	}
	return true, -1
}

func (n *node) String() string {
	return fmt.Sprintf("(%s) - %v, ", n.path, n.start)
}
