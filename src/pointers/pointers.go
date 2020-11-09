package main

import "fmt"

func zeroval(v int) {
	v = 0
}

func zeroptr(v *int) {
	*v = 0
}

func main() {
	var i int = 5
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("call zeroval: ", i)

	zeroptr(&i)
	fmt.Println("call zeroval: ", i)

	fmt.Println("call pointer: ", &i)
}