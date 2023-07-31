package model

// RIMO YAML structure.
type (
	Base   map[string]Table
	Table  map[string]Column
	Column struct {
		Name         string   `yaml:"name"`
		Type         string   `yaml:"type"`
		Concept      string   `yaml:"concept"`
		Constraint   []string `yaml:"constraint"`
		Confidential bool     `yaml:"confidential"`

		MainMetric GenericMetric `yaml:"mainMetric"`

		StringMetric  StringMetric  `yaml:"stringMetric,omitempty"`
		NumericMetric NumericMetric `yaml:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `yaml:"boolMetric,omitempty"`
	}
)

// RIMO YAML metric.
type (
	GenericMetric struct {
		Count  int64         `yaml:"count"`
		Unique int64         `yaml:"unique"`
		Sample []interface{} `yaml:"sample"`
	}
	StringMetric struct {
		MostFreqLen     map[int]float64 `yaml:"mostFrequentLen"`
		LeastFreqLen    map[int]float64 `yaml:"leastFrequentLen"`
		LeastFreqSample []string        `yaml:"leastFrequentSample"`
	}
	NumericMetric struct {
		Min  float64 `yaml:"min"`
		Max  float64 `yaml:"max"`
		Mean float64 `yaml:"mean"`
	}
	BoolMetric struct {
		TrueRatio float64 `yaml:"trueRatio"`
	}
)
