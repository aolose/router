package anysrv

var share0 = [128]int{1}
var share1 = [128]int{1}

//   /a/b
func parseReqPath(path string) *reqPath {
	l := len(path) - 1
	p := 0
	s := share0
	e := share1
	s[0] = 1
	e[0] = 1
	for i := 2; i < l; i++ {
		if path[i] == '/' {
			s[p+1] = i + 1
			e[p] = i
			p++
		}
	}
	if l > 0 && path[l] == '/' {
		e[p] = l
	} else {
		e[p] = l + 1
	}
	return &reqPath{
		length: l,
		deep:   p,
		start:  &s,
		end:    &e,
	}
}
