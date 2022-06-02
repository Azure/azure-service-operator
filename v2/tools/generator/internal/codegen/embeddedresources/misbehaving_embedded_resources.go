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

func findMisbehavingResources(configuration *config.Configuration, defs astmodel.TypeDefinitionSet) (astmodel.TypeNameSet, error) {
	resources := make(astmodel.TypeNameSet)
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			typedCtx := ctx.(misbehavingResourceCtx)
			for _, prop := range ot.Properties().Copy() {
				isResourceLifecycleOwnedByParent, err := configuration.IsResourceLifecycleOwnedByParent(typedCtx.typeName, prop.PropertyName())
				if err != nil && !config.IsNotConfiguredError(err) {
					return nil, errors.Wrap(err, "unexpected error checking config")
				}

				if isResourceLifecycleOwnedByParent {
					resources.Add(typedCtx.resourceName)
				}
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

	err := configuration.VerifyIsResourceLifecycleOwnedByParentConsumed()
	if err != nil {
		return nil, err
	}

	return resources, nil
}
