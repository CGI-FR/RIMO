package main

import (
	"fmt"
	"os"
	"path/filepath"

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
			CheckFile(inputPath)
			CheckFile(outputPath)

			var inputList []string
			var outputList []string

			rimo.Analyse(inputList, outputList)
		},
	}

	rootCmd.AddCommand(analyseCmd)

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
		log.Fatalf("file does not exist: %s", absPath)
	}
	// Check if the file is a regular file
	if !fileInfo.Mode().IsRegular() {
		log.Fatalf("not a regular file: %s", absPath)
	}
}
