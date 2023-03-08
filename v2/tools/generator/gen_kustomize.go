/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/kustomization"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/xcobra"
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
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			crdPath := args[0]

			const bases = "bases"
			const patches = "patches"

			// We have an expectation that the folder structure is: .../config/crd/generated/bases and .../config/crd/generated/patches
			basesPath := filepath.Join(crdPath, bases)
			patchesPath := filepath.Join(crdPath, patches)

			destination := filepath.Join(crdPath, "kustomization.yaml")

			klog.V(3).Infof("Scanning %q for resources", basesPath)

			files, err := ioutil.ReadDir(basesPath)
			if err != nil {
				return logAndExtractStack(fmt.Sprintf("Unable to scan folder %q", basesPath), err)
			}

			err = os.MkdirAll(patchesPath, os.ModePerm)
			if err != nil {
				return logAndExtractStack(fmt.Sprintf("Unable to create output folder %s", patchesPath), err)
			}

			var errs []error
			result := kustomization.NewCRDKustomizeFile()

			for _, f := range files {
				if f.IsDir() {
					continue
				}

				klog.V(3).Infof("Found resource file %s", f.Name())

				patchFile := "webhook-conversion-" + f.Name()
				var def *kustomization.ResourceDefinition
				def, err = kustomization.LoadResourceDefinition(filepath.Join(basesPath, f.Name()))
				if err != nil {
					errs = append(errs, err)
					continue
				}

				klog.V(4).Infof("Resource is %q", def.Name())

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
				return logAndExtractStack("Error creating conversion patches", kerrors.NewAggregate(errs))
			}

			if len(result.Resources) == 0 {
				err = errors.Errorf("no files found in %q", basesPath)
				return logAndExtractStack("No CRD files found", err)
			}

			err = result.Save(destination)
			if err != nil {
				return logAndExtractStack("Error generating "+destination, err)
			}

			return nil
		}),
	}

	return cmd, nil
}

func logAndExtractStack(str string, err error) error {
	klog.Errorf("%s:\n%s\n", str, err)
	if tr, ok := findDeepestTrace(err); ok {
		for _, fr := range tr {
			klog.Errorf("%n (%s:%d)", fr, fr, fr)
		}
	}

	return err
}
