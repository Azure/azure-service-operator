/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package export

import (
	"context"

	"github.com/devigned/pub/pkg/xcobra"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func newExportResourceCommand() *cobra.Command {
	var output *string

	cmd := &cobra.Command{
		Use:   "resource <ARM/ID/of/resource>",
		Short: "exports an ARM resource CRD",
		Args:  cobra.ExactArgs(1),
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			armID := args[0]
			return exportResource(armID, output)
		}),
	}

	output = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write ARM resource CRD to a file")

	return cmd
}

// TODO: export resource logic goes here
func exportResource(armID string, output *string) error {
	klog.Infof("armID : %s", armID)

	if output != nil && *output != "" {
		klog.Infof("output : %s", *output)
	}

	return nil
}
