package Utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

//Creates the connection params as a string
func ProcessDBConnectionString(data [][]string, err error) string {

	if err != nil {
		fmt.Println(fmt.Errorf("encountered error while processing Connection parameters : %v", err))
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

// Read MySQL server configuration
func ProcessDBServerConfiguration(data [][]string, err error) string {
	if err != nil {
		fmt.Println(fmt.Errorf("encountered error while processing Databse Configuration : %v", err))
		return ""
	}

	port, _ := strconv.Atoi(data[3][1])

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		data[0][1],
		data[1][1],
		data[2][1],
		port)
}
