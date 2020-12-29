package fastrouter

// todo:
// 1. 逆向查找
// 2. 多线程查找

type handle func() error

type level struct {
	path    []string
	parent  int
	handles map[string]handle
}

type router struct {
	deep    int
	maxDeep int
	levels  []*level
}

func (r *router) increase() {
	pre := cap(r.levels[r.deep].path)
	r.deep += 1
	if r.maxDeep < r.deep {
		r.maxDeep = router{}.maxDeep * 2
		v := make([]*level, r.deep, r.maxDeep)
		copy(v, r.levels)
		r.levels = v
	} else {
		r.levels = r.levels[:r.deep]
	}
	r.levels[r.deep] = newLevel(pre * 2)
}

func newLevel(cap int) *level {
	return &level{
		path: make([]string, 0, cap),
	}
}

func newRouter(deep, beginCap int) *router {
	r := &router{
		maxDeep: deep,
		deep:    1,
		levels:  make([]*level, 1, deep),
	}
	r.levels[0] = newLevel(beginCap)
	return r
}

func bind(path string, method string, h handle) {

}

func GetDeep(path string) int {
	d := 0
	l := len(path)
	if l > 0 {
		d=1
		l=l-1
		for i := 1; i < l; i++ {
			if path[i] == '/' {
				d++
			}
		}
	}
	return d
}

func Lookup(path string) {
	l := len(path)
	start := 0
	if path[0] == '/' {
		start = 1
	}
	end := start + 1
	for ; end < l; {
		if path[end] == '/' {
			println(path[start:end])
			start = end + 1
		}
		end++
	}
}
