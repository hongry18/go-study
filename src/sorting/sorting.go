package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{1, 6, 2, 3}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)
}
