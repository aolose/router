package anysrv

import (
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
