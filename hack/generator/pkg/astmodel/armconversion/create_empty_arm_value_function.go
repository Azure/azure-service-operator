/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"go/ast"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// CreateEmptyArmValueFunc represents a function that creates
// an empty value suitable for calling PopulateFromARM against.
// It should be equivalent to ConvertToARM("") on a default struct value.
type CreateEmptyArmValueFunc struct {
	idFactory   astmodel.IdentifierFactory
	armTypeName astmodel.TypeName
}

var _ astmodel.Function = &CreateEmptyArmValueFunc{}

func (f CreateEmptyArmValueFunc) AsFunc(c *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *ast.FuncDecl {
	fn := &astbuilder.FuncDetails{
		Name:          ast.NewIdent("CreateEmptyArmValue"),
		ReceiverIdent: ast.NewIdent(f.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)),
		ReceiverType: &ast.StarExpr{
			X: ast.NewIdent(receiver.Name()),
		},
		Body: []ast.Stmt{
			astbuilder.Returns(&ast.CompositeLit{
				Type: ast.NewIdent(f.armTypeName.Name()),
			}),
		},
	}

	fn.AddReturns("interface{}")
	fn.AddComments("returns an empty ARM value suitable for deserializing into")

	return fn.DefineFunc()
}

func (f CreateEmptyArmValueFunc) Equals(other astmodel.Function) bool {
	otherFunc, ok := other.(CreateEmptyArmValueFunc)
	return ok && f.armTypeName == otherFunc.armTypeName
}

func (f CreateEmptyArmValueFunc) Name() string {
	return "CreateEmptyArmValue"
}

func (f CreateEmptyArmValueFunc) References() astmodel.TypeNameSet {
	return nil
}

func (f CreateEmptyArmValueFunc) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet()
}
