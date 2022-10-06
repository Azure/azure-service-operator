/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"fmt"
	"math/big"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// OpenAPISchema implements the Schema abstraction for go-openapi
type OpenAPISchema struct {
	name          string // name of the schema (may be empty if nested)
	inner         spec.Schema
	fileName      string // fully qualified file path of the file from which this schema was loaded
	outputPackage astmodel.LocalPackageReference
	idFactory     astmodel.IdentifierFactory
	loader        OpenAPIFileLoader
}

// MakeOpenAPISchema wraps a spec.Swagger to conform to the Schema abstraction
func MakeOpenAPISchema(
	name string,
	schema spec.Schema,
	fileName string,
	outputPackage astmodel.LocalPackageReference,
	idFactory astmodel.IdentifierFactory,
	cache OpenAPIFileLoader) Schema {
	return &OpenAPISchema{
		name:          name,
		inner:         schema,
		fileName:      fileName,
		outputPackage: outputPackage,
		idFactory:     idFactory,
		loader:        cache}
}

func (schema *OpenAPISchema) withNewSchema(newSchema spec.Schema) Schema {
	result := *schema
	result.inner = newSchema
	return &result
}

var _ Schema = &OpenAPISchema{}

func (schema *OpenAPISchema) transformOpenAPISlice(slice []spec.Schema) []Schema {
	result := make([]Schema, len(slice))
	for i := range slice {
		result[i] = schema.withNewSchema(slice[i])
	}

	return result
}

func (schema *OpenAPISchema) Id() string {
	return schema.name
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

//!! TODO: This is expensive to evaluate, so cache the result
func (schema *OpenAPISchema) discriminatorValues() set.Set[string] {
	// Must have a discriminator
	discriminator := schema.inner.Discriminator
	if discriminator == "" {
		// No options, no error
		return nil
	}

	// Find the discriminator property
	discriminatorProperty, ok := schema.properties()[discriminator]
	if !ok {
		// For now, if there is no discriminator property, we just ignore it
		return nil
	}

	// Expect the discriminator property to be an enum
	enumValues := discriminatorProperty.enumValues()
	if len(enumValues) == 0 {
		return nil
	}

	result := set.Make[string]()
	for _, v := range enumValues {
		result.Add(strings.Trim(v, "\""))
	}

	return result
}

func (schema *OpenAPISchema) discriminator() string {
	return schema.inner.Discriminator
}

func (schema *OpenAPISchema) requiredProperties() []string {
	return schema.inner.Required
}

func (schema *OpenAPISchema) properties() map[string]Schema {
	result := make(map[string]Schema, len(schema.inner.Properties))
	for propName, propSchema := range schema.inner.Properties {
		result[propName] = schema.withNewSchema(propSchema)
	}

	return result
}

func (schema *OpenAPISchema) maxLength() *int64 {
	return schema.inner.MaxLength
}

func (schema *OpenAPISchema) minLength() *int64 {
	return schema.inner.MinLength
}

func (schema *OpenAPISchema) format() string {
	return schema.inner.Format
}

func (schema *OpenAPISchema) pattern() *regexp.Regexp {
	p := schema.inner.Pattern
	if p == "" {
		return nil
	}

	result, err := regexp.Compile(p)
	if err != nil {
		klog.Warningf("Unable to compile regexp, ignoring: %s", p) // use %s instead of %q or everything gets re-escaped
		return nil
	}

	return result
}

func (schema *OpenAPISchema) maxItems() *int64 {
	return schema.inner.MaxItems
}

func (schema *OpenAPISchema) minItems() *int64 {
	return schema.inner.MinItems
}

func (schema *OpenAPISchema) uniqueItems() bool {
	return schema.inner.UniqueItems
}

func (schema *OpenAPISchema) minValue() *big.Rat {
	r := schema.inner.Minimum
	if r == nil {
		return nil
	}

	rat := &big.Rat{}
	rat.SetFloat64(*r)
	return rat
}

func (schema *OpenAPISchema) maxValue() *big.Rat {
	r := schema.inner.Maximum
	if r == nil {
		return nil
	}

	rat := &big.Rat{}
	rat.SetFloat64(*r)
	return rat
}

func (schema *OpenAPISchema) minValueExclusive() bool {
	return schema.inner.ExclusiveMinimum
}

func (schema *OpenAPISchema) maxValueExclusive() bool {
	return schema.inner.ExclusiveMaximum
}

func (schema *OpenAPISchema) multipleOf() *big.Rat {
	r := schema.inner.MultipleOf
	if r == nil {
		return nil
	}

	rat := &big.Rat{}
	rat.SetFloat64(*r)
	return rat
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
			panic(fmt.Sprintf("unable to convert enum value (%s %T) to literal", enumValue, enumValue))
		}
	}

	return result
}

func (schema *OpenAPISchema) enumValues() []string {
	return enumValuesToLiterals(schema.inner.Enum)
}

func (schema *OpenAPISchema) extensionAsString(key string) (string, bool) {
	return schema.inner.Extensions.GetString(key)
}

func (schema *OpenAPISchema) extensionAsBool(key string) bool {
	value, ok := schema.inner.Extensions.GetBool(key)
	return ok && value
}

func (schema *OpenAPISchema) hasExtension(key string) bool {
	_, found := schema.inner.Extensions[strings.ToLower(key)]
	return found
}

func (schema *OpenAPISchema) isRef() bool {
	return schema.inner.Ref.GetURL() != nil
}

func (schema *OpenAPISchema) refTypeName() (astmodel.TypeName, error) {
	absRefPath, err := findFileForRef(schema.fileName, schema.inner.Ref)
	if err != nil {
		return astmodel.EmptyTypeName, err
	}

	// this is the basic type name for the reference
	name := schema.refObjectName()

	// now locate the package name for the reference
	packageAndSwagger, err := schema.loader.loadFile(absRefPath)
	if err != nil {
		return astmodel.EmptyTypeName, err
	}

	// default to using same package as the referring type
	pkg := schema.outputPackage
	// however, if referree type has known package, use that instead
	// this allows us to override e.g. the Microsoft.Common namespace in config
	if packageAndSwagger.Package != nil {
		pkg = *packageAndSwagger.Package
	} else {
		// make sure that pulling the type into the other package wouldn’t conflict with
		// any definitions in that package; we iterate over all files to ensure that
		// we don’t conflict with “sibling” files as well as the pulling-in file
		otherFiles := schema.loader.knownFiles()
		for _, otherFile := range otherFiles {
			if otherFile == filepath.ToSlash(absRefPath) {
				continue // skip containing file
			}

			otherSchema, err := schema.loader.loadFile(otherFile)
			if err != nil {
				panic(err) // assert, not error: file should already be loaded if it is known
			}

			// check only applies if package is the same as the pulling-in package
			// or is nil (so could be set to the pulling-in package)
			if otherSchema.Package == nil || otherSchema.Package.Equals(schema.outputPackage) {
				if _, ok := otherSchema.Swagger.Definitions[name]; ok {
					return astmodel.EmptyTypeName, errors.Errorf(
						"importing type %s from file %s into package %s could generate collision with type in %s",
						name,
						absRefPath,
						pkg,
						otherFile,
					)
				}
			}
		}
	}

	return astmodel.MakeTypeName(pkg, schema.idFactory.CreateIdentifier(name, astmodel.Exported)), nil
}

func (schema *OpenAPISchema) readOnly() bool {
	return schema.inner.ReadOnly
}

func (schema *OpenAPISchema) refSchema() Schema {
	ref := schema.inner.Ref
	fileName, result, packageAndSwagger := loadRefSchema(ref, schema.fileName, schema.loader)

	// if the pkg comes back nil, that means we should keep using the current package
	// this happens for some ‘common’ types defined in files that don’t have groups or versions

	outputPackage := schema.outputPackage
	if packageAndSwagger.Package != nil {
		outputPackage = *packageAndSwagger.Package
	}

	return MakeOpenAPISchema(
		nameFromRef(ref),
		result,
		fileName,
		outputPackage,
		schema.idFactory,
		schema.loader)
}

// findFileForRef identifies the schema path for a ref, relative to the give schema path
func findFileForRef(relativeToSchemaPath string, ref spec.Ref) (string, error) {
	if ref.HasFragmentOnly {
		// same file
		return relativeToSchemaPath, nil
	}

	// an external path
	return resolveAbsolutePath(relativeToSchemaPath, ref.GetURL())
}

func loadRef(
	ref spec.Ref,
	relativeToSchemaPath string,
	loader OpenAPIFileLoader,
) (string, interface{}, PackageAndSwagger) {
	absPath, err := findFileForRef(relativeToSchemaPath, ref)
	if err != nil {
		panic(err)
	}

	packageAndSwagger, err := loader.loadFile(absPath)
	if err != nil {
		panic(err)
	}

	result, _, err := ref.GetPointer().Get(packageAndSwagger.Swagger)
	if err != nil {
		panic(fmt.Sprintf("cannot resolve ref %s in file %s (from %s): %s", ref.String(), absPath, relativeToSchemaPath, err))
	}

	return absPath, result, packageAndSwagger
}

func loadRefSchema(ref spec.Ref, relativeToSchemaPath string, loader OpenAPIFileLoader) (string, spec.Schema, PackageAndSwagger) {
	absPath, result, pkg := loadRef(ref, relativeToSchemaPath, loader)
	return absPath, result.(spec.Schema), pkg
}

func loadRefParameter(ref spec.Ref, relativeToSchemaPath string, loader OpenAPIFileLoader) (string, spec.Parameter) {
	absPath, result, _ := loadRef(ref, relativeToSchemaPath, loader)
	return absPath, result.(spec.Parameter)
}

func (schema *OpenAPISchema) refObjectName() string {
	return objectNameFromPointer(schema.inner.Ref.GetPointer())
}

func objectNameFromPointer(ptr *jsonpointer.Pointer) string {
	// turns a fragment like "#/definitions/Name" into "Name"
	tokens := ptr.DecodedTokens()
	if len(tokens) != 2 || tokens[0] != "definitions" {
		// this condition is never violated by the swagger files
		panic(fmt.Sprintf("not understood: %s", tokens))
	}

	return tokens[1]
}

// resolveAbsolutePath makes an absolute path by combining 'baseFileName' and 'url'
func resolveAbsolutePath(baseFileName string, url *url.URL) (string, error) {
	if url.IsAbs() {
		return "", errors.Errorf("absolute path %q not supported (only relative URLs)", url)
	}

	dir := filepath.Dir(baseFileName)

	result := filepath.Clean(filepath.Join(dir, url.Path))
	return result, nil
}
