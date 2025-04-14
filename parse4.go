package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "ln34k34b1m.gag0anddix.tsdb.cloud.timescale.com"
	port     = 32878
	user     = "tsdbadmin"
	password = "hzcejp9no6nwtf5s"
	dbname   = "tsdb"
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	filter := []string{"host_000004", "host_000003"} // Change this to your filter string
	columnIndex := 0                                 // Change this to the index of the column you want to filter by

	fmt.Println("Filtered rows:")
	for j := 0; j < len(filter); j++ {
		for i, row := range records {
			// Print header
			if i == 0 {
				//fmt.Println(row)
				continue
			}

			if strings.Contains(row[columnIndex], filter[j]) {
				fmt.Println(row[0])
				longquery := fmt.Sprintf("select max(usage), min(usage) from cpu_usage where host='%s' and ts between '%s' and '%s';", row[0], row[1], row[2])
				//fmt.Println(longquery)
				//queryArray := strings.Split(longquery, ";")
				//fmt.Printf("%v\n", queryArray[0])
				rows, err := db.Query(longquery)
				if err != nil {
					log.Fatal("Query1 error:", err)
				}
				defer rows.Close()
				//fmt.Println("Users:")
				for rows.Next() {
					var maxusage float32
					var minusage float32

					err := rows.Scan(&maxusage, &minusage)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("MaxCPU: %f, MinCPU: %f\n", maxusage, minusage)
				}

			}
		}
	}
}
