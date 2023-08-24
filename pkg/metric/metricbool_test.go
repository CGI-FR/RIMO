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
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBooleanMetric(t *testing.T) {
	t.Parallel()

	values := []interface{}{true, true, false, false}
	expectedMetric := model.BoolMetric{
		TrueRatio: float64(1) / float64(2),
	}

	actualMetric := model.BoolMetric{} //nolint:exhaustruct
	err := metric.SetBoolMetric(values, &actualMetric)
	require.NoError(t, err)

	assert.Equal(t, expectedMetric, actualMetric)
}
