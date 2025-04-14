package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Open the CSV file
	file, err := os.Open("query_params.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Exit if there are no records
	if len(records) <= 1 {
		fmt.Println("No data records found in CSV, exiting.")
		return
	}

	// Group data by the first column
	groupedData := make(map[string][][]string)
	for _, record := range records[1:] { // Skip header row
		key := record[0]
		groupedData[key] = append(groupedData[key], record)
	}

	// Order data within groups by the first column
	for _, group := range groupedData {
		sort.Slice(group, func(i, j int) bool {
			return group[i][0] < group[j][0]
		})
	}

	// Print grouped and ordered data
	fmt.Println(records[0]) // Print header
	for _, group := range groupedData {
		for _, record := range group {
			fmt.Println(record)
		}
	}
}
