package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach peach"))
	fmt.Println(r.FindStringIndex("peach peach"))
	fmt.Println(r.FindStringSubmatch("peach peach"))
	fmt.Println(r.FindStringSubmatchIndex("peach peach"))
	fmt.Println(r.FindAllString("peach peach peach", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach peach peach", -1))
	fmt.Println(r.FindAllString("peach peach peach", 2))
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
