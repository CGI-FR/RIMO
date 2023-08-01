package analyse_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
)

func TestColType(t *testing.T) {
	t.Parallel()

	var slice1 []interface{} = []interface{}{1, 2, 3}
	expected1 := "numeric"
	if actual1 := analyse.ColType(slice1); actual1 != expected1 {
		t.Errorf("ColType(%v) = %s; expected %s", slice1, actual1, expected1)
	}

	var slice2 []interface{} = []interface{}{nil, 2, 3}
	expected2 := "numeric"
	if actual2 := analyse.ColType(slice2); actual2 != expected2 {
		t.Errorf("ColType(%v) = %s; expected %s", slice2, actual2, expected2)
	}

	var slice3 []interface{} = []interface{}{nil, "text", nil}
	expected3 := "string"
	if actual3 := analyse.ColType(slice3); actual3 != expected3 {
		t.Errorf("analyse.ColType(%v) = %s; expected %s", slice3, actual3, expected3)
	}

	var slice4 []interface{} = []interface{}{nil, true, false}
	expected4 := "boolean"
	if actual4 := analyse.ColType(slice4); actual4 != expected4 {
		t.Errorf("analyse.ColType(%v) = %s; expected %s", slice4, actual4, expected4)
	}

	var slice5 []interface{} = []interface{}{"text", 2, false}
	expected5 := "unknown"
	if actual5 := analyse.ColType(slice5); actual5 != expected5 {
		t.Errorf("analyse.ColType(%v) = %s; expected %s", slice5, actual5, expected5)
	}

	var slice6 []interface{} = []interface{}{nil, nil, nil}
	expected6 := "unknown"
	if actual6 := analyse.ColType(slice6); actual6 != expected6 {
		t.Errorf("analyse.ColType(%v) = %s; expected %s", slice6, actual6, expected6)
	}
}

func SampleTest(t *testing.T) {
	t.Parallel()

	var slice1 []interface{} = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sample1 := analyse.Sample(slice1, 5)
	sample2 := analyse.Sample(slice1, 5)

	if len(sample1) != 5 {
		t.Errorf("analyse.Sample(%v, 5) = %v; expected %v", slice1, sample1, 5)
	}

	sameOrder := 0
	for i := 0; i < len(sample1); i++ {
		if sample1[i] == sample2[i] {
			sameOrder++
		}
	}

	if sameOrder == len(sample1) {
		t.Errorf("2 analyse.Sample(%v, 5) have same order; most likely expected different", slice1, sample1, 5)
	}

	sample3 := analyse.Sample(slice1, 15)
	if len(sample3) != 15 {
		t.Errorf("analyse.Sample(%v, 15) = %v; expected %v", slice1, sample3, 15)
	}

	var slice2 []interface{} = []interface{}{"Hello", 2, true}
	sample4 := analyse.Sample(slice2, 5)
	if len(sample4) != 3 {
		t.Errorf("analyse.Sample(%v, 5) = %v; expected sample from different type of element", slice2, sample4)
	}
}
