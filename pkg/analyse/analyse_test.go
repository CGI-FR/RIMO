package analyse_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	data1Path = "./testdata/data1/data_input.jsonl" //nolint:gochecknoglobals
	data2Path = "./testdata/data2/data_input.jsonl" //nolint:gochecknoglobals
	testPath  = "./testdata/test/data_input.jsonl"  //nolint:gochecknoglobals

)

// Execute Analyse pipeline and compare with expected result.
func TestAnalyse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		inputPath  string
		outputPath string
		testPath   string
	}{
		{
			name:       "data1",
			inputPath:  data1Path,
			outputPath: "./testdata/data1/data_output.yaml",
			testPath:   "./testdata/data1/data_expected.yaml",
		},
		{
			name:       "data2",
			inputPath:  data2Path,
			outputPath: "./testdata/data2/data_output.yaml",
			testPath:   "./testdata/data2/data_expected.yaml",
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testAnalyse(t, tc.inputPath, tc.outputPath, tc.testPath)
		})
	}
}

func testAnalyse(t *testing.T, inputPath string, outputPath string, testPath string) {
	t.Helper()

	inputList := []string{inputPath}
	err := analyse.Analyse(inputList, outputPath)
	require.NoError(t, err)

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

// Allow to quickly run analyse pipeline on testdata.
func TestRunAnalyse(t *testing.T) {
	t.Parallel()

	inputPath := []string{testPath}
	outputPath := "./testdata/test/data_output.yaml"

	err := analyse.Analyse(inputPath, outputPath)
	require.NoError(t, err)
}

// Benchmark Analyse pipeline.

func BenchmarkAnalyse(b *testing.B) {
	for _, numLines := range []int{100, 1000, 10000, 100000} {
		inputPath := fmt.Sprintf("./testdata/benchmark/mixed/%d_input.jsonl", numLines)
		inputList := []string{inputPath}
		outputPath := fmt.Sprintf("./testdata/benchmark/mixed/%d_output.yaml", numLines)

		b.Run(fmt.Sprintf("numLines=%d", numLines), func(b *testing.B) {
			b.ResetTimer()

			startTime := time.Now()

			for n := 0; n < b.N; n++ {
				err := analyse.Analyse(inputList, outputPath)
				require.NoError(b, err)
			}

			elapsed := time.Since(startTime)
			linesPerSecond := float64(numLines*b.N) / elapsed.Seconds()
			b.ReportMetric(linesPerSecond, "lines/s")
		})
	}
}

func BenchmarkMetric(b *testing.B) {
	listNumValues := []int{100, 1000, 10000}
	listType := []string{"numeric", "text", "bool"}

	for _, dataType := range listType {
		for _, numValues := range listNumValues {
			inputPath := fmt.Sprintf("./testdata/benchmark/%s/%d_input.jsonl", dataType, numValues)
			// Load inputFilePath.
			data, err := analyse.Load(inputPath)
			if err != nil {
				b.Fatalf("failed to load %s: %v", inputPath, err)
			}

			var cols []model.Column

			b.Run(fmt.Sprintf("type= %s, numValues=%d", dataType, numValues), func(b *testing.B) {
				b.ResetTimer()
				startTime := time.Now()

				for n := 0; n < b.N; n++ {
					cols = analyse.BuildColumnMetric(data, cols)
					require.NoError(b, err)
				}
				b.StopTimer()

				elapsed := time.Since(startTime)
				valuesPerSecond := float64(numValues*b.N) / elapsed.Seconds()
				b.ReportMetric(valuesPerSecond, "lines/s")
			})
		}
	}
}

func loadYAML(t *testing.T, path string) model.Base {
	t.Helper()

	// Load output file
	file, err := os.Open(path)
	require.NoError(t, err)

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
	require.NoError(t, err)

	var output string

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	require.NoError(t, err)
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
