/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	flag "github.com/spf13/pflag"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/vbauerster/mpb/v8"
)

var (
	// Global logger and progress bar
	// While these are global here, they can't be referenced from internal packages and should be passed through
	// as parameters where needed.
	log      logr.Logger
	progress *mpb.Progress

	verbose bool
)

func Configure(f *flag.FlagSet) {
	f.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
}

func Logger() logr.Logger {
	if verbose {
		zerologr.SetMaxV(1)
	}

	return log
}

func Progress() *mpb.Progress {
	return progress
}

func init() {

	// Create Progressbar for console output
	progress = mpb.New()

	// Configure console writer for ZeroLog
	output := zerolog.ConsoleWriter{
		Out:        progress,       // Write via the progressbar
		TimeFormat: "15:04:05.999", // Display time to the millisecond
	}

	// Create zerolog logger
	zl := zerolog.New(output).
		With().Timestamp().
		Logger()

	// Use standard interface for logging
	zerologr.VerbosityFieldName = "" // Don't include verbosity in output
	zerologr.SetMaxV(0)              // Default to quiet

	log = zerologr.New(&zl)
}
