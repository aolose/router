package anysrv

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

func match(n *node, real string) bool {
	switch n.cate {
	//:*
	case 0, 1:
		return true
		return real == "" || real[0] != 0
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
		for i := 1; i < l; i++ {
			if p[i] != real[r-l+i] {
				return false
			}
		}
		return true
	case 5:
		return true
	default:
		return false
	}
}

func lookup(path string, h func(start, end int) bool) bool {
	static := true
	l := len(path)
	if l == 0 {
		h(0, 0)
		return static
	}
	if path == "/" {
		h(1, 1)
		return static
	}
	start := 0
	if path[0] == '/' {
		start = 1
		if static && path[1] == ':' {
			static = false
		}
	}
	end := start
	for end < l {
		e := path[end]
		if static && e == '*' {
			static = false
		} else if e == '/' {
			if h(start, end) {
				return static
			}
			start = end + 1
		}
		end++
	}
	if path[l-1] != '/' {
		h(start, l)
	}
	return static
}
