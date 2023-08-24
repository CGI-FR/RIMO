package io

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/model"
	"gopkg.in/yaml.v3"
)

func Export(base model.Base, outputPath string) error {
	// Create output file.
	outputFile, err := os.Create(outputPath)
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
