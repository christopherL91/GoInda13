/*
*	This program is slitely changed. It will return an []int{} containing the numbers.
*	This is for testing purposes.
 */

package Bugs

import (
	"runtime"
	"sort"
)

const (
	//how many numbers are put into one slice.
	count = 11
)

func WriteNumbers() []int {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	//fire of goroutine
	go printNumbers(ch, done)

	for i := 1; i <= count; i++ {
		ch <- i
	}
	//when all the numbers are sent, be sure to close the channel
	close(ch)

	//blocking until goroutine is finished.
	slice := <-done
	return slice
}

//for testing
func Reverse(input []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
}
