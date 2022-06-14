/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"fmt"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AssertPackagesGenerateExpectedCode creates a golden file for each package represented in the set of type definitions,
// asserting that the generated content is expected.
// t is the current test
// definitions is the set of type definitions to be asserted
// options is an optional set of configuration options to control the assertion
func AssertPackagesGenerateExpectedCode(t *testing.T, definitions astmodel.TypeDefinitionSet, options ...AssertionOption) {
	// Group type definitions by package
	groups := make(map[astmodel.PackageReference][]astmodel.TypeDefinition, len(definitions))
	for _, def := range definitions {
		ref := def.Name().PackageReference
		groups[ref] = append(groups[ref], def)
	}

	// Write a file for each package
	for _, defs := range groups {
		ref := defs[0].Name().PackageReference
		group, version := ref.GroupVersion()
		fileName := fmt.Sprintf("%s-%s", group, version)
		AssertTypeDefinitionsGenerateExpectedCode(t, fileName, defs, options...)
	}
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
	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert(name, defs...)
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
	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert(fileName, def)
}
