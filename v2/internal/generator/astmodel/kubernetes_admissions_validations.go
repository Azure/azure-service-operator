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

func NewValidateResourceReferencesFunction(resource *ResourceType, idFactory IdentifierFactory) *resourceFunction {
	return &resourceFunction{
		name:             "validateResourceReferences",
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           validateResourceReferences,
		requiredPackages: NewPackageReferenceSet(GenRuntimeReference, ReflectHelpersReference),
	}
}

func validateResourceReferences(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("error"),
			},
		},
		Body: validateResourceReferencesBody(codeGenerationContext, receiverIdent),
	}

	fn.AddComments("validates all resource references")
	return fn.DefineFunc()
}

// validateResourceReferencesBody helps generate the body of the validateResourceReferences function:
// 	refs, err := reflecthelpers.FindResourceReferences(&<resource>.Spec)
//	if err != nil {
//		return err
//	}
//	return genruntime.ValidateResourceReferences(refs)
func validateResourceReferencesBody(codeGenerationContext *CodeGenerationContext, receiverIdent string) []dst.Stmt {
	reflectHelpers, err := codeGenerationContext.GetImportedPackageName(ReflectHelpersReference)
	if err != nil {
		panic(err)
	}

	genRuntime, err := codeGenerationContext.GetImportedPackageName(GenRuntimeReference)
	if err != nil {
		panic(err)
	}
	var body []dst.Stmt

	body = append(
		body,
		astbuilder.SimpleAssignmentWithErr(
			dst.NewIdent("refs"),
			token.DEFINE,
			astbuilder.CallQualifiedFunc(
				reflectHelpers,
				"FindResourceReferences",
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))))
	body = append(body, astbuilder.CheckErrorAndReturn())
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateResourceReferences",
				dst.NewIdent("refs"))))

	return body
}
