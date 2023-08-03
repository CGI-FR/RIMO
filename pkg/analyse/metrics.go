package analyse

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/model"
)

const (
	sampleSize = 5
)

// For a given dataCol return a model.Column with metrics.
func ComputeMetric(colName string, values []interface{}) model.Column {
	// Main metric
	name := colName
	colType := ColType(values)
	concept := ""
	var confidential *bool = nil //nolint

	// Generic metric
	count := int64(len(values))
	unique := int64(len(values))
	sample := Sample(values, sampleSize)

	genericMetric := model.GenericMetric{
		Count:  count,
		Unique: unique,
		Sample: sample,
	}

	stringMetric := model.StringMetric{}
	numericMetric := model.NumericMetric{}
	boolMetric := model.BoolMetric{}

	// Type specific metric
	switch colType {
	case "string":
		// Length frequency.
		lenCount, err := LenCounter(values)
		if err != nil {
			// Handle the error here, e.g. log it or return it to the caller
			log.Printf("Error counting lengths: %v", err)
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

		mostFreqLenFrequency := GetFrequency(mostFreqLenCount, count)
		leastFreqLenFrequency := GetFrequency(leastFreqLenCount, count)

		// Add metrics to stringMetric
		// TO DO: get 5 most frequent length
		stringMetric.MostFreqLen = map[int]float64{mostFreqLen: mostFreqLenFrequency}
		// TO DO: get 5 least frequent length
		stringMetric.LeastFreqLen = map[int]float64{leastFreqLen: leastFreqLenFrequency}
		// TO DO
		stringMetric.LeastFreqSample = []string{"undefined"}

	case "numeric":
		// Compute numeric metric
		break
	case "bool":
		// Compute bool metric
		break
	}

	// Create the column
	col := model.Column{
		Name:         name,
		Type:         colType,
		Concept:      concept,
		Constraint:   []string{},
		Confidential: *confidential,

		MainMetric: genericMetric,

		StringMetric:  stringMetric,
		NumericMetric: numericMetric,
		BoolMetric:    boolMetric,
	}

	return col
}

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

// StringSpecificMetric

var ErrNonString = fmt.Errorf("non-string value found")

// LenCounter return a map of length and their occurrence.
func LenCounter(values []interface{}) (map[int]int, error) {
	lengthCounter := make(map[int]int)

	for _, value := range values {
		if str, ok := value.(string); ok {
			length := len(str)
			lengthCounter[length]++
		} else {
			return nil, ErrNonString
		}
	}

	return lengthCounter, nil
}

func GetFrequency(occurrence int, count int64) float64 {
	return float64(occurrence) / float64(count)
}
