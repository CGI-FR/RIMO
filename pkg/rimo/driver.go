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

package rimo

import (
	"fmt"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"

	"github.com/rs/zerolog/log"
)

func AnalyseBase(reader Reader, writer Writer) error {
	// log.Logger = zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	baseName := reader.BaseName()

	// log.Debug().Msgf("Processing [%s base]", baseName)

	base := model.NewBase(baseName)

	for reader.Next() { // itère colonne par colonne
		colValues, colName, tableName, err := reader.Value()
		if err != nil {
			return fmt.Errorf("failed to get column value : %w", err)
		}

		column, err := metric.ComputeMetric(colName, colValues)
		if err != nil {
			return fmt.Errorf("failed to compute column : %w", err)
		}

		log.Debug().Msgf("Processing [%s base][%s table][%s column]", baseName, tableName, column.Name)
		// log.Debug().Msg(valast.String(column))

		base.AddColumn(column, tableName)
	}

	base.SortBase()

	// log.Debug().Msg("---------- Finish processing base :")
	// log.Debug().Msg(valast.String(*base))
	// log.Debug().Msg("----------")

	err := writer.Export(base)
	if err != nil {
		return fmt.Errorf("failed to export base : %w", err)
	}

	return nil
}
