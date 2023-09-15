package rimo_test

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/rimo"

	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Run Analyse pipeline with FilesReader and TestWriter and compare with expected result.
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

// PIPELINE TESTS

// Note : numeric value should be converted to float64.
func TestManualPipeline(t *testing.T) {
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
		{
			ColName:   "empty",
			ColValues: []interface{}{nil, nil, nil},
		},
	}

	testReader := TestReader{ //nolint:exhaustruct
		baseName:   baseName,
		tableNames: tableNames,
		data:       testInput,
		index:      0,
	}

	testWriter := TestWriter{} //nolint:exhaustruct

	err := rimo.AnalyseBase(&testReader, &testWriter)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Logf("Base returned : %s", valast.String(*testWriter.Base()))
}

// Ensure that the pipeline produce the same base as expected.
func TestPipeline(t *testing.T) {
	t.Parallel()

	testCases := []testCase{}
	testCases = append(testCases, getTestCase("../../testdata/data1/"))
	testCases = append(testCases, getTestCase("../../testdata/data2/"))

	for _, testCase := range testCases {
		testCase := testCase // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// Actual base

			reader, err := infra.FilesReaderFactory([]string{testCase.inputPath})
			assert.NoError(t, err)

			writer := &TestWriter{} //nolint:exhaustruct

			err = rimo.AnalyseBase(reader, writer)
			assert.NoError(t, err)

			actualBase := writer.Base()

			// Expected base
			expectedBase, err := model.LoadBase(testCase.expectedPath)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			// Remove sample
			model.RemoveSampleFromBase(expectedBase)
			model.RemoveSampleFromBase(actualBase)
			// Compare
			equal, diff := model.SameBase(expectedBase, actualBase)
			if !equal {
				t.Errorf("Base are not equal:\n%s", diff)
			}
		})
	}
}

// Benchmark (same as previous analyse_test.go benchmark).
func BenchmarkAnalyseInterface(b *testing.B) {
	for _, numLines := range []int{100, 1000, 10000, 100000} {
		inputPath := filepath.Join(dataDir, fmt.Sprintf("benchmark/mixed/%d_input.jsonl", numLines))
		inputList := []string{inputPath}
		outputPath := filepath.Join(dataDir, fmt.Sprintf("benchmark/mixed/%dinterface_output.yaml", numLines))

		b.Run(fmt.Sprintf("numLines=%d", numLines), func(b *testing.B) {
			startTime := time.Now()

			reader, err := infra.FilesReaderFactory(inputList)
			require.NoError(b, err)

			writer, err := infra.YAMLWriterFactory(outputPath)
			require.NoError(b, err)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err := rimo.AnalyseBase(reader, writer)
				require.NoError(b, err)
			}
			b.StopTimer()

			elapsed := time.Since(startTime)
			linesPerSecond := float64(numLines*b.N) / elapsed.Seconds()
			b.ReportMetric(linesPerSecond, "lines/s")
		})
	}
}
