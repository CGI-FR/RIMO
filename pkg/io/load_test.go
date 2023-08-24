package io_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/pkg/io"
	"github.com/hexops/valast"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	path := filepath.Join(dataDir, "data1/data_input.jsonl")

	data, err := io.Load(path)
	require.NoError(t, err)

	fmt.Println(valast.String(data))
}
