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

const (
	SampleSize              int = 5
	MostFrequentLenSize     int = 5
	MostFrequentSampleSize  int = 5
	LeastFrequentLenSize    int = 5
	LeastFrequentSampleSize int = 5
)

type (
	Column struct {
		Name string    `json:"name"         jsonschema:"required" yaml:"name"`
		Type ValueType `json:"type"         jsonschema:"required" validate:"oneof=string numeric boolean" yaml:"type"` //nolint:lll

		// The 3 following parameter should be part of a Config struct
		Concept      string   `json:"concept"      jsonschema:"required" yaml:"concept"`
		Constraint   []string `json:"constraint"   jsonschema:"required" yaml:"constraint"`
		Confidential *bool    `json:"confidential" jsonschema:"required" yaml:"confidential"`

		MainMetric GenericMetric `json:"mainMetric"   jsonschema:"required" yaml:"mainMetric"`

		StringMetric  StringMetric  `json:"stringMetric,omitempty"  jsonschema:"required" yaml:"stringMetric,omitempty"`
		NumericMetric NumericMetric `json:"numericMetric,omitempty" jsonschema:"required" yaml:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `json:"boolMetric,omitempty"    jsonschema:"required" yaml:"boolMetric,omitempty"`
	}
)
