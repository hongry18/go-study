package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()

	var msg string
	for i := 0; i < 2; i++ {
		select {
		case msg = <-c1:
			fmt.Println("received: ", msg)
		case msg = <-c2:
			fmt.Println("received: ", msg)
		}
	}
}
