package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	correctAnswers int
	wrongAnswers   int
)

type Question struct {
	question string
	answer   string
}

func main() {
	fileName, _, shuffle := parseUserInput()
	questions := readQuestions(fileName)
	if shuffle {
		scramble(questions)
	}
	mainLoop(questions)
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, correctAnswers+wrongAnswers)
}

func parseUserInput() (string, int, bool) {
	csvPath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the order of questions")
	flag.Parse()
	return *csvPath, *timeLimit, *shuffle
}

func readQuestions(fileName string) []Question {
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

	questions := make([]Question, len(data))
	for idx, line := range data {
		questions[idx] = Question{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return questions
}

func scramble(data []Question) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	random.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

func mainLoop(data []Question) {
	for idx, question := range data {
		fmt.Printf("Problem #%d: %s = ", idx+1, question.question)
		var answer string
		fmt.Scan(&answer)

		if answer == question.answer {
			correctAnswers++
		} else {
			wrongAnswers++
		}
	}
}
