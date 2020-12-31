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

func lookup(path string, h func(start, end int) bool) {
	l := len(path)
	if l == 0 {
		h(0, 0)
		return
	}
	if path == "/" {
		h(1, 1)
		return
	}
	start := 0
	if path[0] == '/' {
		start = 1
	}
	end := start
	for end < l {
		if path[end] == '/' {
			if h(start, end) {
				return
			}
			start = end + 1
		}
		end++
	}
	if path[l-1] != '/' {
		h(start, l)
	}
}
