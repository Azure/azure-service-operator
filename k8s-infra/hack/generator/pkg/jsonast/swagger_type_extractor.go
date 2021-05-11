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

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

type SwaggerTypeExtractor struct {
	idFactory astmodel.IdentifierFactory
	config    *config.Configuration
	cache     OpenAPISchemaCache
	// group for output types (e.g. Microsoft.Network.Frontdoor)
	outputGroup   string
	outputVersion string
}

// NewSwaggerTypeExtractor creates a new SwaggerTypeExtractor
func NewSwaggerTypeExtractor(
	config *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	outputGroup string,
	outputVersion string,
	cache OpenAPISchemaCache) SwaggerTypeExtractor {

	return SwaggerTypeExtractor{
		idFactory:     idFactory,
		outputVersion: outputVersion,
		outputGroup:   outputGroup,
		cache:         cache,
		config:        config,
	}
}

// ExtractTypes finds all operations in the Swagger spec that
// have a PUT verb and a path like "Microsoft.GroupName/…/resourceName/{resourceId}",
// and extracts the types for those operations, into the 'resources' parameter.
// Any additional types required by the resource types are placed into the 'otherTypes' parameter.
func (extractor *SwaggerTypeExtractor) ExtractTypes(
	ctx context.Context,
	filePath string,
	swagger spec.Swagger,
	resources astmodel.Types,
	otherTypes astmodel.Types) error {

	packageName := astmodel.CreateLocalPackageNameFromVersion(extractor.outputVersion)

	scanner := NewSchemaScanner(extractor.idFactory, extractor.config)

	for rawOperationPath, op := range swagger.Paths.Paths {
		put := op.Put
		if put == nil {
			continue
		}

		resourceSchema := extractor.findARMResourceSchema(filePath, swagger, *put)
		if resourceSchema == nil {
			//klog.Warningf("No ARM schema found for %s in %q", rawOperationPath, filePath)
			continue
		}

		for _, operationPath := range expandEnumsInPath(rawOperationPath, put.Parameters) {

			resourceName, err := extractor.resourceNameFromOperationPath(packageName, operationPath)
			if err != nil {
				klog.Errorf("Error extracting resource name (%s): %s", filePath, err.Error())
				continue
			}

			resourceType, err := scanner.RunHandlerForSchema(ctx, *resourceSchema)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}

				return errors.Wrapf(err, "unable to produce type for resource %v", resourceName)
			}

			if resourceType == nil {
				// this indicates a filtered-out type
				continue
			}

			if existingResource, ok := resources[resourceName]; ok {
				if !astmodel.TypeEquals(existingResource.Type(), resourceType) {
					return errors.Errorf("resource already defined differently: %v", resourceName)
				}
			} else {
				resources.Add(astmodel.MakeTypeDefinition(resourceName, resourceType))
			}
		}
	}

	for _, def := range scanner.Definitions() {
		// now add in the additional type definitions required by the resources
		if existingDef, ok := otherTypes[def.Name()]; ok {
			if !astmodel.TypeEquals(existingDef.Type(), def.Type()) {
				klog.Errorf("type already defined differently: %v", def.Name())
			}
		} else {
			otherTypes.Add(def)
		}
	}

	return nil
}

// Look at the responses of the PUT to determine if this represents an ARM resource,
// and if so, return the schema for it.
// see: https://github.com/Azure/autorest/issues/1936#issuecomment-286928591
func (extractor *SwaggerTypeExtractor) findARMResourceSchema(
	filePath string,
	swagger spec.Swagger,
	op spec.Operation) *Schema {

	if op.Responses != nil {
		for statusCode, response := range op.Responses.StatusCodeResponses {
			// only check OK and Created (per above linked comment)
			if statusCode == 200 || statusCode == 201 {
				result := extractor.getARMResourceSchemaFromResponse(filePath, swagger, response)
				if result != nil {
					return result
				}
			}
		}
	}

	// none found
	return nil
}

func (extractor *SwaggerTypeExtractor) getARMResourceSchemaFromResponse(filePath string, swagger spec.Swagger, response spec.Response) *Schema {
	if response.Schema != nil {
		// the schema can either be directly included

		schema := MakeOpenAPISchema(
			*response.Schema,
			swagger,
			filePath,
			extractor.outputGroup,
			extractor.outputVersion,
			extractor.cache)

		if isMarkedAsARMResource(schema) {
			return &schema
		}
	} else if response.Ref.GetURL() != nil {
		// or it can be under a $ref

		refFilePath, refSwagger, refSchema := loadRefSchema(swagger, response.Ref, filePath, extractor.cache)
		schema := MakeOpenAPISchema(
			refSchema,
			refSwagger,
			refFilePath,
			extractor.outputGroup,
			extractor.outputVersion,
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
		if schema.extensions()["x-ms-azure-resource"] == true {
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
			panic(fmt.Sprintf("unable to convert enum value (%v %T) to string", enumValue, enumValue))
		}
	}

	return result
}

func (extractor *SwaggerTypeExtractor) resourceNameFromOperationPath(packageName string, operationPath string) (astmodel.TypeName, error) {
	_, name, err := inferNameFromURLPath(operationPath)
	if err != nil {
		return astmodel.TypeName{}, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
	}

	packageRef := extractor.config.MakeLocalPackageReference(extractor.idFactory.CreateGroupName(extractor.outputGroup), packageName)
	return astmodel.MakeTypeName(packageRef, name), nil
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
