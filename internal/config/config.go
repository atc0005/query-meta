// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/atc0005/query-meta/internal/caller"
)

// see `constants.go`, `logging.go` for other related values

// version reflects the application version. This is overridden via Makefile
// for release builds.
var version = "dev build"

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// String implements the Stringer interface in order to display all
// initialized (user-provided or default) values.
func (c Config) String() string {
	return fmt.Sprintf(
		"{ Host: %v, "+
			"Port: %v, "+
			"Instance: %v, "+
			"Username: %v, "+
			"Password: %v, "+
			"Database: %v, "+
			"EncryptMode: %v, "+
			"TrustCert: %v, "+
			"LogLevel: %v, "+
			"DBTable: %v, "+
			"DBQueryCountAllRecords: %v, "+
			"DBQueryCountInactiveRecords: %v, "+
			"DBQueryRetrieveActivePatronRecords: %v, "+

			c.DBServerHost(),
		c.DBServerPort(),
		c.DBServerInstance(),
		c.DBServerUsername(),
		"REDACTED",
		c.DBName(),
		c.DBServerEncryptMode(),
		c.DBServerTrustCert(),
		c.LogLevel(),
		c.DBTable(),
		c.DBQueryCountAllRecords(),
		c.DBQueryCountInactiveRecords(),
		c.DBQueryRetrieveActivePatronRecords(),
	)
}

// Version reuses the package-level Version function to emit version
// information and associated branding details whenever the user specifies the
// `--version` flag. The application exits after displaying this information.
func (c Config) Version() string {
	return Version() + "\n"
}

// Description emits branding information whenever the user specifies the `-h`
// flag. The application uses this as a header prior to displaying available
// CLI flag options.
func (c Config) Description() string {
	return fmt.Sprintf("\n%s", myAppDescription)
}

// New is a factory function that produces a new Config object based on user
// provided flag and where applicable, default values.
func New() (*Config, error) {

	myFuncName := caller.GetFuncName()

	var config Config

	// Bundle the returned `*.arg.Parser` for potential later use.
	config.flagParser = arg.MustParse(&config)

	// Attempt to load requested config file if specified
	if config.ConfigFile != nil && *config.ConfigFile != "" {

		if err := config.loadConfigFile(*config.ConfigFile); err != nil {
			return nil, fmt.Errorf(
				"%s: failed to load config file: %w",
				myFuncName,
				err,
			)
		}
	}

	if err := config.validate(); err != nil {
		config.flagParser.WriteUsage(os.Stderr)

		return nil, fmt.Errorf(
			"%s: failed to validate configuration: %w",
			myFuncName,
			err,
		)
	}

	config.configureLogging()

	return &config, nil

}
