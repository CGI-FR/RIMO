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

func (a *String) Build(metric *modelv2.Column) {
	a.main.Build(metric)

	metric.StringMetric = &modelv2.String{
		MinLen:   slices.Min(maps.Keys(a.byLen)),
		MaxLen:   slices.Max(maps.Keys(a.byLen)),
		CountLen: len(a.byLen),
		Lengths:  make([]modelv2.StringLen, 0, len(a.byLen)),
	}

	for length, analyser := range a.byLen {
		lenMetric := modelv2.Column{} //nolint:exhaustruct
		analyser.Build(&lenMetric)

		strlen := modelv2.StringLen{
			Length:  length,
			Freq:    float64(lenMetric.MainMetric.Count) / float64(metric.MainMetric.Count),
			Metrics: modelv2.Generic{}, //nolint:exhaustruct
		}
		strlen.Metrics.Count = lenMetric.MainMetric.Count
		strlen.Metrics.Empty = lenMetric.MainMetric.Empty
		strlen.Metrics.Null = lenMetric.MainMetric.Null
		strlen.Metrics.Distinct = lenMetric.MainMetric.Distinct
		strlen.Metrics.Max = lenMetric.MainMetric.Max
		strlen.Metrics.Min = lenMetric.MainMetric.Min
		strlen.Metrics.Samples = lenMetric.MainMetric.Samples
		metric.StringMetric.Lengths = append(metric.StringMetric.Lengths, strlen)
	}

	sort.Slice(metric.StringMetric.Lengths, func(i, j int) bool {
		return metric.StringMetric.Lengths[i].Freq > metric.StringMetric.Lengths[j].Freq
	})
}
