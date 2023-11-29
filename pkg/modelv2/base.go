package modelv2

type Base struct {
	Name   string           `json:"database" yaml:"database" jsonschema:"required"`
	Tables map[string]Table `json:"tables"   yaml:"tables"   jsonschema:"required"`
}

type Table struct {
	Columns []Column `json:"columns" yaml:"columns" jsonschema:"required" `
}

func NewBase(name string) *Base {
	return &Base{
		Name:   name,
		Tables: make(map[string]Table, 10),
	}
}
