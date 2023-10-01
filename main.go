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
Otherwise, default to 'problems.csv'

Ask user all questions, keep track of score

Output number of correct responses vs total questions
*/

func main() {
	// Check if os.Args has -csv filename.csv
	// Otherwise, use problems.csv as a default
	fmt.Println(os.Args)

	// Open problems.csv and use a new csv reader
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(csvFile)

	totalQuestions := 0
	correctAnswers := 0

	// Read and print every line of the csv
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(record) < 2 {
			log.Fatal("Malformed csv file. Each record in csv file must have a question and answer.")
		}

		totalQuestions++

		question, answer := record[0], record[1]

		fmt.Printf("Question #%v: %v = ", totalQuestions, question)

		userAnswer := getUserAnswer()

		if userAnswer == answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect. Your answer: %v (Correct answer: %v)\n", userAnswer, answer)
		}

	}

	fmt.Printf("You answered %v questions correctly out of a total of %v questions.\n", correctAnswers, totalQuestions)
}

func getUserAnswer() string {
	var userAnswer string
	fmt.Scanln(&userAnswer)
	return userAnswer
}
