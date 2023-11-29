package modelv2

import "golang.org/x/exp/constraints"

type Column[T constraints.Ordered] struct {
	Name string `json:"name" yaml:"name" jsonschema:"required"`
	Type string `json:"type" yaml:"type" jsonschema:"required" validate:"oneof=string numeric boolean"`

	Config

	MainMetric    Generic[T] `json:"mainMetric"              yaml:"mainMetric"              jsonschema:"required"`
	StringMetric  String     `json:"stringMetric,omitempty"  yaml:"stringMetric,omitempty"`
	NumericMetric Numeric    `json:"numericMetric,omitempty" yaml:"numericMetric,omitempty"`
	BoolMetric    Bool       `json:"boolMetric,omitempty"    yaml:"boolMetric,omitempty"`
}
