package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
)

func GetJSONSchema() (string, error) {
	resBytes, err := json.MarshalIndent(jsonschema.Reflect(&Base{}), "", "  ") //nolint:exhaustruct
	if err != nil {
		return "", fmt.Errorf("couldn't unmarshall Base in JSON : %w", err)
	}

	return string(resBytes), nil
}

func NewBase(name string) *Base {
	return &Base{
		Name:   name,
		Tables: make([]Table, 0),
	}
}

var ErrBaseFormat = errors.New("error while decoding yaml file in a Base struct")

// Can be improved.
func LoadBase(path string) (*Base, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error while opening file: %w", err)
	}

	decoder := yaml.NewDecoder(file)

	var base Base

	err = decoder.Decode(&base)
	if err != nil {
		return nil, ErrBaseFormat
	}

	file.Close()

	return &base, nil
}

func RemoveSampleFromBase(base *Base) {
	for tableI, table := range base.Tables {
		for columnJ, column := range table.Columns {
			column.MainMetric.Sample = nil

			if column.Type == ColType.String {
				for freqLen := range column.StringMetric.MostFreqLen {
					column.StringMetric.MostFreqLen[freqLen].Sample = nil
				}

				for freqLen := range column.StringMetric.LeastFreqLen {
					column.StringMetric.LeastFreqLen[freqLen].Sample = nil
				}
			}

			base.Tables[tableI].Columns[columnJ] = column
		}
	}
}

func (base *Base) SortBase() {
	for _, table := range base.Tables {
		sort.Slice(table.Columns, func(i, j int) bool {
			return table.Columns[i].Name < table.Columns[j].Name
		})
	}

	sort.Slice(base.Tables, func(i, j int) bool {
		return base.Tables[i].Name < base.Tables[j].Name
	})
}

func (base *Base) AddColumn(column Column, tableName string) {
	mapTableName := make(map[string]int)
	for index, table := range base.Tables {
		mapTableName[table.Name] = index
	}

	if index, ok := mapTableName[tableName]; ok {
		// If the table exists, append the column to the table
		base.Tables[index].Columns = append(base.Tables[index].Columns, column)
	} else {
		// If the table does not exist, create a new table and add it to the base
		table := Table{
			Name:    tableName,
			Columns: []Column{column},
		}
		base.Tables = append(base.Tables, table)
	}
}

// If the table does not exist, create a new table and add it to the base
// table := Table{Name: tableName, Columns: []Column{column}}
// base.Tables = append(base.Tables, table)
