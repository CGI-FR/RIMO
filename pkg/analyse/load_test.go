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

func TestLoadNewFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data := analyse.Load(jsonlNewFormatInputPath, "new")
	fmt.Println(valast.String(data))
}

func TestLoadOldFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data := analyse.Load(jsonlOldFormatInputPath, "old")
	fmt.Println(valast.String(data))
}
