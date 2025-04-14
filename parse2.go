package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("query_params.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	filter := "host_000009" // Change this to your filter string
	columnIndex := 0        // Change this to the index of the column you want to filter by

	fmt.Println("Filtered rows:")
	for i, row := range records {
		// Print header
		if i == 0 {
			fmt.Println(row)
			continue
		}

		if strings.Contains(row[columnIndex], filter) {
			fmt.Println(row)
		}
	}
}
