// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/atc0005/query-meta/internal/config"
)

// Setup basic tests to ensure that alexflint/go-arg package recognizes flags
// (see atc0005/query-meta#15 and alexflint/go-arg#157 for backstory) and that
// config file loading works as expected.

func TestConfigFileParsing(t *testing.T) {

	// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang

	// Save old command-line arguments so that we can restore them later
	oldArgs := os.Args

	// Defer restoring original command-line arguments
	defer func() { os.Args = oldArgs }()

	exampleCfgFile := filepath.Join("../../", "contrib", "qm", "config.example.toml")

	// Note to self: Don't add/escape double-quotes here. The shell strips
	// them away and the application never sees them.
	os.Args = []string{
		"/usr/local/bin/qm", "--config-file", exampleCfgFile,
	}

	_, err := config.New()
	if err != nil {
		t.Errorf("Error encountered when instantiating configuration: %s", err)
	} else {
		t.Log("No errors encountered when instantiating configuration")
	}

}

func TestConfigFileParsingResults(t *testing.T) {

	// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang

	// Save old command-line arguments so that we can restore them later
	oldArgs := os.Args

	// Defer restoring original command-line arguments
	defer func() { os.Args = oldArgs }()

	exampleCfgFile := filepath.Join("../../", "contrib", "qm", "config.example.toml")

	// Note to self: Don't add/escape double-quotes here. The shell strips
	// them away and the application never sees them.
	os.Args = []string{
		"/usr/local/bin/qm", "--config-file", exampleCfgFile,
	}

	cfg, err := config.New()
	if err != nil {
		t.Errorf("Error encountered when instantiating configuration: %s", err)
	} else {
		t.Log("No errors encountered when instantiating configuration")
	}

	queryActivePatronRecords := `
SELECT
    Gid_status,
    Gid,
    Idnum,
    spriden_id,
    ssn,
    givenName,
    name_middle,
    sn,
    Libissuenum,
    empstu_flag,
    student_status,
    banner_last_term_enrolled,
    Student_currently_enrolled,
    Student_level_code,
    employeestatus,
    employeetype,
    sal_table,
    emp_group_code,
    tes_flag,
    campus_code,
    Employee_term_date,
    ou,
    title,
    l,
    address_dept_city,
    address_dept_state,
    address_dept_zip,
    telephoneNumber,
    Address_student_perm_street street1,
    Address_student_perm_street2 street2,
    Address_student_perm_street3 street3,
    Address_student_perm_city,
    Address_student_perm_state,
    Address_student_perm_zip,
    Address_nation_code,
    Phone_student_perm,
    student_type,
    Other_status,
    Other_type_code,
    Student_degree_earned,
    NearFieldBadgeID
FROM
    library_patrons
WHERE
    Gid_status != 'I';
`

	testConfigFileValues := struct {
		host                         string
		port                         int
		instance                     string
		username                     string
		password                     string
		encryptMode                  string
		trustCert                    bool
		databaseName                 string
		databaseTableName            string
		queryRowCountAllRecords      string
		queryRowCountInactiveRecords string
		queryactivePatronRecords     string
		logLevel                     string
	}{
		host:                         "mssql52.example.com",
		port:                         1433,
		instance:                     "mssql07",
		username:                     "chocotaco",
		password:                     "dbPasW0rdZ",
		encryptMode:                  "true",
		trustCert:                    false,
		databaseName:                 "Meta",
		databaseTableName:            "library_patrons",
		queryRowCountAllRecords:      "SELECT COUNT(*) FROM library_patrons",
		queryRowCountInactiveRecords: "SELECT COUNT(*) FROM library_patrons WHERE Gid_status != 'I'",
		logLevel:                     "info",

		// This may require massaging to get the expected equivalency
		queryactivePatronRecords: queryActivePatronRecords,
	}

	if cfg.DBServerHost() != testConfigFileValues.host {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerHost(),
			testConfigFileValues.host,
			"host",
		)
	}

	if cfg.DBServerPort() != testConfigFileValues.port {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerPort(),
			testConfigFileValues.port,
			"port",
		)
	}

	if cfg.DBServerInstance() != testConfigFileValues.instance {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerInstance(),
			testConfigFileValues.instance,
			"instance",
		)
	}

	if cfg.DBServerUsername() != testConfigFileValues.username {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerUsername(),
			testConfigFileValues.username,
			"username",
		)
	}

	if cfg.DBServerPassword() != testConfigFileValues.password {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerPassword(),
			testConfigFileValues.password,
			"password",
		)
	}

	if cfg.DBServerEncryptMode() != testConfigFileValues.encryptMode {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerEncryptMode(),
			testConfigFileValues.encryptMode,
			"encrypt_mode",
		)
	}

	if cfg.DBServerTrustCert() != testConfigFileValues.trustCert {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBServerTrustCert(),
			testConfigFileValues.trustCert,
			"trust_cert",
		)
	}

	if cfg.DBName() != testConfigFileValues.databaseName {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBName(),
			testConfigFileValues.databaseName,
			"database_name",
		)
	}

	if cfg.DBTable() != testConfigFileValues.databaseTableName {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBTable(),
			testConfigFileValues.databaseTableName,
			"database_table_name",
		)
	}

	if cfg.DBQueryCountAllRecords() != testConfigFileValues.queryRowCountAllRecords {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBQueryCountAllRecords(),
			testConfigFileValues.queryRowCountAllRecords,
			"query_row_count_all_records",
		)
	}

	if cfg.DBQueryCountInactiveRecords() != testConfigFileValues.queryRowCountInactiveRecords {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBQueryCountInactiveRecords(),
			testConfigFileValues.queryRowCountInactiveRecords,
			"query_row_count_inactive_records",
		)
	}

	// multi-line SQL query; trim spaces to ensure that comparisons work as
	// expected
	if strings.TrimSpace(cfg.DBQueryRetrieveActivePatronRecords()) !=
		strings.TrimSpace(testConfigFileValues.queryactivePatronRecords) {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.DBQueryRetrieveActivePatronRecords(),
			testConfigFileValues.queryactivePatronRecords,
			"query_active_patron_records",
		)
	}

	if cfg.LogLevel() != testConfigFileValues.logLevel {
		t.Errorf(
			"got %v; want %v for config file setting %q",
			cfg.LogLevel(),
			testConfigFileValues.logLevel,
			"log_level",
		)
	}

}
