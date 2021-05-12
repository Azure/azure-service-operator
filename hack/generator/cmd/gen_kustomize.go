/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/xcobra"
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
				return err
			}

			result := crdKustomizeFile{
				Resources: nil,
			}

			for _, f := range files {
				if f.IsDir() {
					continue
				}
				result.Resources = append(result.Resources, filepath.Join(bases, f.Name()))
			}

			data, err := yaml.Marshal(result)
			if err != nil {
				return logAndExtractStack("Error during kustomize.yaml serialization", err)
			}

			err = ioutil.WriteFile(destination, data, 0644) // #nosec G306
			if err != nil {
				return logAndExtractStack("Error during kustomize.yaml writing", err)
			}

			return nil
		}),
	}

	return cmd, nil
}

func logAndExtractStack(str string, err error) error {
	klog.Errorf("%s:\n%v\n", str, err)
	stackTrace := findDeepestTrace(err)
	if stackTrace != nil {
		klog.V(4).Infof("%+v", stackTrace)
	}
	return err
}

type crdKustomizeFile struct {
	Resources []string `yaml:"resources"`
}
