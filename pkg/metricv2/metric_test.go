package metricv2_test

import (
	"fmt"
	"testing"

	"github.com/cgi-fr/rimo/pkg/metricv2"
	"github.com/cgi-fr/rimo/pkg/modelv2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// Ensure that 1. frequency is correct, 2. order is correct, 3. ties are break by length.
func TestStringMetric(t *testing.T) { //nolint:funlen
	t.Parallel()

	text := []string{"1", "1", "1", "1", "22", "22", "22", "331", "332", "4441", ""}

	min := ""
	max := "4441"

	expectedMetric := modelv2.Generic[string]{
		Count:    12,
		Empty:    1,
		Null:     1,
		Distinct: 6,
		Samples:  []string{"22"},
		Min:      &min,
		Max:      &max,
		String: &modelv2.String{
			MinLen:   0,
			MaxLen:   4,
			CountLen: 5,
			Lengths: []modelv2.StringLen{
				{
					Length: 1,
					Freq:   0.3333333333333333,
					Metrics: modelv2.Generic[string]{
						Count:    4,
						Empty:    0,
						Null:     0,
						Distinct: 1,
						Min:      &text[0],
						Max:      &text[0],
						Samples:  []string{"1", "1", "1", "1"},
						String:   nil,
					},
				},
			},
		},
	}

	actualMetric := modelv2.Generic[string]{}

	analyser := metricv2.NewString(5, true)
	for index := range text {
		analyser.Read(&text[index])
	}

	analyser.Read(nil)

	analyser.Build(&actualMetric)

	out, err := yaml.Marshal(actualMetric)

	assert.NoError(t, err)

	fmt.Println(string(out))

	assert.Equal(t, expectedMetric.Count, actualMetric.Count)
	assert.Equal(t, expectedMetric.Empty, actualMetric.Empty)
	assert.Equal(t, expectedMetric.Null, actualMetric.Null)
	assert.Equal(t, expectedMetric.Distinct, actualMetric.Distinct)
	assert.Equal(t, expectedMetric.Min, actualMetric.Min)
	assert.Equal(t, expectedMetric.Max, actualMetric.Max)
	assert.Equal(t, expectedMetric.String.MinLen, actualMetric.String.MinLen)
	assert.Equal(t, expectedMetric.String.MaxLen, actualMetric.String.MaxLen)
	assert.Equal(t, expectedMetric.String.CountLen, actualMetric.String.CountLen)

	for i := 0; i < len(expectedMetric.String.Lengths); i++ {
		assert.Equal(t, expectedMetric.String.Lengths[i].Length, actualMetric.String.Lengths[i].Length)
		assert.Equal(t, expectedMetric.String.Lengths[i].Freq, actualMetric.String.Lengths[i].Freq)
		assert.Equal(t, expectedMetric.String.Lengths[i].Metrics.Samples, actualMetric.String.Lengths[i].Metrics.Samples)
	}
}
