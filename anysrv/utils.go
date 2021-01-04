package anysrv

func deep(path string) (int, bool) {
	isStatic := true
	d := 1
	l := len(path)
	if l > 0 {
		l = l - 1
		for i := 0; i < l; i++ {
			p := path[i]
			if p == '/' {
				d++
			}
			if isStatic {
				if p == ':' || p == '*' {
					isStatic = false
				}
			}
		}
		if isStatic {
			return d, path[l] != '*'
		}
	}
	return d, isStatic
}

func match(n *node, real string) bool {
	switch n.cate {
	//:*
	case 0, 1:
		return true
	//aa aa
	case 2:
		return n.path == real
	//abc*
	case 3:
		l := len(n.path) - 1
		if l > len(real) {
			return false
		}
		p := n.path
		for i := 0; i < l-1; i++ {
			if p[i] != real[i] {
				return false
			}
		}
		return true
	//*abc
	case 4:
		p := n.path
		l := len(p) - 1
		if l > len(real) {
			return false
		}
		r := len(real)
		for i := 0; i < l; i++ {
			if p[i+1] != real[r-l+i] {
				return false
			}
		}
		return true
	default:
		return false
	}
}
