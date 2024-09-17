/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/vbauerster/mpb/v8"
)

var (
	verbose bool // True if the user specifies --verbose or -v
	quiet   bool // True if the user specifies --quiet or -q
)

// CreateLogger creates a logger  for console output.
// Use this when your command wants to show only log messages
func CreateLogger() logr.Logger {
	return createLoggerCore(os.Stderr)
}

// CreateLoggerAndProgressBar creates a logger and progress bar for console output.
// Use this when your command wants to show progress to the user.
func CreateLoggerAndProgressBar() (logr.Logger, *mpb.Progress) {
	// Create Progressbar for console output
	progress := mpb.New(
		mpb.WithOutput(os.Stderr), // Write to StdErr
	)

	log := createLoggerCore(progress)
	return log, progress
}

func createLoggerCore(writer io.Writer) logr.Logger {
	// Configure console writer for ZeroLog
	output := zerolog.ConsoleWriter{
		Out:        writer,
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
