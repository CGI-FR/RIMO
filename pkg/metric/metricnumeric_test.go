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

package metric_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/stretchr/testify/assert"
)

func TestNumericMetric(t *testing.T) {
	t.Parallel()

	values := []interface{}{1.0, 2.0, 3.0, nil}
	expectedMetric := rimo.NumericMetric{
		Min:  1,
		Max:  3,
		Mean: 2,
	}

	actualMetric := rimo.NumericMetric{} //nolint:exhaustruct

	err := metric.SetNumericMetric(values, &actualMetric)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, expectedMetric, actualMetric)
}
