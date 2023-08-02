package analyse_test

import (
	"fmt"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/hexops/valast"
)

const (
	jsonlNewFormatInputPath = "../../test/data/testcase_newstruct.jsonl"
	jsonlOldFormatInputPath = "../../test/data/testcase_data.jsonl"
)

func TestPipeline(t *testing.T) {
	t.Parallel()

	data := analyse.Load(jsonlNewFormatInputPath)
	fmt.Println(valast.String(data))

	// dataMap := analyse.ColType(data)
	// fmt.Println(valast.String(dataMap))
}
