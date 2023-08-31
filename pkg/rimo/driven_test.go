package rimo_test

import (
	"log"
	"math"
	"testing"

	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/hexops/valast"
)

// TESTS

func TestTestInterface(t *testing.T) {
	t.Parallel()

	var _ rimo.Reader = (*TestReader)(nil)

	var _ rimo.Writer = (*LogWriter)(nil)
}

// Note : numeric value should be converted to float64.
func TestPipeline(t *testing.T) {
	t.Parallel()

	// Set up TestReader
	baseName := "databaseName"
	tableNames := []string{"tableTest"}
	testInput := []colInput{
		{
			ColName:   "string",
			ColValues: []interface{}{"val1", "val2", "val3"},
		},
		{
			ColName:   "col2",
			ColValues: []interface{}{true, false, nil},
		},
		{
			ColName:   "col9",
			ColValues: []interface{}{float64(31), float64(29), float64(42)},
		},
	}

	testReader := TestReader{ //nolint:exhaustruct
		baseName:   baseName,
		tableNames: tableNames,
		data:       testInput,
		index:      0,
	}

	// LogWriter
	testWriter := LogWriter{}

	err := rimo.AnalyseBase(&testReader, &testWriter)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
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

func (r *TestReader) Value() ([]interface{}, string, string, error) {
	log.Printf("Processing %s column in %s table", r.currentTableName, r.currentColName)

	return r.currentValues, r.currentColName, r.currentTableName, nil
}

// LogWritter

type LogWriter struct{}

func (w *LogWriter) Export(base *model.Base) error {
	log.Printf("BASE returned \n \n : %s", valast.String(&base))

	return nil
}

// ReturnWriter : return the base object through GetBase() method.
type ReturnWriter struct {
	base model.Base
}

func (w *ReturnWriter) Export(base *model.Base) error {
	w.base = *base

	return nil
}

func (w *ReturnWriter) GetBase() model.Base {
	return w.base
}
