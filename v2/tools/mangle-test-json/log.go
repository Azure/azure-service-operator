package main

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
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
	zerologr.SetMaxV(0)              // Default to quiet

	log := zerologr.New(&zl)
	return log
}
