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

package metric

import (
	"fmt"
	"math"
	"sort"

	"github.com/cgi-fr/rimo/pkg/model"
)

func SetStringMetric(values []interface{}, metric *model.StringMetric) error {
	// Store strings by length.
	lenMap := make(map[int][]string)
	// Count length occurrence.
	lenCounter := make(map[int]int)
	totalCount := len(values)

	metric.MinLen = math.MaxInt
	metric.MaxLen = 0

	for _, value := range values {
		if value == nil {
			continue
		}

		stringValue, ok := value.(string)
		if !ok {
			return fmt.Errorf("%w : expected string found %T: %v", ErrValueType, value, value)
		}

		length := len(stringValue)
		lenMap[length] = append(lenMap[length], stringValue)
		lenCounter[length]++

		metric.MinLen = min(metric.MinLen, length)
		metric.MaxLen = max(metric.MaxLen, length)
	}

	// Create a list of unique lengths sorted by descending frequency, break ties with ascending length
	sortedLength := uniqueLengthSorted(lenCounter)

	// Get size of MostFreqLen and LeastFreqLen
	mostFrequentLenSize, leastFrequentLenSize := getFreqSize(len(sortedLength))

	// Get ordered slice of least and most frequent length
	lenMostFreqLen := sortedLength[0:mostFrequentLenSize]

	lenLeastFreqLen := make([]int, leastFrequentLenSize)

	for i := 0; i < leastFrequentLenSize; i++ {
		index := len(sortedLength) - 1 - i
		lenLeastFreqLen[i] = sortedLength[index]
	}

	leastFreqLen, err := buildFreqLen(lenLeastFreqLen, lenMap, lenCounter, totalCount, model.LeastFrequentSampleSize)
	if err != nil {
		return fmt.Errorf("error building least frequent length : %w", err)
	}

	metric.LeastFreqLen = leastFreqLen

	mostFreqLen, err := buildFreqLen(lenMostFreqLen, lenMap, lenCounter, totalCount, model.MostFrequentSampleSize)
	if err != nil {
		return fmt.Errorf("error building most frequent length : %w", err)
	}

	metric.MostFreqLen = mostFreqLen

	return nil
}

func buildFreqLen(freqLen []int, lenMap map[int][]string, lenCounter map[int]int, totalCount int, sampleLen int) ([]model.LenFreq, error) { //nolint
	lenFreqs := make([]model.LenFreq, len(freqLen))

	for index, len := range freqLen {
		// Get unique value from lenMap[len]..
		sample, err := Sample(lenMap[len], sampleLen)
		if err != nil {
			return lenFreqs, fmt.Errorf("error getting sample for length %v : %w", len, err)
		}

		lenFreqs[index] = model.LenFreq{
			Length: len,
			Freq:   GetFrequency(lenCounter[len], totalCount),
			Sample: sample,
		}
	}

	return lenFreqs, nil
}

func getFreqSize(nunique int) (int, int) {
	mostFrequentLenSize := model.MostFrequentLenSize
	leastFrequentLenSize := model.LeastFrequentLenSize

	if nunique < model.MostFrequentLenSize+model.LeastFrequentLenSize {
		// Modify MostFrequentLenSize and LeastFrequentLenSize to fit the number of unique length.
		// Should keep ratio of MostFrequentLenSize and LeastFrequentLenSize.
		ratio := float64(model.MostFrequentLenSize) / float64(model.MostFrequentLenSize+model.LeastFrequentLenSize)
		mostFrequentLenSize = int(math.Round(float64(nunique) * ratio))
		leastFrequentLenSize = nunique - mostFrequentLenSize
	}

	return mostFrequentLenSize, leastFrequentLenSize
}

func uniqueLengthSorted(lenCounter map[int]int) []int {
	uniqueLengthSorted := make([]int, 0, len(lenCounter))
	for l := range lenCounter {
		uniqueLengthSorted = append(uniqueLengthSorted, l)
	}

	// Sort the string lengths by descending count of occurrence, breaks ties with ascending length
	sort.Slice(uniqueLengthSorted, func(i, j int) bool {
		if lenCounter[uniqueLengthSorted[i]] == lenCounter[uniqueLengthSorted[j]] {
			return uniqueLengthSorted[i] < uniqueLengthSorted[j]
		}

		return lenCounter[uniqueLengthSorted[i]] > lenCounter[uniqueLengthSorted[j]]
	})

	return uniqueLengthSorted
}

type String struct {
	sampleSize uint
	main       Stateless[string]
	byLen      map[int]Stateless[string]
}

func NewString(sampleSize uint) *String {
	return &String{
		sampleSize: sampleSize,
		main:       NewCounter[string](sampleSize),
		byLen:      map[int]Stateless[string]{},
	}
}

func (s *String) Read(value *string) {
	s.main.Read(value)

	if value != nil {
		length := len(*value)

		analyser, exists := s.byLen[length]
		if !exists {
			analyser = NewCounter[string](s.sampleSize)
		}

		analyser.Read(value)

		s.byLen[length] = analyser
	}
}

func (s *String) Build() model.Col[string] {
	result := model.Col[string]{}

	result.MainMetric.Count = s.main.CountTotal()
	result.MainMetric.Empty = s.main.CountEmpty()
	result.MainMetric.Null = s.main.CountNulls()
	result.MainMetric.Max = s.main.Max()
	result.MainMetric.Min = s.main.Min()
	result.MainMetric.Samples = s.main.Samples()

	lengths := make([]int, 0, len(s.byLen))
	for len := range s.byLen {
		lengths = append(lengths, len)
	}

	sort.Ints(lengths)

	result.StringMetric.CountLen = len(lengths)
	result.StringMetric.MaxLen = lengths[0]
	result.StringMetric.MaxLen = lengths[len(lengths)-1]

	for _, length := range lengths {
		len := model.StringLen{}
		len.Length = length
		len.Metrics.Count = s.byLen[length].CountTotal()
		len.Metrics.Empty = s.byLen[length].CountEmpty()
		len.Metrics.Null = s.byLen[length].CountNulls()
		len.Metrics.Max = s.byLen[length].Max()
		len.Metrics.Min = s.byLen[length].Min()
		len.Metrics.Samples = s.byLen[length].Samples()
		result.StringMetric.Lengths = append(result.StringMetric.Lengths, len)
	}

	return result
}
