package main

import (
	"fmt"
	"time"
)

func main() {
	reqs := make(chan int, 5)
	for i := 1; i < 6; i++ {
		reqs <- i
	}
	close(reqs)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range reqs {
		<-limiter
		fmt.Println("reqeust", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyReq := make(chan int, 5)
	for i := 1; i < 6; i++ {
		burstyReq <- i
	}
	close(burstyReq)
	for req := range burstyReq {
		<-burstyLimiter
		fmt.Println("reqeust", req, time.Now())
	}
}
