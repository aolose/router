package anysrv

import (
	"testing"
)

func TestDeep(t *testing.T) {
	for _, p := range []struct {
		path string
		deep int
	}{
		{"", 1},
		{"/", 1},
		{"a", 1},
		{"/a", 1},
		{"/a/", 1},
		{"a/b", 2},
		{"/a/b", 2},
		{"/a/b/", 2},
	} {
		if v := deep(p.path); v != p.deep {
			t.Errorf("deep(%s) should equal %d, but got %d", p.path, p.deep, v)
		}
	}
}

func TestLookup(t *testing.T) {
	for _, p := range []struct {
		path string
		r    string
	}{
		{"", ""},
		{"/", ""},
		{"a", "a"},
		{"/aa", "aa"},
		{"/a/", "a"},
		{"a/b", "ab"},
		{"/a/bb", "abb"},
		{"/a/b/", "ab"},
	} {
		i := ""
		if lookup(p.path, func(start, end int) bool {
			i += p.path[start:end]
			return false
		}); i != p.r {
			t.Errorf("fn(%s) should equal %s, but got %s", p.path, p.r, i)
		}
	}
}

func TestRouter(t *testing.T) {
	dp := 0
	r := newRouter(1, 5)
	a := ""
	for i, p := range []struct {
		path    string
		method  int
		handler Handler
	}{
		{"", 0, func(c Context) { a = "h0" }},
		{"b/b", 0, func(c Context) { a = "h1" }},
		{"b/c", 0, func(c Context) { a = "h2" }},
		{"a/b", 0, func(c Context) { a = "h3" }},
		{"/a/b/c", 0, func(c Context) { a = "h4" }},
		{"/a/c", 0, func(c Context) { a = "h5" }},
		{"/a/:c/e", 0, func(c Context) { a = "h6" }},
		{"/a/:c/*f", 0, func(c Context) { a = "h7" }},
		{"/a/:c/f*", 0, func(c Context) { a = "h8" }},
	} {
		d := deep(p.path)
		if d > dp {
			dp = d
		}
		r.bind(p.method, p.path, p.handler)
		if r.deep != dp {
			t.Errorf("%v", r)
			t.Errorf("%d. router deep error at  %s, should be %d but got %d", i, p.path, d, r.deep)
		}
	}
	r.initNodeStarts()
	s := r.levels[0].nodes[2].start
	if s[1] != 0 {
		t.Errorf("r[0][2][1] start should be %d,but got %d", 0, s[1])
	}
	if s[0] != 2 {
		t.Errorf("r[1][2][0] start should be %d,but got %d", 2, s[0])
	}
	for _, p := range []struct {
		u string
		r string
		m bool
	}{
		{"", "h0", false},
		{"b/b", "h1", false},
		{"/b/c", "h2", false},
		{"a/b", "h3", false},
		{"a/b/c", "h4", false},
		{"a/c", "h5", false},
		{"a/x/e", "h6", false},
		{"a/b/cf", "h7", false},
		{"a/v/f1", "h8", false},
		{"a/v/f3/ddd", "h8", true},
	} {
		h, _ := r.Lookup("GET", p.u)
		if h == nil {
			if !p.m {
				t.Errorf("lookup %s should return a Handler", p.u)
			}
		} else {
			if p.m {
				t.Errorf("lookup %s should not return  Handler", p.u)
			} else {
				h(nil)
				if a != p.r {
					t.Errorf("lookup %s should change a to %s ,but got %s", p.u, p.r, a)
				}
			}
		}

	}
}

func BenchmarkRouter_Router(b *testing.B) {
	r := newRouter(1, 5)
	for _, p := range []struct {
		path    string
		method  int
		handler Handler
	}{
		{"", 0, nil},
		{"b/b", 0, nil},
		{"b/c", 0, nil},
		{"a/b", 0, nil},
		{"/a/b/c", 0, nil},
		{"/a/c", 0, nil},
		{"/a/:c/e", 0, nil},
		{"/a/:c/*f", 0, nil},
		{"/a/:c/f*", 0, nil},
	} {
		r.bind(p.method, p.path, p.handler)
	}
	r.initNodeStarts()
	for _, p := range []struct {
		u string
		r string
	}{
		{"b/b", "h1"},
		{"/b/c", "h2"},
		{"a/b", "h3"},
		{"a/b/c", "h4"},
		{"a/c", "h5"},
		{"a/x/e", "h6"},
		{"a/b/cf", "h7"},
		{"a/v/f1", "h8"},
	} {
		h, _ := r.Lookup("GET", p.u)
		if h != nil {
			h(nil)
		}
	}
}
