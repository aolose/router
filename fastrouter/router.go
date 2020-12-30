package fastrouter

import (
	"sort"
)

// todo:
// 1. 逆向查找
// 2. 多线程查找

var methods = map[string]int{
	"GET":     0,
	"POST":    1,
	"PUT":     2,
	"HEAD":    3,
	"DELETE":  4,
	"PATCH":   5,
	"OPTIONS": 6,
}

type handle func() error

type node struct {
	path   string
	parent *node
	handle [7]handle
	start  []int
}

type level struct {
	nodes []*node
}

func (v *level) sort() {
	sort.SliceStable(v.nodes, func(i, j int) bool {
		return sort.StringsAreSorted([]string{v.nodes[i].path, v.nodes[j].path})
	})
}

type router struct {
	deep    int
	maxDeep int
	levels  []*level
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
func (n *node) String() string {
	return n.path
}
func (l *level) String() string {
	s := "[ "
	for i := 0; i < len(l.nodes); i++ {
		s = s + l.nodes[i].String() + " "
	}
	return s + "]"
}
func (r *router) String() string {
	s := "{ "
	for i := 0; i < len(r.levels); i++ {
		s = s + r.levels[i].String() + " "
	}
	return s + "}\n"
}

func newLevel(cap int) *level {
	return &level{
		nodes: make([]*node, 0, cap),
	}
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

func (r *router) bind(path string, method string, h handle) {
	m, ok := methods[method]
	dp := 0
	var pr *node
	if ok {
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
}

func deep(path string) int {
	d := 1
	l := len(path)
	if l > 0 {
		l = l - 1
		for i := 1; i < l; i++ {
			if path[i] == '/' {
				d++
			}
		}
	}
	return d
}

func lookup(path string, h func(start, end int) bool) {
	l := len(path)
	if l == 0 {
		h(0, 0)
		return
	}
	start := 0
	if path[0] == '/' {
		start = 1
	}
	end := start
	for end < l {
		if path[end] == '/' {
			if h(start, end) {
				return
			}
			start = end + 1
		}
		end++
	}
	if path[l-1] != '/' {
		h(start, l)
	}
}
