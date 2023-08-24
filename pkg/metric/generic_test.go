package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountEmpty(t *testing.T) {
	t.Parallel()

	slice := []interface{}{1, 2, 3, nil}
	expected := 1
	actual := metric.CountEmpty(slice)

	assert.Equal(t, expected, actual)
}

func TestColType(t *testing.T) {
	t.Parallel()

	t.Run("numeric", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, 2, 3}
		expected := model.ValueType.Numeric

		actual := metric.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, "text", nil}
		expected := model.ValueType.String

		actual := metric.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("boolean", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, true, false}
		expected := model.ValueType.Bool

		actual := metric.ColType(slice)
		require.Equal(t, expected, actual)
	})

	// Treat this case as error would imply to type assert each element of the slice when Loading.
	t.Run("mixed", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{"text", 2, false}
		expected := model.ValueType.String

		actual := metric.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("unknown", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, nil, nil}
		expected := model.ValueType.Undefined

		actual := metric.ColType(slice)
		require.Equal(t, expected, actual)
	})
}

// Implementation questions :
// should Unique() append nil element ?
// should CountUnique() count nil as a unique value ?

func TestUnique(t *testing.T) {
	t.Parallel()

	values := []interface{}{1, 1, 2, 3, nil}
	expected := []interface{}{1, 2, 3}
	actual := metric.Unique(values)

	assert.ElementsMatch(t, expected, actual)
}

func TestCountUnique(t *testing.T) {
	t.Parallel()

	values := []interface{}{1, 1, 2, 3, nil}
	expected := 3
	actual := metric.CountUnique(values)

	assert.Equal(t, expected, actual)
}

func TestSample(t *testing.T) {
	t.Parallel()

	values := []interface{}{1, 2, 3, nil, 5, 6}
	actualOutput, _ := metric.Sample(values, 5)

	assert.Len(t, actualOutput, 5)

	actualOutput, _ = metric.Sample(values, 10)
	assert.Len(t, actualOutput, 5)
}
