package metric

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Sampler implement a basic sampling algorithm.
//
//	see: https://en.wikipedia.org/wiki/Reservoir_sampling#Simple:_Algorithm_R)
type Sampler[T constraints.Ordered] struct {
	size  uint
	count int
	data  []T
}

func NewSampler[T constraints.Ordered](size uint) *Sampler[T] {
	return &Sampler[T]{
		size:  size,
		count: 0,
		data:  make([]T, 0, size),
	}
}

func (s *Sampler[T]) Add(data T) {
	s.count++

	if len(s.data) < int(s.size) {
		s.data = append(s.data, data)

		return
	}

	index := rand.Intn(s.count) //nolint:gosec
	if index < int(s.size) {
		s.data[index] = data
	}
}

func (s *Sampler[T]) Data() []T {
	return s.data
}
