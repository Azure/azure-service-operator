/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/Azure/azure-service-operator/v2/api"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/crd"
)

// newCRDCleanCommand creates a new cobra command for cleaning deprecated CRDs from a cluster
func newCRDCleanCommand() *cobra.Command {
	dryRunFlagName := "dry-run"
	dryRun := false

	cmd := &cobra.Command{
		Use:   "clean",
		Short: "Clean deprecated CRD versions from cluster",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.GetConfig()
			if err != nil {
				return errors.Wrap(err, "unable to get Kubernetes config")
			}

			ctx := cmd.Context()

			apiExtClient, err := v1.NewForConfig(cfg)
			if err != nil {
				return errors.Wrap(err, "unable to create Kubernetes client")
			}

			cl, err := client.New(cfg, client.Options{Scheme: api.CreateScheme()})
			if err != nil {
				return errors.Wrap(err, "unable to create Kubernetes client")
			}

			return crd.NewCleaner(
				apiExtClient.CustomResourceDefinitions(),
				cl,
				dryRun,
				log).Run(ctx)
		},
	}

	cmd.Flags().BoolVar(&dryRun, dryRunFlagName, false, "Dry run (don't actually clean anything)")

	return cmd
}
