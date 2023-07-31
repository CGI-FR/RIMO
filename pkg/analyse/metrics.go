package analyse

import (
	"encoding/json"
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/model"
)

const (
	sampleSize = 5
)

// For a given dataCol return a model.Column with metrics.
func ComputeMetric(dataCol DataCol) model.Column {
	// Main metric
	name := dataCol.ColName
	colType := ColType(dataCol.Values)
	concept := ""
	var confidential *bool = nil //nolint

	// Generic metric
	count := int64(len(dataCol.Values))
	unique := int64(len(dataCol.Values))
	sample := Sample(dataCol.Values, sampleSize)

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
		lenCount := LenCounter(dataCol.Values)

		// MostFreq and LeastFreq
		mostFreqLen := 0
		mostFreqLenCount := 0
		// LeastFreqLen
		leastFreqLen := len(dataCol.Values) + 1
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
		stringMetric.MostFreqLen = map[int]float64{mostFreqLen: mostFreqLenFrequency}    // TODO: get 5 most frequent length
		stringMetric.LeastFreqLen = map[int]float64{leastFreqLen: leastFreqLenFrequency} // TODO: get 5 least frequent length
		stringMetric.LeastFreqSample = []string{"undefined"}                             // TODO

	case "numeric":
		// Compute numeric metric
		break
	case "bool":
		// Compute bool metric
		break
	}

	// Create the column
	col := model.Column{ //nolint:exhaustruct
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

func LenCounter(values []interface{}) map[int]int {
	lengthCounter := make(map[int]int)

	for _, value := range values {
		if str, ok := value.(string); ok {
			length := len(str)
			lengthCounter[length]++
		}
	}

	return lengthCounter
}

func GetFrequency(occurence int, count int64) float64 {
	return float64(occurence) / float64(count)
}
