package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
Read in quiz via CSV file
User can specify '-csv' flag and file name
Otherwise, defaults to 'problems.csv'

Ask user all questions, keep track of score

Output number of correct responses vs total questions
*/

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `question,answer`")
	flag.Parse()

	// Get csv filename and open using a new csv reader
	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(csvFile)

	// Initiate statistics variables
	correctAnswers := 0

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseLines(csvLines)

	// Ask each question in csv and verify answer
	for i, problem := range problems {
		currentProblem := i + 1
		userAnswer := getUserAnswer(problem, currentProblem)

		if userAnswer == problem.answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect. Your answer: %v (Correct answer: %v)\n", userAnswer, problem.answer)
		}
	}

	// Show the user's stats
	fmt.Printf("You answered %v/%v correct!\n", correctAnswers, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		if len(line) != 2 {
			log.Fatal("Malformed csv file. Each line in csv file must be in the format `question,answer`.")
		}
		problems[i] = problem{question: line[0], answer: line[1]}
	}

	return problems
}

// Gets user input from stdin
func getUserAnswer(problem problem, totalQuestions int) string {
	var userAnswer string

	fmt.Printf("Question #%v: %v = ", totalQuestions, problem.question)
	fmt.Scanln(&userAnswer)

	return userAnswer
}
