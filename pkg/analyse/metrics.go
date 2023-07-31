package analyse

import (
	"encoding/json"
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/model"
)

const (
	sampleSize = 5
)

func ComputeMetric(dataCol DataCol) model.Column {
	var confidential *bool = nil //nolint

	col := model.Column{ //nolint:exhaustruct
		Name:         dataCol.ColName,
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
		sample[i] = dataCol.Values[rand.Intn(len(dataCol.Values))] //nolint:gosec
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

// Build a map of column names to column types.
func BuildColType(data DataMap) DataMap {
	for colName, colData := range data {
		// Iterate till colType is not unknown
		for i := 0; i < len(colData.Values) && data[colName].ColType == "unknown"; i++ {
			data[colName] = DataCol{
				ColType: TypeOf(colData.Values[i]),
				Values:  colData.Values,
			}
		}
	}

	return data
}

func TypeOf(v interface{}) string {
	switch v.(type) {
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
