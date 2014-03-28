package main

import (
	reminder "github.com/christopherL91/GoInda13/Uppgift7/ManyRemindMe"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	reminder.Alarm()
	select {}
}
