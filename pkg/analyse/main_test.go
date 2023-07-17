// package name analyse_test and private function : https://stackoverflow.com/questions/24622388/how-to-test-a-unexported-private-function-in-go-golang
// linter Warning :
// unused func, type assertion in _test, braces,
// whiteline / comments below braces in metrics.go

package analyse_test

import (
	"fmt"
	"testing"

	"github.com/hexops/valast"
	// "github.com/cgi-fr/rimo/pkg/analyse"
)

const inputPath = "../../test/data/testcase_newstruct.jsonl"

func TestPipeline(t *testing.T) {
	t.Parallel()
	data := analyse.load(inputPath)
	fmt.Println(valast.String(data))

	dataMap := analyse.buildColType(data)
	fmt.Println(valast.String(dataMap))
}

func TestLoad(t *testing.T) {
	t.Parallel()
	data := load(inputPath)
	fmt.Println(valast.String(data))
}

func TestBuildColType(t *testing.T) {
	t.Parallel()
	data := dataMap{
		"Int": dataCol{
			colType: "unknown",
			values:  []interface{}{nil, 2, 3},
		},
		"Address": dataCol{
			colType: "unknown",
			values: []interface{}{
				"PSC 4713, Box 9649 APO AA 43433",
				"095 Jennifer Turnpike Castrobury, NY 98111",
				"06210 David Court South Kimberly, IL 10236",
			},
		},
		"Float": dataCol{
			colType: "unknown",
			values:  []interface{}{nil, 2.2, 3.3},
		},
		"Bool": dataCol{
			colType: "unknown",
			values:  []interface{}{true, false, true},
		},
		"Empty": dataCol{
			colType: "unknown",
			values:  []interface{}{nil, nil, nil},
		},
	}

	dataMap := buildColType(data)
	fmt.Println(valast.String(dataMap))
}
