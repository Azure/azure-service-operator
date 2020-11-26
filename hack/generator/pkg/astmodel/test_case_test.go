/*
* Copyright (c) Microsoft Corporation.
* Licensed under the MIT license.
 */

package astmodel

import (
	ast "github.com/dave/dst"
)

type FakeTestCase struct {
	name string
}

var _ TestCase = &FakeTestCase{}

func NewFakeTestCase(name string) *FakeTestCase {
	return &FakeTestCase{
		name: name,
	}
}

func (f FakeTestCase) Name() string {
	return f.name
}

func (f FakeTestCase) References() TypeNameSet {
	panic("implement me")
}

func (f FakeTestCase) RequiredImports() *PackageImportSet {
	panic("implement me")
}

func (f FakeTestCase) AsFuncs(subject TypeName, codeGenerationContext *CodeGenerationContext) []ast.Decl {
	panic("implement me")
}

func (f FakeTestCase) Equals(tc TestCase) bool {
	other, ok := tc.(*FakeTestCase)
	if !ok {
		return false
	}

	return f.name == other.name
}
