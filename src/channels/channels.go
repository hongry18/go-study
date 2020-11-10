package main

import (
	"fmt"
	"time"
)

// chan chan receive
func customer(reqCh chan chan string) {
	resCh := make(chan string)
	reqCh <- resCh
	res := <-resCh
	fmt.Println("response: ", res)
}

// chan chan give
func productor(reqCh chan chan string) {
	resCh := <-reqCh
	resCh <- "wassup!"
}

func main() {
	// chan
	messages := make(chan string)

	go func() { messages <- "Ping" }()

	msg := <-messages
	fmt.Println(msg)

	// chan chan
	reqCh := make(chan chan string)
	go customer(reqCh)
	go productor(reqCh)

	time.Sleep(time.Second)
}
