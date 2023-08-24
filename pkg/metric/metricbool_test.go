package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBooleanMetric(t *testing.T) {
	t.Parallel()

	values := []interface{}{true, true, false, false}
	expectedMetric := model.BoolMetric{
		TrueRatio: float64(1) / float64(2),
	}

	actualMetric := model.BoolMetric{} //nolint:exhaustruct
	err := metric.SetBoolMetric(values, &actualMetric)
	require.NoError(t, err)

	assert.Equal(t, expectedMetric, actualMetric)
}
