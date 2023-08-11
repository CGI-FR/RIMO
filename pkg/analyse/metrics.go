package analyse

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/cgi-fr/rimo/pkg/model"
)

const (
	sampleSize = 5
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
		Sample: Sample(values, sampleSize),
	}

	col.MainMetric = genericMetric

	// Type specific metric

	switch colType {
	case "string":
		metric, err := StringMetric(values)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing string metric in column %v : %v", name, err)
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

func ColType(values []interface{}) string {
	colType := "unknown"
	for i := 0; i < len(values) && colType == "unknown"; i++ {
		colType = ValueType(values[i])
	}

	return colType
}

func Sample(values []interface{}, sampleSize int) []interface{} {
	sample := make([]interface{}, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = values[rand.Intn(len(values))]
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

	for i, v := range values { //nolint:varnamelen
		s, ok := v.(string)
		if !ok {
			return metric, fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, v, v)
		}

		strings[i] = s
		lenCounter[len(s)]++
	}

	// Sort the string lengths by descending count
	sorted := make([]int, 0, len(lenCounter))
	for l := range lenCounter {
		sorted = append(sorted, l)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return lenCounter[sorted[i]] > lenCounter[sorted[j]]
	})

	totalCount := int64(len(strings))

	// Find the 5th most and least frequent length
	for i := 0; i < sampleSize && i < len(sorted); i++ {
		metric.MostFreqLen = append(metric.MostFreqLen, model.LenFreq{
			Length: sorted[i],
			Freq:   GetFrequency(lenCounter[sorted[i]], totalCount),
		})

		length := sorted[len(sorted)-i-1]

		metric.LeastFreqLen = append(metric.LeastFreqLen, model.LenFreq{
			Length: length,
			Freq:   GetFrequency(lenCounter[length], totalCount),
		})
	}

	// Find 5 samples of least frequent length, if 5 samples are not found uses 2nd least frequent.
	// Loop through lenCounter till sampleSize samples are known.
	knownSample := 0
	lenSample := []int{}

	for i := len(sorted) - 1; i >= 0 && knownSample < sampleSize; i-- {
		knownSample += lenCounter[sorted[i]]
		lenSample = append(lenSample, sorted[i])
	}

	leastFreqSamples := make([]string, 0, sampleSize)

	// Create a map of the lengths in lenSample
	lenMap := make(map[int]bool)
	for _, l := range lenSample {
		lenMap[l] = true
	}

	for _, string := range strings {
		if len(leastFreqSamples) == sampleSize {
			break
		}

		// Check if the length of the string is in the lenMap.
		if lenMap[len(string)] {
			leastFreqSamples = append(leastFreqSamples, string)
		}
	}

	metric.LeastFreqSample = leastFreqSamples

	return metric, nil
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

// LenCounter return a map of length and their occurrence.
func LenCounter(values []interface{}) (map[int]int, error) {
	lengthCounter := make(map[int]int)

	for _, value := range values {
		if str, ok := value.(string); ok {
			length := len(str)
			lengthCounter[length]++
		} else {
			return nil, fmt.Errorf("%w : expected string found %T: %v", ErrValueType, value, value)
		}
	}

	return lengthCounter, nil
}

func GetFrequency(occurrence int, count int64) float64 {
	return float64(occurrence) / float64(count)
}

func ValueType(value interface{}) string {
	switch value.(type) {
	case int:
		return Numeric
	case float64:
		return Numeric
	case json.Number:
		return Numeric
	case string:
		return String
	case bool:
		return Boolean
	default:
		return "unknown"
	}
}
