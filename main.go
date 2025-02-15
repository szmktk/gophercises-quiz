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

type DefaultConfigurator struct{}
type CsvReader struct{}
type RandomScrambler struct{}
type ConsoleRunner struct{}

func main() {
	configurator := DefaultConfigurator{}
	reader := CsvReader{}
	scrambler := RandomScrambler{}
	runner := ConsoleRunner{}

	runQuizApp(configurator, reader, scrambler, runner)
}

func runQuizApp(configurator QuizConfigurator, reader QuestionReader, scrambler QuestionScrambler, runner QuizRunner) {
	fileName, timeLimit, shuffle := configurator.ParseUserInput()
	questions := reader.ReadQuestions(fileName)

	if shuffle {
		scrambler.Scramble(questions)
	}

	runner.MainLoop(questions, timeLimit)
}

func (DefaultConfigurator) ParseUserInput() (string, int, bool) {
	csvPath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the order of questions")
	flag.Parse()
	return *csvPath, *timeLimit, *shuffle
}

func (CsvReader) ReadQuestions(fileName string) []Question {
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

func (RandomScrambler) Scramble(data []Question) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	random.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

func (ConsoleRunner) MainLoop(questions []Question, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for idx, question := range questions {
		fmt.Printf("Problem #%d: %s = ", idx+1, question.question)

		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(questions))
			return
		case answer := <-answerChannel:
			if answer == question.answer {
				correctAnswers++
			} else {
				wrongAnswers++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(questions))
}
