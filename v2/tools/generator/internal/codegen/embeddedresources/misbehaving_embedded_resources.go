/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type misbehavingResourceCtx struct {
	resourceName astmodel.InternalTypeName
	typeName     astmodel.InternalTypeName
}

type misbehavingResourceDetails struct {
	resourceName astmodel.TypeName
	typeName     astmodel.TypeName
	propertyType astmodel.TypeName
}

func findMisbehavingResources(
	configuration *config.Configuration,
	defs astmodel.TypeDefinitionSet,
) (map[astmodel.InternalTypeName][]misbehavingResourceDetails, error) {
	resources := make(map[astmodel.InternalTypeName][]misbehavingResourceDetails)
	visitor := astmodel.TypeVisitorBuilder[misbehavingResourceCtx]{
		VisitObjectType: func(this *astmodel.TypeVisitor[misbehavingResourceCtx], ot *astmodel.ObjectType, ctx misbehavingResourceCtx) (astmodel.Type, error) {
			// If we don't have any configuration at all for the type, we don't need to check any of the properties
			if configuration.ObjectModelConfiguration.IsTypeConfigured(ctx.typeName) {
				for _, prop := range ot.Properties().Copy() {
					resourceLifecycleOwnedByParent, ok := configuration.ObjectModelConfiguration.ResourceLifecycleOwnedByParent.Lookup(ctx.typeName, prop.PropertyName())
					if !ok {
						continue
					}

					// If the property is a subresource whose lifecycle is owned by a parent resource, but we're not
					// examining the parent resource in question, continue
					if ctx.resourceName.Name() != resourceLifecycleOwnedByParent {
						continue
					}

					// Expected that the property in question is a TypeName
					propertyType, ok := astmodel.ExtractTypeName(prop.PropertyType())
					if !ok {
						return nil, eris.Errorf("property %s of %s doesn't look like a resource because it is not a TypeName", prop.PropertyName(), ctx.typeName.String())
					}

					if _, ok = resources[ctx.resourceName]; !ok {
						resources[ctx.resourceName] = []misbehavingResourceDetails{}
					}

					resources[ctx.resourceName] = append(
						resources[ctx.resourceName],
						misbehavingResourceDetails{
							typeName:     ctx.typeName,
							resourceName: ctx.resourceName,
							propertyType: propertyType,
						})
				}
			}

			return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
		},
	}.Build()

	typeWalker := astmodel.NewTypeWalker(defs, visitor)
	typeWalker.MakeContext = func(
		it astmodel.InternalTypeName,
		ctx misbehavingResourceCtx,
	) (misbehavingResourceCtx, error) {
		if ctx.resourceName.IsEmpty() {
			ctx.resourceName = it
		}

		ctx.typeName = it
		return ctx, nil
	}

	for _, def := range defs.AllResources() {
		_, err := typeWalker.Walk(def)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to walk type %s", def.Name())
		}
	}

	err := configuration.ObjectModelConfiguration.ResourceLifecycleOwnedByParent.VerifyConsumed()
	if err != nil {
		return nil, err
	}

	return resources, nil
}
