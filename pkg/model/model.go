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

package model

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type RIMOType string

const (
	SampleSize              int = 5
	MostFrequentLenSize     int = 5
	MostFrequentSampleSize  int = 5
	LeastFrequentLenSize    int = 5
	LeastFrequentSampleSize int = 5
)

var ValueType = struct { //nolint:gochecknoglobals
	String    RIMOType
	Numeric   RIMOType
	Bool      RIMOType
	Undefined RIMOType
}{
	String:    "string",
	Numeric:   "numeric",
	Bool:      "bool",
	Undefined: "undefined",
}

// RIMO YAML structure.
type (
	Base struct {
		Name   string  `json:"database" jsonschema:"required" yaml:"database"`
		Tables []Table `json:"tables"   jsonschema:"required" yaml:"tables"`
	}

	Table struct {
		Name    string   `json:"name"    jsonschema:"required" yaml:"name"`
		Columns []Column `json:"columns" jsonschema:"required" yaml:"columns"`
	}

	Column struct {
		Name         string        `json:"name"         jsonschema:"required" yaml:"name"`
		Type         RIMOType      `json:"type"         jsonschema:"required" validate:"oneof=string numeric boolean" yaml:"type"` //nolint:lll
		Concept      string        `json:"concept"      jsonschema:"required" yaml:"concept"`
		Constraint   []string      `json:"constraint"   jsonschema:"required" yaml:"constraint"`
		Confidential *bool         `json:"confidential" jsonschema:"required" yaml:"confidential"`
		MainMetric   GenericMetric `json:"mainMetric"   jsonschema:"required" yaml:"mainMetric"`

		StringMetric  StringMetric  `json:"stringMetric,omitempty"  jsonschema:"required" yaml:"stringMetric,omitempty"`
		NumericMetric NumericMetric `json:"numericMetric,omitempty" jsonschema:"required" yaml:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `json:"boolMetric,omitempty"    jsonschema:"required" yaml:"boolMetric,omitempty"`
	}
)

// RIMO YAML metrics.
type (
	GenericMetric struct {
		Count  int           `json:"count"  jsonschema:"required" yaml:"count"`
		Empty  int           `json:"empty"  jsonschema:"required" yaml:"empty"`
		Unique int           `json:"unique" jsonschema:"required" yaml:"unique"`
		Sample []interface{} `json:"sample" jsonschema:"required" yaml:"sample"`
	}
	StringMetric struct {
		MostFreqLen  []LenFreq `json:"mostFrequentLen"  jsonschema:"required" yaml:"mostFrequentLen"`
		LeastFreqLen []LenFreq `json:"leastFrequentLen" jsonschema:"required" yaml:"leastFrequentLen"`
	}

	LenFreq struct {
		Length int      `json:"length" jsonschema:"required" yaml:"length"`
		Freq   float64  `json:"freq"   jsonschema:"required" yaml:"freq"`
		Sample []string `json:"sample" jsonschema:"required" yaml:"sample"`
	}

	NumericMetric struct {
		Min  float64 `json:"min"  jsonschema:"required" yaml:"min"`
		Max  float64 `json:"max"  jsonschema:"required" yaml:"max"`
		Mean float64 `json:"mean" jsonschema:"required" yaml:"mean"`
	}

	BoolMetric struct {
		TrueRatio float64 `json:"trueRatio" jsonschema:"required" yaml:"trueRatio"`
	}
)

func GetJSONSchema() (string, error) {
	resBytes, err := json.MarshalIndent(jsonschema.Reflect(&Base{}), "", "  ") //nolint:exhaustruct
	if err != nil {
		return "", fmt.Errorf("couldn't unmarshall Base in JSON : %w", err)
	}

	return string(resBytes), nil
}
