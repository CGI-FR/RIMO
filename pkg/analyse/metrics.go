package analyse

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/cgi-fr/rimo/pkg/model"
)

// For a given dataCol return a model.Column with metrics.
func ComputeMetric(colName string, values []interface{}) (model.Column, error) {
	// Main metric
	name := colName
	colType := ColType(values)
	concept := ""
	var confidential *bool = nil //nolint

	// Create the column.
	col := model.Column{ //nolint:exhaustruct
		Name:         name,
		Type:         colType,
		Concept:      concept,
		Constraint:   []string{},
		Confidential: confidential,
	}

	// Generic metric
	sample, err := Sample(values, model.SampleSize)
	if err != nil {
		return model.Column{}, fmt.Errorf("error computing sample in column %v : %w", name, err)
	}

	genericMetric := model.GenericMetric{
		Count:  int64(len(values)),
		Unique: int64(len(values)),
		Sample: sample,
	}

	col.MainMetric = genericMetric

	// Type specific metric

	switch colType {
	case model.ValueType.String:
		metric, err := StringMetric(values)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing string metric in column %v : %w", name, err)
		}

		col.StringMetric = metric

	case model.ValueType.Numeric:
		metric, err := NumericMetric(values)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing numeric metric in column %v : %w", name, err)
		}

		col.NumericMetric = metric

	case model.ValueType.Bool:
		metric, err := BoolMetric(values)
		if err != nil {
			return model.Column{}, err
		}

		col.BoolMetric = metric
	}

	return col, nil
}

// Generic metrics.

func ColType(values []interface{}) model.RIMOType {
	colType := model.ValueType.Undefined
	for i := 0; i < len(values) && colType == model.ValueType.Undefined; i++ {
		colType = ValueType(values[i])
	}

	return colType
}

func Sample[T any](values []T, sampleSize int) ([]T, error) {
	if len(values) == 0 {
		return nil, errors.New("values slice is empty")
	}

	if sampleSize >= len(values) {
		return values, nil
	}

	sample := make([]T, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = values[rand.Intn(len(values)-1)] //nolint:gosec
	}

	return sample, nil
}

func Unique(values []interface{}) int64 {
	unique := make(map[interface{}]bool)

	for _, value := range values {
		unique[value] = true
	}

	return int64(len(unique))
}

// Specific type metric.

var ErrValueType = errors.New("value type error")

// String metric : MostFreqLen, LeastFreqLen, LeastFreqSample

func StringMetric(values []interface{}) (model.StringMetric, error) {
	metric := model.StringMetric{} //nolint:exhaustruct

	// Store strings by length.
	lenMap := make(map[int][]string)
	// Count length occurrence.
	lenCounter := make(map[int]int)
	totalCount := len(values)

	for _, value := range values {
		stringValue, ok := value.(string)
		if !ok {
			return metric, fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
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
	lenMostFreqLen := sortedLength[0:leastFrequentLenSize]

	lenLeastFreqLen := make([]int, mostFrequentLenSize)

	for i := 0; i < leastFrequentLenSize; i++ {
		index := len(sortedLength) - 1 - i
		lenLeastFreqLen[i] = sortedLength[index]
	}

	leastFreqLen, err := buildFreqLen(lenLeastFreqLen, lenMap, lenCounter, totalCount)
	if err != nil {
		return metric, fmt.Errorf("error building least frequent length : %w", err)
	}

	metric.LeastFreqLen = leastFreqLen

	mostFreqLen, err := buildFreqLen(lenMostFreqLen, lenMap, lenCounter, totalCount)
	if err != nil {
		return metric, fmt.Errorf("error building most frequent length : %w", err)
	}

	metric.MostFreqLen = mostFreqLen

	return metric, nil
}

func buildFreqLen(leastFreqLen []int, lenMap map[int][]string, lenCounter map[int]int, totalCount int) ([]model.LenFreq, error) { //nolint:lll
	lenFreqs := make([]model.LenFreq, len(leastFreqLen))

	for index, len := range leastFreqLen {
		sample, err := Sample(lenMap[len], model.LeastFrequentSampleSize)
		if err != nil {
			return lenFreqs, fmt.Errorf("error getting sample for length %v : %w", len, err)
		}

		lenFreqs[index] = model.LenFreq{
			Length: len,
			Freq:   GetFrequency(lenCounter[len], int64(totalCount)),
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
		mostFrequentLenSize = int(math.Round(float64(nunique / 2)))  //nolint:gomnd
		leastFrequentLenSize = int(math.Round(float64(nunique / 2))) //nolint:gomnd
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

// Numeric metric : Min, Max, Mean.

func NumericMetric(values []interface{}) (model.NumericMetric, error) {
	totalCount := len(values)

	value, ok := values[0].(float64)
	if !ok {
		return model.NumericMetric{}, fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, values[0], values[0])
	}

	min := value
	max := value
	sum := 0.0

	for _, value := range values {
		valueFloat, ok := value.(float64)
		if !ok {
			return model.NumericMetric{}, fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
		}

		sum += valueFloat

		if valueFloat > max {
			max = valueFloat
		}

		if valueFloat < min {
			min = valueFloat
		}
	}

	mean := sum / float64(totalCount)

	numericMetric := model.NumericMetric{
		Min:  min,
		Max:  max,
		Mean: mean,
	}

	return numericMetric, nil
}

// Bool metric : TrueRatio.
func BoolMetric(values []interface{}) (model.BoolMetric, error) {
	totalCount := len(values)
	trueCount := 0

	for _, value := range values {
		switch valueBool := value.(type) {
		case bool:
			if valueBool {
				trueCount++
			}
		default:
			return model.BoolMetric{}, fmt.Errorf("%w : expected boolean found %T: %v", ErrValueType, value, value)
		}
	}

	trueRatio := GetFrequency(trueCount, int64(totalCount))

	boolMetric := model.BoolMetric{
		TrueRatio: trueRatio,
	}

	return boolMetric, nil
}

// Utils functions.

func GetFrequency(occurrence int, count int64) float64 {
	return float64(occurrence) / float64(count)
}

func ValueType(value interface{}) model.RIMOType {
	switch value.(type) {
	case int:
		return model.ValueType.Numeric
	case float64:
		return model.ValueType.Numeric
	case json.Number:
		return model.ValueType.Numeric
	case string:
		return model.ValueType.String
	case bool:
		return model.ValueType.Bool
	default:
		return model.ValueType.Undefined
	}
}
