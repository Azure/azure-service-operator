/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

/* augmentResourcesWithSwaggerInformation creates a PipelineStage to add status information into the generated resources.

This information is derived from the Azure Swagger specifications. We parse the Swagger specs and look for
any actions that appear to be ARM resources (have PUT methods with types we can use and appropriate names in the
action path). Then for each resource, we use the existing JSON AST parser to extract the status type
(the type-definition part of swagger is the same as JSON Schema).

Next, we walk over all the resources we are currently generating CRDs for and attempt to locate
a match for the resource in the status information we have parsed. If we locate a match, it is
added to the Status field of the Resource type, after we have renamed all the status types to
avoid any conflicts with existing Spec types that have already been defined.

*/
func augmentResourcesWithSwaggerInformation(idFactory astmodel.IdentifierFactory, config *config.Configuration) PipelineStage {
	return MakePipelineStage(
		"augmentWithSwagger",
		"Add information from Swagger specs for 'status' fields",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			if config.Status.SchemaRoot == "" {
				klog.Warningf("No status schema root specified, will not generate status types")
				return types, nil
			}

			klog.V(1).Infof("Loading Swagger data from %q", config.Status.SchemaRoot)

			swaggerTypes, err := loadSwaggerData(ctx, idFactory, config)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load Swagger data")
			}

			klog.V(1).Infof("Loaded Swagger data (%v resources, %v other types)", len(swaggerTypes.resources), len(swaggerTypes.otherTypes))

			newTypes := make(astmodel.Types)

			statusTypes, err := generateStatusTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			found := 0
			for typeName, typeDef := range types {
				// for resources, try to find the matching Status type
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					newStatus, located := statusTypes.findResourceType(typeName)
					if located {
						found++
					}

					newTypes.Add(astmodel.MakeTypeDefinition(typeName, resource.WithStatus(newStatus)))
				} else {
					// other types are simply copied
					newTypes.Add(typeDef)
				}
			}

			// all non-resources from Swagger are added regardless of whether they are used
			// if they are not used they will be pruned off by a later pipeline stage
			// (there will be no name clashes here due to suffixing with "_Status")
			newTypes.AddAll(statusTypes.otherTypes)

			klog.V(1).Infof("Found status information for %v resources", found)
			klog.V(1).Infof("Input %v types, output %v types", len(types), len(newTypes))

			return newTypes, nil
		})
}

// an augment adds information from the Swagger-derived type to
// the main JSON schema-derived type, and returns the new type
type augment func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error)

// fuseAugments merges multiple augments into one by applying each
// augment in order to the result of the previous
func fuseAugments(augments ...augment) augment {
	return func(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
		for _, augment := range augments {
			newMain, err := augment(main, swagger)
			if err != nil {
				return nil, err
			}

			main = newMain
		}

		return main, nil
	}
}

// the overall augmenter we will use
var augmenter augment = fuseAugments(flattenAugment)

// the flattenAugment
func flattenAugment(main astmodel.Type, swagger astmodel.Type) (astmodel.Type, error) {
	var merger = astmodel.NewTypeMerger(func(ctx interface{}, left, right astmodel.Type) (astmodel.Type, error) {
		// return left (= main) if the two sides are not ObjectTypes
		return left, nil
	})

	merger.Add(func(main, swagger *astmodel.ObjectType) (astmodel.Type, error) {
		props := main.Properties()
		for ix, mainProp := range props {
			if swaggerProp, ok := swagger.Property(mainProp.PropertyName()); ok {
				// first copy over flatten property
				mainProp = mainProp.SetFlatten(swaggerProp.Flatten())

				// now recursively merge property types
				newType, err := merger.Merge(mainProp.PropertyType(), swaggerProp.PropertyType())
				if err != nil {
					return nil, err
				}

				props[ix] = mainProp.WithType(newType)
			}
		}

		return main.WithProperties(props...), nil
	})

	return merger.Merge(main, swagger)
}
