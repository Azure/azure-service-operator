/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
)

// CreateEmptyARMValueFunc represents a function that creates
// an empty value suitable for calling PopulateFromARM against.
// It should be equivalent to ConvertToARM("") on a default struct value.
type CreateEmptyARMValueFunc struct {
	idFactory   astmodel.IdentifierFactory
	armTypeName astmodel.TypeName
}

var _ astmodel.Function = &CreateEmptyARMValueFunc{}

func (f CreateEmptyARMValueFunc) AsFunc(c *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	fn := &astbuilder.FuncDetails{
		Name:          "CreateEmptyARMValue",
		ReceiverIdent: f.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported),
		ReceiverType: &dst.StarExpr{
			X: dst.NewIdent(receiver.Name()),
		},
		Body: []dst.Stmt{
			astbuilder.Returns(&dst.CompositeLit{
				Type: dst.NewIdent(f.armTypeName.Name()),
			}),
		},
	}

	fn.AddReturns("interface{}")
	fn.AddComments("returns an empty ARM value suitable for deserializing into")

	return fn.DefineFunc()
}

func (f CreateEmptyARMValueFunc) Equals(other astmodel.Function) bool {
	otherFunc, ok := other.(CreateEmptyARMValueFunc)
	return ok && f.armTypeName == otherFunc.armTypeName
}

func (f CreateEmptyARMValueFunc) Name() string {
	return "CreateEmptyARMValue"
}

func (f CreateEmptyARMValueFunc) References() astmodel.TypeNameSet {
	return nil
}

func (f CreateEmptyARMValueFunc) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet()
}
