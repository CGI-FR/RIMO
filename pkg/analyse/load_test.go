package analyse_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/assert"
)

func TestLoadNewFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data, err := analyse.Load(jsonlNewFormat, "new")
	assert.NoError(t, err)

	fmt.Println(valast.String(data))
}

func TestLoadOldFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	data, err := analyse.Load(jsonlPrevFormat, "old")
	assert.NoError(t, err)

	fmt.Println(valast.String(data))
}

func TestEqualityFormat(t *testing.T) {
	t.Helper()
	t.Parallel()

	dataNew, err := analyse.Load(jsonlNewFormat, "new")
	assert.NoError(t, err)

	dataOld, err := analyse.Load(jsonlPrevFormat, "old")
	assert.NoError(t, err)

	if !reflect.DeepEqual(dataNew, dataOld) {
		t.Errorf("Data mismatch: %v != %v", dataNew, dataOld)
	}
}
