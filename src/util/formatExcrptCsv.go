package util

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func ReadExcrptCsv(path string) ([]string, []string, []string, []string) {
	// Open the CSV file
	file, err := os.Open("Lønkonto.csv")
	if err != nil {
		log.Fatalln("Coud not open excerpt file.", err)
	}
	defer file.Close()

	// Create a new CSV ReadExcrptCsv
	reader := csv.NewReader(file)

	// Dont test on number of fields
	reader.FieldsPerRecord = -1

	reader.Comma = ';'
	// Read all records from the CSV file
	// records, err := reader.ReadAll()

	return getExcrpt(reader)
}

func getExcrpt(reader *csv.Reader) ([]string, []string, []string, []string) {
	// Print each record
	var dates []string
	var amounts []string
	var descriptions []string
	var balances []string

	i := 0
	for {
		row, err := reader.Read()
		// Dont read header-line
		if i > 0 {
			if err != nil {
				// Check for end of file
				if err.Error() == "EOF" {
					break
				} else {
					log.Fatal("Error:", err)
				}
			}

			dates = append(dates, row[0])

			amounts = append(amounts, row[1])

			descriptions = append(descriptions, row[5])
			balances = append(balances, row[6])
			//
		}

		i++
	}
	return dates, amounts, descriptions, balances
}

// Function to encode the value using UTF-8
func encodeValue(inputValue string) string {
	// Encode the value using UTF-8
	encodedValue, err := json.Marshal(inputValue)
	if err != nil {
		log.Fatalf("Unable to encode value: %v", err)
	}
	return string(encodedValue)
}
