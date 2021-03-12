package anysrv

import "sort"

type router struct {
	trees   [7][]*tree
	statics [7][][]*staticNode
}

func (r *router) addStatic(method int, path string, h Handler) {
	st := r.statics[method]
	l := len(st)
	p := len(path)
	var ss []*staticNode
	if p < l {
		ss = st[p]
	} else {
		tt := make([][]*staticNode, p+1, p+1)
		copy(tt, st)
		st = tt
		r.statics[method] = tt
		ss = make([]*staticNode, 0, 0)
		st[p] = ss
	}
	ls := len(ss)
	for i := 0; i < ls; i++ {
		sn := ss[i]
		if sn.path == path {
			return
		}
	}
	ns := make([]*staticNode, ls+1, ls+1)
	copy(ns, ss)
	ss = ns
	st[p] = ss
	ss[ls] = &staticNode{
		path:    path,
		handler: h,
	}
}

func (r *router) getStatic(method, n int, path *string) Handler {
	sa := r.statics[method]
	ll := len(sa)
	if ll > n {
		st := sa[n]
		if st != nil {
			l := len(st)
			if l == 1 {
				if st[0].path == (*path)[1:] {
					return st[0].handler
				}
			} else {
				e := l
				s := -1
				for m := l / 2; m > s && m < e; {
					p := st[m]
					i := 0
					for i < n {
						c0 := p.path[i]
						i++
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
					if i == n {
						return p.handler
					}
				}
			}
		}
	}
	return nil
}

func newRouter() *router {
	r := &router{}
	for i := 0; i < 7; i++ {
		r.trees[i] = make([]*tree, 0, 0)
		r.statics[i] = make([][]*staticNode, 0, 0)
	}
	return r
}

func (r *router) ready() {
	for i := 0; i < 7; i++ {
		for _, t := range r.trees[i] {
			if t != nil {
				t.ready()
			}
		}
		for _, t := range r.statics[i] {
			if t != nil && len(t) > 1 {
				sort.Slice(t, func(i, j int) bool {
					return sort.StringsAreSorted([]string{
						t[i].path,
						t[j].path,
					})
				})
			}
		}
	}
}

func (r *router) bind(code int, path string, h Handler) {
	l := len(path)
	if l > 0 {
		if path[0] == '/' {
			path = path[1:]
			l--
		}
		if l > 0 {
			if path[l-1] == '/' {
				path = path[:l-1]
				l--
			}
		}
	}
	isStatic := l == 0 || path[0] != ':'
	s := make([]int, 0)
	e := make([]int, 0)
	n := make([]int, 0)
	d := 0
	m := 0
	for i := 1; i < l-1; i++ {
		c := path[i]
		if c == ':' {
			mark(&n, d, m)
			m++
			isStatic = false
		} else {
			if c == '/' {
				mark(&s, i+1, d+1)
				mark(&e, i, d)
				d++
			}
		}
	}
	if isStatic {
		r.addStatic(code, path, h)
	} else {
		mark(&e, l, d)
		s = s[:d+1]
		e = e[:d+1]
		n = n[:m+1]
		ts := r.trees[code]
		if d >= len(ts) {
			tt := make([]*tree, d+1, d+1)
			copy(tt, ts)
			tt[d] = &tree{
				nodes: make([][]*node, 0, 0),
			}
			ts = tt
			r.trees[code] = tt
		}
		t := ts[d]
		if t == nil {
			t = &tree{
				nodes: make([][]*node, 0, 0),
			}
			ts[d] = t
		}
		t.addNode(path, h, s, e, n)
	}
}

func (r *router) Lookup(method string, path *string) (Handler, *[]*param) {
	l := len(*path)
	cd := getMethodCode(method)
	if l > 1 && (*path)[l-1] == '/' {
		l--
	}
	s := r.getStatic(cd, l-1, path)
	if s != nil {
		return s, nil
	}
	d := 0
	i := 1
	for ; i < l-1; i++ {
		if (*path)[i] == '/' {
			mark(&mkA, i+1, d+1)
			mark(&mkB, i, d)
			d++
		}
	}
	mark(&mkB, l, d)
	if len(r.trees[cd]) > d {
		t := r.trees[cd][d]
		if t != nil {
			return lookupNs(&t.nodes, t.right, path, 0)
		}
	}
	return nil, nil
}
