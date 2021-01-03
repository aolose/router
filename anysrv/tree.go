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

type deepTree struct {
	max   int
	trees []*tree
	cache []string
}
