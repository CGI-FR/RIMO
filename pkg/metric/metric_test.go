package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Ensure that 1. frequency is correct, 2. order is correct, 3. ties are break by length.
func TestStringMetric(t *testing.T) { //nolint:funlen
	t.Parallel()

	text := []string{"1", "1", "1", "1", "22", "22", "22", "331", "332", "4441", ""}

	min := ""
	max := "4441"

	expectedMetric := model.Column{ //nolint:exhaustruct
		MainMetric: model.Generic{
			Count:    12,
			Empty:    1,
			Null:     1,
			Distinct: 6,
			Samples:  []any{"22"},
			Min:      &min,
			Max:      &max,
		},
		StringMetric: &model.String{
			MinLen:   0,
			MaxLen:   4,
			CountLen: 5,
			Lengths: []model.StringLen{
				{
					Length: 1,
					Freq:   0.3333333333333333,
					Metrics: model.Generic{
						Count:    4,
						Empty:    0,
						Null:     0,
						Distinct: 1,
						Min:      &text[0],
						Max:      &text[0],
						Samples:  []any{"1", "1", "1", "1"},
					},
				},
			},
		},
	}

	actualMetric := model.Column{} //nolint:exhaustruct

	analyser := metric.NewString(5, true)
	for index := range text {
		analyser.Read(&text[index])
	}

	analyser.Read(nil)

	analyser.Build(&actualMetric)

	// out, err := yaml.Marshal(actualMetric)
	// assert.NoError(t, err)
	// fmt.Println(string(out))

	assert.Equal(t, expectedMetric.MainMetric.Count, actualMetric.MainMetric.Count)
	assert.Equal(t, expectedMetric.MainMetric.Empty, actualMetric.MainMetric.Empty)
	assert.Equal(t, expectedMetric.MainMetric.Null, actualMetric.MainMetric.Null)
	assert.Equal(t, expectedMetric.MainMetric.Distinct, actualMetric.MainMetric.Distinct)
	assert.Equal(t, expectedMetric.MainMetric.Min, actualMetric.MainMetric.Min)
	assert.Equal(t, expectedMetric.MainMetric.Max, actualMetric.MainMetric.Max)
	assert.Equal(t, expectedMetric.StringMetric.MinLen, actualMetric.StringMetric.MinLen)
	assert.Equal(t, expectedMetric.StringMetric.MaxLen, actualMetric.StringMetric.MaxLen)
	assert.Equal(t, expectedMetric.StringMetric.CountLen, actualMetric.StringMetric.CountLen)

	for i := 0; i < len(expectedMetric.StringMetric.Lengths); i++ {
		assert.Equal(t, expectedMetric.StringMetric.Lengths[i].Length, actualMetric.StringMetric.Lengths[i].Length)
		assert.Equal(t, expectedMetric.StringMetric.Lengths[i].Freq, actualMetric.StringMetric.Lengths[i].Freq)
		assert.Equal(t, expectedMetric.StringMetric.Lengths[i].Metrics.Samples, actualMetric.StringMetric.Lengths[i].Metrics.Samples) //nolint:lll
	}
}
