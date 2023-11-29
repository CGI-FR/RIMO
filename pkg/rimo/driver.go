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

	"github.com/cgi-fr/rimo/pkg/metricv2"
	"github.com/cgi-fr/rimo/pkg/modelv2"

	"github.com/rs/zerolog/log"
)

func AnalyseBase(reader Reader, writer Writer) error {
	// log.Logger = zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	baseName := reader.BaseName()

	// log.Debug().Msgf("Processing [%s base]", baseName)

	base := modelv2.NewBase(baseName)

	for reader.Next() { // itère colonne par colonne
		valreader, err := reader.Col()
		if err != nil {
			return fmt.Errorf("failed to get column reader : %w", err)
		}

		nilcount := 0

		for valreader.Next() {
			val, err := valreader.Value()
			if err != nil {
				return fmt.Errorf("failed to read value : %w", err)
			}

			log.Debug().Msgf("Processing [%s base][%s table][%s column]", baseName, valreader.TableName(), valreader.ColName())

			switch valtyped := val.(type) {
			case string:
				col, err := AnalyseString(nilcount, valtyped, valreader)
				if err != nil {
					return fmt.Errorf("failed to analyse column : %w", err)
				}

				table, exists := base.Tables[valreader.TableName()]
				if !exists {
					table = modelv2.Table{
						Columns: []modelv2.Column{},
					}
				}

				table.Columns = append(table.Columns, col)

				base.Tables[valreader.TableName()] = table
			}
		}
	}

	// base.SortBase()

	// log.Debug().Msg("---------- Finish processing base :")
	// log.Debug().Msg(valast.String(*base))
	// log.Debug().Msg("----------")

	err := writer.Export(base)
	if err != nil {
		return fmt.Errorf("failed to export base : %w", err)
	}

	return nil
}

func AnalyseString(nilcount int, firstValue string, reader ColReader) (modelv2.Column, error) {
	column := modelv2.Column{
		Name:          reader.ColName(),
		Type:          "string",
		Config:        modelv2.Config{},  //nolint:exhaustruct
		MainMetric:    modelv2.Generic{}, //nolint:exhaustruct
		StringMetric:  &modelv2.String{}, //nolint:exhaustruct
		NumericMetric: nil,
		BoolMetric:    nil,
	}

	analyser := metricv2.NewString(5, true)

	for i := 0; i < nilcount; i++ {
		analyser.Read(nil)
	}

	analyser.Read(&firstValue)

	for reader.Next() {
		val, err := reader.Value()
		if err != nil {
			return column, fmt.Errorf("failed to read value : %w", err)
		}

		switch valtyped := val.(type) {
		case string:
			analyser.Read(&valtyped)
		default:
			return column, fmt.Errorf("invalue value type : %w", err)
		}
	}

	analyser.Build(&column)

	return column, nil
}
