/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// AssertPackagesGenerateExpectedCode creates a golden file for each package represented in the set of type definitions,
// asserting that the generated content is expected.
// t is the current test
// definitions is the set of type definitions to be asserted
// options is an optional set of configuration options to control the assertion
func AssertPackagesGenerateExpectedCode(
	t *testing.T,
	definitions astmodel.TypeDefinitionSet,
	options ...AssertionOption,
) {
	t.Helper()

	defs := definitions.AsSlice()
	packages := createSetOfPackages(defs)

	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert("", defs, packages)
}

// AssertTypeDefinitionsGenerateExpectedCode serialises the given FileDefinition as a golden file test, checking that the expected
// results are generated
// t is the current test
// name is a unique name for the current assertion
// defs is a set of type definitions to be asserted
// options is an optional set of configuration options to control the assertion
func AssertTypeDefinitionsGenerateExpectedCode(
	t *testing.T,
	name string,
	defs []astmodel.TypeDefinition,
	options ...AssertionOption,
) {
	t.Helper()

	packages := createSetOfPackages(defs)

	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert(name, defs, packages)
}

// AssertSingleTypeDefinitionGeneratesExpectedCode serialises the given TypeDefinition as a golden file test, checking
// that the expected results are generated
// t is the current test
// name is a unique name for the current assertion
// def is the type definition to be asserted
// options is an optional set of configuration options to control the assertion
func AssertSingleTypeDefinitionGeneratesExpectedCode(
	t *testing.T,
	fileName string,
	def astmodel.TypeDefinition,
	options ...AssertionOption,
) {
	t.Helper()

	defs := []astmodel.TypeDefinition{
		def,
	}

	packages := createSetOfPackages(defs)

	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert(fileName, defs, packages)
}

// AssertDefinitionHasExpectedShape fails the test if the given definition does not have the expected shape.
// t is the current test.
// filename is the name of the golden file to write.
// def is the definition to be asserted.
func AssertDefinitionHasExpectedShape(
	t *testing.T,
	filename string,
	def astmodel.TypeDefinition,
) {
	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(def)
	AssertDefinitionsHaveExpectedShapes(t, filename, defs)
}

// AssertDefinitionsHaveExpectedShapes fails the test if the given definition does not have the expected shape.
// t is the current test.
// filename is the name of the golden file to write.
// def is the definition to be asserted.
func AssertDefinitionsHaveExpectedShapes(
	t *testing.T,
	filename string,
	defs astmodel.TypeDefinitionSet,
) {
	t.Helper()
	g := goldie.New(t)

	err := g.WithTestNameForDir(true)
	if err != nil {
		t.Fatalf("Unable to configure goldie output folder %s", err)
	}

	buf := &bytes.Buffer{}
	report := reporting.NewTypeCatalogReport(defs)
	err = report.WriteTo(buf)
	if err != nil {
		t.Fatalf("unable to write report to buffer: %s", err)
	}

	g.Assert(t, filename, buf.Bytes())
}

// AssertPropertyExists fails the test if the given object does not have a property with the given name and type
// t is the current test.
// atype is the type that's expected to have the property.
// expectedName is the name of the property we expect to be present.
func AssertPropertyExists(
	t *testing.T,
	aType astmodel.Type,
	expectedName astmodel.PropertyName,
) *astmodel.PropertyDefinition {
	t.Helper()
	container, ok := astmodel.AsPropertyContainer(aType)
	if !ok {
		t.Fatalf("Expected %s to be a property container", astmodel.DebugDescription(aType))
	}

	property, ok := container.Property(expectedName)
	if !ok {
		t.Fatalf("Expected object to have property %q", expectedName)
	}

	return property
}

// AssertPropertyExistsWithType fails the test if the given object does not have a property with the given name and type
// t is the current test.
// atype is the type that's expected to have the property.
// expectedName is the name of the property we expect to be present.
// expectedType is the type of the property we expect to be present.
func AssertPropertyExistsWithType(
	t *testing.T,
	aType astmodel.Type,
	expectedName astmodel.PropertyName,
	expectedType astmodel.Type,
) *astmodel.PropertyDefinition {
	t.Helper()
	property := AssertPropertyExists(t, aType, expectedName)
	if !astmodel.TypeEquals(property.PropertyType(), expectedType) {
		t.Fatalf(
			"Expected property %q to have type %q, but was %q",
			expectedName,
			astmodel.DebugDescription(expectedType),
			astmodel.DebugDescription(property.PropertyType()))
	}

	return property
}

// AssertPropertyCount fails the test if the given object does not have the expected number of properties.
func AssertPropertyCount(
	t *testing.T,
	aType astmodel.Type,
	expected int,
) {
	t.Helper()

	container, ok := astmodel.AsPropertyContainer(aType)
	if !ok {
		t.Fatalf("Expected %s to be a property container", astmodel.DebugDescription(aType))
	}

	actual := container.Properties().Len()
	if actual != expected {
		t.Fatalf("Expected object to have %d properties, but had %d", expected, actual)
	}
}
