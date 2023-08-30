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

package infra_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/stretchr/testify/require"
)

func TestLoaderJSONL(t *testing.T) {
	t.Parallel()

	path := filepath.Join(testdataDir, "data1/data_input.jsonl")

	LoaderJSONL := infra.JSONLinesLoader{}

	data, err := LoaderJSONL.Load(path)
	require.NoError(t, err)
	fmt.Printf("dataMap: %v\n", data)
}
