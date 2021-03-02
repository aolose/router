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
			e := l
			s := -1
			for m := l / 2; s < e && m > s && m < e; {
				p := st[m]
				i := 0
				for ; i < n; i++ {
					c0 := p.path[i]
					c1 := (*path)[i+1]
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
					return p.handler
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
	allParams := !isStatic
	s := make([]int, l/2+1)
	e := make([]int, l/2+1)
	n := make([]int, l/2+1)
	d := 0
	m := 0
	for i := 1; i < l-1; i++ {
		c := path[i]
		if c == ':' {
			n[m] = d
			m++
			isStatic = false
		} else {
			if c == '/' {
				if path[i+1] != ':' {
					allParams = false
				}
				s[d+1] = i + 1
				e[d] = i
				d++
			}
		}
	}
	e[d] = l
	s = s[:d+1]
	e = e[:d+1]
	n = n[:m+1]
	ts := r.trees[code]
	if d >= len(ts) {
		tt := make([]*tree, d+1, d+1)
		copy(tt, ts)
		tt[d] = &tree{
			static: make([][]*staticNode, 0, 0),
			raw:    make([]*node, 0, 0),
			nodes:  make([]*node, 0, 0),
		}
		ts = tt
		r.trees[code] = tt
	}
	t := ts[d]
	if t == nil {
		t = &tree{
			static: make([][]*staticNode, 0, 0),
			raw:    make([]*node, 0, 0),
		}
		ts[d] = t
	}
	if isStatic {
		r.addStatic(code, path, h)
	} else if allParams {
		if len(t.nodes) < d+1 {
			a := make([]*node, d+1)
			copy(a, t.nodes)
			t.nodes = a
		}
		t.nodes = make([]*node, d+1)
		for i := 0; i <= d; i++ {
			pm := make([]*param, d+1)
			nd := &node{
				handler: h,
				deep:    i,
				params:  &pm,
			}
			t.nodes[i] = nd
			for j := 0; j <= d; j++ {
				(*nd.params)[j] = &param{
					name: path[s[j]+1 : e[j]],
					deep: j,
				}
			}
		}
	} else {
		t.addNode(path, h, s, e, n)
	}
}

func (r *router) Lookup(method string, path string) (Handler, *[]*param) {
	l := len(path)
	cd := getMethodCode(method)
	if l > 1 && path[l-1] == '/' {
		l--
	}
	s := r.getStatic(cd, l-1, &path)
	if s != nil {
		return s, nil
	}
	d := 0
	i := 1
	for ; i < l-1; i++ {
		if path[i] == '/' {
			share[0][d+1] = i + 1
			share[1][d] = i
			d++
		}
	}
	share[1][d] = l
	ts := r.trees[cd]
	if len(ts) > d {
		t := ts[d]
		if t != nil {
			return t.lookup(&path, d, l-1)
		}
	}
	return nil, nil
}
