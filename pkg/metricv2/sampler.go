package metricv2

import (
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/modelv2"
)

type Sampler[T Accepted] struct {
	size    uint
	count   int
	samples []T
}

func NewSampler[T Accepted](size uint) *Sampler[T] {
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

func (s *Sampler[T]) Build(metric *modelv2.Column) {
	metric.MainMetric.Samples = make([]any, len(s.samples))
	for i, s := range s.samples {
		metric.MainMetric.Samples[i] = s
	}
}
