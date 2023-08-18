package analyse_test

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
)

var (
	data1Path = "./testdata/data1/data_input.jsonl" //nolint:gochecknoglobals
	data2Path = "./testdata/data2/data_input.jsonl" //nolint:gochecknoglobals
)

// Execute Analyse pipeline and compare with expected result.
func TestAnalyse(t *testing.T) {
	t.Parallel()

	inputList := []string{data2Path}
	outputPath := "./testdata/data2/data_output.yaml"
	testPath := "./testdata/data2/data_expected.yaml"

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

// Benchmark Analyse pipeline.

func BenchmarkAnalyse(b *testing.B) {
	jsonLine := `{"address": "PSC 4713, Box 9649 APO AA 43433", "age": 29, "major": false}`

	for _, numLines := range []int{100, 1000, 10000, 100000} {
		filepath := fmt.Sprintf("./testdata/benchmark/%d_input.jsonl", numLines)
		outputPath := fmt.Sprintf("./testdata/benchmark/%d_output.yaml", numLines)

		// Create a file with n lines.
		err := createBenchFile(filepath, numLines, jsonLine)
		assert.NoError(b, err)

		inputList := []string{filepath}

		b.Run(fmt.Sprintf("numLines=%d", numLines), func(b *testing.B) {
			b.ResetTimer()

			startTime := time.Now()

			for n := 0; n < b.N; n++ {
				err := analyse.Analyse(inputList, outputPath)
				assert.NoError(b, err)
			}

			elapsed := time.Since(startTime)
			linesPerSecond := float64(numLines*b.N) / elapsed.Seconds()
			b.ReportMetric(linesPerSecond, "lines/s")
		})
	}
}

// Create a JSON Line for benchmark.
func createBenchFile(path string, lines int, jsonLine string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("error while creating directory: %w", err)
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error while closing file: %v", err)
		}
	}()

	for i := 0; i < lines; i++ {
		_, err := file.WriteString(jsonLine + "\n")
		if err != nil {
			return fmt.Errorf("error while writing to file: %w", err)
		}
	}

	if err := file.Sync(); err != nil {
		return fmt.Errorf("error while syncing file: %w", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error while getting file info: %w", err)
	}

	if fileInfo.Size() != int64(lines*(len(jsonLine)+1)) {
		return fmt.Errorf("file size does not match expected size")
	}

	return nil
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
