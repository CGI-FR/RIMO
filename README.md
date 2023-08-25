# rimo

## Description

Rimo contains tools that helps creating a *masking.yaml* for [PIMO](https://github.com/CGI-FR/PIMO).

## Usage

### `rimo analyse`

```console
rimo analyse [inputDir] [outputDir]
```

- `inputDir` : path to a directory containing *jsonl* files.
- `output` : path to a directory where *rimo.yaml* will be created.

**inputDir** must contain .jsonl files named basename_tablename.jsonl and respecting this format :

```json
{"colName1": "value1", "colName2": "value2" } 
{"colName1": "value2", "colName2": "value2" }
```

such files can be generated using [LINO](https://github.com/CGI-FR/LINO)

**outputDir** will generate basename.yaml in output directory containing various metrics. An example can be found in *testdata/data1/data_expected.yaml*.

## Tests

Run `neon test-int` to execute unit-test and Venom test.

## Project status

In active development

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Copyright (C) 2023 CGI France

This file is part of RIMO.

RIMO is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

RIMO is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with RIMO.  If not, see <http://www.gnu.org/licenses/>.
