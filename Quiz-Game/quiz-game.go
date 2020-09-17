package game

import (
	"encoding/csv"
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
