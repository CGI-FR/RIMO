package analyse_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
)

const (
	TestDir = "../../test/analyseTest/"
)

var (
	JsonlNewFormat  = filepath.Join(TestDir, "/input/testcase_newstruct.jsonl")
	JsonlPrevFormat = filepath.Join(TestDir, "/input/testcase_prevstruct.jsonl")
)

// Compare output file with expected output file.
func TestAnalyseFileComparison(t *testing.T) {
	t.Parallel()

	inputList := []string{JsonlNewFormat}
	outputPath := filepath.Join(TestDir, "/output/rimo_output.yaml")
	analyse.Analyse(inputList, outputPath)

	// Load output file
	file, err := os.Open(outputPath)
	assert.NoError(t, err)

	var actualOutput string

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	assert.NoError(t, err)

	actualOutput = buf.String()

	// Load expected output file
	testPath := filepath.Join(TestDir, "/expected/rimo_output.yaml")
	expectedFile, err := os.Open(testPath)
	assert.NoError(t, err)

	t.Cleanup(func() {
		file.Close()
		expectedFile.Close()
	})

	var expectedOutput string

	buf = new(bytes.Buffer)
	_, err = buf.ReadFrom(expectedFile)
	assert.NoError(t, err)

	expectedOutput = buf.String()

	// Call removeSampleFromStrings
	actualOutput = removeSampleFromStrings(actualOutput)
	expectedOutput = removeSampleFromStrings(expectedOutput)

	// Compare the expected output and actual output
	assert.Equal(t, expectedOutput, actualOutput)
}

// Compare loaded output file with loaded expected output file.
// EqualBase() is used to compare two model.Base.
func TestAnalyseObjectComparison(t *testing.T) {
	t.Parallel()

	inputList := []string{JsonlNewFormat}
	outputPath := filepath.Join(TestDir, "/output/rimo_output.yaml")
	analyse.Analyse(inputList, outputPath)

	// Load output file
	file, err := os.Open(outputPath)
	assert.NoError(t, err)

	// Load expected output file
	testPath := filepath.Join(TestDir, "/expected/rimo_output.yaml")
	expectedFile, err := os.Open(testPath)
	assert.NoError(t, err)

	t.Cleanup(func() {
		file.Close()
		expectedFile.Close()
	})

	// Load file in a model.Base.
	decoder := yaml.NewDecoder(file)

	var actualOutputBase model.Base
	err = decoder.Decode(&actualOutputBase)

	if err != nil {
		t.Errorf("error while decoding yaml file: %v", err)
	}

	// Load expected file in a model.Base.
	decoder = yaml.NewDecoder(expectedFile)

	var expectedOutputBase model.Base
	err = decoder.Decode(&expectedOutputBase)

	if err != nil {
		t.Errorf("error while decoding yaml file: %v", err)
	}

	// Remove sample fields from both model.Base.
	actualOutputBase = removeSampleFromBase(actualOutputBase)
	expectedOutputBase = removeSampleFromBase(expectedOutputBase)

	// Compare the expected output and actual output except all sample fields.
	equal, diff := EqualBase(expectedOutputBase, actualOutputBase)
	if !equal {
		t.Errorf("base are not similar : %s", diff)
	}
}

// UTILS FUNCTIONS.

// DeepEqual two model.Base.
func EqualBase(base1, base2 model.Base) (equal bool, diff string) {
	if !reflect.DeepEqual(base1, base2) {
		return false, fmt.Sprintf("base is different : %s \n \n %s", valast.String(base1), valast.String(base2))
	} else {
		return true, ""
	}
}

// NOT IN USE : order of table and column in yaml is ensured by analyse.Load()
func EqualBaseWithoutOrder(base1, base2 model.Base) (equal bool, diff string) {
	if base1.Name != base2.Name {
		return false, fmt.Sprintf("base name does not match : %s and %s", base1.Name, base2.Name)
	}

	// Use case : order of tables and columns are not ensured.
	// We need to map table and column to their respective index to be able to compare them.
	type (
		tableLocator  map[string]int            // Index of table.
		columnLocator map[string]map[string]int // Index of column.
	)

	// base1
	tableLocator1 := make(tableLocator)
	columnLocator1 := make(columnLocator)

	for tableIndex, table := range base1.Tables {
		tableLocator1[table.Name] = tableIndex

		for columnIndex, column := range table.Columns {
			if columnLocator1[table.Name] == nil {
				columnLocator1[table.Name] = make(map[string]int)
			}

			columnLocator1[table.Name][column.Name] = columnIndex
		}
	}

	// base2
	tableLocator2 := make(tableLocator)
	columnLocator2 := make(columnLocator)

	for tableIndex, table := range base2.Tables {
		tableLocator2[table.Name] = tableIndex

		for columnIndex, column := range table.Columns {
			if columnLocator2[table.Name] == nil {
				columnLocator2[table.Name] = make(map[string]int)
			}

			columnLocator2[table.Name][column.Name] = columnIndex
		}
	}

	// compare tables and columns of base1 and base2.
	for tableName, columnLocator := range columnLocator1 {
		if _, ok := columnLocator2[tableName]; !ok {
			return false, fmt.Sprintf("table %s does not exist in Base", tableName)
		}

		for columnName := range columnLocator {
			if _, ok := columnLocator2[tableName][columnName]; !ok {
				return false, fmt.Sprintf("column %s does not exist in table %s", columnName, tableName)
			}
		}
	}

	// A second loop is required to check if all columns of base2 are present in base1.
	for tableName, columnLocator := range columnLocator2 {
		if _, ok := columnLocator1[tableName]; !ok {
			return false, fmt.Sprintf("table %s does not exist in Base", tableName)
		}

		for columnName := range columnLocator {
			if _, ok := columnLocator1[tableName][columnName]; !ok {
				return false, fmt.Sprintf("column %s does not exist in table %s", columnName, tableName)
			}
		}
	}

	// DeepEqual columns of base1 and base2
	for tableName, tableIndex := range tableLocator1 {
		table1 := base1.Tables[tableIndex]
		table2 := base2.Tables[tableLocator2[tableName]]

		for columnName, columnIndex := range columnLocator1[tableName] {
			column1 := table1.Columns[columnIndex]
			column2 := table2.Columns[columnLocator2[tableName][columnName]]

			if !reflect.DeepEqual(column1, column2) {
				return false, fmt.Sprintf(
					"column %s in table %s is different : %s \n %s",
					columnName, tableName, valast.String(column1), valast.String(column2))
			}
		}
	}

	return true, ""
}

func removeSampleFromBase(base model.Base) model.Base {
	for i, table := range base.Tables {
		for j, column := range table.Columns {
			column.MainMetric.Sample = nil

			if column.Type == "string" {
				column.StringMetric.LeastFreqSample = nil
			}
			base.Tables[i].Columns[j] = column
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
		switch {
		case skipLine > 0:
			skipLine--
		case strings.Contains(line, "sample:") || strings.Contains(line, "leastFrequentSample:"):
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
