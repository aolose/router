package anysrv

import "sort"

var share = [3][128]int{
	{1}, {1}, {1},
}

//   /a/b
func parseReqPath(path string) *reqPath {
	l := len(path) - 1
	p := 0
	s := share[0]
	e := share[1]
	s[0] = 1
	e[0] = 1
	for i := 2; i < l; i++ {
		if path[i] == '/' {
			s[p+1] = i + 1
			e[p] = i
			p++
		}
	}
	e[p] = l + 1
	return &reqPath{
		length: l,
		deep:   p,
		start:  &s,
		end:    &e,
	}
}

func addRawNode(r *rawNode, rs *[]*rawNode, path string) *rawNode {
	s := len(path) > 0 && path[0] == ':'
	if s {
		path = path[1:]
	}
	if r != nil {
		rs = &r.nodes
	}
	for _, nd := range *rs {
		if nd.path == path && nd.skip == s {
			return nd
		}
	}
	n := &rawNode{
		parent: r,
		nodes:  make([]*rawNode, 0, 0),
		deep:   0,
		path:   path,
		skip:   s,
	}
	if r != nil {
		n.deep = r.deep + 1
	}
	l := len(*rs) + 1
	nn := make([]*rawNode, l, l)
	copy(nn, *rs)
	nn[l-1] = n
	*rs = nn
	return n
}

func sortRawNode(ns []*rawNode) {
	sort.Slice(ns, func(i, j int) bool {
		if ns[i].skip != ns[j].skip {
			return !ns[i].skip
		}
		return sort.StringsAreSorted([]string{
			ns[i].path,
			ns[j].path,
		})
	})
	var d *rawNode
	for _, n := range ns {
		if d != nil {
			d.right = n
		}
		d = n
	}
	if d != nil {
		if d.parent != nil {
			d.right = d.parent.right
		}
	}
}
