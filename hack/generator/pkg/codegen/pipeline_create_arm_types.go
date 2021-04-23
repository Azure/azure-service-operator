/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// createARMTypes walks the type graph and builds new types for communicating
// with ARM
func createARMTypes(idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"createArmTypes",
		"Creates ARM types",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {

			armTypeCreator := &armTypeCreator{definitions: definitions, idFactory: idFactory}
			armTypes, err := armTypeCreator.createARMTypes()

			if err != nil {
				return nil, err
			}

			result := astmodel.TypesDisjointUnion(armTypes, definitions)

			return result, nil
		})
}

type armTypeCreator struct {
	definitions astmodel.Types
	idFactory   astmodel.IdentifierFactory
}

func getAllSpecDefinitions(definitions astmodel.Types) (astmodel.Types, error) {
	result := make(astmodel.Types)
	for _, def := range definitions {
		if resourceType, ok := astmodel.AsResourceType(def.Type()); ok {
			spec, err := definitions.ResolveResourceSpecDefinition(resourceType)
			if err != nil {
				return nil, err
			}
			result.Add(spec)
		}
	}

	return result, nil
}

func (c *armTypeCreator) createARMTypes() (astmodel.Types, error) {

	result := make(astmodel.Types)
	resourceSpecDefs, err := getAllSpecDefinitions(c.definitions)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get all resource spec definitions")
	}

	for _, def := range resourceSpecDefs {
		armSpecDef, err := c.createARMResourceSpecDefinition(def)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create arm resource spec definition for resource %s", def.Name())
		}

		result.Add(armSpecDef)
	}

	otherDefs := c.definitions.Except(resourceSpecDefs).Where(func(def astmodel.TypeDefinition) bool {
		_, ok := astmodel.AsObjectType(def.Type())
		return ok
	})

	for _, def := range otherDefs {
		armDef, err := c.createARMTypeDefinition(
			false, // not Spec type
			def)
		if err != nil {
			return nil, err
		}

		result.Add(armDef)
	}

	return result, nil
}

func (c *armTypeCreator) createARMResourceSpecDefinition(
	resourceSpecDef astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {

	emptyDef := astmodel.TypeDefinition{}

	armTypeDef, err := c.createARMTypeDefinition(true, resourceSpecDef)
	if err != nil {
		return emptyDef, err
	}

	// ARM specs have a special interface that they need to implement, go ahead and create that here
	if !astmodel.ARMFlag.IsOn(armTypeDef.Type()) {
		return emptyDef, errors.Errorf("arm spec %q isn't a flagged object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	// Safe because above test passed
	flagged := armTypeDef.Type().(*astmodel.FlaggedType)

	specObj, ok := flagged.Element().(*astmodel.ObjectType)
	if !ok {
		return emptyDef, errors.Errorf("arm spec %q isn't an object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	iface, err := astmodel.NewARMSpecInterfaceImpl(c.idFactory, specObj)
	if err != nil {
		return emptyDef, err
	}

	updatedSpec := specObj.WithInterface(iface)
	armTypeDef = armTypeDef.WithType(astmodel.ARMFlag.ApplyTo(updatedSpec))

	return armTypeDef, nil
}

func removeValidations(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	for _, p := range t.Properties() {

		// set all properties as not-required
		p = p.SetRequired(false)

		// remove all validation types by promoting inner type
		if validated, ok := p.PropertyType().(*astmodel.ValidatedType); ok {
			p = p.WithType(validated.ElementType())
		}

		t = t.WithProperty(p)
	}

	return t, nil
}

func (c *armTypeCreator) createARMTypeDefinition(isSpecType bool, def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	convertPropertiesToARMTypesWrapper := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return c.convertPropertiesToARMTypes(t, isSpecType)
	}

	addOneOfConversionFunctionIfNeeded := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if astmodel.OneOfFlag.IsOn(def.Type()) {
			klog.V(4).Infof("Type %s is a OneOf type, adding MarshalJSON()", def.Name())
			return t.WithFunction(astmodel.NewOneOfJSONMarshalFunction(t, c.idFactory)), nil
		}

		return t, nil
	}

	armName := astmodel.CreateARMTypeName(def.Name())
	armDef, err := def.WithName(armName).ApplyObjectTransformations(removeValidations, convertPropertiesToARMTypesWrapper, addOneOfConversionFunctionIfNeeded)
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM prototype %v from Kubernetes definition %v", armName, def.Name())
	}

	result, err := armDef.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
		return astmodel.ARMFlag.ApplyTo(objectType), nil
	})
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM definition %v from Kubernetes definition %v", armName, def.Name())
	}

	return result, nil
}

func (c *armTypeCreator) convertARMPropertyTypeIfNeeded(t astmodel.Type) (astmodel.Type, error) {

	visitor := astmodel.MakeTypeVisitor()
	visitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		// Allow json type to pass through.
		if it == astmodel.JSONType {
			return it, nil
		}

		def, ok := c.definitions[it]
		if !ok {
			return nil, errors.Errorf("Failed to lookup %v", it)
		}

		if _, ok := def.Type().(*astmodel.ObjectType); ok {
			return astmodel.CreateARMTypeName(def.Name()), nil
		}

		// We may or may not need to use an updated type name (i.e. if it's an aliased primitive type we can
		// just keep using that alias)
		updatedType, err := this.Visit(def.Type(), ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to update definition %v", def.Name())
		}

		if updatedType.Equals(def.Type()) {
			return it, nil
		}

		return astmodel.CreateARMTypeName(def.Name()), nil
	}

	return visitor.Visit(t, nil)
}

func (c *armTypeCreator) convertPropertiesToARMTypes(t *astmodel.ObjectType, isSpecType bool) (*astmodel.ObjectType, error) {
	result := t

	var errs []error
	for _, prop := range result.Properties() {
		if isSpecType && prop.HasName("Name") {
			// all resource Spec Name properties must be strings on their way to ARM
			// as nested resources will have the owner etc added to the start:
			result = result.WithProperty(prop.WithType(astmodel.StringType))
		} else {
			propType := prop.PropertyType()
			newType, err := c.convertARMPropertyTypeIfNeeded(propType)
			if err != nil {
				errs = append(errs, err)
			} else if newType != propType {
				newProp := prop.WithType(newType)
				result = result.WithProperty(newProp)
			}
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}
