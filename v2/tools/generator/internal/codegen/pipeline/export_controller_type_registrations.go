/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ExportControllerResourceRegistrations creates a Stage to generate type registrations
// for resources.
func ExportControllerResourceRegistrations(outputPath string) Stage {
	return MakeLegacyStage(
		"exportControllerResourceRegistrations",
		fmt.Sprintf("Export resource registrations to %q", outputPath),
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			// If the configuration doesn't specify an output destination for us, just do nothing
			if outputPath == "" {
				return types, nil
			}

			var resources []astmodel.TypeName
			var storageVersionResources []astmodel.TypeName
			var resourceExtensions []astmodel.TypeName

			// We need to register each version
			for _, def := range types {

				if resource, ok := astmodel.AsResourceType(def.Type()); ok {

					if resource.IsStorageVersion() {
						storageVersionResources = append(storageVersionResources, def.Name())
					}

					resources = append(resources, def.Name())
				} else if resourceExtension, ok := astmodel.AsObjectType(def.Type()); ok {

					if resourceExtension.HasFunctionWithName(functions.ExtendedResourcesFunctionName) {
						resourceExtensions = append(resourceExtensions, def.Name())
					}
				}

			}

			file := NewResourceRegistrationFile(resources, storageVersionResources, resourceExtensions)
			fileWriter := astmodel.NewGoSourceFileWriter(file)

			err := fileWriter.SaveToFile(outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to write controller type registration file to %q", outputPath)
			}

			return types, nil
		})
}
