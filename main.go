package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	correctAnswers int
	wrongAnswers   int
)

func main() {
	fileName, _, shuffle := parseUserInput()
	data := readQuestions(fileName)
	if shuffle {
		scramble(data)
	}
	mainLoop(data)
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, correctAnswers+wrongAnswers)
}

func parseUserInput() (string, int, bool) {
	csvPath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the order of questions")
	flag.Parse()
	return *csvPath, *timeLimit, *shuffle
}

func readQuestions(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	return data
}

func scramble(data [][]string) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	random.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
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
