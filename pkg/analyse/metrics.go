package analyse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"

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

func Sample(values []interface{}, sampleSize int) []interface{} {
	sample := make([]interface{}, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = values[rand.Intn(len(values))]
	}

	return sample
}

// Specific type metric.

var ErrValueType = errors.New("value type error")

// String metric : MostFreqLen, LeastFreqLen, LeastFreqSample

func StringMetric(values []interface{}) (model.StringMetric, error) {
	// Length frequency.
	lenCount, err := LenCounter(values)
	if err != nil {
		// Handle the error here, e.g. log it or return it to the caller
		log.Printf("Error counting lengths of strings : %v", err)
	}

	// MostFreq and LeastFreq
	mostFreqLen := 0
	mostFreqLenCount := 0
	// LeastFreqLen
	leastFreqLen := len(values) + 1
	leastFreqLenCount := 0

	for len, count := range lenCount {
		if count > mostFreqLenCount {
			mostFreqLen = len
			mostFreqLenCount = count
		}

		if count < leastFreqLenCount {
			leastFreqLen = len
			leastFreqLenCount = count
		}
	}

	mostFreqLenFrequency := GetFrequency(mostFreqLenCount, int64(len(values)))
	leastFreqLenFrequency := GetFrequency(leastFreqLenCount, int64(len(values)))

	// Add metrics to stringMetric
	stringMetric := model.StringMetric{
		MostFreqLen:     []model.LenFreq{{Length: mostFreqLen, Freq: mostFreqLenFrequency}},
		LeastFreqLen:    []model.LenFreq{{Length: leastFreqLen, Freq: leastFreqLenFrequency}},
		LeastFreqSample: []string{"undefined"},
	}

	return stringMetric, nil
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
