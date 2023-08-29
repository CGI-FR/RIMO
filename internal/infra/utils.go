package infra

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrFileDoesNotExist   = fmt.Errorf("file does not exist")
	ErrDirDoesNotExist    = fmt.Errorf("directory does not exist")
	ErrPathIsNotDir       = fmt.Errorf("path is not a directory")
	ErrNotRegularFile     = fmt.Errorf("path is not a regular file")
	ErrReadPermission     = fmt.Errorf("user does not have read permission for file")
	ErrWriteDirPermission = fmt.Errorf("user does not have write permission for directory")
)

const (
	ReadPerm     os.FileMode = 0o400
	WriteDirPerm os.FileMode = 0o200
)

func ValidateFilePath(path string) error {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileDoesNotExist, path)
	} else if err != nil {
		return fmt.Errorf("%w: failed to get file info %s", err, path)
	}

	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%w: %s", ErrNotRegularFile, path)
	}

	if fileInfo.Mode().Perm()&ReadPerm != ReadPerm {
		return fmt.Errorf("%w: %s", ErrReadPermission, path)
	}

	return nil
}

func ValidateDirPath(path string) error {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrDirDoesNotExist, path)
	} else if err != nil {
		return fmt.Errorf("failed to get directory info: %w", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%w: %s", ErrPathIsNotDir, path)
	}

	if fileInfo.Mode().Perm()&WriteDirPerm != WriteDirPerm {
		return fmt.Errorf("%w: %s", ErrWriteDirPermission, path)
	}

	return nil
}

var ErrNonExtractibleValue = errors.New("couldn't extract base or table name from path")

func ExtractName(path string) (string, string, error) {
	// path format : /path/to/jsonl/BASE_TABLE.jsonl
	fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(filepath.Base(path)))

	parts := strings.Split(fileName, "_")
	if len(parts) != 2 { //nolint:gomnd
		return "", "", fmt.Errorf("%w : %s", ErrNonExtractibleValue, path)
	}

	baseName := parts[0]
	if baseName == "" {
		return "", "", fmt.Errorf("%w : base name is empty from %s", ErrNonExtractibleValue, path)
	}

	tableName := parts[1]
	if tableName == "" {
		return "", "", fmt.Errorf("%w : table name is empty from %s", ErrNonExtractibleValue, path)
	}

	return baseName, tableName, nil
}
