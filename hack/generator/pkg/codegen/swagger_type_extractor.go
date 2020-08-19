/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"regexp"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

type typeExtractor struct {
	idFactory astmodel.IdentifierFactory
	config    *config.Configuration
	cache     jsonast.OpenAPISchemaCache
	// group for output types (e.g. Microsoft.Network.Frontdoor)
	outputGroup   string
	outputVersion string
}

// extractTypes finds all operations in the Swagger spec that
// have a PUT verb and a path like "Microsoft.GroupName/…/resourceName/{resourceId}",
// and extracts the types for those operations, into the 'resources' parameter.
// Any additional types required by the resource types are placed into the 'otherTypes' parameter.
func (extractor *typeExtractor) extractTypes(
	ctx context.Context,
	filePath string,
	swagger spec.Swagger,
	resources astmodel.Types,
	otherTypes astmodel.Types) error {

	packageName := extractor.idFactory.CreatePackageNameFromVersion(extractor.outputVersion)

	scanner := jsonast.NewSchemaScanner(extractor.idFactory, extractor.config)

	for operationPath, op := range swagger.Paths.Paths {
		put := op.Put
		if put == nil {
			continue
		}

		resourceName, err := extractor.resourceNameFromOperationPath(packageName, operationPath)
		if err != nil {
			klog.Error(err)
			continue
		}

		resourceType, err := extractor.resourceTypeFromOperation(ctx, scanner, swagger, filePath, put)
		if err != nil {
			if err == context.Canceled {
				return err
			}

			return errors.Wrapf(err, "unable to produce type for resource %v", resourceName)
		}

		if resourceType == nil {
			continue
		}

		if existingResource, ok := resources[resourceName]; ok {
			if !astmodel.TypeEquals(existingResource.Type(), resourceType) {
				klog.Errorf("RESOURCE already defined differently 😱: %v", resourceName)
			}
		} else {
			resources.Add(astmodel.MakeTypeDefinition(resourceName, resourceType))
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

func (extractor *typeExtractor) resourceNameFromOperationPath(packageName string, operationPath string) (astmodel.TypeName, error) {
	_, name, err := inferNameFromURLPath(operationPath)
	if err != nil {
		return astmodel.TypeName{}, errors.Wrapf(err, "unable to infer name from path %q", operationPath)
	}

	packageRef := astmodel.MakeLocalPackageReference(extractor.idFactory.CreateGroupName(extractor.outputGroup), packageName)
	return astmodel.MakeTypeName(packageRef, name), nil
}

func (extractor *typeExtractor) resourceTypeFromOperation(
	ctx context.Context,
	scanner *jsonast.SchemaScanner,
	schemaRoot spec.Swagger,
	filePath string,
	operation *spec.Operation) (astmodel.Type, error) {

	for _, param := range operation.Parameters {
		if param.In == "body" && param.Required { // assume this is the Resource
			schema := jsonast.MakeOpenAPISchema(
				*param.Schema,
				schemaRoot,
				filePath,
				extractor.outputGroup,
				extractor.outputVersion,
				extractor.cache)

			return scanner.RunHandlerForSchema(ctx, schema)
		}
	}

	return nil, nil
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
		if reading {
			if len(urlPart) > 0 && urlPart[0] != '{' {
				name += strings.ToUpper(urlPart[0:1]) + urlPart[1:]
				skippedLast = false
			} else {
				if skippedLast {
					// this means two {parameters} in a row
					return "", "", errors.Errorf("multiple parameters in path")
				}

				skippedLast = true
			}
		} else if swaggerGroupRegex.MatchString(urlPart) {
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

// based on: https://github.com/Azure/autorest/blob/85de19623bdce3ccc5000bae5afbf22a49bc4665/core/lib/pipeline/metadata-generation.ts#L25
var swaggerGroupRegex = regexp.MustCompile(`[Mm]icrosoft\.[^/\\]+`)
