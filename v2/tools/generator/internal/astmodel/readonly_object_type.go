/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// ReadonlyObjectType allows exposing an ObjectType for querying while prohibiting any modification.
// This is useful as a convenience to consumers for working with ObjectTypes, while ensuring they
// make any modifications using a TypeVisitor (or one of the Injectors based on TypeVisitor).
type ReadonlyObjectType interface {
	// Properties returns all the property definitions
	Properties() ReadOnlyPropertySet

	// Property returns the details of a specific property based on its unique case-sensitive name
	Property(name PropertyName) (*PropertyDefinition, bool)

	// EmbeddedProperties returns all the embedded properties
	// A sorted slice is returned to preserve immutability and provide determinism
	EmbeddedProperties() []*PropertyDefinition

	IsResource() bool
	Resources() TypeNameSet

	// Functions returns all the function implementations
	// A sorted slice is returned to preserve immutability and provide determinism
	Functions() []Function

	// HasFunctionWithName determines if this object has a function with the given name
	HasFunctionWithName(name string) bool

	TestCases() []TestCase
}
