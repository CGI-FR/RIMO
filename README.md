# rimo

## Description

Rimo contains a series of tools that helps to create a *masking.yaml* for [PIMO](https://github.com/CGI-FR/PIMO).
<!-- It works as a 6 steps process : -->
<!-- ![rimo steps](.github/img/rimo_steps.png "rimo steps") -->

<!-- 1. `analyse` : extract meaningful information on database from *.jsonl* -->

<!-- ## Installation
`rimo` command line work in relative project's directory, like `git` or `docker` -->

## Usage

### `rimo analyse`

```console
rimo analyse input output
```

- `input` : path to a directory containing to *jsonl* files
- `output` : path to a directory where output the *.yaml* file

**input.jsonl** is a JSON single line that contains a pair of (column_name, value) for every row of the database table

**output.yaml** contain various metrics on table's columns and a small default configuration for PIMO. An example can be found in *src/unit_test/testcase_output.yaml*.

### `rimo jsonschema`

Generate the json schema of rimo.

## Tests

To run tests execute `neon test-int`.

## Project status

In active development

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
`testcase_data.jsonl` in *tests/data* is generated using Faker and does contain any real information.

## License

Copyright (C) 2023 CGI France

PIMO is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

PIMO is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
 along with PIMO.  If not, see <http://www.gnu.org/licenses/>.
