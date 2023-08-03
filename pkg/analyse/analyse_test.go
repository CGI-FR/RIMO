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
	yamlData                = `
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

func TestPipeline(t *testing.T) {
	t.Parallel()

	data := analyse.Load(jsonlNewFormatInputPath, "new")
	fmt.Println(valast.String(data))

	// dataMap := analyse.ColType(data)
	// fmt.Println(valast.String(dataMap))
}
