package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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

func startQuiz(problems [][]string) (score int) {
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
	return
}

func main() {
	csv := flag.String("csv", "problems.csv", "a CSV file in the format of 'question,answer'")
	flag.Parse()
	problems, err := getProblems(*csv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	score := startQuiz(problems)
	maxScore := len(problems)
	fmt.Printf("You scored %d out of %d.\n", score, maxScore)
}
