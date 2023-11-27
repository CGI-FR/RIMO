package metricv2

import (
	"sort"

	"github.com/cgi-fr/rimo/pkg/modelv2"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type String struct {
	sampleSize uint
	distinct   bool
	main       Multi[string]
	byLen      map[int]Multi[string]
}

func NewString(sampleSize uint, countDistinct bool) *String {
	mainAnalyser := []Analyser[string]{
		NewCounter[string](),           // count total, count null, count empty
		NewMinMax[string](),            // store min and max values
		NewSampler[string](sampleSize), // store few samples
	}

	if countDistinct {
		mainAnalyser = append(mainAnalyser, NewDistinct[string]())
	}

	return &String{
		sampleSize: sampleSize,
		distinct:   countDistinct,
		main:       Multi[string]{mainAnalyser},
		byLen:      make(map[int]Multi[string], 0),
	}
}

func (a *String) Read(value *string) {
	a.main.Read(value)

	if value != nil {
		length := len(*value)

		analyser, exists := a.byLen[length]
		if !exists {
			analyser = Multi[string]{
				[]Analyser[string]{
					NewCounter[string](),             // count total, count null, count empty
					NewMinMax[string](),              // store min and max values
					NewSampler[string](a.sampleSize), // store few samples
				},
			}

			if a.distinct {
				analyser.analyser = append(analyser.analyser, NewDistinct[string]())
			}
		}

		analyser.Read(value)

		a.byLen[length] = analyser
	}
}

func (a *String) Build(metric *modelv2.Generic[string]) {
	a.main.Build(metric)

	metric.String = &modelv2.String{
		MinLen:   slices.Min(maps.Keys(a.byLen)),
		MaxLen:   slices.Max(maps.Keys(a.byLen)),
		CountLen: len(a.byLen),
		Lengths:  make([]modelv2.StringLen, 0, len(a.byLen)),
	}

	for length, analyser := range a.byLen {
		lenMetric := modelv2.Generic[string]{}
		analyser.Build(&lenMetric)

		strlen := modelv2.StringLen{
			Length:  length,
			Freq:    float64(lenMetric.Count) / float64(metric.Count),
			Metrics: modelv2.Generic[string]{},
		}
		strlen.Metrics.Count = lenMetric.Count
		strlen.Metrics.Empty = lenMetric.Empty
		strlen.Metrics.Null = lenMetric.Null
		strlen.Metrics.Distinct = lenMetric.Distinct
		strlen.Metrics.Max = lenMetric.Max
		strlen.Metrics.Min = lenMetric.Min
		strlen.Metrics.Samples = lenMetric.Samples
		metric.String.Lengths = append(metric.String.Lengths, strlen)
	}

	sort.Slice(metric.String.Lengths, func(i, j int) bool {
		return metric.String.Lengths[i].Freq > metric.String.Lengths[j].Freq
	})
}
