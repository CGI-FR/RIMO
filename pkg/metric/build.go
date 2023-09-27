// Copyright (C) 2023 CGI France
//
// This file is part of RIMO.
//
// RIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// RIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with RIMO.  If not, see <http://www.gnu.org/licenses/>.

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
		Type:          GetColType(values),
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
	case model.ColType.String:
		err := SetStringMetric(values, &col.StringMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing string metric in column %v : %w", col.Name, err)
		}

	case model.ColType.Numeric:
		err := SetNumericMetric(values, &col.NumericMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing numeric metric in column %v : %w", col.Name, err)
		}

	case model.ColType.Bool:
		err := SetBoolMetric(values, &col.BoolMetric)
		if err != nil {
			return model.Column{}, fmt.Errorf("error computing bool metric in column %v : %w", col.Name, err)
		}
	}

	return col, nil
}

func GetColType(values []interface{}) model.ValueType {
	colType := model.ColType.Undefined
	for i := 0; i < len(values) && colType == model.ColType.Undefined; i++ {
		colType = ColType(values[i])
	}

	return colType
}

// Utils functions.

func GetFrequency(occurrence int, count int) float64 {
	return float64(occurrence) / float64(count)
}

// To check why not using isNil() ?
func GetFirstValue(values []interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}

	return nil
}

func ColType(value interface{}) model.ValueType {
	switch value.(type) {
	case int:
		return model.ColType.Numeric
	case float64:
		return model.ColType.Numeric
	case json.Number:
		return model.ColType.Numeric
	case string:
		return model.ColType.String
	case bool:
		return model.ColType.Bool
	default:
		return model.ColType.Undefined
	}
}
