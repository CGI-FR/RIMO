package model

// RIMO YAML structure.
type (
	Base struct {
		Name string `json:"database" jsonschema:"required" yaml:"database"`
		// Tables should be map[string][]Column
		Tables []Table `json:"tables"   jsonschema:"required" yaml:"tables"`
	}

	Table struct {
		Name    string   `json:"name"    jsonschema:"required" yaml:"name"`
		Columns []Column `json:"columns" jsonschema:"required" yaml:"columns"`
	}
)
