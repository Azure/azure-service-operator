/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astbuilder"
)

func NewDefaultAzureNameFunction(resource *ResourceType, idFactory IdentifierFactory) *resourceFunction {
	return &resourceFunction{
		name:             "defaultAzureName",
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           defaultAzureNameFunction,
		requiredPackages: NewPackageReferenceSet(GenRuntimeReference),
	}
}

// defaultAzureNameFunction returns a function that defaults the AzureName property of the resource spec
// to the Name property of the resource spec
func defaultAzureNameFunction(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	specSelector := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Spec"),
	}

	azureNameProp := &dst.SelectorExpr{
		X:   specSelector,
		Sel: dst.NewIdent(AzureNameProperty),
	}

	nameProp := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Name"), // this comes from ObjectMeta
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			&dst.IfStmt{
				Cond: &dst.BinaryExpr{
					X:  dst.Clone(azureNameProp).(dst.Expr),
					Op: token.EQL,
					Y:  &dst.BasicLit{Kind: token.STRING, Value: "\"\""},
				},
				Body: astbuilder.StatementBlock(
					astbuilder.SimpleAssignment(
						azureNameProp,
						nameProp)),
			},
		},
	}

	fn.AddComments("defaults the Azure name of the resource to the Kubernetes name")
	return fn.DefineFunc()
}
