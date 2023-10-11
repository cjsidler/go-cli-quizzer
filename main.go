package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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
	var csvFilename string
	var quizTimer int
	flag.StringVar(&csvFilename, "csv", "problems.csv", "a csv file in the format of `question,answer`")
	flag.IntVar(&quizTimer, "timer", 30, "a time limit in seconds for the duration of the quiz")
	flag.Parse()

	fmt.Println(`

	 ██████  ██    ██ ██ ███████     ████████ ██ ███    ███ ███████ ██ 
	██    ██ ██    ██ ██    ███         ██    ██ ████  ████ ██      ██ 
	██    ██ ██    ██ ██   ███          ██    ██ ██ ████ ██ █████   ██ 
	██ ▄▄ ██ ██    ██ ██  ███           ██    ██ ██  ██  ██ ██         
	 ██████   ██████  ██ ███████        ██    ██ ██      ██ ███████ ██ 
	    ▀▀                                                             
	`)
	fmt.Printf("You have %v seconds. Time starts now!\n\n", quizTimer)

	// Get csv filename and open using a new csv reader
	csvFile, err := os.Open(csvFilename)
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

	timer := time.NewTimer(time.Duration(quizTimer) * time.Second)
	answerChan := make(chan string)

	// Ask each question in csv and verify answer
	for i, problem := range problems {
		currentProblem := i + 1

		go getUserAnswer(problem, currentProblem, answerChan)

		select {
		case <-timer.C:
			fmt.Println()
			fmt.Printf("Time is up! You answered %v/%v correct!\n", correctAnswers, len(problems))
			return
		case userAnswer := <-answerChan:
			if userAnswer == problem.answer {
				fmt.Println("Correct!")
				correctAnswers++
			} else {
				fmt.Printf("Incorrect. Your answer: %v (Correct answer: %v)\n", userAnswer, problem.answer)
			}
			fmt.Println()
		}
	}

	fmt.Printf("You answered %v/%v correct!\n", correctAnswers, len(problems))
}

// Parse csv file into a slice of problems
func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problems
}

// Gets user input from stdin and send it into answer channel
func getUserAnswer(problem problem, problemNumber int, answerChan chan string) {
	var userInput string
	fmt.Printf("Question #%v: %v = ", problemNumber, problem.question)
	fmt.Scanln(&userInput)
	answerChan <- userInput
}
