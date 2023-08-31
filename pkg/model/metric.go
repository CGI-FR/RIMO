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
