package rimo_test

import (
	"github.com/cgi-fr/rimo/pkg/rimo"
)

type TestReader struct {
	index  int
	values []interface{}
}

func (r *TestReader) BaseName() string {
	return ""
}

func (r *TestReader) Next() bool {
	return r.index < len(r.values)
}

func (r *TestReader) Value() ([]interface{}, string, string, error) {
	return nil, "", "", nil
}

type TestWriter struct{}

func (w *TestWriter) Export(base *rimo.Base) error {
	return nil
}
