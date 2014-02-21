package main

import (
	"github.com/christopherL91/GoInda13/Uppgift7/RemindMe"
	"time"
)

func main() {
	clock := 2 * time.Second
	go RemindMe.Remind("Good morning", clock)
	select {}
}
