package game

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

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
