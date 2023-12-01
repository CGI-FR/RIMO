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

package infra

import (
	"fmt"
	"os"

	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/modelv2"
	"gopkg.in/yaml.v3"
)

// Terminal writter interface

type StdoutWriter struct{}

func StdoutWriterFactory() *StdoutWriter {
	writer := StdoutWriter{}

	return &writer
}

func (w *StdoutWriter) Export(base *modelv2.Base) error {
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
