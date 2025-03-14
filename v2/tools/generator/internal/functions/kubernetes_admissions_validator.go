/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/dave/dst"

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

// ValidatorBuilder helps in building an interface implementation for admissions.Validator.
type ValidatorBuilder struct {
	resourceName astmodel.InternalTypeName
	resource     *astmodel.ResourceType
	idFactory    astmodel.IdentifierFactory

	validations map[ValidationKind][]*ResourceFunction
}

// NewValidatorBuilder creates a new ValidatorBuilder for the given object type.
func NewValidatorBuilder(
	resourceName astmodel.InternalTypeName,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) *ValidatorBuilder {
	return &ValidatorBuilder{
		resourceName: resourceName,
		resource:     resource,
		idFactory:    idFactory,
		validations: map[ValidationKind][]*ResourceFunction{
			ValidationKindCreate: nil,
			ValidationKindUpdate: nil,
			ValidationKindDelete: nil,
		},
	}
}

// AddValidation adds a validation function to the set of validation functions to be applied to the given object.
func (v *ValidatorBuilder) AddValidation(kind ValidationKind, f *ResourceFunction) {
	if !v.resource.Equals(f.resource, astmodel.EqualityOverrides{}) {
		panic("cannot add validation function on non-matching object types")
	}
	v.validations[kind] = append(v.validations[kind], f)
}

// ToInterfaceImplementation creates an InterfaceImplementation that implements the admissions.Validator interface.
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
		NewResourceFunction(
			"ValidateCreate",
			v.resource,
			v.idFactory,
			v.validateCreate,
			astmodel.GenRuntimeReference,
			astmodel.APIMachineryErrorsReference,
			astmodel.APIMachineryRuntimeReference),
		NewResourceFunction(
			"ValidateUpdate",
			v.resource,
			v.idFactory,
			v.validateUpdate,
			astmodel.GenRuntimeReference,
			astmodel.APIMachineryErrorsReference,
			astmodel.APIMachineryRuntimeReference),
		NewResourceFunction(
			"ValidateDelete",
			v.resource,
			v.idFactory,
			v.validateDelete,
			astmodel.GenRuntimeReference,
			astmodel.APIMachineryErrorsReference,
			astmodel.APIMachineryRuntimeReference),
		NewResourceFunction(
			"createValidations",
			v.resource,
			v.idFactory,
			v.localCreateValidations),
		NewResourceFunction(
			"updateValidations",
			v.resource,
			v.idFactory,
			v.localUpdateValidations),
		NewResourceFunction(
			"deleteValidations",
			v.resource,
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
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.validateBody(
		codeGenerationContext,
		receiverIdent,
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

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates the creation of the resource")
	return fn.DefineFunc(), nil
}

// validateUpdate returns a function that performs validation of update for the resource
func (v *ValidatorBuilder) validateUpdate(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	retType := getValidationFuncType(ValidationKindUpdate, codeGenerationContext)

	body, err := v.validateBody(
		codeGenerationContext,
		receiverIdent,
		"updateValidations",
		"UpdateValidations",
		"ValidateUpdate",
		"old")
	if err != nil {
		return nil, eris.Wrap(err, "creating validation body")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		Params:        retType.Params.List,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Returns:       retType.Results.List,
		Body:          body,
	}

	fn.AddComments("validates an update of the resource")
	return fn.DefineFunc(), nil
}

// validateDelete returns a function that performs validation of deletion for the resource
func (v *ValidatorBuilder) validateDelete(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	body, err := v.validateBody(
		codeGenerationContext,
		receiverIdent,
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

	fn.AddReturn(astbuilder.QualifiedTypeName(codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeAdmission), "Warnings"))
	fn.AddReturn(dst.NewIdent("error"))
	fn.AddComments("validates the deletion of the resource")
	return fn.DefineFunc(), nil
}

// validateBody returns the body for the generic validation function which invokes all local (code generated) validations
// as well as checking if there are any handcrafted validations and invoking them too:
// For example:
//
//	validations := <receiverIdent>.createValidations()
//	var temp interface{} = <receiverIdent>
//	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
//		validations = append(validations, runtimeValidator.CreateValidations()...)
//	}
//	return genruntime.<validationFunctionName>(validations)
func (v *ValidatorBuilder) validateBody(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiverIdent string,
	implFunctionName string,
	overrideFunctionName string,
	validationFunctionName string,
	funcParamIdent string,
) ([]dst.Stmt, error) {
	overrideInterfaceType, err := astmodel.GenRuntimeValidatorInterfaceName.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", astmodel.GenRuntimeValidatorInterfaceName)
	}

	validationsIdent := "validations"
	tempVarIdent := "temp"
	runtimeValidatorIdent := "runtimeValidator"

	var args []dst.Expr
	if funcParamIdent != "" {
		args = append(args, dst.NewIdent(funcParamIdent))
	}
	args = append(args, dst.NewIdent(validationsIdent))

	hack := astbuilder.CallQualifiedFunc(runtimeValidatorIdent, overrideFunctionName)
	hack.Ellipsis = true

	appendFuncCall := astbuilder.CallFunc("append", dst.NewIdent(validationsIdent), astbuilder.CallQualifiedFunc(runtimeValidatorIdent, overrideFunctionName))
	appendFuncCall.Ellipsis = true

	body := astbuilder.Statements(
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
	_ *ResourceFunction,
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
	_ *ResourceFunction,
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
	_ *ResourceFunction,
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

	return &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Returns: []*dst.Field{
			{
				Type: &dst.ArrayType{
					Elt: getValidationFuncType(kind, codeGenerationContext),
				},
			},
		},
		Body: body,
	}, nil
}

// localValidationFuncBody returns the body of the local (code generated) validation functions:
//
//	return []func() error {
//		<receiver>.<validationFunc1>,
//		<receiver>.<validationFunc2>,
//		...
//	}
//
// or in the case of update functions (that may not need the old parameter):
//
//	return []func(old runtime.Object) (admission.Warnings, error) {
//		func(old runtime.Object) (admission.Warnings, error) {
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
			Elt: getValidationFuncType(kind, codeGenerationContext),
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
//	func(old runtime.Object) error {
//		return <receiver>.<validationFunc>
//	}
func (v *ValidatorBuilder) makeLocalValidationElement(
	kind ValidationKind,
	validation *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) (dst.Expr, error) {
	receiverIdent := v.idFactory.CreateReceiver(receiver.Name())

	if kind == ValidationKindUpdate {
		// It's common that updates don't actually need the "old" variable. If the function that we're going to be calling
		// doesn't take any parameters, provide a wrapper
		f, err := validation.asFunc(validation, codeGenerationContext, receiver, validation.name)
		if err != nil {
			return nil, eris.Wrapf(err, "unable to create function for %s", validation.name)
		}

		if f.Type.Params.NumFields() == 0 {
			return &dst.FuncLit{
				Decs: dst.FuncLitDecorations{
					NodeDecs: dst.NodeDecs{
						// Start:  doc,
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
				Type: getValidationFuncType(kind, codeGenerationContext),
				Body: &dst.BlockStmt{
					List: []dst.Stmt{
						astbuilder.Returns(astbuilder.CallQualifiedFunc(receiverIdent, validation.name)),
					},
				},
			}, nil
		}
	}

	return astbuilder.Selector(dst.NewIdent(receiverIdent), validation.name), nil
}

func getValidationFuncType(kind ValidationKind, codeGenerationContext *astmodel.CodeGenerationContext) *dst.FuncType {
	runtime, err := codeGenerationContext.GetImportedPackageName(astmodel.APIMachineryRuntimeReference)
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

	if kind == ValidationKindUpdate {
		return &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{
							dst.NewIdent("old"),
						},
						Type: &dst.SelectorExpr{
							X:   dst.NewIdent(runtime),
							Sel: dst.NewIdent("Object"),
						},
					},
				},
			},
			Results: &dst.FieldList{
				List: result,
			},
		}
	}

	return &dst.FuncType{
		Results: &dst.FieldList{
			List: result,
		},
	}
}
