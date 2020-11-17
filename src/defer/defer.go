package main

import (
	"fmt"
	"os"
)

var root string = os.Getenv("GOPATH")

func main() {

	f := createFile("/tmp/sample.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("create File")
	f, err := os.Create(root + path)
	if err != nil {
		panic(err)
	}

	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
