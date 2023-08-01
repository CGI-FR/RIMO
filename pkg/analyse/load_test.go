package analyse_test

import (
	"fmt"
	"testing"

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
