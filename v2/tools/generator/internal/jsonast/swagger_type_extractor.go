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
	ResourceDefinitions astmodel.TypeDefinitionSet
	OtherDefinitions    astmodel.TypeDefinitionSet
}

// ExtractTypes finds all operations in the Swagger spec that
// have a PUT verb and a path like "Microsoft.GroupName/…/resourceName/{resourceId}",
// and extracts the types for those operations, into the 'resourceTypes' result.
// Any additional types required by the resource types are placed into the 'otherTypes' result.
func (extractor *SwaggerTypeExtractor) ExtractTypes(ctx context.Context) (SwaggerTypes, error) {
	result := SwaggerTypes{
		ResourceDefinitions: make(astmodel.TypeDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	scanner := NewSchemaScanner(extractor.idFactory, extractor.config)

	for rawOperationPath, op := range extractor.swagger.Paths.Paths {
		put := op.Put
		if put == nil {
			continue
		}

		resourceSchema := extractor.findARMResourceSchema(*put)
		if resourceSchema == nil {
			// klog.Warningf("No ARM schema found for %s in %q", rawOperationPath, filePath)
			continue
		}

		for _, operationPath := range expandEnumsInPath(rawOperationPath, put.Parameters) {

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

			if existingResource, ok := result.ResourceDefinitions[resourceName]; ok {
				if !astmodel.TypeEquals(existingResource.Type(), resourceType) {
					return SwaggerTypes{}, errors.Errorf("resource already defined differently: %s\ndiff: %s",
						resourceName,
						astmodel.DiffTypes(existingResource.Type(), resourceType))
				}
			} else {
				result.ResourceDefinitions.Add(astmodel.MakeTypeDefinition(resourceName, resourceType))
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
func (extractor *SwaggerTypeExtractor) findARMResourceSchema(op spec.Operation) *Schema {
	if op.Responses != nil {
		for statusCode, response := range op.Responses.StatusCodeResponses {
			// only check OK and Created (per above linked comment)
			if statusCode == 200 || statusCode == 201 {
				result := extractor.getARMResourceSchemaFromResponse(response)
				if result != nil {
					return result
				}
			}
		}
	}

	// none found
	return nil
}

func (extractor *SwaggerTypeExtractor) getARMResourceSchemaFromResponse(response spec.Response) *Schema {
	if response.Schema != nil {
		// the schema can either be directly included

		schema := MakeOpenAPISchema(
			*response.Schema,
			extractor.swaggerPath,
			extractor.outputPackage,
			extractor.idFactory,
			extractor.cache)

		if isMarkedAsARMResource(schema) {
			return &schema
		}
	} else if response.Ref.GetURL() != nil {
		// or it can be under a $ref

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

		if isMarkedAsARMResource(schema) {
			return &schema
		}
	}

	return nil
}

// a Schema represents an ARM resource if it (or anything reachable via $ref or AllOf)
// is marked with x-ms-azure-resource, or if it (or anything reachable via $ref or AllOf)
// has each of the properties: id,name,type.
func isMarkedAsARMResource(schema Schema) bool {
	hasID := false
	hasName := false
	hasType := false

	var recurse func(schema Schema) bool
	recurse = func(schema Schema) bool {
		if schema.extensions("x-ms-azure-resource") == true {
			return true
		}

		props := schema.properties()
		if !hasID {
			if idProp, ok := props["id"]; ok {
				if idProp.hasType("string") {
					hasID = true
				}
			}
		}

		if !hasName {
			if nameProp, ok := props["name"]; ok {
				if nameProp.hasType("string") {
					hasName = true
				}
			}
		}

		if !hasType {
			if typeProp, ok := props["type"]; ok {
				if typeProp.hasType("string") {
					hasType = true
				}
			}
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
		return astmodel.EmptyTypeName, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
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
