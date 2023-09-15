package model

import (
	"fmt"
	"reflect"

	"github.com/hexops/valast"
)

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

// Should be improved with more detail about difference.
func SameBase(base1, base2 *Base) (bool, string) {
	if !reflect.DeepEqual(base1, base2) {
		msg := fmt.Sprintf("base is different : %s \n \n %s", valast.String(base1), valast.String(base2))

		return false, msg
	}

	return true, ""
}
