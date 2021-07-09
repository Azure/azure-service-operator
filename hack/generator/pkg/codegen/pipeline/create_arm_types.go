/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// CreateARMTypes walks the type graph and builds new types for communicating
// with ARM
func CreateARMTypes(idFactory astmodel.IdentifierFactory) Stage {
	return MakeStage(
		"createArmTypes",
		"Create types for interaction with ARM",
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
		p = p.WithKubebuilderRequiredValidation(false)

		// remove all validation types by promoting inner type
		if validated, ok := p.PropertyType().(*astmodel.ValidatedType); ok {
			p = p.WithType(validated.ElementType())
		}

		t = t.WithProperty(p)
	}

	return t, nil
}

func removeFlattening(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	return removeFlattenFromObject(t), nil
}

func (c *armTypeCreator) createARMTypeDefinition(isSpecType bool, def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	convertObjectPropertiesForARM := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return c.convertObjectPropertiesForARM(t, isSpecType)
	}

	addOneOfConversionFunctionIfNeeded := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if astmodel.OneOfFlag.IsOn(def.Type()) {
			klog.V(4).Infof("Type %s is a OneOf type, adding MarshalJSON()", def.Name())
			return t.WithFunction(astmodel.NewOneOfJSONMarshalFunction(t, c.idFactory)), nil
		}

		return t, nil
	}

	armName := astmodel.CreateARMTypeName(def.Name())
	armDef, err := def.WithName(armName).ApplyObjectTransformations(removeValidations, convertObjectPropertiesForARM, addOneOfConversionFunctionIfNeeded, removeFlattening)
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
	createArmTypeName := func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		// Allow json type to pass through.
		if it == astmodel.JSONTypeName {
			return it, nil
		}

		def, ok := c.definitions[it]
		if !ok {
			return nil, errors.Errorf("failed to lookup %v", it)
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

	visitor := astmodel.TypeVisitorBuilder{
		VisitTypeName: createArmTypeName,
	}.Build()

	return visitor.Visit(t, nil)
}

// convertObjectPropertiesForARM returns the given object type with
// any properties updated that need to be changed for ARM
func (c *armTypeCreator) convertObjectPropertiesForARM(t *astmodel.ObjectType, isSpecType bool) (*astmodel.ObjectType, error) {
	result := t

	var errs []error
	for _, prop := range result.Properties() {
		if isSpecType && prop.HasName("Name") {
			// all resource Spec Name properties must be strings on their way to ARM
			// as nested resources will have the owner etc added to the start:
			result = result.WithProperty(prop.WithType(astmodel.StringType))
		} else if prop.PropertyType().Equals(astmodel.ResourceReferenceTypeName) || prop.PropertyType().Equals(astmodel.NewOptionalType(astmodel.ResourceReferenceTypeName)) {
			isRequired := prop.PropertyType().Equals(astmodel.ResourceReferenceTypeName)

			// Extract expected property name
			values, ok := prop.Tag(astmodel.ARMReferenceTag)
			if !ok {
				errs = append(errs, errors.Errorf("ResourceReference property missing %q tag", astmodel.ARMReferenceTag))
				continue
			}

			if len(values) != 1 {
				errs = append(errs, errors.Errorf("ResourceReference %q tag len(values) != 1", astmodel.ARMReferenceTag))
				continue
			}

			armPropName := values[0]

			// Remove the property we had since it doesn't apply to ARM
			result = result.WithoutProperty(prop.PropertyName())

			newProp := astmodel.NewPropertyDefinition(
				c.idFactory.CreatePropertyName(armPropName, astmodel.Exported),
				c.idFactory.CreateIdentifier(armPropName, astmodel.NotExported),
				astmodel.StringType)

			if isRequired {
				// We want to be required but don't need any kubebuidler annotations on this type because it's an ARM type
				newProp = newProp.MakeRequired().WithKubebuilderRequiredValidation(false)
			} else {
				newProp = newProp.MakeOptional()
			}

			result = result.WithProperty(newProp)
		} else {
			newType, err := c.convertARMPropertyTypeIfNeeded(prop.PropertyType())

			if err != nil {
				errs = append(errs, err)
			} else {
				result = result.WithProperty(prop.WithType(newType))
			}
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}
