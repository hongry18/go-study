package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e == nil {
		return
	}

	panic(e)
}

func main() {
	root := os.Getenv("GOPATH")
	if root == "" {
		panic("GO PATH NOT FOUND")
	}

	f1 := fmt.Sprintf("%s/tmp/dat1", root)

	d1 := []byte("hello/ngo/n")
	err := ioutil.WriteFile(f1, d1, 0644)
	check(err)

	f2 := fmt.Sprintf("%s/tmp/dat2", root)
	f, err := os.Create(f2)
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
