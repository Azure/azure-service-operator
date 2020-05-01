/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"fmt"
	"log"
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
		Structs      map[string]*astmodel.StructDefinition
		TypeHandlers map[SchemaType]TypeHandler
		Filters      []string
		idFactory    astmodel.IdentifierFactory
	}
)

// FindStruct looks to see if we have seen the specified struct before, returning its definition if we have.
func (scanner *SchemaScanner) FindStruct(name string, version string) (*astmodel.StructDefinition, bool) {
	key := name + "/" + version
	result, ok := scanner.Structs[key]
	return result, ok
}

// AddStruct makes a record of the specified struct so that FindStruct() can return it when it is needed again.
func (scanner *SchemaScanner) AddStruct(structDefinition *astmodel.StructDefinition) {
	key := structDefinition.Name() + "/" + structDefinition.Version()
	scanner.Structs[key] = structDefinition
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
		Structs:      make(map[string]*astmodel.StructDefinition),
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
	rootStructVersion, err := versionOf(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract version for schema: %w", err)
	}

	nodes, err := scanner.RunHandler(ctx, schemaType, schema)
	if err != nil {
		return nil, err
	}

	// TODO: make safer:
	structName := scanner.idFactory.CreateIdentifier(rootStructName)
	root := astmodel.NewStructDefinition(structName, rootStructVersion, nodes.(*astmodel.StructType).Fields()...)
	description := "Generated from: " + url.String()
	root = root.WithDescription(&description)

	scanner.AddStruct(root)

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
	ctx, span := tab.StartSpan(ctx, "enumHandler")
	defer span.End()

	// if there is an underlying primitive type, return that
	for _, t := range []SchemaType{Bool, Int, Number, String} {
		if schema.Types.Contains(string(t)) {
			return getPrimitiveType(t)
		}
	}

	// assume string
	return astmodel.StringType, nil

	//TODO Create an Enum field that captures the permitted options too
}

func fixedTypeHandler(typeToReturn astmodel.Type, handlerName string) TypeHandler {
	return func(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) (astmodel.Type, error) {
		ctx, span := tab.StartSpan(ctx, handlerName+"Handler")
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

	structDefinition := astmodel.NewStructType(fields)
	return structDefinition, nil
}

func getFields(ctx context.Context, scanner *SchemaScanner, schema *gojsonschema.SubSchema) ([]*astmodel.FieldDefinition, error) {
	ctx, span := tab.StartSpan(ctx, "getFields")
	defer span.End()

	var fields []*astmodel.FieldDefinition
	for _, prop := range schema.PropertiesChildren {
		schemaType, err := getSubSchemaType(prop)
		if _, ok := err.(*UnknownSchemaError); ok {
			// if we don't know the type, we still need to provide the property, we will just provide open interface
			fieldName := scanner.idFactory.CreateIdentifier(prop.Property)
			field := astmodel.NewFieldDefinition(fieldName, prop.Property, astmodel.AnyType).WithDescription(schema.Description)
			fields = append(fields, field)
			continue
		}

		if err != nil {
			return nil, err
		}

		propType, err := scanner.RunHandler(ctx, schemaType, prop)
		if _, ok := err.(*UnknownSchemaError); ok {
			// if we don't know the type, we still need to provide the property, we will just provide open interface
			fieldName := scanner.idFactory.CreateIdentifier(prop.Property)
			field := astmodel.NewFieldDefinition(fieldName, prop.Property, astmodel.AnyType).WithDescription(schema.Description)
			fields = append(fields, field)
			continue
		}

		if err != nil {
			return nil, err
		}

		fieldName := scanner.idFactory.CreateIdentifier(prop.Property)
		field := astmodel.NewFieldDefinition(fieldName, prop.Property, propType).WithDescription(prop.Description)
		fields = append(fields, field)
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
		additionalPropsField := astmodel.NewFieldDefinition("additionalProperties", "additionalProperties", astmodel.NewStringMap(additionalPropsType))
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

	log.Printf("INF $ref to %s\n", url)

	schemaType, err := getSubSchemaType(schema.RefSchema)
	if err != nil {
		return nil, err
	}

	// make a new topic based on the ref URL
	name, err := objectTypeOf(url)
	if err != nil {
		return nil, err
	}

	version, err := versionOf(url)
	if err != nil {
		return nil, err
	}

	structName := scanner.idFactory.CreateIdentifier(name)
	if schemaType == Object {
		// see if we already generated a struct for this ref
		// TODO: base this on URL?
		if definition, ok := scanner.FindStruct(structName, version); ok {
			return &definition.StructReference, nil
		}

		// Add a placeholder to avoid recursive calls
		sd := astmodel.NewStructDefinition(structName, version)
		scanner.AddStruct(sd)
	}

	result, err := scanner.RunHandler(ctx, schemaType, schema.RefSchema)
	if err != nil {
		return nil, err
	}

	// if we got back a struct type, give it a name
	// (i.e. emit it as a "type X struct {}")
	// and return that instead
	if std, ok := result.(*astmodel.StructType); ok {

		description := "Generated from: " + url.String()

		sd := astmodel.NewStructDefinition(structName, version, std.Fields()...).WithDescription(&description)

		// this will overwrite placeholder added above
		scanner.AddStruct(sd)
		return &sd.StructReference, nil
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
		switch d.(type) {
		case *astmodel.StructType:
			// if it's a struct type get all its fields:
			s := d.(*astmodel.StructType)
			fields = append(fields, s.Fields()...)

		case *astmodel.StructReference:
			// if it's a reference to a struct type, embed it inside:
			s := d.(*astmodel.StructReference)
			fields = append(fields, astmodel.NewEmbeddedStructDefinition(s))

		default:
			log.Printf("Unhandled type in allOf: %T\n", d)
		}
	}

	result := astmodel.NewStructType(fields)
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
		log.Printf("WRN Interface assumption unproven\n")

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

func isPrimitiveType(name SchemaType) bool {
	switch name {
	case String, Int, Number, Bool:
		return true
	default:
		return false
	}
}

func asComment(text *string) string {
	if text == nil {
		return ""
	}

	return "// " + *text
}

// Extract the name of an object from the supplied schema URL
func objectTypeOf(url *url.URL) (string, error) {
	isPathSeparator := func(c rune) bool {
		return c == '/'
	}

	fragmentParts := strings.FieldsFunc(url.Fragment, isPathSeparator)

	return fragmentParts[len(fragmentParts)-1], nil
}

// Extract the name of an object from the supplied schema URL
func versionOf(url *url.URL) (string, error) {
	isPathSeparator := func(c rune) bool {
		return c == '/'
	}

	pathParts := strings.FieldsFunc(url.Path, isPathSeparator)
	versionRegex, err := regexp.Compile("\\d\\d\\d\\d-\\d\\d-\\d\\d")
	if err != nil {
		return "", fmt.Errorf("Invalid Regex format %w", err)
	}

	for _, p := range pathParts {
		if versionRegex.MatchString(p) {
			return p, nil
		}
	}

	// No version found, that's fine
	return "", nil
}
