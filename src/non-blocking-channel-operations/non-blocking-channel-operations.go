package main

import (
	"fmt"
)

func main() {
	//messages := make(chan string)
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "hi"
		messages <- "hi"
	}()

	fmt.Println(<-messages)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	var msg string = "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("sent message", msg)
	case sig := <-signals:
		fmt.Println("sent signal", sig)
	default:
		fmt.Println("no activity")
	}

}
