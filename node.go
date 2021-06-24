package anysrv

import "sort"

type node struct {
	parent  *node
	right   *node
	nodes   [][]*node
	deep    int
	path    string
	handler Handler
	params  []*param
}

func (n *node) Add(path string) *node {
	l := len(path)
	if path[0] == ':' {
		if n.right == nil {
			n.right = &node{
				parent: n,
				path:   path,
				deep:   n.deep + 1,
				nodes:  make([][]*node, 0, 0),
			}
		}
		return n.right
	}
	if l > len(n.nodes) {
		ls := make([][]*node, l, l)
		copy(ls, n.nodes)
		n.nodes = ls
	}
	l--
	if n.nodes[l] == nil {
		n.nodes[l] = make([]*node, 0, 0)
	}

	la := len(n.nodes[l])
	for i := 0; i < la; i++ {
		if n.nodes[l][i].path == path {
			return n.nodes[l][i]
		}
	}
	nd := &node{
		parent: n,
		path:   path,
		deep:   n.deep + 1,
		nodes:  make([][]*node, 0, 0),
	}
	ln := len(n.nodes[l])
	nn := make([]*node, ln+1, ln+1)
	copy(nn, n.nodes[l])
	nn[ln] = nd
	n.nodes[l] = nn
	return nd
}

func lookupNs(ns *[][]*node, right *node, path *string, deep int) (Handler, []*param) {
	st := mkA[deep]
	en := mkB[deep]
	l := en - st - 1
	if len(*ns) > l {
		n := (*ns)[l]
		if n != nil {
			l = len(n)
			if l == 1 {
				if n[0].path == (*path)[st:en] {
					if n[0].handler != nil {
						return n[0].handler, n[0].params
					}
					return lookupNs(&n[0].nodes, n[0].right, path, n[0].deep)
				}
			} else {
				e := l
				s := -1
				for m := l / 2; m > s && m < e; {
					p := n[m]
					i := st
					for ; i < en; i++ {
						c0 := p.path[i-st]
						c1 := (*path)[i]
						if c0 == c1 {
							continue
						}
						if c0 > c1 {
							e = m
							m = (e + s) / 2
							break
						} else {
							s = m
							m = (e + s) / 2
							break
						}
					}
					if i == en {
						if p.handler != nil {
							return p.handler, p.params
						}
						return lookupNs(&p.nodes, p.right, path, p.deep)
					}
				}
			}
		}
	}
	if right != nil {
		if right.handler != nil {
			return right.handler, right.params
		}
		return lookupNs(&right.nodes, right.right, path, right.deep)
	}
	return nil, nil
}

func readNs(ns [][]*node, right *node) {
	if right != nil {
		p := right.parent
		if p != nil && p.parent != nil && len(p.nodes) == 0 {
			p = p.parent
			p.right = right
		}
		readNs(right.nodes, right.right)
	}
	for _, n := range ns {
		if n != nil {
			if len(n) > 1 {
				sort.Slice(n, func(i, j int) bool {
					return sort.StringsAreSorted([]string{
						n[i].path,
						n[j].path,
					})
				})
			}
			for _, nn := range n {
				readNs(nn.nodes, nn.right)
			}
		}
	}
}
