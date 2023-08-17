package main

import (
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
		Use:   "schema",
		Short: "Export rimo json schema",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Print current working directory
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal().Msgf("error getting current working directory: %v", err)
			}
			log.Info().Msgf("current working directory: %s", cwd)

			err = model.ExportSchema()
			if err != nil {
				log.Fatal().Msgf("error generating rimo schema: %v", err)
			}
		},
	}

	rimoAnalyseCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse [inputPath] [outputPath]",
		Short: "Generate a rimo.yaml from a directory of .jsonl files",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Run: func(cmd *cobra.Command, args []string) {
			inputPath := args[0]
			CheckDir(inputPath)

			outputPath := args[1]
			CheckDir(outputPath)

			// List of .jsonl files in input directory
			inputList, err := FilesList(inputPath, ".jsonl")
			if err != nil {
				log.Fatal().Msgf("error listing files: %v", err)
			}
			if len(inputList) == 0 {
				log.Fatal().Msgf("no .jsonl files found in %s", inputPath)
			}

			// Output path
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

func CheckFile(path string) {
	fileInfo, err := os.Stat(path)

	absPath, _ := filepath.Abs(path)
	// Check if the file exists
	if os.IsNotExist(err) {
		log.Fatal().Msgf("file does not exist: %s", absPath)
	}
	// Check if the file is a regular file
	if !fileInfo.Mode().IsRegular() {
		log.Fatal().Msgf("not a regular file: %s", absPath)
	}
}

func CheckDir(path string) {
	fileInfo, err := os.Stat(path)

	absPath, _ := filepath.Abs(path)
	// Check if the file exists
	if os.IsNotExist(err) {
		log.Fatal().Msgf("file does not exist: %s", absPath)
	}
	// Check if the file is a directory
	if !fileInfo.Mode().IsDir() {
		log.Fatal().Msgf("not a directory: %s", absPath)
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
