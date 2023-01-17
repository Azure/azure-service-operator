/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/crd"
)

func newCRDCleanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clean",
		Short: "clean deprecated CRD storedVersions",
		RunE: func(cmd *cobra.Command, args []string) error { // TODO: Should consider xcobra.RunWithCtx here
			cfg := config.GetConfigOrDie()
			cl, err := v1.NewForConfig(cfg)
			if err != nil {
				return errors.Wrap(err, "unable to create kubernetes client")
			}

			return crd.CleanDeprecatedCRDVersions(context.TODO(), cl.CustomResourceDefinitions())
		},
	}

	return cmd
}
