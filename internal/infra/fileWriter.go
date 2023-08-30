package infra

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/rimo"
	"gopkg.in/yaml.v3"
)

// Terminal writter interface

type TerminalWriter struct{}

func TerminalWriterFactory() *TerminalWriter {
	writer := TerminalWriter{}

	return &writer
}

func (w *TerminalWriter) Export(base *rimo.Base) error {
	fmt.Printf("%v\n", base)

	return nil
}

// YAML Writter interface

type YAMLWriter struct {
	outputPath string
}

func YAMLWriterFactory(filepath string) *YAMLWriter {
	writer := YAMLWriter{
		outputPath: filepath,
	}

	return &writer
}

// Write a YAML file from RIMO base at outputPath.
func (w *YAMLWriter) Export(base *rimo.Base) error {
	err := ValidateOutputPath(w.outputPath)
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
