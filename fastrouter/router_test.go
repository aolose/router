package fastrouter

import "testing"

func TestGetDeep(t *testing.T) {
	for _, p := range []struct {
		path string
		deep int
	}{
		{"", 0},
		{"/", 1},
		{"a", 1},
		{"/a", 1},
		{"/a/", 1},
		{"a/b", 2},
		{"/a/b", 2},
		{"/a/b/", 2},
	} {
		if v := GetDeep(p.path); v != p.deep {
			t.Errorf("GetDeep(%s) should equal %d, but got %d", p.path, p.deep, v)
		}
	}
}
