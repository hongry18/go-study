package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func sample1() {
	fmt.Println("")
	fmt.Println("# sample1 start")
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("wait end")
	fmt.Println("ops: ", ops)
}

func sample2() {
	fmt.Println("")
	fmt.Println("# sample2 start")
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//time.Sleep(time.Second)
	fmt.Println("after a few second")
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
}

func main() {
	sample1()
	sample2()
}
