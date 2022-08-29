// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// loadConfigFile is a helper function to handle opening a specified config
// file and importing the settings for use
func (c *Config) loadConfigFile(configFile string) error {

	fh, err := os.Open(filepath.Clean(configFile))
	if err != nil {
		return err
	}

	// #nosec G307
	// Believed to be a false-positive from recent gosec release
	// https://github.com/securego/gosec/issues/714
	defer func() {
		if err := fh.Close(); err != nil {
			// Ignore "file already closed" errors
			if !errors.Is(err, os.ErrClosed) {
				fmt.Printf(
					"loadConfigFile: failed to close file %q: %s",
					configFile,
					err.Error(),
				)
			}
		}
	}()

	if err := c.ImportConfigFile(fh); err != nil {
		return err
	}

	return err
}

// ImportConfigFile reads from an io.Reader and unmarshals a configuration file
// in TOML format into the associated Config struct.
func (c *Config) ImportConfigFile(fh io.Reader) error {

	configFile, err := io.ReadAll(fh)
	if err != nil {
		return err
	}

	// target nested config struct dedicated to TOML config file settings
	if err := toml.Unmarshal(configFile, &c.fileConfig); err != nil {
		return err
	}

	return nil
}
