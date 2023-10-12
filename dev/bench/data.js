window.BENCHMARK_DATA = {
  "lastUpdate": 1697140673225,
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
      },
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
          "id": "cdbbd75279fc5b7d3b11ddff6a0d909023cdf256",
          "message": "refactor: expose interfaces and a driver to client\n\n* refactor: renamed model to rimo\r\n\r\n* feat: defined interface\r\n\r\n* refactor: .gitignore\r\n\r\n* feat: added FileWriter interface\r\n\r\n* feat: fileWriter interface with test\r\n\r\n* feat: loader for JSONL\r\n\r\n* test: rimo interface test\r\n\r\n* feat: added filesReader interface\r\n\r\n* test: FilesReader with 2 files\r\n\r\n* refactor: re adding to model package to avoid circular dependency\r\n\r\n* feat(rimo): driven_test.go\r\n\r\n* refactor(rimo): improve Writer name\r\n\r\n* refactor: renamed testWriter() and added GetBase() method\r\n\r\n* feat: TestWriter improv (similar to prev commit)\r\n\r\n* refactor: minor typo\r\n\r\n* test: RIMO pipeline infra_test.go\r\n\r\n* refactor: added cobra command using interface\r\n\r\n* refactor: more explicit variable naming\r\n\r\n* refactor: removed unusued function\r\n\r\n* refactor: added test to compare pipeline with expected output\r\n\r\n* refactor: fix : giving filesReader proper filepath\r\n\r\n* refactor: almost work as expected\r\n\r\n* refactor: updated schema from rimo pkg to model pkg\r\n\r\n* refactor: work as expected\r\n\r\n* fix: remove old analyse command\r\n\r\n* chore: remove dead code\r\n\r\n* docs: add GPLv3 license header in new files\r\n\r\n* fix: remove output test from git\r\n\r\n---------\r\n\r\nCo-authored-by: Youen Péron <youen.peron@cgi.com>",
          "timestamp": "2023-09-27T22:43:03+02:00",
          "tree_id": "65d970da16ca6e2778c0d418c670283c4585c1e6",
          "url": "https://github.com/CGI-FR/RIMO/commit/cdbbd75279fc5b7d3b11ddff6a0d909023cdf256"
        },
        "date": 1695848462800,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - ns/op",
            "value": 6906414531,
            "unit": "ns/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - lines/s",
            "value": 14479,
            "unit": "lines/s",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - B/op",
            "value": 1358506976,
            "unit": "B/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - allocs/op",
            "value": 13663849,
            "unit": "allocs/op",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "89643755+giraud10@users.noreply.github.com",
            "name": "giraud10",
            "username": "giraud10"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "b654755e714c2cf537a5dc20b8cd69f3f67d8b7c",
          "message": "docs: add rimo schema (#30)",
          "timestamp": "2023-10-12T21:41:28+02:00",
          "tree_id": "2ac92db707ef21b94d9e73ac76ba0123e9376d9d",
          "url": "https://github.com/CGI-FR/RIMO/commit/b654755e714c2cf537a5dc20b8cd69f3f67d8b7c"
        },
        "date": 1697140672815,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - ns/op",
            "value": 6195392178,
            "unit": "ns/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - lines/s",
            "value": 16141,
            "unit": "lines/s",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - B/op",
            "value": 1358497248,
            "unit": "B/op",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "BenchmarkAnalyseInterface/numLines=100000 - allocs/op",
            "value": 13663645,
            "unit": "allocs/op",
            "extra": "1 times\n2 procs"
          }
        ]
      }
    ]
  }
}