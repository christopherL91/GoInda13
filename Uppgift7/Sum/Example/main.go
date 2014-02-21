package main

import (
	"fmt"
	"github.com/christopherL91/GoInda13/Uppgift7/Sum"
)

func main() {
	var sum int = 0
	buffer := []int{3, 4, 5, 3}
	length := len(buffer)
	ch := make(chan int)

	go Sum.Add(buffer[:length/2], ch)
	go Sum.Add(buffer[length/2:], ch)

	sum += <-ch
	sum += <-ch
	fmt.Println(sum)
}
