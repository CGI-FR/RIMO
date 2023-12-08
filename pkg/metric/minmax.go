package metric

import (
	"github.com/cgi-fr/rimo/pkg/model"
	"golang.org/x/exp/constraints"
)

type MinMax[T constraints.Ordered] struct {
	min *T
	max *T
}

func NewMinMax[T constraints.Ordered]() *MinMax[T] {
	return &MinMax[T]{
		min: nil,
		max: nil,
	}
}

func (a *MinMax[T]) Read(value *T) {
	if value != nil {
		if a.min == nil {
			a.min = value
		}

		if a.max == nil {
			a.max = value
		}

		if *value < *a.min {
			a.min = value
		} else if *value > *a.max {
			a.max = value
		}
	}
}

func (a *MinMax[T]) Build(metric *model.Column) {
	metric.MainMetric.Min = a.min
	metric.MainMetric.Max = a.max
}
