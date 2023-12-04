package metricv2

type Bool struct {
	Multi[bool]
}

func NewBool(sampleSize uint, countDistinct bool) *Bool {
	mainAnalyser := []Analyser[bool]{
		NewCounter[bool](), // count total, count null
		NewTrueRatio(),     // calculate true ratio
	}

	if countDistinct {
		mainAnalyser = append(mainAnalyser, NewDistinct[bool]())
	}

	return &Bool{
		Multi: Multi[bool]{mainAnalyser},
	}
}
