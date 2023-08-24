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
	"time"

	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/io"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dataDir      = "../../testdata/"
	inputName    = "data_input.jsonl"
	outputName   = "data_output.yaml"
	expectedName = "data_expected.yaml"
)

type testCase struct {
	name         string
	inputPath    string
	outputPath   string
	expectedPath string
}

func getTestCase(dataFolder string) testCase {
	return testCase{
		name:         filepath.Base(dataFolder),
		inputPath:    filepath.Join(dataFolder, inputName),
		outputPath:   filepath.Join(dataFolder, outputName),
		expectedPath: filepath.Join(dataFolder, expectedName),
	}
}

// Execute Analyse pipeline and compare with expected result.
func TestAnalyse(t *testing.T) {
	t.Parallel()

	testCases := []testCase{}
	testCases = append(testCases, getTestCase("../../testdata/data1/"))
	testCases = append(testCases, getTestCase("../../testdata/data2/"))

	for _, testCase := range testCases {
		testCase := testCase // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			runAnalyse(t, testCase.inputPath, testCase.outputPath)
			compareFileOutput(t, testCase.outputPath, testCase.expectedPath)
			compareObjectOutput(t, testCase.outputPath, testCase.expectedPath)
		})
	}
}

func runAnalyse(t *testing.T, inputPath string, outputPath string) {
	t.Helper()

	inputList := []string{inputPath}

	base, err := analyse.Build(inputList)
	require.NoError(t, err)

	if outputPath != "" {
		err = io.Export(base, outputPath)
		require.NoError(t, err)
	}
}

func compareFileOutput(t *testing.T, outputPath string, testPath string) {
	t.Helper()

	actualOutput := getText(t, outputPath)
	expectedOutput := getText(t, testPath)

	// Call removeSampleFromStrings
	actualOutput = removeSampleFromStrings(actualOutput)
	expectedOutput = removeSampleFromStrings(expectedOutput)

	// Compare the expected output and actual output
	assert.Equal(t, expectedOutput, actualOutput)
}

func compareObjectOutput(t *testing.T, outputPath string, testPath string) {
	t.Helper()

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
				base, err := analyse.Build(inputList)
				require.NoError(b, err)

				err = io.Export(base, outputPath)
				require.NoError(b, err)
			}

			elapsed := time.Since(startTime)
			linesPerSecond := float64(numLines*b.N) / elapsed.Seconds()
			b.ReportMetric(linesPerSecond, "lines/s")
		})
	}
}

func TestExtractName(t *testing.T) {
	t.Parallel()

	path := "path/to/dir/basename_tablename.jsonl"
	expectedBase, expectedName := "basename", "tablename"
	actualBase, actualName, err := analyse.ExtractName(path)
	assert.NoError(t, err)

	assert.Equal(t, expectedBase, actualBase)
	assert.Equal(t, expectedName, actualName)

	path = "basename_tablename.jsonl"
	expectedBase, expectedName = "basename", "tablename"
	actualBase, actualName, err = analyse.ExtractName(path)
	assert.NoError(t, err)

	assert.Equal(t, expectedBase, actualBase)
	assert.Equal(t, expectedName, actualName)

	invalidPath := ""

	_, _, err = analyse.ExtractName(invalidPath)
	if !errors.Is(err, analyse.ErrNonExtractibleValue) {
		t.Errorf("expected error %v, but got %v", analyse.ErrNonExtractibleValue, err)
	}
}

func TestBaseIsUnique(t *testing.T) {
	t.Parallel()

	inputList := []string{
		"/data/somewhere/BASE_test.jsonl",
		"/data/somewhere/BASE3221_test.jsonl",
	}

	err := analyse.BaseIsUnique(inputList)
	assert.ErrorIs(t, err, analyse.ErrNonUniqueBase)
}

// Helper functions

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

// DeepEqual two model.Base.
func EqualBase(base1, base2 model.Base) (bool, string) {
	if !reflect.DeepEqual(base1, base2) {
		return false, fmt.Sprintf("base is different : %s \n \n %s", valast.String(base1), valast.String(base2))
	}

	return true, ""
}
