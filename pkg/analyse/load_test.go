package analyse_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/hexops/valast"
)

const inputPath = "../../test/data/testcase_newstruct.jsonl"

func TestPipeline(t *testing.T) {
	t.Parallel()

	data := analyse.Load(inputPath)
	fmt.Println(valast.String(data))

	dataMap := analyse.BuildColType(data)
	fmt.Println(valast.String(dataMap))
}

func TestLoad(t *testing.T) {
	t.Parallel()

	data := analyse.Load(inputPath)
	fmt.Println(valast.String(data))
}

func TestBuildColType(t *testing.T) {
	t.Parallel()

	data := analyse.DataMap{
		"Int": analyse.DataCol{
			ColType: "unknown",
			Values:  []interface{}{nil, 2, 3},
		},
		"String": analyse.DataCol{
			ColType: "unknown",
			Values: []interface{}{
				"text1",
				"text2",
			},
		},
		"Float": analyse.DataCol{
			ColType: "unknown",
			Values:  []interface{}{nil, 2.2, 3.3},
		},
		"Bool": analyse.DataCol{
			ColType: "unknown",
			Values:  []interface{}{true, false, true},
		},
		"Empty": analyse.DataCol{
			ColType: "unknown",
			Values:  []interface{}{nil, nil, nil},
		},
	}

	dataMap := analyse.BuildColType(data)

	correctDataMap := analyse.DataMap{
		"Int": analyse.DataCol{
			ColType: "numeric",
			Values:  []interface{}{nil, 2, 3},
		},
		"String": analyse.DataCol{
			ColType: "string",
			Values: []interface{}{
				"text1",
				"text2",
			},
		},
		"Float": analyse.DataCol{
			ColType: "numeric",
			Values:  []interface{}{nil, 2.2, 3.3},
		},
		"Bool": analyse.DataCol{
			ColType: "boolean",
			Values:  []interface{}{true, false, true},
		},
		"Empty": analyse.DataCol{
			ColType: "unknown",
			Values:  []interface{}{nil, nil, nil},
		},
	}
	// Assert equality
	assert.Equal(t, correctDataMap, dataMap)
}
