package anysrv

type router struct {
	deep    int
	maxDeep int
	levels  []*level
	cache   []string
}

func (r *router) increase() {
	pre := cap(r.levels[r.deep-1].nodes)
	r.deep++
	if r.maxDeep < r.deep {
		r.maxDeep = r.maxDeep * 2
		v := make([]*level, r.deep, r.maxDeep)
		copy(v, r.levels)
		r.levels = v
	} else {
		r.levels = r.levels[:r.deep]
	}
	r.levels[r.deep-1] = newLevel(pre * 2)
}

func (r *router) String() string {
	s := "{"
	for i := 0; i < len(r.levels); i++ {
		s = s + r.levels[i].String() + " "
	}
	return s + "\n}\n"
}

func newRouter(deep, beginCap int) *router {
	r := &router{
		maxDeep: deep,
		deep:    1,
		levels:  make([]*level, 1, deep),
	}
	r.levels[0] = newLevel(beginCap)
	return r
}

func (r *router) ready() {
	r.cache = make([]string, r.deep, r.deep)
	w := len(r.levels)
	for i := 0; i < w; i++ {
		if i > 0 {
			r.levels[i].sort()
		}
		c := r.deep - i - 1
		ns := r.levels[i].nodes
		l := len(ns)
		for j := 0; j < l; j++ {
			if c > 0 {
				s := make([]int, c, c)
				ns[j].start = s
				for t := 0; t < c; t++ {
					s[t] = -1
				}
			}
		}
	}
	for i := r.deep - 1; i > 0; i-- {
		lv := r.levels[i]
		var p *node
		ns := lv.nodes
		l := len(ns)
		for j := 0; j < l; j++ {
			if ns[j].parent != p {
				p = ns[j].parent
				if len(ns[j].start) > 0 {
					copy(p.start[1:], ns[j].start)
				}
				p.start[0] = j
			}
		}
	}
}

func (r *router) Lookup(method, path string) (Handler, *node) {
	m := getMethodCode(method)
	d := 0
	lookup(path, func(start, end int) bool {
		if d > r.deep-1 {
			d = r.deep + 1
			return true
		}
		r.cache[d] = path[start:end]
		d++
		return false
	})
	if d <= r.deep {
		ns := r.levels[d-1].nodes
		for e := len(ns) - 1; e > -1; e-- {
			n := ns[e]
			h := n.handle[m]
			if h != nil {
				ok, i := n.match(r.cache[:d])
				if ok {
					return h, n
				}
				if i != -1 {
					e = i
				}
			}
		}
	}
	return nil, nil
}

func (r *router) bind(m int, path string, h Handler) {
	dp := 0
	var pr *node
	lookup(path, func(start, end int) bool {
		p := path[start:end]
		if dp == r.deep {
			r.increase()
		}
		pr = r.levels[dp].bind(pr, p)
		dp++
		return false
	})
	if pr != nil {
		pr.handle[m] = h
	}
}