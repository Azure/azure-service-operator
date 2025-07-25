/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ValidationKind determines when a particular validation should be run
type ValidationKind string

const (
	ValidationKindCreate = ValidationKind("Create")
	ValidationKindUpdate = ValidationKind("Update")
	ValidationKindDelete = ValidationKind("Delete")
)

type ValidateFunction = DataFunction[astmodel.InternalTypeName]

var validateFunctionRequiredPackages = []astmodel.PackageReference{
	astmodel.ContextReference,
	astmodel.ControllerRuntimeAdmission,   // admissions.Warnings return
	astmodel.ControllerRuntimeWebhook,     // Used for interface assertion
	astmodel.APIMachineryRuntimeReference, // Used for runtime.Object parameter of Defaulter interface
}

func NewValidateFunction(
	name string,
	data astmodel.InternalTypeName,
	idFactory astmodel.IdentifierFactory,
	asFunc DataFunctionHandler[astmodel.InternalTypeName],
	requiredPackages ...astmodel.PackageReference,
) *ValidateFunction {
	// Add the default set of required packages
	requiredPackages = append(requiredPackages, validateFunctionRequiredPackages...)

	return NewDataFunction[astmodel.InternalTypeName](
		name,
		data,
		idFactory,
		asFunc,
		requiredPackages...)
}

// ValidatorBuilder helps in building an interface implementation for webhook.CustomValidator.
type ValidatorBuilder struct {
	resourceName astmodel.InternalTypeName
	idFactory    astmodel.IdentifierFactory

	validations map[ValidationKind][]*ValidateFunction
}

// NewValidatorBuilder creates a new ValidatorBuilder for the given object type.
func NewValidatorBuilder(
	resourceName astmodel.InternalTypeName,
	idFactory astmodel.IdentifierFactory,
) *ValidatorBuilder {
	return &ValidatorBuilder{
		resourceName: resourceName,
		idFactory:    idFactory,
		validations: map[ValidationKind][]*ValidateFunction{
			ValidationKindCreate: nil,
			ValidationKindUpdate: nil,
			ValidationKindDelete: nil,
		},
	}
}

// AddValidation adds a validation function to the set of validation functions to be applied to the given object.
func (v *ValidatorBuilder) AddValidation(kind ValidationKind, f *ValidateFunction) {
	if !f.Data().Equals(v.resourceName, astmodel.EqualityOverrides{}) {
		panic(fmt.Sprintf("cannot add validation function on non-matching resources. Expected %s, got %s", v.resourceName, f.Data()))
	}
	v.validations[kind] = append(v.validations[kind], f)
}

// ToInterfaceImplementation creates an InterfaceImplementation that implements the webhook.CustomValidator interface.
// This implementation includes calls to all validations registered with this ValidatorBuilder via the AddValidation function,
// as well as helper functions that allow additional handcrafted validations to be injected by
// implementing the genruntime.Validator interface.
func (v *ValidatorBuilder) ToInterfaceImplementation() *astmodel.InterfaceImplementation {
	group, version := v.resourceName.InternalPackageReference().GroupVersion()

	// e.g. group = "microsoft.network.azure.com"
	// e.g. resource = "backendaddresspools"
	// e.g. version = "v1"

	resource := v.resourceName.Name()

	group = strings.ToLower(group + astmodel.GroupSuffix)
	nonPluralResource := strings.ToLower(resource)
	resource = strings.ToLower(v.resourceName.Plural().Name())

	// e.g. "validate-microsoft-network-azure-com-v1-backendaddresspool"
	// note that this must match _exactly_ how controller-runtime generates the path,
	// or it will not work!
	path := fmt.Sprintf("/validate-%s-%s-%s", strings.ReplaceAll(group, ".", "-"), version, nonPluralResource)

	// e.g.  "default.v123.backendaddresspool.azure.com"
	name := fmt.Sprintf("validate.%s.%s.%s", version, resource, group)

	annotation := fmt.Sprintf(
		"+kubebuilder:webhook:path=%s,mutating=false,sideEffects=None,"+
			"matchPolicy=Exact,failurePolicy=fail,groups=%s,resources=%s,"+
			"verbs=create;update,versions=%s,name=%s,admissionReviewVersions=v1",
		path,
		group,
		resource,
		version,
		name)

	funcs := []astmodel.Function{
		NewValidateFunction(
			"ValidateCreate",
			v.resourceName,
			v.idFactory,
			v.validateCreate,
			v.resourceName.PackageReference(),
			astmodel.FmtReference,
			astmodel.GenRuntimeReference),
		NewValidateFunction(
			"ValidateUpdate",
			v.resourceName,
			v.idFactory,
			v.validateUpdate,
			v.resourceName.PackageReference(),
			astmodel.FmtReference,
			astmodel.GenRuntimeReference),
		NewValidateFunction(
			"ValidateDelete",
			v.resourceName,
			v.idFactory,
			v.validateDelete,
			v.resourceName.PackageReference(),
			astmodel.FmtReference,
			astmodel.GenRuntimeReference),
		NewValidateFunction(
			"createValidations",
			v.resourceName,
			v.idFactory,
			v.localCreateValidations),
		NewValidateFunction(
			"updateValidations",
			v.resourceName,
			v.idFactory,
			v.localUpdateValidations),
		NewValidateFunction(
			"deleteValidations",
			v.resourceName,
			v.idFactory,
			v.localDeleteValidations),
	}

	// Add the actual individual validation functions
	for _, validations := range v.validations {
		for _, validation := range validations {
			funcs = append(funcs, validation)
		}
	}

	return astmodel.NewInterfaceImplementation(
		astmodel.ValidatorInterfaceName,
		funcs...,
	).WithAnnotation(annotation)
}

// validateCreate returns a function that performs validation of creation for the resource
func (v *ValidatorBuilder) validateCreate(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	resourceIdent := "resource"
	contextIdent := "ctx"

	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.validateBody(
		codeGenerationContext,
		k,
		receiverIdent,
		contextIdent,
		resourceIdent,
		"createValidations",
		"CreateValidations",
		"ValidateCreate",
		"")
	if err != nil {
		return nil, eris.Wrap(err, "creating validation body")
	}

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

	runtimeObjectTypeExpr, err := astmodel.APIMachineryObject.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(resourceIdent, runtimeObjectTypeExpr)

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))

	fn.AddComments("validates the creation of the resource")
	return fn.DefineFunc(), nil
}

// validateUpdate returns a function that performs validation of update for the resource
func (v *ValidatorBuilder) validateUpdate(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	resourceIdent := "newResource"
	oldResourceIdent := "oldResource"
	contextIdent := "ctx"

	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.validateBody(
		codeGenerationContext,
		k,
		receiverIdent,
		contextIdent,
		resourceIdent,
		"updateValidations",
		"UpdateValidations",
		"ValidateUpdate",
		oldResourceIdent)
	if err != nil {
		return nil, eris.Wrap(err, "creating validation body")
	}

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

	runtimeObjectTypeExpr, err := astmodel.APIMachineryObject.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(oldResourceIdent, runtimeObjectTypeExpr)
	fn.AddParameter(resourceIdent, runtimeObjectTypeExpr)

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))

	fn.AddComments("validates an update of the resource")
	return fn.DefineFunc(), nil
}

// validateDelete returns a function that performs validation of deletion for the resource
func (v *ValidatorBuilder) validateDelete(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	resourceIdent := "resource"
	contextIdent := "ctx"

	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.validateBody(
		codeGenerationContext,
		k,
		receiverIdent,
		contextIdent,
		resourceIdent,
		"deleteValidations",
		"DeleteValidations",
		"ValidateDelete",
		"")
	if err != nil {
		return nil, eris.Wrap(err, "creating validation body")
	}

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

	runtimeObjectTypeExpr, err := astmodel.APIMachineryObject.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object type expression")
	}
	fn.AddParameter(resourceIdent, runtimeObjectTypeExpr)

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates the deletion of the resource")
	return fn.DefineFunc(), nil
}

// validateBody returns the body for the generic validation function which invokes all local (code generated) validations
// as well as checking if there are any handcrafted validations and invoking them too:
// For example:
//
//	resource, ok := obj.(*<resourceName>)
//	if !ok {
//		return nil, fmt.Errorf("Expected <resourceName>, but got %T", obj)
//	}
//	validations := <receiverIdent>.createValidations()
//	var temp interface{} = <receiverIdent>
//	if runtimeValidator, ok := temp.(genruntime.Validator[T]); ok {
//		validations = append(validations, runtimeValidator.CreateValidations()...)
//	}
//	return genruntime.<validationFunctionName>(ctx, obj, validations)
func (v *ValidatorBuilder) validateBody(
	codeGenerationContext *astmodel.CodeGenerationContext,
	k *ValidateFunction,
	receiverIdent string,
	contextIdent string,
	objectIdent string,
	implFunctionName string,
	overrideFunctionName string,
	validationFunctionName string,
	oldObjectIdent string,
) ([]dst.Stmt, error) {
	var objIdent string
	var oldObjIdent string
	if oldObjectIdent == "" {
		objIdent = "obj"
	} else {
		objIdent = "newObj"
		oldObjIdent = "oldObj"
	}

	var castStmts []dst.Stmt

	fmtPkg, err := codeGenerationContext.GetImportedPackageName(astmodel.FmtReference)
	if err != nil {
		return nil, eris.Wrap(err, "getting fmt package name")
	}

	resourceTypeExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource type expression")
	}
	castStmts = append(castStmts, astbuilder.TypeAssert(dst.NewIdent(objIdent), dst.NewIdent(objectIdent), astbuilder.PointerTo(resourceTypeExpr)))
	castStmts = append(castStmts, astbuilder.IfNotOk(astbuilder.Returns(astbuilder.Nil(), astbuilder.FormatError(fmtPkg, fmt.Sprintf("expected %s, but got %%T", k.data), dst.NewIdent(objectIdent)))))
	if oldObjectIdent != "" {
		castStmts = append(castStmts, astbuilder.TypeAssert(dst.NewIdent(oldObjIdent), dst.NewIdent(oldObjectIdent), astbuilder.PointerTo(resourceTypeExpr)))
		castStmts = append(castStmts, astbuilder.IfNotOk(astbuilder.Returns(astbuilder.Nil(), astbuilder.FormatError(fmtPkg, fmt.Sprintf("expected %s, but got %%T", k.data), dst.NewIdent(oldObjectIdent)))))
	}

	overrideInterfaceType, err := astmodel.GenRuntimeValidatorInterfaceName.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", astmodel.GenRuntimeValidatorInterfaceName)
	}
	overrideInterfaceType = &dst.IndexExpr{
		X:     overrideInterfaceType,
		Index: astbuilder.PointerTo(resourceTypeExpr),
	}

	validationsIdent := "validations"
	tempVarIdent := "temp"
	runtimeValidatorIdent := "runtimeValidator"

	var args []dst.Expr
	args = append(args, dst.NewIdent(contextIdent))
	if oldObjectIdent != "" {
		args = append(args, dst.NewIdent(oldObjIdent))
	}
	args = append(args, dst.NewIdent(objIdent))
	args = append(args, dst.NewIdent(validationsIdent))

	hack := astbuilder.CallQualifiedFunc(runtimeValidatorIdent, overrideFunctionName)
	hack.Ellipsis = true

	appendFuncCall := astbuilder.CallFunc("append", dst.NewIdent(validationsIdent), astbuilder.CallQualifiedFunc(runtimeValidatorIdent, overrideFunctionName))
	appendFuncCall.Ellipsis = true

	body := astbuilder.Statements(
		castStmts,
		astbuilder.ShortDeclaration(
			validationsIdent,
			astbuilder.CallQualifiedFunc(receiverIdent, implFunctionName)),
		astbuilder.AssignToInterface(tempVarIdent, dst.NewIdent(receiverIdent)),
		astbuilder.IfType(
			dst.NewIdent(tempVarIdent),
			overrideInterfaceType,
			runtimeValidatorIdent,
			// Not using astbuilder.AppendList here as we want to tack on a "..." at the end
			astbuilder.SimpleAssignment(dst.NewIdent(validationsIdent), appendFuncCall)),
		astbuilder.Returns(astbuilder.CallQualifiedFunc(astmodel.GenRuntimeReference.PackageName(), validationFunctionName, args...)),
	)

	return body, nil
}

func (v *ValidatorBuilder) localCreateValidations(
	_ *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	fn, err := v.makeLocalValidationFuncDetails(ValidationKindCreate, codeGenerationContext, receiver, methodName)
	if err != nil {
		// error already has name of method, doesn't need wrapping
		return nil, err
	}

	fn.AddComments("validates the creation of the resource")
	return fn.DefineFunc(), nil
}

func (v *ValidatorBuilder) localUpdateValidations(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	fn, err := v.makeLocalValidationFuncDetails(ValidationKindUpdate, codeGenerationContext, receiver, methodName)
	if err != nil {
		// error already has name of method, doesn't need wrapping
		return nil, err
	}

	fn.AddComments("validates the update of the resource")
	return fn.DefineFunc(), nil
}

func (v *ValidatorBuilder) localDeleteValidations(
	k *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	fn, err := v.makeLocalValidationFuncDetails(ValidationKindDelete, codeGenerationContext, receiver, methodName)
	if err != nil {
		// error already has name of method, doesn't need wrapping
		return nil, err
	}

	fn.AddComments("validates the deletion of the resource")
	return fn.DefineFunc(), nil
}

func (v *ValidatorBuilder) makeLocalValidationFuncDetails(
	kind ValidationKind,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*astbuilder.FuncDetails, error) {
	receiverIdent := v.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.localValidationFuncBody(kind, codeGenerationContext, receiver)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create local validation function body for %s", methodName)
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Returns: []*dst.Field{
			{
				Type: &dst.ArrayType{
					Elt: getValidationFuncType(kind, v.resourceName, codeGenerationContext),
				},
			},
		},
		Body: body,
	}

	return fn, nil
}

// localValidationFuncBody returns the body of the local (code generated) validation functions:
//
//	return []func(ctx context.Context, resource *T) error {
//		<receiver>.<validationFunc1>,
//		<receiver>.<validationFunc2>,
//		...
//	}
//
// or in the case of update functions (that may not need the old parameter):
//
//	return []func(old *T, new *T) (admission.Warnings, error) {
//		func(old *T, new *T) (admission.Warnings, error) {
//			return <receiver>.<validationFunc1>
//		},
//		<receiver>.<validationFunc2>,
//	}
func (v *ValidatorBuilder) localValidationFuncBody(
	kind ValidationKind,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) ([]dst.Stmt, error) {
	elements := make([]dst.Expr, 0, len(v.validations[kind]))
	var errs []error
	for _, validationFunc := range v.validations[kind] {
		expr, err := v.makeLocalValidationElement(kind, validationFunc, codeGenerationContext, receiver)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		expr.Decorations().Before = dst.NewLine
		expr.Decorations().After = dst.NewLine

		elements = append(elements, expr)
	}

	if len(errs) > 0 {
		return nil, eris.Wrap(
			kerrors.NewAggregate(errs),
			"failed to create local validation function body")
	}

	if len(elements) == 0 {
		return astbuilder.Statements(
			astbuilder.Returns(astbuilder.Nil()),
		), nil
	}

	returnStmt := astbuilder.Returns(&dst.CompositeLit{
		Type: &dst.ArrayType{
			Elt: getValidationFuncType(kind, v.resourceName, codeGenerationContext),
		},
		Elts: elements,
	})

	return astbuilder.Statements(returnStmt), nil
}

// makeLocalValidationElement creates a validation expression, automatically removing the old parameter for update
// validations if it's not needed. These elements are used to build the list of validation functions.
// If validation != ValidationKindUpdate or validation == ValidationKindUpdate that DOES use the old parameter:
//
//	<receiver>.<validationFunc>
//
// If validate == ValidationKindUpdate that doesn't use the old parameter:
//
//	func(old *T) error {
//		return <receiver>.<validationFunc>
//	}
func (v *ValidatorBuilder) makeLocalValidationElement(
	kind ValidationKind,
	validation *ValidateFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) (dst.Expr, error) {
	receiverIdent := v.idFactory.CreateReceiver(receiver.Name())
	ctxIdent := "ctx"
	newObjIdent := "newObj"

	if kind == ValidationKindUpdate {
		// It's common that updates don't actually need the "old" variable. If the function that we're going to be calling
		// doesn't take any parameters, provide a wrapper
		f, err := validation.asFunc(validation, codeGenerationContext, receiver, validation.name)
		if err != nil {
			return nil, eris.Wrapf(err, "unable to create function for %s", validation.name)
		}

		if f.Type.Params.NumFields() == 2 {
			return &dst.FuncLit{
				Decs: dst.FuncLitDecorations{
					NodeDecs: dst.NodeDecs{
						// Start:  doc,
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
				Type: getValidationFuncType(kind, v.resourceName, codeGenerationContext),
				Body: astbuilder.StatementBlock(
					astbuilder.Returns(astbuilder.CallQualifiedFunc(receiverIdent, validation.name, dst.NewIdent(ctxIdent), dst.NewIdent(newObjIdent))),
				),
			}, nil
		}
	}

	return astbuilder.Selector(dst.NewIdent(receiverIdent), validation.name), nil
}

func getValidationFuncType(kind ValidationKind, resourceName astmodel.InternalTypeName, codeGenerationContext *astmodel.CodeGenerationContext) *dst.FuncType {
	typeExpr, err := resourceName.AsTypeExpr(codeGenerationContext)
	if err != nil {
		panic(err)
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		panic(err)
	}

	result := []*dst.Field{
		{
			Type: astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"),
		},
		{
			Type: dst.NewIdent("error"),
		},
	}

	switch kind {
	case ValidationKindUpdate:
		return &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{
							dst.NewIdent("ctx"),
						},
						Type: contextTypeExpr,
					},
					{
						Names: []*dst.Ident{
							dst.NewIdent("oldObj"),
						},
						Type: astbuilder.PointerTo(typeExpr),
					},
					{
						Names: []*dst.Ident{
							dst.NewIdent("newObj"),
						},
						Type: astbuilder.PointerTo(typeExpr),
					},
				},
			},
			Results: &dst.FieldList{
				List: result,
			},
		}
	default:
		return &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{
							dst.NewIdent("ctx"),
						},
						Type: contextTypeExpr,
					},
					{
						Names: []*dst.Ident{
							dst.NewIdent("obj"),
						},
						Type: astbuilder.PointerTo(typeExpr),
					},
				},
			},
			Results: &dst.FieldList{
				List: result,
			},
		}
	}
}
