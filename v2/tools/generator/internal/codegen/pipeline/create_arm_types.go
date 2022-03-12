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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/armconversion"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// CreateARMTypesStageID is the unique identifier for this pipeline stage
const CreateARMTypesStageID = "createArmTypes"

// CreateARMTypes walks the type graph and builds new types for communicating
// with ARM
func CreateARMTypes(idFactory astmodel.IdentifierFactory) *Stage {
	return NewLegacyStage(
		CreateARMTypesStageID,
		"Create types for interaction with ARM",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			armTypeCreator := &armTypeCreator{definitions: definitions, idFactory: idFactory}
			armTypes, err := armTypeCreator.createARMTypes()
			if err != nil {
				return nil, err
			}

			result := astmodel.TypesDisjointUnion(armTypes, definitions)

			return result, nil
		})
}

type armPropertyTypeConversionHandler func(prop *astmodel.PropertyDefinition, isSpec bool) (*astmodel.PropertyDefinition, error)

type armTypeCreator struct {
	definitions astmodel.TypeDefinitionSet
	idFactory   astmodel.IdentifierFactory
}

func (c *armTypeCreator) createARMTypes() (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)
	resourceSpecDefs := make(astmodel.TypeDefinitionSet)

	resourceDefs := astmodel.FindResourceDefinitions(c.definitions)

	for _, def := range resourceDefs {
		resolved, err := c.definitions.ResolveResourceSpecAndStatus(def)
		if err != nil {
			return nil, errors.Wrapf(err, "resolving resource spec and status for %s", def.Name())
		}

		resourceSpecDefs.Add(resolved.SpecDef)

		armSpecDef, err := c.createARMResourceSpecDefinition(resolved.ResourceType, resolved.SpecDef)
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
	resource *astmodel.ResourceType,
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

	iface, err := armconversion.NewARMSpecInterfaceImpl(c.idFactory, resource, specObj)
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
		p = p.MakeOptional()

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

	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())

	addOneOfConversionFunctionIfNeeded := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if isOneOf {
			klog.V(4).Infof("Type %s is a OneOf type, adding MarshalJSON and UnmarshalJSON", def.Name())
			marshal := functions.NewOneOfJSONMarshalFunction(t, c.idFactory)
			unmarshal := functions.NewOneOfJSONUnmarshalFunction(t, c.idFactory)
			return t.WithFunction(marshal).WithFunction(unmarshal), nil
		}

		return t, nil
	}

	armName := astmodel.CreateARMTypeName(def.Name())
	armDef, err := def.WithName(armName).ApplyObjectTransformations(
		removeValidations,
		convertObjectPropertiesForARM,
		addOneOfConversionFunctionIfNeeded,
		removeFlattening)
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM prototype %s from Kubernetes definition %s", armName, def.Name())
	}

	result, err := armDef.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
		return astmodel.ARMFlag.ApplyTo(objectType), nil
	})
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM definition %s from Kubernetes definition %s", armName, def.Name())
	}

	// copy OneOf flag over to ARM type, if applicable
	// this is needed so the gopter generators can tell if the type should be OneOf,
	// see json_serialization_test_cases.go: AddTestTo
	if isOneOf {
		result = result.WithType(astmodel.OneOfFlag.ApplyTo(result.Type()))
	}

	return result, nil
}

func (c *armTypeCreator) createARMTypeIfNeeded(t astmodel.Type) (astmodel.Type, error) {
	createArmTypeName := func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		// Allow json type to pass through.
		if it == astmodel.JSONType {
			return it, nil
		}

		def, ok := c.definitions[it]
		if !ok {
			return nil, errors.Errorf("failed to lookup %s", it)
		}

		if _, ok := def.Type().(*astmodel.ObjectType); ok {
			return astmodel.CreateARMTypeName(def.Name()), nil
		}

		// We may or may not need to use an updated type name (i.e. if it's an aliased primitive type we can
		// just keep using that alias)
		updatedType, err := this.Visit(def.Type(), ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to update definition %s", def.Name())
		}

		if astmodel.TypeEquals(updatedType, def.Type()) {
			return it, nil
		}

		return astmodel.CreateARMTypeName(def.Name()), nil
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitTypeName: createArmTypeName,
	}.Build()

	return visitor.Visit(t, nil)
}

func (c *armTypeCreator) createARMNameProperty(prop *astmodel.PropertyDefinition, isSpec bool) (*astmodel.PropertyDefinition, error) {
	if isSpec && prop.HasName("Name") {
		// all resource Spec Name properties must be strings on their way to ARM
		// as nested resources will have the owner etc added to the start:
		return prop.WithType(astmodel.StringType), nil
	}

	return nil, nil
}

func (c *armTypeCreator) createResourceReferenceProperty(prop *astmodel.PropertyDefinition, _ bool) (*astmodel.PropertyDefinition, error) {
	if !astmodel.TypeEquals(prop.PropertyType(), astmodel.ResourceReferenceType) &&
		!astmodel.TypeEquals(prop.PropertyType(), astmodel.NewOptionalType(astmodel.ResourceReferenceType)) {
		return nil, nil
	}

	// Extract expected property name
	values, ok := prop.Tag(astmodel.ARMReferenceTag)
	if !ok {
		return nil, errors.Errorf("ResourceReference property missing %q tag", astmodel.ARMReferenceTag)
	}

	if len(values) != 1 {
		return nil, errors.Errorf("ResourceReference %q tag len(values) != 1", astmodel.ARMReferenceTag)
	}

	armPropName := values[0]
	newProp := astmodel.NewPropertyDefinition(
		c.idFactory.CreatePropertyName(armPropName, astmodel.Exported),
		c.idFactory.CreateIdentifier(armPropName, astmodel.NotExported),
		astmodel.StringType).MakeTypeOptional()

	return newProp, nil
}

func (c *armTypeCreator) createSecretReferenceProperty(prop *astmodel.PropertyDefinition, _ bool) (*astmodel.PropertyDefinition, error) {
	if !astmodel.TypeEquals(prop.PropertyType(), astmodel.SecretReferenceType) &&
		!astmodel.TypeEquals(prop.PropertyType(), astmodel.NewOptionalType(astmodel.SecretReferenceType)) {
		return nil, nil
	}

	isRequired := astmodel.TypeEquals(prop.PropertyType(), astmodel.SecretReferenceType)

	var newType astmodel.Type
	if isRequired {
		newType = astmodel.StringType
	} else {
		newType = astmodel.NewOptionalType(astmodel.StringType)
	}

	return prop.WithType(newType), nil
}

func (c *armTypeCreator) createARMProperty(prop *astmodel.PropertyDefinition, _ bool) (*astmodel.PropertyDefinition, error) {
	newType, err := c.createARMTypeIfNeeded(prop.PropertyType())

	if err != nil {
		return nil, err
	}
	return prop.WithType(newType), nil
}

// convertObjectPropertiesForARM returns the given object type with
// any properties updated that need to be changed for ARM
func (c *armTypeCreator) convertObjectPropertiesForARM(t *astmodel.ObjectType, isSpecType bool) (*astmodel.ObjectType, error) {
	propertyHandlers := []armPropertyTypeConversionHandler{
		c.createARMNameProperty,
		c.createResourceReferenceProperty,
		c.createSecretReferenceProperty,
		c.createARMProperty,
	}

	result := t.WithoutProperties()
	var errs []error
	for _, prop := range t.Properties() {
		for _, handler := range propertyHandlers {
			newProp, err := handler(prop, isSpecType)
			if err != nil {
				errs = append(errs, err)
				break // Stop calling handlers and proceed to the next property
			}

			if newProp != nil {
				result = result.WithProperty(newProp)
				// Once we've matched a handler, stop looking for more
				break
			}
		}
	}

	embeddedPropertyHandlers := []armPropertyTypeConversionHandler{
		c.createARMProperty,
	}

	// Also convert embedded properties if there are any
	result = result.WithoutEmbeddedProperties() // Clear them out first so we're starting with a clean slate
	for _, prop := range t.EmbeddedProperties() {
		for _, handler := range embeddedPropertyHandlers {
			newProp, err := handler(prop, isSpecType)
			if err != nil {
				errs = append(errs, err)
				break // Stop calling handlers and proceed to the next property
			}

			if newProp != nil {
				result, err = result.WithEmbeddedProperty(newProp)
				if err != nil {
					errs = append(errs, err)
				}
				// Once we've matched a handler, stop looking for more
				break
			}
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}
