package analyse_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/hexops/valast"
)

func TestLoadNewFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data := analyse.Load(jsonlNewFormat, "new")
	fmt.Println(valast.String(data))
}

func TestLoadOldFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data := analyse.Load(jsonlPrevFormat, "old")
	fmt.Println(valast.String(data))
}

func TestEqualityFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	dataNew := analyse.Load(jsonlNewFormat, "new")
	dataOld := analyse.Load(jsonlPrevFormat, "old")

	if !reflect.DeepEqual(dataNew, dataOld) {
		t.Errorf("Data mismatch: %v != %v", dataNew, dataOld)
	}
}
