package metric

type Numeric struct {
	Multi[float64]
}

func NewNumeric(sampleSize uint, countDistinct bool) *Numeric {
	mainAnalyser := []Analyser[float64]{
		NewCounter[float64](),           // count total, count null, count empty
		NewMinMax[float64](),            // store min and max values
		NewSampler[float64](sampleSize), // store few samples
		NewMean(),                       // calculate running mean
	}

	if countDistinct {
		mainAnalyser = append(mainAnalyser, NewDistinct[float64]())
	}

	return &Numeric{
		Multi: Multi[float64]{mainAnalyser},
	}
}
