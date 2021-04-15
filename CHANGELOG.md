# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/query-meta/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.1.1] - 2021-04-15

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.16.3

### Changed

- dependencies
  - built using `Go 1.16.3`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `denisenkom/go-mssqldb`
    - `v0.0.0-20201104001602-427686ac8ec1` to `v0.10.0`
  - `golang.org/x/text`
    - `v0.3.5` to `v0.3.6`
  - `pelletier/go-toml`
    - `v1.8.1` to `v1.9.0`
  - `rs/zerolog`
    - `v1.20.0` to `v1.21.0`
  - `actions/setup-node`
    - `v2.1.4` to `v2.1.5`

### Fixed

- `denisenkom/go-mssqldb` library is incompatible with Go 1.16
- Fix environment variable names
- Minor README issues
- PURPOSE statement in `doc.go` file

## [v0.1.0] - 2021-03-24

### Added

Initial release!

This release provides an early version of a prototype tool intended to replace
the existing `query_meta.php` script used to retrieve patron records from Meta
and write out a pipe-delimited CSV file for further processing.

Due to known issues with the `denisenkom/go-mssqldb` package, Go 1.16 is not
supported at this time. Go 1.15 should be used instead until upstream
[GH-639](https://github.com/denisenkom/go-mssqldb/issues/639) is resolved.

[Unreleased]: https://github.com/atc0005/query-meta/compare/v0.1.1...HEAD
[v0.1.1]: https://github.com/atc0005/query-meta/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/query-meta/releases/tag/v0.1.0
