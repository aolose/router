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

func match(route, real string) bool {
	if route == real {
		return true
	}
	n0 := len(route)
	if n0 == 0 {
		return false
	}
	if route[0] == ':' {
		return true
	}
	if route == "*" {
		return true
	}
	n1 := len(real)
	if n1 == 0 {
		return false
	}
	if n1 >= n0-1 {
		if route[0] == '*' && (route[1:] == real[n1-n0+1:]) || (route[n0-1] == '*' && route[:n0-1] == real[:n0-1]) {
			return true
		}
	}
	return false
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
