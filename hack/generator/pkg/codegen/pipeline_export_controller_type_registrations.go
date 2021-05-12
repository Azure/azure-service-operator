/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

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
			var storageVersionResources []astmodel.TypeName

			// We need to register each version
			for _, def := range types {
				resource, ok := astmodel.AsResourceType(def.Type())
				if !ok {
					continue
				}

				if resource.IsStorageVersion() {
					storageVersionResources = append(storageVersionResources, def.Name())
				}

				resources = append(resources, def.Name())
			}

			file := NewResourceRegistrationFile(resources, storageVersionResources)
			fileWriter := astmodel.NewGoSourceFileWriter(file)

			err := fileWriter.SaveToFile(outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to write controller type registration file to %q", outputPath)
			}

			return types, nil
		})
}
