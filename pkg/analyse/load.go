//nolint:forcetypeassert
package analyse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	Numeric = "numeric"
	String  = "string"
	Boolean = "boolean"
)

var jsonFormatFunc = LoadNewJSONStruct

// Load .jsonl and return DataMap.
func Load(inputPath string, format string) DataMap {
	if format == "old" {
		jsonFormatFunc = LoadOldJSONStruct
	} else if format == "new" {
		jsonFormatFunc = LoadNewJSONStruct
	} else {
		panic("Format not supported")
	}
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := jsonFormatFunc(scanner)

	return data
}

type DataMap map[string][]interface{}

// Reads JSON new structure.
// { "col_name" : [value1, value2, ...] }.
func LoadNewJSONStruct(scanner *bufio.Scanner) DataMap {
	// Instantiate dataMap map[string]dataCol
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			panic(err)
		}

		for colName := range lineMap {
			// Check if colName is already in dataMap.
			if _, ok := data[colName]; ok {
				// Column already exist in dataMap.
				// Raise error
				panic("Column already exist in dataMap")
			} else {
				// Column does not exist in dataMap.
				// Assert linemap[colName] is []interface{}
				if _, ok := lineMap[colName].([]interface{}); !ok {
					// print lineMap[colName]
					fmt.Println(lineMap[colName])
					panic("Column is not []interface{}")
				}
				data[colName] = lineMap[colName].([]interface{})
			}
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// Reads previous JSON structure.
// { "col_name" : value, "col_name2" : value2, ... }.
func LoadOldJSONStruct(scanner *bufio.Scanner) DataMap {
	// Instantiate dataMap map[string][]interface{}
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			panic(err)
		}

		for colName := range lineMap {
			if _, ok := data[colName]; !ok {
				// data[colName] does not exist, instantiate it
				data[colName] = []interface{}{lineMap[colName]}
			} else {
				// Add values to data[colName]
				data[colName] = append(data[colName], lineMap[colName])
			}
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}
