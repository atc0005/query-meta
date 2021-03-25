<!-- omit in toc -->
# query-meta

Go-based tooling to process patron records from the Meta Microsoft SQL Server
database.

[![Latest Release](https://img.shields.io/github/release/atc0005/query-meta.svg?style=flat-square)](https://github.com/atc0005/query-meta/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/query-meta.svg)](https://pkg.go.dev/github.com/atc0005/query-meta)
[![Validate Codebase](https://github.com/atc0005/query-meta/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/query-meta/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/query-meta/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/query-meta/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/query-meta/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/query-meta/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/query-meta/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/query-meta/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Limitations](#limitations)
  - [Builds](#builds)
  - [Output files](#output-files)
- [Features](#features)
- [Status](#status)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Command-line Arguments](#command-line-arguments)
  - [Environment Variables](#environment-variables)
  - [Configuration File](#configuration-file)
- [Examples](#examples)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

Go-based tooling to process patron records from the Meta Microsoft SQL
Server database. The initial `qm` binary provided by this project is intended
to replace the existing `query_meta.php` script used to retrieve patron
records from Meta and write out a pipe-delimited CSV file for further
processing.

## Limitations

### Builds

Due to known issues with the `denisenkom/go-mssqldb` package, Go 1.16 is not
supported at this time. Go 1.15 should be used instead until upstream
[GH-639](https://github.com/denisenkom/go-mssqldb/issues/639) is resolved.

### Output files

The `qm` CLI app writes patron records to `stdout` and log messages to
`stderr` in an effort to mimic the behavior of the original PHP script. Future
versions of this tool will likely write data directly to files specified by
configuration file.

## Features

The `qm` CLI app queries a Microsoft SQL Server (MSSQL) database and writes
out a pipe-delimited CSV file for further processing by other tools in a
patron records processing pipeline.

## Status

Beta quality. The shape of the application is still under development, but
functional parity with the original PHP script has been achieved.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application.

## Requirements

The following is a loose guideline. Other combinations of Go (see below) and
operating systems for building and running tools from this repo may work, but
have not been tested.

### Building source code

- Go 1.15
  - Due to known issues with the `denisenkom/go-mssqldb` package, Go 1.16 is
    not supported at this time. Go 1.15 should be used instead for building
    applications from this project's repo.
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10 Version 1909+
- Ubuntu Linux 18.04

## Installation

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/query-meta`
   1. `cd query-meta`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     - `sudo yum install make gcc`
   - for Windows
     - Emulated environments (*easier*)
       - Skip all of this and build using the default `go build` command in
         Windows (see below for use of the `-mod=vendor` flag)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed
       - If already familiar with LXD, create a container and follow the
         installation steps given previously to install required dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - for the current operating system, explicitly using bundled dependencies
         in top-level `vendor` folder
     - `go build -mod=vendor ./cmd/qm/`
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for use on Windows
      - `make windows`
   - for use on Linux
     - `make linux`
1. Copy the newly compiled binary from the applicable `/tmp` subdirectory path
   (based on the clone instructions in this section) below and deploy where
   needed.
   - if using `Makefile`
     - look in `/tmp/query-meta/release_assets/qm/`
   - if using `go build`
     - look in `/tmp/query-meta/`

## Configuration

This application is configured via a TOML-formatted configuration file. The
only CLI flags currently supported are `help` and and a flag to specify the
location of the configuration file.

### Command-line Arguments

| Option        | Required | Default        | Repeat | Possible               | Description                                                                                                                    |
| ------------- | -------- | -------------- | ------ | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `h`, `help`   | No       | `false`        | No     | `h`, `help`            | Show Help text along with the list of supported flags.                                                                         |
| `config-file` | Yes      | *empty string* | No     | *valid path to a file* | Fully-qualified path to required TOML-formatted configuration file. See contrib/qm/config.example.toml for a starter template. |

### Environment Variables

| Flag Name     | Environment Variable Name | Notes | Example (mostly using default values)                            |
| ------------- | ------------------------- | ----- | ---------------------------------------------------------------- |
| `config-file` | `QUERY_META_CONFIG_FILE`  |       | `QUERY_META_CONFIG_FILE="/usr/local/etc/query-meta/config.toml"` |

### Configuration File

Configuration file settings have the lowest priority and are overridden by
settings specified in other configuration sources, except for default values.
See the [Command-line Arguments](#command-line-arguments) table for more
information, including the available values for the listed configuration
settings.

| Config file Setting Name           | Required | Default        | Possible                                                                | Description                                                                                                                                                                                |
| ---------------------------------- | -------- | -------------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `host`                             | **Yes**  |                | *valid fqdn or IP Address*                                              | The database server hosting the database used by this application. If using encryption, this value should match one of the Subject Alternate Name (SANs) values listed on the certificate. |
| `port`                             | No       | 1433           | *valid TCP port*                                                        | The TCP port used to connect to the database server. If not specified, the default port will be used.                                                                                      |
| `instance`                         | No       | *empty string* | *valid instance name*                                                   | The the database server instance name. This may be blank.                                                                                                                                  |
| `username`                         | **Yes**  | *empty string* | *valid username*                                                        | The username used to connect to the database server. An account with read-only access to the database table is sufficient.                                                                 |
| `password`                         | **Yes**  | *empty string* | *valid password*                                                        | The plaintext password used to connect to the database server. An account with read-only access to the database table is sufficient.                                                       |
| `encrypt_mode`                     | No       | `false`        | `true`, `false`, `disable`                                              | Specifies whether data sent between client and server is encrypted. true for yes, false for login packet only and disable for no encryption.                                               |
| `trust_cert`                       | No       | `false`        | `true`, `false`                                                         | Indicates whether the certificate should be trusted as-is without validation. WARNING: TLS is susceptible to man-in-the-middle attacks if enabling this option.                            |
| `log_level`                        | No       | `info`         | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | The maximum log level at which messages will be logged. Log messages below this threshold will be discarded.                                                                               |
| `database_name`                    | **Yes**  | *empty string* | *valid database name*                                                   | The name of the database which holds patron records.                                                                                                                                       |
| `database_table_name`              | **Yes**  | *empty string* | *valid database table or view name*                                     | The name of the database table or view. Used primarily in logging output but also for a few other use cases.                                                                               |
| `query_row_count_all_records`      | **Yes**  | *empty string* | *valid SQL query*                                                       | The query used to obtain a count of all patron records.                                                                                                                                    |
| `query_row_count_inactive_records` | **Yes**  | *empty string* | *valid SQL query*                                                       | The query used to obtain a count of all inactive patron records                                                                                                                            |
| `query_active_patron_records`      | **Yes**  | *empty string* | *valid SQL query*                                                       | The query used to obtain all active patron records (by intentionally excluding inactive records).                                                                                          |

The [`config.example.toml`](config.example.toml) file is intended as a
starting point for your own configuration file and attempts to illustrate
working values for the available command-line flags.

Once reviewed and potentially adjusted, your copy of the configuration file
can be placed wherever you like and can be referred to using the `config-file`
flag. Though you can name the file whatever you like, it is recommended to
keep the `toml` file extension to help signal to future sysadmins that they're
working with a TOML formatted file and not a traditional INI file.

Suggested filenames:

- `config.toml`
- `query-meta.toml`

See the [Examples](#examples) and [Command-line
arguments](#command-line-arguments) sections for usage details.

## Examples

```ShellSession
cd /path/to/workspace
/usr/local/bin/qm --config-file query-meta.toml > view_extract.txt 2> log/query_exceptions.txt
```

## License

```license
MIT License

Copyright 2021 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## References

- <https://github.com/denisenkom/go-mssqldb>
  - (GH-639) [Not implemented panic in go 1.16](https://github.com/denisenkom/go-mssqldb/issues/639)
- <https://github.com/rs/zerolog>
- <https://github.com/alexflint/go-arg>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/query-meta>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
