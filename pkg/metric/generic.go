package metric

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/model"
)

func SetGenericMetric(values []interface{}, met *model.GenericMetric) error {
	sample, err := Sample(values, model.SampleSize)
	if err != nil {
		return fmt.Errorf("error computing sample: %w", err)
	}

	met.Count = len(values)
	met.Unique = CountUnique(values)
	met.Empty = CountEmpty(values)
	met.Sample = sample

	return nil
}

func CountEmpty(values []interface{}) int {
	empty := 0

	for _, value := range values {
		if value == nil {
			empty++
		}
	}

	return empty
}

var ErrEmptySlice = errors.New("slice is empty")

// Return a sample of size sampleSize from values.
func Sample[T comparable](values []T, sampleSize int) ([]T, error) {
	uniqueValues := Unique(values)

	if sampleSize >= len(uniqueValues) {
		return uniqueValues, nil
	}

	sample := make([]T, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = uniqueValues[rand.Intn(len(uniqueValues)-1)] //nolint:gosec
	}

	return sample, nil
}

func CountUnique[T comparable](values []T) int {
	unique := make(map[T]bool)

	for _, value := range values {
		if isNil(value) {
			continue
		}

		unique[value] = true
	}

	return len(unique)
}

func Unique[T comparable](values []T) []T {
	unique := make(map[T]bool)

	for _, value := range values {
		if isNil(value) {
			continue
		}

		unique[value] = true
	}

	uniqueValues := make([]T, 0, len(unique))
	for value := range unique {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}

func isNil[T comparable](v T) bool {
	return v == *new(T)
}

// // Return a sample of size sampleSize from values.
// func Sample(values []interface{}, sampleSize int) ([]interface{}, error) {
// 	uniqueValues := Unique(values)

// 	if sampleSize >= len(uniqueValues) {
// 		return uniqueValues, nil
// 	}

// 	sample := make([]interface{}, sampleSize)
// 	for i := 0; i < sampleSize; i++ {
// 		sample[i] = uniqueValues[rand.Intn(len(uniqueValues)-1)] //nolint:gosec
// 	}

// 	return sample, nil
// }

// func CountUnique(values []interface{}) int {
// 	unique := make(map[interface{}]bool)

// 	for _, value := range values {
// 		if value == nil {
// 			continue
// 		}

// 		unique[value] = true
// 	}

// 	return len(unique)
// }

// func Unique(values []interface{}) []interface{} {
// 	unique := make(map[interface{}]bool)

// 	for _, value := range values {
// 		if value == nil {
// 			continue
// 		}

// 		unique[value] = true
// 	}

// 	uniqueValues := make([]interface{}, 0, len(unique))
// 	for value := range unique {
// 		uniqueValues = append(uniqueValues, value)
// 	}

// 	return uniqueValues
// }
