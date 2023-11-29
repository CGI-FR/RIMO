package modelv2

type Generic struct {
	Count    uint  `json:"count"              yaml:"count"              jsonschema:"required"`
	Empty    uint  `json:"empty,omitempty"    yaml:"empty,omitempty"`
	Null     uint  `json:"nulls,omitempty"    yaml:"nulls,omitempty"`
	Distinct uint  `json:"distinct,omitempty" yaml:"distinct,omitempty"`
	Min      any   `json:"min,omitempty"      yaml:"min,omitempty"`
	Max      any   `json:"max,omitempty"      yaml:"max,omitempty"`
	Samples  []any `json:"samples"            yaml:"samples"            jsonschema:"required"`
}

type String struct {
	MinLen   int         `json:"minLen"   yaml:"minLen"`
	MaxLen   int         `json:"maxLen"   yaml:"maxLen"`
	CountLen int         `json:"countLen,omitempty" yaml:"countLen,omitempty"`
	Lengths  []StringLen `json:"lengths,omitempty"  yaml:"lengths,omitempty"`
}

type StringLen struct {
	Length  int     `json:"length"  yaml:"length"  jsonschema:"required"`
	Freq    float64 `json:"freq"    yaml:"freq"    jsonschema:"required"`
	Metrics Generic `json:"metrics" yaml:"metrics" jsonschema:"required"`
}

type Numeric struct{}

type Bool struct{}
