package rimo_test

import (
	"log"
	"math"
	"testing"

	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/rimo"
)

// TESTS

func TestTestInterface(t *testing.T) {
	t.Parallel()

	var _ rimo.Reader = (*TestReader)(nil)

	var _ rimo.Writer = (*TestWriter)(nil)
}

// TestReader implementation

type colInput struct {
	ColName   string
	ColValues []interface{}
}

type TestReader struct {
	baseName   string
	data       []colInput
	tableNames []string // Next() will progressively change tableName
	// internal
	index            int
	currentValues    []interface{}
	currentColName   string
	currentTableName string
}

func (r *TestReader) BaseName() string {
	return r.baseName
}

func (r *TestReader) Next() bool {
	if r.index == len(r.data) {
		log.Println("End of data")

		return false
	}

	// update tableName
	if len(r.tableNames) == len(r.data) {
		r.currentTableName = r.tableNames[r.index]
	} else {
		// use a percentage to determine the table name to use from the list
		percentageComplete := float64(r.index) / float64(len(r.data))
		expectedTableIndex := percentageComplete * float64(len(r.tableNames))
		roundedTableIndex := math.Floor(expectedTableIndex)
		tableNameIndex := int(roundedTableIndex)

		r.currentTableName = r.tableNames[tableNameIndex]
	}

	r.currentColName = r.data[r.index].ColName
	r.currentValues = r.data[r.index].ColValues
	r.index++

	return true
}

func (r *TestReader) Value() ([]interface{}, string, string, error) { //nolint:wsl
	// log.Printf("Processing %s column in %s table", r.currentTableName, r.currentColName)

	return r.currentValues, r.currentColName, r.currentTableName, nil
}

// TestWriter implementation

type TestWriter struct {
	base model.Base
}

func (w *TestWriter) Export(base *model.Base) error {
	w.base = *base

	return nil
}

func (w *TestWriter) Base() *model.Base {
	return &w.base
}
