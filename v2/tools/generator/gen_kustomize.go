/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/kustomization"
)

// NewGenKustomizeCommand creates a new cobra Command when invoked from the command line
func NewGenKustomizeCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		// TODO: there's not great support for required
		// TODO: arguments in cobra so this is the best we get... see:
		// TODO: https://github.com/spf13/cobra/issues/395
		Use:   "gen-kustomize <path to config/crd/generated folder>",
		Short: "generate K8s Kustomize file in the spirit of Kubebuilder, based on the specified config folder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			crdPath := args[0]

			const bases = "bases"
			const patches = "patches"

			log := CreateLogger()

			// We have an expectation that the folder structure is: .../config/crd/generated/bases and .../config/crd/generated/patches
			basesPath := filepath.Join(crdPath, bases)
			patchesPath := filepath.Join(crdPath, patches)

			destination := filepath.Join(crdPath, "kustomization.yaml")

			log.Info(
				"Scanning for resources",
				"basePath", basesPath)

			files, err := os.ReadDir(basesPath)
			if err != nil {
				log.Error(err, "Unable to scan folder", "folder", basesPath)
				return err
			}

			err = os.MkdirAll(patchesPath, os.ModePerm)
			if err != nil {
				log.Error(err, "Unable to create output folder", "folder", patchesPath)
				return err
			}

			var errs []error
			result := kustomization.NewCRDKustomizeFile()

			for _, f := range files {
				if f.IsDir() {
					continue
				}

				log.V(1).Info(
					"Found resource file",
					"file", f.Name())

				patchFile := "webhook-conversion-" + f.Name()
				var def *kustomization.ResourceDefinition
				def, err = kustomization.LoadResourceDefinition(filepath.Join(basesPath, f.Name()))
				if err != nil {
					errs = append(errs, err)
					continue
				}

				log.V(1).Info(
					"Loaded Resource",
					"name", def.Name())

				patch := kustomization.NewConversionPatchFile(def.Name())
				err = patch.Save(filepath.Join(patchesPath, patchFile))
				if err != nil {
					errs = append(errs, err)
					continue
				}

				result.AddResource(filepath.Join(bases, f.Name()))
				result.AddPatch(filepath.Join(patches, patchFile))
			}

			if len(errs) > 0 {
				err = kerrors.NewAggregate(errs)
				log.Error(
					err,
					"Error creating conversion patches")
				return err
			}

			if len(result.Resources) == 0 {
				err = errors.Errorf("no files found in %q", basesPath)
				log.Error(
					err,
					"No CRD files found")
				return err
			}

			err = result.Save(destination)
			if err != nil {
				log.Error(
					err,
					"Error generating",
					"destination", destination)
				return err
			}

			return nil
		},
	}

	return cmd, nil
}
