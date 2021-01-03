package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{
		"cba",
		"abc",
		"acb",
		"bca",
	}
	sort.Slice(a, func(i, j int) bool {
		return sort.StringsAreSorted([]string{
			a[i], a[j],
		})
	})
	print(3 / 2)

	fmt.Printf("%v", a)
}
