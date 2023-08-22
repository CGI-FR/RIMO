package analyse_test

import (
	"fmt"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Helper()
	t.Parallel()

	data, err := analyse.Load(data1Path)
	require.NoError(t, err)

	fmt.Println(valast.String(data))
}
