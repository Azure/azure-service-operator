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

// jsonType is the type of fields storing arbitrary JSON content in
// custom resources - apiextensions/v1.JSON.
var (
	// replace map[string]map[string]interface{} with maps of JSON to
	// work around the controller-gen limitation that it barfs in maps
	// of maps of things.
	// TODO: remove this when it can handle them correctly.
	mapOfMapOfAnyType = astmodel.NewMapType(
		astmodel.StringType,
		astmodel.NewMapType(
			astmodel.StringType,
			astmodel.AnyType,
		),
	)
	mapOfJSON = astmodel.NewMapType(astmodel.StringType, astmodel.JSONType)
)

func ReplaceAnyTypeWithJSON() *Stage {
	return NewLegacyStage(
		"replaceAnyTypeWithJSON",
		"Replace properties using interface{} with arbitrary JSON",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			replaceAnyWithJson := func(it *astmodel.PrimitiveType) astmodel.Type {
				if it == astmodel.AnyType {
					return astmodel.JSONType
				}

				return it
			}

			replaceMapOfMapOfAnyWithJSON := func(v *astmodel.TypeVisitor[any], it *astmodel.MapType, ctx interface{}) (astmodel.Type, error) {
				if astmodel.TypeEquals(it, mapOfMapOfAnyType) {
					return mapOfJSON, nil
				}
				return astmodel.IdentityVisitOfMapType(v, it, ctx)
			}

			visitor := astmodel.TypeVisitorBuilder[any]{
				VisitPrimitive: replaceAnyWithJson,
				VisitMapType:   replaceMapOfMapOfAnyWithJSON,
			}.Build()

			results := make(astmodel.TypeDefinitionSet)
			for _, def := range definitions {
				d, err := visitor.VisitDefinition(def, nil)
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}
				results.Add(d)
			}

			return results, nil
		},
	)
}
