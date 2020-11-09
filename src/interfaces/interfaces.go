package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.w * r.h
}

func (r rect) perim() float64 {
	return (2 * r.w) + (2 * r.h)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func measure(g geometry) {
	fmt.Println("geometry: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	var r rect = rect{3, 4}
	var c circle = circle{5}

	measure(r)
	measure(c)
}
