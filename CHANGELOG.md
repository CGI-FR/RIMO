# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## Unreleased (detailled)

- `Added` rimo.yaml format 
- `Changed` loading .jsonl with Numbers json type
- `Changed` use empty slices of type [json.Numbers, string, bool] to load jsonl
- `Changed` deal with null cases at first row of .jsonl

## TODO / Notes

- Venom test
- neon test / compile / lint / test-int

## [0.1.0]

- `Added` analyse.go : load .jsonl, process, output rimo.yaml
