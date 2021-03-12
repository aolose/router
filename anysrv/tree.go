package anysrv

type tree struct {
	nodes [][]*node
	right *node
}

func (t *tree) addNode(path string, h Handler, start, end, pm []int) {
	ps := make([]*param, len(pm))
	for i, d := range pm {
		ps[i] = &param{
			name: path[start[d]+1 : end[d]],
			deep: d,
		}
	}
	var r *node
	l := len(start)
Loop:
	for i := 0; i < l; i++ {
		s := start[i]
		e := end[i]
		p := path[s:e]
		if path[s] == ':' && i == 0 {
			if t.right == nil {
				r = &node{
					deep:  1,
					nodes: make([][]*node, 0, 0),
				}
				t.right = r
			}
			continue
		} else {
			if r == nil {
				l0 := e - s
				l1 := len(t.nodes)
				if l1 < l0 {
					tt := make([][]*node, l0, l0)
					copy(tt, t.nodes)
					t.nodes = tt
				} else {
					for _, rr := range t.nodes[l0-1] {
						if rr.path == p {
							r = rr
							continue Loop
						}
					}
				}
				r = &node{
					deep:  1,
					path:  p,
					nodes: make([][]*node, 0, 0),
				}
				tt := t.nodes[l0-1]
				if tt == nil {
					tt = make([]*node, 0, 0)
					t.nodes[l0-1] = tt
				}
				ln := len(tt) + 1
				tn := make([]*node, ln, ln)
				copy(tn, tt)
				tn[ln-1] = r
				t.nodes[l0-1] = tn
			} else {
				r = r.Add(p)
			}
		}
	}
	r.handler = h
	r.params = &ps
}
func (t *tree) ready() {
	readNs(t.nodes, t.right)
}
