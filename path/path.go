package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/christopherL91/GoInda13/path/graph"
	// "github.com/christopherL91/GoInda13/path/stack"
	"os"
	"strconv"
	"strings"
)

var (
	fileName string
	from     int
	to       int
)

func onError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
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
	fromString, err := stringToNum(args[0])
	from = fromString

	//convert string to number
	fromString, err = stringToNum(args[1])
	to = fromString

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

	stack := make([]int, g.NumEdges())
	visited := make([]bool, g.NumVertices())
	count := 0
	found := false
	graph.BFS(g, from, visited, func(w int) {
		if !found {
			stack[count] = w
			count++
		}
		if w == to {
			stack = stack[0:count]
			found = true
		}
	})
	if found {
		fmt.Println(stack)
	}
}
