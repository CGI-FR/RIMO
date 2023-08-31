package rimo

import (
	"github.com/cgi-fr/rimo/pkg/model"
)

type Reader interface {
	BaseName() string
	Next() bool                                    // itère sur les colonnes.
	Value() ([]interface{}, string, string, error) // colValues, colName, tableName
}

type Writer interface {
	Export(base *model.Base) error
}
