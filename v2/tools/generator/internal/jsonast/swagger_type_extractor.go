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
	// TODO: use ARMURI for generating Resource URIs (only used for documentation at the moment)
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

		operationPaths, _ := extractor.expandEnumsInPath(ctx, rawOperationPath, scanner, op.Put.Parameters)
		for _, operationPath := range operationPaths {

			armType, resourceName, err := extractor.resourceNameFromOperationPath(operationPath)
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
						return SwaggerTypes{}, err
					}

					return SwaggerTypes{}, errors.Wrapf(err, "unable to produce spec type for resource %s", resourceName)
				}
			}

			if resourceSpec == nil {
				// this indicates a type filtered out by RunHandlerForSchema, skip
				continue
			}

			var resourceStatus astmodel.Type
			if statusSchema == nil {
				// nil indicates empty body
				resourceStatus = astmodel.NewObjectType()
			} else {
				resourceStatus, err = scanner.RunHandlerForSchema(ctx, *statusSchema)
				if err != nil {
					if errors.Is(err, context.Canceled) {
						return SwaggerTypes{}, err
					}

					return SwaggerTypes{}, errors.Wrapf(err, "unable to produce status type for resource %s", resourceName)
				}
			}

			if existingResource, ok := result.ResourceDefinitions[resourceName]; ok {
				// TODO: check status types as well
				if !astmodel.TypeEquals(existingResource.SpecType, resourceSpec) {
					return SwaggerTypes{}, errors.Errorf("resource already defined differently: %s\ndiff: %s",
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
					ARMURI:     operationPath,
				}
			}
		}
	}

	for _, def := range scanner.Definitions() {
		// now add in the additional type definitions required by the resources
		if existingDef, ok := result.OtherDefinitions[def.Name()]; ok {
			if !astmodel.TypeEquals(existingDef.Type(), def.Type()) {
				return SwaggerTypes{}, errors.Errorf("type already defined differently: %s\nwas %s is %s\ndiff: %s",
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
			klog.Warningf("Empty body for %s", rawOperationPath)
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

	paramPath, param, _ := loadRefParameter(param.Ref, extractor.swaggerPath, extractor.cache)
	if param.Ref.GetURL() == nil {
		return paramPath, param
	}

	return extractor.fullyResolveParameter(param)
}

func (extractor *SwaggerTypeExtractor) schemaFromParameter(param spec.Parameter) *Schema {
	paramPath, param := extractor.fullyResolveParameter(param)

	if param.Schema == nil {
		klog.Warningf("schemaFromParameter invoked on parameter without schema")
		return nil
	}

	result := MakeOpenAPISchema(
		*param.Schema, // all params marked as 'body' have schemas
		paramPath,
		extractor.outputPackage,
		extractor.idFactory,
		extractor.cache)

	return &result
}

func (extractor *SwaggerTypeExtractor) doesResponseRepresentARMResource(response spec.Response, rawOperationPath string) (*Schema, bool) {
	// the schema can either be directly included
	if response.Schema != nil {

		schema := MakeOpenAPISchema(
			*response.Schema,
			extractor.swaggerPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)

		return &schema, isMarkedAsARMResource(schema)
	}

	// or it can be under a $ref
	if response.Ref.GetURL() != nil {

		refFilePath, refSchema, pkg := loadRefSchema(response.Ref, extractor.swaggerPath, extractor.cache)
		outputPackage := extractor.outputPackage
		if pkg != nil {
			outputPackage = *pkg
		}

		schema := MakeOpenAPISchema(
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
		if schema.extensions("x-ms-azure-resource") == true {
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

func (extractor *SwaggerTypeExtractor) expandEnumsInPath(ctx context.Context, operationPath string, scanner *SchemaScanner, parameters []spec.Parameter) ([]string, map[string]astmodel.Type) {
	results := []string{operationPath}
	urlParams := make(map[string]astmodel.Type)

	for _, px := range parameters {
		// ensure that any $ref params are fully resolved
		_, parameter := extractor.fullyResolveParameter(px)

		if parameter.In == "path" &&
			parameter.Required {
			// if an enum:
			if len(parameter.Enum) > 0 {

				// found an enum that needs expansion, replace '{parameterName}' with
				// each value of the enum

				var newResults []string

				replace := fmt.Sprintf("{%s}", parameter.Name)
				values := enumValuesToStrings(parameter.Enum)

				for _, result := range results {
					for _, enumValue := range values {
						newResults = append(newResults, strings.ReplaceAll(result, replace, enumValue))
					}
				}

				results = newResults
			} else {
				// non-enum parameter
				var err error
				var paramType astmodel.Type
				if parameter.SimpleSchema.Type != "" {
					// handle "SimpleSchema"
					paramType, err = GetPrimitiveType(SchemaType(parameter.SimpleSchema.Type))
					if err != nil {
						panic(err)
					}
				} else {
					schema := extractor.schemaFromParameter(parameter)
					if schema == nil {
						panic(fmt.Sprintf("no schema generated for parameter %s in path %q", parameter.Name, operationPath))
					}

					paramType, err = scanner.RunHandlerForSchema(ctx, *schema)
					if err != nil {
						panic(err)
					}
				}

				urlParams[parameter.Name] = paramType
			}
		}
	}

	return results, urlParams
}

// if you update this you might also need to update "jsonast.enumValuesToLiterals"
func enumValuesToStrings(enumValues []interface{}) []string {
	result := make([]string, len(enumValues))
	for i, enumValue := range enumValues {
		if enumString, ok := enumValue.(string); ok {
			result[i] = enumString
		} else if enumStringer, ok := enumValue.(fmt.Stringer); ok {
			result[i] = enumStringer.String()
		} else if enumFloat, ok := enumValue.(float64); ok {
			result[i] = fmt.Sprintf("%g", enumFloat)
		} else {
			panic(fmt.Sprintf("unable to convert enum value (%s %T) to string", enumValue, enumValue))
		}
	}

	return result
}

func (extractor *SwaggerTypeExtractor) resourceNameFromOperationPath(operationPath string) (string, astmodel.TypeName, error) {
	group, resource, name, err := inferNameFromURLPath(operationPath)
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
func inferNameFromURLPath(operationPath string) (string, string, string, error) {
	group := ""
	nameParts := []string{}

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
	nameParts[len(nameParts)-1] = astmodel.Singularize(nameParts[len(nameParts)-1])

	name := strings.Join(nameParts, "")

	return group, resource, name, nil
}

// SwaggerGroupRegex matches a “group” (Swagger ‘namespace’)
// based on: https://github.com/Azure/autorest/blob/85de19623bdce3ccc5000bae5afbf22a49bc4665/core/lib/pipeline/metadata-generation.ts#L25
var SwaggerGroupRegex = regexp.MustCompile(`[Mm]icrosoft\.[^/\\]+`)
