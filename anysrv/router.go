package anysrv

import (
	"sort"
)

type router [7]*deepTree

func (r *router) Lookup(method, path string) (Handler, *node) {
	dt := r[getMethodCode(method)]
	d := 0
	var st *staticData
	path = path[1:]
	l := len(path)
	if l == 0 {
		st = quickFind(dt.static, 0)
	} else {
		if path[l-1] == '/' {
			l--
			path = path[:l]
		}
		st = quickFind(dt.static, l)
	}
	if st != nil {
		h := st.macth(path)
		if h != nil {
			return h, nil
		}
	}
	ss := 0
	for s := 0; s < l; s++ {
		if path[ss] == '/' {
			dt.cache[d] = path[:ss]
			d++
			path = path[ss+1:]
			ss = 0
		} else {
			ss++
		}
	}
	dt.cache[d] = path
	d++
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
	n, isStatic := deep(path)
	if isStatic {
		if path[0] == '/' {
			path = path[1:]
		}
		l := len(path)
		if l > 0 && path[l-1] == '/' {
			l--
			path = path[:l]
		}
		sl := r[m].static
		nl := len(sl)
		for i := 0; i < nl; i++ {
			if sl[i].length == l {
				sl[i].add(path, h)
				return
			}
		}
		if cap(sl) == nl {
			v := make([]*staticData, nl+1, nl+1)
			copy(v, sl)
			sl = v
		}
		r[m].static = sl[:nl+1]
		sd := &staticData{
			length: l,
			paths:  make([]*staticHandler, 0, 0),
		}
		sd.add(path, h)
		r[m].static[nl] = sd
		return
	}
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
		sort.Slice(dt.static, func(i, j int) bool {
			return dt.static[i].length < dt.static[j].length
		})
		for _, v := range dt.static {
			v.startIndex = len(v.paths) / 2
			sort.Slice(v.paths, func(i, j int) bool {
				return sort.StringsAreSorted([]string{
					v.paths[i].path, v.paths[j].path,
				})
			})
		}
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
