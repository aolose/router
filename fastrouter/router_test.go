package fastrouter

import (
	"fmt"
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
		handler handle
	}{
		{"", 0, func() error { a = "h0"; return nil }},
		{"b/b", 0, func() error { a = "h1"; return nil }},
		{"b/c", 0, func() error { a = "h2"; return nil }},
		{"a/b", 0, func() error { a = "h3"; return nil }},
		{"/a/b/c", 0, func() error { a = "h4"; return nil }},
		{"/a/c", 0, func() error { a = "h5"; return nil }},
		{"/a/:c/e", 0, func() error { a = "h6"; return nil }},
		{"/a/:c/*f", 0, func() error { a = "h7"; return nil }},
		{"/a/:c/f*", 0, func() error { a = "h8"; return nil }},
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
	fmt.Printf("%v", r)
	if s[1] != 0 {
		t.Errorf("r[0][2][1] start should be %d,but got %d", 0, s[1])
	}
	if s[0] != 2 {
		t.Errorf("r[1][2][0] start should be %d,but got %d", 2, s[0])
	}
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
		h := r.Lookup("GET", p.u)
		if h == nil {
			t.Errorf("lookup %s should return a handle", p.u)
		} else {
			_ = h()
			if a != p.r {
				t.Errorf("lookup %s should change a to %s ,but got %s", p.u, p.r, a)
			}
		}

	}
}
