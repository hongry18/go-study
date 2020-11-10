package main

import (
	"fmt"
	"time"
)

func customer(req chan chan string) {
	rc := make(chan string)
	req <- rc

	fmt.Println(<-rc)
	fmt.Println(<-rc)
}

func productor(req chan chan string) {
	res := <-req
	res <- "buffered"
	res <- "channel"
}

func main() {
	msgCh := make(chan string, 2)
	msgCh <- "buffered"
	msgCh <- "channel"

	fmt.Println(<-msgCh)
	fmt.Println(<-msgCh)

	// chan chan
	rc := make(chan chan string)
	go customer(rc)
	go productor(rc)

	time.Sleep(time.Second)
}
