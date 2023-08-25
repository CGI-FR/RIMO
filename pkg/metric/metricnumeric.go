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
	"fmt"

	"github.com/cgi-fr/rimo/pkg/model"
)

func SetNumericMetric(values []interface{}, metric *model.NumericMetric) error {
	nonNullCount := 0

	value := GetFirstValue(values)

	floatValue, ok := value.(float64)
	if !ok {
		return fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
	}

	min := floatValue
	max := floatValue
	sum := 0.0

	for _, value := range values {
		floatValue, ok := value.(float64)
		if !ok {
			if value == nil {
				continue
			}

			return fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
		}

		sum += floatValue
		nonNullCount++

		if floatValue > max {
			max = floatValue
		}

		if floatValue < min {
			min = floatValue
		}
	}

	mean := sum / float64(nonNullCount)

	metric.Min = min
	metric.Max = max
	metric.Mean = mean

	return nil
}
