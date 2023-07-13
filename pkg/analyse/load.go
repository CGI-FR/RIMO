package analyse

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type dataMap map[string]dataCol

type dataCol struct {
	colType string
	values  []interface{}
}

// Load .jsonl and return DataMap
func load(inputPath string) dataMap {
	// Open the file
	CheckFile(inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Instantiate dataMap map[string]interface{}
	data := dataMap{}
	// Read each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineMap := make(map[string]interface{})
		err = json.Unmarshal(scanner.Bytes(), &lineMap)
		if err != nil {
			panic(err)
		}

		for colName := range lineMap {
			if _, ok := data[colName]; !ok {
				data[colName] = dataCol{
					colType: "unknown",
					values:  lineMap[colName].([]interface{}),
				}
			} else {
				log.Fatalf("2 columns with same name: %s", colName)
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// Build a map of column names to column types
func buildColType(data dataMap) dataMap {
	for colName, colData := range data {
		// Iterate till colType is not unknown
		for i := 0; i < len(colData.values) && data[colName].colType == "unknown"; i++ {
			data[colName] = dataCol{
				colType: typeOf(colData.values[i]),
				values:  colData.values,
			}
		}
	}
	return data
}

func typeOf(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case json.Number:
		return "json.Number"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		// return string(reflect.TypeOf(row_value))
		return "unknown"
	}
}
