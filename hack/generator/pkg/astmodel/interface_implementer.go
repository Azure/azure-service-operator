/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/token"
	"sort"
)

type InterfaceImplementer struct {
	interfaces map[TypeName]*InterfaceImplementation
}

// MakeInterfaceImplementer returns an interface implementer
func MakeInterfaceImplementer() InterfaceImplementer {
	return InterfaceImplementer{}
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
	_ []string) []ast.Decl {

	var result []ast.Decl

	// First interfaces must be ordered by name for deterministic output
	var ifaceNames []TypeName
	for ifaceName := range i.interfaces {
		ifaceNames = append(ifaceNames, ifaceName)
	}

	sort.Slice(ifaceNames, func(i int, j int) bool {
		return ifaceNames[i].name < ifaceNames[j].name
	})

	for _, ifaceName := range ifaceNames {
		iface := i.interfaces[ifaceName]

		result = append(result, i.generateInterfaceImplAssertion(codeGenerationContext, iface, typeName))

		var funcNames []string
		for funcName := range iface.functions {
			funcNames = append(funcNames, funcName)
		}

		sort.Strings(funcNames)

		for _, methodName := range funcNames {
			function := iface.functions[methodName]
			result = append(result, function.AsFunc(codeGenerationContext, typeName, methodName))
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

func (i InterfaceImplementer) RequiredImports() []PackageReference {
	var result []PackageReference

	for _, i := range i.interfaces {
		result = append(result, i.RequiredImports()...)
	}

	return result
}

func (i InterfaceImplementer) generateInterfaceImplAssertion(
	codeGenerationContext *CodeGenerationContext,
	iface *InterfaceImplementation,
	typeName TypeName) ast.Decl {

	ifacePackageName, err := codeGenerationContext.GetImportedPackageName(iface.name.PackageReference)
	if err != nil {
		panic(err)
	}

	typeAssertion := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent(ifacePackageName),
					Sel: ast.NewIdent(iface.name.name),
				},
				Names: []*ast.Ident{
					ast.NewIdent("_"),
				},
				Values: []ast.Expr{
					&ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: ast.NewIdent(typeName.name),
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
