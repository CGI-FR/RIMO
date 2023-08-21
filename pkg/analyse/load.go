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
func Load(inputPath string) (DataMap, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load %s : %w", inputPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data, err := LoadJSONLines(scanner)
	if err != nil {
		return nil, fmt.Errorf("couldn't load %s : %w", inputPath, err)
	}

	return data, nil
}

type DataMap map[string][]interface{}

// Reads JSON lines  structure: { "col_name1" : value1, "col_name2" : value1, ... }.
func LoadJSONLines(scanner *bufio.Scanner) (DataMap, error) {
	var data map[string][]interface{} = DataMap{}

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		lineMap := make(map[string]interface{})

		// Check for BOM character at beginning of file.
		bytes := scanner.Bytes()
		if lineNumber == 1 && len(bytes) > 2 {
			if scanner.Bytes()[0] == 0xEF && scanner.Bytes()[1] == 0xBB && scanner.Bytes()[2] == 0xBF {
				bytes = scanner.Bytes()[3:]
			}
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
