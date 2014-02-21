package RemindMe

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
			fmt.Println("Time is now " + strconv.Itoa(clock.Hour()) + ":" + strconv.Itoa(clock.Minute()) + "." + strconv.Itoa(clock.Second()) + " " + text)
		default:
			//nothing to do.
		}
	}
}
