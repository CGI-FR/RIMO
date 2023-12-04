package metricv2

import (
	"github.com/cgi-fr/rimo/pkg/modelv2"
)

type Analyser[T Accepted] interface {
	Read(*T)
	Build(*modelv2.Column)
}

type Multi[T Accepted] struct {
	analyser []Analyser[T]
}

func (m Multi[T]) Read(value *T) {
	for _, a := range m.analyser {
		a.Read(value)
	}
}

func (m Multi[T]) Build(metric *modelv2.Column) {
	for _, a := range m.analyser {
		a.Build(metric)
	}
}
