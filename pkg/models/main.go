package models

type Column struct {
	Name         string   `yaml:"name"`
	Type         string   `yaml:"type"`
	Concept      string   `yaml:"concept"`
	Constraint   []string `yaml:"constraint"`
	Confidential bool     `yaml:"confidential"`

	MainMetric GenericMetric `yaml:"main_metric"`

	StringMetric  StringMetric  `yaml:"string_metric,omitempty"`
	NumericMetric NumericMetric `yaml:"numeric_metric,omitempty"`
	BoolMetric    BoolMetric    `yaml:"bool_metric,omitempty"`
}

type GenericMetric struct {
	Count  int64         `yaml:"count"`
	Unique int64         `yaml:"unique"`
	Sample []interface{} `yaml:"sample"`
}

type StringMetric struct {
	MostFreqLen     map[int]float64 `yaml:"most_frequent_len"`
	LeastFreqLen    map[int]float64 `yaml:"least_frequent_len"`
	LeastFreqSample []string        `yaml:"least_frequent_sample"`
}

type NumericMetric struct {
	Min  float64 `yaml:"min"`
	Max  float64 `yaml:"max"`
	Mean float64 `yaml:"mean"`
}

type BoolMetric struct {
	TrueRatio float64 `yaml:"true_ratio"`
}
