/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"fmt"
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AssertPackagesGenerateExpectedCode creates a golden file for each package represented in the set of type definitions,
// asserting that the generated content is expected.
// t is the current test
// types is the set of type definitions to be asserted
// options is an optional set of configuration options to control the assertion
func AssertPackagesGenerateExpectedCode(t *testing.T, types astmodel.Types, options ...AssertionOption) {
	// Group type definitions by package
	groups := make(map[astmodel.PackageReference][]astmodel.TypeDefinition)
	for _, def := range types {
		ref := def.Name().PackageReference
		groups[ref] = append(groups[ref], def)
	}

	// Write a file for each package
	for _, defs := range groups {
		ref := defs[0].Name().PackageReference
		local, ok := ref.AsLocalPackage()
		if !ok {
			panic("Must only have types from local packages - fix your test")
		}

		fileName := fmt.Sprintf("%s-%s", local.Group(), local.Version())
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
	options ...AssertionOption) {

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
	options ...AssertionOption) {

	asserter := newTypeAsserter(t)
	asserter.configure(options)
	asserter.assert(fileName, def)
}
