package analyse_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
)

var data1Path = "./testdata/data1/data_input.jsonl" //nolint:gochecknoglobals

// Execute Analyse pipeline and compare with expected result.
func TestAnalyseFileComparison(t *testing.T) {
	t.Parallel()

	inputList := []string{data1Path}
	outputPath := "./testdata/data1/data_output.yaml"
	testPath := "./testdata/data1/data_expected.yaml"

	err := analyse.Analyse(inputList, outputPath)
	assert.NoError(t, err)

	// Compare output file with expected output file.
	t.Run("output file comparison", func(t *testing.T) {
		t.Parallel()

		actualOutput := getText(t, outputPath)
		expectedOutput := getText(t, testPath)

		// Call removeSampleFromStrings
		actualOutput = removeSampleFromStrings(actualOutput)
		expectedOutput = removeSampleFromStrings(expectedOutput)

		// Compare the expected output and actual output
		assert.Equal(t, expectedOutput, actualOutput)
	})

	// Compare loaded output file with loaded expected output file.
	// EqualBase() is used to compare two model.Base.
	t.Run("loaded object comparison", func(t *testing.T) {
		t.Parallel()

		actualOutputBase := loadYAML(t, outputPath)
		expectedOutputBase := loadYAML(t, testPath)

		// Remove sample fields from both model.Base.
		actualOutputBase = removeSampleFromBase(actualOutputBase)
		expectedOutputBase = removeSampleFromBase(expectedOutputBase)

		// Compare the expected output and actual output except all sample fields.
		equal, diff := EqualBase(expectedOutputBase, actualOutputBase)
		if !equal {
			t.Errorf("base are not similar : %s", diff)
		}
	})
}

func loadYAML(t *testing.T, path string) model.Base {
	t.Helper()

	// Load output file
	file, err := os.Open(path)
	assert.NoError(t, err)

	decoder := yaml.NewDecoder(file)

	var base model.Base
	err = decoder.Decode(&base)

	if err != nil {
		t.Errorf("error while decoding yaml file: %v", err)
	}

	file.Close()

	return base
}

func getText(t *testing.T, outputPath string) string {
	t.Helper()

	file, err := os.Open(outputPath)
	assert.NoError(t, err)

	var output string

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	assert.NoError(t, err)
	file.Close()

	output = buf.String()

	return output
}

// UTILS FUNCTIONS.

// DeepEqual two model.Base.
func EqualBase(base1, base2 model.Base) (bool, string) {
	if !reflect.DeepEqual(base1, base2) {
		return false, fmt.Sprintf("base is different : %s \n \n %s", valast.String(base1), valast.String(base2))
	}

	return true, ""
}

func removeSampleFromBase(base model.Base) model.Base {
	for tableI, table := range base.Tables {
		for columnJ, column := range table.Columns {
			column.MainMetric.Sample = nil

			if column.Type == model.ValueType.String {
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

	return base
}

func removeSampleFromStrings(rimoString string) string {
	// Split at every new line
	lines := strings.Split(rimoString, "\n")

	// Filter out sample by skipping sampleSize + 1 lines when a line contain "sample" or "leastFrequentSample:"
	var filteredLines []string

	var skipLine int

	sampleSizeSkip := model.SampleSize + 1

	for _, line := range lines {
		// sample of stringMetric.MostFreqLen and stringMetric.LeastFreqLen may be of different length, skipping when nex
		if skipLine > 0 && strings.Contains(line, "   - length:") || strings.Contains(line, "    - name:") {
			skipLine = 0
		}

		switch {
		case skipLine > 0:
			skipLine--
		case strings.Contains(line, "sample:"):
			skipLine = sampleSizeSkip
		default:
			filteredLines = append(filteredLines, line)
		}
	}

	// Join the filtered lines back into a string
	rimoString = strings.Join(filteredLines, "\n")

	return rimoString
}

// TESTS .......

func TestGetBaseName(t *testing.T) {
	t.Helper()
	t.Parallel()

	path := "path/to/dir/basename_tablename.jsonl"
	expected := "basename"

	if baseName, err := analyse.GetBaseName(path); baseName != expected || err != nil {
		t.Errorf("GetBaseName(%q) = (%q, %v), expected (%q, %v)", path, baseName, err, expected, nil)
	}

	path2 := "basename_tablename.jsonl"
	expected2 := "basename"

	if baseName, err := analyse.GetBaseName(path2); baseName != expected2 || err != nil {
		t.Errorf("GetBaseName(%q) = (%q, %v), expected (%q, %v)", path2, baseName, err, expected2, nil)
	}

	invalidPath := ""

	_, err := analyse.GetBaseName(invalidPath)
	if !errors.Is(err, analyse.ErrNonExtractibleValue) {
		t.Errorf("expected error %v, but got %v", analyse.ErrNonExtractibleValue, err)
	}
}

func TestGetTableName(t *testing.T) {
	t.Helper()
	t.Parallel()

	path := "path/to/dir/basename_tablename.jsonl"
	expected := "tablename"

	if tableName, err := analyse.GetTableName(path); tableName != expected || err != nil {
		t.Errorf("GetTableName(%q) = (%q, %v), expected (%q, %v)", path, tableName, err, expected, nil)
	}

	path2 := "basename_tablename.jsonl"
	expected2 := "tablename"

	if tableName, err := analyse.GetTableName(path2); tableName != expected2 || err != nil {
		t.Errorf("GetTableName(%q) = (%q, %v), expected (%q, %v)", path2, tableName, err, expected2, nil)
	}

	invalidPath := ""

	_, err := analyse.GetTableName(invalidPath)
	if !errors.Is(err, analyse.ErrNonExtractibleValue) {
		t.Errorf("expected error %v, but got %v", analyse.ErrNonExtractibleValue, err)
	}
}
