package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func getCommandLineFlags() (*string, *int) {
	csvFlag := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitFlag := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return csvFlag, limitFlag
}

func parsecsv(csvFile string) ([]problem, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var quiz []problem
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quiz = append(quiz, problem{question: record[0], answer: strings.TrimSpace(record[1])})
	}

	return quiz, nil
}

func playQuizGame(quiz []problem) int {
	var answer string
	numCorrectAnswers := 0

	for idx, problem := range quiz {
		question, correctAnswer := problem.question, problem.answer
		fmt.Printf("Problem #%d: %s = ", idx+1, question)
		fmt.Scanln(&answer)

		if answer == correctAnswer {
			numCorrectAnswers++
		}
	}
	return numCorrectAnswers
}

func main() {
	csvFile, _ := getCommandLineFlags()
	quiz, err := parsecsv(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	numCorrectAnswers := playQuizGame(quiz)
	fmt.Printf("You scored %d out of %d.\n", numCorrectAnswers, len(quiz))
}
