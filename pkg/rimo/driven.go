package rimo

type Reader interface {
	BaseName() string
	Next() bool // itère sur une colonne
	Value() ColumnIterator
}

type ColumnIterator interface {
	Next() bool
	Value() (interface{}, error)
	ColumnName() string
	TableName() string
}

type Writer interface {
	Export(base *Base) error
}
