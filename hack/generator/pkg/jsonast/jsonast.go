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
	"strings"

	"github.com/devigned/tab"
	"github.com/pkg/errors"

	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
)

type (

	// TypeHandler is a standard delegate used for walking the schema tree.
	// Note that it is permissible for a TypeHandler to return `nil, nil`, which indicates that
	// there is no type to be included in the output.
	TypeHandler func(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error)

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
		panic(fmt.Sprintf("overwriting existing definition for %v", def.Name()))
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
func NewSchemaScanner(idFactory astmodel.IdentifierFactory, configuration *config.Configuration) *SchemaScanner {
	return &SchemaScanner{
		definitions:   make(map[astmodel.TypeName]*astmodel.TypeDefinition),
		TypeHandlers:  defaultTypeHandlers(),
		configuration: configuration,
		idFactory:     idFactory,
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
	return handler(ctx, scanner, schema)
}

// RunHandlerForSchema inspects the passed schema to identify what kind it is, then runs the appropriate handler
func (scanner *SchemaScanner) RunHandlerForSchema(ctx context.Context, schema Schema) (astmodel.Type, error) {
	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	return scanner.RunHandler(ctx, schemaType, schema)
}

// GenerateDefinitionsFromDeploymentTemplate takes in the resources section of the Azure deployment template schema and returns golang AST Packages
//    containing the types described in the schema which match the {resource_type}/{version} filters provided.
//
// 		The schema we are working with is something like the following (in yaml for brevity):
//
// 		resources:
// 			items:
// 				oneOf:
// 					allOf:
// 						$ref: {{ base resource schema for ARM }}
// 						oneOf:
// 							- ARM resources
// 				oneOf:
// 					allOf:
// 						$ref: {{ base resource for external resources, think SendGrid }}
// 						oneOf:
// 							- External ARM resources
// 				oneOf:
// 					allOf:
// 						$ref: {{ base resource for ARM specific stuff like locks, deployments, etc }}
// 						oneOf:
// 							- ARM specific resources. I'm not 100% sure why...
//
// 		allOf acts like composition which composites each schema from the child oneOf with the base reference from allOf.
func (scanner *SchemaScanner) GenerateDefinitionsFromDeploymentTemplate(ctx context.Context, root Schema) (astmodel.Types, error) {
	ctx, span := tab.StartSpan(ctx, "GenerateDefinitionsFromDeploymentTemplate")
	defer span.End()

	resourcesProp, ok := root.properties()["resources"]
	if !ok {
		return nil, errors.Errorf("unable to find 'resources' property in deployment template")
	}

	resourcesTypes, err := scanner.RunHandlerForSchema(ctx, resourcesProp)
	if err != nil {
		return nil, err
	}

	resourcesArray, ok := resourcesTypes.(*astmodel.ArrayType)
	if !ok {
		return nil, errors.Errorf("expected 'resources' property to be an array")
	}

	resourcesOneOf, ok := resourcesArray.Element().(astmodel.OneOfType)
	if !ok {
		return nil, errors.Errorf("expected 'resources' property to be an array containing oneOf")
	}

	err = resourcesOneOf.Types().ForEachError(func(oneType astmodel.Type, _ int) error {
		allOf, ok := oneType.(astmodel.AllOfType)
		if !ok {
			return errors.Errorf("unexpected resource shape: not an allOf")
		}

		var resourceRef astmodel.TypeName
		var objectBase astmodel.TypeName
		found := 0
		allOf.Types().ForEach(func(t astmodel.Type, _ int) {
			if typeName, ok := t.(astmodel.TypeName); ok {
				if !strings.Contains(strings.ToLower(typeName.Name()), "resourcebase") {
					resourceRef = typeName
				} else {
					objectBase = typeName
				}
				found++
			}
		})

		if found != 2 {
			return errors.Errorf("unexpected resource shape: expected a ref to base and ref to object")
		}

		resourceDef, ok := scanner.findTypeDefinition(resourceRef)
		if !ok {
			return errors.Errorf("unable to resolve resource definition for %v", resourceRef)
		}

		resourceType, ok := resourceDef.Type().(*astmodel.ResourceType)
		if !ok {
			// safety check
			return errors.Errorf("resource reference %v in deployment template did not resolve to resource type", resourceRef)
		}

		// now we will remove the existing resource definition and replace it with a new one that includes the base type
		// first, reconstruct the allof with an anonymous type instead of the typename
		specType := astmodel.MakeAllOfType(objectBase, resourceType.SpecType())
		// now replace it
		scanner.removeTypeDefinition(resourceRef)
		scanner.addTypeDefinition(resourceDef.WithType(astmodel.NewAzureResourceType(specType, nil, resourceDef.Name())))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return scanner.Definitions(), nil
}

func (scanner *SchemaScanner) GenerateAllDefinitions(ctx context.Context, schema Schema) (astmodel.Types, error) {
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

	rootPackage := astmodel.MakeLocalPackageReference(
		scanner.idFactory.CreateGroupName(rootGroup),
		scanner.idFactory.CreatePackageNameFromVersion(rootVersion))

	rootTypeName := astmodel.MakeTypeName(rootPackage, rootName)

	_, err = generateDefinitionsFor(ctx, scanner, rootTypeName, schema)
	if err != nil {
		return nil, err
	}

	return scanner.Definitions(), nil
}

// Definitions produces a set of all the types defined so far
func (scanner *SchemaScanner) Definitions() astmodel.Types {

	defs := make(astmodel.Types)
	for defName, def := range scanner.definitions {
		if def == nil {
			// safety check/assert:
			panic(fmt.Sprintf("%v was nil", defName))
		}

		if defName != def.Name() {
			// this indicates a serious programming error
			panic(fmt.Sprintf("mismatched typenames: %v != %v", defName, def.Name()))
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

func stringHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	t := astmodel.StringType

	maxLength := schema.maxLength()
	minLength := schema.minLength()
	pattern := schema.pattern()

	if maxLength != nil || minLength != nil || pattern != nil {
		return astmodel.MakeValidatedType(t, astmodel.StringValidations{
			MaxLength: maxLength,
			MinLength: minLength,
			Pattern:   pattern,
		}), nil
	}

	return t, nil
}

func numberHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	t := astmodel.FloatType
	v := getNumberValidations(schema)
	if v != nil {
		// for controller-gen anything with min/max/multipleof must be based on int
		// double-check that all of these are integral
		if v.Minimum != nil {
			if !v.Minimum.IsInt() {
				return nil, errors.Errorf("'minimum' validation must be an integer")
			}
		}

		if v.Maximum != nil {
			if !v.Maximum.IsInt() {
				return nil, errors.Errorf("'maximum' validation must be an integer")
			}
		}

		if v.MultipleOf != nil {
			if !v.MultipleOf.IsInt() {
				return nil, errors.Errorf("'multipleOf' validation must be an integer")
			}
		}

		t = astmodel.IntType
		return astmodel.MakeValidatedType(t, *v), nil
	}

	return t, nil
}

var zero *big.Rat = big.NewRat(0, 1)
var maxUint32 *big.Rat = big.NewRat(1, 1).SetUint64(math.MaxUint32)

func intHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
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

		return astmodel.MakeValidatedType(t, *v), nil
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

func boolHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	return astmodel.BoolType, nil
}

func enumHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
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

	var values []astmodel.EnumValue
	for _, v := range schema.enumValues() {

		vTrimmed := strings.Trim(v, "\"")

		// Some specs include boolean enums with quotes around the boolean literals
		if baseType == astmodel.BoolType {
			v = vTrimmed
		}

		// TODO: This is a bit of a hack as we don't have a way to handle this generically right now
		// TODO: for an arbitrary non-renderable character
		// use vTrimmed as seed for identifier as it doesn't have quotes surrounding it
		id := scanner.idFactory.CreateIdentifier(vTrimmed, astmodel.Exported)
		values = append(values, astmodel.EnumValue{Identifier: id, Value: v})
	}

	enumType := astmodel.NewEnumType(baseType, values)

	return enumType, nil
}

func objectHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "objectHandler")
	defer span.End()

	properties, err := getProperties(ctx, scanner, schema)
	if err != nil {
		return nil, err
	}

	// if we _only_ have an 'additionalProperties' property, then we are making
	// a dictionary-like type, and we won't generate an object type; instead, we
	// will just use the 'additionalProperties' type directly
	if len(properties) == 1 && properties[0].PropertyName() == "additionalProperties" {
		return properties[0].PropertyType(), nil
	}

	objectType := astmodel.NewObjectType().WithProperties(properties...)
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

	property := astmodel.NewPropertyDefinition(propertyName, rawPropName, propType)
	return property, nil
}

func getProperties(
	ctx context.Context,
	scanner *SchemaScanner,
	schema Schema) ([]*astmodel.PropertyDefinition, error) {

	ctx, span := tab.StartSpan(ctx, "getProperties")

	defer span.End()

	var properties []*astmodel.PropertyDefinition
	for propName, propSchema := range schema.properties() {

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
			// TODO: https://github.com/Azure/k8s-infra/issues/138
			klog.V(2).Infof("Property %s omitted due to nil propType (probably due to type filter)", propName)
			continue
		}

		// add documentation
		if propSchema.description() != nil {
			property = property.WithDescription(*propSchema.description())
		}

		// add validations
		isRequired := false
		for _, required := range schema.requiredProperties() {
			if propName == required {
				isRequired = true
				break
			}
		}

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
					"additionalProperties",
					"additionalProperties",
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
				astmodel.PropertyName("additionalProperties"),
				"additionalProperties",
				astmodel.NewStringMapType(additionalPropsType))

			properties = append(properties, additionalProperties)
		}
	}

	return properties, nil
}

func refHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "refHandler")
	defer span.End()

	// make a new topic based on the ref URL
	name, err := schema.refObjectName()
	if err != nil {
		return nil, err
	}

	group, err := schema.refGroupName()
	if err != nil {
		return nil, err
	}

	version, err := schema.refVersion()
	if err != nil {
		return nil, err
	}

	// produce a usable name:
	typeName := astmodel.MakeTypeName(
		astmodel.MakeLocalPackageReference(
			scanner.idFactory.CreateGroupName(group),
			scanner.idFactory.CreatePackageNameFromVersion(version)),
		scanner.idFactory.CreateIdentifier(name, astmodel.Exported))

	// Prune the graph according to the configuration
	shouldPrune, because := scanner.configuration.ShouldPrune(typeName)
	if shouldPrune == config.Prune {
		klog.V(3).Infof("Skipping %s because %s", typeName, because)
		return nil, nil // Skip entirely
	}

	// Target types according to configuration
	transformation, because := scanner.configuration.TransformType(typeName)
	if transformation != nil {
		klog.V(2).Infof("Transforming %s -> %s because %s", typeName, transformation, because)
		return transformation, nil
	}

	return generateDefinitionsFor(ctx, scanner, typeName, schema.refSchema())
}

func generateDefinitionsFor(
	ctx context.Context,
	scanner *SchemaScanner,
	typeName astmodel.TypeName,
	schema Schema) (astmodel.Type, error) {

	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	url := schema.url()

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

	if isResource(url) {
		result = astmodel.NewAzureResourceType(result, nil, typeName)
	}

	description := []string{
		fmt.Sprintf("Generated from: %s", schema.url().String()),
	}
	definition := astmodel.MakeTypeDefinition(typeName, result).WithDescription(description)

	scanner.addTypeDefinition(definition)

	if def, ok := scanner.findTypeDefinition(typeName); !ok || def == nil {
		// safety check in case of breaking changes
		panic(fmt.Sprintf("didn't set type definition for %v", typeName))
	}

	return definition.Name(), nil
}

func allOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "allOfHandler")
	defer span.End()

	var types []astmodel.Type
	for _, all := range schema.allOf() {

		d, err := scanner.RunHandlerForSchema(ctx, all)
		if err != nil {
			var unknownSchema *UnknownSchemaError
			if errors.As(err, &unknownSchema) {
				if unknownSchema.Schema.description() != nil {
					// some Swagger types (e.g. ServiceFabric Cluster) use allOf with a description-only schema
					klog.V(2).Infof("skipping description-only schema type with description %q", *unknownSchema.Schema.description())
					continue
				}
			}

			return nil, err
		}

		if d != nil {
			types = append(types, d)
		}
	}

	// if the node that contains the allOf defines other properties, create an object type with them inside to merge
	if len(schema.properties()) > 0 {
		objectType, err := scanner.RunHandler(ctx, Object, schema)
		if err != nil {
			return nil, err
		}

		types = append(types, objectType)
	}

	return astmodel.MakeAllOfType(types...), nil
}

func oneOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "oneOfHandler")
	defer span.End()

	return generateOneOfUnionType(ctx, schema, schema.oneOf(), scanner)
}

func generateOneOfUnionType(ctx context.Context, schema Schema, subschemas []Schema, scanner *SchemaScanner) (astmodel.Type, error) {
	// make sure we visit everything before bailing out,
	// to get all types generated even if we can't use them
	var types []astmodel.Type
	for _, one := range subschemas {
		result, err := scanner.RunHandlerForSchema(ctx, one)
		if err != nil {
			return nil, err
		}

		if result != nil {
			types = append(types, result)
		}
	}

	result := astmodel.MakeOneOfType(types...)

	// if the node that contains the oneOf(/anyOf) defines other properties, create an object type with them inside to merge
	if len(schema.properties()) > 0 {
		objectType, err := scanner.RunHandler(ctx, Object, schema)
		if err != nil {
			return nil, err
		}

		result = astmodel.MakeAllOfType(objectType, result)
	}

	return result, nil
}

func anyOfHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "anyOfHandler")
	defer span.End()

	// See https://github.com/Azure/k8s-infra/issues/111 for details about why this is treated as oneOf
	klog.V(2).Infof("Handling anyOf type as if it were oneOf: %v\n", schema.url()) // TODO: was Ref.URL
	return generateOneOfUnionType(ctx, schema, schema.anyOf(), scanner)
}

func arrayHandler(ctx context.Context, scanner *SchemaScanner, schema Schema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "arrayHandler")
	defer span.End()

	items := schema.items()
	if len(items) > 1 {
		return nil, errors.Errorf("item contains more children than expected: %v", schema.items())
	}

	if len(items) == 0 {
		// there is no type to the elements, so we must assume interface{}
		klog.Warningf("Interface assumption unproven for %v\n", schema.url())

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
		return astmodel.MakeValidatedType(t, astmodel.ArrayValidations{
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
	case schema.hasOneOf():
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
	if len(schema.properties()) > 0 {
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

func isResource(url *url.URL) bool {
	fragmentParts := strings.FieldsFunc(url.Fragment, isURLPathSeparator)

	for _, fragmentPart := range fragmentParts {
		if fragmentPart == "resourceDefinitions" ||

			// EventGrid does this, unsure why:
			fragmentPart == "unknown_resourceDefinitions" ||

			// Treat all resourceBase things as resources so that "resourceness"
			// is inherited:
			strings.Contains(strings.ToLower(fragmentPart), "resourcebase") {

			return true
		}
	}

	return false
}
