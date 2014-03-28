package ManyRemindMe

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Alarm() {
	go remindeMe("Time to sleep", 24*time.Hour)
	go remindeMe("Time to work", 8*time.Hour)
	go remindeMe("Time to eat", 3*time.Hour)
}

//not exported.
func remindeMe(text string, duration time.Duration) {
	if duration <= 0 {
		log.Fatal("duration >= 0")
		os.Exit(1)
	}
	ticker := time.NewTicker(duration)
	for {
		select {
		case clock := <-ticker.C:
			fmt.Println("The time is " + clock.Format("15:04:05") + " " + text)
		default:
			//nothing to do.
		}
	}
}
