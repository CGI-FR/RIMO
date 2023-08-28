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

	"github.com/cgi-fr/rimo/pkg/rimo"
)

var ErrValueType = errors.New("value type error")

// Return a rimo.Column.
func ComputeMetric(colName string, values []interface{}) (rimo.Column, error) {
	var confidential *bool = nil //nolint

	// Create the column.
	col := rimo.Column{
		Name:          colName,
		Type:          ColType(values),
		Concept:       "",
		Constraint:    []string{},
		Confidential:  confidential,
		MainMetric:    rimo.GenericMetric{}, //nolint:exhaustruct
		StringMetric:  rimo.StringMetric{},  //nolint:exhaustruct
		NumericMetric: rimo.NumericMetric{}, //nolint:exhaustruct
		BoolMetric:    rimo.BoolMetric{},    //nolint:exhaustruct
	}

	// Generic metric
	err := SetGenericMetric(values, &col.MainMetric)
	if err != nil {
		return rimo.Column{}, fmt.Errorf("error computing generic metric in column %v : %w", col.Name, err)
	}

	// Type specific metric
	switch col.Type {
	case rimo.ColType.String:
		err := SetStringMetric(values, &col.StringMetric)
		if err != nil {
			return rimo.Column{}, fmt.Errorf("error computing string metric in column %v : %w", col.Name, err)
		}

	case rimo.ColType.Numeric:
		err := SetNumericMetric(values, &col.NumericMetric)
		if err != nil {
			return rimo.Column{}, fmt.Errorf("error computing numeric metric in column %v : %w", col.Name, err)
		}

	case rimo.ColType.Bool:
		err := SetBoolMetric(values, &col.BoolMetric)
		if err != nil {
			return rimo.Column{}, fmt.Errorf("error computing bool metric in column %v : %w", col.Name, err)
		}
	}

	return col, nil
}

func ColType(values []interface{}) rimo.ValueType {
	colType := rimo.ColType.Undefined
	for i := 0; i < len(values) && colType == rimo.ColType.Undefined; i++ {
		colType = ValueType(values[i])
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

func ValueType(value interface{}) rimo.ValueType {
	switch value.(type) {
	case int:
		return rimo.ColType.Numeric
	case float64:
		return rimo.ColType.Numeric
	case json.Number:
		return rimo.ColType.Numeric
	case string:
		return rimo.ColType.String
	case bool:
		return rimo.ColType.Bool
	default:
		return rimo.ColType.Undefined
	}
}
