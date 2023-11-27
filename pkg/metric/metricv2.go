package metric

import "golang.org/x/exp/constraints"

type Analyser[T constraints.Ordered] interface {
	Read(*T)
}

type Stateless[T constraints.Ordered] interface {
	Analyser[T]
	CountTotal() uint
	CountNulls() uint
	CountEmpty() uint
	Min() *T
	Max() *T
	Samples() []T
}

type Statefull[T constraints.Ordered] interface {
	Stateless[T]
	CountDistinct() uint
}

type Counter[T constraints.Ordered] struct {
	countTotal uint
	countNulls uint
	countEmpty uint
	min        *T
	max        *T
	samples    *Sampler[T]
	zero       T
}

func NewCounter[T constraints.Ordered](samplerSize uint) *Counter[T] {
	return &Counter[T]{
		countTotal: 0,
		countNulls: 0,
		countEmpty: 0,
		samples:    NewSampler[T](samplerSize),
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

	if value != nil {
		c.samples.Add(*value)

		if c.min == nil {
			c.min = value
		}

		if c.max == nil {
			c.max = value
		}

		if *value < *c.min {
			c.min = value
		} else if *value > *c.max {
			c.max = value
		}
	}
}

// CountEmpty implements Stateless.
func (c *Counter[T]) CountEmpty() uint {
	return c.countEmpty
}

// CountNulls implements Stateless.
func (c *Counter[T]) CountNulls() uint {
	return c.countNulls
}

// CountTotal implements Stateless.
func (c *Counter[T]) CountTotal() uint {
	return c.countTotal
}

// Samples implements Stateless.
func (c *Counter[T]) Samples() []T {
	return c.samples.Data()
}

// Min implements Stateless.
func (c *Counter[T]) Min() *T {
	return c.min
}

// Max implements Stateless.
func (c *Counter[T]) Max() *T {
	return c.max
}

type Distinctcounter[T constraints.Ordered] struct {
	Counter[T]
	values map[T]int
}

func NewDistinctCounter[T constraints.Ordered](samplerSize uint) Statefull[T] {
	return &Distinctcounter[T]{
		Counter: Counter[T]{
			countTotal: 0,
			countNulls: 0,
			countEmpty: 0,
			samples:    NewSampler[T](samplerSize),
			zero:       *new(T),
		},
		values: make(map[T]int, 1024), //nolint:gomnd
	}
}

// Read implements Statefull.
func (c *Distinctcounter[T]) Read(value *T) {
	c.Counter.Read(value)

	if value != nil {
		c.values[*value] = 0
	}
}

// CountDistinct implements Statefull.
func (c *Distinctcounter[T]) CountDistinct() uint {
	return uint(len(c.values))
}
