package anysrv

type tree struct {
	deep   int
	levels []*level
}

func (r *tree) String() string {
	s := "{"
	for i := 0; i < len(r.levels); i++ {
		s = s + r.levels[i].String() + " "
	}
	return s + "\n}\n"
}

func newTree(deep, beginCap int) *tree {
	r := &tree{
		deep:   deep,
		levels: make([]*level, deep, deep),
	}
	for i := 0; i < deep; i++ {
		r.levels[i] = newLevel(beginCap)
	}
	return r
}

func (r *tree) ready() {
	w := len(r.levels)
	for i := 0; i < w; i++ {
		if i > 0 {
			r.levels[i].sort()
		}
	}
	for d := 0; d < r.deep; d++ {
		ns := r.levels[d].nodes
		l := len(ns)
		var before *node
		for i := 0; i < l; i++ {
			n := ns[i]
			n.deep = d
			if n.parent != nil && n.parent.next == nil {
				n.parent.next = n
			}
			if n != before {
				if before != nil {
					if before.parent == n.parent {
						before.right = n
					} else {
						if before.parent != nil {
							before.right = before.parent.right
						}
					}
				}
				before = n
			}
		}
	}
}

type staticHandler struct {
	path    string
	handler Handler
}

type staticData struct {
	length     int
	paths      []*staticHandler
	startIndex int
}

func (sd *staticData) add(path string, handler Handler) {
	ps := sd.paths
	l := len(ps)
	if l == cap(ps) {
		v := make([]*staticHandler, l+1, l+1)
		copy(v, ps)
		ps = v
	}
	sd.paths = ps[:l+1]
	sd.paths[l] = &staticHandler{path: path, handler: handler}
}

func (sd *staticData) macth(path string) Handler {
	i := sd.startIndex
	e := len(sd.paths)
	s := -1
	m := sd.length
	for s < i && i < e {
		c := sd.paths[i]
		p := c.path
		n := 0
		for ; n < m; n++ {
			if p[n] < path[n] {
				s = i
				i = (i + e + 1) / 2
				break
			}
			if p[n] > path[n] {
				e = i
				i = (i + s) / 2
				break
			}
		}
		if n == m {
			return c.handler
		}
	}
	return nil
}

type deepTree struct {
	max    int
	trees  []*tree
	cache  []string
	static []*staticData
}

func quickFind(sd []*staticData, n int) *staticData {
	e := len(sd)
	s := -1
	for m := e / 2; s < m && m < e; {
		v := sd[m].length
		if n > v {
			s = m
			m = (m + e + 1) / 2
			continue
		}
		if n < v {
			e = m
			m = (s + m) / 2
			continue
		}
		return sd[m]
	}
	return nil
}
