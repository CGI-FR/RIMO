package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Ensure that 1. frequency is correct, 2. order is correct, 3. ties are break by length.
func TestStringMetric(t *testing.T) {
	t.Parallel()

	text := []interface{}{"1", "1", "1", "1", "22", "22", "22", "331", "332", "4441"}
	expectedMetric := model.StringMetric{
		MostFreqLen: []model.LenFreq{
			{Length: 1, Freq: 0.4, Sample: []string{"1"}},
			{Length: 2, Freq: 0.3, Sample: []string{"22"}},
		},
		LeastFreqLen: []model.LenFreq{
			{Length: 4, Freq: 0.1, Sample: []string{"4441"}},
			{Length: 3, Freq: 0.2, Sample: []string{"331", "332"}},
		},
	}

	actualMetric := model.StringMetric{} //nolint:exhaustruct

	err := metric.SetStringMetric(text, &actualMetric)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// t.Logf(valast.String(actualMetric))

	for i := 0; i < len(expectedMetric.MostFreqLen); i++ {
		assert.Equal(t, expectedMetric.MostFreqLen[i].Length, actualMetric.MostFreqLen[i].Length)
		assert.Equal(t, expectedMetric.MostFreqLen[i].Freq, actualMetric.MostFreqLen[i].Freq)
		assert.Equal(t, expectedMetric.MostFreqLen[i].Sample, actualMetric.MostFreqLen[i].Sample)
	}

	for i := 0; i < len(expectedMetric.LeastFreqLen); i++ {
		assert.Equal(t, expectedMetric.LeastFreqLen[i].Length, actualMetric.LeastFreqLen[i].Length)
		assert.Equal(t, expectedMetric.LeastFreqLen[i].Freq, actualMetric.LeastFreqLen[i].Freq)
		assert.ElementsMatch(t, expectedMetric.LeastFreqLen[i].Sample, actualMetric.LeastFreqLen[i].Sample)
	}
}
