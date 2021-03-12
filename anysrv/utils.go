package anysrv

var mkA = []int{1}
var mkB = []int{1}

func mark(a *[]int, b, idx int) {
	if idx > len(*a)-1 {
		c := make([]int, idx+16, idx+16)
		copy(c, *a)
		*a = c
	}
	(*a)[idx] = b
}

//   /a/b
func parseReqPath(path string) {
	l := len(path) - 1
	p := 0
	mkA[0] = 1
	mkB[0] = 1
	for i := 2; i < l; i++ {
		if path[i] == '/' {
			mark(&mkA, i+1, p+1)
			mark(&mkB, i, p)
			p++
		}
	}
	if l > 0 && path[l] == '/' {
		mark(&mkB, l, p)
	} else {
		mark(&mkB, l+1, p)
	}
}
