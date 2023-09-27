// Copyright (C) 2023 CGI France
//
// This file is part of RIMO.
//
// RIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// RIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with RIMO.  If not, see <http://www.gnu.org/licenses/>.

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
