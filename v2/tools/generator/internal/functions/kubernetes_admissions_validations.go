/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewValidateResourceReferencesFunction creates a function for validating resource references
//
//	func (group *<obj>) validateResourceReferences(ctx context.Context, obj *<obj>) (admission.Warnings, error) {
//		refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
//		if err != nil {
//			return nil, err
//		}
//		return genruntime.ValidateResourceReferences(refs)
//	}
func NewValidateResourceReferencesFunction(resource astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *ValidateFunction {
	return NewValidateFunction(
		"validateResourceReferences",
		resource.Name(),
		idFactory,
		validateResourceReferences,
		astmodel.GenRuntimeReference,
		astmodel.ReflectHelpersReference)
}

// NewValidateOwnerReferenceFunction creates a function for validating the owner reference, if it exists
//
//	func (domain *<obj>) validateOwnerReference(ctx context.Context, obj *<obj>) (admission.Warnings, error) {
//		return genruntime.ValidateOwner(obj)
//	}
func NewValidateOwnerReferenceFunction(resource astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *ValidateFunction {
	return NewValidateFunction(
		"validateOwnerReference",
		resource.Name(),
		idFactory,
		validateOwnerReferences,
		astmodel.GenRuntimeReference)
}

// NewValidateWriteOncePropertiesFunction creates a function for validating write-once properties
//
//	func (group *<obj>) validateWriteOnceProperties(ctx context.Context, oldObj *<obj>, newObj *<obj>) (admission.Warnings, error) {
//		return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
//	}
func NewValidateWriteOncePropertiesFunction(resource astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *ValidateFunction {
	return NewValidateFunction(
		"validateWriteOnceProperties",
		resource.Name(),
		idFactory,
		validateWriteOncePropertiesFunction)
}

// NewValidateOptionalConfigMapReferenceFunction creates a function for validating optional configmap references
//
//	func (encryptionSet *<obj>) validateOptionalConfigMapReferences(ctx context.Context, obj *<obj>) (admission.Warnings, error) {
//		refs, err := reflecthelpers.FindOptionalConfigMapReferences(&obj.Spec)
//		if err != nil {
//			return nil, err
//		}
//		return configmaps.ValidateOptionalReferences(refs)
//	}
func NewValidateOptionalConfigMapReferenceFunction(resource astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *ValidateFunction {
	return NewValidateFunction(
		"validateOptionalConfigMapReferences",
		resource.Name(),
		idFactory,
		validateOptionalConfigMapReferences,
		astmodel.GenRuntimeConfigMapsReference,
		astmodel.ReflectHelpersReference)
}

func validateResourceReferences(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objectIdent := "obj"
	contextIdent := "ctx"

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          validateResourceReferencesBody(codeGenerationContext, objectIdent),
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	typedObjExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(objectIdent, astbuilder.PointerTo(typedObjExpr))

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
func validateResourceReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, objIdent string) []dst.Stmt {
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
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(objIdent), "Spec")))))
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
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objectIdent := "obj"
	contextIdent := "ctx"

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating receiver type expression for %s", receiver)
	}

	admissionPkg := codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          validateOwnerReferencesBody(codeGenerationContext, objectIdent),
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	typedObjExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(objectIdent, astbuilder.PointerTo(typedObjExpr))

	fn.AddReturn(astbuilder.QualifiedTypeName(admissionPkg, "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates the owner field")
	return fn.DefineFunc(), nil
}

// validateOwnerReferencesBody helps generate the body of the validateOwnerReferences function:
//
//	return genruntime.ValidateOwner(<resource>)
func validateOwnerReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, objIdent string) []dst.Stmt {
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	var body []dst.Stmt

	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntime,
				"ValidateOwner",
				dst.NewIdent(objIdent))))

	return body
}

func validateWriteOncePropertiesFunction(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objIdent := "newObj"
	oldObjIdent := "oldObj"
	contextIdent := "ctx"

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating receiver type expression for %s", receiver)
	}

	body := validateWriteOncePropertiesFunctionBody(codeGenerationContext, oldObjIdent, objIdent)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          body,
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	typedObjExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(oldObjIdent, astbuilder.PointerTo(typedObjExpr))
	fn.AddParameter(objIdent, astbuilder.PointerTo(typedObjExpr))

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates all WriteOnce properties")

	return fn.DefineFunc(), nil
}

// validateWriteOncePropertiesFunctionBody helps generate the body of the validateWriteOncePropertiesFunctionBody function:
//
//	oldObj, ok := old.(*T)
//	if !ok {
//	    return nil
//	}
//
// return genruntime.ValidateWriteOnceProperties(oldObj, <objIdent>)
func validateWriteOncePropertiesFunctionBody(
	codeGenerationContext *astmodel.CodeGenerationContext,
	oldObjIdent string,
	newObjIdent string,
) []dst.Stmt {
	genRuntime := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	returnStmt := astbuilder.Returns(
		astbuilder.CallQualifiedFunc(
			genRuntime,
			"ValidateWriteOnceProperties",
			dst.NewIdent(oldObjIdent),
			dst.NewIdent(newObjIdent)))
	return astbuilder.Statements(returnStmt)
}

func validateOptionalConfigMapReferences(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objectIdent := "obj"
	contextIdent := "ctx"

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          validateOptionalConfigMapReferencesBody(codeGenerationContext, objectIdent),
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	typedObjExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(objectIdent, astbuilder.PointerTo(typedObjExpr))

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
func validateOptionalConfigMapReferencesBody(codeGenerationContext *astmodel.CodeGenerationContext, objIdent string) []dst.Stmt {
	reflectHelpers := codeGenerationContext.MustGetImportedPackageName(astmodel.ReflectHelpersReference)
	genRuntimeConfigMaps := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeConfigMapsReference)

	var body []dst.Stmt

	body = append(
		body,
		astbuilder.SimpleAssignmentWithErr(
			dst.NewIdent("refs"),
			token.DEFINE,
			astbuilder.CallQualifiedFunc(
				reflectHelpers,
				"FindOptionalConfigMapReferences",
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(objIdent), "Spec")))))
	body = append(body, astbuilder.CheckErrorAndReturn(astbuilder.Nil()))
	body = append(
		body,
		astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				genRuntimeConfigMaps,
				"ValidateOptionalReferences",
				dst.NewIdent("refs"))))

	return body
}
