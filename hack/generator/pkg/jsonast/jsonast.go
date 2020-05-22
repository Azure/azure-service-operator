/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"fmt"
	"k8s.io/klog/v2"
	"net/url"
	"regexp"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/devigned/tab"
	"github.com/xeipuuv/gojsonschema"
)

type (
	// SchemaType defines the type of JSON schema node we are currently processing
	SchemaType string

	// TypeHandler is a standard delegate used for walking the schema tree.
	// Note that it is permissible for a TypeHandler to return `nil, nil`, which indicates that
	// there is no type to be included in the output.
	TypeHandler func(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error)

	// UnknownSchemaError is used when we find a JSON schema node that we don't know how to handle
	UnknownSchemaError struct {
		Schema  *gojsonschema.SubSchema
		Filters []string
	}

	// A BuilderOption is used to provide custom configuration for our scanner
	BuilderOption func(scanner *SchemaScanner) error

	// A SchemaScanner is used to scan a JSON Schema extracting and collecting type definitions
	SchemaScanner struct {
		Definitions  map[astmodel.DefinitionName]astmodel.Definition
		TypeHandlers map[SchemaType]TypeHandler
		Filters      []string
		idFactory    astmodel.IdentifierFactory
	}
)

// FindDefinition looks to see if we have seen the specified definiton before, returning its definition if we have.
func (scanner *SchemaScanner) FindDefinition(ref astmodel.DefinitionName) (astmodel.Definition, bool) {
	result, ok := scanner.Definitions[ref]
	return result, ok
}

// AddDefinition makes a record of the specified struct so that FindStruct() can return it when it is needed again.
func (scanner *SchemaScanner) AddDefinition(def astmodel.Definition) {
	scanner.Definitions[*def.Reference()] = def
}

// Definitions for different kinds of JSON schema
const (
	AnyOf   SchemaType = "anyOf"
	AllOf   SchemaType = "allOf"
	OneOf   SchemaType = "oneOf"
	Ref     SchemaType = "ref"
	Array   SchemaType = "array"
	Bool    SchemaType = "boolean"
	Int     SchemaType = "integer"
	Number  SchemaType = "number"
	Object  SchemaType = "object"
	String  SchemaType = "string"
	Enum    SchemaType = "enum"
	Unknown SchemaType = "unknown"

	expressionFragment = "/definitions/expression"
)

func (use *UnknownSchemaError) Error() string {
	if use.Schema == nil || use.Schema.ID == nil {
		return fmt.Sprint("unable to determine schema type for nil schema or one without an ID")
	}
	return fmt.Sprintf("unable to determine the schema type for %s", use.Schema.ID.String())
}

// NewSchemaScanner constructs a new scanner, ready for use
func NewSchemaScanner(idFactory astmodel.IdentifierFactory) *SchemaScanner {
	return &SchemaScanner{
		Definitions:  make(map[astmodel.DefinitionName]astmodel.Definition),
		TypeHandlers: DefaultTypeHandlers(),
		idFactory:    idFactory,
	}
}

// AddTypeHandler will override a default type handler for a given SchemaType. This allows for a consumer to customize
// AST generation.
func (scanner *SchemaScanner) AddTypeHandler(schemaType SchemaType, handler TypeHandler) {
	scanner.TypeHandlers[schemaType] = handler
}

// RunHandler triggers the appropriate handler for the specified schemaType
func (scanner *SchemaScanner) RunHandler(ctx context.Context, schemaType SchemaType, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	handler := scanner.TypeHandlers[schemaType]
	return handler(ctx, scanner, schema)
}

// RunHandlerForSchema inspects the passed schema to identify what kind it is, then runs the appropriate handler
func (scanner *SchemaScanner) RunHandlerForSchema(ctx context.Context, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	return scanner.RunHandler(ctx, schemaType, schema)
}

// AddFilters will add a filter (perhaps not currently used?)
func (scanner *SchemaScanner) AddFilters(filters []string) {
	scanner.Filters = append(scanner.Filters, filters...)
}

// ToNodes takes in the resources section of the Azure deployment template schema and returns golang AST Packages
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
func (scanner *SchemaScanner) ToNodes(ctx context.Context, schema *gojsonschema.SubSchema, opts ...BuilderOption) (astmodel.Definition, error) {
	ctx, span := tab.StartSpan(ctx, "ToNodes")
	defer span.End()

	for _, opt := range opts {
		if err := opt(scanner); err != nil {
			return nil, err
		}
	}

	schemaType, err := getSubSchemaType(schema)
	if err != nil {
		return nil, err
	}

	// get initial topic from ID and Title:
	url := schema.ID.GetUrl()
	if schema.Title == nil {
		return nil, fmt.Errorf("Given schema has no Title")
	}

	rootStructName := *schema.Title

	rootStructGroup, err := groupOf(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract group for schema: %w", err)
	}

	rootStructVersion, err := versionOf(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract version for schema: %w", err)
	}

	nodes, err := scanner.RunHandler(ctx, schemaType, schema)
	if err != nil {
		return nil, err
	}

	rootStructRef := astmodel.NewStructReference(
		scanner.idFactory.CreateIdentifier(rootStructName),
		scanner.idFactory.CreateGroupName(rootStructGroup),
		scanner.idFactory.CreatePackageNameFromVersion(rootStructVersion),
		false)

	// TODO: make safer:
	root := astmodel.NewStructDefinition(rootStructRef, nodes.(*astmodel.StructType))
	description := "Generated from: " + url.String()
	root = root.WithDescription(&description)

	scanner.AddDefinition(root)

	return root, nil
}

// DefaultTypeHandlers will create a default map of JSONType to AST transformers
func DefaultTypeHandlers() map[SchemaType]TypeHandler {
	return map[SchemaType]TypeHandler{
		Array:  arrayHandler,
		OneOf:  oneOfHandler,
		AnyOf:  anyOfHandler,
		AllOf:  allOfHandler,
		Ref:    refHandler,
		Object: objectHandler,
		Enum:   enumHandler,
		String: fixedTypeHandler(astmodel.StringType, "string"),
		Int:    fixedTypeHandler(astmodel.IntType, "int"),
		Number: fixedTypeHandler(astmodel.FloatType, "number"),
		Bool:   fixedTypeHandler(astmodel.BoolType, "bool"),
	}
}

func enumHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	_, span := tab.StartSpan(ctx, "enumHandler")
	defer span.End()

	// Default to a string base type
	baseType := astmodel.StringType
	for _, t := range []SchemaType{Bool, Int, Number, String} {
		if schema.Types.Contains(string(t)) {
			bt, err := getPrimitiveType(t)
			if err != nil {
				return nil, err
			}

			baseType = bt
		}
	}

	var values []astmodel.EnumValue
	for _, v := range schema.Enum {
		id := scanner.idFactory.CreateIdentifier(v)
		values = append(values, astmodel.EnumValue{Identifier: id, Value: v})
	}

	enumType := astmodel.NewEnumType(baseType, values)

	return enumType, nil
}

func fixedTypeHandler(typeToReturn astmodel.Type, handlerName string) TypeHandler {
	return func(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
		_, span := tab.StartSpan(ctx, handlerName+"Handler")
		defer span.End()

		return typeToReturn, nil
	}
}

func objectHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "objectHandler")
	defer span.End()

	fields, err := getFields(ctx, scanner, schema)
	if err != nil {
		return nil, err
	}

	// if we _only_ have an 'additionalProperties' field, then we are making
	// a dictionary-like type, and we won't generate a struct; instead, we
	// will just use the 'additionalProperties' type directly
	if len(fields) == 1 && fields[0].FieldName() == "additionalProperties" {
		return fields[0].FieldType(), nil
	}

	structDefinition := astmodel.NewStructType(fields...)
	return structDefinition, nil
}

func generateFieldDefinition(ctx context.Context, scanner *SchemaScanner, prop *gojsonschema.SubSchema) (*astmodel.FieldDefinition, error) {
	fieldName := scanner.idFactory.CreateFieldName(prop.Property)

	schemaType, err := getSubSchemaType(prop)
	if _, ok := err.(*UnknownSchemaError); ok {
		// if we don't know the type, we still need to provide the property, we will just provide open interface
		field := astmodel.NewFieldDefinition(fieldName, prop.Property, astmodel.AnyType)
		return field, nil
	}

	if err != nil {
		return nil, err
	}

	propType, err := scanner.RunHandler(ctx, schemaType, prop)
	if _, ok := err.(*UnknownSchemaError); ok {
		// if we don't know the type, we still need to provide the property, we will just provide open interface
		field := astmodel.NewFieldDefinition(fieldName, prop.Property, astmodel.AnyType)
		return field, nil
	}

	if err != nil {
		return nil, err
	}

	field := astmodel.NewFieldDefinition(fieldName, prop.Property, propType)
	return field, nil
}

func getFields(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) ([]*astmodel.FieldDefinition, error) {
	ctx, span := tab.StartSpan(ctx, "getFields")
	defer span.End()

	var fields []*astmodel.FieldDefinition
	for _, prop := range schema.PropertiesChildren {

		fieldDefinition, err := generateFieldDefinition(ctx, scanner, prop)
		if err != nil {
			return nil, err
		}

		// add documentation
		fieldDefinition = fieldDefinition.WithDescription(prop.Description)

		// add validations
		isRequired := false
		for _, required := range schema.Required {
			if prop.Property == required {
				isRequired = true
				break
			}
		}

		if isRequired {
			fieldDefinition = fieldDefinition.MakeRequired()
		} else {
			fieldDefinition = fieldDefinition.MakeOptional()
		}

		fields = append(fields, fieldDefinition)
	}

	// see: https://json-schema.org/understanding-json-schema/reference/object.html#properties
	if schema.AdditionalProperties == nil {
		// if not specified, any additional properties are allowed (TODO: tell all Azure teams this fact and get them to update their API definitions)
		// for now we aren't following the spec 100% as it pollutes the generated code
		// only generate this field if there are no other fields:
		if len(fields) == 0 {
			// TODO: for JSON serialization this needs to be unpacked into "parent"
			additionalPropsField := astmodel.NewFieldDefinition("additionalProperties", "additionalProperties", astmodel.NewStringMap(astmodel.AnyType))
			fields = append(fields, additionalPropsField)
		}
	} else if schema.AdditionalProperties != false {
		// otherwise, if not false then it is a type for all additional fields
		// TODO: for JSON serialization this needs to be unpacked into "parent"
		additionalPropsType, err := scanner.RunHandlerForSchema(ctx, schema.AdditionalProperties.(*gojsonschema.SubSchema))
		if err != nil {
			return nil, err
		}

		additionalPropsField := astmodel.NewFieldDefinition(astmodel.FieldName("additionalProperties"), "additionalProperties", astmodel.NewStringMap(additionalPropsType))
		fields = append(fields, additionalPropsField)
	}

	return fields, nil
}

func refHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "refHandler")
	defer span.End()

	url := schema.Ref.GetUrl()

	if url.Fragment == expressionFragment {
		return nil, nil
	}

	schemaType, err := getSubSchemaType(schema.RefSchema)
	if err != nil {
		return nil, err
	}

	// make a new topic based on the ref URL
	name, err := objectTypeOf(url)
	if err != nil {
		return nil, err
	}

	group, err := groupOf(url)
	if err != nil {
		return nil, err
	}

	version, err := versionOf(url)
	if err != nil {
		return nil, err
	}

	isResource := isResource(url)

	// produce a usable struct name:
	structReference := astmodel.NewStructReference(
		scanner.idFactory.CreateIdentifier(name),
		scanner.idFactory.CreateGroupName(group),
		scanner.idFactory.CreatePackageNameFromVersion(version),
		isResource)

	if schemaType == Object {
		// see if we already generated a struct for this ref
		// TODO: base this on URL?
		if definition, ok := scanner.FindDefinition(structReference.DefinitionName); ok {
			return definition.Reference(), nil
		}

		// Add a placeholder to avoid recursive calls
		sd := astmodel.NewStructDefinition(structReference, astmodel.NewStructType())
		scanner.AddDefinition(sd)
	}

	result, err := scanner.RunHandler(ctx, schemaType, schema.RefSchema)
	if err != nil {
		return nil, err
	}

	// if we got back a struct type, give it a name
	// (i.e. emit it as a "type X struct {}")
	// and return that instead
	if structType, ok := result.(*astmodel.StructType); ok {

		description := "Generated from: " + url.String()

		sd := astmodel.NewStructDefinition(structReference, structType).WithDescription(&description)

		// this will overwrite placeholder added above
		scanner.AddDefinition(sd)

		// Add any further definitions related to this
		relatedDefinitions := sd.StructType.CreateRelatedDefinitions(structReference.PackageReference, structReference.Name(), scanner.idFactory)
		for _, d := range relatedDefinitions {
			scanner.AddDefinition(d)
		}

		return sd.Reference(), nil
	}

	return result, err
}

func allOfHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "allOfHandler")
	defer span.End()

	var fields []*astmodel.FieldDefinition
	for _, all := range schema.AllOf {

		d, err := scanner.RunHandlerForSchema(ctx, all)
		if err != nil {
			return nil, err
		}

		if d == nil {
			continue // ignore skipped types
		}

		// unpack the contents of what we got from subhandlers:
		switch s := d.(type) {
		case *astmodel.StructType:
			// if it's a struct type get all its fields:
			fields = append(fields, s.Fields()...)

		case *astmodel.StructReference:
			// if it's a reference to a defined struct, embed it inside:
			fields = append(fields, astmodel.NewEmbeddedStructDefinition(s))

		default:
			klog.Errorf("Unhandled type in allOf: %#v\n", d)
		}
	}

	result := astmodel.NewStructType(fields...)
	return result, nil
}

func oneOfHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "oneOfHandler")
	defer span.End()

	// make sure we visit everything before bailing out,
	// to get all types generated even if we can't use them
	var results []astmodel.Type
	for _, one := range schema.OneOf {
		result, err := scanner.RunHandlerForSchema(ctx, one)
		if err != nil {
			return nil, err
		}

		if result != nil {
			results = append(results, result)
		}
	}

	if len(results) == 1 {
		return results[0], nil
	}

	// bail out, can't handle this yet:
	return astmodel.AnyType, nil
}

func anyOfHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "anyOfHandler")
	defer span.End()

	// again, make sure we walk everything first
	// to generate types:
	var results []astmodel.Type
	for _, any := range schema.AnyOf {
		result, err := scanner.RunHandlerForSchema(ctx, any)
		if err != nil {
			return nil, err
		}

		if result != nil {
			results = append(results, result)
		}
	}

	if len(results) == 1 {
		return results[0], nil
	}

	// return all possibilities...
	return astmodel.AnyType, nil
}

func arrayHandler(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
	ctx, span := tab.StartSpan(ctx, "arrayHandler")
	defer span.End()

	if len(schema.ItemsChildren) > 1 {
		return nil, fmt.Errorf("item contains more children than expected: %v", schema.ItemsChildren)
	}

	if len(schema.ItemsChildren) == 0 {
		// there is no type to the elements, so we must assume interface{}
		klog.Warning("Interface assumption unproven\n")

		return astmodel.NewArrayType(astmodel.AnyType), nil
	}

	// get the only child type and wrap it up as an array type:

	onlyChild := schema.ItemsChildren[0]

	astType, err := scanner.RunHandlerForSchema(ctx, onlyChild)
	if err != nil {
		return nil, err
	}

	return astmodel.NewArrayType(astType), nil
}

func getSubSchemaType(schema *gojsonschema.SubSchema) (SchemaType, error) {
	// handle special nodes:
	switch {
	case schema.Enum != nil: // this should come before the primitive checks below
		return Enum, nil
	case schema.OneOf != nil:
		return OneOf, nil
	case schema.AllOf != nil:
		return AllOf, nil
	case schema.AnyOf != nil:
		return AnyOf, nil
	case schema.RefSchema != nil:
		return Ref, nil
	}

	if schema.Types.IsTyped() {
		for _, t := range []SchemaType{Object, String, Number, Int, Bool, Array} {
			if schema.Types.Contains(string(t)) {
				return t, nil
			}
		}
	}

	// TODO: this whole switch is a bit wrong because type: 'object' can
	// be combined with OneOf/AnyOf/etc. still, it works okay for now...
	if !schema.Types.IsTyped() && schema.PropertiesChildren != nil {
		// no type but has properties, treat it as an object
		return Object, nil
	}

	return Unknown, &UnknownSchemaError{Schema: schema}
}

func getPrimitiveType(name SchemaType) (*astmodel.PrimitiveType, error) {
	switch name {
	case String:
		return astmodel.StringType, nil
	case Int:
		return astmodel.IntType, nil
	case Number:
		return astmodel.FloatType, nil
	case Bool:
		return astmodel.BoolType, nil
	default:
		return astmodel.AnyType, fmt.Errorf("%s is not a simple type and no ast.NewIdent can be created", name)
	}
}

func isURLPathSeparator(c rune) bool {
	return c == '/'
}

// Extract the name of an object from the supplied schema URL
func objectTypeOf(url *url.URL) (string, error) {
	fragmentParts := strings.FieldsFunc(url.Fragment, isURLPathSeparator)

	return fragmentParts[len(fragmentParts)-1], nil
}

// Extract the 'group' (here filename) of an object from the supplied schemaURL
func groupOf(url *url.URL) (string, error) {
	pathParts := strings.FieldsFunc(url.Path, isURLPathSeparator)

	file := pathParts[len(pathParts)-1]
	if !strings.HasSuffix(file, ".json") {
		return "", fmt.Errorf("Unexpected URL format (doesn't point to .json file)")
	}

	return strings.TrimSuffix(file, ".json"), nil
}

func isResource(url *url.URL) bool {
	fragmentParts := strings.FieldsFunc(url.Fragment, isURLPathSeparator)

	for _, fragmentPart := range fragmentParts {
		if fragmentPart == "resourceDefinitions" {
			return true
		}
	}

	return false
}

var versionRegex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

// Extract the name of an object from the supplied schema URL
func versionOf(url *url.URL) (string, error) {
	pathParts := strings.FieldsFunc(url.Path, isURLPathSeparator)

	for _, p := range pathParts {
		if versionRegex.MatchString(p) {
			return p, nil
		}
	}

	// No version found, that's fine
	return "", nil
}
