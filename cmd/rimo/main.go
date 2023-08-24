package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/io"
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
		Use:   "schema",
		Short: "Export rimo json schema",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Print current working directory
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal().Msgf("error getting current working directory: %v", err)
			}

			err = model.ExportSchema()
			if err != nil {
				log.Fatal().Msgf("error generating rimo schema: %v", err)
			}

			log.Info().Msgf("rimo schema successfully exported in %s", cwd)
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
