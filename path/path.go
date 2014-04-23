package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/christopherL91/GoInda13/path/graph"
	"os"
	"strconv"
	"strings"
)

var (
	printStack bool
	fileName   string
	from       int
	to         int
)

func onError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func reverse(input []int) {
	length := len(input)
	for i := 0; i < length/2; i++ {
		input[i], input[length-1-i] = input[length-1-i], input[i]
	}
}

func nextNum(reader *bufio.Reader, delim byte) (int, error) {
	raw, err := reader.ReadString(delim)
	if err != nil {
		return 0, err
	}

	num, err := stringToNum(raw)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func stringToNum(input string) (int, error) {
	rawDigit := strings.Trim(input, " \n")
	digit, err := strconv.Atoi(rawDigit)
	if err != nil {
		return 0, errors.New("Could not convert number")
	}
	return digit, nil
}

func init() {
	flag.StringVar(&fileName, "filename", "input.txt", "Name of file")
	flag.BoolVar(&printStack, "stack", false, "Print the stack")
	flag.Parse()
}

func main() {
	//contains all the arguments from command line
	args := flag.Args()
	//check number of arguments
	if len(args) != 2 {
		onError("Check number of arguments")
	}

	//convert string to number
	var input int
	input, err := stringToNum(args[0])
	from = input

	//convert string to number
	input, err = stringToNum(args[1])
	to = input

	if err != nil {
		onError("Could not convert values from command line")
	}

	rawFile, err := os.Open(fileName)
	if err != nil {
		onError(err.Error())
	}

	reader := bufio.NewReader(rawFile)
	rawSize, err := reader.ReadString('\n')
	if err != nil {
		onError(err.Error())
	}

	size, err := stringToNum(rawSize)
	if err != nil {
		onError(err.Error())
	}
	//since hash is not yet implemented.
	g := graph.NewMatrix(size)
	reader.ReadString('\n') //jump one line

	var v, w, cost int
	for {
		v, err = nextNum(reader, ' ')
		if err != nil {
			break
		}
		w, err = nextNum(reader, ' ')
		if err != nil {
			break
		}
		cost, err = nextNum(reader, '\n')
		if err != nil {
			break
		}
		g.AddLabel(v, w, cost)
	}

	// stack := make([]int, g.NumEdges())
	// visited := make([]bool, g.NumVertices())
	// count := 0
	// found := false

	// graph.BFS(g, from, visited, func(w int) {
	// 	if !found {
	// 		stack[count] = w
	// 		count++
	// 	}
	// 	if w == to {
	// 		stack = stack[0:count]
	// 		found = true
	// 	}
	// })

	visited := make([]bool, g.NumEdges())
	stack := make([]int, g.NumEdges())

	for index, _ := range stack {
		stack[index] = -1
	}

	graph.BFS(g, from, visited, func(prev, w int) {
		stack[w] = prev
	})

	index := to
	found := false
	var previous, totalCost int
	totalCost = 0
	var path []int
	path = append(path, to)
	for stack[index] != -1 {
		previous = stack[index]
		cost = g.Label(previous, index).(int)
		totalCost += cost
		index = previous
		path = append(path, index)
		if index == from {
			found = true
			break
		}
	}

	if found {
		if printStack {
			fmt.Println(stack)
		}
		reverse(path)
		fmt.Printf("Path:%v Cost:%d\n", path, totalCost)
	} else {
		onError("Could not find a way")
	}
}

// [6 5 1 0] //path
// [-1 0 1 1 0 1 5 4 -1 -1 -1] //stack
//	 0 1 2 3 4 5 6 7  8  9 10
/*
	1: path = [6]
	2: path = []
*/
