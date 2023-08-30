package infra_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	t.Parallel()

	inputFile := filepath.Join(dataDir, "data0/data_input.jsonl")

	reader, err := infra.FilesReaderFactory([]string{inputFile})
	assert.NoError(t, err)

	// Assertions.

	actualBaseName := reader.BaseName()
	expectedBaseName := "data"
	assert.Equal(t, expectedBaseName, actualBaseName)

	expectedTableName := "input"
	expectedDataMap := map[string][]interface{}{
		"address": {"PSC", "095", "06210"},
		"age":     {nil, nil, float64(61)},
		"major":   {true, false, true},
		"empty":   {nil, nil, nil},
	}

	for reader.Next() {
		values, colName, tableName, err := reader.Value()
		if err != nil {
			assert.NoError(t, err)
		}

		expectedColData, ok := expectedDataMap[colName]
		if !ok {
			assert.Fail(t, "column name not found : %s", colName)
		}

		assert.Equal(t, expectedColData, values)
		assert.Equal(t, expectedTableName, tableName)
	}
}

func TestReaderMultipleFiles(t *testing.T) {
	t.Parallel()

	inputFile := filepath.Join(dataDir, "data0/data_input.jsonl")
	inputFile2 := filepath.Join(dataDir, "data0/data_input2.jsonl")
	reader, err := infra.FilesReaderFactory([]string{inputFile, inputFile2})
	assert.NoError(t, err)

	for reader.Next() {
		values, colName, tableName, err := reader.Value()
		if err != nil {
			assert.NoError(t, err)
		}

		fmt.Printf("%s.%s: %v\n", tableName, colName, values)
	}
}
