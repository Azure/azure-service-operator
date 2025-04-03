/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const promoteRootOneOfProperties = "promoteRootOneOfProperties"

func PromoteRootOneOfProperties() *Stage {
	stage := NewStage(
		promoteRootOneOfProperties,
		"Promote selected properties on OneOf objects directly referenced by resources",
		func(ctx context.Context, state *State) (*State, error) {
			promoter := newRootPropertyPromoter(state.Definitions())
			newDefs, err := promoter.modifyDefinitions()
			if err != nil {
				return nil, eris.Wrapf(err, "assembling OneOf types")
			}

			return state.WithOverlaidDefinitions(newDefs), nil
		})

	return stage
}

type rootPropertyPromoter struct {
	definitions          astmodel.TypeDefinitionSet
	promotableProperties set.Set[string]
	rootVisitor          astmodel.TypeVisitor[astmodel.TypeDefinitionSet]
	leafVisitor          astmodel.TypeVisitor[astmodel.PropertySet]
}

func newRootPropertyPromoter(
	defs astmodel.TypeDefinitionSet,
) *rootPropertyPromoter {
	result := &rootPropertyPromoter{
		definitions:          defs,             // All known definitions
		promotableProperties: set.Make("Name"), // Properties we want to promote (hard coded for now)
	}

	result.rootVisitor = astmodel.TypeVisitorBuilder[astmodel.TypeDefinitionSet]{
		VisitObjectType: result.visitRootObjectType,
	}.Build()

	result.leafVisitor = astmodel.TypeVisitorBuilder[astmodel.PropertySet]{
		VisitObjectType: result.visitLeafObjectType,
	}.Build()

	return result
}

func (p *rootPropertyPromoter) modifyDefinitions() (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	// Find any OneOf types that are directly referenced by resources
	var errs []error
	for _, r := range p.definitions.AllResources() {
		info, err := p.definitions.ResolveResourceSpecAndStatus(r)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// If our spec is an object made from a OneOf, visit it and the associated leaves
		if astmodel.OneOfFlag.IsOn(info.SpecDef.Type()) {
			def, err := p.rootVisitor.VisitDefinition(info.SpecDef, result)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			result.Add(def)
		}

		// Ditto for the status
		if astmodel.OneOfFlag.IsOn(info.StatusDef.Type()) {
			def, err := p.rootVisitor.VisitDefinition(info.StatusDef, result)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			result.Add(def)
		}
	}

	if len(errs) > 0 {
		return nil, eris.Wrap(kerrors.NewAggregate(errs), "failed to promote properties")
	}

	return result, nil
}

func (p *rootPropertyPromoter) visitRootObjectType(
	_ *astmodel.TypeVisitor[astmodel.TypeDefinitionSet],
	ot *astmodel.ObjectType,
	modifiedDefs astmodel.TypeDefinitionSet,
) (astmodel.Type, error) {
	// First iterate through our leaves and remove any properties that need promotion
	propertiesRemoved := astmodel.NewPropertySet()
	var errs []error
	for _, prop := range ot.Properties().AsSlice() {
		if tn, ok := astmodel.AsInternalTypeName(prop.PropertyType()); ok {
			def, ok := p.definitions[tn]
			if !ok {
				errs = append(errs, eris.Errorf("could not find definition for %q", tn))
			}

			newDef, err := p.leafVisitor.VisitDefinition(def, propertiesRemoved)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			modifiedDefs.Add(newDef)
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	// Now inject any properties we removed from leaves into the root object
	resultType := ot.WithProperties(propertiesRemoved.AsSlice()...)
	return resultType, nil
}

func (p *rootPropertyPromoter) visitLeafObjectType(
	_ *astmodel.TypeVisitor[astmodel.PropertySet],
	ot *astmodel.ObjectType,
	propertiesRemoved astmodel.PropertySet,
) (astmodel.Type, error) {
	for n := range p.promotableProperties {
		propertyName := astmodel.PropertyName(n)
		prop, ok := ot.Property(propertyName)
		if ok {
			propertiesRemoved.Add(prop)
			ot = ot.WithoutProperty(propertyName)
		}
	}

	return ot, nil
}
