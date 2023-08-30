package rimo

import "fmt"

func AnalyseBase(reader Reader, writer Writer) error {
	baseName := reader.BaseName()

	base := NewBase(baseName)

	for reader.Next() { // itère colonne par colonne
		colValues, colName, tableName, err := reader.Value()
		if err != nil {
			return fmt.Errorf("failed to get column value : %w", err)
		}

		column, err := Analyse(colValues, colName)
		if err != nil {
			return fmt.Errorf("failed to compute model.column : %w", err)
		}

		base.AddColumn(column, tableName)
	}

	base.SortBase()

	err := writer.Export(base)
	if err != nil {
		return fmt.Errorf("failed to export base : %w", err)
	}

	return nil
}

// Process a input and return a model.Column.
func Analyse(colValues []interface{}, colName string) (Column, error) {
	return Column{}, nil //nolint:exhaustruct
}
