// Copyright (C) 2023 CGI France
//
// This file is part of RIMO.
//
// RIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// RIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with RIMO.  If not, see <http://www.gnu.org/licenses/>.

package metric

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/cgi-fr/rimo/pkg/model"
)

var ErrEmptySlice = errors.New("slice is empty")

func SetGenericMetric(values []interface{}, metric *model.GenericMetric) error {
	sample, err := Sample(values, model.SampleSize)
	if err != nil {
		return fmt.Errorf("error computing sample: %w", err)
	}

	metric.Count = len(values)
	metric.Unique = CountUnique(values)
	metric.Empty = CountEmpty(values)
	metric.Sample = sample

	return nil
}

func CountEmpty[T comparable](values []T) int {
	empty := 0

	for _, value := range values {
		if isNil(value) {
			empty++
		}
	}

	return empty
}

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
