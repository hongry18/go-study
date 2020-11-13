package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong1(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func pong2(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message 1")
	pong1(pings, pongs)
	fmt.Println(<-pongs)

	ping(pings, "passed message 2")
	pong2(pings, pongs)
	fmt.Println(<-pongs)
}
