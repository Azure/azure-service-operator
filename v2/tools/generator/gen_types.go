/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/xcobra"
)

// NewGenTypesCommand creates a new cobra Command when invoked from the command line
func NewGenTypesCommand() (*cobra.Command, error) {
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

			err = cg.Generate(ctx)
			if err != nil {
				return logAndExtractStack("Error during code generation", err)
			}

			return nil
		}),
	}

	return cmd, nil
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// findDeepestTrace returns the stack trace from the furthest error
// down the chain that has one. We can't just use errors.Cause(err)
// here because the innermost error may not have been created by
// pkg/errors (gasp).
func findDeepestTrace(err error) errors.StackTrace {
	var tracer stackTracer
	if errors.As(err, &tracer) {
		return tracer.StackTrace()
	}

	return nil
}
