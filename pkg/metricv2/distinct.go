package metricv2

import (
	"github.com/cgi-fr/rimo/pkg/modelv2"
	"golang.org/x/exp/constraints"
)

type Distinct[T constraints.Ordered] struct {
	values map[T]int
}

func NewDistinct[T constraints.Ordered]() *Distinct[T] {
	return &Distinct[T]{
		values: make(map[T]int, 1024), //nolint:gomnd
	}
}

func (a *Distinct[T]) Read(value *T) {
	if value != nil {
		a.values[*value] = 0
	}
}

func (a *Distinct[T]) Build(metric *modelv2.Column) {
	metric.MainMetric.Distinct = uint(len(a.values))
}
