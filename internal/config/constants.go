// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

const (

	// MyAppName is the public name of this application.
	myAppName string = "query-meta"

	// MyAppURL is the location of the repo for this application.
	myAppURL string = "https://github.com/atc0005/query-meta"

	// MyAppDescription is the description for this application shown in
	// HelpText output.
	myAppDescription string = "Generate patron records CSV file from the Meta Microsoft SQL Server database"
)

// Default (flag, config file, etc) settings if not overridden by user input.
const (
	defaultLogLevel                           string = "info"
	defaultDBServerHost                       string = ""
	defaultDBServerInstance                   string = ""
	defaultDBServerPort                       int    = 1433
	defaultDBServerUsername                   string = ""
	defaultDBServerPassword                   string = ""
	defaultDBServerTrustCert                  bool   = false
	defaultDBName                             string = ""
	defaultDBTableName                        string = ""
	defaultDBQueryCountAllRecords             string = ""
	defaultDBQueryCountInactiveRecords        string = ""
	defaultDBQueryRetrieveActivePatronRecords string = ""
	// defaultFileRuntimeLog    string = ""
	// defaultFileErrorLog      string = ""
	// defaultFileOutput        string = ""

	// defaultDBServerEncryptMode is a string value representing one of
	// `disable`, `false` or `true`. These values directly match the valid
	// values for the `denisenkom/go-mssqldb` driver parameter named
	// `encrypt`.
	defaultDBServerEncryptMode string = EncryptModeFalse
)

// These values directly map to the `encrypt` parameter settings used by the
// database driver.
// https://github.com/denisenkom/go-mssqldb#common-parameters
const (
	EncryptModeDisable string = "disable"
	EncryptModeTrue    string = "true"
	EncryptModeFalse   string = "false"
)

// https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008-r2/ms143531(v=sql.105)
// https://docs.microsoft.com/en-us/sql/relational-databases/system-stored-procedures/sp-server-info-transact-sql
// https://docs.microsoft.com/en-us/sql/t-sql/statements/create-database-transact-sql
// https://docs.microsoft.com/en-us/sql/t-sql/statements/create-table-transact-sql?
// https://docs.microsoft.com/en-us/sql/t-sql/statements/create-user-transact-sql
const (
	MSSQLInstanceNameMaxChars      int = 16
	MSSQLUsernameMaxChars          int = 128
	MSSQLPasswordMaxChars          int = 128
	MSSQLDatabaseNameMaxChars      int = 123
	MSSQLDatabaseTableNameMaxChars int = 128
)

// TCP port ranges
// http://www.iana.org/assignments/port-numbers
// Port numbers are assigned in various ways, based on three ranges: System
// Ports (0-1023), User Ports (1024-49151), and the Dynamic and/or Private
// Ports (49152-65535)
const (
	TCPReservedPort            int = 0
	TCPSystemPortStart         int = 1
	TCPSystemPortEnd           int = 1023
	TCPUserPortStart           int = 1024
	TCPUserPortEnd             int = 49151
	TCPDynamicPrivatePortStart int = 49152
	TCPDynamicPrivatePortEnd   int = 65535
)
