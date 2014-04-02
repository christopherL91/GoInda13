/*
Pythia, oraklet i Delphi

 Oracle
Filen oracle.go innehåller ett kodskelett till ett orakelprogram som besvarar frågor.

Gör klart Oracle-metoden. Du får inte ändra i main-metoden och du får inte heller ändra metodsignaturerna. Observera att svaren inte ska komma direkt, utan med fördröjning. Glöm inte heller att oraklet ska skriva ut meddelanden även om det inte kommer några frågor. Du får gärna dela upp din lösning på flera metoder.
Ditt program ska innehålla två stycken kanaler: en kanal för frågor samt en kanal för svar och förutsägelser. I Oracle-metoden ska du starta tre stycken permanenta gorutiner:

En gorutin som tar emot alla frågor och för varje inkommande fråga skapar en separat gorutin som besvarar frågan. OK
En gorutin som genererar förutsägelser.
En gorutin som tar emot alla svar och förutsägelser och skriver ut dem på stdout.
Oracle-metoden är den viktigaste delen av uppgiften. Om du vill får du också förbättra svarsalgoritmen. Även här får gärna dela upp algoritmen på flera metoder. Här är några tips:

Paketen strings och regexp kan vara användbara.
Programmet kan verka mera mänskligt om oraklet skriver ut sina svar en bokstav i taget.
Ta en titt på ELIZA, det första programmet av det här slaget.
*/

// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star       = "Pythia"
	venue      = "Delphi"
	prompt     = "> "
	buffersize = 10
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string, buffersize)
	printch := make(chan string, buffersize)
	go generateProphecy2(questions, printch)
	go printOut(printch)

	return questions
}

func generateProphecy1(question string, ch1 chan<- string) {
	go prophecy(question, ch1)
}

func generateProphecy2(ch1 <-chan string, ch2 chan<- string) {
	for {
		// Keep them waiting. Pythia, the original oracle at Delphi,
		// only gave prophecies on the seventh day of each month.
		time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)
		select {
		case question := <-ch1:
			go generateProphecy1(question, ch2)
		default:
			go generateProphecy1("", ch2)
		}
	}
}

func printOut(ch <-chan string) {
	for {
		select {
		case in := <-ch:
			fmt.Print("\r")
			for _, val := range in {
				time.Sleep(time.Duration(30+rand.Intn(20)) * time.Millisecond)
				fmt.Print(string(val))
			}
			fmt.Print("\n", prompt)
		}
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {

	if question == "" {
		gibberish := []string{
			"mohaha you fucking loser!",
			"can't believe you're so fucking dumb!",
			"Leonidas: THIS IS SPARTA!!!!! Pythia: No! This is ATHENS!!!!! Dumb spartans...",
			"Can't believe you falled for that!",
		}
		answer <- gibberish[rand.Intn(len(gibberish))]
		return
	}

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"There will be blood",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
