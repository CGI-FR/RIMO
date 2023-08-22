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

package analyse_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExport(t *testing.T) {
	t.Parallel()
	t.Helper()

	base := model.Base{
		Name: "databaseName",
		Tables: []model.Table{
			{
				Name:    "tableName",
				Columns: []model.Column{},
			},
		},
	}

	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("./testdata/", "export_test")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	// Create a temporary file for the output
	outputFile := filepath.Join(tempDir, "output.yaml")

	// Export the base to the output file
	err = analyse.Export(base, outputFile)
	require.NoError(t, err)

	// Read the output file and check its contents
	file, err := os.Open(outputFile)
	require.NoError(t, err)

	defer file.Close()

	stat, err := file.Stat()
	require.NoError(t, err)

	outputData := make([]byte, stat.Size())
	_, err = file.Read(outputData)
	require.NoError(t, err)

	expectedData := `database: databaseName
tables:
    - name: tableName
      columns: []
`

	assert.Equal(t, expectedData, string(outputData))
}
