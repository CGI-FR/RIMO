package model

// RIMO YAML structure.
type (
	Base struct {
		Name   string  `yaml:"database"`
		Tables []Table `yaml:"tables"`
	}
	Table struct {
		Name    string   `yaml:"name"`
		Columns []Column `yaml:"columns"`
	}
	Column struct {
		Name         string   `yaml:"name"`
		Type         string   `yaml:"type"`
		Concept      string   `yaml:"concept"`
		Constraint   []string `yaml:"constraint"`
		Confidential *bool    `yaml:"confidential"`

		MainMetric GenericMetric `yaml:"mainMetric"`

		StringMetric  StringMetric  `yaml:"stringMetric,omitempty"`
		NumericMetric NumericMetric `yaml:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `yaml:"boolMetric,omitempty"`
	}
)

// RIMO YAML metrics.
type (
	GenericMetric struct {
		Count  int64         `yaml:"count"`
		Unique int64         `yaml:"unique"`
		Sample []interface{} `yaml:"sample"`
	}
	StringMetric struct {
		MostFreqLen     []LenFreq `yaml:"mostFrequentLen"`
		LeastFreqLen    []LenFreq `yaml:"leastFrequentLen"`
		LeastFreqSample []string  `yaml:"leastFrequentSample"`
	}

	LenFreq struct {
		Length int
		Freq   float64
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
