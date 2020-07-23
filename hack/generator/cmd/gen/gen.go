/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package gen

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/codegen"
	"github.com/Azure/k8s-infra/hack/generator/pkg/xcobra"
)

// NewGenCommand creates a new cobra Command when invoked from the command line
func NewGenCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		// TODO: there's not great support for required
		// TODO: arguments in cobra so this is the best we get... see:
		// TODO: https://github.com/spf13/cobra/issues/395
		Use:   "gen <config>",
		Short: "generate K8s infrastructure resources from Azure deployment template schema",
		Args:  cobra.ExactArgs(1),
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			configFile := args[0]

			cg, err := codegen.NewCodeGeneratorFromConfigFile(configFile)
			if err != nil {
				klog.Errorf("Error creating code generator: %v\n", err)
				return err
			}

			err = cg.Generate(ctx)
			if err != nil {
				klog.Errorf("Error during code generation:\n%v\n", err)
				stackTrace := findDeepestTrace(err)
				if stackTrace != nil {
					klog.V(4).Infof("%+v", stackTrace)
				}
				return err
			}

			return nil
		}),
	}

	cmd.Flags().StringArrayP("resources", "r", nil, "list of resource type / versions to generate")
	if err := viper.BindPFlag("resources", cmd.Flags().Lookup("resources")); err != nil {
		return cmd, err
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
	if err == nil {
		return nil
	}
	deeperTrace := findDeepestTrace(errors.Unwrap(err))
	if deeperTrace != nil {
		return deeperTrace
	}
	tracer, ok := err.(stackTracer)
	if !ok {
		return nil
	}
	return tracer.StackTrace()
}
