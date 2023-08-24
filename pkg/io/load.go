package io

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	Numeric = "numeric"
	String  = "string"
	Boolean = "boolean"
)

var (
	ErrNotInterface = errors.New("line is not an interface")
	ErrScanJSON     = errors.New("couldn't scan JSON")
	ErrSameColumn   = errors.New("column found twice in JSON")
)

type DataMap map[string][]interface{}

// Load .jsonl and return DataMap.
func Load(inputPath string) (DataMap, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load %s : %w", inputPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data, err := LoadJSONLines(scanner)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Reads JSON lines  structure: { "col_name1" : value1, "col_name2" : value1, ... }.
func LoadJSONLines(scanner *bufio.Scanner) (DataMap, error) {
	var data map[string][]interface{} = DataMap{}

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		bytes := scanner.Bytes()
		lineMap := make(map[string]interface{})

		if lineNumber == 1 {
			bytes = stripBOM(bytes)
		}

		err := json.Unmarshal(bytes, &lineMap)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal JSON in %s : %w", scanner.Text(), err)
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
		return nil, fmt.Errorf("couldn't scan JSON: %w", err)
	}

	return data, nil
}

func stripBOM(bytes []byte) []byte {
	if len(bytes) > 2 && bytes[0] == 0xEF && bytes[1] == 0xBB && bytes[2] == 0xBF {
		return bytes[3:]
	}

	return bytes
}
