package infra_test

import (
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/internal/infra"
	"github.com/cgi-fr/rimo/pkg/rimo"
	"github.com/stretchr/testify/require"
)

const (
	testdataDir = "../../testdata/"
)

// Test RIMO pipeline with FilesReader, JSONLinesLoader and YAMLWriter.
func TestPipeline(t *testing.T) {
	t.Parallel()

	inputPath := filepath.Join(testdataDir, "data1/data_input.jsonl")

	reader, err := infra.FilesReaderFactory([]string{inputPath})
	require.NoError(t, err)

	writer := infra.StdoutWriterFactory()

	err = rimo.AnalyseBase(reader, writer)
	require.NoError(t, err)
}

// var (
// 	Readers []*rimo.Reader
// 	Writers []*rimo.Writer
// )

// // List of implemented readers and writers.
// func GetReaders(filepathList []string) []*rimo.Reader {
// 	filesReader, err := infra.FilesReaderFactory(filepathList)
// 	if err != nil {
// 		panic(err)
// 	}

// 	Readers = []*rimo.Reader{filesReader}

// 	return Readers
// }

// func GetWriters() []*rimo.Writer {
// 	yamlWriter := infra.YAMLWriterFactory("../../testdata/data1/data_output.yaml")

// 	Writers = []*rimo.Writer{yamlWriter, infra.StdoutWriter{}}

// 	return Writers
// }

// func TestInterface(t *testing.T) {
// 	t.Parallel()

// 	Writers = GetWriters()
// 	Readers = GetReaders([]string{"../../testdata/data1/data_input.jsonl"})
// 	// Assert that all readers and writers implement the Reader and Writer interfaces.
// 	for _, reader := range Readers {
// 		var _ rimo.Reader = (reader)(nil)
// 	}
// 	for _, writer := range Writers {
// 		var _ rimo.Reader = (writer)(nil)
// 	}

// 	// Assert that all combinations of readers and writers can be used in the pipeline.
// 	for _, reader := range Readers {
// 		for _, writer := range Writers {
// 			err := rimo.AnalyseBase(reader, writer)
// 			require.NoError(t, err)
// 		}
// 	}
// }
