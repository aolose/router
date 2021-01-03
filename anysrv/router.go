package anysrv

type router [7]*deepTree

func (r *router) Lookup(method, path string) (Handler, *node) {
	dt := r[getMethodCode(method)]
	d := 0
	if path == "/" {
		d = 1
		dt.cache[0] = ""
	} else {
		path = path[1:]
		l := len(path)
		s := 0
		e := path[l-1] != '/'
		ss := 0
		for ; s < l; s++ {
			if path[ss] == '/' {
				dt.cache[d] = path[:ss]
				d++
				path = path[ss+1:]
				ss = 0
			} else {
				ss++
			}
		}
		if e {
			dt.cache[d] = path
			d++
		}
	}
	if d > dt.max {
		return nil, nil
	}
	tre := dt.trees[d-1]
	if tre != nil {
		n := tre.levels[0].nodes[0].match(dt.cache)
		if n != nil {
			return n.handle, n
		}
	}
	return nil, nil
}

func (r *router) bind(m int, path string, h Handler) {
	dp := 0
	n := deep(path)
	if n > r[m].max {
		r[m].max = n
		tr := r[m].trees
		if n > cap(tr) {
			v := make([]*tree, n+4, n+4)
			copy(v, tr)
			r[m].trees = v
		}
		r[m].trees = r[m].trees[:n]
	}
	var pr *node
	u := n - 1
	dt := r[m].trees[u]
	if dt == nil {
		dt = newTree(n, 4)
		r[m].trees[u] = dt
	}
	lookup(path, func(start, end int) bool {
		p := path[start:end]
		pr = dt.levels[dp].bind(pr, p)
		dp++
		return false
	})
	if pr != nil {
		pr.handle = h
	}
}

func (r *router) ready() {
	for i := 0; i < 7; i++ {
		dt := r[i]
		for _, v := range dt.trees {
			if v != nil {
				v.ready()
			}
		}
		dt.cache = make([]string, dt.max, dt.max)
	}
}

func newRouter() *router {
	return &router{
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
		&deepTree{trees: make([]*tree, 0, 0)},
	}
}
