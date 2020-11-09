package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	var p person = person{name: name}
	p.age = 20
	return &p
}

func main() {
	fmt.Println("### usages ")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	var s person = person{name: "Sean", age: 25}
	fmt.Println(s.name)

	// reference copy
	// var sp *person = &s
	fmt.Println("### reference copy")
	sp := &s
	fmt.Println(s)
	fmt.Println(sp)

	sp.age = 50
	
	fmt.Println(s)
	fmt.Println(sp)

	fmt.Println("### value copy")
	// value copy
	np := s

	fmt.Println(s)
	fmt.Println(np)

	np.age = 3
	fmt.Println(s)
	fmt.Println(np)
}