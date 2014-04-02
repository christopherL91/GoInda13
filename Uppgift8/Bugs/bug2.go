/*
	This program is modified for testing purposes.
*/

package Bugs

import (
	"runtime"
	"sort"
)

const (
	//the number of numbers in one slice.
	count = 11
)

func init() {
	//maximum power!
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func WriteNumbers() []int {
	//create channels
	ch := make(chan int)
	done := make(chan []int)

	//closure
	printNumbers := func(ch1 chan int, done chan []int) {
		buffer := []int{}
		//reads until close(ch1)
		for n := range ch1 {
			buffer = append(buffer, n)
		}
		done <- buffer
	}
	//fire dah lazer!
	go printNumbers(ch, done)

	//send all the numbers!11!!!!!!1111!!!!!!
	for i := 1; i <= count; i++ {
		ch <- i
	}
	//when all the numbers are sent, be sure to close the channel
	close(ch)

	//blocking until goroutine is finished.
	return <-done
}

//for testing
func Reverse(input []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
}
