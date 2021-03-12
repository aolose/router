package main

import (
	"fmt"
	_ "net/http/pprof"
	"sort"
)

func main() {
	a := []int{2, 3, 4, 6, 9, 8, 1}
	sort.Ints(a[1:])
	fmt.Printf("%v", a)
}
