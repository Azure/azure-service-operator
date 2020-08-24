/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"

	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

// OpenAPISchema implements the Schema abstraction for go-openapi
type OpenAPISchema struct {
	inner     spec.Schema
	root      spec.Swagger
	fileName  string
	groupName string
	version   string
	cache     OpenAPISchemaCache
}

// MakeOpenAPISchema wraps a spec.Swagger to conform to the Schema abstraction
func MakeOpenAPISchema(
	schema spec.Schema,
	root spec.Swagger,
	fileName string,
	groupName string,
	version string,
	cache OpenAPISchemaCache) Schema {
	return &OpenAPISchema{schema, root, fileName, groupName, version, cache}
}

func (schema *OpenAPISchema) withNewSchema(newSchema spec.Schema) Schema {
	return &OpenAPISchema{
		newSchema,
		schema.root,
		schema.fileName,
		schema.groupName,
		schema.version,
		schema.cache,
	}
}

var _ Schema = &OpenAPISchema{}

func (schema *OpenAPISchema) transformOpenAPISlice(slice []spec.Schema) []Schema {
	result := make([]Schema, len(slice))
	for i := range slice {
		result[i] = schema.withNewSchema(slice[i])
	}

	return result
}

func (schema *OpenAPISchema) title() *string {
	if len(schema.inner.Title) == 0 {
		return nil // translate to optional
	}

	return &schema.inner.Title
}

func (schema *OpenAPISchema) url() *url.URL {
	url, err := url.Parse(schema.inner.ID)
	if err != nil {
		return nil
	}

	return url
}

func (schema *OpenAPISchema) hasType(schemaType SchemaType) bool {
	return schema.inner.Type.Contains(string(schemaType))
}

func (schema *OpenAPISchema) hasAllOf() bool {
	return len(schema.inner.AllOf) > 0
}

func (schema *OpenAPISchema) allOf() []Schema {
	return schema.transformOpenAPISlice(schema.inner.AllOf)
}

func (schema *OpenAPISchema) hasAnyOf() bool {
	return len(schema.inner.AnyOf) > 0
}

func (schema *OpenAPISchema) anyOf() []Schema {
	return schema.transformOpenAPISlice(schema.inner.AnyOf)
}

func (schema *OpenAPISchema) hasOneOf() bool {
	return len(schema.inner.OneOf) > 0
}

func (schema *OpenAPISchema) oneOf() []Schema {
	return schema.transformOpenAPISlice(schema.inner.OneOf)
}

func (schema *OpenAPISchema) requiredProperties() []string {
	return schema.inner.Required
}

func (schema *OpenAPISchema) properties() map[string]Schema {
	result := make(map[string]Schema)
	for propName, propSchema := range schema.inner.Properties {
		result[propName] = schema.withNewSchema(propSchema)
	}

	return result
}

func (schema *OpenAPISchema) description() *string {
	if len(schema.inner.Description) == 0 {
		return nil
	}

	return &schema.inner.Description
}

func (schema *OpenAPISchema) items() []Schema {
	if schema.inner.Items.Schema != nil {
		return []Schema{schema.withNewSchema(*schema.inner.Items.Schema)}
	}

	return schema.transformOpenAPISlice(schema.inner.Items.Schemas)
}

func (schema *OpenAPISchema) additionalPropertiesAllowed() bool {
	return schema.inner.AdditionalProperties == nil || schema.inner.AdditionalProperties.Allows
}

func (schema *OpenAPISchema) additionalPropertiesSchema() Schema {
	if schema.inner.AdditionalProperties == nil {
		return nil
	}

	result := schema.inner.AdditionalProperties.Schema
	if result == nil {
		return nil
	}

	return schema.withNewSchema(*result)
}

// enumValuesToLiterals converts interface{}-typed values to their
// literal go-lang representations
// if you update this you might also need to update "codegen.enumValuesToStrings"
func enumValuesToLiterals(enumValues []interface{}) []string {
	result := make([]string, len(enumValues))
	for i, enumValue := range enumValues {
		if enumString, ok := enumValue.(string); ok {
			result[i] = fmt.Sprintf("%q", enumString)
		} else if enumStringer, ok := enumValue.(fmt.Stringer); ok {
			result[i] = fmt.Sprintf("%q", enumStringer.String())
		} else if enumFloat, ok := enumValue.(float64); ok {
			result[i] = fmt.Sprintf("%g", enumFloat)
		} else {
			panic(fmt.Sprintf("unable to convert enum value (%v %T) to literal", enumValue, enumValue))
		}
	}

	return result
}

func (schema *OpenAPISchema) enumValues() []string {
	return enumValuesToLiterals(schema.inner.Enum)
}

func (schema *OpenAPISchema) extensions() map[string]interface{} {
	return schema.inner.Extensions
}

func (schema *OpenAPISchema) isRef() bool {
	return schema.inner.Ref.GetURL() != nil
}

func (schema *OpenAPISchema) refSchema() Schema {
	fileName, root, result := loadRefSchema(schema.root, schema.inner.Ref, schema.fileName, schema.cache)
	return &OpenAPISchema{
		result,
		root,
		fileName,
		// Note that we preserve the groupName and version that were input at the start,
		// even if we are reading a file from a different group or version. this is intentional;
		// essentially all imported types are copied into the target group/version, which avoids
		// issues with types from the 'common-types' files which have no group and a version of 'v1'.
		schema.groupName,
		schema.version,
		schema.cache,
	}
}

func loadRefSchema(
	schemaRoot spec.Swagger,
	ref spec.Ref,
	schemaPath string,
	cache OpenAPISchemaCache) (string, spec.Swagger, spec.Schema) {

	var fileName string
	var root spec.Swagger
	var err error

	if !ref.HasFragmentOnly {
		fileName, root, err = cache.fetchFileRelative(schemaPath, ref.GetURL())
		if err != nil {
			panic(err)
		}
	} else {
		fileName, root = schemaPath, schemaRoot
	}

	reffed := objectNameFromPointer(ref.GetPointer())
	if result, ok := root.Definitions[reffed]; !ok {
		panic(fmt.Sprintf("couldn't find: %s in %s", reffed, fileName))
	} else {
		return fileName, root, result
	}
}

func (schema *OpenAPISchema) refVersion() (string, error) {
	return schema.version, nil
}

func (schema *OpenAPISchema) refGroupName() (string, error) {
	return schema.groupName, nil
}

func (schema *OpenAPISchema) refObjectName() (string, error) {
	return objectNameFromPointer(schema.inner.Ref.GetPointer()), nil
}

func objectNameFromPointer(ptr *jsonpointer.Pointer) string {
	// turns a fragment like "#/definitions/Name" into "Name"
	tokens := ptr.DecodedTokens()
	if len(tokens) != 2 || tokens[0] != "definitions" {
		// this condition is never violated by the swagger files
		panic(fmt.Sprintf("not understood: %v", tokens))
	}

	return tokens[1]
}

// OpenAPISchemaCache is a cache of schema that have been loaded,
// identified by file path
type OpenAPISchemaCache struct {
	files map[string]spec.Swagger
}

// NewOpenAPISchemaCache creates an OpenAPISchemaCache with the initial
// file path → spec mapping
func NewOpenAPISchemaCache(specs map[string]spec.Swagger) OpenAPISchemaCache {
	files := make(map[string]spec.Swagger)
	for specPath, spec := range specs {
		files[specPath] = spec
	}

	return OpenAPISchemaCache{files}
}

// fetchFileRelative fetches the schema for the relative path created by combining 'baseFileName' and 'url'
func (fileCache OpenAPISchemaCache) fetchFileRelative(baseFileName string, url *url.URL) (string, spec.Swagger, error) {
	path := ""
	swagger := spec.Swagger{}

	if url.IsAbs() {
		return path, swagger, errors.Errorf("only relative URLs can be handled")
	}

	fileURL, err := url.Parse("file://" + filepath.ToSlash(baseFileName))

	if err != nil {
		return path, swagger, errors.Wrapf(err, "cannot convert filename to file URI")
	}

	path = fileURL.ResolveReference(url).Path
	swagger, err = fileCache.fetchFileAbsolute(path)

	return path, swagger, err
}

// fetchFileAbsolute fetches the schema for the absolute path specified
func (fileCache OpenAPISchemaCache) fetchFileAbsolute(filePath string) (spec.Swagger, error) {
	if swagger, ok := fileCache.files[filePath]; ok {
		return swagger, nil
	}

	var swagger spec.Swagger

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return swagger, errors.Wrapf(err, "unable to read swagger file %q", filePath)
	}

	err = swagger.UnmarshalJSON(fileContent)
	if err != nil {
		return swagger, errors.Wrapf(err, "unable to parse swagger file %q", filePath)
	}

	fileCache.files[filePath] = swagger

	return swagger, err
}
