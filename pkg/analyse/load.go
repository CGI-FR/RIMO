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

type DataMap map[string]DataCol

type DataCol struct {
	ColType string
	Values  []interface{}
}

// Load .jsonl and return DataMap.
func Load(inputPath string) DataMap {
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
func LoadJSONStruct(scanner *bufio.Scanner) DataMap {
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

// Build a map of column names to column types.
func BuildColType(data DataMap) DataMap {
	for colName, colData := range data {
		// Iterate till colType is not unknown
		for i := 0; i < len(colData.Values) && data[colName].ColType == "unknown"; i++ {
			data[colName] = DataCol{
				ColType: TypeOf(colData.Values[i]),
				Values:  colData.Values,
			}
		}
	}

	return data
}

func TypeOf(v interface{}) string {
	switch v.(type) {
	case int:
		return Numeric
	case float64:
		return Numeric
	case json.Number:
		return Numeric
	case string:
		return String
	case bool:
		return Boolean
	default:
		return "unknown"
	}
}
