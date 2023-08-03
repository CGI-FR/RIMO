package analyse_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColType(t *testing.T) { //nolint:funlen
	t.Parallel()

	t.Run("numeric", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{1, 2, 3}
		expected := "numeric"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("numeric with nil", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{nil, 2, 3}
		expected := "numeric"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("string", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{nil, "text", nil}
		expected := "string"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("boolean", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{nil, true, false}
		expected := "boolean"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("mixed", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{"text", 2, false}
		expected := "string"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("unknown", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{nil, nil, nil}
		expected := "unknown"

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})
}

func TestSample(t *testing.T) {
	t.Helper()
	t.Parallel()

	slice1 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sample1 := analyse.Sample(slice1, 5)
	sample2 := analyse.Sample(slice1, 5)

	t.Run("Sample len", func(t *testing.T) {
		t.Parallel()

		if len(sample1) != 5 {
			t.Errorf("analyse.Sample(%v, 5) = %v; expected %v", slice1, sample1, 5)
		}
	})

	t.Run("Sample is random", func(t *testing.T) {
		t.Parallel()

		sameOrder := 0

		for i := 0; i < len(sample1); i++ {
			if sample1[i] == sample2[i] {
				sameOrder++
			}
		}

		if sameOrder == len(sample1) {
			t.Errorf("2 analyse.Sample(%v, 5) have same order; most likely expected different", slice1)
		}
	})

	t.Run("Sample len greater than input len", func(t *testing.T) {
		t.Parallel()

		sample3 := analyse.Sample(slice1, 15)
		if len(sample3) != 15 {
			t.Errorf("analyse.Sample(%v, 15) = %v; expected %v", slice1, sample3, 15)
		}
	})
}

func TestLenCouter(t *testing.T) {
	t.Parallel()

	t.Run("valid input", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{"Hello", "Hello", "Hi", ""}
		expected := map[int]int{5: 2, 2: 1, 0: 1}

		actual, err := analyse.LenCounter(slice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("invalid input", func(t *testing.T) {
		t.Helper()
		t.Parallel()

		slice := []interface{}{"Hello", 2, true}

		_, err := analyse.LenCounter(slice)
		assert.ErrorIs(t, err, analyse.ErrNonString)
	})
}
