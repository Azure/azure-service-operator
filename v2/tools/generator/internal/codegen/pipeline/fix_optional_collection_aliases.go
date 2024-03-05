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
			fixer.visitor = astmodel.TypeVisitorBuilder[any]{
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
	visitor     astmodel.TypeVisitor[any]
}

func (f *optionalCollectionAliasFixer) fixOptionalCollectionAliases(
	this *astmodel.TypeVisitor[any],
	it *astmodel.OptionalType,
	ctx any,
) (astmodel.Type, error) {
	typeName, ok := astmodel.AsInternalTypeName(it)
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
