package analyse_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cgi-fr/rimo/pkg/analyse"
	"github.com/cgi-fr/rimo/pkg/model"
	"github.com/hexops/valast"
	"gopkg.in/yaml.v3"
)

const (
	jsonlNewFormatInputPath = "../../test/data/testcase_newstruct.jsonl"
	jsonlOldFormatInputPath = "../../test/data/testcase_data.jsonl"
	outputDir               = "../../test/data/outputTest"
	outputFileNameAnalyse   = "testcase_newstruct.jsonl"

	yamlData = `
	database: ""
	tables:
	  - name: testcase_data
		columns:
		  - name: address
			type: object
			concept: ""
			constraint: []
			confidential: false
			sample:
			  - 2035 Simmons Islands Heatherchester, IN 46152
			  - 38432 Moreno Turnpike Garrettland, TN 72939
			  - 25545 Cole Court Newtonfurt, KY 13882
			  - 9038 Frye Ramp South Cheryltown, CT 54262
			  - 06210 David Court South Kimberly, IL 10236
			statistics:
			  count: 10
			  unique: 10
			  length_histogram:
				min_length: 31
				max_length: 52
				25%_length: 41
				50%_length: 42
				75%_length: 42
			  most_freq_len:
				- 42
				- 41
				- 31
				- 45
				- 52
			  most_freq_len_freq:
				- 0.3
				- 0.2
				- 0.1
				- 0.1
				- 0.1
			  least_frequent_len:
				- 31
				- 45
				- 52
				- 43
				- 37
			  least_frequent_len_freq:
				- 0.1
				- 0.1
				- 0.1
				- 0.1
				- 0.1
			  least_frequent_value:
				- PSC 4713, Box 9649 APO AA 43433
				- 2035 Simmons Islands Heatherchester, IN 46152
				- 275 Stone Ridges Suite 885 East Aliciafurt, MH 15407
				- 38432 Moreno Turnpike Garrettland, TN 72939
				- 25545 Cole Court Newtonfurt, KY 13882
		  - name: age
			type: int64
			concept: ""
			constraint: []
			confidential: false
			sample:
			  - 80
			  - 47
			  - 95
			  - 61
			  - 45
			statistics:
			  count: 10
			  unique: 9
			  mean: 57.7
			  value_histogram:
				min: 29.0
				25%: 45.5
				50%: 54.0
				75%: 71.0
				max: 95.0
		  - name: date
			type: datetime64[ns]
			concept: ""
			constraint: []
			confidential: false
			sample:
			  - "2005-05-10 00:00:00"
			  - "2010-11-18 00:00:00"
			  - "2003-10-11 00:00:00"
			  - "2014-07-24 00:00:00"
			  - "2022-04-23 00:00:00"
			statistics:
			  count: 10
			  unique: 10
			  date_histogram:
				earliest: "2001-08-23 00:00:00"
				latest: "2022-04-23 00:00:00"
		  - name: phone
			type: object
			concept: ""
			constraint: []
			confidential: false
			sample:
			  - (517)819-3454
			  - +1-407-997-8293x68130
			  - 001-845-854-2110
			  - "7795418893"
			  - 828-755-3826
			statistics:
			  count: 10
			  unique: 10
			  length_histogram:
				min_length: 10
				max_length: 21
				25%_length: 12
				50%_length: 16
				75%_length: 16
			  most_freq_len:
				- 16
				- 12
				- 13
				- 21
				- 10
			  most_freq_len_freq:
				- 0.4
				- 0.2
				- 0.1
				- 0.1
				- 0.1
			  least_frequent_len:
				- 12
				- 13
				- 21
				- 10
				- 18
			  least_frequent_len_freq:
				- 0.2
				- 0.1
				- 0.1
				- 0.1
				- 0.1
			  least_frequent_value:
				- 260-587-0590
				- (517)819-3454
				- +1-407-997-8293x68130
				- "7795418893"
				- (330)616-7639x7810
`
)

func TestAnalyse(t *testing.T) {
	t.Parallel()

	inputList := []string{jsonlNewFormatInputPath}
	outputPath := filepath.Join(outputDir, outputFileNameAnalyse)
	analyse.Analyse(inputList, outputPath)

	// Read output file
	file, err := os.Open(outputPath)
	if err != nil {
		t.Errorf("Error opening output file: %v", err)
	}
	defer file.Close()

	// Decode the YAML data into a model.Base object
	var baseData model.Base

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&baseData)

	if err != nil {
		panic(err)
	}

	// Print the base data
	fmt.Println(valast.String(baseData))
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
