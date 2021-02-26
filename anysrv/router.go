package anysrv

type router struct {
	trees [7][]*tree
}

func newRouter() *router {
	r := &router{}
	for i := 0; i < 7; i++ {
		r.trees[i] = make([]*tree, 0, 0)
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
			if i-1 == e[d] {
				allParams = false
			}
			if c == '/' {
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
			raw:    make([]*rawNode, 0, 0),
			nodes:  make([]*node, 0, 0),
		}
		ts = tt
		r.trees[code] = tt
	}
	t := ts[d]
	if t == nil {
		t = &tree{
			static: make([][]*staticNode, 0, 0),
			raw:    make([]*rawNode, 0, 0),
		}
		ts[d] = t
	}
	if isStatic {
		t.addStatic(path, h)
	} else if allParams {
		if len(t.nodes) < d+1 {
			a := make([]*node, d+1)
			copy(a, t.nodes)
			t.nodes = a
		}
		t.nodes = make([]*node, d+1)
		for i := 0; i <= d; i++ {
			nd := &node{
				handler: h,
				deep:    i,
				params:  make([]*param, d+1),
			}
			t.nodes[i] = nd
			for j := 0; j <= d; j++ {
				nd.params[j] = &param{
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
	d := 0
	i := 1
	for ; i < l-1; i++ {
		if path[i] == '/' {
			share[0][d+1] = i + 1
			share[1][d] = i
			d++
		}
	}
	if i < l && path[i] == '/' {
		l--
	}
	share[1][d] = l
	ts := r.trees[getMethodCode(method)]
	if len(ts) > d {
		t := ts[d]
		if t != nil {
			return t.lookup(&path, d, l-1)
		}
	}
	return nil, nil
}
