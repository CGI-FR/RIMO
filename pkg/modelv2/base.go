package modelv2

const DefaultTableSize = 10

type Base struct {
	Name   string  `json:"database" yaml:"database" jsonschema:"required"`
	Tables []Table `json:"tables"   yaml:"tables"   jsonschema:"required"`
}

type Table struct {
	Name    string   `json:"name"    yaml:"name"    jsonschema:"required"`
	Columns []Column `json:"columns" yaml:"columns" jsonschema:"required"`
}

func NewBase(name string) *Base {
	return &Base{
		Name:   name,
		Tables: make([]Table, 0, DefaultTableSize),
	}
}
