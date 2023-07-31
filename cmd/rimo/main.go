package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cgi-fr/rimo/pkg/analyse"
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

func main() {
	//nolint: exhaustruct
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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

	analyseCmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "analyse [input_path] [output_path]",
		Short: "Analyse a jsonl file and output a yaml file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World!")
			inputPath := args[0]
			outputPath := args[1]
			// Check if the input path is a directory
			if err := CheckDir(inputPath); err != nil {
				log.Fatal().Msgf("error checking input path: %v", err)
			}

			// Check if the output path is a regular file
			if err := CheckFile(outputPath); err != nil {
				log.Fatal().Msgf("error checking output path: %v", err)
			}

			// List of .jsonl files in input directory
			inputList, err := FilesList(inputPath, ".jsonl")
			if err != nil {
				log.Fatal().Msgf("error listing files: %v", err)
			}
			if len(inputList) == 0 {
				log.Fatal().Msgf("no .jsonl files found in %s", inputPath)
			}

			analyse.Analyse(inputList, outputPath)
		},
	}

	rootCmd.AddCommand(analyseCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("Error when executing command")
		os.Exit(1)
	}
}

func CheckFile(path string) error {
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
	return nil
}

func CheckDir(path string) error {
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
	return nil
}

func FilesList(path string, extension string) ([]string, error) {
	var inputList []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == extension {
			inputList = append(inputList, path)
		}
		return nil
	})
	return inputList, err
}
