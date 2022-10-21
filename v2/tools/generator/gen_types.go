/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/xcobra"
)

// NewGenTypesCommand creates a new cobra Command when invoked from the command line
func NewGenTypesCommand() (*cobra.Command, error) {
	var debugMode *string

	cmd := &cobra.Command{
		// TODO: there's not great support for required
		// TODO: arguments in cobra so this is the best we get... see:
		// TODO: https://github.com/spf13/cobra/issues/395
		Use:   "gen-types <config>",
		Short: "generate K8s resources from Azure deployment template schema",
		Args:  cobra.ExactArgs(1),
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			configFile := args[0]

			cg, err := codegen.NewCodeGeneratorFromConfigFile(configFile)
			if err != nil {
				klog.Errorf("Error creating code generator: %s\n", err)
				return err
			}

			if debugMode != nil && *debugMode != "" {
				var tmpDir string
				tmpDir, err = ioutil.TempDir("", createDebugPrefix(*debugMode))
				if err != nil {
					klog.Errorf("Error creating temporary directory: %s\n", err)
					return err
				}

				klog.V(0).Infof("Debug output will be written to the folder %s\n", tmpDir)
				cg.UseDebugMode(*debugMode, tmpDir)
				defer func() {
					// Write the debug folder again so the user doesn't have to scroll back
					klog.V(0).Infof("Debug output is available in folder %s\n", tmpDir)
				}()
			}

			err = cg.Generate(ctx)

			if err != nil {
				return logAndExtractStack("Error during code generation", err)
			}

			return nil
		}),
	}

	debugMode = cmd.Flags().StringP(
		"debug",
		"d",
		"",
		"Write debug logs to a temp folder for a group (e.g. compute), multiple groups (e.g. compute;network), or groups matching a wildcard (e.g. net*)")

	return cmd, nil
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// findDeepestTrace returns the stack trace from the furthest error
// down the chain that has one. We can't just use errors.Cause(err)
// here because the innermost error may not have been created by
// pkg/errors (gasp).
func findDeepestTrace(err error) (errors.StackTrace, bool) {
	nested := errors.Unwrap(err)
	if nested != nil {
		if tr, ok := findDeepestTrace(nested); ok {
			// We've found the deepest trace, ,return it
			return tr, true
		}
	}

	// No stack trace found (yet), see if we have it at this level
	//nolint:errorlint // We're walking wrapped errors ourselves
	if tracer, ok := err.(stackTracer); ok {
		return tracer.StackTrace(), true
	}

	// No stack found at this, or any deeper, level
	return nil, false
}

func createDebugPrefix(debugMode string) string {
	var builder strings.Builder
	builder.WriteString("aso-gen-debug-")

	for _, r := range debugMode {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}

	builder.WriteRune('-')

	return builder.String()
}
