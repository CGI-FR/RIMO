package analyse_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
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
	tempDir, err := os.MkdirTemp(TestDir, "export_test")
	assert.NoError(t, err)

	defer os.RemoveAll(tempDir)

	// Create a temporary file for the output
	outputFile := filepath.Join(tempDir, "output.yaml")

	// Export the base to the output file
	err = analyse.Export(base, outputFile)
	assert.NoError(t, err)

	// Read the output file and check its contents
	file, err := os.Open(outputFile)
	assert.NoError(t, err)

	defer file.Close()

	stat, err := file.Stat()
	assert.NoError(t, err)

	outputData := make([]byte, stat.Size())
	_, err = file.Read(outputData)
	assert.NoError(t, err)

	expectedData := `database: databaseName
tables:
    - name: tableName
      columns: []
`

	assert.Equal(t, expectedData, string(outputData))
}
