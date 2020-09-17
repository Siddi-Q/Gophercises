package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func getCommandLineFlags() (*string, *int) {
	csvFlag := flag.String("csv", "", "a string")
	limitFlag := flag.Int("limit", 0, "an int")
	flag.Parse()
	return csvFlag, limitFlag
}

func readcsv() ([][]string, error) {
	f, err := os.Open("C:\\Users\\saddi\\Downloads\\problems.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var quiz [][]string
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quiz = append(quiz, record)
	}

	return quiz, nil
}

func playQuizGame(quiz [][]string) int {
	var answer string
	numCorrectAnswers := 0

	for idx, arr := range quiz {
		question, correctAnswer := arr[0], arr[1]
		fmt.Printf("Problem #%d: %s = ", idx+1, question)
		fmt.Scanln(&answer)

		if answer == correctAnswer {
			numCorrectAnswers++
		}
	}
	return numCorrectAnswers
}

func main() {
	quiz, err := readcsv()
	if err != nil {
		log.Fatal(err)
	}
	numCorrectAnswers := playQuizGame(quiz)
	fmt.Printf("You scored %d out of %d.\n", numCorrectAnswers, len(quiz))
}
