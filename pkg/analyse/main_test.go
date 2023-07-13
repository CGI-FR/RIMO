package analyse

import (
	"fmt"
	"testing"

	"github.com/hexops/valast"
)

const input_path = "../../test/data/testcase_newstruct.jsonl"

func TestPipeline(t *testing.T) {
	data := load(input_path)
	fmt.Println(valast.String(data))
	dataMap := buildColType(data)
	fmt.Println(valast.String(dataMap))
}

func TestLoad(t *testing.T) {
	data := load(input_path)
	fmt.Println(valast.String(data))
}
func TestBuildColType(t *testing.T) {
	data := dataMap{
		"Int": dataCol{
			colType: "unknown",
			values:  []interface{}{nil, 2, 3},
		},
		"Address": dataCol{
			colType: "unknown",
			values:  []interface{}{"PSC 4713, Box 9649 APO AA 43433", "095 Jennifer Turnpike Castrobury, NY 98111", "06210 David Court South Kimberly, IL 10236"},
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
	// if err != nil {
	// 	t.Errorf("Error: %v", err)
	// }
	fmt.Println(valast.String(dataMap))
}
