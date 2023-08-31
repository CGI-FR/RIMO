package infra

import (
	"errors"
	"fmt"
)

// Errors declaration.
var (
	ErrInvalidFilePath = errors.New("failed to validate path")
	ErrNoFilePath      = errors.New("no file path provided")
	ErrNonUniqueBase   = errors.New("base name is not unique")
)

// FilesReader can read multiple type of file and feed data to rimo.
// FilesReader is responsible of :
// - BaseName() return the name of the base
// - Next() return true if there is a next value to read
// - Value() return the value of the current column, the name of the column and the name of the table
// Interface itself with a Loader interface. Which currently only supports YAML files.
// Loader and FilesReader can be initialized with LoaderFactory and FilesReaderFactory.
type FilesReader struct {
	filepathList []string
	loader       JSONLinesLoader // responsible of loading a file format
	baseName     string
	// variable for looping over columns
	fileIndex       int
	colNameMapIndex map[int]string // map of column name by index
	colIndex        int            // value of current column index
	// given by Value()
	dataMap   map[string][]interface{}
	tableName string // filled by FilesReader
}

// Constructor for FilesReader.
func FilesReaderFactory(filepathList []string) (*FilesReader, error) {
	var err error

	// Process inputDirList
	if len(filepathList) == 0 {
		return nil, ErrNoFilePath
	}

	for _, path := range filepathList {
		err := ValidateFilePath(path)
		if err != nil {
			return nil, ErrInvalidFilePath
		}
	}

	// Initialize FilesReader
	var filesReader FilesReader
	filesReader.filepathList = filepathList
	filesReader.fileIndex = -1

	filesReader.baseName, err = filesReader.isBaseUnique()
	if err != nil {
		return nil, fmt.Errorf("base is not unique: %w", err)
	}

	// Use of JSONLinesLoader
	filesReader.loader = JSONLinesLoader{}

	return &filesReader, nil
}

// RIMO.Reader interface implementation

func (r *FilesReader) BaseName() string {
	return r.baseName
}

func (r *FilesReader) Next() bool {
	// First call to Next()
	if r.fileIndex == -1 {
		r.fileIndex = 0
		r.colIndex = 0

		return true
	}

	// Current file contain column left to process.
	if r.colIndex < len(r.dataMap) {
		r.colIndex++
	}

	// Current file contain no columns left to process.
	if r.colIndex == len(r.dataMap) {
		// Current file is last file.
		if r.fileIndex == len(r.filepathList)-1 {
			return false
		}
		// There is a next file.
		r.fileIndex++
		r.colIndex = 0
	}

	return true
}

// Charger les fichiers un à un dans une dataMap.
// Retourne les valeurs d'une colonne, son nom et le nom de table.
func (r *FilesReader) Value() ([]interface{}, string, string, error) {
	var err error

	// colIndex = 0 : new file to load
	if r.colIndex == 0 {
		filepath := r.filepathList[r.fileIndex]

		// Extract table name from file name
		_, r.tableName, err = ExtractName(filepath)
		if err != nil {
			return nil, "", "", fmt.Errorf("failed to extract table name: %w", err)
		}

		// Load file in dataMap
		r.dataMap, err = r.loader.Load(r.filepathList[r.fileIndex])
		if err != nil {
			panic(err)
		}

		// Create a map of column name by index
		r.colNameMapIndex = make(map[int]string, 0)
		i := 0

		for k := range r.dataMap {
			r.colNameMapIndex[i] = k
			i++
		}
	}

	// colIndex = n : current file have been partially processed

	// return values, colName, tableName
	return r.dataMap[r.colNameMapIndex[r.colIndex]], r.colNameMapIndex[r.colIndex], r.tableName, nil
}

func (r *FilesReader) isBaseUnique() (string, error) {
	baseName, _, err := ExtractName(r.filepathList[0])
	if err != nil {
		return "", err
	}

	for _, path := range r.filepathList {
		baseNameI, _, err := ExtractName(path)
		if err != nil {
			return "", err
		}

		if baseName != baseNameI {
			return "", fmt.Errorf("%w : %s and %s", ErrNonUniqueBase, baseName, baseNameI)
		}
	}

	return baseName, nil
}
