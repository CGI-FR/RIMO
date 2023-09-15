package infra

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/model"
	"gopkg.in/yaml.v3"
)

// Terminal writter interface

type StdoutWriter struct{}

func StdoutWriterFactory() *StdoutWriter {
	writer := StdoutWriter{}

	return &writer
}

func (w *StdoutWriter) Export(base *model.Base) error {
	fmt.Printf("%v\n", base)

	return nil
}

// YAML Writter interface

type YAMLWriter struct {
	outputPath string
}

func YAMLWriterFactory(filepath string) (*YAMLWriter, error) {
	err := ValidateOutputPath(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to validate file path: %w", err)
	}

	writer := YAMLWriter{
		outputPath: filepath,
	}

	return &writer, nil
}

// Write a YAML file from RIMO base at outputPath.
func (w *YAMLWriter) Export(base *model.Base) error {
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
