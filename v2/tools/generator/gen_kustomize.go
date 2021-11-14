/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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
		Use:   "gen-kustomize <path to config/crd folder>",
		Short: "generate K8s Kustomize file in the spirit of Kubebuilder, based on the specified config folder",
		Args:  cobra.ExactArgs(1),
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			crdPath := args[0]

			bases := "bases"
			// We have an expectation that the folder structure is: .../config/crd/bases
			basesPath := filepath.Join(crdPath, bases)
			destination := filepath.Join(crdPath, "kustomization.yaml")

			files, err := ioutil.ReadDir(basesPath)
			if err != nil {
				return logAndExtractStack(fmt.Sprintf("Unable to scan folder %q", basesPath), err)
			}

			result := kustomization.NewCRDKustomizeFile()
			for _, f := range files {
				if f.IsDir() {
					continue
				}

				result.AddResource(filepath.Join(bases, f.Name()))
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

//Extend this command verb to also export files into crd/patches to enable versioning
//Look at v1/config/default/crds to see what's needed, it's essentially identical for each resource

func logAndExtractStack(str string, err error) error {
	klog.Errorf("%s:\n%s\n", str, err)
	stackTrace := findDeepestTrace(err)
	if stackTrace != nil {
		klog.V(4).Infof("%s", stackTrace)
	}
	return err
}

