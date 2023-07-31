//nolint:forcetypeassert
package analyse

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

const (
	Numeric = "numeric"
	String  = "string"
	Boolean = "boolean"
)

// Define which format of JSON is used.
var jsonFormatFunc = LoadNewJSONStruct //nolint:gochecknoglobals

// Load .jsonl and return DataMap.
func Load(inputPath string) []DataCol {
	// Open the file
	CheckFile(inputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := jsonFormatFunc(scanner)

	return data
}

// Reads JSON new structure.
// { "col_name" : [value1, value2, ...] }.
func LoadNewJSONStruct(scanner *bufio.Scanner) []DataCol {
	// Instantiate dataMap map[string]dataCol
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			panic(err)
		}

		for colName := range lineMap {
			if _, ok := data[colName]; ok {
				log.Fatalf("column was found twice: %s", colName)
			}

			dataCol := DataCol{
				ColType: "unknown",
				Values:  lineMap[colName].([]interface{}),
			}
			data[colName] = dataCol
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// Reads JSON structure.
// { "col_name" : value, "col_name2" : value2, ... }.
func LoadJSONStruct(scanner *bufio.Scanner) []DataCol {
	// Instantiate dataMap map[string]dataCol
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			panic(err)
		}

		for colName := range lineMap {
			if _, ok := data[colName]; !ok {
				data[colName] = DataCol{
					ColType: "unknown",
					Values:  lineMap[colName].([]interface{}),
				}
			} else {
				log.Fatalf("2 columns with same name: %s", colName)
			}
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}
