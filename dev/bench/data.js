window.BENCHMARK_DATA = {
  "lastUpdate": 1693232609229,
  "repoUrl": "https://github.com/CGI-FR/RIMO",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "116900975+mathisdrn@users.noreply.github.com",
            "name": "mathisdrn",
            "username": "mathisdrn"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "dc04ea24b63a196618743a007c41cd6ba9f1469a",
          "message": "Add benchmark (#28)\n\n* fix: typo readme.md\r\n\r\n* refactor: sortBase() as a method of model.Base\r\n\r\n* feat: added benchmark\r\n\r\n* fix: rm venom output",
          "timestamp": "2023-08-28T16:20:26+02:00",
          "tree_id": "15fc859141170ba81d5f5e696ec4981459faec4a",
          "url": "https://github.com/CGI-FR/RIMO/commit/dc04ea24b63a196618743a007c41cd6ba9f1469a"
        },
        "date": 1693232608687,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAnalyse/numLines=100 - ns/op",
            "value": 5180797,
            "unit": "ns/op",
            "extra": "229 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100 - lines/s",
            "value": 19301,
            "unit": "lines/s",
            "extra": "229 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100 - B/op",
            "value": 2279321,
            "unit": "B/op",
            "extra": "229 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100 - allocs/op",
            "value": 17623,
            "unit": "allocs/op",
            "extra": "229 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=1000 - ns/op",
            "value": 41512502,
            "unit": "ns/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=1000 - lines/s",
            "value": 24088,
            "unit": "lines/s",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=1000 - B/op",
            "value": 15362881,
            "unit": "B/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=1000 - allocs/op",
            "value": 140632,
            "unit": "allocs/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=10000 - ns/op",
            "value": 419291194,
            "unit": "ns/op",
            "extra": "3 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=10000 - lines/s",
            "value": 23849,
            "unit": "lines/s",
            "extra": "3 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=10000 - B/op",
            "value": 136304285,
            "unit": "B/op",
            "extra": "3 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=10000 - allocs/op",
            "value": 1361524,
            "unit": "allocs/op",
            "extra": "3 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100000 - ns/op",
            "value": 4692771309,
            "unit": "ns/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100000 - lines/s",
            "value": 21309,
            "unit": "lines/s",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100000 - B/op",
            "value": 1358328536,
            "unit": "B/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyse/numLines=100000 - allocs/op",
            "value": 13662987,
            "unit": "allocs/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=100 - ns/op",
            "value": 176535,
            "unit": "ns/op",
            "extra": "6799 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=100 - lines/s",
            "value": 566461,
            "unit": "lines/s",
            "extra": "6799 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=100 - B/op",
            "value": 81373,
            "unit": "B/op",
            "extra": "6799 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=100 - allocs/op",
            "value": 1140,
            "unit": "allocs/op",
            "extra": "6799 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=1000 - ns/op",
            "value": 1714174,
            "unit": "ns/op",
            "extra": "690 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=1000 - lines/s",
            "value": 583372,
            "unit": "lines/s",
            "extra": "690 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=1000 - B/op",
            "value": 855614,
            "unit": "B/op",
            "extra": "690 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=1000 - allocs/op",
            "value": 11096,
            "unit": "allocs/op",
            "extra": "690 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=10000 - ns/op",
            "value": 17260782,
            "unit": "ns/op",
            "extra": "67 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=10000 - lines/s",
            "value": 579349,
            "unit": "lines/s",
            "extra": "67 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=10000 - B/op",
            "value": 8196446,
            "unit": "B/op",
            "extra": "67 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_numeric,_numValues=10000 - allocs/op",
            "value": 110374,
            "unit": "allocs/op",
            "extra": "67 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=100 - ns/op",
            "value": 181943,
            "unit": "ns/op",
            "extra": "5845 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=100 - lines/s",
            "value": 549623,
            "unit": "lines/s",
            "extra": "5845 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=100 - B/op",
            "value": 93175,
            "unit": "B/op",
            "extra": "5845 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=100 - allocs/op",
            "value": 1164,
            "unit": "allocs/op",
            "extra": "5845 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=1000 - ns/op",
            "value": 1775744,
            "unit": "ns/op",
            "extra": "670 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=1000 - lines/s",
            "value": 563145,
            "unit": "lines/s",
            "extra": "670 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=1000 - B/op",
            "value": 1019863,
            "unit": "B/op",
            "extra": "670 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=1000 - allocs/op",
            "value": 11150,
            "unit": "allocs/op",
            "extra": "670 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=10000 - ns/op",
            "value": 19259840,
            "unit": "ns/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=10000 - lines/s",
            "value": 519216,
            "unit": "lines/s",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=10000 - B/op",
            "value": 9769628,
            "unit": "B/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_text,_numValues=10000 - allocs/op",
            "value": 110567,
            "unit": "allocs/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=100 - ns/op",
            "value": 135668,
            "unit": "ns/op",
            "extra": "8767 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=100 - lines/s",
            "value": 737096,
            "unit": "lines/s",
            "extra": "8767 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=100 - B/op",
            "value": 65128,
            "unit": "B/op",
            "extra": "8767 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=100 - allocs/op",
            "value": 921,
            "unit": "allocs/op",
            "extra": "8767 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=1000 - ns/op",
            "value": 1227709,
            "unit": "ns/op",
            "extra": "963 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=1000 - lines/s",
            "value": 814526,
            "unit": "lines/s",
            "extra": "963 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=1000 - B/op",
            "value": 615465,
            "unit": "B/op",
            "extra": "963 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=1000 - allocs/op",
            "value": 9025,
            "unit": "allocs/op",
            "extra": "963 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=10000 - ns/op",
            "value": 12278479,
            "unit": "ns/op",
            "extra": "91 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=10000 - lines/s",
            "value": 814434,
            "unit": "lines/s",
            "extra": "91 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=10000 - B/op",
            "value": 6290355,
            "unit": "B/op",
            "extra": "91 times\n2 procs"
          },
          {
            "name": "BenchmarkMetric/type=_bool,_numValues=10000 - allocs/op",
            "value": 90032,
            "unit": "allocs/op",
            "extra": "91 times\n2 procs"
          }
        ]
      }
    ]
  }
}