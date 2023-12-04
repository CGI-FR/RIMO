package infra

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/goccy/go-json"
)

var ErrReadFile = errors.New("error while reading file")

type JSONLFolderReader struct {
	basename string
	readers  []*JSONLFileReader
	current  int
}

func NewJSONLFolderReader(folderpath string) (*JSONLFolderReader, error) {
	basename := path.Base(folderpath)

	pattern := filepath.Join(folderpath, "*.jsonl")

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}

	readers := make([]*JSONLFileReader, len(files))
	for index, filepath := range files {
		readers[index], err = NewJSONLFileReader(basename, filepath)
		if err != nil {
			return nil, fmt.Errorf("error opening files: %w", err)
		}
	}

	return &JSONLFolderReader{
		basename: basename,
		readers:  readers,
		current:  0,
	}, nil
}

func (r *JSONLFolderReader) BaseName() string {
	return r.basename
}

func (r *JSONLFolderReader) Next() bool {
	if r.current < len(r.readers) && !r.readers[r.current].Next() {
		r.current++

		return r.Next()
	}

	return r.current < len(r.readers)
}

func (r *JSONLFolderReader) Col() (rimo.ColReader, error) { //nolint:ireturn
	return r.readers[r.current].Col()
}

type JSONLFileReader struct {
	tablename string
	source    *os.File
	columns   []string
	current   int
	decoder   *json.Decoder
	basename  string
}

func NewJSONLFileReader(basename string, filepath string) (*JSONLFileReader, error) {
	source, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	template := map[string]any{}

	decoder := json.NewDecoder(source)
	if err := decoder.Decode(&template); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	if _, err := source.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	columns := make([]string, 0, len(template))
	for column := range template {
		columns = append(columns, column)
	}

	return &JSONLFileReader{
		tablename: strings.TrimSuffix(path.Base(filepath), path.Ext(filepath)),
		source:    source,
		columns:   columns,
		current:   -1,
		decoder:   json.NewDecoder(source),
		basename:  basename,
	}, nil
}

func (fr *JSONLFileReader) BaseName() string {
	return fr.basename
}

func (fr *JSONLFileReader) Next() bool {
	fr.current++

	if _, err := fr.source.Seek(0, 0); err != nil {
		panic(err)
	}

	fr.decoder = json.NewDecoder(fr.source)

	return fr.current < len(fr.columns)
}

func (fr *JSONLFileReader) Col() (rimo.ColReader, error) { //nolint:ireturn
	return NewJSONLColReader(fr.tablename, fr.columns[fr.current], fr.decoder), nil
}

type JSONLColReader struct {
	table   string
	column  string
	decoder *json.Decoder
}

func NewJSONLColReader(table, column string, decoder *json.Decoder) *JSONLColReader {
	return &JSONLColReader{
		table:   table,
		column:  column,
		decoder: decoder,
	}
}

func (cr *JSONLColReader) ColName() string {
	return cr.column
}

func (cr *JSONLColReader) TableName() string {
	return cr.table
}

func (cr *JSONLColReader) Next() bool {
	return cr.decoder.More()
}

func (cr *JSONLColReader) Value() (any, error) {
	row := map[string]any{}

	if err := cr.decoder.Decode(&row); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	return row[cr.column], nil
}
