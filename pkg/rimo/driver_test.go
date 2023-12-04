package rimo_test

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/stretchr/testify/require"
)

const (
	dataDir      = "../../testdata/"
	inputName    = "data_input.jsonl"
	outputName   = "interface_data_output.yaml"
	expectedName = "data_expected.yaml"
)

// Benchmark (same as previous analyse_test.go benchmark).
func BenchmarkAnalyseInterface(b *testing.B) {
	for _, numLines := range []int{100, 1000, 10000, 100000} {
		inputPath := filepath.Join(dataDir, fmt.Sprintf("benchmark/mixed/%d", numLines))
		outputPath := filepath.Join(dataDir, fmt.Sprintf("benchmark/mixed/%dinterface_output.yaml", numLines))

		b.Run(fmt.Sprintf("numLines=%d", numLines), func(b *testing.B) {
			startTime := time.Now()

			reader, err := infra.NewJSONLFolderReader(inputPath)
			require.NoError(b, err)

			writer, err := infra.YAMLWriterFactory(outputPath)
			require.NoError(b, err)

			driver := rimo.Driver{SampleSize: 5, Distinct: true}

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err := driver.AnalyseBase(reader, writer)
				require.NoError(b, err)
			}
			b.StopTimer()

			elapsed := time.Since(startTime)
			linesPerSecond := float64(numLines*b.N) / elapsed.Seconds()
			b.ReportMetric(linesPerSecond, "lines/s")
		})
	}
}
