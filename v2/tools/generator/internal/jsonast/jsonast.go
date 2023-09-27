/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"regexp"
	"strings"

	"github.com/devigned/tab"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type (

	// TypeHandler is a standard delegate used for walking the schema tree.
	// Note that it is permissible for a TypeHandler to return `nil, nil`, which indicates that
	// there is no type to be included in the output.
	TypeHandler func(ctx context.Context, scanner *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error)

	// UnknownSchemaError is used when we find a JSON schema node that we don't know how to handle
	UnknownSchemaError struct {
		Schema  Schema
		Filters []string
	}

	// A SchemaScanner is used to scan a JSON Schema extracting and collecting type definitions
	SchemaScanner struct {
		definitions   map[astmodel.TypeName]*astmodel.TypeDefinition
		TypeHandlers  map[SchemaType]TypeHandler
		configuration *config.Configuration
		idFactory     astmodel.IdentifierFactory
		log           logr.Logger
	}
)

// findTypeDefinition looks to see if we have seen the specified definition before, returning its definition if we have.
func (scanner *SchemaScanner) findTypeDefinition(name astmodel.TypeName) (*astmodel.TypeDefinition, bool) {
	result, ok := scanner.definitions[name]
	return result, ok
}

// addTypeDefinition adds a type definition to emit later
func (scanner *SchemaScanner) addTypeDefinition(def astmodel.TypeDefinition) {
	if existing, ok := scanner.definitions[def.Name()]; ok && existing != nil {
		panic(fmt.Sprintf("overwriting existing definition for %s", def.Name()))
	}

	scanner.definitions[def.Name()] = &def
}

// addEmptyTypeDefinition adds a placeholder definition; it should always be replaced later
func (scanner *SchemaScanner) addEmptyTypeDefinition(name astmodel.TypeName) {
	scanner.definitions[name] = nil
}

// removeTypeDefinition removes a type definition
func (scanner *SchemaScanner) removeTypeDefinition(name astmodel.TypeName) {
	delete(scanner.definitions, name)
}

func (use *UnknownSchemaError) Error() string {
	if use.Schema == nil || use.Schema.url() == nil {
		return "unable to determine schema type for nil schema or one without a URL"
	}

	return fmt.Sprintf("unable to determine the schema type for %s", use.Schema.url().String())
}

// NewSchemaScanner constructs a new scanner, ready for use
func NewSchemaScanner(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	log logr.Logger,
) *SchemaScanner {
	return &SchemaScanner{
		definitions:   make(map[astmodel.TypeName]*astmodel.TypeDefinition),
		TypeHandlers:  defaultTypeHandlers(),
		configuration: configuration,
		idFactory:     idFactory,
		log:           log,
	}
}

// AddTypeHandler will override a default type handler for a given SchemaType. This allows for a consumer to customize
// AST generation.
func (scanner *SchemaScanner) AddTypeHandler(schemaType SchemaType, handler TypeHandler) {
	scanner.TypeHandlers[schemaType] = handler
}

// RunHandler triggers the appropriate handler for the specified schemaType
func (scanner *SchemaScanner) RunHandler(ctx context.Context, schemaType SchemaType, schema Schema) (astmodel.Type, error) {
	if ctx.Err() != nil { // check for cancellation
		return nil, ctx.Err()
	}

	handler := scanner.TypeHandlers[schemaType]
	return handler(ctx, scanner, schema, scanner.log)
}

// RunHandlerForSchema inspects the passed schema to identify what kind it is, then runs the appropriate handler
func (scanner *SchemaScanner) RunHandlerForSchema(ctx context.Context, schema Schema) (astmodel.Type, error) {
	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	return scanner.RunHandler(ctx, schemaType, schema)
}

// RunHandlersForSchemas inspects each passed schema and runs the appropriate handler
func (scanner *SchemaScanner) RunHandlersForSchemas(ctx context.Context, schemas []Schema) ([]astmodel.Type, error) {
	var results []astmodel.Type
	var errs []error
	for _, schema := range schemas {
		t, err := scanner.RunHandlerForSchema(ctx, schema)
		if err != nil {
			var unknownSchema *UnknownSchemaError
			if errors.As(err, &unknownSchema) {
				if unknownSchema.Schema.description() != nil {
					// some Swagger types (e.g. ServiceFabric Cluster) use allOf with a description-only schema
					scanner.log.V(2).Info(
						"skipping description-only schema type",
						"schema", unknownSchema.Schema.url(),
						"description", *unknownSchema.Schema.description())
					continue
				}
			}

			errs = append(errs, errors.Wrapf(err, "unable to handle schema %s", schema.Id()))
		}

		if t != nil {
			results = append(results, t)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (scanner *SchemaScanner) GenerateAllDefinitions(ctx context.Context, schema Schema) (astmodel.TypeDefinitionSet, error) {
	title := schema.title()
	if title == nil {
		return nil, errors.New("given schema has no title")
	}

	rootName := *title
	rootURL := schema.url()
	rootGroup, err := groupOf(rootURL)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to extract group for schema")
	}

	rootVersion := versionOf(rootURL)

	rootPackage := scanner.configuration.MakeLocalPackageReference(
		scanner.idFactory.CreateGroupName(rootGroup),
		rootVersion)
	rootTypeName := astmodel.MakeInternalTypeName(rootPackage, rootName)

	_, err = generateDefinitionsFor(ctx, scanner, rootTypeName, schema)
	if err != nil {
		return nil, err
	}

	return scanner.Definitions(), nil
}

// Definitions produces a set of all the types defined so far
func (scanner *SchemaScanner) Definitions() astmodel.TypeDefinitionSet {
	defs := make(astmodel.TypeDefinitionSet)
	for defName, def := range scanner.definitions {
		if def == nil {
			// safety check/assert:
			panic(fmt.Sprintf("%s was nil", defName))
		}

		if defName != def.Name() {
			// this indicates a serious programming error
			panic(fmt.Sprintf("mismatched typenames: %s != %s", defName, def.Name()))
		}

		defs.Add(*def)
	}

	return defs
}

// defaultTypeHandlers will create a default map of JSONType to AST transformers
func defaultTypeHandlers() map[SchemaType]TypeHandler {
	return map[SchemaType]TypeHandler{
		Array:  arrayHandler,
		OneOf:  oneOfHandler,
		AnyOf:  anyOfHandler,
		AllOf:  allOfHandler,
		Ref:    refHandler,
		Object: objectHandler,
		Enum:   enumHandler,
		String: stringHandler,
		Int:    intHandler,
		Number: numberHandler,
		Bool:   boolHandler,
	}
}

func stringHandler(_ context.Context, _ *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error) {
	t := astmodel.StringType

	maxLength := schema.maxLength()
	minLength := schema.minLength()
	pattern := schema.pattern()
	format := schema.format()

	if maxLength != nil || minLength != nil || pattern != nil || format != "" {
		patterns := make([]*regexp.Regexp, 0, 2)
		if pattern != nil {
			patterns = append(patterns, pattern)
		}

		if format != "" {
			formatPattern := formatToPattern(format, log)
			if formatPattern != nil {
				patterns = append(patterns, formatPattern)
			}
		}

		if format == "arm-id" {
			t = astmodel.ARMIDType
		}

		// If there are no validations except format, and that format didn't result in any patterns,
		// then there's no need for a validated type so just return t
		if len(patterns) == 0 && maxLength == nil && minLength == nil {
			return t, nil
		}

		validations := astmodel.StringValidations{
			MaxLength: maxLength,
			MinLength: minLength,
			Patterns:  patterns,
		}

		return astmodel.NewValidatedType(t, validations), nil
	}

	return t, nil
}

// copied from ARM implementation
var uuidRegex = regexp.MustCompile("^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$")

func formatToPattern(format string, log logr.Logger) *regexp.Regexp {
	switch format {
	case "uuid":
		return uuidRegex
	case "date-time", "date", "duration", "date-time-rfc1123", "arm-id":
		// TODO: don’t bother validating for now
		return nil
	case "password":
		// This is handled later in the status_augment phase of processing, so just
		// ignore it for now
		return nil
	default:
		log.V(1).Info(
			"Unknown property format",
			"format", format)
		return nil
	}
}

func numberHandler(_ context.Context, _ *SchemaScanner, schema Schema, _ logr.Logger) (astmodel.Type, error) {
	t := astmodel.FloatType
	v := getNumberValidations(schema)
	if v != nil {
		// for controller-gen anything with min/max/multipleof must be based on int
		// double-check that all of these are integral
		var errs []string

		if v.Minimum != nil {
			if !v.Minimum.IsInt() {
				errs = append(errs, "'minimum' validation must be an integer")
			}
		}

		if v.Maximum != nil {
			if !v.Maximum.IsInt() {
				errs = append(errs, "'maximum' validation must be an integer")
			}
		}

		if v.MultipleOf != nil {
			if !v.MultipleOf.IsInt() {
				errs = append(errs, "'multipleOf' validation must be an integer")
			}
		}

		result := astmodel.NewValidatedType(t, *v)
		if len(errs) > 0 {
			return astmodel.NewErroredType(result, errs, nil), nil
		}

		// we have checked they are all integers:
		return result, nil
	}

	return t, nil
}

var (
	zero      *big.Rat = big.NewRat(0, 1)
	maxUint32 *big.Rat = big.NewRat(1, 1).SetUint64(math.MaxUint32)
)

func intHandler(_ context.Context, _ *SchemaScanner, schema Schema, _ logr.Logger) (astmodel.Type, error) {
	t := astmodel.IntType
	v := getNumberValidations(schema)
	if v != nil {
		// special-case some things to return different types
		if !v.ExclusiveMaximum && v.Maximum != nil &&
			v.MultipleOf == nil &&
			!v.ExclusiveMinimum && v.Minimum != nil && v.Minimum.Cmp(zero) == 0 {

			if v.Maximum.Cmp(maxUint32) == 0 {
				return astmodel.UInt32Type, nil
			}
		}

		return astmodel.NewValidatedType(t, *v), nil
	}

	return t, nil
}

func getNumberValidations(schema Schema) *astmodel.NumberValidations {
	minValue := schema.minValue()
	minExclusive := schema.minValueExclusive()
	maxValue := schema.maxValue()
	maxExclusive := schema.maxValueExclusive()
	multipleOf := schema.multipleOf()

	if minValue != nil || maxValue != nil || multipleOf != nil {
		return &astmodel.NumberValidations{
			Maximum:          maxValue,
			Minimum:          minValue,
			ExclusiveMaximum: maxExclusive,
			ExclusiveMinimum: minExclusive,
			MultipleOf:       multipleOf,
		}
	}

	return nil
}

func boolHandler(_ context.Context, _ *SchemaScanner, _ Schema, _ logr.Logger) (astmodel.Type, error) {
	return astmodel.BoolType, nil
}

func enumHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, _ logr.Logger) (astmodel.Type, error) {
	_, span := tab.StartSpan(ctx, "enumHandler")
	defer span.End()

	// Default to a string base type
	baseType := astmodel.StringType
	for _, t := range []SchemaType{Bool, Int, Number, String} {
		if schema.hasType(t) {
			bt, err := GetPrimitiveType(t)
			if err != nil {
				return nil, err
			}

			baseType = bt
		}
	}

	enumValues := schema.enumValues()
	values := make([]astmodel.EnumValue, 0, len(enumValues))
	for _, v := range enumValues {
		vTrimmed := strings.Trim(v, "\"")

		// Some specs include boolean (or float, int) enums with quotes around the literals.
		// Trim quotes from anything that's not a string
		if baseType != astmodel.StringType {
			v = vTrimmed
		}

		// TODO: This is a bit of a hack as we don't have a way to handle this generically right now
		// TODO: for an arbitrary non-renderable character
		// use vTrimmed as seed for identifier as it doesn't have quotes surrounding it
		id := scanner.idFactory.CreateIdentifier(vTrimmed, astmodel.Exported)
		values = append(values, astmodel.MakeEnumValue(id, v))
	}

	enumType := astmodel.NewEnumType(baseType, values...)

	return enumType, nil
}

func objectHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "objectHandler")
	defer span.End()

	properties, err := getProperties(ctx, scanner, schema, log)
	if err != nil {
		return nil, err
	}

	// if we _only_ have an 'additionalProperties' property, then we are making
	// a dictionary-like type, and we won't generate an object type; instead, we
	// will just use the 'additionalProperties' type directly
	if len(properties) == 1 && properties[0].PropertyName() == astmodel.AdditionalPropertiesPropertyName {
		return properties[0].PropertyType(), nil
	}

	isResource := schema.extensionAsBool("x-ms-azure-resource")

	// If we're a resource, our 'Id' property needs to have a special type
	if isResource {
		for i, prop := range properties {
			if prop.HasName("Id") || prop.HasName("ID") {
				properties[i] = prop.WithType(astmodel.NewOptionalType(astmodel.ARMIDType))
			}
		}
	}

	objectType := astmodel.NewObjectType().WithProperties(properties...).WithIsResource(isResource)

	return objectType, nil
}

func generatePropertyDefinition(ctx context.Context, scanner *SchemaScanner, rawPropName string, prop Schema) (*astmodel.PropertyDefinition, error) {
	propertyName := scanner.idFactory.CreatePropertyName(rawPropName, astmodel.Exported)

	schemaType, err := getSubSchemaType(prop)
	var use *UnknownSchemaError
	if errors.As(err, &use) {
		// if we don't know the type, we still need to provide the property, we will just provide open interface
		property := astmodel.NewPropertyDefinition(propertyName, rawPropName, astmodel.AnyType)
		return property, nil
	}

	if err != nil {
		return nil, err
	}

	propType, err := scanner.RunHandler(ctx, schemaType, prop)
	if errors.As(err, &use) {
		// if we don't know the type, we still need to provide the property, we will just provide open interface
		property := astmodel.NewPropertyDefinition(propertyName, rawPropName, astmodel.AnyType)
		return property, nil
	}

	if err != nil {
		return nil, err
	}

	// This can happen if the property type was pruned away by a type filter.
	if propType == nil {
		// returning nil here is a signal to the caller that this property cannot be constructed.
		return nil, nil
	}

	property := astmodel.NewPropertyDefinition(propertyName, rawPropName, propType).WithReadOnly(prop.readOnly())
	return property, nil
}

func getProperties(
	ctx context.Context,
	scanner *SchemaScanner,
	schema Schema,
	log logr.Logger,
) ([]*astmodel.PropertyDefinition, error) {
	ctx, span := tab.StartSpan(ctx, "getProperties")
	defer span.End()

	props := schema.properties()
	properties := make([]*astmodel.PropertyDefinition, 0, len(props))
	for propName, propSchema := range props {

		property, err := generatePropertyDefinition(ctx, scanner, propName, propSchema)
		if err != nil {
			return nil, err
		}

		// This can happen if the property type was pruned away by a type filter.
		// There are a few options here: We can skip this property entirely, we can emit it
		// with no type (won't compile), or we can emit with with interface{}.
		// Currently emitting a warning and skipping
		if property == nil {
			// TODO: This log shouldn't happen in cases where the type in question is later excluded, see:
			// TODO: https://github.com/Azure/azure-service-operator/issues/1517
			log.V(2).Info(
				"Property omitted due to nil propType (probably due to type filter)",
				"property", propName)
			continue
		}

		// add documentation
		if propSchema.description() != nil {
			property = property.WithDescription(*propSchema.description())
		}

		// add flattening
		property = property.SetFlatten(propSchema.extensionAsBool("x-ms-client-flatten") == true)

		// add secret flag
		hasSecretExtension := propSchema.extensionAsBool("x-ms-secret")

		hasFormatPassword := propSchema.format() == "password"
		if hasSecretExtension || hasFormatPassword {
			property = property.WithIsSecret(true)
		}

		// add validations
		isRequired := false
		for _, required := range schema.requiredProperties() {
			if propName == required {
				isRequired = true
				break
			}
		}

		// All types are optional (regardless of if the property is required or not) because of
		// non-optional types (int, string, MyType, etc) interaction with omitempty.
		// If a field is json:omitempty but its type is not optional (not a ptr) then the default value
		// for that type will be omitted from the JSON payload. For example 0 would be omitted for ints.
		// On the other hand if a field is NOT json:omitempty then the type is always serialized in the payload
		// which causes issues for kubebuilder:validation:Required (how can we tell the user didn't specify that value?)
		// See https://github.com/Azure/azure-service-operator/issues/1999 for more details.
		property = property.MakeTypeOptional()
		if isRequired {
			property = property.MakeRequired()
		} else {
			property = property.MakeOptional()
		}

		properties = append(properties, property)
	}

	// see: https://json-schema.org/understanding-json-schema/reference/object.html#properties
	if schema.additionalPropertiesAllowed() {
		additionalPropSchema := schema.additionalPropertiesSchema()
		if additionalPropSchema == nil {
			// if not specified, any additional properties are allowed
			// (TODO: tell all Azure teams this fact and get them to update their API definitions!)
			// for now we aren't following the spec 100% as it pollutes the generated code
			// only generate this field if there are no other fields:
			if len(properties) == 0 {
				// TODO: for JSON serialization this needs to be unpacked into "parent"
				additionalProperties := astmodel.NewPropertyDefinition(
					astmodel.AdditionalPropertiesPropertyName,
					astmodel.AdditionalPropertiesJsonName,
					astmodel.NewStringMapType(astmodel.AnyType))

				properties = append(properties, additionalProperties)
			}
		} else {
			// otherwise, it is a type for all additional fields
			// TODO: for JSON serialization this needs to be unpacked into "parent"
			additionalPropsType, err := scanner.RunHandlerForSchema(ctx, additionalPropSchema)
			if err != nil {
				return nil, err
			}

			// This can happen if the property type was pruned away by a type filter.
			// There are a few options here: We can skip this property entirely, we can emit it
			// with no type (won't compile), or we can emit with with interface{}.
			// TODO: Currently setting this to anyType as that's easiest to deal with and will generate
			// TODO: a warning during controller-gen
			if additionalPropsType == nil {
				additionalPropsType = astmodel.AnyType
			}

			additionalProperties := astmodel.NewPropertyDefinition(
				astmodel.AdditionalPropertiesPropertyName,
				astmodel.AdditionalPropertiesJsonName,
				astmodel.NewStringMapType(additionalPropsType))

			properties = append(properties, additionalProperties)
		}
	}

	return properties, nil
}

func refHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "refHandler")
	defer span.End()

	// Allow for inclusions of other full-fledged files, such as the autogenerated list in the 2019-04-01
	// deploymentTemplate.json
	refSchema := schema.refSchema()
	// TODO: How hacky is this?
	if refSchema.url().Host != "" && refSchema.url().Fragment == "" {
		return scanner.RunHandlerForSchema(ctx, refSchema)
	}

	typeName, err := schema.refTypeName()
	if err != nil {
		return nil, err
	}

	// Prune the graph according to the configuration
	shouldPrune, because := scanner.configuration.ShouldPrune(typeName)
	if shouldPrune == config.Prune {
		log.V(2).Info(
			"Skipping type",
			"type", typeName,
			"because", because)
		return nil, nil // Skip entirely
	}

	return generateDefinitionsFor(ctx, scanner, typeName, schema.refSchema())
}

func generateDefinitionsFor(
	ctx context.Context,
	scanner *SchemaScanner,
	typeName astmodel.InternalTypeName,
	schema Schema,
) (astmodel.TypeName, error) {
	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	schemaUrl := schema.url()

	// see if we already generated something for this ref
	if _, ok := scanner.findTypeDefinition(typeName); ok {
		return typeName, nil
	}

	// Add a placeholder to avoid recursive calls
	// we will overwrite this later (this is checked below)
	scanner.addEmptyTypeDefinition(typeName)
	result, err := scanner.RunHandler(ctx, schemaType, schema)
	if err != nil {
		scanner.removeTypeDefinition(typeName) // we weren't able to generate it, remove placeholder
		return nil, err
	}

	// TODO: This code and below does nothing in the Swagger path as schema.url() is always empty.
	// TODO: It's still used in the JSON schema path for golden tests and should be removed once those
	// TODO: are retired.
	resourceType := categorizeResourceType(schemaUrl)
	if resourceType != nil {
		result = astmodel.NewAzureResourceType(result, nil, typeName, *resourceType)
	}
	definition := astmodel.MakeTypeDefinition(typeName, result)

	// Add a description of the type
	var description []string
	if desc := schema.description(); desc != nil {
		description = astbuilder.WordWrap(*desc, 120)
	}

	// Add URL reference if we have one
	if schema.url().String() != "" {
		description = append(description, fmt.Sprintf("Generated from: %s", schema.url().String()))
	}

	if len(description) > 0 {
		definition = definition.WithDescription(description...)
	}

	scanner.addTypeDefinition(definition)

	if def, ok := scanner.findTypeDefinition(typeName); !ok || def == nil {
		// safety check in case of breaking changes
		panic(fmt.Sprintf("didn't set type definition for %s", typeName))
	}

	return definition.Name(), nil
}

func allOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, _ logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "allOfHandler")
	defer span.End()

	types, err := scanner.RunHandlersForSchemas(ctx, schema.allOf())
	if err != nil {
		return nil, err
	}

	// if the node that contains the allOf defines other properties, create an object type with them inside to merge
	if len(schema.properties()) > 0 {
		objectType, err := scanner.RunHandler(ctx, Object, schema)
		if err != nil {
			return nil, err
		}

		types = append(types, objectType)
	}

	return astmodel.BuildAllOfType(types...), nil
}

func oneOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, _ logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "oneOfHandler")
	defer span.End()

	result := astmodel.NewOneOfType(schema.Id())

	// Capture the discriminator property name, if we have one
	if schema.discriminator() != "" {
		result = result.WithDiscriminatorProperty(schema.discriminator())
	}

	// Capture the discriminator value, if we have one
	if discriminatorValue, ok := schema.extensionAsString("x-ms-discriminator-value"); ok {
		result = result.WithDiscriminatorValue(discriminatorValue)
	}

	// Handle any nested OneOf options
	// These will each be either a TypeName or an Object
	types, err := scanner.RunHandlersForSchemas(ctx, schema.oneOf())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to generate oneOf types for %s", schema.Id())
	}

	result = result.WithTypes(types)

	// Also need to pick up any nested AllOf options
	// If an object, these are properties that are required for this option
	// Otherwise, add as an option
	allOfTypes, err := scanner.RunHandlersForSchemas(ctx, schema.allOf())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to generate allOf types for %s", schema.Id())
	}

	// Our AllOf contains either properties to add to our OneOf, or references to parent types
	for _, t := range allOfTypes {
		if obj, ok := asCommonProperties(t, scanner.definitions); ok {
			// If we have an object, add its properties
			result = result.WithAdditionalPropertyObject(obj)
			continue
		}

		// Treat it as an option
		result = result.WithType(t)
	}

	// If there are any properties, we need to create an object type to wrap them
	if len(schema.properties()) > 0 {
		t, err := scanner.RunHandler(ctx, Object, schema)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate object for properties of %s", schema.Id())
		}

		obj, ok := t.(*astmodel.ObjectType)
		if !ok {
			return nil, errors.Errorf(
				"expected object type for properties of %s, got %T", schema.Id(), t)
		}

		result = result.WithAdditionalPropertyObject(obj)
	}

	return result, nil
}

// attempts to resolve the passed type to an object type that represents properties
func asCommonProperties(
	t astmodel.Type,
	defs map[astmodel.TypeName]*astmodel.TypeDefinition,
) (*astmodel.ObjectType, bool) {
	if obj, isObject := astmodel.AsObjectType(t); isObject {
		return obj, true
	}

	if tn, isTypeName := astmodel.AsTypeName(t); isTypeName {
		if def, found := defs[tn]; found {
			return asCommonProperties(def.Type(), defs)
		}
	}

	return nil, false
}

func anyOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "anyOfHandler")
	defer span.End()

	// See https://github.com/Azure/azure-service-operator/issues/1518 for details about why this is treated as oneOf
	log.V(1).Info(
		"Handling anyOf type as if it were oneOf",
		"url", schema.url())
	return oneOfHandler(ctx, scanner, schema, log)
}

func arrayHandler(ctx context.Context, scanner *SchemaScanner, schema Schema, log logr.Logger) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "arrayHandler")
	defer span.End()

	items := schema.items()
	if len(items) > 1 {
		return nil, errors.Errorf("item contains more children than expected: %s", schema.items())
	}

	if len(items) == 0 {
		// there is no type to the elements, so we must assume interface{}
		log.V(1).Info(
			"Interface assumption unproven",
			"url", schema.url())

		result := astmodel.NewArrayType(astmodel.AnyType)
		return withArrayValidations(schema, result), nil
	}

	// get the only child type and wrap it up as an array type:

	onlyChild := items[0]

	astType, err := scanner.RunHandlerForSchema(ctx, onlyChild)
	if err != nil {
		return nil, err
	}

	// astType can be nil if it was pruned from the tree
	if astType == nil {
		return nil, nil
	}

	result := astmodel.NewArrayType(astType)
	return withArrayValidations(schema, result), nil
}

func withArrayValidations(schema Schema, t *astmodel.ArrayType) astmodel.Type {
	maxItems := schema.maxItems()
	minItems := schema.minItems()
	uniqueItems := schema.uniqueItems()

	if maxItems != nil || minItems != nil || uniqueItems {
		return astmodel.NewValidatedType(t, astmodel.ArrayValidations{
			MaxItems:    maxItems,
			MinItems:    minItems,
			UniqueItems: uniqueItems,
		})
	}

	return t
}

func getSubSchemaType(schema Schema) (SchemaType, error) {

	// handle special nodes:
	switch {
	case len(schema.enumValues()) > 0: // this should come before the primitive checks below
		return Enum, nil
	// Handle the three cases of oneOf: complete/root/leaf
	case schema.hasOneOf() || schema.discriminator() != "" || schema.hasExtension("x-ms-discriminator-value"):
		return OneOf, nil
	case schema.hasAllOf():
		return AllOf, nil
	case schema.hasAnyOf():
		return AnyOf, nil
	case schema.isRef():
		return Ref, nil
	}

	for _, t := range []SchemaType{Object, String, Number, Int, Bool, Array} {
		if schema.hasType(t) {
			return t, nil
		}
	}

	// TODO: this whole switch is a bit wrong because type: 'object' can
	// be combined with OneOf/AnyOf/etc. still, it works okay for now...
	if len(schema.properties()) > 0 || schema.additionalPropertiesSchema() != nil {
		// haven't figured out a type but it has properties, treat it as an object
		return Object, nil
	}

	return Unknown, &UnknownSchemaError{Schema: schema}
}

// GetPrimitiveType returns the primtive type for this schema type
func GetPrimitiveType(name SchemaType) (*astmodel.PrimitiveType, error) {
	switch name {
	case String:
		return astmodel.StringType, nil
	case Int:
		return astmodel.IntType, nil
	case Number:
		return astmodel.FloatType, nil
	case Bool:
		return astmodel.BoolType, nil
	case AllOf:
	case AnyOf:
	case Array:
	case Enum:
	case Object:
	case OneOf:
	case Ref:
	case Unknown:
		return astmodel.AnyType, errors.Errorf("%s is not a simple type and no ast.NewIdent can be created", name)
	}

	panic(fmt.Sprintf("unhandled case in getPrimitiveType: %s", name)) // this is also checked by linter
}

// categorizeResourceType determines if this URL represents an ARM resource or not.
// If the URL represents a resource, a non-nil value is returned. If the URL does not represent
// a resource, nil is returned.
func categorizeResourceType(url *url.URL) *astmodel.ResourceScope {
	fragmentParts := strings.FieldsFunc(url.Fragment, isURLPathSeparator)

	resourceGroup := astmodel.ResourceScopeResourceGroup
	extension := astmodel.ResourceScopeExtension
	tenant := astmodel.ResourceScopeTenant

	for _, fragmentPart := range fragmentParts {
		// resourceDefinitions are "normal" resources
		if fragmentPart == "resourceDefinitions" ||
			// Treat all resourceBase things as resources so that "resourceness"
			// is inherited:
			strings.Contains(strings.ToLower(fragmentPart), "resourcebase") {
			return &resourceGroup
		}

		if fragmentPart == "tenant_resourceDefinitions" {
			return &tenant
		}

		// unknown_ResourceDefinitions or extension_resourceDefinitions are extension resources, see
		// https://github.com/Azure/azure-resource-manager-schemas/blob/069dc7cbff0725aea3a3595e4bb777da966dbb6f/generator/generate.ts#L186
		// to learn more.
		if fragmentPart == "unknown_resourceDefinitions" ||
			fragmentPart == "extension_resourceDefinitions" {
			return &extension
		}
	}

	return nil
}
