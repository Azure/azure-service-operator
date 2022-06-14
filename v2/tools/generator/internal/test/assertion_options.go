/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// AssertionOption represents options that can be used to customize the behavour of the assert
type AssertionOption interface {
	// configure applies this option to the passed typeAsserter
	configure(ta *typeAsserter)
}

/*
 * diffOption
 */

// DiffWith specifies the type definitions being tested should be diff'd with the provided type definitions to highlight
// differences.
// TypeDefinitionSet are matched by fully qualified name.
func DiffWith(defs ...astmodel.TypeDefinition) AssertionOption {
	return &diffOption{
		references: defs,
	}
}

// DiffWithTypes specifies the type definitions being tested should be diff'd with the provided type definitions to
// highlight differences.
// TypeDefinitionSet are matched by fully qualified name.
func DiffWithTypes(definitions astmodel.TypeDefinitionSet) AssertionOption {
	defs := make([]astmodel.TypeDefinition, 0, len(definitions))
	for _, d := range definitions {
		defs = append(defs, d)
	}

	return &diffOption{
		references: defs,
	}
}

// diffOption captures a GoSourceFile which acts as a base for comparison
type diffOption struct {
	references []astmodel.TypeDefinition
}

var _ AssertionOption = &diffOption{}

// configure sets up the typeAsserter with the types to use as a reference
func (d *diffOption) configure(ta *typeAsserter) {
	ta.addReferences(d.references...)
}

/*
 * includeTestFiles
 */

// IncludeTestFiles creates an AssertionOption to also assert generated tests
func IncludeTestFiles() AssertionOption {
	return &includeTestFiles{}
}

type includeTestFiles struct{}

var _ AssertionOption = &includeTestFiles{}

// configure sets up the typeAsserter to also write test files
func (i *includeTestFiles) configure(ta *typeAsserter) {
	ta.writeTests = true
}

/*
 * excludeCodeFiles
 */

// ExcludeCodeFiles creates an AssertionOption to suppress regular code files
func ExcludeCodeFiles() AssertionOption {
	return &excludeCodeFiles{}
}

type excludeCodeFiles struct{}

var _ AssertionOption = &excludeCodeFiles{}

// configure sets up the typeAsserter to exclude code files
func (i *excludeCodeFiles) configure(ta *typeAsserter) {
	ta.writeCode = false
}
