/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "github.com/dave/dst"

// TestCase represents a test we generate to ensure the generated code works as expected
type TestCase interface {
	// Name returns the unique name of this test case
	Name() string

	// References returns the set of types to which this test case refers.
	References() TypeNameSet

	// RequiredImports returns a set of the package imports required by this test case
	RequiredImports() *PackageImportSet

	// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
	// subject is the name of the type under test
	// codeGenerationContext contains reference material to use when generating
	AsFuncs(subject TypeName, codeGenerationContext *CodeGenerationContext) []dst.Decl

	// Equals determines if this TestCase is equal to another one
	Equals(f TestCase) bool
}

// TestCaseContainer represents types that can contain test cases
// These types allow us to generate tests to verify the generated code does the right thing
type TestCaseContainer interface {
	TestCases() []TestCase
}

// AsTestCaseContainer converts a type into a function container
// Only use this readonly access as we must use a TypeVisitor for modifications to preserve type wrapping
func AsTestCaseContainer(theType Type) (TestCaseContainer, bool) {
	switch t := theType.(type) {
	case TestCaseContainer:
		return t, true
	case MetaType:
		return AsTestCaseContainer(t.Unwrap())
	default:
		return nil, false
	}
}
