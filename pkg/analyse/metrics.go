package analyse

import (
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/models"
)

const (
	sampleSize = 5
)

func ComputeMetric(dataCol DataCol, colName string) models.Column {
	var confidential *bool = nil

	col := models.Column{
		Name:         colName,
		Type:         dataCol.ColType,
		Concept:      "",
		Constraint:   []string{},
		Confidential: *confidential,
	}
	// Main metric
	col.MainMetric.Count = int64(len(dataCol.Values))
	col.MainMetric.Unique = int64(len(dataCol.Values))
	sample := make([]interface{}, sampleSize)

	for i := 0; i < 5; i++ {
		sample[i] = dataCol.Values[rand.Intn(len(dataCol.Values))]
	}

	col.MainMetric.Sample = sample

	// Type specific metric
	switch dataCol.ColType {
	case "string":
		// Create a counter for each string length
		lengthCounter := make(map[int]int)

		for _, value := range dataCol.Values {
			if str, ok := value.(string); ok {
				length := len(str)
				lengthCounter[length]++
			}
		}
	case "numeric":
		// Compute numeric metric
		break
	case "bool":
		// Compute bool metric
		break
	}

	return col
}
