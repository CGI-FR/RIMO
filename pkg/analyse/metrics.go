package analyse

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/cgi-fr/rimo/pkg/model" //nolint:depguard
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

	genericMetric := model.GenericMetric{
		Count:  int64(len(values)),
		Unique: int64(len(values)),
		Sample: Sample(values, model.SampleSize),
	}

	col.MainMetric = genericMetric

	// Type specific metric

	switch colType {
	case "string":
		metric, err := StringMetric(values)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing string metric in column %v : %w", name, err)
		}

		col.StringMetric = metric

	case "numeric":
		metric, err := NumericMetric(values)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing numeric metric in column %v : %w", name, err)
		}

		col.NumericMetric = metric

	case "bool":
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

func Sample(values []interface{}, sampleSize int) []interface{} {
	sample := make([]interface{}, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = values[rand.Intn(len(values))] //nolint:gosec
	}

	return sample
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
	// Initialize the StringMetric struct
	metric := model.StringMetric{} //nolint:exhaustruct

	// Convert the input values to a slice of strings
	strings := make([]string, len(values))
	// Count the frequency of each string length
	lenCounter := make(map[int]int)

	for vIndex, value := range values {
		stringValue, ok := value.(string)
		if !ok {
			return metric, fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
		}

		strings[vIndex] = stringValue
		lenCounter[len(stringValue)]++
	}

	// Create a list of unique lengths sorted by ascending frequency, break ties with ascending length
	sortedLength := uniqueLengthSorted(lenCounter)

	totalCount := int64(len(strings))

	// Find the n_th most and least frequent length
	for index := 0; index < model.SampleSize && index < len(sortedLength); index++ {
		metric.MostFreqLen = append(metric.MostFreqLen, model.LenFreq{
			Length: sortedLength[index],
			Freq:   GetFrequency(lenCounter[sortedLength[index]], totalCount),
		})

		length := sortedLength[len(sortedLength)-index-1]

		metric.LeastFreqLen = append(metric.LeastFreqLen, model.LenFreq{
			Length: length,
			Freq:   GetFrequency(lenCounter[length], totalCount),
		})
	}

	metric.LeastFreqSample = getLeastFrequentSample(sortedLength, lenCounter, strings)

	return metric, nil
}

// Find n samples of least frequent length.
func getLeastFrequentSample(sortedLength []int, lenCounter map[int]int, strings []string) []string {
	// Strategy :
	// 	1. build a pool of length to take sample from :
	//		if lenCounter[leastFrequentLength] does not provide enough sample :
	// 			select n-i least frequent length till n samples are found.
	//  2. once length pool provide at least n samples,
	// 		iterate over values []interface{} and add to sample if length match length pool.
	// Note : 1. this strategy isn't random 2. is subject to over representation of n-i least frequent length.
	// ---------
	// Create a map of the lengths in lenSample
	// Check if the length of the string is in the lenMap.
	knownSample := 0
	lenSample := []int{}

	for i := len(sortedLength) - 1; i >= 0 && knownSample < model.SampleSize; i-- {
		knownSample += lenCounter[sortedLength[i]]
		lenSample = append(lenSample, sortedLength[i])
	}

	leastFreqSamples := make([]string, 0, model.SampleSize)

	lenMap := make(map[int]bool)
	for _, l := range lenSample {
		lenMap[l] = true
	}

	for _, string := range strings {
		if len(leastFreqSamples) == model.SampleSize {
			break
		}

		if lenMap[len(string)] {
			leastFreqSamples = append(leastFreqSamples, string)
		}
	}

	return leastFreqSamples
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
