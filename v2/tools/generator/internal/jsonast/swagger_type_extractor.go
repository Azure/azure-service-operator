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
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/names"
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
	if extractor.swagger.Paths == nil {
		// No paths, nothing to extract
		return nil
	}

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
		operationPaths := extractor.expandAndCanonicalizePath(ctx, rawOperationPath, scanner, fullPutParameters)

		for _, operationPath := range operationPaths {
			err := extractor.extractOneResourceType(
				ctx,
				scanner,
				result,
				specSchema,
				statusSchema,
				nameParameterType,
				operationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (extractor *SwaggerTypeExtractor) extractOneResourceType(
	ctx context.Context,
	scanner *SchemaScanner,
	result SwaggerTypes,
	specSchema *Schema,
	statusSchema *Schema,
	nameParameterType astmodel.Type,
	operationPath string) error {

	armType, resourceName, err := extractor.resourceNameFromOperationPath(operationPath)
	if err != nil {
		klog.Errorf("Error extracting resource name (%s): %s", extractor.swaggerPath, err.Error())
		return nil
	}

	shouldPrune, because := scanner.configuration.ShouldPrune(resourceName)
	if shouldPrune == config.Prune {
		klog.V(3).Infof("Skipping %s because %s", resourceName, because)
		return nil
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
		return nil
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
			return nil
		}
		result.ResourceDefinitions[resourceName] = ResourceDefinition{
			SourceFile: extractor.swaggerPath,
			SpecType:   resourceSpec,
			StatusType: resourceStatus,
			ARMType:    armType,
			ARMURI:     operationPath,
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

// expandAndCanonicalizePath expands enums in the path. There are a number of cases here which make enum expansion and
// subsequent resource name extraction tricky:
//  1. Not every "resource" follows the standard ARM ID pattern of "/resourceType/{resourceName}/childResourceType/{childResourceName}".
//     For example see https://github.com/Azure/azure-rest-api-specs/blob/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2014-04-01/cosmos-db.json#L3195.
//     This API has a PUT "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}"
//     where /apis/gremlin looks like it should be the name of a resource called "apis", but in fact it has a different
//     body payload than apis/sql.
//     There are other examples of this as well, such as Web.Sites config: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/web/resource-manager/Microsoft.Web/stable/2020-12-01/WebApps.json#L1797
//  2. While most enums are for "names", some enums are for "types", see for example PrivateDNSZone RecordSets:
//     https://github.com/Azure/azure-rest-api-specs/blob/main/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json#L744.
//     This is the primary reason why this method must return a collection of paths, as a single URL may actually refer to
//     multiple resources. We want to make sure to expand these enums.
//  3. Many services have enums for resource names with a single value "default". In some cases, some paths for this
//     resource may use a parameter of type enum with a single value "default", while other usages in the same Swagger
//     put default directly into the path. This makes it important to expand these values, to ensure that resource
//     ownership ends up working later.
//     Ref: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/blob.json#L58
//  4. There are many other instances of single-value enums used for names that aren't "default". It just so happens that
//     service teams that use this pattern seem to consistently use either a hardcoded name or a parameter. They don't seem
//     to mix/match like the storage example in #3 above.
//     Ref: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADAdministrators.json#L61
//
// The above realities result in the following algorithm:
//   - Expand all enums that are in the resourceType section of the path (case 2 above)
//   - Expand all enums with a single value "default", to protect against hardcoding of "default" in other URL paths (case 3 above)
//   - Don't expand any other enums, which protects us from considering those parameters as part of the resource type and
//     including them in the resource name later
func (extractor *SwaggerTypeExtractor) expandAndCanonicalizePath(
	ctx context.Context,
	operationPath string,
	scanner *SchemaScanner,
	parameters []spec.Parameter) []string {

	var results = []string{
		operationPath,
	}

	_, subpath, err := extractor.extractResourceSubpath(operationPath)
	if err != nil {
		// This error is safe to ignore in this case as it means that we can't parse the path as a resource.
		// Just return the raw path so we do further processing and log a message
		klog.Errorf("Error expanding enums in path (%s): %s", extractor.swaggerPath, err.Error())
		return results
	}

	paramMap := make(map[string]spec.Parameter)
	for _, param := range parameters {
		_, resolved := extractor.fullyResolveParameter(param)

		paramMap[resolved.Name] = resolved
	}

	split := strings.Split(subpath, "/")
	// Loop over split to ensure we're processing the url paths left to right, which is important
	// for counting which parameters correspond to a "name" and which correspond to a "type"
	for i, urlPart := range split {
		if len(urlPart) == 0 {
			// skip empty parts
			continue
		}

		if urlPart[0] != '{' {
			// Not a parameter, so by definition doesn't need to be expanded
			continue
		}

		paramName := strings.Trim(urlPart, "{}")
		// Note that Swagger requires a case-sensitive match, so we don't need to worry about case insensitivity here
		param, ok := paramMap[paramName]
		if !ok {
			// This shouldn't happen, but if it does the best bet is to just continue as without the actual
			// parameter there's no way to perform substitution
			continue
		}

		var err error
		var paramType astmodel.Type
		schema := extractor.schemaFromParameter(param)

		if schema == nil {
			panic(fmt.Sprintf("no schema generated for parameter %s in path %q", param.Name, operationPath))
		}

		paramType, err = scanner.RunHandlerForSchema(ctx, *schema)
		if err != nil {
			panic(err)
		}

		enum, isEnum := paramType.(*astmodel.EnumType)
		if !isEnum {
			// We only expand enums, skip everything else
			continue
		}

		if len(enum.Options()) == 0 {
			// This shouldn't happen but guard against it anyway
			continue
		}

		if !astmodel.TypeEquals(enum.BaseType(), astmodel.StringType) {
			// We only expand string enums
			continue
		}

		// resource types come on the even indices after the group.
		// Example: Microsoft.Group/resourceType/{resourceName}/{childResourceTypeParameter}/{childName},
		// has "resourceType" at index=0, "{childResourceTypeParameter}" at index=2
		isResourceTypeParameter := i%2 == 0
		isDefaultEnum := false
		if len(enum.Options()) == 1 {
			opt := enum.Options()[0]
			theOption := strings.Trim(opt.Value, "\"")
			isDefaultEnum = strings.EqualFold(theOption, "default")
		}

		if !isResourceTypeParameter && !isDefaultEnum {
			continue
		}

		toReplace := fmt.Sprintf("{%s}", param.Name)
		var expanded []string

		for _, path := range results {
			for _, opt := range enum.Options() {
				theOption := strings.Trim(opt.Value, "\"")
				expanded = append(expanded, strings.ReplaceAll(path, toReplace, theOption))
			}
		}

		results = expanded
	}

	return results
}

func (extractor *SwaggerTypeExtractor) resourceNameFromOperationPath(operationPath string) (string, astmodel.TypeName, error) {
	group, resource, name, err := extractor.inferNameFromURLPath(operationPath)
	if err != nil {
		return "", astmodel.EmptyTypeName, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
	}

	return group + "/" + resource, astmodel.MakeTypeName(extractor.outputPackage, name), nil
}

// extractResourceSubpath gets the subpath focused on the actual resource.
// For example: “…/Microsoft.GroupName/resourceType/{parameterId}/differentType/{otherId}/something/{moreId}”
// would return "Microsoft.GroupName", "resourceType/{parameterId}/differentType/{otherId}/something/{moreId}"
func (extractor *SwaggerTypeExtractor) extractResourceSubpath(operationPath string) (string, string, error) {
	urlParts := strings.Split(operationPath, "/")
	for i, urlPart := range urlParts {
		if len(urlPart) == 0 {
			// skip empty parts
			continue
		}

		if SwaggerGroupRegex.MatchString(urlPart) {
			return urlPart, strings.Join(urlParts[i+1:], "/"), nil
		}
	}

	return "", "", errors.Errorf("no group name (‘Microsoft…’) found in %s", operationPath)
}

// inferNameFromURLPath attempts to extract a name from a Swagger operation path
// for example “…/Microsoft.GroupName/resourceType/{resourceId}” would result
// in the name “ResourceType”. Child resources are treated by converting (e.g.)
// “…/Microsoft.GroupName/resourceType/{parameterId}/differentType/{otherId}/something/{moreId}”
// to “ResourceTypeDifferentTypeSomething”.
func (extractor *SwaggerTypeExtractor) inferNameFromURLPath(operationPath string) (string, string, string, error) {
	group := ""
	var nameParts []string

	if strings.EqualFold(operationPath, "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}") {
		return "Microsoft.Resources", "resourceGroups", "ResourceGroup", nil
	}

	group, subpath, err := extractor.extractResourceSubpath(operationPath)
	if err != nil {
		return "", "", "", err
	}

	urlParts := strings.Split(subpath, "/")
	skippedLast := false
	for _, urlPart := range urlParts {
		if len(urlPart) == 0 {
			// skip empty parts
			continue
		}

		// If default was defined as an enum in the Swagger, this check is not needed as it wouldn't have been expanded by
		// the expandEnumsInPath method, but some services hardcode default into their URLs like : blobService/default/containers/{containerName}
		// and we don't want the "default" name to be part of the resource type name for those cases so we ignore it here.
		if strings.EqualFold(urlPart, "default") {
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
	nameParts[len(nameParts)-1] = names.Singularize(nameParts[len(nameParts)-1])

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
				Name:        getUnusedParameterName(params),
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

func getUnusedParameterName(params []spec.Parameter) string {
	name := "name" // Default to name if we can
	clashes := 0
	for _, param := range params {
		if param.Name == name {
			clashes += 1
			name = fmt.Sprintf("name%d", clashes)
		}
	}

	return name
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
