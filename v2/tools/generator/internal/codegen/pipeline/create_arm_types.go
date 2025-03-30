/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
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
	armDefs       astmodel.TypeDefinitionSet // Our set of new ARM types
	convertedDefs astmodel.TypeDefinitionSet // The definitions we've already converted
	skipTypes     []func(it astmodel.TypeDefinition) bool
	log           logr.Logger
	visitor       astmodel.TypeVisitor[any]
	configuration *config.ObjectModelConfiguration

	// leafProperties is a set of properties found on the root of a one-of type that need to be
	// pushed out to the leaves of that one-of, keyed by the leaf type that needs to be modified.
	leafProperties map[astmodel.InternalTypeName]astmodel.PropertySet
}

func newARMTypeCreator(
	definitions astmodel.TypeDefinitionSet,
	configuration *config.ObjectModelConfiguration,
	idFactory astmodel.IdentifierFactory,
	log logr.Logger,
) *armTypeCreator {
	result := &armTypeCreator{
		definitions:   definitions,
		idFactory:     idFactory,
		armDefs:       make(astmodel.TypeDefinitionSet),
		convertedDefs: make(astmodel.TypeDefinitionSet),
		skipTypes: []func(it astmodel.TypeDefinition) bool{
			skipUserAssignedIdentity,
		},
		log:            log,
		configuration:  configuration,
		leafProperties: make(map[astmodel.InternalTypeName]astmodel.PropertySet),
	}

	result.visitor = astmodel.TypeVisitorBuilder[any]{
		VisitInternalTypeName: result.visitARMTypeName,
		VisitValidatedType:    result.visitARMValidatedType,
	}.Build()

	return result
}

func (c *armTypeCreator) createARMTypes() (astmodel.TypeDefinitionSet, error) {
	// Create ARM variants for all Spec and Status types
	// For OneOf types to end up correctly shaped, we must do both spec and status before we do any
	// other types, so that we discover any properties needing to be pushed from root to leaf
	// before we create the leaf types.
	for _, def := range c.definitions.AllResources() {
		err := c.createARMTypesForResource(def)
		if err != nil {
			// No need to wrap as it already identifies the resource
			return nil, err
		}
	}

	// Collect other object and enum definitions, as we may need to create ARM variants of those too
	otherDefs := c.definitions.Except(c.convertedDefs).
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

		c.addARMType(def, armDef)
	}

	return c.armDefs, nil
}

// createARMTypesForResource creates ARM types for the spec and status of the resource type.
func (c *armTypeCreator) createARMTypesForResource(
	rsrc astmodel.TypeDefinition,
) error {
	resolved, err := c.definitions.ResolveResourceSpecAndStatus(rsrc)
	if err != nil {
		return eris.Wrapf(
			err,
			"resolving resource spec and status for %s",
			rsrc.Name())
	}

	// Create ARM type for the Spec
	spec, err := c.createARMResourceSpecDefinition(resolved.ResourceDef, resolved.SpecDef)
	if err != nil {
		return eris.Wrapf(
			err,
			"unable to create arm resource spec definition for resource %s",
			rsrc.Name())
	}

	c.addARMType(resolved.SpecDef, spec)

	// Create ARM type for the Status, if we have one
	if resolved.ResourceType.StatusType() != nil {
		convContext := c.createConversionContext(resolved.StatusDef.Name())
		status, err := c.createARMTypeDefinition(resolved.StatusDef, convContext)
		if err != nil {
			return eris.Wrapf(
				err,
				"unable to create arm resource status definition for resource %s",
				rsrc.Name())
		}

		c.addARMType(resolved.StatusDef, status)
	}

	return nil
}

func (c *armTypeCreator) createARMResourceSpecDefinition(
	rsrcDef astmodel.TypeDefinition,
	specDef astmodel.TypeDefinition,
) (astmodel.TypeDefinition, error) {
	resource, ok := astmodel.AsResourceType(rsrcDef.Type())
	if !ok {
		return astmodel.TypeDefinition{},
			eris.Errorf(
				"expected resource %s to be a resource type, but got %s",
				rsrcDef.Name(),
				astmodel.DebugDescription(rsrcDef.Type(), rsrcDef.Name().InternalPackageReference()))
	}

	emptyDef := astmodel.TypeDefinition{}

	convContext := c.createSpecConversionContext(specDef.Name())

	armTypeDef, err := c.createARMTypeDefinition(specDef, convContext)
	if err != nil {
		return emptyDef, err
	}

	// ARM specs have a special interface that they need to implement, go ahead and create that here
	flagged, isFlagged := astmodel.AsFlaggedType(armTypeDef.Type())
	if !isFlagged || !flagged.HasFlag(astmodel.ARMFlag) {
		return emptyDef, eris.Errorf("arm spec %q isn't a flagged object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	specObj, ok := astmodel.AsObjectType(flagged.Element())
	if !ok {
		return emptyDef, eris.Errorf("arm spec %q isn't an object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	iface, err := armconversion.NewARMSpecInterfaceImpl(c.idFactory, resource, specObj)
	if err != nil {
		return emptyDef, eris.Wrapf(err, "creating ARM spec interface for %s", armTypeDef.Name())
	}

	injector := astmodel.NewInterfaceInjector()
	armTypeDef, err = injector.Inject(armTypeDef, iface)
	if err != nil {
		return emptyDef, eris.Wrapf(err, "injecting ARM spec interface into %s", armTypeDef.Name())
	}

	return armTypeDef, nil
}

func (c *armTypeCreator) removeValidations(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
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
	return astmodel.TypeDefinition{}, eris.New("unsupported type: " + debug)
}

func (c *armTypeCreator) createARMObjectTypeDefinition(
	def astmodel.TypeDefinition,
	convContext *armPropertyTypeConversionContext,
) (astmodel.TypeDefinition, error) {
	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())

	armName := astmodel.CreateARMTypeName(def.Name())
	armDef, err := def.WithName(armName).ApplyObjectTransformations(
		c.removeValidations,
		c.convertObjectPropertiesForARM(convContext),
		c.findOneOfLeafPropertiesStillOnRoot(def),
		c.addOneOfConversionFunctionIfNeeded(def),
		c.addLeafOneOfPropertiesIfNeeded(def),
		removeFlattening)
	if err != nil {
		return astmodel.TypeDefinition{},
			eris.Wrapf(err, "creating ARM prototype %s from Kubernetes definition %s", armName, def.Name())
	}

	result, err := armDef.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
		return astmodel.ARMFlag.ApplyTo(objectType), nil
	})
	if err != nil {
		return astmodel.TypeDefinition{},
			eris.Wrapf(err, "creating ARM definition %s from Kubernetes definition %s", armName, def.Name())
	}

	// copy OneOf flag over to ARM type, if applicable
	// this is needed so the gopter generators can tell if the type should be OneOf,
	// see json_serialization_test_cases.go: AddTestTo
	if isOneOf {
		result = result.WithType(astmodel.OneOfFlag.ApplyTo(result.Type()))
	}

	return result, nil
}

// findOneOfLeafPropertiesStillOnRoot verifies the provided TypeDefinition is a one-of type, and
// then scans the provided TypeDefinition to find any properties that are NOT references to one-of
// the leaf options.
// Normally properties are pushed all the way to leaves, but in rare cases we might have a property
// kept on the root that needs to be manually propagated to the leaves at runtime in order to
// correctly construct the ARM payload.
// For example, with a Kusto ClusterDatabase object, we still have `Name` on the root, but need to
// push that down to either `ReadWriteDatabase` or `ReadOnlyFollowingDatabase` when constructing
// the PUT payload for ARM.
func (c *armTypeCreator) findOneOfLeafPropertiesStillOnRoot(
	def astmodel.TypeDefinition,
) func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())
	if !isOneOf {
		return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
			return t, nil
		}
	}

	return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		// Iterate through all the available properties and find any that are not references
		// to OneOf leaf objects. We keep those properties to one side, as we need to inject
		// them on the leaves themselves later
		result := t
		rootProperties := astmodel.NewPropertySet()
		leafTypes := astmodel.NewInternalTypeNameSet()
		for _, p := range result.Properties().Copy() {
			if tn, ok := astmodel.AsInternalTypeName(p.PropertyType()); ok {
				leafTypes.Add(tn)
			} else {
				rootProperties.Add(p.WithTag(armconversion.ConversionTag, armconversion.NoARMConversionValue))
				// Put a flag on the root property so we can recognize it later
				result = result.WithProperty(
					p.WithTag(armconversion.ConversionTag, armconversion.PushToOneOfLeaf),
				)
			}
		}

		// If we have any root properties to push to the leaves, store them for later
		if len(rootProperties) > 0 {
			for leafType := range leafTypes {
				if known, ok := c.leafProperties[leafType]; ok {
					known.AddSet(rootProperties)
				} else {
					c.leafProperties[leafType] = rootProperties.Copy()
				}
			}
		}

		return result, nil
	}
}

// addLeafOneOfPropertiesIfNeeded adds any properties that were pushed from the root to the leaf
func (c *armTypeCreator) addLeafOneOfPropertiesIfNeeded(
	def astmodel.TypeDefinition,
) func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	armName := astmodel.CreateARMTypeName(def.Name())
	properties, ok := c.leafProperties[armName]
	if !ok {
		// No extra properties to inject
		return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
			return t, nil
		}
	}

	return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return t.WithProperties(properties.AsSlice()...), nil
	}
}

func (c *armTypeCreator) addOneOfConversionFunctionIfNeeded(
	def astmodel.TypeDefinition,
) func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	isOneOf := astmodel.OneOfFlag.IsOn(def.Type())
	return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
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

	err := c.armDefs.AddAllowDuplicates(newDef)
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
		return nil, eris.Errorf("ResourceReference property missing %q tag", astmodel.ARMReferenceTag)
	}

	if len(values) != 1 {
		return nil, eris.Errorf("ResourceReference %q tag len(values) != 1", astmodel.ARMReferenceTag)
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
	convContext *armPropertyTypeConversionContext,
) func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	return func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
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
					if eris.As(err, &skipError{}) {
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
			return nil, eris.Wrapf(err, "failed to look through definition %s", def.Name())
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
		return nil, eris.Wrapf(err, "failed to update definition %s", def.Name())
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

func (c *armTypeCreator) addARMType(
	original astmodel.TypeDefinition,
	def astmodel.TypeDefinition,
) {
	c.armDefs.Add(def)
	c.convertedDefs.Add(original)
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
