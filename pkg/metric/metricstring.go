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

	"github.com/cgi-fr/rimo/pkg/rimo"
)

func SetStringMetric(values []interface{}, metric *rimo.StringMetric) error {
	// Store strings by length.
	lenMap := make(map[int][]string)
	// Count length occurrence.
	lenCounter := make(map[int]int)
	totalCount := len(values)

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

	leastFreqLen, err := buildFreqLen(lenLeastFreqLen, lenMap, lenCounter, totalCount, rimo.LeastFrequentSampleSize)
	if err != nil {
		return fmt.Errorf("error building least frequent length : %w", err)
	}

	metric.LeastFreqLen = leastFreqLen

	mostFreqLen, err := buildFreqLen(lenMostFreqLen, lenMap, lenCounter, totalCount, rimo.MostFrequentSampleSize)
	if err != nil {
		return fmt.Errorf("error building most frequent length : %w", err)
	}

	metric.MostFreqLen = mostFreqLen

	return nil
}

func buildFreqLen(freqLen []int, lenMap map[int][]string, lenCounter map[int]int, totalCount int, sampleLen int) ([]rimo.LenFreq, error) { //nolint
	lenFreqs := make([]rimo.LenFreq, len(freqLen))

	for index, len := range freqLen {
		// Get unique value from lenMap[len]..
		sample, err := Sample(lenMap[len], sampleLen)
		if err != nil {
			return lenFreqs, fmt.Errorf("error getting sample for length %v : %w", len, err)
		}

		lenFreqs[index] = rimo.LenFreq{
			Length: len,
			Freq:   GetFrequency(lenCounter[len], totalCount),
			Sample: sample,
		}
	}

	return lenFreqs, nil
}

func getFreqSize(nunique int) (int, int) {
	mostFrequentLenSize := rimo.MostFrequentLenSize
	leastFrequentLenSize := rimo.LeastFrequentLenSize

	if nunique < rimo.MostFrequentLenSize+rimo.LeastFrequentLenSize {
		// Modify MostFrequentLenSize and LeastFrequentLenSize to fit the number of unique length.
		// Should keep ratio of MostFrequentLenSize and LeastFrequentLenSize.
		ratio := float64(rimo.MostFrequentLenSize) / float64(rimo.MostFrequentLenSize+rimo.LeastFrequentLenSize)
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
