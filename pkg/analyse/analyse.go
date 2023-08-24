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

package analyse

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cgi-fr/rimo/pkg/io"
	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
)

var ErrWrongParameter = errors.New("wrong parameter")

// Handle execution pipeline of rimo analyse.
func Orchestrator(inputList []string, outputPath string) error {
	// Process input
	err := ProcessInput(inputList, outputPath)
	if err != nil {
		return err
	}

	// Compute model.base
	base, err := Build(inputList)
	if err != nil {
		return err
	}

	// Export rimo.yaml
	outputPath = filepath.Join(outputPath, base.Name+".yaml")

	err = io.Export(base, outputPath)
	if err != nil {
		return fmt.Errorf("%w : cannot export to %s", err, outputPath)
	}

	return nil
}

func ProcessInput(inputList []string, outputPath string) error {
	// verify output dirPath
	err := io.ValidateDirPath(outputPath)
	if err != nil {
		return fmt.Errorf("failed to validate output path: %w", err)
	}

	// validate input filepath
	for i := range inputList {
		err := io.ValidateFilePath(inputList[i])
		if err != nil {
			return fmt.Errorf("failed to validate input file: %w", err)
		}
	}

	// verify that input files relates to the same base
	err = BaseIsUnique(inputList)
	if err != nil {
		return fmt.Errorf("failed to validate input file: %w", err)
	}

	return nil
}

// Return a model.Base from inputList.
func Build(inputList []string) (model.Base, error) {
	baseName, _, err := ExtractName(inputList[0])
	if err != nil {
		return model.Base{}, fmt.Errorf("failed to extract base name for %s: %w", inputList[0], err)
	}

	base := model.Base{
		Name:   baseName,
		Tables: []model.Table{},
	}

	for _, inputPath := range inputList {
		_, tableName, err := ExtractName(inputPath)
		if err != nil {
			return model.Base{}, fmt.Errorf("failed to extract table name for %s: %w", inputPath, err)
		}

		columns, err := Analyse(inputPath)
		if err != nil {
			return model.Base{}, fmt.Errorf("failed to analyse %s: %w", inputPath, err)
		}

		// Add columns to base
		table := model.Table{
			Name:    tableName,
			Columns: columns,
		}
		base.Tables = append(base.Tables, table)
	}

	SortBase(&base)

	return base, nil
}

// Sort tables and columns by name.
func SortBase(base *model.Base) {
	for i := range base.Tables {
		sort.Slice(base.Tables[i].Columns, func(j, k int) bool {
			return base.Tables[i].Columns[j].Name < base.Tables[i].Columns[k].Name
		})
	}

	sort.Slice(base.Tables, func(i, j int) bool {
		return base.Tables[i].Name < base.Tables[j].Name
	})
}

// Return a list of column from a jsonl file.
func Analyse(path string) ([]model.Column, error) {
	// Load file in a dataMap.
	data, err := io.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load jsonl file: %w", err)
	}

	columns := []model.Column{}

	for colName, values := range data {
		column, err := metric.ComputeMetric(colName, values)
		if err != nil {
			return nil, fmt.Errorf("failed to compute metric: %w", err)
		}

		columns = append(columns, column)
	}

	return columns, nil
}

// Error definitions.

var ErrNonExtractibleValue = errors.New("couldn't extract base or table name from path")

func ExtractName(path string) (string, string, error) {
	// path format : /path/to/jsonl/BASE_TABLE.jsonl
	fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(filepath.Base(path)))

	parts := strings.Split(fileName, "_")
	if len(parts) != 2 { //nolint:gomnd
		return "", "", fmt.Errorf("%w : %s", ErrNonExtractibleValue, path)
	}

	baseName := parts[0]
	if baseName == "" {
		return "", "", fmt.Errorf("%w : base name is empty from %s", ErrNonExtractibleValue, path)
	}

	tableName := parts[1]
	if tableName == "" {
		return "", "", fmt.Errorf("%w : table name is empty from %s", ErrNonExtractibleValue, path)
	}

	return baseName, tableName, nil
}

var ErrNonUniqueBase = errors.New("base name is not unique")

func BaseIsUnique(pathList []string) error {
	baseName, _, err := ExtractName(pathList[0])
	if err != nil {
		return err
	}

	for _, path := range pathList {
		baseNameI, _, err := ExtractName(path)
		if err != nil {
			return err
		}

		if baseName != baseNameI {
			return fmt.Errorf("%w : %s and %s", ErrNonUniqueBase, baseName, baseNameI)
		}
	}

	return nil
}
