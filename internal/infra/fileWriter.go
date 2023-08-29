package infra

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/rimo"
	"gopkg.in/yaml.v3"
)

type FileWriter struct {
	rimo.Writer
	outputPath string
}

// Write a YAML file from a RIMO base at outputPath.
func (w *FileWriter) Export(base rimo.Base) error {
	err := ValidateFilePath(w.outputPath)
	if err != nil {
		return fmt.Errorf("failed to validate file path: %w", err)
	}

	outputFile, err := os.Create(w.outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Encode Base to YAML.
	encoder := yaml.NewEncoder(outputFile)
	defer encoder.Close()

	err = encoder.Encode(base)
	if err != nil {
		return fmt.Errorf("failed to encode Base to YAML: %w", err)
	}

	return nil
}
