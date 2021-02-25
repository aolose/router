package anysrv

import (
	"fmt"
	"testing"
)

func notEqual(k, p string, a, b interface{}, t *testing.T) bool {
	if a != b {
		t.Errorf(
			"%s error: %s should %d,but got %d",
			p, k, a, b,
		)
		return true
	}
	return false
}

func TestParseReqPath(t *testing.T) {
	for _, p := range []struct {
		p string
		r reqPath
	}{
		{"/", reqPath{0, 0, []int{0}, []int{0}}},
		{"/a", reqPath{1, 0, []int{0}, []int{1}}},
		{"/b", reqPath{2, 0, []int{0}, []int{2}}},
		{"/a/b", reqPath{3, 2, []int{0, 2}, []int{2}}},
	} {
		p0 := parseReqPath(p.p)
		if notEqual(p.p, "deep", p.r.deep, p0.deep, t) {
			return
		}
		if notEqual(p.p, "length", p.r.length, p0.length, t) {
			return
		}
		s0 := fmt.Sprintf("%v", p0.start)
		s1 := fmt.Sprintf("%v", p.r.start)
		if notEqual(p.p, "start", s1, s0, t) {
			return
		}
		s3 := fmt.Sprintf("%v", p0.end)
		s4 := fmt.Sprintf("%v", p.r.end)
		if notEqual(p.p, "end", s4, s3, t) {
			return
		}
	}
}
