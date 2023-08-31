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
	"fmt"
	"os"
	"path/filepath"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/io"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/rimo"
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

	// Make use of interface instead of analyse/pkg
	rimoAnalyseInterfaceCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse2 [inputDir] [outputDir]",
		Short: "Generate a rimo.yaml from a directory of .jsonl files",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Run: func(cmd *cobra.Command, args []string) {
			inputDir := args[0]
			outputDir := args[1]

			inputList, err := BuildFilepathList(inputDir, ".jsonl")
			if err != nil {
				log.Fatal().Msgf("error listing files: %v", err)
			}

			reader, err := infra.FilesReaderFactory(inputList)
			if err != nil {
				log.Fatal().Msgf("error creating reader: %v", err)
			}

			writer := infra.YAMLWriterFactory(outputDir)

			err = rimo.AnalyseBase(reader, writer)
			if err != nil {
				log.Fatal().Msgf("error generating rimo.yaml: %v", err)
			}

			log.Info().Msgf("Successfully generated rimo.yaml in %s", outputDir)
		},
	}

	rimoAnalyseCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse [inputDir] [outputDir]",
		Short: "Generate a rimo.yaml from a directory of .jsonl files",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Run: func(cmd *cobra.Command, args []string) {
			inputDir := args[0]
			outputDir := args[1]

			// List .jsonl files in input directory
			if err := io.ValidateDirPath(inputDir); err != nil {
				log.Fatal().Msgf("error validating input directory: %v", err)
			}

			inputList, err := FilesList(inputDir, ".jsonl")
			if err != nil {
				log.Fatal().Msgf("error listing files: %v", err)
			}

			if len(inputList) == 0 {
				log.Fatal().Msgf("no .jsonl files found in %s", inputDir)
			}

			err = analyse.Orchestrator(inputList, outputDir)
			if err != nil {
				log.Fatal().Msgf("error generating rimo.yaml: %v", err)
			}

			log.Info().Msgf("Successfully generated rimo.yaml in %s", outputDir)
		},
	}

	rootCmd.AddCommand(rimoAnalyseCmd)
	rootCmd.AddCommand(rimoAnalyseInterfaceCmd)
	rootCmd.AddCommand(rimoSchemaCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

func FilesList(path string, extension string) ([]string, error) {
	pattern := filepath.Join(path, "*"+extension)

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}

	return files, nil
}

var ErrNoFile = fmt.Errorf("no file found")

func BuildFilepathList(path string, extension string) ([]string, error) {
	err := infra.ValidateDirPath(path)
	if err != nil {
		return nil, fmt.Errorf("failed to validate input directory: %w", err)
	}

	pattern := filepath.Join(path, "*"+extension)

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("%w : no %s files found in %s", ErrNoFile, extension, path)
	}

	return files, nil
}
