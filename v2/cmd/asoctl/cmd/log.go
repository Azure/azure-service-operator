/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/vbauerster/mpb/v8"
)

var (
	verbose bool // True if the user specifies --verbose or -v
	quiet   bool // True if the user specifies --quiet or -s
)

// CreateLogger creates a logger  for console output.
// Use this when your command wants to show only log messages
func CreateLogger() logr.Logger {
	// Configure console writer for ZeroLog
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,      // Write to StdErr
		TimeFormat: "15:04:05.999", // Display time to the millisecond
	}

	// Create zerolog logger
	zl := zerolog.New(output).
		With().Timestamp().
		Logger()

	// Use standard interface for logging
	zerologr.VerbosityFieldName = "" // Don't include verbosity in output
	log := zerologr.New(&zl)

	return log
}

// CreateLoggerAndProgressBar creates a logger and progress bar for console output.
// Use this when your command wants to show progress to the user.
func CreateLoggerAndProgressBar() (logr.Logger, *mpb.Progress) {
	// Create Progressbar for console output
	progress := mpb.New(
		mpb.WithOutput(os.Stderr), // Write to StdErr
	)

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

	log := zerologr.New(&zl)
	return log, progress
}
