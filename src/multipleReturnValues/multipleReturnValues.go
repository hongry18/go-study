package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func main() {
	r1, r2 := vals()

	fmt.Println(r1, r2)

	_, r3 := vals()
	fmt.Println(r3)

	var r4, r5 int = vals()
	fmt.Println(r4, r5)
}
