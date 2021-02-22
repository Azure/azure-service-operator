/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// exportControllerResourceRegistrations creates a PipelineStage to generate type registrations
// for resources.
func exportControllerResourceRegistrations(outputPath string) PipelineStage {
	description := fmt.Sprintf("Export resource registrations to %q", outputPath)
	return MakePipelineStage(
		"exportControllerResourceRegistrations",
		description,
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			// If the configuration doesn't specify an output destination for us, just do nothing
			if outputPath == "" {
				return types, nil
			}

			var resources []astmodel.TypeName

			// This is a bit hacky but it's just for logging
			excludedGroupVersions := make(map[string]struct{})

			// We need to register each version
			for _, def := range types {
				resource, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					continue
				}
				// TODO: Stop excluding storage versions when they're ready for prime time
				if resource.IsStorageVersion() {
					localPkg, err := astmodel.PackageAsLocalPackage(def.Name().PackageReference)
					if err != nil {
						return nil, err
					}
					excludedGroupVersions[localPkg.Group()+localPkg.Version()] = struct{}{}
					continue
				}

				resources = append(resources, def.Name())
			}

			for excluded := range excludedGroupVersions {
				klog.Warningf("Excluded package %s from registration in %s", excluded, outputPath)
			}

			file := NewResourceRegistrationFile(resources)
			fileWriter := astmodel.NewGoSourceFileWriter(file)

			err := fileWriter.SaveToFile(outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to write controller type registration file to %q", outputPath)
			}

			return types, nil
		})
}
