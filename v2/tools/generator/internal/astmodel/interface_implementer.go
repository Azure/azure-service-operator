/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"

	"github.com/dave/dst"
	"golang.org/x/exp/maps"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// InterfaceImplementer represents a container that may contain multiple interface implementations
// Both resources and objects are examples of interface implementers
type InterfaceImplementer struct {
	interfaces map[TypeName]*InterfaceImplementation
}

// MakeInterfaceImplementer returns an interface implementer
func MakeInterfaceImplementer() InterfaceImplementer {
	return InterfaceImplementer{}
}

// FindInterface is used to find a particular interface implementation when given the type name of the interface being
// implemented.
func (i InterfaceImplementer) FindInterface(name TypeName) (*InterfaceImplementation, bool) {
	result, ok := i.interfaces[name]
	return result, ok
}

// WithoutInterface returns a new interface implementer that doesn't contain the specified interface
func (i InterfaceImplementer) WithoutInterface(name TypeName) InterfaceImplementer {
	result := i.copy()
	delete(result.interfaces, name)
	return result
}

// WithInterface creates a new InterfaceImplementer with the specified implementation included.
// Any existing implementation of the same interface will be replaced.
func (i InterfaceImplementer) WithInterface(iface *InterfaceImplementation) InterfaceImplementer {
	result := i.copy()
	result.interfaces[iface.Name()] = iface

	return result
}

func (i InterfaceImplementer) References() TypeNameSet {
	results := NewTypeNameSet()
	for _, iface := range i.interfaces {
		for ref := range iface.References() {
			results.Add(ref)
		}
	}

	return results
}

func (i InterfaceImplementer) AsDeclarations(
	codeGenerationContext *CodeGenerationContext,
	typeName TypeName,
	_ []string,
) []dst.Decl {
	// interfaces must be ordered by name for deterministic output
	// (We sort them directly to skip future lookups)
	interfaces := make([]*InterfaceImplementation, 0, len(i.interfaces))
	for _, iface := range i.interfaces {
		interfaces = append(interfaces, iface)
	}

	sort.Slice(interfaces, func(i int, j int) bool {
		return interfaces[i].Name().name < interfaces[j].Name().name
	})

	result := make([]dst.Decl, 0, len(interfaces))
	for _, iface := range interfaces {
		result = append(result, i.generateInterfaceImplAssertion(codeGenerationContext, iface, typeName))

		functions := maps.Values(iface.functions)
		sort.Slice(functions, func(i int, j int) bool {
			return functions[i].Name() < functions[j].Name()
		})

		for _, f := range functions {
			decl := generateMethodDeclForFunction(typeName, f, codeGenerationContext)
			result = append(result, decl)
		}
	}

	return result
}

func (i InterfaceImplementer) Equals(other InterfaceImplementer, overrides EqualityOverrides) bool {
	if len(i.interfaces) != len(other.interfaces) {
		return false
	}

	for ifaceName, iface := range i.interfaces {
		otherIface, ok := other.interfaces[ifaceName]
		if !ok {
			return false
		}

		if !iface.Equals(otherIface, overrides) {
			return false
		}
	}

	return true
}

func (i InterfaceImplementer) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()
	for _, i := range i.interfaces {
		result.Merge(i.RequiredPackageReferences())
	}

	return result
}

func (i InterfaceImplementer) generateInterfaceImplAssertion(
	codeGenerationContext *CodeGenerationContext,
	iface *InterfaceImplementation,
	typeName TypeName,
) dst.Decl {
	ifacePackageName, err := codeGenerationContext.GetImportedPackageName(iface.name.PackageReference)
	if err != nil {
		panic(err)
	}

	var doc dst.Decorations
	if iface.annotation != "" {
		doc.Append("// " + iface.annotation)
		doc.Append("\n")
	}

	typeAssertion := &dst.GenDecl{
		Tok: token.VAR,
		Decs: dst.GenDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				Start:  doc,
			},
		},
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Type: astbuilder.Selector(
					dst.NewIdent(ifacePackageName),
					iface.name.name),
				Names: []*dst.Ident{
					dst.NewIdent("_"),
				},
				Values: astbuilder.Expressions(
					astbuilder.AddrOf(
						&dst.CompositeLit{
							Type: dst.NewIdent(typeName.name),
						})),
			},
		},
	}

	return typeAssertion
}

func (i InterfaceImplementer) copy() InterfaceImplementer {
	result := i

	result.interfaces = make(map[TypeName]*InterfaceImplementation, len(i.interfaces))
	for k, v := range i.interfaces {
		result.interfaces[k] = v
	}

	return result
}
