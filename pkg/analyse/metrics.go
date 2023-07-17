package analyse

import (
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/models"
)

const (
	sampleSize = 5
)

func computeMetric(dataCol dataCol, colName string) models.Column {
	var confidential *bool = nil

	col := models.Column{
		Name:         colName,
		Type:         dataCol.colType,
		Concept:      "",
		Constraint:   []string{},
		Confidential: *confidential,
	}
	// Main metric
	col.MainMetric.Count = int64(len(dataCol.values))
	col.MainMetric.Unique = int64(len(dataCol.values))
	sample := make([]interface{}, sampleSize)
	for i := 0; i < 5; i++ {
		sample[i] = dataCol.values[rand.Intn(len(dataCol.values))]
	}
	col.MainMetric.Sample = sample

	// Type specific metric
	switch dataCol.colType {
	case "string":
		// Create a counter for each string length
		lengthCounter := make(map[int]int)
		for _, value := range dataCol.values {
			if str, ok := value.(string); ok {
				length := len(str)
				lengthCounter[length]++
			}
		}
	case "numeric":
		// Compute numeric metric
	case "bool":
		// Compute bool metric
	}

	return col
}
