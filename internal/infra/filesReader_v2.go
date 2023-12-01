package infra

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/goccy/go-json"
)

var ErrReadFile = errors.New("error while reading file")

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

	source.Seek(0, 0)

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

	fr.source.Seek(0, 0)
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
