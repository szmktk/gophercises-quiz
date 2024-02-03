package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	correctAnswers int
	wrongAnswers   int
)

func main() {
	fileName, _ := parseUserInput()
	data := readQuestions(fileName)
	mainLoop(data)
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, correctAnswers+wrongAnswers)
}

func parseUserInput() (string, int) {
	csvPath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return *csvPath, *timeLimit
}

func readQuestions(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	return data
}

func mainLoop(data [][]string) {
	for idx, line := range data {
		fmt.Printf("Problem #%d: %s = ", idx+1, line[0])
		var answer string
		fmt.Scan(&answer)

		if answer == line[1] {
			correctAnswers++
		} else {
			wrongAnswers++
		}
	}
}
