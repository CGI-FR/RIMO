package modelv2

import "golang.org/x/exp/constraints"

type Generic[T constraints.Ordered] struct {
	Count    uint    `json:"count"              yaml:"count"              jsonschema:"required"`
	Empty    uint    `json:"empty,omitempty"    yaml:"empty,omitempty"`
	Null     uint    `json:"nulls,omitempty"    yaml:"nulls,omitempty"`
	Distinct uint    `json:"distinct,omitempty" yaml:"distinct,omitempty"`
	Min      *T      `json:"min,omitempty"      yaml:"min,omitempty"`
	Max      *T      `json:"max,omitempty"      yaml:"max,omitempty"`
	Samples  []T     `json:"samples"            yaml:"samples"            jsonschema:"required"`
	String   *String `json:"string,omitempty"   yaml:"string,omitempty"`
}

type String struct {
	MinLen   int         `json:"minLen"   yaml:"minLen"`
	MaxLen   int         `json:"maxLen"   yaml:"maxLen"`
	CountLen int         `json:"countLen,omitempty" yaml:"countLen,omitempty"`
	Lengths  []StringLen `json:"lengths,omitempty"  yaml:"lengths,omitempty"`
}

type StringLen struct {
	Length  int             `json:"length"  yaml:"length"  jsonschema:"required"`
	Freq    float64         `json:"freq"    yaml:"freq"    jsonschema:"required"`
	Metrics Generic[string] `json:"metrics" yaml:"metrics" jsonschema:"required"`
}
