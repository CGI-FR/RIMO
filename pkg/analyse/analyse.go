package analyse

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cgi-fr/rimo/pkg/model"
)

// Handle execution pipeline of analyse pkg.
func Analyse(inputList []string, outputPath string) {
	// Ensure all input files relate to same Base.
	baseName, err := GetBaseName(inputList[0])
	if err != nil {
		log.Fatalf("failed to extract database name: %v", err)
	}

	for i := range inputList {
		newBaseName, err := GetBaseName(inputList[i])
		if err != nil {
			log.Fatalf("failed to extract database name: %v", err)
		}

		if newBaseName != baseName {
			log.Fatalf("input files do not relate to same Base: %s", baseName)
		}
	}

	// Treatment of input file.

	tables := make([]model.Table, 0, len(inputList))

	for i := range inputList {
		inputPath := inputList[i]
		// Extract Base and Table name from inputFilePath.
		tableName, err := GetTableName(inputPath)
		if err != nil {
			log.Fatalf("failed to extract table name: %v", err)
		}
		// Load inputFilePath.
		data := Load(inputPath, "new")

		// Analyse
		var cols []model.Column

		for colName, values := range data {
			column, err := ComputeMetric(colName, values)
			if err != nil {
				log.Fatalf("failed to compute metric: %v", err)
			}

			cols = append(cols, column)
		}

		table := model.Table{
			Name:    tableName,
			Columns: cols,
		}

		// Append tables to Base structure.
		tables = append(tables, table)
	}

	base := model.Base{
		Name:   baseName,
		Tables: tables,
	}

	// Export
	err = Export(base, outputPath)
	if err != nil {
		log.Fatalf("failed to export: %v", err)
	}
}

// Error definitions.

var ErrNonExtractibleValue = errors.New("couldn't extract base or table name from path")

func GetBaseName(path string) (string, error) {
	// path format : /path/to/jsonl/BASE_TABLE.jsonl
	baseName := filepath.Base(path)
	baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
	// Split at _ to get Base name.
	parts := strings.Split(baseName, "_")
	if len(parts) < 2 {
		return "", fmt.Errorf("%w : unable to extract base name from %s", ErrNonExtractibleValue, path)
	}

	baseName = parts[0]
	if baseName == "" {
		return "", fmt.Errorf("%w : base name is empty from %s", ErrNonExtractibleValue, path)
	}

	return baseName, nil
}

func GetTableName(path string) (string, error) {
	// path format : /path/to/jsonl/BASE_TABLE.jsonl
	tableName := filepath.Base(path)
	tableName = tableName[:len(tableName)-len(filepath.Ext(tableName))]
	// Split at _ to get Table name.
	parts := strings.Split(tableName, "_")
	if len(parts) < 2 {
		return "", fmt.Errorf("%w : unable to extract table name from %s", ErrNonExtractibleValue, path)
	}

	tableName = parts[1]
	if tableName == "" {
		return "", fmt.Errorf("%w : table name is empty from %s", ErrNonExtractibleValue, path)
	}

	return tableName, nil
}

func CheckFile(path string) {
	fileInfo, err := os.Stat(path)

	absPath, _ := filepath.Abs(path)
	// Check if the file exists
	if os.IsNotExist(err) {
		log.Fatalf("file does not exist: %s", absPath)
	}
	// Check if the file is a regular file
	if !fileInfo.Mode().IsRegular() {
		log.Fatalf("not a regular file: %s", absPath)
	}
}
