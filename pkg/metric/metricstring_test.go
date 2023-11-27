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

package metric_test

import (
	"encoding/json"
	"fmt"
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
		MinLen:       1,
		MaxLen:       4,
		MostFreqLen:  []model.LenFreq{{Length: 1, Freq: 0.4, Sample: []string{"1"}}, {Length: 2, Freq: 0.3, Sample: []string{"22"}}},            //nolint:lll
		LeastFreqLen: []model.LenFreq{{Length: 4, Freq: 0.1, Sample: []string{"4441"}}, {Length: 3, Freq: 0.2, Sample: []string{"331", "332"}}}, //nolint:lll
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

func TestStringMetricV2(t *testing.T) {
	analyser := metric.NewString(5)

	strings := []string{"1", "1", "1", "1", "22", "22", "22", "331", "332", "4441", ""}

	for _, s := range strings {
		s := s
		analyser.Read(&s)
	}

	analyser.Read(nil)

	bytes, err := json.Marshal(analyser.Build())

	assert.NoError(t, err)

	fmt.Printf("%s\n", string(bytes))
}
