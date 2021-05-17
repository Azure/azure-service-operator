/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InterfaceImplementation specifies how a type will satisfy an interface implementation
type InterfaceImplementation struct {
	name       TypeName
	annotation string
	functions  map[string]Function
}

// NewInterfaceImplementation creates a new interface implementation with the given name and set of functions
func NewInterfaceImplementation(name TypeName, functions ...Function) *InterfaceImplementation {
	result := &InterfaceImplementation{name: name, functions: make(map[string]Function)}
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
	var results TypeNameSet
	for _, f := range iface.functions {
		for ref := range f.References() {
			results.Add(ref)
		}
	}

	return results
}

// Equals determines if this interface is equal to another interface
func (iface *InterfaceImplementation) Equals(other *InterfaceImplementation) bool {
	if len(iface.functions) != len(other.functions) {
		return false
	}

	for name, f := range iface.functions {
		otherF, ok := other.functions[name]
		if !ok {
			return false
		}

		if !f.Equals(otherF) {
			return false
		}
	}

	return true
}
