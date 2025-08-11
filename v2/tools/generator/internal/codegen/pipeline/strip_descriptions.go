/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// StripDocumentationStageID is the unique identifier for this pipeline stage
const StripDocumentationStageID = "stripDocumentation"

// StripDocumentation strips property descriptions from certain types to
func StripDocumentation(configuration *config.Configuration, log logr.Logger) *Stage {
	stage := NewStage(
		StripDocumentationStageID,
		"Strip descriptions for CRDs that have the $stripDocumentation flag set",
		func(ctx context.Context, state *State) (*State, error) {
			visitor := createPropertyDescriptionRemovalVisitor()

			walker := astmodel.NewTypeWalker(
				state.Definitions(),
				visitor)
			walker.AfterVisit = stripTypeDefinitionDescription

			// TODO: This should be used sparingly as a stop-gap if a CRD gets too large.
			//typesToStrip := map[string]set.Set[string]{
			//	"containerservice": set.Make[string]("ManagedCluster"),
			//}

			result := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions().AllResources() {
				group := def.Name().InternalPackageReference().Group()
				name := def.Name().Name()

				shouldStrip, ok := configuration.ObjectModelConfiguration.StripDocumentation.Lookup(def.Name())
				if !ok || !shouldStrip {
					continue
				}
				//types, ok := typesToStrip[group]
				//if !ok {
				//	continue
				//}
				//
				//if !types.Contains(name) {
				//	continue
				//}

				log.V(1).Info("Stripping property descriptions from %s.%s", group, name)
				modifiedDefs, err := walker.Walk(def)
				if err != nil {
					return nil, err
				}
				result.AddTypes(modifiedDefs)
			}

			err := configuration.ObjectModelConfiguration.StripDocumentation.VerifyConsumed()
			if err != nil {
				return nil, eris.Wrap(
					err,
					"Found unused $stripDocumentation configurations; these can only be specified on top-level resources.")
			}

			return state.WithOverlaidDefinitions(result), nil
		})

	return stage
}

func stripTypeDefinitionDescription(_ astmodel.TypeDefinition, updated astmodel.TypeDefinition, _ any) (astmodel.TypeDefinition, error) {
	updated = updated.WithDescription("")
	return updated, nil
}

func createPropertyDescriptionRemovalVisitor() astmodel.TypeVisitor[any] {
	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: func(
			this *astmodel.TypeVisitor[any],
			it *astmodel.ObjectType,
			ctx any,
		) (astmodel.Type, error) {
			result := it
			for _, prop := range it.Properties().AsSlice() {
				// If the property already has an empty description, we're done
				if prop.Description() == "" {
					continue
				}

				// Set the description to empty
				prop = prop.WithDescription("")
				result = result.WithProperty(prop)
			}

			return astmodel.OrderedIdentityVisitOfObjectType(this, result, ctx)
		},
	}

	return visitor.Build()
}
