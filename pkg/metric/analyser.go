package metric

import (
	"github.com/cgi-fr/rimo/pkg/model"
)

type Analyser[T Accepted] interface {
	Read(*T)
	Build(*model.Column)
}

type Multi[T Accepted] struct {
	analyser []Analyser[T]
}

func (m Multi[T]) Read(value *T) {
	for _, a := range m.analyser {
		a.Read(value)
	}
}

func (m Multi[T]) Build(metric *model.Column) {
	for _, a := range m.analyser {
		a.Build(metric)
	}
}
