package analyse

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cgi-fr/rimo/pkg/model"
)

// Handle execution pipeline of analyse pkg.
func Analyse(inputList []string, outputPath string) {
	base := make(model.Base)

	// Iterate over inputList.
	for i := range inputList {
		inputPath := inputList[i]
		// Extract Base and Table name from inputFilePath.
		baseName := filepath.Base(filepath.Dir(inputPath))
		tableName := filepath.Base(inputPath)

		// Load inputFilePath.
		data := Load(inputPath, "new")

		// Analyse
		for colName, values := range data {
			// Append each column to Base structure.
			column := ComputeMetric(colName, values)
			base[baseName][tableName] = column
		}
	}

	// Export
	Export(base, outputPath)
}

func CheckFile(path string) {
	fileInfo, err := os.Stat(path)

	absPath, _ := filepath.Abs(path)
	// Check if the file exists
	if os.IsNotExist(err) {
		log.Fatalf("file does not exist: %s", absPath)
	}
	// Check if the file is a regular file
	if !fileInfo.Mode().IsRegular() {
		log.Fatalf("not a regular file: %s", absPath)
	}
}
