package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question         string `json:"question"`
	Answer           string `json:"answer"`
	CorrectIndicator bool   `json:"correctIndicator"`
}

const time_in_seconds = 8

func main() {
	// Parse flags
	var userCsvFile = flag.String("csvFile", "problems.csv", "Defaults to problems.csv, pass flag if you want to use your own csv file.")
	var timeInSeconds = flag.Int("timeInSeconds", 30, "Seconds that user has to take the quiz")

	flag.Parse()

	// Read in CSV file
	csvFile, _ := os.Open(*userCsvFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Declare correct or incorrect answer counter
	correctAnswers := 0
	incorrectAnswers := 0

	// Set timer
	timer := time.NewTimer(time.Duration(*timeInSeconds) * time.Second)

	// Initiate
	go func() {
		<-timer.C
		fmt.Printf("\n\n\nTimes up! Your %d second timer finished.", *timeInSeconds)
		fmt.Println("\nCorrect answers:", correctAnswers)
		fmt.Println("Incorrect answers:", incorrectAnswers)
		os.Exit(0)
	}()

	// Iterate lines in CSV file.
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		userInputReader := bufio.NewReader(os.Stdin)
		fmt.Printf("What is %s: ", line[0])
		userInput, _ := userInputReader.ReadString('\n')
		if strings.TrimSpace(userInput) == line[1] {
			fmt.Println("Correct\n")
			// Increment correctAnswers variable
			correctAnswers += 1

		} else {
			fmt.Printf("Incorrect. Answer is %s. You typed %s\n", line[1], userInput)
			// Increment incorrectAnswers
			incorrectAnswers += 1
		}
	}
	stop := timer.Stop()
	if stop {
		fmt.Println("You ran out of time")
	}
	fmt.Println("\nCorrect answers:", correctAnswers)
	fmt.Println("Incorrect answers:", incorrectAnswers)
}

// questionJson, _ := json.Marshal(question)
// fmt.Println(string(questionJson))
