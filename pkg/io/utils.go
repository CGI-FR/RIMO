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
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w : %s", ErrFileDoesNotExist, path)
		}

		return fmt.Errorf("%w : failed to get file info %s", err, path)
	}

	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%w : %s", ErrNotRegularFile, path)
	}

	if fileInfo.Mode().Perm()&ReadPerm != ReadPerm {
		return fmt.Errorf("%w : %s", ErrReadPermission, path)
	}

	return nil
}

func ValidateDirPath(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w : %s", ErrDirDoesNotExist, path)
		}

		return fmt.Errorf("failed to get directory info: %w", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%w : %s", ErrPathIsNotDir, path)
	}

	if fileInfo.Mode().Perm()&WriteDirPerm != WriteDirPerm {
		return fmt.Errorf("%w : %s", ErrWriteDirPermission, path)
	}

	return nil
}
