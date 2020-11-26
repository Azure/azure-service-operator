/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import ast "github.com/dave/dst"

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
	AsFuncs(subject TypeName, codeGenerationContext *CodeGenerationContext) []ast.Decl

	// Equals determines if this TestCase is equal to another one
	Equals(f TestCase) bool
}

type TestCaseDefiner interface {
	TestCases() []TestCase
}
