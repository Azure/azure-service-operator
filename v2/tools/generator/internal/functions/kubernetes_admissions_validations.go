/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/dave/dst"
)

func NewValidateResourceReferencesFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateResourceReferences",
		resource,
		idFactory,
		validateResourceReferences,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference, astmodel.ReflectHelpersReference))
}

func NewValidateImmutablePropertiesFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateImmutableProperties",
		resource,
		idFactory,
		validateImmutablePropertiesFunction,
		astmodel.NewPackageReferenceSet())
}

func validateResourceReferences(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
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
func validateResourceReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, receiverIdent string) []dst.Stmt {
	reflectHelpers := codeGenerationContext.MustGetImportedPackageName(astmodel.ReflectHelpersReference)
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

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

func validateImmutablePropertiesFunction(resourceFn *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {

	receiverIdent := resourceFn.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	runtimePackage := codeGenerationContext.MustGetImportedPackageName(astmodel.APIMachineryRuntimeReference)

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
		Body: validateImmutablePropertiesFunctionBody(receiver, codeGenerationContext, receiverIdent),
	}

	fn.AddParameter("old", astbuilder.QualifiedTypeName(runtimePackage, "Object"))
	fn.AddComments("validates all immutable properties")

	return fn.DefineFunc()
}

func validateImmutablePropertiesFunctionBody(receiver astmodel.TypeName, codeGenerationContext *astmodel.CodeGenerationContext, receiverIdent string) []dst.Stmt {

	obj := dst.NewIdent("oldObj")

	var validations []dst.Stmt

	cast := astbuilder.TypeAssert(obj, dst.NewIdent("old"), astbuilder.Dereference(receiver.AsType(codeGenerationContext)))
	checkAssert := astbuilder.ReturnIfNotOk(astbuilder.Nil())

	azureNameRef := []string{"AzureName()"}
	azureNameCheck := checkImmutableProperty(obj, receiverIdent, azureNameRef, codeGenerationContext)

	resourceOwnerRef := []string{"Owner()", "Name"}
	resourceOwnerCheck := checkImmutableProperty(obj, receiverIdent, resourceOwnerRef, codeGenerationContext)

	validations = append(validations, azureNameCheck, resourceOwnerCheck)

	return astbuilder.Statements(cast, checkAssert, validations, astbuilder.ReturnNoError())

}

func checkImmutableProperty(obj *dst.Ident, receiverIdent string, ref []string, codeGenerationContext *astmodel.CodeGenerationContext) *dst.IfStmt {
	errorsPkg := codeGenerationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	oldIndent := astbuilder.Selector(obj, ref...)
	newIndent := astbuilder.Selector(dst.NewIdent(receiverIdent), ref...)

	errorMessage := astbuilder.StringLiteralf(
		"update for '%s' is not allowed", getRefString(ref))

	azureNameCheck := astbuilder.ReturnIfExpr(
		astbuilder.AreNotEqual(oldIndent, newIndent),
		astbuilder.CallQualifiedFunc(errorsPkg, "New", errorMessage))
	azureNameCheck.Decorations().Before = dst.EmptyLine

	return azureNameCheck
}

func getRefString(ref []string) string {

	return strings.Join(ref, ".")

}
