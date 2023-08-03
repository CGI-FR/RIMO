package analyse_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	t.Parallel()
	t.Helper()

	// Create a temporary directory for the test
	tempDir, err := ioutil.TempDir("", "export_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a temporary file for the output
	outputFile := filepath.Join(tempDir, "output.yaml")

	// Create a base for the test
	column := model.Column{
		Name:         "Test Column",
		Type:         "string",
		Concept:      "Test Concept",
		Constraint:   []string{"Test Constraint 1", "Test Constraint 2"},
		Confidential: true,
		MainMetric: model.GenericMetric{
			Count:  10,
			Unique: 5,
			Sample: []interface{}{"Test Value 1", "Test Value 2", "Test Value 3", "Test Value 4", "Test Value 5"},
		},
		StringMetric: model.StringMetric{
			MostFreqLen:     map[int]float64{5: 0.5},
			LeastFreqLen:    map[int]float64{10: 0.5},
			LeastFreqSample: []string{"Test Value 1", "Test Value 2"},
		},
		NumericMetric: model.NumericMetric{
			Min:  0,
			Max:  100,
			Mean: 50,
		},
		BoolMetric: model.BoolMetric{
			TrueRatio: 0.6,
		},
	}

	base := model.Base{
		Name: "Test Database",
		Tables: []struct {
			Name    string         `yaml:"name"`
			Columns []model.Column `yaml:"columns"`
		}{
			{
				Name: "Test Table",
				Columns: []model.Column{
					column,
				},
			},
		},
	}
	// Export the base to the output file
	err = analyse.Export(base, outputFile)
	assert.NoError(t, err)

	// Read the output file and check its contents
	outputData, err := ioutil.ReadFile(outputFile)
	assert.NoError(t, err)

	expectedData := `name: Test Base
tables:
- name: Test Table
  columns:
  - name: id
    type: integer
  - name: name
    type: string
`
	assert.Equal(t, expectedData, string(outputData))
}
