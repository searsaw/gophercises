package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	text   string
	answer string
}

func getProblemData(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	return csvReader.ReadAll()
}

func createProblems(questionData [][]string) []problem {
	questions := make([]problem, len(questionData))

	for i, row := range questionData {
		questions[i] = problem{
			text:   row[0],
			answer: strings.TrimSpace(row[1]),
		}
	}

	return questions
}

func main() {
	problemsFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	questionData, err := getProblemData(*problemsFile)
	if err != nil {
		fmt.Printf("Error when reading question data file at path %s\n", *problemsFile)
		os.Exit(1)
	}

	problems := createProblems(questionData)
	correctAnswers := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.text)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime's up!\nYou got %d out of %d questions right!\n", correctAnswers, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("You got %d out of %d questions right!\n", correctAnswers, len(problems))
}
