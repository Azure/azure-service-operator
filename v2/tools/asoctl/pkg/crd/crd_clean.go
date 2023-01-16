/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/client"
	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/crd"
)

func newCRDCleanCommand() (*cobra.Command, error) { //nolint:unparam
	cmd := &cobra.Command{
		Use:   "clean",
		Short: "clean deprecated CRD storedVersions",
		RunE: func(cmd *cobra.Command, args []string) error { // TODO: Should consider xcobra.RunWithCtx here
			cl, err := client.NewClient()
			if err != nil {
				return err
			}

			return crd.CleanDeprecatedCRDVersions(context.TODO(), cl.CustomResourceDefinitions())
		},
	}

	return cmd, nil
}
