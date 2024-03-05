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

func NewValidateOwnerReferenceFunction(resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *ResourceFunction {
	return NewResourceFunction(
		"validateOwnerReference",
		resource,
		idFactory,
		validateOwnerReferences,
		astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference))
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

func validateResourceReferences(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body:          validateResourceReferencesBody(codeGenerationContext, receiverIdent),
	}

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates all resource references")
	return fn.DefineFunc(), nil
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
	body = append(body, astbuilder.CheckErrorAndReturn(astbuilder.Nil()))
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateResourceReferences",
				dst.NewIdent("refs"))))

	return body
}

func validateOwnerReferences(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	admissionPkg := codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body:          validateOwnerReferencesBody(codeGenerationContext, receiverIdent),
	}

	fn.AddReturn(astbuilder.QualifiedTypeName(admissionPkg, "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates the owner field")
	return fn.DefineFunc(), nil
}

// validateOwnerReferencesBody helps generate the body of the validateOwnerReferences function:
//
//	return genruntime.ValidateOwner(<resource>)
func validateOwnerReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, receiverIdent string) []dst.Stmt {
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	var body []dst.Stmt

	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateOwner",
				dst.NewIdent(receiverIdent))))

	return body
}

func validateWriteOncePropertiesFunction(
	resourceFn *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := resourceFn.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	runtimePackage := codeGenerationContext.MustGetImportedPackageName(astmodel.APIMachineryRuntimeReference)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body:          validateWriteOncePropertiesFunctionBody(receiver, codeGenerationContext, receiverIdent),
	}

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddParameter("old", astbuilder.QualifiedTypeName(runtimePackage, "Object"))
	fn.AddComments("validates all WriteOnce properties")

	return fn.DefineFunc(), nil
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

	cast := astbuilder.TypeAssert(obj, dst.NewIdent("old"), astbuilder.PointerTo(receiver.AsType(codeGenerationContext)))
	checkAssert := astbuilder.ReturnIfNotOk(astbuilder.Nil(), astbuilder.Nil())

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

func validateOptionalConfigMapReferences(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body:          validateOptionalConfigMapReferencesBody(codeGenerationContext, receiverIdent),
	}

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates all optional configmap reference pairs to ensure that at most 1 is set")
	return fn.DefineFunc(), nil
}

// validateOptionalConfigMapReferencesBody helps generate the body of the validateResourceReferences function:
//
//	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&<resource>.Spec)
//	if err != nil {
//		return nil, err
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
	body = append(body, astbuilder.CheckErrorAndReturn(astbuilder.Nil()))
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateOptionalConfigMapReferences",
				dst.NewIdent("refs"))))

	return body
}
