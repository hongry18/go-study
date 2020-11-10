package main

import "fmt"

type rect struct {
	w, h int
}

func (r *rect) area() int {
	return r.w * r.h
}

func (r rect) perim() int {
	return (2 * r.w) + (2 * r.h)
}

// call by value test
func (r rect) add1(a int) {
	r.w += a
	r.h += a
}

// call by reference test
func (r *rect) add2(a int) {
	r.w += a
	r.h += a
}

func main() {
	var r rect = rect{10, 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	r.add1(1)
	r.add1(2)
	fmt.Println("added: ", r)

	r.add2(1)
	r.add2(2)
	fmt.Println("added: ", r)
}
