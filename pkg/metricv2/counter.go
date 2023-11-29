package metricv2

import (
	"github.com/cgi-fr/rimo/pkg/modelv2"
	"golang.org/x/exp/constraints"
)

type Counter[T constraints.Ordered] struct {
	countTotal uint
	countNulls uint
	countEmpty uint
	zero       T
}

func NewCounter[T constraints.Ordered]() *Counter[T] {
	return &Counter[T]{
		countTotal: 0,
		countNulls: 0,
		countEmpty: 0,
		zero:       *new(T),
	}
}

func (c *Counter[T]) Read(value *T) {
	c.countTotal++

	switch {
	case value == nil:
		c.countNulls++
	case *value == c.zero:
		c.countEmpty++
	}
}

func (c *Counter[T]) Build(metric *modelv2.Column) {
	metric.MainMetric.Count = c.countTotal
	metric.MainMetric.Null = c.countNulls
	metric.MainMetric.Empty = c.countEmpty
}
