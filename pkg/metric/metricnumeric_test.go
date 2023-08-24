package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNumericMetric(t *testing.T) {
	t.Parallel()

	values := []interface{}{1.0, 2.0, 3.0, nil}
	expectedMetric := model.NumericMetric{
		Min:  1,
		Max:  3,
		Mean: 2,
	}

	actualMetric := model.NumericMetric{} //nolint:exhaustruct

	err := metric.SetNumericMetric(values, &actualMetric)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, expectedMetric, actualMetric)
}
