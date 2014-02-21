/*
Author: Christopher Lillthors
Different approaches for printing out "Hello World" with several goroutines

*/

package main

import (
	"fmt"
	"sync"
	"time"
	"log"
)
------------------------------------------------------------
const (
	numberOfHello = 1000
)

func sayHello(ch chan string, wg *sync.WaitGroup) {
	for {
		select {
		case input := <-ch:
			fmt.Println(input)
			wg.Done()
		}
	}
}

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	wg.Add(numberOfHello)
	for i := 1; i <= numberOfHello; i++ {
		go sayHello(ch, wg)
		ch <- "Hello world "
	}
	wg.Wait()
}
-----------------------------------------------------------

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	done := make(chan bool)
	go sayHello(ch,done)
	ch <- "Hello World!"
	//wait for goroutine to finish
	time.Sleep(0.5 * time.Second)
}
-----------------------------------------------------------


func sayHello(ch chan string,done chan bool) {
	fmt.Println(<-ch)
	done <- true
}

func main() {
	ch := make(chan string)
	go sayHello(ch)
	ch <- "Hello World!"
	if done == true {
		return
	} else {
		log.Fatal("Something went wrong.")
	}
}

-----------------------------------------------------------