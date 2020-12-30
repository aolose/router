package fastrouter

import (
	"fmt"
	"sort"
	"unsafe"
)

// todo:
// 1. 逆向查找
// 2. 多线程查找

const (
	GET = iota
	POST
	PUT
	HEAD
	DELETE
	PATCH
	OPTIONS
)

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

func (n *node) match(ps []string) (bool, int) {
	p := n
	l := len(ps)
	d := -1
	var p0, p1 string
	var n0, n1 int
	for i := l - 1; i > -1 && p != nil; i-- {
		p0 = p.path
		p1 = ps[i]
		n0 = len(p0)
		n1 = len(p1)
		if p0[0] == ':' || p0 == "*" || p0 == p1 ||
			(n1 >= n0 &&
				(p0[0] == '*' && (p0[1:] == p1[n1-n0+1:]) ||
					(p0[n0-1] == '*' && p0[:n0-1] == p1[:n0-1]))) {
			p = p.parent
			d++
		} else {
			if d == -1 || len(p.start) == 0 {
				return false, -1
			}
			return false, p.start[d]
		}
	}
	return true, -1
}

func (v *level) sort() {
	sort.SliceStable(v.nodes, func(i, j int) bool {
		p0 := uintptr(unsafe.Pointer(v.nodes[i].parent))
		p1 := uintptr(unsafe.Pointer(v.nodes[j].parent))
		return p0 < p1
	})
}

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
func (n *node) String() string {
	return fmt.Sprintf("(%s) - %v, ", n.path, n.start)
}
func (v *level) String() string {
	s := "[ "
	for i := 0; i < len(v.nodes); i++ {
		s = s + v.nodes[i].String() + " "
	}
	return s + "]\n"
}
func (r *router) String() string {
	s := "{\n "
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
func getMethodCode(method string) int {
	switch method {
	case "GET", "":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "HEAD":
		return HEAD
	case "DELETE":
		return DELETE
	case "PATCH":
		return PATCH
	case "OPTIONS":
		return OPTIONS
	default:
		return GET
	}
}

func (r *router) initNodeStarts() {
	r.cache = make([]string, r.deep, r.deep)
	for i, l := range r.levels {
		if i > 0 {
			l.sort()
		}
		c := r.deep - i - 1
		for _, n := range l.nodes {
			if c > 0 {
				s := make([]int, c, c)
				n.start = s
				for t := 0; t < c; t++ {
					s[t] = -1
				}
			}
		}
	}
	for i := r.deep - 1; i > 0; i-- {
		lv := r.levels[i]
		var p *node
		for x, n := range lv.nodes {
			if n.parent != p {
				p = n.parent
				if len(n.start) > 0 {
					copy(p.start[1:], n.start)
				}
				p.start[0] = x
			}
		}
	}
}

func (r *router) Lookup(method, path string) handle {
	m := getMethodCode(method)
	d := 0
	lookup(path, func(start, end int) bool {
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
					return h
				}
				if i != -1 {
					e = i
				}
			}
		}
	}
	return nil
}

func (r *router) Any(path string, h handle) {
	for i := 0; i < 7; i++ {
		r.bind(i, path, h)
	}
}

func (r *router) Get(path string, h handle) {
	r.bind(GET, path, h)
}

func (r *router) Post(path string, h handle) {
	r.bind(POST, path, h)
}

func (r *router) Put(path string, h handle) {
	r.bind(PUT, path, h)
}

func (r *router) Head(path string, h handle) {
	r.bind(HEAD, path, h)
}

func (r *router) Delete(path string, h handle) {
	r.bind(DELETE, path, h)
}

func (r *router) Patch(path string, h handle) {
	r.bind(PATCH, path, h)
}

func (r *router) Options(path string, h handle) {
	r.bind(OPTIONS, path, h)
}

func (r *router) bind(m int, path string, h handle) {
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
