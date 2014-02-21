package RemindMe

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Remind(text string, paus time.Duration) {
	if paus == 0 {
		log.Fatal("Time must be greater than 0")
		os.Exit(1)
	}
	ticker := time.NewTicker(paus)
	for {
		select {
		case clock := <-ticker.C:
			fmt.Println("Time is now " + clock.Format("15:04:05") + " " + text)
		default:
			//nothing to do.
		}
	}
}
