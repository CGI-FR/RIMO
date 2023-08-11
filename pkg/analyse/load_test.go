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

	data := analyse.Load(JsonlNewFormat, "new")
	fmt.Println(valast.String(data))
}

func TestLoadOldFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data := analyse.Load(JsonlPrevFormat, "old")
	fmt.Println(valast.String(data))
}

func TestEqualityFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	dataNew := analyse.Load(JsonlNewFormat, "new")
	dataOld := analyse.Load(JsonlPrevFormat, "old")

	if !reflect.DeepEqual(dataNew, dataOld) {
		t.Errorf("Data mismatch: %v != %v", dataNew, dataOld)
	}
}
