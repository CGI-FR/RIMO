package analyse_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/hexops/valast"
	"gopkg.in/yaml.v3"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

const (
	TestDir = "../../test/analyseTest/"
)

var (
	JsonlNewFormat  = filepath.Join(TestDir, "/input/testcase_newstruct.jsonl")
	JsonlPrevFormat = filepath.Join(TestDir, "/input/testcase_prevstruct.jsonl")
)

func TestAnalyse(t *testing.T) {
	t.Parallel()

	inputList := []string{JsonlNewFormat}
	outputPath := filepath.Join(TestDir, "/output/rimo_output.yaml")
	analyse.Analyse(inputList, outputPath)

	// Load output file
	file, err := os.Open(outputPath)
	assert.NoError(t, err)

	t.Cleanup(func() {
		file.Close()
	})

	var actualOutput string

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	assert.NoError(t, err)

	actualOutput = buf.String()
	t.Log(actualOutput)

	// Load expected output file
	testPath := filepath.Join(TestDir, "/expected/rimo_output.yaml")
	expectedFile, err := os.Open(testPath)
	assert.NoError(t, err)

	t.Cleanup(func() {
		expectedFile.Close()
	})

	var expectedOutput string

	buf = new(bytes.Buffer)
	_, err = buf.ReadFrom(expectedFile)
	assert.NoError(t, err)

	expectedOutput = buf.String()
	t.Log(expectedOutput)

	// Compare the expected output and actual output
	// WILL FAIL for now as sample is compared too.
	t.Run("TestAnalyseFileComparison", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, expectedOutput, actualOutput)
	})

	t.Run("TestAnalyseObjectComparison", func(t *testing.T) {
		t.Parallel()

		// Load file in a model.Base.
		decoder := yaml.NewDecoder(file)

		var actualOutputBase model.Base
		err = decoder.Decode(&actualOutputBase)
		if err != nil {
			t.Errorf("error while decoding yaml file: %v", err)
		}

		// Load expected file in a model.Base.
		decoder = yaml.NewDecoder(expectedFile)

		var expectedOutputBase model.Base
		err = decoder.Decode(&expectedOutputBase)
		if err != nil {
			t.Errorf("error while decoding yaml file: %v", err)
		}

		// Remove sample fields from both model.Base.
		removeSampleFields(&actualOutputBase)
		removeSampleFields(&expectedOutputBase)

		// Compare the expected output and actual output except all sample fields.
		if !reflect.DeepEqual(expectedOutputBase, actualOutputBase) {
			t.Errorf("output does not match expected output")
		}

		// Print actual output.
		t.Log(valast.String(actualOutputBase))
	})
}

func TestGetBaseName(t *testing.T) {
	t.Helper()
	t.Parallel()

	path := "path/to/dir/basename_tablename.jsonl"
	expected := "basename"

	if baseName, err := analyse.GetBaseName(path); baseName != expected || err != nil {
		t.Errorf("GetBaseName(%q) = (%q, %v), expected (%q, %v)", path, baseName, err, expected, nil)
	}

	path2 := "basename_tablename.jsonl"
	expected2 := "basename"

	if baseName, err := analyse.GetBaseName(path2); baseName != expected2 || err != nil {
		t.Errorf("GetBaseName(%q) = (%q, %v), expected (%q, %v)", path2, baseName, err, expected2, nil)
	}

	invalidPath := ""

	_, err := analyse.GetBaseName(invalidPath)
	if !errors.Is(err, analyse.ErrNonExtractibleValue) {
		t.Errorf("expected error %v, but got %v", analyse.ErrNonExtractibleValue, err)
	}
}

func TestGetTableName(t *testing.T) {
	t.Helper()
	t.Parallel()

	path := "path/to/dir/basename_tablename.jsonl"
	expected := "tablename"

	if tableName, err := analyse.GetTableName(path); tableName != expected || err != nil {
		t.Errorf("GetTableName(%q) = (%q, %v), expected (%q, %v)", path, tableName, err, expected, nil)
	}

	path2 := "basename_tablename.jsonl"
	expected2 := "tablename"

	if tableName, err := analyse.GetTableName(path2); tableName != expected2 || err != nil {
		t.Errorf("GetTableName(%q) = (%q, %v), expected (%q, %v)", path2, tableName, err, expected2, nil)
	}

	invalidPath := ""

	_, err := analyse.GetTableName(invalidPath)
	if !errors.Is(err, analyse.ErrNonExtractibleValue) {
		t.Errorf("expected error %v, but got %v", analyse.ErrNonExtractibleValue, err)
	}
}

func removeSampleFields(base *model.Base) {
	for _, table := range base.Tables {
		for _, column := range table.Columns {
			column.MainMetric.Sample = nil

			if column.Type == "string" {
				column.StringMetric.LeastFreqSample = nil
			}
		}
	}
}
