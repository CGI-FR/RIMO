package metric

import (
	"github.com/cgi-fr/rimo/pkg/model"
)

type Distinct[T Accepted] struct {
	values map[T]int
}

func NewDistinct[T Accepted]() *Distinct[T] {
	return &Distinct[T]{
		values: make(map[T]int, 1024), //nolint:gomnd
	}
}

func (a *Distinct[T]) Read(value *T) {
	if value != nil {
		a.values[*value] = 0
	}
}

func (a *Distinct[T]) Build(metric *model.Column) {
	metric.MainMetric.Distinct = uint(len(a.values))
}
