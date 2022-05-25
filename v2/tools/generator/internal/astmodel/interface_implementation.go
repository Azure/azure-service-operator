/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"

	"golang.org/x/exp/maps"
)

// InterfaceImplementation specifies how a type will satisfy an interface implementation
type InterfaceImplementation struct {
	name       TypeName
	annotation string
	functions  map[string]Function
}

var _ FunctionContainer = &InterfaceImplementation{}

// NewInterfaceImplementation creates a new interface implementation with the given name and set of functions
func NewInterfaceImplementation(name TypeName, functions ...Function) *InterfaceImplementation {
	result := &InterfaceImplementation{name: name, functions: make(map[string]Function, len(functions))}
	for _, f := range functions {
		result.functions[f.Name()] = f
	}

	return result
}

func (iface *InterfaceImplementation) WithAnnotation(annotation string) *InterfaceImplementation {
	result := *iface
	result.annotation = annotation
	return &result
}

// Name returns the name of the interface
func (iface *InterfaceImplementation) Name() TypeName {
	return iface.name
}

// RequiredPackageReferences returns a list of packages required by this
func (iface *InterfaceImplementation) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()
	result.AddReference(iface.Name().PackageReference)

	for _, f := range iface.functions {
		result.Merge(f.RequiredPackageReferences())
	}

	return result
}

// References indicates whether this type includes any direct references to the given type
func (iface *InterfaceImplementation) References() TypeNameSet {
	results := NewTypeNameSet()
	for _, f := range iface.functions {
		for ref := range f.References() {
			results.Add(ref)
		}
	}

	return results
}

// FunctionCount returns the number of included functions
func (iface *InterfaceImplementation) FunctionCount() int {
	return len(iface.functions)
}

// Functions returns all the function implementations
// A sorted slice is returned to preserve immutability and provide determinism
func (iface *InterfaceImplementation) Functions() []Function {
	functions := maps.Values(iface.functions)

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	return functions
}

// HasFunctionWithName determines if this interface has a function with the given name
func (iface *InterfaceImplementation) HasFunctionWithName(functionName string) bool {
	_, found := iface.functions[functionName]
	return found
}

// Equals determines if this interface is equal to another interface
func (iface *InterfaceImplementation) Equals(other *InterfaceImplementation, overrides EqualityOverrides) bool {
	if len(iface.functions) != len(other.functions) {
		return false
	}

	for name, f := range iface.functions {
		otherF, ok := other.functions[name]
		if !ok {
			return false
		}

		if !f.Equals(otherF, overrides) {
			return false
		}
	}

	return true
}
