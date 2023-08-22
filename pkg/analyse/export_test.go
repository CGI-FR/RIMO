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
