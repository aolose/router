package anysrv

import "sort"

type tree struct {
	static [][]*staticNode
	raw    []*rawNode
	node   *node
	nodes  []*node
}

func (t *tree) addNode(path string, h Handler, start, end, pm []int) {
	ps := make([]*param, len(pm))
	for i, d := range pm {
		ps[i] = &param{
			name: path[start[d]:end[d]],
			deep: d,
		}
	}
	var r *rawNode

	l := len(start)
	for i := 0; i < l; i++ {
		s := start[i]
		e := end[i]
		p := path[s:e]
		r = addRawNode(r, &t.raw, p)
	}
	if r != nil {
		r.handler = h
		r.params = ps
	}
}
func (t *tree) addStatic(path string, h Handler) {
	l := len(path)
	st := t.static
	if len(st) <= l {
		ss := make([][]*staticNode, l+1, l+1)
		copy(ss, st)
		ss[l] = make([]*staticNode, 0, 0)
		t.static = ss
	}
	s := t.static[l]
	for _, p := range s {
		if p.path == path {
			return
		}
	}
	n := len(s) + 1
	ss := make([]*staticNode, n, n)
	n--
	copy(ss, s)
	ss[n] = &staticNode{
		path:    path,
		handler: h,
	}
	t.static[l] = ss
}
func (t *tree) ready() {
	sortRawNode(t.raw)
	for _, r := range t.raw {
		r.ready()
	}
	for _, a := range t.static {
		if a != nil {
			sort.Slice(a, func(i, j int) bool {
				return sort.StringsAreSorted([]string{
					a[i].path,
					a[j].path,
				})
			})
		}
	}

	nd := &node{}
	if len(t.raw) > 0 {
		t.raw[0].toNode(nd, nd)
	}
	t.node = nd
	t.raw = nil
}

func (t *tree) lookup(path *string, rq *reqPath) (Handler, []*param) {
	n := rq.length
	if len(t.static) > n {
		st := t.static[n]
		if st != nil {
			l := len(st)
			e := l
			s := -1
			pt := (*path)[1:]
			for m := l / 2; s < e && m > s && m < e; {
				p := st[m]
				i := 0
				for ; i < n; i++ {
					c0 := p.path[i]
					c1 := pt[i]
					if c0 == c1 {
						continue
					}
					if c0 > c1 {
						e = m
						m = (e + s + 1) / 2
						break
					} else {
						s = m
						m = (e + s + 1) / 2
						break
					}
				}
				if i == n {
					return p.handler, nil
				}
			}
		}
	}
	if t.node != nil {
		a, b := t.node.lookup(path, rq)
		if a != nil {
			return a, b
		}
	}
	if len(t.nodes) > rq.deep {
		d := t.nodes[rq.deep]
		if d != nil {
			return d.handler, d.params
		}
	}
	return nil, nil
}
