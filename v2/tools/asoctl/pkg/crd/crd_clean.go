/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/devigned/pub/pkg/xcobra"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/crd"
)

func newCRDCleanCommand() *cobra.Command {
	dryRunFlagName := "dry-run"

	cmd := &cobra.Command{
		Use:   "clean",
		Short: "clean deprecated CRD storedVersions",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			cfg := config.GetConfigOrDie()

			apiExtClient, err := v1.NewForConfig(cfg)
			if err != nil {
				return errors.Wrap(err, "unable to create kubernetes client")
			}

			cl, err := client.New(cfg, client.Options{Scheme: controllers.CreateScheme()})
			if err != nil {
				return errors.Wrap(err, "unable to create kubernetes client")
			}

			dryRun, err := cmd.Flags().GetBool(dryRunFlagName)
			if err != nil {
				return err
			}

			return crd.NewCleaner(
				apiExtClient.CustomResourceDefinitions(),
				cl,
				dryRun).Run(ctx)
		}),
	}

	cmd.PersistentFlags().Bool(dryRunFlagName, false, "")

	return cmd
}
