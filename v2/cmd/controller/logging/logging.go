/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package logging

import (
	"flag"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2/textlogger"
)

var (
	// verbose indicates whether we should use verbose logging.
	// Default is no
	verbose *bool

	// useJson indicates whether we should output logs in JSON format.
	// Default is no
	useJson *bool
)

// Create returns a new logger, ready for use.
// This can be called multiple times if required.
func Create() logr.Logger {
	if useJson != nil && *useJson {
		log, err := createJSONLogger()
		if err != nil {
			log = createTextLogger()
			log.Error(err, "failed to create JSON logger, falling back to text")
		}

		return log
	}

	return createTextLogger()
}

func createTextLogger() logr.Logger {
	opts := []textlogger.ConfigOption{}
	if verbose != nil && *verbose {
		opts = append(opts, textlogger.Verbosity(3))
	}

	cfg := textlogger.NewConfig(opts...)
	return textlogger.NewLogger(cfg)
}

func createJSONLogger() (logr.Logger, error) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	if verbose != nil && *verbose {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.Config{
		Level:            level,
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := cfg.Build()
	if err != nil {
		return logr.Logger{}, err
	}

	return zapr.NewLogger(logger), nil
}

// InitFlags initializes the flags for the logging package
func InitFlags(fs *flag.FlagSet) {
	verbose = fs.Bool("verbose", false, "Enable verbose logging")
	fs.BoolVar(verbose, "v", false, "Enable verbose logging")

	useJson = fs.Bool("json-logging", false, "Enable JSON logging")
}
