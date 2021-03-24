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

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/rs/zerolog/log"

	"github.com/atc0005/query-meta/internal/config"
	"github.com/atc0005/query-meta/internal/meta"
)

func main() {

	cfg, configErr := config.New()
	if configErr != nil {
		log.Err(configErr).Msg("error validating configuration")

		// no need to go any further, we *want* to exit right away; we don't
		// have a working configuration and there isn't anything further to do
		return
	}

	// TODO: Decide whether to use dedicated files, or just stdout and stderr
	// as the prior query_meta.php script was using. For now, I'll mirror the
	// behavior of the previous tool.
	//
	// Setup output log file for capturing normal runtime info
	// Setup output file for writing CSV records (pipe delimited)
	// Setup error log file for capturing error messages

	// Use ADO connection string format. This format does not appear to
	// require character encoding in order to handle special characters (e.g., encoding `{` as `%7B`)
	// https://github.com/denisenkom/go-mssqldb#common-parameters
	// https://github.com/denisenkom/go-mssqldb#less-common-parameters
	dsn := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=%s;TrustServerCertificate=%t;app name=%s;",
		cfg.DBServerHost(),
		cfg.DBServerUsername(),
		cfg.DBServerPassword(),
		cfg.DBServerPort(),
		cfg.DBName(),
		cfg.DBServerEncryptMode(),
		cfg.DBServerTrustCert(),
		config.Version(),
	)

	if cfg.DBServerInstance() != "" {
		dsn += fmt.Sprintf("instance=%s;", cfg.DBServerInstance())
	}

	// Setup connection (which is established lazily)
	db, dbOpenErr := sql.Open("mssql", dsn)
	if dbOpenErr != nil {
		cfg.Log.Error().Err(dbOpenErr).Msg("open connection failed")

		// no need to go any further, we *want* to exit right away; we don't
		// have a connection to the remote server and there isn't anything
		// further we can do
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			cfg.Log.Error().Err(err).Msg("error closing connection to server")
		} else {
			cfg.Log.Debug().Msg("successfully closed connection to server")
		}
	}()

	// Use Ping() to create a connection and check for any errors
	if err := db.Ping(); err != nil {
		cfg.Log.Error().Err(err).Msg("error verifying connection to server")

		return
	}

	cfg.Log.Debug().
		Str("host", cfg.DBServerHost()).
		Str("instance", cfg.DBServerInstance()).
		Int("port", cfg.DBServerPort()).
		Str("database", cfg.DBName()).
		Str("username", cfg.DBServerHost()).
		Str("encrypt_mode", cfg.DBServerEncryptMode()).
		Bool("validate_cert", !cfg.DBServerTrustCert()).
		Msg("connection established")

	// ensure *any* rows are present in the source database table
	allPatronRecordsCount, allRecordsQueryErr := rowsCount(
		db,
		cfg.DBTable(),
		cfg.DBQueryCountAllRecords(),
	)
	if allRecordsQueryErr != nil {
		cfg.Log.Error().
			Err(allRecordsQueryErr).
			Str("table_name", cfg.DBTable()).
			Str("query_template", cfg.DBQueryCountAllRecords()).
			Msg("failed to execute query")

		return
	}

	// ensure there are rows in the source database table that match our
	// criteria
	activePatronRecordsCount, activeRecordsQueryErr := rowsCount(
		db,
		cfg.DBTable(),
		cfg.DBQueryCountInactiveRecords(),
	)
	if activeRecordsQueryErr != nil {
		cfg.Log.Error().
			Err(allRecordsQueryErr).
			Str("table_name", cfg.DBTable()).
			Str("query_template", cfg.DBQueryCountInactiveRecords()).
			Msg("failed to execute query")

		return
	}

	// Fatal conditions
	switch {
	case allPatronRecordsCount == 0:
		cfg.Log.Error().
			Err(allRecordsQueryErr).
			Str("table_name", cfg.DBTable()).
			Msg("no matching rows found in table")

		return

	case activePatronRecordsCount == 0:
		cfg.Log.Error().
			Err(activeRecordsQueryErr).
			Str("table_name", cfg.DBTable()).
			Msg("no matching rows found in table")

		return

	default:
		cfg.Log.Debug().
			Str("table_name", cfg.DBTable()).
			Int("all_records_num", allPatronRecordsCount).
			Int("active_records_num", activePatronRecordsCount).
			Int("inactive_records_num", allPatronRecordsCount-activePatronRecordsCount).
			Msg("matching rows found in table")

	}

	rows, queryErr := db.Query(cfg.DBQueryRetrieveActivePatronRecords())
	if queryErr != nil {
		cfg.Log.Error().Err(queryErr).Msg("failed to execute query")

		return
	}

	defer func() {
		if err := rows.Close(); err != nil {
			cfg.Log.Error().Err(err).Msg("error closing rows object")
		} else {
			cfg.Log.Debug().Msg("successfully closed rows object")
		}
	}()

	// patronRecords := make([]config.LibraryPatronRecord, 0, activePatronRecordsCount)
	var patronRecord meta.LibraryPatronRecord
	var patronRecordNum int

	for rows.Next() {

		scanErr := rows.Scan(
			&patronRecord.GIDStatus,
			&patronRecord.GID,
			&patronRecord.IDNum,
			&patronRecord.SprIdenID,
			&patronRecord.SSN,
			&patronRecord.GivenName,
			&patronRecord.NameMiddle,
			&patronRecord.Surname,
			&patronRecord.LibIssueNumber,
			&patronRecord.EmpStudentFlag,
			&patronRecord.StudentStatus,
			&patronRecord.BannerLastTermEnrolled,
			&patronRecord.StudentCurrentlyEnrolled,
			&patronRecord.StudentLevelCode,
			&patronRecord.EmployeeStatus,
			&patronRecord.EmployeeType,
			&patronRecord.SalaryTable,
			&patronRecord.EmployeeGroupCode,
			&patronRecord.TESFlag,
			&patronRecord.CampusCode,
			&patronRecord.EmployeeTerminationDate,
			&patronRecord.OU,
			&patronRecord.Title,
			&patronRecord.Location,
			&patronRecord.AddressDeptCity,
			&patronRecord.AddressDeptState,
			&patronRecord.AddressDeptZip,
			&patronRecord.TelephoneNumber,
			&patronRecord.AddressStudentPermStreet1,
			&patronRecord.AddressStudentPermStreet2,
			&patronRecord.AddressStudentPermStreet3,
			&patronRecord.AddressStudentPermCity,
			&patronRecord.AddressStudentPermState,
			&patronRecord.AddressStudentPermZip,
			&patronRecord.AddressNationCode,
			&patronRecord.PhoneStudentPerm,
			&patronRecord.StudentType,
			&patronRecord.OtherStatus,
			&patronRecord.OtherTypeCode,
			&patronRecord.StudentDegreeEarned,
		)
		if scanErr != nil {
			cfg.Log.Error().Err(scanErr).Msg("failed to retrieve patron record")

			return
		}
		patronRecordNum++
		cfg.Log.Debug().
			Int("record_num", patronRecordNum).
			Msg("completed scanning record")

		// perform any necessary cleanup of patron data
		processRecord(&patronRecord, cfg)

		// Explicitly encode formatted patron details in a compatible ingest
		// format for parsing by other tools in the pipeline.
		encodedRecord, inputBytesUsed, err := patronRecord.Extract()
		if err != nil {
			cfg.Log.Error().
				Err(err).
				Str("patron", patronRecord.GID.String).
				Msg("failed to encode patron record")

			return
		}
		cfg.Log.Debug().
			Str("patron", patronRecord.GID.String).
			Int("bytes_used_from_input", inputBytesUsed).
			Msg("successfully encoded record")
		fmt.Print(encodedRecord)

	}
	cfg.Log.Debug().Msg("completed scanning all row objects")

	if err := rows.Err(); err != nil {
		cfg.Log.Error().Err(err).Msg("error occurred during rows scan loop")

		return
	}
	cfg.Log.Debug().Msg("no row object scanning errors encountered")

	cfg.Log.Debug().
		Int("patron_records", patronRecordNum).
		Msg("all records fetched")

}
