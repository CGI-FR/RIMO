package rimo

import (
	"fmt"
	"log"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"

	"github.com/hexops/valast"
)

func AnalyseBase(reader Reader, writer Writer) error {
	baseName := reader.BaseName()
	log.Printf("Processing base : %s", baseName)
	base := model.NewBase(baseName)

	for reader.Next() { // itère colonne par colonne
		colValues, colName, tableName, err := reader.Value()
		if err != nil {
			return fmt.Errorf("failed to get column value : %w", err)
		}

		log.Printf("Processing %s column in %s table", tableName, colName)
		log.Printf("Column values : %v", colValues)

		column, err := metric.ComputeMetric(colName, colValues)
		if err != nil {
			return fmt.Errorf("failed to compute column : %w", err)
		}

		fmt.Printf("Column : %s\n", valast.String(column))

		base.AddColumn(column, tableName)
	}

	base.SortBase()
	log.Printf("Base processed : %s", valast.String(base))

	err := writer.Export(base)
	if err != nil {
		return fmt.Errorf("failed to export base : %w", err)
	}

	return nil
}
