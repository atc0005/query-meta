// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"strings"

	"github.com/atc0005/query-meta/internal/config"
	"github.com/atc0005/query-meta/internal/meta"
)

// processRecord receives a patron record and performs any required
// manipulation to sanitize or cleanup data.
func processRecord(record *meta.LibraryPatronRecord, cfg *config.Config) {

	// Replace any commas in surnames to help prevent CSV parsing issues
	// by other tooling.
	if strings.Contains(record.Surname.String, ",") {
		modifiedSurname := strings.ReplaceAll(record.Surname.String, ",", "")
		cfg.Log.Warn().
			Str("action", "modify").
			Str("patron_id", record.GID.String).
			Str("original_surname", record.Surname.String).
			Str("modified_surname", modifiedSurname).
			Msg("surname contains comma")

		record.Surname.String = modifiedSurname
	}

	// search all fields in formatted record
	if strings.Contains(record.CSV(), `"`) {

		// Skipping the record instead of trying to manipulate it (as a
		// workaround) so that the patron record could (potentially) be parsed
		// later once data quality issues are resolved.
		cfg.Log.Error().
			Str("action", "skipping").
			Str("patron_id", record.GID.String).
			Msg("record has forbidden characters")

	}

}
