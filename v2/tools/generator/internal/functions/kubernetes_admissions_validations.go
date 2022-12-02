/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func NewValidateResourceReferencesFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateResourceReferences",
		resource,
		idFactory,
		validateResourceReferences,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference, astmodel.ReflectHelpersReference))
}

func NewValidateWriteOncePropertiesFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateWriteOnceProperties",
		resource,
		idFactory,
		validateWriteOncePropertiesFunction,
		astmodel.NewPackageReferenceSet())
}

func NewValidateOptionalConfigMapReferenceFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateOptionalConfigMapReferences",
		resource,
		idFactory,
		validateOptionalConfigMapReferences,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference, astmodel.ReflectHelpersReference))
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
//
//	refs, err := reflecthelpers.FindResourceReferences(&<resource>.Spec)
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

func validateWriteOncePropertiesFunction(resourceFn *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {

	receiverIdent := resourceFn.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	runtimePackage := codeGenerationContext.MustGetImportedPackageName(astmodel.APIMachineryRuntimeReference)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.Dereference(receiverType),
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("error"),
			},
		},
		Body: validateWriteOncePropertiesFunctionBody(receiver, codeGenerationContext, receiverIdent),
	}

	fn.AddParameter("old", astbuilder.QualifiedTypeName(runtimePackage, "Object"))
	fn.AddComments("validates all WriteOnce properties")

	return fn.DefineFunc()
}

// validateWriteOncePropertiesFunctionBody helps generate the body of the validateWriteOncePropertiesFunctionBody function:
//
//	oldObj, ok := old.(*Receiver)
//	if !ok {
//	    return nil
//	}
//
// return genruntime.ValidateWriteOnceProperties(oldObj, <receiverIndent>)
func validateWriteOncePropertiesFunctionBody(receiver astmodel.TypeName, codeGenerationContext *astmodel.CodeGenerationContext, receiverIdent string) []dst.Stmt {

	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	obj := dst.NewIdent("oldObj")

	cast := astbuilder.TypeAssert(obj, dst.NewIdent("old"), astbuilder.Dereference(receiver.AsType(codeGenerationContext)))
	checkAssert := astbuilder.ReturnIfNotOk(astbuilder.Nil())

	returnStmt := astbuilder.Returns(
		astbuilder.CallQualifiedFunc(
			genRuntime,
			"ValidateWriteOnceProperties",
			obj,
			dst.NewIdent(receiverIdent)))
	returnStmt.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		cast,
		checkAssert,
		returnStmt)
}

func validateOptionalConfigMapReferences(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
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
		Body: validateOptionalConfigMapReferencesBody(codeGenerationContext, receiverIdent),
	}

	fn.AddComments("validates all optional configmap reference pairs to ensure that at most 1 is set")
	return fn.DefineFunc()
}

// validateOptionalConfigMapReferencesBody helps generate the body of the validateResourceReferences function:
//
//	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&<resource>.Spec)
//	if err != nil {
//		return err
//	}
//	return genruntime.ValidateOptionalConfigMapReferences(refs)
func validateOptionalConfigMapReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, receiverIdent string) []dst.Stmt {
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
				"FindOptionalConfigMapReferences",
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))))
	body = append(body, astbuilder.CheckErrorAndReturn())
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateOptionalConfigMapReferences",
				dst.NewIdent("refs"))))

	return body
}
