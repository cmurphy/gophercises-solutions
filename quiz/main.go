package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func getProblems(csvFile string) ([][]string, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
	}()
	r := csv.NewReader(f)
	problems, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return problems, nil
}

var score int

func startQuiz(problems [][]string, done chan bool) {
	for i, p := range problems {
		question := p[0]
		answer := p[1]
		var input string
		fmt.Printf("Problem #%d: %s = ", i+1, question)
		fmt.Scanln(&input)
		if input == answer {
			score++
		}
	}
	close(done)
}

func main() {
	csv := flag.String("csv", "problems.csv", "a CSV file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	problems, err := getProblems(*csv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	maxScore := len(problems)
	done := make(chan bool)
	go startQuiz(problems, done)
	select {
	case <-done:
		fmt.Printf("You scored %d out of %d.\n", score, maxScore)
	case <-time.After(time.Duration(*limit) * time.Second):
		fmt.Printf("\nYou scored %d out of %d.\n", score, maxScore)
	}
}
