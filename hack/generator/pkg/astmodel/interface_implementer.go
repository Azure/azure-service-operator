/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"

	"github.com/dave/dst"
)

type InterfaceImplementer struct {
	interfaces map[TypeName]*InterfaceImplementation
}

// MakeInterfaceImplementer returns an interface implementer
func MakeInterfaceImplementer() InterfaceImplementer {
	return InterfaceImplementer{}
}

func (i InterfaceImplementer) HasInterface(name TypeName) bool {
	_, ok := i.interfaces[name]
	return ok
}

func (i InterfaceImplementer) WithoutInterface(name TypeName) InterfaceImplementer {
	result := i.copy()
	delete(result.interfaces, name)
	return result
}

// WithInterface creates a new ObjectType with a function (method) attached to it
func (i InterfaceImplementer) WithInterface(iface *InterfaceImplementation) InterfaceImplementer {
	result := i.copy()
	result.interfaces[iface.Name()] = iface

	return result
}

func (i InterfaceImplementer) References() TypeNameSet {
	var results TypeNameSet
	for _, iface := range i.interfaces {
		for ref := range iface.References() {
			results = results.Add(ref)
		}
	}

	return results
}

func (i InterfaceImplementer) AsDeclarations(
	codeGenerationContext *CodeGenerationContext,
	typeName TypeName,
	_ []string) []dst.Decl {

	var result []dst.Decl

	// interfaces must be ordered by name for deterministic output
	// (We sort them directly to skip future lookups)
	var interfaces []*InterfaceImplementation
	for _, iface := range i.interfaces {
		interfaces = append(interfaces, iface)
	}

	sort.Slice(interfaces, func(i int, j int) bool {
		return interfaces[i].Name().name < interfaces[j].Name().name
	})

	for _, iface := range interfaces {
		result = append(result, i.generateInterfaceImplAssertion(codeGenerationContext, iface, typeName))

		var functions []Function
		for _, f := range iface.functions {
			functions = append(functions, f)
		}

		sort.Slice(functions, func(i int, j int) bool {
			return functions[i].Name() < functions[j].Name()
		})

		for _, f := range functions {
			result = append(result, f.AsFunc(codeGenerationContext, typeName))
		}
	}

	return result
}

func (i InterfaceImplementer) Equals(other InterfaceImplementer) bool {

	if len(i.interfaces) != len(other.interfaces) {
		return false
	}

	for ifaceName, iface := range i.interfaces {
		otherIface, ok := other.interfaces[ifaceName]
		if !ok {
			return false
		}

		if !iface.Equals(otherIface) {
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
	typeName TypeName) dst.Decl {

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
				Type: &dst.SelectorExpr{
					X:   dst.NewIdent(ifacePackageName),
					Sel: dst.NewIdent(iface.name.name),
				},
				Names: []*dst.Ident{
					dst.NewIdent("_"),
				},
				Values: []dst.Expr{
					&dst.UnaryExpr{
						Op: token.AND,
						X: &dst.CompositeLit{
							Type: dst.NewIdent(typeName.name),
						},
					},
				},
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
