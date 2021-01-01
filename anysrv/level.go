package anysrv

import (
	"sort"
	"unsafe"
)

type level struct {
	nodes []*node
}

func (v *level) sort() {
	sort.SliceStable(v.nodes, func(i, j int) bool {
		p0 := uintptr(unsafe.Pointer(v.nodes[i].parent))
		p1 := uintptr(unsafe.Pointer(v.nodes[j].parent))
		return p0 < p1
	})
}

func (v *level) String() string {
	s := "\n ["
	for i := 0; i < len(v.nodes); i++ {
		s = s + v.nodes[i].String() + " "
	}
	return s + "\n ]"
}

func newLevel(cap int) *level {
	return &level{
		nodes: make([]*node, 0, cap),
	}
}

func (v *level) bind(p *node, path string) *node {
	l := len(v.nodes)
	for i := 0; i < l; i++ {
		n := v.nodes[i]
		if p != nil && n.parent != p {
			continue
		}
		if path != n.path {
			continue
		}
		return n
	}
	c := cap(v.nodes)
	if c == l {
		ns := make([]*node, c, c*2)
		copy(ns, v.nodes)
		v.nodes = ns
	}
	v.nodes = v.nodes[:l+1]
	n := &node{path: path, parent: p}
	v.nodes[l] = n
	return n
}
