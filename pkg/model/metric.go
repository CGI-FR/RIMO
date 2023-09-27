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

// Type that a column can be.
type ValueType string

var ColType = struct { //nolint:gochecknoglobals
	String    ValueType
	Numeric   ValueType
	Bool      ValueType
	Undefined ValueType
}{
	String:    "string",
	Numeric:   "numeric",
	Bool:      "bool",
	Undefined: "undefined",
}
