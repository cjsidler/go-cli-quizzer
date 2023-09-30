package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

/*
read in quiz via CSV file
user can specify '-csv' flag and file name
otherwise, default to 'problems.csv'

ask user all questions, keep track of score

output number of correct responses vs total questions
*/

func main() {
	// Open problems.csv
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

		fmt.Println(record)
	}
}