package analyse

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/model"
	"gopkg.in/yaml.v3"
)

// Export base to outputPath in YAML format.
const filePerm = 0o600

func Export(base model.Base, outputPath string) error {
	// Convert the base to YAML format
	yamlData, err := yaml.Marshal(base)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML data: %w", err)
	}

	// Write the YAML data to the output file
	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		return fmt.Errorf("failed to write YAML data to file: %w", err)
	}

	return nil
}
