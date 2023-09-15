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
