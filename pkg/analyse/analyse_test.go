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

	"github.com/cgi-fr/rimo/pkg/analyse" //nolint:depguard
	"github.com/cgi-fr/rimo/pkg/model"   //nolint:depguard
	"github.com/hexops/valast"           //nolint:depguard
	"github.com/stretchr/testify/assert" //nolint:depguard
)

const (
	TestDir = "../../test/analyseTest/"
)

var (
	jsonlNewFormat  = filepath.Join(TestDir, "/input/testcase_newstruct.jsonl")  //nolint:gochecknoglobals
	jsonlPrevFormat = filepath.Join(TestDir, "/input/testcase_prevstruct.jsonl") //nolint:gochecknoglobals
)

// Compare output file with expected output file.
func TestAnalyseFileComparison(t *testing.T) {
	t.Parallel()

	inputList := []string{jsonlNewFormat}
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

	inputList := []string{jsonlNewFormat}
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
				column.StringMetric.LeastFreqSample = nil
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
