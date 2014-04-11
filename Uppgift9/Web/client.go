package main

/*
*	This program will now quit when a request returns 200 as a status code.
*	Original author: Stefan Nilsson.
*	Modified by: Christopher Lillthors.
 */

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

var (
	serverList = []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	numberofGoroutines = len(serverList)
)

//an response object.
type Response struct {
	Body       string
	StatusCode int
}

func init() {
	//maximum power!
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// res := Get(server[0])
// res := Read(serverList[0], time.Second)

func main() {
	before := time.Now()
	res := MultiRead(serverList, time.Second)
	after := time.Now()
	fmt.Printf("Temperature %s", res.Body)
	fmt.Print("Time: ", after.Sub(before))
	// fmt.Println()
	// time.Sleep(500 * time.Millisecond)
}

// Get makes an HTTP Get request and returns an abbreviated response.
// Status code 200 means that the request was successful.
// The function returns &Response{"", 0} if the request fails
// and it blocks forever if the server doesn't respond.
func Get(url string) *Response {
	//does the actual GET request to the server.
	res, err := http.Get(url)
	if err != nil {
		//return empty response
		return &Response{}
	}
	// res.Body != nil when err == nil.
	//close or you will leak memory.
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	//fill the response with data.
	//warning. body is a []byte{}, therefore cast it to an string.
	return &Response{string(body), res.StatusCode}
}

/*
*	This function is now fixed. It will return an response from the server if the
*	response code is 200.
 */
func Read(url string, timeout time.Duration) (res *Response) {
	done := make(chan *Response)

	// create a function pointer.
	worker := func(ch chan *Response) {
		//a new response is created in Get() and as long there's an reference to the object, the GC will not take it away.
		resp := Get(url)
		//200 is the http code for success.
		//http.StatusOK = 200
		if resp.StatusCode != http.StatusOK {
			return
		}
		//send the response.
		ch <- resp
	}
	//launch goroutine to do the work.
	go worker(done)

	//listen on values on the channel.
	select {
	case resp := <-done:
		return resp
	//if the server is to slow...
	case <-time.After(timeout):
		return &Response{"Gateway timeout", http.StatusGatewayTimeout}
	}
}

// MultiRead makes an HTTP Get request to each url and returns
// the response of the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is
// 503 – Service unavailable.
func MultiRead(urls []string, timeout time.Duration) (res *Response) {
	done := make(chan *Response)

	//worker function. Does all the work.
	worker := func(url string, ch chan *Response) {
		resp := Get(url)
		for {
			//statusOK = 200
			if resp.StatusCode != http.StatusOK {
				//try again
				resp = Get(url)
				//if you don't want to try again, then just return.
				//You only have 3 goroutines to spare!
			} else {
				//send back data
				ch <- resp
				//die!
				return
			}
		}
	}

	//fire of one goroutine/url in array.
	for _, url := range urls {
		go worker(url, done)
	}

	//check channel.
	select {
	case resp := <-done:
		return resp
	case <-time.After(timeout):
		return &Response{"Gateway timeout", http.StatusGatewayTimeout}
	}
}
