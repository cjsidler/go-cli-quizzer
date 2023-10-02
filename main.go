package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

func main() {

	// Get csv filename and open using a new csv reader
	filename := getFilename()
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(csvFile)

	// Initiate statistics variables
	totalQuestions := 0
	correctAnswers := 0

	// Ask each question in csv and get answer from user
	for {
		csvLine, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(csvLine) != 2 {
			log.Fatal("Malformed csv file. Each line in csv file must have one question and one answer.")
		}

		totalQuestions++

		question, answer := csvLine[0], csvLine[1]
		userAnswer := getUserAnswer(question, answer, totalQuestions)

		if userAnswer == answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect. Your answer: %v (Correct answer: %v)\n", userAnswer, answer)
		}

	}

	// Show the user's stats
	fmt.Printf("You answered %v/%v correct!\n", correctAnswers, totalQuestions)
}

// Checks if os.Args has "-csv filename.csv"
// Otherwise, returns "problems.csv" as a default
func getFilename() string {
	if len(os.Args) == 3 && os.Args[1] == "-csv" {
		return os.Args[2]
	} else {
		return "problems.csv"
	}
}

// Gets user input from stdin
func getUserAnswer(question string, answer string, totalQuestions int) string {
	var userAnswer string

	fmt.Printf("Question #%v: %v = ", totalQuestions, question)
	fmt.Scanln(&userAnswer)
	return userAnswer
}
