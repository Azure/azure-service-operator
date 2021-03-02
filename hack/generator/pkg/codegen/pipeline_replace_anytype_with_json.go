/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
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

func replaceAnyTypeWithJSON() PipelineStage {
	return MakePipelineStage(
		"replaceAnyTypeWithJSON",
		"Replacing interface{}s with arbitrary JSON",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			visitor := astmodel.MakeTypeVisitor()

			visitor.VisitPrimitive = func(_ *astmodel.TypeVisitor, it *astmodel.PrimitiveType, _ interface{}) (astmodel.Type, error) {
				if it == astmodel.AnyType {
					return astmodel.JSONType, nil
				}
				return it, nil
			}

			originalVisitMapType := visitor.VisitMapType
			visitor.VisitMapType = func(v *astmodel.TypeVisitor, it *astmodel.MapType, ctx interface{}) (astmodel.Type, error) {
				if it.Equals(mapOfMapOfAnyType) {
					return mapOfJSON, nil
				}
				return originalVisitMapType(v, it, ctx)
			}

			results := make(astmodel.Types)
			for _, def := range types {
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
