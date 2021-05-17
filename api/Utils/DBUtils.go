package Utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Read CSV file and parse the DB Configuration
func ReadDBConfiguration(file string) (resp [][]string, err error) {
	// Open the file
	recordFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("an error encountered :: %v", err)
	}

	//Initialize the reader
	reader := csv.NewReader(recordFile)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("an error encountered :: %v", err)
	}

	return records, nil
}

//Creates the DB Connection string
func ProcessDBConnectionString(data [][]string, err error) string {

	if err != nil {
		fmt.Println(fmt.Errorf("encountered error while processing Databse Configuration : %v", err))
		return ""
	}
	connStr := "?"
	params := make([]string, len(data))

	for _, v := range data {
		if v != nil {
			params = append(params, v[0]+"="+v[1])
		}
	}
	connStr += strings.Join(params, "&")

	return connStr
}
