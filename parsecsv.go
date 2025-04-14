package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	file, err := os.Open("query_params.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader.
	reader := csv.NewReader(file)

	// Specify the column index to print (0-based)
	columnIndex := 0 // Change this to the desired column

	// Read the CSV data line by line.
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			log.Fatal(err)
		}

		// Process the record (a slice of strings).
		//fmt.Println(record)
		// You can access individual fields like record[0], record[1], etc.
		if columnIndex < len(record) {
			// Print the specified column
			fmt.Println(record[columnIndex])
		} else {
			fmt.Printf("Column index %d out of range for row with %d columns\n", columnIndex, len(record))
		}
	}
}
