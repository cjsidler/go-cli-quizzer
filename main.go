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

		question, answer := record[0], record[1]

		fmt.Println("Question:", question)

		userAnswer := getUserAnswer()

		fmt.Printf("Your answer: %v (Correct answer: %v)\n", userAnswer, answer)
	}
}

func getUserAnswer() string {
	var userAnswer string
	fmt.Scanln(&userAnswer)
	return userAnswer
}
