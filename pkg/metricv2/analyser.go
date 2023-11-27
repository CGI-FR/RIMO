package metricv2

import (
	"github.com/cgi-fr/rimo/pkg/modelv2"
	"golang.org/x/exp/constraints"
)

type Analyser[T constraints.Ordered] interface {
	Read(*T)
	Build(*modelv2.Generic[T])
}

type Multi[T constraints.Ordered] struct {
	analyser []Analyser[T]
}

func (m Multi[T]) Read(value *T) {
	for _, a := range m.analyser {
		a.Read(value)
	}
}

func (m Multi[T]) Build(metric *modelv2.Generic[T]) {
	for _, a := range m.analyser {
		a.Build(metric)
	}
}
