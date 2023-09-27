// Copyright (C) 2023 CGI France
//
// This file is part of RIMO.
//
// RIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// RIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with RIMO.  If not, see <http://www.gnu.org/licenses/>.

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
