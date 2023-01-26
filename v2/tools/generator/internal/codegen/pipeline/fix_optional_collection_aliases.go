/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const FixOptionalCollectionAliasesStageId = "fixOptionalCollectionAliases"

func FixOptionalCollectionAliases() *Stage {
	return NewStage(
		FixOptionalCollectionAliasesStageId,
		"Replace types which are optional aliases to collections with just the collection alias",
		func(ctx context.Context, state *State) (*State, error) {

			fixer := optionalCollectionAliasFixer{
				definitions: state.Definitions(),
			}
			fixer.visitor = astmodel.TypeVisitorBuilder{
				VisitOptionalType: fixer.fixOptionalCollectionAliases,
			}.Build()

			results := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				d, err := fixer.visitor.VisitDefinition(def, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}
				results.Add(d)
			}

			return state.WithDefinitions(results), nil
		},
	)
}

type optionalCollectionAliasFixer struct {
	definitions astmodel.TypeDefinitionSet
	visitor     astmodel.TypeVisitor
}

func (f *optionalCollectionAliasFixer) fixOptionalCollectionAliases(this *astmodel.TypeVisitor, it *astmodel.OptionalType, ctx interface{}) (astmodel.Type, error) {
	typeName, ok := astmodel.AsTypeName(it)
	if !ok {
		return astmodel.IdentityVisitOfOptionalType(this, it, ctx)
	}

	// Make sure we're a local reference
	_, _, ok = typeName.PackageReference.TryGroupVersion()
	if !ok {
		return astmodel.IdentityVisitOfOptionalType(this, it, ctx)
	}

	resolved, err := f.definitions.FullyResolve(typeName)
	if err != nil {
		return nil, err
	}

	_, isArray := astmodel.AsArrayType(resolved)
	_, isMap := astmodel.AsMapType(resolved)
	isCollection := isArray || isMap

	if !isCollection {
		return astmodel.IdentityVisitOfOptionalType(this, it, ctx)
	}

	// Return the type unwrapped
	return it.Element(), nil
}
