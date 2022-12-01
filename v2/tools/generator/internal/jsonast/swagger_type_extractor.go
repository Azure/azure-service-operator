/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

type SwaggerTypeExtractor struct {
	idFactory   astmodel.IdentifierFactory
	config      *config.Configuration
	cache       CachingFileLoader
	swagger     spec.Swagger
	swaggerPath string
	// package for output types (e.g. Microsoft.Network.Frontdoor/v20200101)
	outputPackage astmodel.LocalPackageReference
}

// NewSwaggerTypeExtractor creates a new SwaggerTypeExtractor
func NewSwaggerTypeExtractor(
	config *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	swagger spec.Swagger,
	swaggerPath string,
	outputPackage astmodel.LocalPackageReference,
	cache CachingFileLoader,
) SwaggerTypeExtractor {
	return SwaggerTypeExtractor{
		idFactory:     idFactory,
		swaggerPath:   swaggerPath,
		swagger:       swagger,
		outputPackage: outputPackage,
		cache:         cache,
		config:        config,
	}
}

type SwaggerTypes struct {
	ResourceDefinitions ResourceDefinitionSet
	OtherDefinitions    astmodel.TypeDefinitionSet
}

type ResourceDefinitionSet map[astmodel.TypeName]ResourceDefinition

type ResourceDefinition struct {
	SpecType   astmodel.Type
	StatusType astmodel.Type
	SourceFile string
	ARMType    string // e.g. Microsoft.XYZ/resourceThings
	ARMURI     string
	// TODO: use ARMURI for generating Resource URIs (only used for documentation & ownership at the moment)
}

// ExtractTypes finds all operations in the Swagger spec that
// have a PUT verb and a path like "Microsoft.GroupName/…/resourceName/{resourceId}",
// and extracts the types for those operations, into the 'resourceTypes' result.
// Any additional types required by the resource types are placed into the 'otherTypes' result.
func (extractor *SwaggerTypeExtractor) ExtractTypes(ctx context.Context) (SwaggerTypes, error) {
	result := SwaggerTypes{
		ResourceDefinitions: make(ResourceDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	scanner := NewSchemaScanner(extractor.idFactory, extractor.config)

	err := extractor.ExtractResourceTypes(ctx, scanner, result)
	if err != nil {
		return SwaggerTypes{}, errors.Wrap(err, "error extracting resource types")
	}

	err = extractor.ExtractOneOfTypes(ctx, scanner, result)
	if err != nil {
		return SwaggerTypes{}, errors.Wrap(err, "error extracting one-of option types")
	}

	for _, def := range scanner.Definitions() {
		// Add additional type definitions required by the resources
		if existingDef, ok := result.OtherDefinitions[def.Name()]; ok {
			if !astmodel.TypeEquals(existingDef.Type(), def.Type()) {
				return SwaggerTypes{}, errors.Errorf("type already defined differently: %s\nwas %s is %s\ndiff:\n%s",
					def.Name(),
					existingDef.Type(),
					def.Type(),
					astmodel.DiffTypes(existingDef.Type(), def.Type()))
			}
		} else {
			result.OtherDefinitions.Add(def)
		}
	}

	return result, nil
}

func (extractor *SwaggerTypeExtractor) ExtractResourceTypes(ctx context.Context, scanner *SchemaScanner, result SwaggerTypes) error {
	for rawOperationPath, op := range extractor.swagger.Paths.Paths {
		// a resource must have both PUT and GET
		if op.Put == nil || op.Get == nil {
			continue
		}

		specSchema, statusSchema, ok := extractor.findARMResourceSchema(op, rawOperationPath)
		if !ok {
			// klog.Warningf("No ARM schema found for %s in %q", rawOperationPath, filePath)
			continue
		}

		// Parameters may be defined at the URL level or the individual verb level (PUT, GET, etc). The union is the final parameter set
		// for the operation
		fullPutParameters := append(op.Parameters, op.Put.Parameters...)
		nameParameterType := extractor.getNameParameterType(ctx, rawOperationPath, scanner, fullPutParameters)

		armType, resourceName, err := extractor.resourceNameFromOperationPath(rawOperationPath)
		if err != nil {
			klog.Errorf("Error extracting resource name (%s): %s", extractor.swaggerPath, err.Error())
			continue
		}

		shouldPrune, because := scanner.configuration.ShouldPrune(resourceName)
		if shouldPrune == config.Prune {
			klog.V(3).Infof("Skipping %s because %s", resourceName, because)
			continue
		}

		var resourceSpec astmodel.Type
		if specSchema == nil {
			// nil indicates empty body
			resourceSpec = astmodel.NewObjectType()
		} else {
			resourceSpec, err = scanner.RunHandlerForSchema(ctx, *specSchema)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}

				return errors.Wrapf(err, "unable to produce spec type for resource %s", resourceName)
			}
		}

		if resourceSpec == nil {
			// this indicates a type filtered out by RunHandlerForSchema, skip
			continue
		}

		// Use the name field documented on the PUT parameter as it's likely to have better
		// validation attached to it than the name specified in the request body (which is usually readonly and has no validation).
		// If by some chance the body property has validation, the AllOf below will merge it anyway, so it will not be lost.
		nameProperty := astmodel.NewPropertyDefinition(astmodel.NameProperty, "name", nameParameterType)
		resourceSpec = astmodel.NewAllOfType(resourceSpec, astmodel.NewObjectType().WithProperty(nameProperty))

		var resourceStatus astmodel.Type
		if statusSchema == nil {
			// nil indicates empty body
			resourceStatus = astmodel.NewObjectType()
		} else {
			resourceStatus, err = scanner.RunHandlerForSchema(ctx, *statusSchema)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}

				return errors.Wrapf(err, "unable to produce status type for resource %s", resourceName)
			}
		}

		if existingResource, ok := result.ResourceDefinitions[resourceName]; ok {
			// TODO: check status types as well
			if !astmodel.TypeEquals(existingResource.SpecType, resourceSpec) {
				return errors.Errorf("resource already defined differently: %s\ndiff:\n%s",
					resourceName,
					astmodel.DiffTypes(existingResource.SpecType, resourceSpec))
			}
		} else {
			if resourceSpec != nil && resourceStatus == nil {
				fmt.Printf("generated nil resourceStatus for %s\n", resourceName)
				continue
			}
			result.ResourceDefinitions[resourceName] = ResourceDefinition{
				SourceFile: extractor.swaggerPath,
				SpecType:   resourceSpec,
				StatusType: resourceStatus,
				ARMType:    armType,
				ARMURI:     rawOperationPath,
			}
		}
	}

	return nil
}

// ExtractOneOfTypes ensures we haven't missed any of the required OneOf type definitions.
// The depth-first search of the Swagger spec done by ExtractResourcetypes() won't have found any "loose" one of options
// so we need this extra step.
func (extractor *SwaggerTypeExtractor) ExtractOneOfTypes(
	ctx context.Context,
	scanner *SchemaScanner,
	result SwaggerTypes,
) error {
	var errs []error
	for name, def := range extractor.swagger.Definitions {
		// Looking for definitions that either
		//  o  specify discriminator property
		//  o  contains a 'x-ms-discriminator-value' extension value
		if _, found := def.Extensions.GetString("x-ms-discriminator-value"); !found && def.Discriminator == "" {
			continue
		}

		// Create a wrapper type for this definition
		schema := MakeOpenAPISchema(
			name,
			def,
			extractor.swaggerPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)

		// Run a handler to generate our type
		t, err := scanner.RunHandlerForSchema(ctx, schema)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "unable to produce type for definition %s", name))
			continue
		}

		if t == nil {
			// No type created
			continue
		}

		typeName := astmodel.MakeTypeName(extractor.outputPackage, name)

		if action, _ := scanner.configuration.ShouldPrune(typeName); action == config.Prune {
			// Pruned type
			continue
		}

		result.OtherDefinitions[typeName] = astmodel.MakeTypeDefinition(typeName, t)
	}

	return nil
}

// Look at the responses of the PUT to determine if this represents an ARM resource,
// and if so, return the schema for it.
// see: https://github.com/Azure/autorest/issues/1936#issuecomment-286928591
func (extractor *SwaggerTypeExtractor) findARMResourceSchema(op spec.PathItem, rawOperationPath string) (*Schema, *Schema, bool) {
	// to decide if something is a resource, we must look at the GET responses
	isResource := false

	var foundStatus *Schema
	if op.Get.Responses != nil {
		for statusCode, response := range op.Get.Responses.StatusCodeResponses {
			// only check OK and Created (per above linked comment)
			// TODO: we should really check that the results are the same in each status result
			if statusCode == 200 || statusCode == 201 {
				schema, ok := extractor.doesResponseRepresentARMResource(response, rawOperationPath)
				if ok {
					foundStatus = schema
					isResource = true
					break
				}
			}
		}
	}

	if !isResource {
		return nil, nil, false
	}

	var foundSpec *Schema

	params := op.Put.Parameters
	if op.Parameters != nil {
		klog.Warningf("overriding parameters for %q in %s", rawOperationPath, extractor.swaggerPath)
		params = op.Parameters
	}

	// the actual Schema must come from the PUT parameters
	noBody := true
	for _, param := range params {
		inBody := param.In == "body"
		_, innerParam := extractor.fullyResolveParameter(param)
		inBody = inBody || innerParam.In == "body"

		// note: bug avoidance: we must not pass innerParam to schemaFromParameter
		// since it will treat it as relative to the current file, which might not be correct.
		// instead, schemaFromParameter must re-fully-resolve the parameter so that
		// it knows it comes from another file (if it does).

		if inBody { // must be a (the) body parameter
			noBody = false
			result := extractor.schemaFromParameter(param)
			if result != nil {
				foundSpec = result
				break
			}
		}
	}

	if foundSpec == nil {
		if noBody {
			klog.V(3).Infof("Empty body for %s", rawOperationPath)
		} else {
			klog.Warningf("Response indicated that type was ARM resource but no schema found for %s in %q", rawOperationPath, extractor.swaggerPath)
			return nil, nil, false
		}
	}

	return foundSpec, foundStatus, true
}

// fullyResolveParameter resolves the parameter and returns the file that contained the parameter and the parameter
func (extractor *SwaggerTypeExtractor) fullyResolveParameter(param spec.Parameter) (string, spec.Parameter) {
	if param.Ref.GetURL() == nil {
		// it is not a $ref parameter, we already have it
		return extractor.swaggerPath, param
	}

	paramPath, param := loadRefParameter(param.Ref, extractor.swaggerPath, extractor.cache)
	if param.Ref.GetURL() == nil {
		return paramPath, param
	}

	return extractor.fullyResolveParameter(param)
}

func (extractor *SwaggerTypeExtractor) schemaFromParameter(param spec.Parameter) *Schema {
	paramPath, param := extractor.fullyResolveParameter(param)

	var result Schema
	if param.Schema == nil {
		// We're dealing with a simple schema here
		if param.SimpleSchema.Type == "" {
			klog.Warningf("schemaFromParameter invoked on parameter without schema or simpleschema")
			return nil
		}

		result = MakeOpenAPISchema(
			param.Name,
			*makeSchemaFromSimpleSchemaParam(schemaAndValidations{param.SimpleSchema, param.CommonValidations}),
			paramPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)
	} else {
		result = MakeOpenAPISchema(
			nameFromRef(param.Schema.Ref),
			*param.Schema,
			paramPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)
	}

	return &result
}

func (extractor *SwaggerTypeExtractor) doesResponseRepresentARMResource(response spec.Response, rawOperationPath string) (*Schema, bool) {
	// the schema can either be directly included
	if response.Schema != nil {

		schema := *response.Schema
		result := MakeOpenAPISchema(
			nameFromRef(schema.Ref),
			*response.Schema,
			extractor.swaggerPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)

		return &result, isMarkedAsARMResource(result)
	}

	// or it can be under a $ref
	ref := response.Ref
	if ref.GetURL() != nil {
		refFilePath, refSchema, packageAndSwagger := loadRefSchema(ref, extractor.swaggerPath, extractor.cache)
		outputPackage := extractor.outputPackage
		if packageAndSwagger.Package != nil {
			outputPackage = *packageAndSwagger.Package
		}

		schema := MakeOpenAPISchema(
			nameFromRef(ref),
			refSchema,
			refFilePath,
			outputPackage,
			extractor.idFactory,
			extractor.cache)

		return &schema, isMarkedAsARMResource(schema)
	}

	klog.Warningf("Unable to locate schema on response for %q in %s", rawOperationPath, extractor.swaggerPath)
	return nil, false
}

// a Schema represents an ARM resource if it (or anything reachable via $ref or AllOf)
// is marked with x-ms-azure-resource, or if it (or anything reachable via $ref or AllOf)
// has each of the properties: id,name,type, and they are all readonly and required.
// see: https://github.com/Azure/autorest/issues/1936#issuecomment-286928591
// and: https://github.com/Azure/autorest/issues/2127
func isMarkedAsARMResource(schema Schema) bool {
	hasID := false
	idRequired := false
	hasName := false
	nameRequired := false
	hasType := false
	typeRequired := false

	var recurse func(schema Schema) bool
	recurse = func(schema Schema) bool {
		if schema.extensionAsBool("x-ms-azure-resource") {
			return true
		}

		props := schema.properties()
		if !hasID {
			if idProp, ok := props["id"]; ok {
				if idProp.hasType("string") && idProp.readOnly() {
					hasID = true
				}
			}
		}

		if !idRequired {
			idRequired = slices.Contains(schema.requiredProperties(), "id")
		}

		if !hasName {
			if nameProp, ok := props["name"]; ok {
				if nameProp.hasType("string") && nameProp.readOnly() {
					hasName = true
				}
			}
		}

		if !nameRequired {
			nameRequired = slices.Contains(schema.requiredProperties(), "name")
		}

		if !hasType {
			if typeProp, ok := props["type"]; ok {
				if typeProp.hasType("string") && typeProp.readOnly() {
					hasType = true
				}
			}
		}

		if !typeRequired {
			typeRequired = slices.Contains(schema.requiredProperties(), "type")
		}

		if schema.isRef() {
			return recurse(schema.refSchema())
		}

		if schema.hasAllOf() {
			for _, allOf := range schema.allOf() {
				if recurse(allOf) {
					return true
				}
			}
		}

		return hasID && hasName && hasType
	}

	return recurse(schema)
}

func (extractor *SwaggerTypeExtractor) getNameParameterType(
	ctx context.Context,
	operationPath string,
	scanner *SchemaScanner,
	parameters []spec.Parameter) astmodel.Type {

	lastParam, ok := extractor.extractLastPathParam(operationPath, parameters)
	if !ok {
		panic(fmt.Sprintf("couldn't find path parameter for %s", operationPath))
	}

	// non-enum parameter
	var err error
	var paramType astmodel.Type
	schema := extractor.schemaFromParameter(lastParam)
	if schema == nil {
		panic(fmt.Sprintf("no schema generated for parameter %s in path %q", lastParam.Name, operationPath))
	}

	paramType, err = scanner.RunHandlerForSchema(ctx, *schema)
	if err != nil {
		panic(err)
	}

	return paramType
}

func (extractor *SwaggerTypeExtractor) resourceNameFromOperationPath(operationPath string) (string, astmodel.TypeName, error) {
	group, resource, name, err := extractor.inferNameFromURLPath(operationPath)
	if err != nil {
		return "", astmodel.EmptyTypeName, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
	}

	return group + "/" + resource, astmodel.MakeTypeName(extractor.outputPackage, name), nil
}

// inferNameFromURLPath attempts to extract a name from a Swagger operation path
// for example “…/Microsoft.GroupName/resourceType/{resourceId}” would result
// in the name “ResourceType”. Child resources are treated by converting (e.g.)
// “…/Microsoft.GroupName/resourceType/{parameterId}/differentType/{otherId}/something/{moreId}”
// to “ResourceTypeDifferentTypeSomething”.
func (extractor *SwaggerTypeExtractor) inferNameFromURLPath(operationPath string) (string, string, string, error) {
	group := ""
	var nameParts []string

	urlParts := strings.Split(operationPath, "/")
	reading := false
	skippedLast := false
	for _, urlPart := range urlParts {
		if len(urlPart) == 0 {
			// skip empty parts
			continue
		}

		if reading {
			if urlPart == "default" {
				// skip; shouldn’t be part of name
				// TODO: I haven’t yet found where this is done in autorest/autorest.armresource to document this
			} else if urlPart[0] == '{' {
				// this is a url parameter

				if skippedLast {
					// this means two {parameters} in a row
					return "", "", "", errors.Errorf("multiple parameters in path")
				}

				skippedLast = true
			} else {
				// normal part of path, uppercase first character
				nameParts = append(nameParts, urlPart)
				skippedLast = false
			}
		} else if SwaggerGroupRegex.MatchString(urlPart) {
			group = urlPart
			reading = true
		}
	}

	if !reading {
		return "", "", "", errors.Errorf("no group name (‘Microsoft…’ = %q) found", group)
	}

	if len(nameParts) == 0 {
		return "", "", "", errors.Errorf("couldn’t infer name")
	}

	resource := strings.Join(nameParts, "/") // capture this before uppercasing/singularizing

	// Uppercase first letter in each part:
	for ix := range nameParts {
		nameParts[ix] = strings.ToUpper(nameParts[ix][0:1]) + nameParts[ix][1:]
	}

	// Now singularize last part of name:
	nameParts[len(nameParts)-1] = astmodel.Singularize(nameParts[len(nameParts)-1], extractor.idFactory)

	name := strings.Join(nameParts, "_")

	return group, resource, name, nil
}

func nameFromRef(ref spec.Ref) string {
	url := ref.GetURL()
	if url == nil {
		return ""
	}

	parts := strings.Split(url.Fragment, "/")
	if len(parts) == 0 {
		return "'"
	}

	return parts[len(parts)-1]
}

// SwaggerGroupRegex matches a “group” (Swagger ‘namespace’)
// based on: https://github.com/Azure/autorest/blob/85de19623bdce3ccc5000bae5afbf22a49bc4665/core/lib/pipeline/metadata-generation.ts#L25
var SwaggerGroupRegex = regexp.MustCompile(`[Mm]icrosoft\.[^/\\]+`)

func (extractor *SwaggerTypeExtractor) extractLastPathParam(operationPath string, params []spec.Parameter) (spec.Parameter, bool) {
	// The parameter we want is the one following the final /
	split := strings.Split(operationPath, "/")
	rawLastParam := split[len(split)-1]

	if !strings.HasPrefix(rawLastParam, "{") {
		return spec.Parameter{
			SimpleSchema: spec.SimpleSchema{
				Type: "string",
			},
			ParamProps: spec.ParamProps{
				Name:        "name", // TODO: We don't know that this name isn't taken already
				In:          "path",
				Required:    true,
				Description: "The name",
			},
			CommonValidations: spec.CommonValidations{
				Enum: []interface{}{
					rawLastParam,
				},
			},
		}, true
	}

	lastParamStr := strings.Trim(rawLastParam, "{}")

	var lastParam spec.Parameter
	var found bool

	for _, param := range params {
		_, resolved := extractor.fullyResolveParameter(param)
		if resolved.In != "path" || resolved.Name != lastParamStr {
			continue
		}

		lastParam = resolved
		found = true
	}

	return lastParam, found
}

type schemaAndValidations struct {
	spec.SimpleSchema
	spec.CommonValidations
}

func makeSchemaFromSimpleSchemaParam(param schemaAndValidations) *spec.Schema {
	if param.SimpleSchema.Type == "" {
		panic("cannot make schema from simple schema for non-simple-schema param")
	}
	var itemsSchema *spec.Schema
	if param.Items != nil {
		// TODO: Possibly need to consume CollectionFormat here
		itemsSchema = makeSchemaFromSimpleSchemaParam(schemaAndValidations{param.Items.SimpleSchema, param.Items.CommonValidations})
	}

	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type:    spec.StringOrArray{param.SimpleSchema.Type},
			Default: param.SimpleSchema.Default,
			Format:  param.SimpleSchema.Format,
			Items: &spec.SchemaOrArray{
				Schema: itemsSchema,
			},
			Nullable: param.SimpleSchema.Nullable,
		},
	}
	schema = schema.WithValidations(param.Validations())

	return schema
}
