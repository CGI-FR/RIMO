package metric_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/require"
)

const (
	dataDir = "../../testdata/benchmark"
)

var result model.Base //nolint:gochecknoglobals // used in benchmark to avoid misleading compiler optimisation.

func BenchmarkMetric(b *testing.B) {
	listNumValues := []int{100, 1000, 10000}
	listType := []string{"numeric", "text", "bool"}

	for _, dataType := range listType {
		for _, numValues := range listNumValues {
			inputList := []string{fmt.Sprintf("%s/%s/%d_input.jsonl", dataDir, dataType, numValues)}

			b.Run(fmt.Sprintf("type= %s, numValues=%d", dataType, numValues), func(b *testing.B) {
				startTime := time.Now()

				base := model.Base{} //nolint:exhaustruct
				var err error

				for n := 0; n < b.N; n++ {
					base, err = analyse.Build(inputList)
					require.NoError(b, err)
				}

				result = base

				elapsed := time.Since(startTime)
				valuesPerSecond := float64(numValues*b.N) / elapsed.Seconds()
				b.ReportMetric(valuesPerSecond, "lines/s")
			})
		}
	}
}
