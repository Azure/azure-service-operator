/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// AddKubernetesResourceInterfaceImpls adds the required interfaces for
// the resource to be a Kubernetes resource.
// Returns a set of modified definitions.
func AddKubernetesResourceInterfaceImpls(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	definitions astmodel.TypeDefinitionSet,
	log logr.Logger,
) (astmodel.TypeDefinitionSet, error) {
	resolved, err := definitions.ResolveResourceSpecAndStatus(resourceDef)
	if err != nil {
		return nil, eris.Wrapf(err, "unable to resolve resource %s", resourceDef.Name())
	}

	specDef := resolved.SpecDef
	r := resolved.ResourceType

	// Check the spec first to ensure it looks how we expect
	if r.Scope() == astmodel.ResourceScopeResourceGroup || r.Scope() == astmodel.ResourceScopeExtension {
		ownerProperty := idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported)
		_, ok := resolved.SpecType.Property(ownerProperty)
		if !ok {
			return nil, eris.Errorf("resource spec doesn't have %q property", ownerProperty)
		}
	}

	azureNameProp, ok := resolved.SpecType.Property(astmodel.AzureNameProperty)
	if !ok {
		return nil, eris.Errorf("resource spec doesn't have %q property", astmodel.AzureNameProperty)
	}

	nameFns, err := createAzureNameFunctionHandlersForType(azureNameProp.PropertyType(), definitions, log)
	if err != nil {
		return nil, err
	}

	// Sometimes we need to remove the AzureName property from our Spec because the name is forced
	if nameFns.removeAzureNameProperty {
		// remove the AzureName property from the spec of the resource
		remover := astmodel.NewPropertyRemover()
		var updated astmodel.TypeDefinition
		updated, err = remover.Remove(resolved.SpecDef, astmodel.AzureNameProperty)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to remove AzureName property from resource %s", resourceDef.Name())
		}

		specDef = updated
	}

	getAzureNameProperty := functions.NewObjectFunction(
		astmodel.AzureNameProperty,
		idFactory,
		nameFns.getNameFunction,
		astmodel.GenRuntimeReference)

	getOwnerProperty := functions.NewResourceFunction(
		astmodel.OwnerProperty,
		r,
		idFactory,
		getOwnerFunction,
		astmodel.GenRuntimeReference)

	getSpecFunction := functions.NewGetSpecFunction(idFactory)

	getTypeFunction := functions.NewGetTypeFunction(r.ARMType(), idFactory, functions.ReceiverTypePtr)

	getResourceScopeFunc := functions.NewResourceFunction(
		"GetResourceScope",
		r,
		idFactory,
		getResourceScopeFunction,
		astmodel.GenRuntimeReference)

	getSupportedOperationsFunc := functions.NewResourceFunction(
		"GetSupportedOperations",
		r,
		idFactory,
		getSupportedOperationsFunction,
		astmodel.GenRuntimeReference)

	getAPIVersionFunc := functions.NewGetAPIVersionFunction(r.APIVersionEnumValue(), idFactory)

	fns := []astmodel.Function{
		getAzureNameProperty,
		getOwnerProperty,
		getSpecFunction,
		getTypeFunction,
		getResourceScopeFunc,
		getSupportedOperationsFunc,
		getAPIVersionFunc,
	}

	if r.StatusType() != nil {
		// Skip Status functions if no status
		status, ok := astmodel.AsTypeName(r.StatusType())
		if !ok {
			msg := fmt.Sprintf(
				"Unable to create NewEmptyStatus() for resource %s (expected Status to be a TypeName but had %T)",
				resourceDef.Name(),
				r.StatusType())
			return nil, eris.New(msg)
		}

		emptyStatusFunction := functions.NewEmptyStatusFunction(status, idFactory)
		getStatusFunction := functions.NewGetStatusFunction(idFactory)
		setStatusFunction := functions.NewResourceStatusSetterFunction(r, idFactory)

		fns = append(fns, emptyStatusFunction, getStatusFunction, setStatusFunction)
	}

	kubernetesResourceImplementation := astmodel.NewInterfaceImplementation(astmodel.KubernetesResourceType, fns...)

	interfaceInjector := astmodel.NewInterfaceInjector()
	updatedResource, err := interfaceInjector.Inject(resolved.ResourceDef, kubernetesResourceImplementation)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to inject KubernetesResource interface into resource %s", resourceDef.Name())
	}

	if nameFns.setNameFunction != nil {
		// this function applies to Spec not the resource
		functionInjector := astmodel.NewFunctionInjector()

		setFn := functions.NewObjectFunction(
			astmodel.SetAzureNameFunc,
			idFactory,
			nameFns.setNameFunction,
			astmodel.GenRuntimeReference)

		updated, err := functionInjector.Inject(specDef, setFn)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to inject SetAzureName function into resource %s", resourceDef.Name())
		}

		specDef = updated
	}

	result := astmodel.MakeTypeDefinitionSetFromDefinitions(
		updatedResource,
		specDef,
	)

	return result, nil
}

func createAzureNameFunctionHandlersForType(
	t astmodel.Type,
	definitions astmodel.TypeDefinitionSet,
	log logr.Logger,
) (createAzureNameFunctionsForTypeResult, error) {
	if opt, ok := astmodel.AsOptionalType(t); ok {
		t = opt.BaseType()
	}

	// handle different definitions of AzureName property
	switch azureNamePropType := t.(type) {
	case *astmodel.ValidatedType:
		if !astmodel.TypeEquals(azureNamePropType.ElementType(), astmodel.StringType) {
			return createAzureNameFunctionsForTypeResult{},
				eris.Errorf("unable to handle non-string validated definitions in AzureName property")
		}

		validations := azureNamePropType.Validations().(astmodel.StringValidations)
		if len(validations.Patterns) != 0 {
			if len(validations.Patterns) == 1 &&
				validations.Patterns[0].String() == "^.*/default$" {
				// Validation requires the resource be named exactly "default"
				return createAzureNameFunctionsForTypeResult{
					getNameFunction:         fixedValueGetAzureNameFunction("default"),
					removeAzureNameProperty: true,
				}, nil
			}

			// ignoring for now:
			log.V(1).Info("ignoring pattern validation on Name property", "pattern", validations.Patterns[0].String())
			return createAzureNameFunctionsForTypeResult{
				getNameFunction: getStringAzureNameFunction,
				setNameFunction: setStringAzureNameFunction,
			}, nil
		}

		// ignoring length validations for now
		// return nil, errors.Errorf("unable to handle validations on Name property …TODO")
		return createAzureNameFunctionsForTypeResult{
			getNameFunction: getStringAzureNameFunction,
			setNameFunction: setStringAzureNameFunction,
		}, nil

	case astmodel.TypeName:
		// resolve property type if it is a typename
		resolvedPropType, err := definitions.FullyResolve(azureNamePropType)
		if err != nil {
			return createAzureNameFunctionsForTypeResult{},
				eris.Wrapf(err, "unable to resolve type of resource Name property: %s", azureNamePropType.String())
		}

		if t, ok := resolvedPropType.(*astmodel.EnumType); ok {
			if !astmodel.TypeEquals(t.BaseType(), astmodel.StringType) {
				return createAzureNameFunctionsForTypeResult{},
					eris.Errorf("unable to handle non-string enum base type in Name property")
			}

			options := t.Options()
			if len(options) == 1 {
				// if there is only one possible value,
				// we make an AzureName function that returns it, and do not
				// provide an AzureName property on the spec
				return createAzureNameFunctionsForTypeResult{
					getNameFunction:         fixedValueGetAzureNameFunction(options[0].Value),
					removeAzureNameProperty: true,
				}, nil
			}

			// with multiple values, provide an AzureName function that casts from the
			// enum-valued AzureName property:
			return createAzureNameFunctionsForTypeResult{
				getNameFunction: getEnumAzureNameFunction(azureNamePropType),
				setNameFunction: setEnumAzureNameFunction(azureNamePropType),
			}, nil
		}

		return createAzureNameFunctionsForTypeResult{},
			eris.Errorf("unable to produce AzureName()/SetAzureName() for Name property with type %s", resolvedPropType.String())

	case *astmodel.PrimitiveType:
		if !astmodel.TypeEquals(azureNamePropType, astmodel.StringType) {
			return createAzureNameFunctionsForTypeResult{},
				eris.Errorf("cannot use type %s as type of AzureName property", azureNamePropType.String())
		}

		return createAzureNameFunctionsForTypeResult{
			getNameFunction: getStringAzureNameFunction,
			setNameFunction: setStringAzureNameFunction,
		}, nil

	default:
		return createAzureNameFunctionsForTypeResult{},
			eris.Errorf("unsupported type for AzureName property: %s", azureNamePropType.String())
	}
}

type createAzureNameFunctionsForTypeResult struct {
	getNameFunction         functions.ObjectFunctionHandler
	setNameFunction         functions.ObjectFunctionHandler
	removeAzureNameProperty bool
}

// getEnumAzureNameFunction adds an AzureName() function that casts the AzureName property
// with an enum value to a string
func getEnumAzureNameFunction(enumType astmodel.TypeName) functions.ObjectFunctionHandler {
	return func(
		f *functions.ObjectFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) (*dst.FuncDecl, error) {
		receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
		receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating receiver type expression")
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  astbuilder.PointerTo(receiverExpr),
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.CallFunc("string", astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec", astmodel.AzureNameProperty)))),
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (string representation of %s)", enumType.String()))
		fn.AddReturns("string")
		return fn.DefineFunc(), nil
	}
}

// setEnumAzureNameFunction returns a function that sets the AzureName property to the result of casting
// the argument string to the given enum type
func setEnumAzureNameFunction(enumType astmodel.TypeName) functions.ObjectFunctionHandler {
	return func(
		f *functions.ObjectFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) (*dst.FuncDecl, error) {
		receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
		receiverTypeExpr, err := receiver.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating receiver type expression")
		}

		azureNameProp := astbuilder.Selector(dst.NewIdent(receiverIdent), astmodel.AzureNameProperty)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  astbuilder.PointerTo(receiverTypeExpr),
			Body: astbuilder.Statements(
				astbuilder.SimpleAssignment(
					azureNameProp,
					astbuilder.CallFunc(enumType.Name(), dst.NewIdent("azureName")))),
		}

		fn.AddComments(fmt.Sprintf("sets the Azure name from the given %s value", enumType.String()))
		fn.AddParameter("azureName", dst.NewIdent("string"))
		return fn.DefineFunc(), nil
	}
}

// fixedValueGetAzureNameFunction adds an AzureName() function that returns a fixed value
func fixedValueGetAzureNameFunction(fixedValue string) functions.ObjectFunctionHandler {
	// ensure fixedValue is quoted. This is always the case with enum values we pass,
	// but let's be safe:
	if len(fixedValue) == 0 {
		panic("cannot created fixed value AzureName function with empty fixed value")
	}

	if fixedValue[0] != '"' || fixedValue[len(fixedValue)-1] != '"' {
		fixedValue = fmt.Sprintf("%q", fixedValue)
	}

	return func(
		f *functions.ObjectFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) (*dst.FuncDecl, error) {
		receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
		receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating receiver type expression")
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  astbuilder.PointerTo(receiverExpr),
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.TextLiteral(fixedValue))),
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (always %s)", fixedValue))
		fn.AddReturns("string")
		return fn.DefineFunc(), nil
	}
}

// getOwnerFunction creates the Owner function declaration. This has two possible formats.
// For normal resources:
//
//	func (<receiver> *<receiver>) Owner() *genruntime.ResourceReference {
//		group, kind := genruntime.LookupOwnerGroupKind(<receiver>.Spec)
//		return <receiver>.Spec.Owner.AsKnownResourceReference(group, kind)
//	}
//
// For extension resources:
//
//	func (<receiver> *<receiver>) Owner() *genruntime.ResourceReference {
//		return <receiver>.Spec.Owner.AsKnownResourceReference()
//	}
func getOwnerFunction(
	r *functions.ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := r.IDFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	specSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")
	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Params:        nil,
	}
	resourceReferenceTypeExpr, err := astmodel.ResourceReferenceType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource reference type expression")
	}

	fn.AddReturn(astbuilder.Dereference(resourceReferenceTypeExpr))

	groupLocal := "group"
	kindLocal := "kind"
	if receiverIdent == groupLocal || receiverIdent == kindLocal {
		groupLocal = "ownerGroup"
		kindLocal = "ownerKind"
	}

	owner := astbuilder.Selector(specSelector, astmodel.OwnerProperty)
	nilOwnerCheck := astbuilder.IfNil(owner, astbuilder.Returns(astbuilder.Nil()))
	nilOwnerCheck.Decs.After = dst.EmptyLine

	switch r.Resource().Scope() {
	case astmodel.ResourceScopeResourceGroup:
		fn.AddComments("returns the ResourceReference of the owner")
		fn.AddStatements(
			nilOwnerCheck,
			lookupGroupAndKindStmt(groupLocal, kindLocal, specSelector),
			astbuilder.Returns(createResourceReferenceWithGroupKind(dst.NewIdent(groupLocal), dst.NewIdent(kindLocal), owner)))

	case astmodel.ResourceScopeExtension:
		fn.AddComments("returns the ResourceReference of the owner")
		fn.AddStatements(
			nilOwnerCheck,
			astbuilder.Returns(createResourceReference(owner)))

	case astmodel.ResourceScopeTenant:
		// Tenant resources never have an owner, just return nil
		fn.AddComments("returns nil as Tenant scoped resources never have an owner")
		fn.AddStatements(astbuilder.Returns(astbuilder.Nil()))

	case astmodel.ResourceScopeLocation:
		// Location resources never have an owner, just return nil
		fn.AddComments("returns nil as Location scoped resources never have an owner")
		fn.AddStatements(astbuilder.Returns(astbuilder.Nil()))

	default:
		panic(fmt.Sprintf("unknown resource kind: %s", r.Resource().Scope()))
	}

	return fn.DefineFunc(), nil
}

// getResourceScopeFunction creates a function that returns the scope of the resource.
//
//	func (<receiver> *<receiver>) GetResourceScope() genruntime.ResourceScope {
//		return genruntime.ResourceScopeResourceGroup
//	}
func getResourceScopeFunction(
	r *functions.ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := r.IDFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	var resourceScope string
	switch r.Resource().Scope() {
	case astmodel.ResourceScopeLocation:
		resourceScope = "ResourceScopeLocation"
	case astmodel.ResourceScopeResourceGroup:
		resourceScope = "ResourceScopeResourceGroup"
	case astmodel.ResourceScopeExtension:
		resourceScope = "ResourceScopeExtension"
	case astmodel.ResourceScopeTenant:
		resourceScope = "ResourceScopeTenant"
	default:
		panic(fmt.Sprintf("unknown resource scope %s", r.Resource().Scope()))
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Params:        nil,
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.Selector(dst.NewIdent(astmodel.GenRuntimeReference.PackageName()), resourceScope))),
	}

	fn.AddComments("returns the scope of the resource")
	resourceScopeTypeExpr, err := astmodel.ResourceScopeType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource scope type expression")
	}

	fn.AddReturn(resourceScopeTypeExpr)

	return fn.DefineFunc(), nil
}

// getSupportedOperationsFunction creates a function that returns the supported operations of the resource.
//
//	func (<receiver> *<receiver>) GetSupportedOperations() []genruntime.ResourceOperation {
//		return []genruntime.ResourceOperation{
//			genruntime.ResourceOperationGet,
//			genruntime.ResourceOperationPut,
//			genruntime.ResourceOperationDelete,
//		}
//	}
func getSupportedOperationsFunction(
	r *functions.ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := r.IDFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	genruntimePackage := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)
	supportedOperations := set.AsSortedSlice(r.Resource().SupportedOperations()) // Sorted to ensure ordered codegen
	idents := make([]dst.Expr, 0, len(supportedOperations))
	for _, op := range supportedOperations {
		var genruntimeOpName string
		switch op {
		case astmodel.ResourceOperationPut:
			genruntimeOpName = "ResourceOperationPut"
		case astmodel.ResourceOperationGet:
			genruntimeOpName = "ResourceOperationGet"
		case astmodel.ResourceOperationHead:
			genruntimeOpName = "ResourceOperationHead"
		case astmodel.ResourceOperationDelete:
			genruntimeOpName = "ResourceOperationDelete"
		default:
			panic(fmt.Sprintf("unknown resource operation %s", op))
		}
		idents = append(idents, astbuilder.Selector(dst.NewIdent(genruntimePackage), genruntimeOpName))
	}

	resourceOperationTypeExpr, err := astmodel.ResourceOperationType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource operation type expression")
	}

	sliceBuilder := astbuilder.NewSliceLiteralBuilder(resourceOperationTypeExpr, true)
	for _, id := range idents {
		sliceBuilder.AddElement(id)
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Params:        nil,
		Body: astbuilder.Statements(
			astbuilder.Returns(sliceBuilder.Build()),
		),
	}

	fn.AddComments("returns the operations supported by the resource")
	resourceOperationTypeArrayExpr, err := astmodel.ResourceOperationTypeArray.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource operation type array expression")
	}

	fn.AddReturn(resourceOperationTypeArrayExpr)

	return fn.DefineFunc(), nil
}

func lookupGroupAndKindStmt(
	groupIdent string,
	kindIdent string,
	specSelector *dst.SelectorExpr,
) *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.NewIdent(groupIdent),
			dst.NewIdent(kindIdent),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			astbuilder.CallExpr(
				dst.NewIdent(astmodel.GenRuntimeReference.PackageName()),
				"LookupOwnerGroupKind",
				specSelector),
		},
	}
}

func createResourceReferenceWithGroupKind(
	group dst.Expr,
	kind dst.Expr,
	ownerSelector dst.Expr,
) dst.Expr {
	return astbuilder.CallExpr(ownerSelector, "AsResourceReference", group, kind)
}

func createResourceReference(
	ownerSelector dst.Expr,
) dst.Expr {
	return astbuilder.CallExpr(ownerSelector, "AsResourceReference")
}

// setStringAzureNameFunction returns a function that sets the Name property of
// the resource spec to the argument string
func setStringAzureNameFunction(
	k *functions.ObjectFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverTypeExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverTypeExpr),
		Body: astbuilder.Statements(
			astbuilder.QualifiedAssignment(
				dst.NewIdent(receiverIdent),
				astmodel.AzureNameProperty,
				token.ASSIGN,
				dst.NewIdent("azureName")),
		),
	}

	fn.AddComments("sets the Azure name of the resource")
	fn.AddParameter("azureName", dst.NewIdent("string"))
	return fn.DefineFunc(), nil
}

// getStringAzureNameFunction returns a function that returns the Name property of the resource spec
func getStringAzureNameFunction(
	k *functions.ObjectFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverTypeExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverTypeExpr),
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.Selector(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec"), astmodel.AzureNameProperty))),
	}

	fn.AddComments("returns the Azure name of the resource")
	fn.AddReturns("string")
	return fn.DefineFunc(), nil
}
