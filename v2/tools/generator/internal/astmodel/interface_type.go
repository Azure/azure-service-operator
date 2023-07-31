/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/readonly"
)

var EmptyInterfaceType = NewInterfaceType()

// InterfaceType represents an interface
type InterfaceType struct {
	functions readonly.Map[string, Function]
}

var _ Type = &InterfaceType{}

// NewInterfaceType creates a new InterfaceType
func NewInterfaceType(functions ...Function) *InterfaceType {
	var f readonly.Map[string, Function]
	if len(functions) > 0 {
		inner := make(map[string]Function, len(functions))
		for _, v := range functions {
			inner[v.Name()] = v
		}
		f = readonly.CreateMap(inner)
	} else {
		f = readonly.EmptyMap[string, Function]()
	}

	return &InterfaceType{
		functions: f,
	}
}

// WithFunction creates a new InterfaceType with a function (method) attached to it
func (i *InterfaceType) WithFunction(function Function) *InterfaceType {
	// Create a copy of objectType to preserve immutability
	result := i.copy()
	result.functions = result.functions.With(function.Name(), function)

	return result
}

// Functions returns all the function definitions
// A sorted slice is returned to preserve immutability and provide determinism
func (i *InterfaceType) Functions() []Function {
	functions := i.functions.Values()

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	return functions
}

// References returns any type referenced by this interface
func (i *InterfaceType) References() TypeNameSet {
	result := NewTypeNameSet()
	i.functions.ForEach(
		func(name string, method Function) {
			result.AddAll(method.References())
		})
	return result
}

// AsType renders the Go abstract syntax tree for the interface type
func (i *InterfaceType) AsType(codeGenerationContext *CodeGenerationContext) dst.Expr {
	fields := make([]*dst.Field, 0, i.functions.Len())

	functions := i.functions.Values()

	sort.Slice(functions, func(i int, j int) bool {
		return functions[i].Name() < functions[j].Name()
	})

	for _, method := range functions {
		field := functionToField(codeGenerationContext, method.Name(), method)
		fields = append(fields, field)
	}

	return &dst.InterfaceType{
		Methods: &dst.FieldList{
			List: fields,
		},
	}
}

func (i *InterfaceType) AsDeclarations(codeGenerationContext *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	declaration := &dst.GenDecl{
		Decs: dst.GenDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
			},
		},
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent(declContext.Name.Name()),
				Type: i.AsType(codeGenerationContext),
			},
		},
	}

	astbuilder.AddUnwrappedComments(&declaration.Decs.Start, declContext.Description)
	AddValidationComments(&declaration.Decs.Start, declContext.Validations)

	result := []dst.Decl{declaration}
	return result
}

// AsZero always panics; Interface does not have a zero type
func (i *InterfaceType) AsZero(_ TypeDefinitionSet, _ *CodeGenerationContext) dst.Expr {
	panic("cannot create a zero value for an interface type")
}

// RequiredPackageReferences returns the set of package references which this interface requires
func (i *InterfaceType) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()

	i.functions.ForEach(
		func(_ string, method Function) {
			result.Merge(method.RequiredPackageReferences())
		})

	return result
}

func (i *InterfaceType) Equals(t Type, overrides EqualityOverrides) bool {
	if i == t {
		return true // short circuit
	}

	other, ok := t.(*InterfaceType)
	if !ok {
		return false
	}

	equalityFunc := func(l, r Function) bool {
		return l.Equals(r, overrides)
	}
	if !i.functions.Equals(other.functions, equalityFunc) {
		return false
	}

	return true
}

// String implements fmt.Stringer
func (i *InterfaceType) String() string {
	return "(interface)"
}

// WriteDebugDescription adds a description of the current InterfaceType to the passed builder.
func (i *InterfaceType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if i == nil {
		builder.WriteString("<nilInterface>")
	} else {
		builder.WriteString("Interface")
	}
}

func (i *InterfaceType) copy() *InterfaceType {
	result := &InterfaceType{
		// no need to clone these, they are all readonly
		functions: i.functions,
	}

	return result
}

// AsInterfaceType unwraps any wrappers around the provided type and returns either the underlying InterfaceType and true,
// or nil and false.
func AsInterfaceType(t Type) (*InterfaceType, bool) {
	if iface, ok := t.(*InterfaceType); ok {
		return iface, true
	}

	if wrapper, ok := t.(MetaType); ok {
		return AsInterfaceType(wrapper.Unwrap())
	}

	return nil, false
}

func functionToField(codeGenerationContext *CodeGenerationContext, name string, function Function) *dst.Field {
	var names []*dst.Ident
	if name != "" {
		names = []*dst.Ident{dst.NewIdent(name)}
	}

	f := function.AsFunc(codeGenerationContext, nil) // Empty typename here because we have no receiver

	return &dst.Field{
		Names: names,
		Type:  f.Type,
	}
}
