package Url

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type serverResponse struct {
	//the server adress
	Name string
	//the short adress from google
	Id string `json:"id"`
}

const (
	serverURL = "https://www.googleapis.com/urlshortener/v1/url"
)

func Request(longUrl string, ws *sync.WaitGroup) {
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
