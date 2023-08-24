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
