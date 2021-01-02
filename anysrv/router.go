package anysrv

type router [7]*deepTree

func (r *router) Lookup(method, path string) (Handler, *node) {
	dt := r[getMethodCode(method)]
	d := 0
	lookup(path, func(start, end int) bool {
		dt.cache[d] = path[start:end]
		if d > dt.max {
			dt.cache[d] += "?"
			return true
		}
		d++
		return false
	})
	tre := dt.trees[d-1]
	if tre != nil {
		ns := tre.levels[d-1].nodes
		for e := len(ns) - 1; e > -1; e-- {
			n := ns[e]
			o, i := n.match(dt.cache[:d])
			if o {
				return n.handle, n
			}
			if i != -1 {
				e = i
			}
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
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
		&deepTree{trees: make([]*tree, 0, 32)},
	}
}
