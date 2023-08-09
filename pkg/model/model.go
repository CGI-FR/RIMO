package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/invopop/jsonschema"
)

// RIMO YAML structure.
type (
	Base struct {
		Name   string  `yaml:"database" json:"database" jsonschema:"required"`
		Tables []Table `yaml:"tables" json:"tables" jsonschema:"required"`
	}

	Table struct {
		Name    string   `yaml:"name" json:"name" jsonschema:"required"`
		Columns []Column `yaml:"columns" json:"columns" jsonschema:"required"`
	}

	Column struct {
		Name         string        `yaml:"name" json:"name" jsonschema:"required"`
		Type         string        `yaml:"type" json:"type" jsonschema:"required"`
		Concept      string        `yaml:"concept" json:"concept" jsonschema:"required"`
		Constraint   []string      `yaml:"constraint" json:"constraint" jsonschema:"required"`
		Confidential *bool         `yaml:"confidential" json:"confidential" jsonschema:"required"`
		MainMetric   GenericMetric `yaml:"mainMetric" json:"mainMetric" jsonschema:"required"`

		StringMetric  StringMetric  `yaml:"stringMetric,omitempty" json:"stringMetric,omitempty"`
		NumericMetric NumericMetric `yaml:"numericMetric,omitempty" json:"numericMetric,omitempty"`
		BoolMetric    BoolMetric    `yaml:"boolMetric,omitempty" json:"boolMetric,omitempty"`
	}
)

// RIMO YAML metrics.
type (
	GenericMetric struct {
		Count  int64         `yaml:"count" json:"count" jsonschema:"required"`
		Unique int64         `yaml:"unique" json:"unique" jsonschema:"required"`
		Sample []interface{} `yaml:"sample" json:"sample" jsonschema:"required"`
	}
	StringMetric struct {
		MostFreqLen     []LenFreq `yaml:"mostFrequentLen" json:"mostFrequentLen" jsonschema:"required"`
		LeastFreqLen    []LenFreq `yaml:"leastFrequentLen" json:"leastFrequentLen" jsonschema:"required"`
		LeastFreqSample []string  `yaml:"leastFrequentSample" json:"leastFrequentSample" jsonschema:"required"`
	}

	LenFreq struct {
		Length int     `json:"length" jsonschema:"required"`
		Freq   float64 `json:"freq" jsonschema:"required"`
	}

	NumericMetric struct {
		Min  float64 `yaml:"min" json:"min" jsonschema:"required"`
		Max  float64 `yaml:"max" json:"max" jsonschema:"required"`
		Mean float64 `yaml:"mean" json:"mean" jsonschema:"required"`
	}

	BoolMetric struct {
		TrueRatio float64 `yaml:"trueRatio" json:"trueRatio" jsonschema:"required"`
	}
)

func main() {
	err := ExportBaseSchema()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

// ExportBaseSchema exports the YAML schema for the Base struct to the current folder.
func ExportBaseSchema() error {
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
