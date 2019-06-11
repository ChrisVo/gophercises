package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Question struct {
	Question         string `json:"question"`
	Answer           string `json:"answer"`
	CorrectIndicator bool   `json:"correctIndicator"`
}

func main() {
	// Read in CSV file
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var questionData []Question
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
			fmt.Println("Correct")
			questionData = append(questionData, Question{
				Question:         line[0],
				Answer:           line[1],
				CorrectIndicator: true,
			})
		} else {
			fmt.Printf("Incorrect. Answer is %s. You typed %s", line[1], userInput)
			questionData = append(questionData, Question{
				Question:         line[0],
				Answer:           line[1],
				CorrectIndicator: false,
			})
		}
		// fmt.Println(line[0], line[1])
	}
}

// questionJson, _ := json.Marshal(question)
// fmt.Println(string(questionJson))
