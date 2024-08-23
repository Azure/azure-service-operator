/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/armconversion"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// CreateARMTypesStageID is the unique identifier for this pipeline stage
const CreateARMTypesStageID = "createArmTypes"

// CreateARMTypes walks the type graph and builds new types for communicating
// with ARM
func CreateARMTypes(
	configuration *config.ObjectModelConfiguration,
	idFactory astmodel.IdentifierFactory,
	log logr.Logger,
) *Stage {
	return NewStage(
		CreateARMTypesStageID,
		"Create types for interaction with ARM",
		func(ctx context.Context, state *State) (*State, error) {
			typeCreator := newARMTypeCreator(state.Definitions(), configuration, idFactory, log)
			armTypes, err := typeCreator.createARMTypes()
			if err != nil {
				return nil, err
			}

			newDefs := astmodel.TypesDisjointUnion(armTypes, state.Definitions())

			if err := configuration.PayloadType.VerifyConsumed(); err != nil {
				return nil, err
			}

			return state.WithDefinitions(newDefs), nil
		})
}

type skipError struct{}

func (s skipError) Error() string {
	return "skip"
}

var _ error = skipError{}

type armPropertyTypeConversionHandler func(
	prop *astmodel.PropertyDefinition,
	convContext *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error)

type armPropertyTypeConversionContext struct {
	isSpec   bool
	typeName astmodel.InternalTypeName
}

type armTypeCreator struct {
	definitions   astmodel.TypeDefinitionSet
	idFactory     astmodel.IdentifierFactory
	newDefs       astmodel.TypeDefinitionSet
	skipTypes     []func(it astmodel.TypeDefinition) bool
	log           logr.Logger
	visitor       astmodel.TypeVisitor[any]
	configuration *config.ObjectModelConfiguration
}

func newARMTypeCreator(
	definitions astmodel.TypeDefinitionSet,
	configuration *config.ObjectModelConfiguration,
	idFactory astmodel.IdentifierFactory,
	log logr.Logger,
) *armTypeCreator {
	result := &armTypeCreator{
		definitions: definitions,
		idFactory:   idFactory,
		newDefs:     make(astmodel.TypeDefinitionSet),
		skipTypes: []func(it astmodel.TypeDefinition) bool{
			skipUserAssignedIdentity,
		},
		log:           log,
		configuration: configuration,
	}

	result.visitor = astmodel.TypeVisitorBuilder[any]{
		VisitInternalTypeName: result.visitARMTypeName,
		VisitValidatedType:    result.visitARMValidatedType,
	}.Build()

	return result
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

	// Collect other object and enum definitions, as we need to create ARM variants of those too
	otherDefs := c.definitions.Except(resourceSpecDefs).
		Where(
			func(def astmodel.TypeDefinition) bool {
				_, isObject := astmodel.AsObjectType(def.Type())
				_, isEnum := astmodel.AsEnumType(def.Type())
				return isObject || isEnum
			})

	for name, def := range otherDefs {
		if !requiresARMType(def) {
			continue
		}

		convContext := c.createConversionContext(name)

		armDef, err := c.createARMTypeDefinition(def, convContext)
		if err != nil {
			return nil, err
		}

		result.Add(armDef)
	}

	result.AddTypes(c.newDefs)

	return result, nil
}

func (c *armTypeCreator) createARMResourceSpecDefinition(
	resource *astmodel.ResourceType,
	resourceSpecDef astmodel.TypeDefinition,
) (astmodel.TypeDefinition, error) {
	emptyDef := astmodel.TypeDefinition{}

	convContext := c.createSpecConversionContext(resourceSpecDef.Name())

	armTypeDef, err := c.createARMTypeDefinition(resourceSpecDef, convContext)
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
	for _, p := range t.Properties().Copy() {

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

func (c *armTypeCreator) createARMTypeDefinition(
	def astmodel.TypeDefinition,
	convContext *armPropertyTypeConversionContext,
) (astmodel.TypeDefinition, error) {
	if _, isObject := astmodel.AsObjectType(def.Type()); isObject {
		return c.createARMObjectTypeDefinition(def, convContext)
	}

	if _, isEnum := astmodel.AsEnumType(def.Type()); isEnum {
		return c.createARMEnumTypeDefinition(def, convContext)
	}

	debug := astmodel.DebugDescription(def.Type())
	return astmodel.TypeDefinition{}, errors.New("unsupported type: " + debug)
}

func (c *armTypeCreator) createARMObjectTypeDefinition(
	def astmodel.TypeDefinition,
	convContext *armPropertyTypeConversionContext,
) (astmodel.TypeDefinition, error) {
	convertObjectPropertiesForARM := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return c.convertObjectPropertiesForARM(t, convContext)
	}

	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())

	addOneOfConversionFunctionIfNeeded := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if isOneOf {
			c.log.V(1).Info(
				"Adding MarshalJSON and UnmarshalJSON to OneOf",
				"type", def.Name())
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

func (c *armTypeCreator) createARMEnumTypeDefinition(
	def astmodel.TypeDefinition,
	_ *armPropertyTypeConversionContext,
) (astmodel.TypeDefinition, error) {
	name := astmodel.CreateARMTypeName(def.Name())
	return def.WithName(name), nil
}

func (c *armTypeCreator) createARMTypeIfNeeded(t astmodel.Type) (astmodel.Type, error) {
	return c.visitor.Visit(t, nil)
}

func (c *armTypeCreator) createARMNameProperty(
	prop *astmodel.PropertyDefinition,
	convContext *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	if convContext.isSpec && prop.HasName("Name") {
		// all resource Spec Name properties must be strings on their way to ARM
		// as nested resources will have the owner etc. added to the start:
		return prop.WithType(astmodel.StringType), nil
	}

	return nil, nil
}

func (c *armTypeCreator) createUserAssignedIdentitiesProperty(
	prop *astmodel.PropertyDefinition,
	_ *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	typeName, ok := astmodel.IsUserAssignedIdentityProperty(prop)
	if !ok {
		return nil, nil
	}

	newTypeName := astmodel.CreateARMTypeName(typeName)
	// TODO: Currently the shape of the this type is always empty. This is safe now because the expected value type
	// TODO: of the map for all RPs is entirely readonly. If at some point in the future ARM allows users to pass
	// TODO: values to the RP in the value of this map, we will need to be more intelligent about how we construct
	// TODO: the map value type.
	// TODO: Given that evolution of the UserAssignedIdentities API is infrequent and may never happen, we don't currently
	// TODO: support that.
	newType := astmodel.ARMFlag.ApplyTo(astmodel.DoNotPrune.ApplyTo(astmodel.EmptyObjectType))
	newDef := astmodel.MakeTypeDefinition(newTypeName, newType).WithDescription(UserAssignedIdentityTypeDescription)

	newPropType := astmodel.NewMapType(astmodel.StringType, newTypeName)

	newProp := astmodel.NewPropertyDefinition(
		c.idFactory.CreatePropertyName(astmodel.UserAssignedIdentitiesProperty, astmodel.Exported),
		c.idFactory.CreateStringIdentifier(astmodel.UserAssignedIdentitiesProperty, astmodel.NotExported),
		newPropType).MakeTypeOptional()

	err := c.newDefs.AddAllowDuplicates(newDef)
	if err != nil {
		return nil, err
	}

	return newProp, nil
}

func (c *armTypeCreator) createResourceReferenceProperty(
	prop *astmodel.PropertyDefinition,
	_ *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	if !astmodel.IsTypeResourceReference(prop.PropertyType()) {
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

	var newPropType astmodel.Type
	if astmodel.IsTypeResourceReferenceSlice(prop.PropertyType()) {
		newPropType = astmodel.NewArrayType(astmodel.StringType)
	} else if astmodel.IsTypeResourceReferenceMap(prop.PropertyType()) {
		newPropType = astmodel.MapOfStringStringType
	} else {
		newPropType = astmodel.StringType
	}

	armPropName := values[0]
	newProp := astmodel.NewPropertyDefinition(
		c.idFactory.CreatePropertyName(armPropName, astmodel.Exported),
		c.idFactory.CreateStringIdentifier(armPropName, astmodel.NotExported),
		newPropType).MakeTypeOptional()

	return newProp, nil
}

func (c *armTypeCreator) createSecretReferenceProperty(
	prop *astmodel.PropertyDefinition,
	_ *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	if !astmodel.IsTypeSecretReference(prop.PropertyType()) {
		return nil, nil
	}

	var newType astmodel.Type
	if astmodel.IsTypeSecretReferenceSlice(prop.PropertyType()) {
		newType = astmodel.NewArrayType(astmodel.StringType)
	} else if astmodel.IsTypeSecretReferenceMap(prop.PropertyType()) {
		newType = astmodel.MapOfStringStringType
	} else if astmodel.TypeEquals(prop.PropertyType(), astmodel.OptionalSecretReferenceType) {
		newType = astmodel.OptionalStringType
	} else {
		newType = astmodel.StringType
	}

	return prop.WithType(newType), nil
}

func (c *armTypeCreator) createConfigMapReferenceProperty(
	prop *astmodel.PropertyDefinition,
	_ *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	if !astmodel.TypeEquals(prop.PropertyType(), astmodel.ConfigMapReferenceType) &&
		!astmodel.TypeEquals(prop.PropertyType(), astmodel.OptionalConfigMapReferenceType) {
		return nil, nil
	}

	if prop.HasTag(astmodel.OptionalConfigMapPairTag) {
		return nil, skipError{}
	}

	isRequired := astmodel.TypeEquals(prop.PropertyType(), astmodel.ConfigMapReferenceType)

	var newType astmodel.Type
	if isRequired {
		newType = astmodel.StringType
	} else {
		newType = astmodel.OptionalStringType
	}

	return prop.WithType(newType), nil
}

func (c *armTypeCreator) createARMProperty(
	prop *astmodel.PropertyDefinition,
	convContext *armPropertyTypeConversionContext,
) (*astmodel.PropertyDefinition, error) {
	newType, err := c.createARMTypeIfNeeded(prop.PropertyType())
	if err != nil {
		return nil, err
	}

	// Return a property with (potentially) a new type
	result := prop.WithType(newType)

	var payloadType config.PayloadType
	payloadType, ok := c.configuration.PayloadType.Lookup(convContext.typeName, prop.PropertyName())
	if !ok {
		payloadType = config.OmitEmptyProperties
	}

	switch payloadType {
	case config.OmitEmptyProperties:
		// NOP

	case config.ExplicitProperties:
		// With PayloadType 'explicit' we remove the `omitempty` tag, because we always want to send *all* ARM properties
		// to the server, even if they are empty.
		// See https://github.com/Azure/azure-service-operator/issues/2914 for the problem we're solving here.
		// See https://azure.github.io/azure-service-operator/design/adr-2023-04-patch-collections/ for how we're solving it.
		result = result.WithoutTag("json", "omitempty")

	case config.ExplicitEmptyCollections:
		fallthrough
	case config.ExplicitCollections:
		// With PayloadType 'explicitCollections' or 'ExplicitEmptyCollections' we remove the `omitempty`
		// tag from arrays and maps, because we always want to explicitly send collections to the server, even if empty.
		_, isMap := astmodel.AsMapType(newType)
		_, isArray := astmodel.AsArrayType(newType)
		if isMap || isArray {
			result = result.WithoutTag("json", "omitempty")
		}
	}

	return result, nil
}

// convertObjectPropertiesForARM returns the given object type with
// any properties updated that need to be changed for ARM
func (c *armTypeCreator) convertObjectPropertiesForARM(
	t *astmodel.ObjectType,
	convContext *armPropertyTypeConversionContext,
) (*astmodel.ObjectType, error) {
	propertyHandlers := []armPropertyTypeConversionHandler{
		c.createARMNameProperty,
		c.createUserAssignedIdentitiesProperty,
		c.createResourceReferenceProperty,
		c.createSecretReferenceProperty,
		c.createConfigMapReferenceProperty,
		c.createARMProperty,
	}

	result := t.WithoutProperties()
	var errs []error
	for _, prop := range t.Properties().Copy() {
		for _, handler := range propertyHandlers {
			newProp, err := handler(prop, convContext)
			if err != nil {
				if errors.As(err, &skipError{}) {
					break
				}
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
	result = result.WithoutEmbeddedProperties() // Clear them out first, so we're starting with a clean slate
	for _, prop := range t.EmbeddedProperties() {
		for _, handler := range embeddedPropertyHandlers {
			newProp, err := handler(prop, convContext)
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

func (c *armTypeCreator) visitARMTypeName(
	this *astmodel.TypeVisitor[any],
	it astmodel.InternalTypeName,
	ctx any,
) (astmodel.Type, error) {
	// Look up the definition
	def, err := c.definitions.GetDefinition(it)
	if err != nil {
		// Don't need to wrap, it already has everything we want in the error
		return nil, err
	}

	// If the name references an object type, we need an updated name, create it and return
	if _, ok := astmodel.AsObjectType(def.Type()); ok {
		return astmodel.CreateARMTypeName(def.Name()), nil
	}

	// If the name is an alias for a primitive type, we look through it to the underlying type,
	// allowing us to define the property using the primitive type directly
	// and avoiding any potential circular dependencies
	_, isPrimitiveType := astmodel.AsPrimitiveType(def.Type())
	_, isArrayType := astmodel.AsArrayType(def.Type())
	_, isMapType := astmodel.AsMapType(def.Type())
	if isPrimitiveType || isArrayType || isMapType {
		// We use the visitor to simplify the type
		var updatedType astmodel.Type
		updatedType, err = this.Visit(def.Type(), ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to look through definition %s", def.Name())
		}

		return updatedType, nil
	}

	// If the name references an enum type, we need an updated name, create it and return
	if _, ok := astmodel.AsEnumType(def.Type()); ok {
		return astmodel.CreateARMTypeName(def.Name()), nil
	}

	// We may or may not need to use an updated type name (i.e. if it's an aliased primitive type we can
	// just keep using that alias)
	updatedType, err := this.Visit(def.Type(), ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to update definition %s", def.Name())
	}

	// If no change, keep the existing name
	if astmodel.TypeEquals(updatedType, def.Type()) {
		return it, nil
	}

	// Otherwise give it a new name, and use that
	return astmodel.CreateARMTypeName(def.Name()), nil
}

func (c *armTypeCreator) visitARMValidatedType(
	this *astmodel.TypeVisitor[any],
	it *astmodel.ValidatedType,
	ctx any,
) (astmodel.Type, error) {
	return this.Visit(it.ElementType(), ctx)
}

func (c *armTypeCreator) createSpecConversionContext(name astmodel.InternalTypeName) *armPropertyTypeConversionContext {
	result := c.createConversionContext(name)
	result.isSpec = true
	return result
}

func (c *armTypeCreator) createConversionContext(name astmodel.InternalTypeName) *armPropertyTypeConversionContext {
	result := &armPropertyTypeConversionContext{
		typeName: name,
	}

	return result
}

var skipARMFuncs = []func(it astmodel.TypeDefinition) bool{
	skipUserAssignedIdentity,
}

func skipUserAssignedIdentity(def astmodel.TypeDefinition) bool {
	return def.Name().Name() == astmodel.UserAssignedIdentitiesTypeName
}

func requiresARMType(def astmodel.TypeDefinition) bool {
	for _, f := range skipARMFuncs {
		skip := f(def)
		if skip {
			return false
		}
	}

	return true
}
