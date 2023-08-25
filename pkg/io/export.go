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
