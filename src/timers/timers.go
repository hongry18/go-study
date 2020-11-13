package main

import (
	"fmt"
	"time"
)

func main() {
	var t1 = time.NewTimer(time.Second * 2)

	<-t1.C
	fmt.Println("timer 1 expired")

	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		fmt.Println("timer 2 expired")
	}()

	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
}
