/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "go/ast"

type FakeFunction struct {
	name       string
	Imported   map[PackageReference]struct{}
	Referenced TypeNameSet
}

func NewFakeFunction(name string) *FakeFunction {
	return &FakeFunction{
		name:     name,
		Imported: nil,
	}
}

func (fake *FakeFunction) Name() string {
	return fake.name
}

func (fake *FakeFunction) RequiredPackageReferences() []PackageReference {
	var result []PackageReference
	for k := range fake.Imported {
		result = append(result, k)
	}

	return result
}

func (fake *FakeFunction) References() TypeNameSet {
	return fake.Referenced
}

func (fake *FakeFunction) AsFunc(_ *CodeGenerationContext, _ TypeName) *ast.FuncDecl {
	panic("implement me")
}

func (fake *FakeFunction) Equals(f Function) bool {
	if fake == nil && f == nil {
		return true
	}

	if fake == nil || f == nil {
		return false
	}

	fn, ok := f.(*FakeFunction)
	if !ok {
		return false
	}

	// Check to see if they have the same references
	if !fake.Referenced.Equals(fn.Referenced) {
		return false
	}

	// Check to see if they have the same imports
	if len(fake.Imported) != len(fn.Imported) {
		return false
	}

	for imp := range fake.Imported {
		if _, ok := fn.Imported[imp]; !ok {
			// Missing key, not equal
			return false
		}
	}

	return true
}

var _ Function = &FakeFunction{}
