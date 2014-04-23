package GraphPath

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/christopherL91/GoInda13/path/graph"
	"os"
	"strconv"
	"strings"
)

//convenience function.
func OnError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func Reverse(input []int) {
	length := len(input)
	for i := 0; i < length/2; i++ {
		input[i], input[length-1-i] = input[length-1-i], input[i]
	}
}

func NextNum(reader *bufio.Reader, delim byte) (int, error) {
	raw, err := reader.ReadString(delim)
	if err != nil {
		return 0, err
	}

	num, err := StringToNum(raw)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func StringToNum(input string) (int, error) {
	rawDigit := strings.Trim(input, " \n")
	digit, err := strconv.Atoi(rawDigit)
	if err != nil {
		return 0, errors.New("Could not convert number")
	}
	return digit, nil
}

func FindPath(to, from int, fileName string, printStack bool) ([]int, []int, int, error) {
	rawFile, err := os.Open(fileName)
	if err != nil {
		OnError(err.Error())
	}

	reader := bufio.NewReader(rawFile)
	rawSize, err := reader.ReadString('\n')
	if err != nil {
		OnError(err.Error())
	}

	size, err := StringToNum(rawSize)
	if err != nil {
		OnError(err.Error())
	}
	//since hash is not yet implemented.
	g := graph.NewMatrix(size)
	reader.ReadString('\n') //jump one line

	var v, w, cost int
	for {
		v, err = NextNum(reader, ' ')
		if err != nil {
			break
		}
		w, err = NextNum(reader, ' ')
		if err != nil {
			break
		}
		cost, err = NextNum(reader, '\n')
		if err != nil {
			break
		}
		g.AddLabel(v, w, cost)
	}

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
			//reverse the order of path.
			Reverse(path)
			break
		}
	}

	if found {
		if printStack {
			return path, stack, totalCost, nil
		} else {
			return path, nil, totalCost, nil
		}
	}
	//didn't find anything. Print error.
	return nil, nil, 0, errors.New("Could not find a way")
}
