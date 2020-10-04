/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/pkg/errors"
	"go/ast"
	"go/token"
)

// These are some magical field names which we're going to use or generate
const (
	AzureNameProperty = "AzureName"
	OwnerProperty     = "Owner"
)

// NewArmTransformerImpl creates a new interface with the specified ARM conversion functions
func NewKubernetesResourceInterfaceImpl(
	idFactory IdentifierFactory,
	spec *ObjectType) (*InterfaceImplementation, error) {

	// Check the spec first to ensure it looks how we expect
	ownerProperty := idFactory.CreatePropertyName(OwnerProperty, Exported)
	_, ok := spec.Property(ownerProperty)
	if !ok {
		return nil, errors.Errorf("Resource spec doesn't have %q property", ownerProperty)
	}

	azureNameProperty := idFactory.CreatePropertyName(AzureNameProperty, Exported)
	_, ok = spec.Property(azureNameProperty)
	if !ok {
		return nil, errors.Errorf("Resource spec doesn't have %q property", azureNameProperty)
	}

	ownerFunc := &objectFunction{
		name:      OwnerProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    ownerFunction,
	}

	azureNameFunc := &objectFunction{
		name:      AzureNameProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    azureNameFunction,
	}

	result := NewInterfaceImplementation(
		MakeTypeName(MakeGenRuntimePackageReference(), "KubernetesResource"),
		ownerFunc,
		azureNameFunc)

	return result, nil
}

// objectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type objectFunction struct {
	name      string
	o         *ObjectType
	idFactory IdentifierFactory

	asFunc func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl
}

var _ Function = &objectFunction{}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *objectFunction) Name() string {
	return k.name
}

func (k *objectFunction) RequiredImports() []PackageReference {
	// We only require GenRuntime
	return []PackageReference{
		MakeGenRuntimePackageReference(),
	}
}

func (k *objectFunction) References() TypeNameSet {
	return k.o.References()
}

func (k *objectFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *ast.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

func (k *objectFunction) Equals(f Function) bool {
	typedF, ok := f.(*objectFunction)
	if !ok {
		return false
	}

	return k.o.Equals(typedF.o)
}

func ownerFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	receiverIdent := ast.NewIdent(k.idFactory.CreateIdentifier(receiver.Name(), NotExported))
	receiverType := receiver.AsType(codeGenerationContext)

	specSelector := &ast.SelectorExpr{
		X:   receiverIdent,
		Sel: ast.NewIdent("Spec"),
	}

	groupIdent := ast.NewIdent("group")
	kindIdent := ast.NewIdent("kind")

	return astbuilder.DefineFunc(
		astbuilder.FuncDetails{
			Name:          ast.NewIdent(methodName),
			ReceiverIdent: receiverIdent,
			ReceiverType: &ast.StarExpr{
				X: receiverType,
			},
			Comment: "returns the ResourceReference of the owner, or nil if there is no owner",
			Params:  nil,
			Returns: []*ast.Field{
				{
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(GenRuntimePackageName),
							Sel: ast.NewIdent("ResourceReference"),
						},
					},
				},
			},
			Body: []ast.Stmt{
				lookupGroupAndKindStmt(groupIdent, kindIdent, specSelector),
				&ast.ReturnStmt{
					Results: []ast.Expr{
						createResourceReference(groupIdent, kindIdent, specSelector),
					},
				},
			},
		})
}

func lookupGroupAndKindStmt(
	groupIdent *ast.Ident,
	kindIdent *ast.Ident,
	specSelector *ast.SelectorExpr) *ast.AssignStmt {

	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			groupIdent,
			kindIdent,
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent(GenRuntimePackageName),
					Sel: ast.NewIdent("LookupOwnerGroupKind"),
				},
				Args: []ast.Expr{
					specSelector,
				},
			},
		},
	}
}

func createResourceReference(
	groupIdent *ast.Ident,
	kindIdent *ast.Ident,
	specSelector *ast.SelectorExpr) ast.Expr {

	return astbuilder.AddrOf(
		&ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(GenRuntimePackageName),
				Sel: ast.NewIdent("ResourceReference"),
			},
			Elts: []ast.Expr{
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Name"),
					Value: &ast.SelectorExpr{
						X: &ast.SelectorExpr{
							X:   specSelector,
							Sel: ast.NewIdent(OwnerProperty),
						},
						Sel: ast.NewIdent("Name"),
					},
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("Group"),
					Value: groupIdent,
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("Kind"),
					Value: kindIdent,
				},
			},
		})
}

func azureNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	receiverIdent := ast.NewIdent(k.idFactory.CreateIdentifier(receiver.Name(), NotExported))
	receiverType := receiver.AsType(codeGenerationContext)

	specSelector := &ast.SelectorExpr{
		X:   receiverIdent,
		Sel: ast.NewIdent("Spec"),
	}

	return astbuilder.DefineFunc(
		astbuilder.FuncDetails{
			Name:          ast.NewIdent(methodName),
			ReceiverIdent: receiverIdent,
			ReceiverType: &ast.StarExpr{
				X: receiverType,
			},
			Comment: "returns the Azure name of the resource",
			Params:  nil,
			Returns: []*ast.Field{
				{
					Type: ast.NewIdent("string"),
				},
			},
			Body: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.SelectorExpr{
							X:   specSelector,
							Sel: ast.NewIdent(AzureNameProperty),
						},
					},
				},
			},
		})
}
