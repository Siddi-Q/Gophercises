package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func getCommandLineFlags() (*string, *int, *bool) {
	csvFlag := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitFlag := flag.Int("limit", 30, "the time limit for the quiz in seconds. set to 0 if you don't want a limit")
	shuffleFlag := flag.Bool("shuffle", false, "set to true inorder to shuffle the quiz problems")
	flag.Parse()
	return csvFlag, limitFlag, shuffleFlag
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

func shuffleQuiz(quiz []problem) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(quiz), func(i, j int) {
		quiz[i], quiz[j] = quiz[j], quiz[i]
	})
}

func playQuizGame(quiz []problem) int {
	numCorrectAnswers := 0
	reader := bufio.NewReader(os.Stdin)

	for idx, problem := range quiz {
		question, correctAnswer := problem.question, problem.answer
		fmt.Printf("Problem #%d: %s = ", idx+1, question)

		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-2] // Remove \r\n from the user's answer in Windows
		if strings.TrimSpace(strings.ToLower(answer)) == strings.ToLower(correctAnswer) {
			numCorrectAnswers++
		}
	}
	return numCorrectAnswers
}

func playTimedQuizGame(quiz []problem, timeLimit int) int {
	numCorrectAnswers := 0
	reader := bufio.NewReader(os.Stdin)
	answerCh := make(chan string)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for idx, problem := range quiz {
		question, correctAnswer := problem.question, problem.answer
		fmt.Printf("Problem #%d: %s = ", idx+1, question)

		go func() {
			answer, _ := reader.ReadString('\n')
			answer = answer[:len(answer)-2] // Remove \r\n from the user's answer in Windows
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			return numCorrectAnswers
		case answer := <-answerCh:
			if strings.TrimSpace(strings.ToLower(answer)) == strings.ToLower(correctAnswer) {
				numCorrectAnswers++
			}
		}
	}
	return numCorrectAnswers
}

func main() {
	csvFile, timeLimit, shuffle := getCommandLineFlags()
	quiz, err := parsecsv(*csvFile)
	if err != nil {
		log.Fatal(err)
	}

	if *shuffle {
		shuffleQuiz(quiz)
	}

	var numCorrectAnswers int
	if *timeLimit == 0 {
		numCorrectAnswers = playQuizGame(quiz)
	} else {
		numCorrectAnswers = playTimedQuizGame(quiz, *timeLimit)
	}
	fmt.Printf("You scored %d out of %d.\n", numCorrectAnswers, len(quiz))
}
