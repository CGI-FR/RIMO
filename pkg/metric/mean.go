package metric

import "github.com/cgi-fr/rimo/pkg/model"

type Mean struct {
	count uint
	mean  float64
}

func NewMean() *Mean {
	return &Mean{
		count: 0,
		mean:  0,
	}
}

func (a *Mean) Read(value *float64) {
	if value == nil {
		return
	}

	a.count++

	a.mean += (*value - a.mean) / float64(a.count)
}

func (a *Mean) Build(metric *model.Column) {
	metric.NumericMetric = &model.Numeric{
		Mean: a.mean,
	}
}
