package main

import (
	"fmt"
	"runtime"
)

const (
	count = 11
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	//create channel
	ch := make(chan int)
	done := make(chan bool)

	//closure
	printNumbers := func(ch1 chan int, done chan bool) {
		//reads until close(ch1)
		for n := range ch1 {
			fmt.Println(n)
		}
		done <- true
	}
	//fire of goroutine
	go printNumbers(ch, done)

	for i := 1; i <= count; i++ {
		ch <- i
	}
	//when all the numbers are sent, be sure to close the channel
	close(ch)

	//blocking until goroutine is finished.
	if <-done == true {
		close(done)
	}
}
