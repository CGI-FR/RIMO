//nolint:forcetypeassert
package analyse

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

// Load .jsonl and return DataMap.
func Load(inputPath string, format string) (DataMap, error) {
	var jsonFormatFunc func(*bufio.Scanner) (DataMap, error)

	switch format {
	case "old":
		jsonFormatFunc = LoadOldJSONStruct
	case "new":
		jsonFormatFunc = LoadNewJSONStruct
	default:
		jsonFormatFunc = LoadNewJSONStruct
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load %s : %w", inputPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data, err := jsonFormatFunc(scanner)
	if err != nil {
		return nil, fmt.Errorf("couldn't load %s : %w", inputPath, err)
	}

	return data, nil
}

type DataMap map[string][]interface{}

// Reads JSON new structure.
// { "col_name" : [value1, value2, ...] }.
func LoadNewJSONStruct(scanner *bufio.Scanner) (DataMap, error) {
	// Instantiate dataMap map[string]dataCol
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			return nil, ErrScanJSON
		}

		for colName := range lineMap {
			// Check if colName is already in dataMap.
			if _, ok := data[colName]; ok {
				// Column already exist in dataMap.
				return nil, ErrSameColumn
			}
			// Column does not exist in dataMap.
			if _, ok := lineMap[colName].([]interface{}); !ok {
				return nil, fmt.Errorf("%w: %s of type %T", ErrNotInterface, colName, lineMap[colName])
			}

			data[colName] = lineMap[colName].([]interface{})
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("couldn't scan JSON: %w", err)
	}

	return data, nil
}

// Reads previous JSON structure.
// { "col_name" : value, "col_name2" : value2, ... }.
func LoadOldJSONStruct(scanner *bufio.Scanner) (DataMap, error) {
	// Instantiate dataMap map[string][]interface{}
	data := DataMap{}

	for scanner.Scan() {
		lineMap := make(map[string]interface{})

		err := json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal JSON: %w", err)
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
