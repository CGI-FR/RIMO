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
