package fastrouter

import "testing"

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
	for i, p := range []struct {
		path    string
		method  int
		handler handle
	}{
		{"", 0, nil},
		{"a/b", 0, nil},
		{"a/b/", 0, nil},
		{"/a/b/c", 0, nil},
		{"/a/c", 0, nil},
	} {
		d := deep(p.path)
		if d > dp {
			dp = d
		}
		r.bind(p.method, p.path, nil)
		if r.deep != dp {
			t.Errorf("%v", r)
			t.Errorf("%d. router deep error at  %s, should be %d but got %d", i, p.path, d, r.deep)
		}
	}
}
