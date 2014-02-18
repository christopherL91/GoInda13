package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

type serverResponse struct {
	Name    string
	Id      string `json:"id"`
	Kind    string `json:"kind"`
	LongUrl string `json:"longUrl"`
}

const (
	serverURL = "https://www.googleapis.com/urlshortener/v1/url"
)

func request(longUrl string, ws *sync.WaitGroup) {
	data := new(serverResponse)
	data.Name = longUrl
	body := bytes.NewBufferString(fmt.Sprintf(`{"longUrl":"%s"}`, longUrl))
	client := new(http.Client)
	req, _ := http.NewRequest("POST", serverURL, body)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Could not do request to server")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("No response from server. URL: " + longUrl)
	}
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Could not read response")
	}
	err = json.Unmarshal(resbody, data)
	if err != nil {
		log.Fatal("could not unmarshal json")
	}
	fmt.Println(data.Name + " -> " + data.Id)
	//we are done here...
	ws.Done()
}

func init() {
	cores := flag.Int("cores", runtime.NumCPU(), "Number of cores used, Default is the number of cores in this machine.")
	runtime.GOMAXPROCS(*cores)
	flag.Parse()
}

func main() {
	ws := new(sync.WaitGroup)
	defer ws.Wait()
	for _, url := range flag.Args() {
		go request(url, ws)
		ws.Add(1)
	}
}
