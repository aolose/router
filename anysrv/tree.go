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
	ns := r.levels[r.deep-1].nodes
	l := len(ns)
	for i := 0; i < l; i++ {
		for p := ns[i].parent; p != nil; p = p.parent {
			if p.start == -1 {
				p.start = i
			}
		}
	}
}

type deepTree struct {
	max   int
	trees []*tree
	cache []string
}
