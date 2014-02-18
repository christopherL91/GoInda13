package main

import (
	"flag"
	urlshort "github.com/christopherL91/GoInda13/Shorturl/Url"
	"runtime"
	"sync"
)

func init() {
	cores := flag.Int("cores", runtime.NumCPU(), "Number of cores used, Default is the number of cores in this machine.")
	runtime.GOMAXPROCS(*cores)
	flag.Parse()
}

func main() {
	ws := new(sync.WaitGroup)
	defer ws.Wait()
	for _, url := range flag.Args() {
		go urlshort.Request(url, ws)
		ws.Add(1)
	}
}
