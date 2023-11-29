package metricv2

import (
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/modelv2"
	"golang.org/x/exp/constraints"
)

type Sampler[T constraints.Ordered] struct {
	size    uint
	count   int
	samples []T
}

func NewSampler[T constraints.Ordered](size uint) *Sampler[T] {
	return &Sampler[T]{
		size:    size,
		count:   0,
		samples: make([]T, 0, size),
	}
}

func (s *Sampler[T]) Read(value *T) {
	if value != nil {
		s.count++

		if len(s.samples) < int(s.size) {
			s.samples = append(s.samples, *value)

			return
		}

		index := rand.Intn(s.count) //nolint:gosec
		if index < int(s.size) {
			s.samples[index] = *value
		}
	}
}

func (s *Sampler[T]) Build(metric *modelv2.Column[T]) {
	metric.MainMetric.Samples = s.samples
}
