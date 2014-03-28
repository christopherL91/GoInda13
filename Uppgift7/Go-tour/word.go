package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	list := make(map[string]int)
	for _, element := range strings.Fields(s) {
		//very clever trick!
		list[element]++
	}
	return list
}

func main() {
	wc.Test(WordCount)
}
