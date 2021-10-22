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
	cache CachingFileLoader) SwaggerTypeExtractor {

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
	ResourceTypes, OtherTypes astmodel.Types
}

// ExtractTypes finds all operations in the Swagger spec that
// have a PUT verb and a path like "Microsoft.GroupName/…/resourceName/{resourceId}",
// and extracts the types for those operations, into the 'resourceTypes' result.
// Any additional types required by the resource types are placed into the 'otherTypes' result.
func (extractor *SwaggerTypeExtractor) ExtractTypes(ctx context.Context) (SwaggerTypes, error) {
	result := SwaggerTypes{
		ResourceTypes: make(astmodel.Types),
		OtherTypes:    make(astmodel.Types),
	}

	scanner := NewSchemaScanner(extractor.idFactory, extractor.config)

	for rawOperationPath, op := range extractor.swagger.Paths.Paths {
		// a Resource must have both PUT and GET
		if op.Put == nil || op.Get == nil {
			continue
		}

		resourceSchema := extractor.findARMResourceSchema(op, rawOperationPath)
		if resourceSchema == nil {
			continue
		}

		for _, operationPath := range expandEnumsInPath(rawOperationPath, op.Put.Parameters) {

			resourceName, err := extractor.resourceNameFromOperationPath(operationPath)
			if err != nil {
				klog.Errorf("Error extracting resource name (%s): %s", extractor.swaggerPath, err.Error())
				continue
			}

			resourceType, err := scanner.RunHandlerForSchema(ctx, *resourceSchema)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return SwaggerTypes{}, err
				}

				return SwaggerTypes{}, errors.Wrapf(err, "unable to produce type for resource %s", resourceName)
			}

			if resourceType == nil {
				// this indicates a filtered-out type
				continue
			}

			if existingResource, ok := result.ResourceTypes[resourceName]; ok {
				if !astmodel.TypeEquals(existingResource.Type(), resourceType) {
					return SwaggerTypes{}, errors.Errorf("resource already defined differently: %s\ndiff: %s",
						resourceName,
						astmodel.DiffTypes(existingResource.Type(), resourceType))
				}
			} else {
				result.ResourceTypes.Add(astmodel.MakeTypeDefinition(resourceName, resourceType))
			}
		}
	}

	for _, def := range scanner.Definitions() {
		// now add in the additional type definitions required by the resources
		if existingDef, ok := result.OtherTypes[def.Name()]; ok {
			if !astmodel.TypeEquals(existingDef.Type(), def.Type()) {
				return SwaggerTypes{}, errors.Errorf("type already defined differently: %s\nwas %s is %s\ndiff: %s",
					def.Name(),
					existingDef.Type(),
					def.Type(),
					astmodel.DiffTypes(existingDef.Type(), def.Type()))
			}
		} else {
			result.OtherTypes.Add(def)
		}
	}

	return result, nil
}

// Look at the responses of the PUT to determine if this represents an ARM resource,
// and if so, return the schema for it.
// see: https://github.com/Azure/autorest/issues/1936#issuecomment-286928591
func (extractor *SwaggerTypeExtractor) findARMResourceSchema(op spec.PathItem, rawOperationPath string) *Schema {
	// to decide if something is a resource, we must look at the responses from the GET
	isResource := false

	if op.Get.Responses != nil {
		for statusCode, response := range op.Get.Responses.StatusCodeResponses {
			// only check OK and Created (per above linked comment)
			// TODO: we really should check that the results of all of these are the same
			if statusCode == 200 || statusCode == 201 {
				if extractor.doesResponseRepresentARMResource(response, rawOperationPath) {
					isResource = true
					break
				}
			}
		}
	}

	if !isResource {
		return nil // not a resource
	}

	params := op.Put.Parameters
	if op.Parameters != nil {
		klog.Warningf("overriding parameters for %q in %s", rawOperationPath, extractor.swaggerPath)
		params = op.Parameters
	}

	// the actual Schema must come from the PUT parameters
	for _, param := range params {
		result := extractor.schemaFromParameter(param)
		if result != nil {
			return result
		}
	}

	klog.Warningf("Response indicated that type was ARM resource but no schema found for %s in %q", rawOperationPath, extractor.swaggerPath)
	return nil
}

func (extractor *SwaggerTypeExtractor) fullyResolveParameter(param spec.Parameter) spec.Parameter {
	if param.Ref.GetURL() == nil {
		return param
	}

	_, param, _ = loadRefParameter(param.Ref, extractor.swaggerPath, extractor.cache)
	return extractor.fullyResolveParameter(param)
}

func (extractor *SwaggerTypeExtractor) schemaFromParameter(param spec.Parameter) *Schema {
	param = extractor.fullyResolveParameter(param)

	if param.In != "body" {
		return nil
	}

	result := MakeOpenAPISchema(
		*param.Schema, // all params marked as 'body' have schemas
		extractor.swaggerPath,
		extractor.outputPackage,
		extractor.idFactory,
		extractor.cache)

	return &result
}

func (extractor *SwaggerTypeExtractor) doesResponseRepresentARMResource(response spec.Response, rawOperationPath string) bool {
	// the schema can either be directly included
	if response.Schema != nil {
		schema := MakeOpenAPISchema(
			*response.Schema,
			extractor.swaggerPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)

		return isMarkedAsARMResource(schema)
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

		return isMarkedAsARMResource(schema)
	}

	klog.Warningf("Unable to locate schema on response for %q in %s", rawOperationPath, extractor.swaggerPath)
	return false
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
				if idProp.hasType("string") &&
					idProp.readOnly() {
					hasID = true
				}
			}
		}

		if !idRequired {
			idRequired = containsString(schema.requiredProperties(), "id")
		}

		if !hasName {
			if nameProp, ok := props["name"]; ok {
				if nameProp.hasType("string") &&
					nameProp.readOnly() {
					hasName = true
				}
			}
		}

		if !nameRequired {
			nameRequired = containsString(schema.requiredProperties(), "name")
		}

		if !hasType {
			if typeProp, ok := props["type"]; ok {
				if typeProp.hasType("string") &&
					typeProp.readOnly() {
					hasType = true
				}
			}
		}

		if !typeRequired {
			typeRequired = containsString(schema.requiredProperties(), "type")
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

func expandEnumsInPath(operationPath string, parameters []spec.Parameter) []string {
	results := []string{operationPath}

	for _, parameter := range parameters {
		if parameter.In == "path" &&
			parameter.Required &&
			len(parameter.Enum) > 0 {

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
		}
	}

	return results
}

func containsString(xs []string, needle string) bool {
	for _, x := range xs {
		if x == needle {
			return true
		}
	}

	return false
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

func (extractor *SwaggerTypeExtractor) resourceNameFromOperationPath(operationPath string) (astmodel.TypeName, error) {
	_, name, err := inferNameFromURLPath(operationPath)
	if err != nil {
		return astmodel.TypeName{}, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
	}

	return astmodel.MakeTypeName(extractor.outputPackage, name), nil
}

// inferNameFromURLPath attempts to extract a name from a Swagger operation path
// for example “…/Microsoft.GroupName/resourceType/{resourceId}” would result
// in the name “ResourceType”. Child resources are treated by converting (e.g.)
// “…/Microsoft.GroupName/resourceType/{parameterId}/differentType/{otherId}/something/{moreId}”
// to “ResourceTypeDifferentTypeSomething”.
func inferNameFromURLPath(operationPath string) (string, string, error) {
	group := ""
	name := ""

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
					return "", "", errors.Errorf("multiple parameters in path")
				}

				skippedLast = true
			} else {
				// normal part of path, uppercase first character
				name += strings.ToUpper(urlPart[0:1]) + urlPart[1:]
				skippedLast = false
			}
		} else if SwaggerGroupRegex.MatchString(urlPart) {
			group = urlPart
			reading = true
		}
	}

	if !reading {
		return "", "", errors.Errorf("no group name (‘Microsoft…’ = %q) found", group)
	}

	if name == "" {
		return "", "", errors.Errorf("couldn’t infer name")
	}

	return group, name, nil
}

// SwaggerGroupRegex matches a “group” (Swagger ‘namespace’)
// based on: https://github.com/Azure/autorest/blob/85de19623bdce3ccc5000bae5afbf22a49bc4665/core/lib/pipeline/metadata-generation.ts#L25
var SwaggerGroupRegex = regexp.MustCompile(`[Mm]icrosoft\.[^/\\]+`)
