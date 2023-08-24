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

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Provisioned by ldflags.
var (
	name      string //nolint: gochecknoglobals
	version   string //nolint: gochecknoglobals
	commit    string //nolint: gochecknoglobals
	buildDate string //nolint: gochecknoglobals
	builtBy   string //nolint: gochecknoglobals
)

func main() { //nolint:funlen
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //nolint: exhaustruct

	log.Info().Msgf("%v %v (commit=%v date=%v by=%v)", name, version, commit, buildDate, builtBy)

	rootCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "rimo",
		Short: "Series of tool to help generate PIMO masking",
		Version: fmt.Sprintf(`%v (commit=%v date=%v by=%v)
		Copyright (C) 2021 CGI France
		License GPLv3: GNU GPL version 3 <https://gnu.org/licenses/gpl.html>.
		This is free software: you are free to change and redistribute it.
		There is NO WARRANTY, to the extent permitted by law.`, version, commit, buildDate, builtBy),
	}

	rimoSchemaCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "jsonschema",
		Short: "Return rimo jsonschema",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			jsonschema, err := model.GetJSONSchema()
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(jsonschema)
		},
	}

	rimoAnalyseCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse [inputPath] [outputPath]",
		Short: "Generate a rimo.yaml from a directory of .jsonl files",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Run: func(cmd *cobra.Command, args []string) {
			inputPath := args[0]
			if err := CheckDir(inputPath); err != nil {
				log.Fatal().Msgf("error checking input directory: %v", err)
			}

			outputPath := args[1]
			if err := CheckDir(outputPath); err != nil {
				log.Fatal().Msgf("error checking output directory: %v", err)
			}

			// List .jsonl files in input directory
			inputList, err := FilesList(inputPath, ".jsonl")
			if err != nil {
				log.Fatal().Msgf("error listing files: %v", err)
			}
			if len(inputList) == 0 {
				log.Fatal().Msgf("no .jsonl files found in %s", inputPath)
			}

			// Output path : outpathPath + basename + .yaml
			basename, err := analyse.GetBaseName(inputList[0])
			if err != nil {
				log.Fatal().Msgf("error getting basename: %v", err)
			}
			outputPath = filepath.Join(outputPath, basename+".yaml")

			err = analyse.Analyse(inputList, outputPath)
			if err != nil {
				log.Fatal().Msgf("error generating rimo.yaml: %v", err)
			}

			log.Info().Msgf("Successfully generated rimo.yaml at %s", outputPath)
		},
	}

	rootCmd.AddCommand(rimoAnalyseCmd)
	rootCmd.AddCommand(rimoSchemaCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

var (
	ErrNotExist = errors.New("path does not exist")
	ErrNotDir   = errors.New("path is not a directory")
	ErrNotFile  = errors.New("path is not a file")
)

func CheckFile(path string) error {
	// Get absPath
	path, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}

	fileInfo, err := os.Stat(path)
	// Check if the file exists
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrNotExist, path)
	}
	// Check if the file is a regular file
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%w: %s", ErrNotFile, path)
	}

	return nil
}

func CheckDir(path string) error {
	// Get absPath
	path, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}

	fileInfo, err := os.Stat(path)
	// Check if path exists
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrNotExist, path)
	}
	// Check if path is a directory
	if !fileInfo.Mode().IsDir() {
		return fmt.Errorf("%w: %s", ErrNotDir, path)
	}

	return nil
}

func FilesList(path string, extension string) ([]string, error) {
	pattern := filepath.Join(path, "*"+extension)

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}

	return files, nil
}
