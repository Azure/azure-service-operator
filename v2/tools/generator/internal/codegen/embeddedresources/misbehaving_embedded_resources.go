/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type misbehavingResourceCtx struct {
	resourceName astmodel.TypeName
	typeName     astmodel.TypeName
}

type misbehavingResourceDetails struct {
	resourceName astmodel.TypeName
	typeName     astmodel.TypeName
	propertyType astmodel.TypeName
}

func findMisbehavingResources(configuration *config.Configuration, defs astmodel.TypeDefinitionSet) (map[astmodel.TypeName][]misbehavingResourceDetails, error) {
	resources := make(map[astmodel.TypeName][]misbehavingResourceDetails)
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			typedCtx := ctx.(misbehavingResourceCtx)
			for _, prop := range ot.Properties().Copy() {
				resourceLifecycleOwnedByParent, err := configuration.ResourceLifecycleOwnedByParent(typedCtx.typeName, prop.PropertyName())
				if err != nil {
					if config.IsNotConfiguredError(err) {
						continue
					}
					return nil, errors.Wrap(err, "unexpected error checking config")
				}

				// If the property is a subresource whose lifecycle is owned by a parent resource, but we're not
				// examining the parent resource in question, continue
				if typedCtx.resourceName.Name() != resourceLifecycleOwnedByParent {
					continue
				}

				// Expected that the property in question is a TypeName
				propertyType, ok := astmodel.ExtractTypeName(prop.PropertyType())
				if !ok {
					return nil, errors.Errorf("property %s of %s doesn't look like a resource because it is not a TypeName", prop.PropertyName(), typedCtx.typeName.String())
				}

				if _, ok = resources[typedCtx.resourceName]; !ok {
					resources[typedCtx.resourceName] = []misbehavingResourceDetails{}
				}

				resources[typedCtx.resourceName] = append(
					resources[typedCtx.resourceName],
					misbehavingResourceDetails{
						typeName:     typedCtx.typeName,
						resourceName: typedCtx.resourceName,
						propertyType: propertyType,
					})
			}

			return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
		},
	}.Build()

	typeWalker := astmodel.NewTypeWalker(defs, visitor)
	typeWalker.MakeContext = func(it astmodel.TypeName, ctx interface{}) (interface{}, error) {
		if ctx == nil {
			ctx = misbehavingResourceCtx{resourceName: it}
		}

		typedCtx := ctx.(misbehavingResourceCtx)
		typedCtx.typeName = it
		return typedCtx, nil
	}

	for _, def := range astmodel.FindResourceDefinitions(defs) {
		_, err := typeWalker.Walk(def)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to walk type %s", def.Name())
		}
	}

	err := configuration.VerifyResourceLifecycleOwnedByParentConsumed()
	if err != nil {
		return nil, err
	}

	return resources, nil
}
