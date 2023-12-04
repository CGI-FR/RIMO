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
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const DefaultSampleSize = uint(5)

// Provisioned by ldflags.
var (
	name      string //nolint: gochecknoglobals
	version   string //nolint: gochecknoglobals
	commit    string //nolint: gochecknoglobals
	buildDate string //nolint: gochecknoglobals
	builtBy   string //nolint: gochecknoglobals

	sampleSize uint //nolint: gochecknoglobals
	distinct   bool //nolint: gochecknoglobals
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
	rimoAnalyseCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse [inputDir] [outputDir]",
		Short: "Generate a rimo.yaml from a directory of .jsonl files",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Run: func(cmd *cobra.Command, args []string) {
			inputDir := args[0]
			outputDir := args[1]

			// Reader
			reader, err := infra.NewJSONLFolderReader(inputDir)
			if err != nil {
				log.Fatal().Msgf("error creating reader: %v", err)
			}

			outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.yaml", reader.BaseName()))

			writer, err := infra.YAMLWriterFactory(outputPath)
			if err != nil {
				log.Fatal().Msgf("error creating writer: %v", err)
			}

			driver := rimo.Driver{SampleSize: sampleSize, Distinct: distinct}

			err = driver.AnalyseBase(reader, writer)
			if err != nil {
				log.Fatal().Msgf("error generating rimo.yaml: %v", err)
			}

			log.Info().Msgf("Successfully generated rimo.yaml in %s", outputDir)
		},
	}

	rimoAnalyseCmd.Flags().UintVar(&sampleSize, "sample-size", DefaultSampleSize, "number of sample value to collect")
	rimoAnalyseCmd.Flags().BoolVarP(&distinct, "distinct", "d", false, "count distinct values")

	rootCmd.AddCommand(rimoAnalyseCmd)
	rootCmd.AddCommand(rimoSchemaCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}
