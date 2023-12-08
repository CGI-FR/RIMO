package metric

import "github.com/cgi-fr/rimo/pkg/model"

type TrueRatio struct {
	countTrue uint
	count     uint
}

func NewTrueRatio() *TrueRatio {
	return &TrueRatio{
		countTrue: 0,
		count:     0,
	}
}

func (a *TrueRatio) Read(value *bool) {
	if value == nil {
		return
	}

	a.count++

	if *value {
		a.countTrue++
	}
}

func (a *TrueRatio) Build(metric *model.Column) {
	metric.BoolMetric = &model.Bool{
		TrueRatio: float64(a.countTrue) / float64(a.count),
	}
}
