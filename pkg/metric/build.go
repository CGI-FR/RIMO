package metric

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cgi-fr/rimo/pkg/model"
)

var ErrValueType = errors.New("value type error")

// Return a model.Column.
func ComputeMetric(colName string, values []interface{}) (model.Column, error) {
	var confidential *bool = nil //nolint

	// Create the column.
	col := model.Column{
		Name:          colName,
		Type:          ColType(values),
		Concept:       "",
		Constraint:    []string{},
		Confidential:  confidential,
		MainMetric:    model.GenericMetric{}, //nolint:exhaustruct
		StringMetric:  model.StringMetric{},  //nolint:exhaustruct
		NumericMetric: model.NumericMetric{}, //nolint:exhaustruct
		BoolMetric:    model.BoolMetric{},    //nolint:exhaustruct
	}

	// Generic metric
	err := SetGenericMetric(values, &col.MainMetric)
	if err != nil {
		return model.Column{}, fmt.Errorf("error computing generic metric in column %v : %w", col.Name, err)
	}

	// Type specific metric
	switch col.Type {
	case model.ValueType.String:
		err := SetStringMetric(values, &col.StringMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing string metric in column %v : %w", col.Name, err)
		}

	case model.ValueType.Numeric:
		err := SetNumericMetric(values, &col.NumericMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing numeric metric in column %v : %w", col.Name, err)
		}

	case model.ValueType.Bool:
		err := SetBoolMetric(values, &col.BoolMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing bool metric in column %v : %w", col.Name, err)
		}
	}

	return col, nil
}

func ColType(values []interface{}) model.RIMOType {
	colType := model.ValueType.Undefined
	for i := 0; i < len(values) && colType == model.ValueType.Undefined; i++ {
		colType = ValueType(values[i])
	}

	return colType
}

// Utils functions.

func GetFrequency(occurrence int, count int) float64 {
	return float64(occurrence) / float64(count)
}

func GetFirstValue(values []interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}

	return nil
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
