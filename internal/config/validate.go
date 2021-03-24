// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"unicode"

	"github.com/atc0005/query-meta/internal/caller"
)

// validate verifies that user-provided and/or default values are acceptable.
//
// In most cases (e.g., where a setting is required), getter methods are
// checked instead of directly referencing the config struct. getter methods
// pass user-provided values through without modification; if a user did not
// specify a value, the default value is passed through for validation.
func (c Config) validate() error {

	myFuncName := caller.GetFuncName()

	if c.cliConfig.ConfigFile == nil || *c.cliConfig.ConfigFile == "" {
		return fmt.Errorf(
			"%s: fully-qualified path to configuration file not provided",
			myFuncName,
		)
	}

	switch c.LogLevel() {
	case LogLevelDisabled:
	case LogLevelPanic:
	case LogLevelFatal:
	case LogLevelError:
	case LogLevelWarn:
	case LogLevelInfo:
	case LogLevelDebug:
	case LogLevelTrace:
	default:
		return fmt.Errorf(
			"%s: invalid log level provided: %v",
			myFuncName,
			c.LogLevel(),
		)
	}

	if c.DBServerHost() == "" {
		return fmt.Errorf(
			"%s: missing database host name or IP Address",
			myFuncName,
		)
	}

	if !(c.DBServerPort() >= TCPSystemPortStart) && (c.DBServerPort() <= TCPDynamicPrivatePortEnd) {
		return fmt.Errorf(
			"%s: invalid port %d specified; outside of the range of %d and %d",
			myFuncName,
			c.DBServerPort(),
			TCPSystemPortStart,
			TCPDynamicPrivatePortEnd,
		)
	}

	// TODO: Extend this as needed to apply proper validation
	// https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008-r2/ms143531(v=sql.105)
	// https://stackoverflow.com/questions/5260650/max-length-of-sql-server-instance-name
	switch {
	case c.DBServerInstance() == "":
		if c.fileConfig.Instance != nil {
			return fmt.Errorf(
				"%s: missing database host instance",
				myFuncName,
			)
		}

	case !unicode.IsLetter(rune(c.DBServerInstance()[0])):
		return fmt.Errorf(
			"%s: specified instance name does not begin with a known Unicode letter; got %q, expected letter instead",
			myFuncName,
			c.DBServerInstance(),
		)

	case len(c.DBServerInstance()) > MSSQLInstanceNameMaxChars:
		return fmt.Errorf(
			"%s: specified instance name too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerInstance()),
			MSSQLInstanceNameMaxChars,
		)
	}

	switch {
	case c.DBServerUsername() == "":
		return fmt.Errorf(
			"%s: missing database host username",
			myFuncName,
		)
	case len(c.DBServerUsername()) > MSSQLUsernameMaxChars:
		return fmt.Errorf(
			"%s: specified username too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerUsername()),
			MSSQLUsernameMaxChars,
		)
	}

	switch {
	case c.DBServerPassword() == "":
		return fmt.Errorf(
			"%s: missing database host password",
			myFuncName,
		)
	case len(c.DBServerPassword()) > MSSQLPasswordMaxChars:
		return fmt.Errorf(
			"%s: specified password too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerPassword()),
			MSSQLPasswordMaxChars,
		)
	}

	switch c.DBServerEncryptMode() {
	case EncryptModeDisable:
	case EncryptModeTrue:
	case EncryptModeFalse:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for connection encrypt mode",
			myFuncName,
			c.DBServerEncryptMode(),
		)
	}

	// DBServerTrustCert returns a boolean value, so nothing to test here.

	switch {
	case c.DBName() == "":
		return fmt.Errorf(
			"%s: missing database name",
			myFuncName,
		)

	case len(c.DBName()) > MSSQLDatabaseNameMaxChars:
		return fmt.Errorf(
			"%s: specified database name too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBName()),
			MSSQLDatabaseNameMaxChars,
		)

	}

	switch {
	case c.DBTable() == "":
		return fmt.Errorf(
			"%s: missing database table name",
			myFuncName,
		)

	case len(c.DBTable()) > MSSQLDatabaseTableNameMaxChars:
		return fmt.Errorf(
			"%s: specified database name too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBTable()),
			MSSQLDatabaseNameMaxChars,
		)

	}

	if c.DBQueryCountAllRecords() == "" {
		return fmt.Errorf(
			"%s: missing SQL query for determining count of all patron records",
			myFuncName,
		)
	}

	if c.DBQueryCountInactiveRecords() == "" {
		return fmt.Errorf(
			"%s: missing SQL query for determining count of inactive patron records",
			myFuncName,
		)
	}

	if c.DBQueryRetrieveActivePatronRecords() == "" {
		return fmt.Errorf(
			"%s: missing SQL query for retrieving active patron records",
			myFuncName,
		)
	}

	// 	if c.FileRuntimeLog() == "" {
	// 		return fmt.Errorf(
	// 			"%s: missing fully-qualified path to runtime or operational messages log file",
	// 			myFuncName,
	// 		)
	// 	}
	//
	// 	if c.FileErrorLog() == "" {
	// 		return fmt.Errorf(
	// 			"%s: missing fully-qualified path to error messages log file",
	// 			myFuncName,
	// 		)
	// 	}
	//
	// 	if c.FileOutput() == "" {
	// 		return fmt.Errorf(
	// 			"%s: missing fully-qualified path to output file to be generated by this application",
	// 			myFuncName,
	// 		)
	// 	}

	return nil

}
