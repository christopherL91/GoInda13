package main

import (
	"flag"
	"fmt"
	"github.com/christopherL91/GoInda13/path/GraphPath"
	"os"
	"runtime/pprof"
)

var (
	printStack bool
	profile    bool
	fileName   string
	from       int
	to         int
)

func init() {
	flag.StringVar(&fileName, "filename", "input.txt", "Name of file")
	flag.BoolVar(&profile, "profile", false, "Profile this program")
	flag.BoolVar(&printStack, "stack", false, "Print the stack")
	flag.Parse()
}

func main() {
	//contains all the arguments from command line
	args := flag.Args()
	//check number of arguments
	if len(args) != 2 {
		GraphPath.OnError("Check number of arguments")
	}

	if profile {
		file, err := os.Create("profile.txt")
		if err != nil {
			GraphPath.OnError(err.Error())
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	//convert string to number
	var input int
	input, err := GraphPath.StringToNum(args[0])
	from = input

	//convert string to number
	input, err = GraphPath.StringToNum(args[1])
	to = input

	if err != nil {
		GraphPath.OnError("Could not convert values from command line")
	}

	path, stack, totalCost, err := GraphPath.FindPath(to, from, fileName, printStack)
	if err != nil {
		GraphPath.OnError(err.Error())
	}

	if stack != nil {
		fmt.Printf("Path: %v \nCost: %d \nStack: %v\n", path, totalCost, stack)
	} else {
		fmt.Printf("Path: %v \nCost: %d\n", path, totalCost)
	}
}
