package model

import "golang.org/x/exp/constraints"

type Col[T constraints.Ordered] struct {
	Name string    `json:"name"         jsonschema:"required" yaml:"name"`
	Type ValueType `json:"type"         jsonschema:"required" validate:"oneof=string numeric boolean" yaml:"type"` //nolint:lll

	// The 3 following parameter should be part of a Config struct
	Concept      string   `json:"concept,omitempty"      yaml:"concept,omitempty"`
	Constraint   []string `json:"constraint,omitempty"   yaml:"constraint,omitempty"`
	Confidential *bool    `json:"confidential,omitempty" yaml:"confidential,omitempty"`

	MainMetric Generic[T] `json:"mainMetric"   jsonschema:"required" yaml:"mainMetric"`

	StringMetric  String        `json:"stringMetric,omitempty"  yaml:"stringMetric,omitempty"`
	NumericMetric NumericMetric `json:"numericMetric,omitempty" yaml:"numericMetric,omitempty"`
	BoolMetric    BoolMetric    `json:"boolMetric,omitempty"    yaml:"boolMetric,omitempty"`
}

type Generic[T constraints.Ordered] struct {
	Count    uint  `json:"count"    jsonschema:"required" yaml:"count"`
	Empty    uint  `json:"empty"    jsonschema:"required" yaml:"empty"`
	Null     uint  `json:"null"     jsonschema:"required" yaml:"null"`
	Distinct *uint `json:"distinct" jsonschema:"required" yaml:"distinct"`
	Min      *T    `json:"min"      jsonschema:"required" yaml:"min"`
	Max      *T    `json:"max"      jsonschema:"required" yaml:"max"`
	Samples  []T   `json:"samples"  jsonschema:"required" yaml:"samples"`
}

type String struct {
	MinLen   int         `json:"minLen"    jsonschema:"required" yaml:"minLen"`
	MaxLen   int         `json:"maxLen"    jsonschema:"required" yaml:"maxLen"`
	CountLen int         `json:"countLen"  jsonschema:"required" yaml:"countLen"`
	Lengths  []StringLen `json:"lengths"   jsonschema:"required" yaml:"lengths"`
}

type StringLen struct {
	Length  int             `json:"length"  jsonschema:"required" yaml:"length"`
	Freq    float64         `json:"freq"    jsonschema:"required" yaml:"freq"`
	Metrics Generic[string] `json:"metrics" jsonschema:"required" yaml:"metrics"`
}
