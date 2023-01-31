/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/crd"
)

func newCRDCleanCommand() *cobra.Command {
	dryRun := "dry-run"

	cmd := &cobra.Command{
		Use:   "clean",
		Short: "clean deprecated CRD storedVersions",
		RunE: func(cmd *cobra.Command, args []string) error { // TODO: Should consider xcobra.RunWithCtx here
			cfg := config.GetConfigOrDie()
			apiExtClient, err := v1.NewForConfig(cfg)
			if err != nil {
				return errors.Wrap(err, "unable to create kubernetes client")
			}
			cl, err := client.New(cfg, client.Options{Scheme: controllers.CreateScheme()})
			if err != nil {
				return err
			}

			dryrun, err := cmd.Flags().GetBool(dryRun)
			if err != nil {
				return err
			}

			return crd.NewCleaner(
				apiExtClient.CustomResourceDefinitions(),
				cl,
				dryrun).Run(context.TODO())
		},
	}

	cmd.PersistentFlags().Bool(dryRun, false, "")

	return cmd
}
