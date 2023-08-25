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

// Bool metric : TrueRatio.
func SetBoolMetric(values []interface{}, metric *model.BoolMetric) error {
	nullCount := 0
	trueCount := 0

	for _, value := range values {
		if value == nil {
			nullCount++

			continue
		}

		boolValue, ok := value.(bool)
		if !ok {
			return fmt.Errorf("%w : expected numeric found %T: %v", ErrValueType, value, value)
		}

		if boolValue {
			trueCount++
		}
	}

	metric.TrueRatio = GetFrequency(trueCount, len(values)-nullCount)

	return nil
}
