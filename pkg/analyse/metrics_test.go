// Copyright (C) 2023 CGI France
//
// This file is part of RIMO.
//
// RIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// RIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with RIMO.  If not, see <http://www.gnu.org/licenses/>.

package analyse_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Main metrics tests.

func TestColType(t *testing.T) { //nolint:funlen
	t.Parallel()

	t.Run("numeric", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{1, 2, 3}
		expected := model.ValueType.Numeric

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("numeric with nil", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, 2, 3}
		expected := model.ValueType.Numeric

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, "text", nil}
		expected := model.ValueType.String

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("boolean", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, true, false}
		expected := model.ValueType.Bool

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("mixed", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{"text", 2, false}
		expected := model.ValueType.String

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})

	t.Run("unknown", func(t *testing.T) {
		t.Parallel()

		slice := []interface{}{nil, nil, nil}
		expected := model.ValueType.Undefined

		actual := analyse.ColType(slice)
		require.Equal(t, expected, actual)
	})
}

func TestSample(t *testing.T) {
	t.Parallel()

	slice1 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sample1, _ := analyse.Sample(slice1, 5)
	sample2, _ := analyse.Sample(slice1, 5)

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

		sample3, _ := analyse.Sample(slice1, 15)
		if len(sample3) != len(slice1) {
			t.Errorf("analyse.Sample(%v, 15) = %v; expected len to be %v", slice1, sample3, len(slice1))
		}
	})
}

func TestCountUnique(t *testing.T) {
	t.Parallel()

	sample := []interface{}{1, 1, 2, 3}
	expected := int64(3)
	actual := analyse.CountUnique(sample)

	assert.Equal(t, expected, actual)
}

func TestUnique(t *testing.T) {
	t.Parallel()

	values := []interface{}{1, 1, 2, 3}
	expected := []interface{}{1, 2, 3}
	actual := analyse.Unique(values)

	assert.ElementsMatch(t, expected, actual)
}

// Metrics tests.

func TestNumericMetric(t *testing.T) {
	t.Parallel()

	t.Run("numeric metric", func(t *testing.T) {
		t.Parallel()

		numeric := []interface{}{1.0, 2.0, 3.0}
		expectedMetric := model.NumericMetric{
			Min:  1,
			Max:  3,
			Mean: 2,
		}

		actualMetric, err := analyse.NumericMetric(numeric)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, expectedMetric, actualMetric)
	})
}

// Ensure that 1. frequency is correct, 2. order goes from least frequent and most frequent.
// and 3. ties are break by length.
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

	actualMetric, err := analyse.StringMetric(text)
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
	}
}

func TestBooleanMetric(t *testing.T) {
	t.Parallel()

	boolean := []interface{}{true, true, false}
	expectedMetric := model.BoolMetric{
		TrueRatio: float64(2) / float64(3),
	}

	boolMetric, _ := analyse.BoolMetric(boolean)
	assert.Equal(t, expectedMetric, boolMetric)
}
