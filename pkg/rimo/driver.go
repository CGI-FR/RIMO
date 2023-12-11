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
	"sort"

	"github.com/cgi-fr/rimo/pkg/metric"
	"github.com/cgi-fr/rimo/pkg/model"

	"github.com/rs/zerolog/log"
)

type Driver struct {
	SampleSize uint
	Distinct   bool
}

func (d Driver) AnalyseBase(reader Reader, writer Writer) (err error) {
	if err := reader.Open(); err != nil {
		return fmt.Errorf("failed to open column reader : %w", err)
	}

	defer func() {
		if localerr := reader.Close(); err != nil {
			err = localerr
		}
	}()

	baseName := reader.BaseName()

	base := model.NewBase(baseName)
	tables := map[string]model.Table{}

	for reader.Next() { // itère colonne par colonne
		if err := d.analyse(reader, tables); err != nil {
			return err
		}
	}

	for _, table := range tables {
		sort.SliceStable(table.Columns, func(i, j int) bool {
			return table.Columns[i].Name < table.Columns[j].Name
		})

		base.Tables = append(base.Tables, table)
	}

	sort.SliceStable(base.Tables, func(i, j int) bool {
		return base.Tables[i].Name < base.Tables[j].Name
	})

	err = writer.Export(base)
	if err != nil {
		return fmt.Errorf("failed to export base : %w", err)
	}

	return nil
}

//nolint:funlen,cyclop
func (d Driver) analyse(reader Reader, tables map[string]model.Table) (err error) {
	valreader, err := reader.Col()
	if err != nil {
		return fmt.Errorf("failed to get column reader : %w", err)
	}

	if err := valreader.Open(); err != nil {
		return fmt.Errorf("failed to open value reader : %w", err)
	}

	defer func() {
		if localerr := reader.Close(); err != nil {
			err = localerr
		}
	}()

	nilcount := 0

	for valreader.Next() {
		val, err := valreader.Value()
		if err != nil {
			return fmt.Errorf("failed to read value : %w", err)
		}

		log.Debug().Msgf("Processing [%s][%s][%s]", reader.BaseName(), valreader.TableName(), valreader.ColName())

		switch valtyped := val.(type) {
		case string:
			col, err := d.AnalyseString(nilcount, valtyped, valreader)
			if err != nil {
				return fmt.Errorf("failed to analyse column : %w", err)
			}

			table, exists := tables[valreader.TableName()]
			if !exists {
				table = model.Table{
					Name:    valreader.TableName(),
					Columns: []model.Column{},
				}
			}

			table.Columns = append(table.Columns, col)

			tables[valreader.TableName()] = table
		case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			col, err := d.AnalyseNumeric(nilcount, valtyped, valreader)
			if err != nil {
				return fmt.Errorf("failed to analyse column : %w", err)
			}

			table, exists := tables[valreader.TableName()]
			if !exists {
				table = model.Table{
					Name:    valreader.TableName(),
					Columns: []model.Column{},
				}
			}

			table.Columns = append(table.Columns, col)

			tables[valreader.TableName()] = table
		case bool:
			col, err := d.AnalyseBool(nilcount, valtyped, valreader)
			if err != nil {
				return fmt.Errorf("failed to analyse column : %w", err)
			}

			table, exists := tables[valreader.TableName()]
			if !exists {
				table = model.Table{
					Name:    valreader.TableName(),
					Columns: []model.Column{},
				}
			}

			table.Columns = append(table.Columns, col)

			tables[valreader.TableName()] = table
		case nil:
			nilcount++
		}
	}

	return nil
}

func (d Driver) AnalyseString(nilcount int, firstValue string, reader ColReader) (model.Column, error) {
	column := model.Column{
		Name:          reader.ColName(),
		Type:          "string",
		Config:        model.Config{},  //nolint:exhaustruct
		MainMetric:    model.Generic{}, //nolint:exhaustruct
		StringMetric:  &model.String{}, //nolint:exhaustruct
		NumericMetric: nil,
		BoolMetric:    nil,
	}

	analyser := metric.NewString(d.SampleSize, d.Distinct)

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
		case nil:
			analyser.Read(nil)
		default:
			return column, fmt.Errorf("invalue value type : %w", err)
		}
	}

	analyser.Build(&column)

	return column, nil
}

func (d Driver) AnalyseNumeric(nilcount int, firstValue any, reader ColReader) (model.Column, error) {
	column := model.Column{
		Name:          reader.ColName(),
		Type:          "numeric",
		Config:        model.Config{},  //nolint:exhaustruct
		MainMetric:    model.Generic{}, //nolint:exhaustruct
		StringMetric:  nil,
		NumericMetric: &model.Numeric{}, //nolint:exhaustruct
		BoolMetric:    nil,
	}

	analyser := metric.NewNumeric(d.SampleSize, d.Distinct)

	for i := 0; i < nilcount; i++ {
		analyser.Read(nil)
	}

	valtyped, err := GetFloat64(firstValue)
	if err != nil {
		return column, fmt.Errorf("failed to read value : %w", err)
	}

	analyser.Read(valtyped)

	for reader.Next() {
		val, err := reader.Value()
		if err != nil {
			return column, fmt.Errorf("failed to read value : %w", err)
		}

		valtyped, err := GetFloat64(val)
		if err != nil {
			return column, fmt.Errorf("failed to read value : %w", err)
		}

		analyser.Read(valtyped)
	}

	analyser.Build(&column)

	return column, nil
}

func (d Driver) AnalyseBool(nilcount int, firstValue bool, reader ColReader) (model.Column, error) {
	column := model.Column{
		Name:          reader.ColName(),
		Type:          "bool",
		Config:        model.Config{},  //nolint:exhaustruct
		MainMetric:    model.Generic{}, //nolint:exhaustruct
		StringMetric:  nil,
		NumericMetric: nil,
		BoolMetric:    &model.Bool{}, //nolint:exhaustruct
	}

	analyser := metric.NewBool(d.SampleSize, d.Distinct)

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
		case bool:
			analyser.Read(&valtyped)
		case nil:
			analyser.Read(nil)
		default:
			return column, fmt.Errorf("invalue value type : %w", err)
		}
	}

	analyser.Build(&column)

	return column, nil
}

//nolint:cyclop
func GetFloat64(value any) (*float64, error) {
	var converted float64

	switch valtyped := value.(type) {
	case float64:
		converted = valtyped
	case float32:
		converted = float64(valtyped)
	case int:
		converted = float64(valtyped)
	case int8:
		converted = float64(valtyped)
	case int16:
		converted = float64(valtyped)
	case int32:
		converted = float64(valtyped)
	case int64:
		converted = float64(valtyped)
	case uint:
		converted = float64(valtyped)
	case uint8:
		converted = float64(valtyped)
	case uint16:
		converted = float64(valtyped)
	case uint32:
		converted = float64(valtyped)
	case uint64:
		converted = float64(valtyped)
	case nil:
		return nil, nil //nolint:nilnil
	default:
		return nil, fmt.Errorf("%w : %T", ErrInvalidValueType, value)
	}

	return &converted, nil
}
