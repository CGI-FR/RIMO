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
