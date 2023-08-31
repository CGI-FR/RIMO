package model

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/invopop/jsonschema"
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
	// Check if the table already exists in the base
	for _, table := range base.Tables {
		if table.Name == tableName {
			// Add the column to the existing table
			table.Columns = append(table.Columns, column)

			return
		}
	}

	// If the table does not exist, create a new table and add it to the base
	table := Table{Name: tableName, Columns: []Column{column}}
	base.Tables = append(base.Tables, table)
}
