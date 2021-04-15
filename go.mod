// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

module github.com/atc0005/query-meta

// NOTE: Go 1.15 is needed for the time being until GH-639 is fixed.

go 1.15

require (
	github.com/alexflint/go-arg v1.3.0
	github.com/denisenkom/go-mssqldb v0.0.0-20201104001602-427686ac8ec1
	github.com/pelletier/go-toml v1.9.0
	github.com/rs/zerolog v1.21.0
	golang.org/x/text v0.3.6
)
