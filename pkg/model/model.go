package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/invopop/jsonschema"
)

type RIMOType string

const (
	SampleSize              int = 5
	MostFrequentLenSize     int = 5
	MostFrequentSampleSize  int = 5
	LeastFrequentLenSize    int = 5
	LeastFrequentSampleSize int = 5
)

var ValueType = struct { //nolint:gochecknoglobals
	String    RIMOType
	Numeric   RIMOType
	Bool      RIMOType
	Undefined RIMOType
}{
	String:    "string",
	Numeric:   "numeric",
	Bool:      "bool",
	Undefined: "undefined",
}

// RIMO YAML structure.
type (
	Base struct {
		Name   string  `json:"database" jsonschema:"required" yaml:"database"`
		Tables []Table `json:"tables"   jsonschema:"required" yaml:"tables"`
	}

	Table struct {
		Name    string   `json:"name"    jsonschema:"required" yaml:"name"`
		Columns []Column `json:"columns" jsonschema:"required" yaml:"columns"`
	}

	Column struct {
		Name         string        `json:"name"         jsonschema:"required" yaml:"name"`
		Type         RIMOType      `json:"type"         jsonschema:"required" validate:"oneof=string numeric boolean" yaml:"type"` //nolint:lll
		Concept      string        `json:"concept"      jsonschema:"required" yaml:"concept"`
		Constraint   []string      `json:"constraint"   jsonschema:"required" yaml:"constraint"`
		Confidential *bool         `json:"confidential" jsonschema:"required" yaml:"confidential"`
		MainMetric   GenericMetric `json:"mainMetric"   jsonschema:"required" yaml:"mainMetric"`

		StringMetric  StringMetric  `json:"stringMetric,omitempty"  jsonschema:"required" yaml:"stringMetric,omitempty"`
		NumericMetric NumericMetric `json:"numericMetric,omitempty" jsonschema:"required" yaml:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `json:"boolMetric,omitempty"    jsonschema:"required" yaml:"boolMetric,omitempty"`
	}
)

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

// ExportBaseSchema exports the YAML schema for the Base struct to the current folder.
func ExportSchema() error {
	// Marshal the struct into a JSON string
	s := jsonschema.Reflect(&Base{}) //nolint:exhaustruct

	schema, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON schema: %w", err)
	}

	// Create a new file for the JSON schema
	file, err := os.Create("rimo-schema.json")
	if err != nil {
		return fmt.Errorf("error creating JSON schema file: %w", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(schema)
	if err != nil {
		return fmt.Errorf("error writing JSON data to file: %w", err)
	}

	return nil
}
