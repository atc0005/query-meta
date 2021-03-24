// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"database/sql"
	"fmt"
)

// rowsCount returns the number of rows for a specified table.
func rowsCount(db *sql.DB, table string, query string) (int, error) {
	var rowsCount int
	if err := db.QueryRow(query).Scan(&rowsCount); err != nil {
		return -1, fmt.Errorf(
			"failed to retrieve row count for table %s: %w",
			table,
			err,
		)
	}

	return rowsCount, nil
}
