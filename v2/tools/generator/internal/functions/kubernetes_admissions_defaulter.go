/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var defaultFunctionRequiredPackages = []astmodel.PackageReference{
	astmodel.ContextReference,
	astmodel.ControllerRuntimeWebhook,     // Used for interface assertion
	astmodel.APIMachineryRuntimeReference, // Used for runtime.Object parameter of Defaulter interface
}

type DefaultFunction = DataFunction[astmodel.InternalTypeName]

func NewDefaultFunction(
	name string,
	data astmodel.InternalTypeName,
	idFactory astmodel.IdentifierFactory,
	asFunc DataFunctionHandler[astmodel.InternalTypeName],
	requiredPackages ...astmodel.PackageReference,
) *DefaultFunction {
	// Add the default set of required packages
	requiredPackages = append(requiredPackages, defaultFunctionRequiredPackages...)

	return NewDataFunction[astmodel.InternalTypeName](
		name,
		data,
		idFactory,
		asFunc,
		requiredPackages...)
}

// DefaulterBuilder helps in building an interface implementation for webhook.CustomDefaulter.
type DefaulterBuilder struct {
	resourceName astmodel.InternalTypeName
	idFactory    astmodel.IdentifierFactory

	defaults []*DefaultFunction
}

// NewDefaulterBuilder creates a new DefaulterBuilder for the given object type.
func NewDefaulterBuilder(
	resourceName astmodel.InternalTypeName,
	idFactory astmodel.IdentifierFactory,
) *DefaulterBuilder {
	return &DefaulterBuilder{
		resourceName: resourceName,
		idFactory:    idFactory,
	}
}

// AddDefault adds a default function to the set of default functions to be applied to the given object
func (d *DefaulterBuilder) AddDefault(f *DefaultFunction) {
	if !f.Data().Equals(d.resourceName, astmodel.EqualityOverrides{}) {
		panic(fmt.Sprintf("cannot add validation function on non-matching resources. Expected %s, got %s", d.resourceName, f.Data()))
	}
	d.defaults = append(d.defaults, f)
}

// ToInterfaceImplementation creates an InterfaceImplementation that implements the webhook.CustomDefaulter interface.
// This implementation includes calls to all defaults registered with this DefaulterBuilder via the AddDefault function,
// as well as helper functions that allow additional handcrafted defaults to be injected by
// implementing the genruntime.Defaulter interface.
func (d *DefaulterBuilder) ToInterfaceImplementation() *astmodel.InterfaceImplementation {
	grp, ver := d.resourceName.InternalPackageReference().GroupVersion()

	// e.g. grp = "microsoft.network.azure.com"
	// e.g. resource = "backendaddresspools"
	// e.g. ver = "v1"

	resource := d.resourceName.Name()

	grp = strings.ToLower(grp + astmodel.GroupSuffix)
	nonPluralResource := strings.ToLower(resource)
	resource = strings.ToLower(d.resourceName.Plural().Name())

	// e.g. "mutate-microsoft-network-azure-com-v1-backendaddresspool"
	// note that this must match _exactly_ how controller-runtime generates the path,
	// or it will not work!
	path := fmt.Sprintf("/mutate-%s-%s-%s", strings.ReplaceAll(grp, ".", "-"), ver, nonPluralResource)

	// e.g.  "default.v123.backendaddresspool.azure.com"
	name := fmt.Sprintf("default.%s.%s.%s", ver, resource, grp)

	annotation := fmt.Sprintf(
		"+kubebuilder:webhook:path=%s,mutating=true,sideEffects=None,"+
			"matchPolicy=Exact,failurePolicy=fail,groups=%s,resources=%s,"+
			"verbs=create;update,versions=%s,name=%s,admissionReviewVersions=v1",
		path,
		grp,
		resource,
		ver,
		name)

	funcs := []astmodel.Function{
		NewDefaultFunction(
			"Default",
			d.resourceName,
			d.idFactory,
			d.defaultFunction,
			d.resourceName.PackageReference(),
			astmodel.FmtReference,
			astmodel.GenRuntimeReference),
		NewDefaultFunction(
			"defaultImpl",
			d.resourceName,
			d.idFactory,
			d.localDefault,
			astmodel.FmtReference,
			astmodel.GenRuntimeReference),
	}

	// Add the actual individual default functions
	for _, def := range d.defaults {
		funcs = append(funcs, def)
	}

	return astmodel.NewInterfaceImplementation(
		astmodel.DefaulterInterfaceName,
		funcs...).WithAnnotation(annotation)
}

func (d *DefaulterBuilder) localDefault(
	k *DefaultFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	objIdent := "obj"
	contextIdent := "ctx"

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	defaults := make([]dst.Stmt, 0, len(d.defaults))
	tok := token.DEFINE
	for _, def := range d.defaults {
		call := astbuilder.AssignmentStatement(
			dst.NewIdent("err"),
			tok,
			astbuilder.CallQualifiedFunc(receiverIdent, def.Name(), dst.NewIdent(contextIdent), dst.NewIdent(objIdent)))
		check := astbuilder.CheckErrorAndReturn()
		tok = token.ASSIGN
		defaults = append(
			defaults,
			call,
			check)
	}

	defaults = append(defaults, astbuilder.Returns(astbuilder.Nil()))

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body:          defaults,
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter(contextIdent, contextTypeExpr)

	resourceTypeExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource type expression")
	}
	fn.AddParameter(objIdent, astbuilder.PointerTo(resourceTypeExpr))
	fn.AddReturns("error")

	fn.AddComments(fmt.Sprintf("applies the code generated defaults to the %s resource", receiver.Name()))
	return fn.DefineFunc(), nil
}

// defaultFunction returns a function with the following body:
//
//	func (r *<reciever>) Default(ctx context.Context, obj runtime.Object) error {
//		resource, ok := obj.(*<resourceName>)
//		if !ok {
//			return fmt.Errorf("Expected <resourceName>, but got %T", obj)
//		}
//
//		err := r.defaultImpl(ctx, resource)
//		if err != nil {
//			return err
//		}
//		var temp any = disk
//		if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
//			err = runtimeDefaulter.CustomDefault(ctx, resource)
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	}
func (d *DefaulterBuilder) defaultFunction(
	k *DefaultFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	resourceIdent := "resource"
	objectIdent := "obj"
	contextIdent := "ctx"

	fmtPkg, err := codeGenerationContext.GetImportedPackageName(astmodel.FmtReference)
	if err != nil {
		return nil, eris.Wrap(err, "getting fmt package name")
	}

	resourceTypeExpr, err := k.data.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating resource type expression")
	}

	assertStmt := astbuilder.TypeAssert(dst.NewIdent(resourceIdent), dst.NewIdent(objectIdent), astbuilder.PointerTo(resourceTypeExpr))
	ifAssertFails := astbuilder.IfNotOk(astbuilder.Returns(astbuilder.FormatError(fmtPkg, fmt.Sprintf("expected %s, but got %%T", k.data), dst.NewIdent(objectIdent))))

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverType, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	tempVarIdent := "temp"
	runtimeDefaulterIdent := "runtimeDefaulter"

	overrideInterfaceType, err := astmodel.GenRuntimeDefaulterInterfaceName.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating runtime defaulter interface type expression for method %s", methodName)
	}

	defaultImplCall := astbuilder.AssignmentStatement(
		dst.NewIdent("err"),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(receiverIdent, "defaultImpl", dst.NewIdent(contextIdent), dst.NewIdent(resourceIdent)))

	customDefaultCall := astbuilder.AssignmentStatement(
		dst.NewIdent("err"),
		token.ASSIGN,
		astbuilder.CallQualifiedFunc(runtimeDefaulterIdent, "CustomDefault", dst.NewIdent(contextIdent), dst.NewIdent(resourceIdent)))

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body: astbuilder.Statements(
			assertStmt,
			ifAssertFails,
			// TODO: This part should maybe be conditional if there are no defaults to define?
			defaultImplCall,
			astbuilder.CheckErrorAndReturn(),
			astbuilder.AssignToInterface(tempVarIdent, dst.NewIdent(receiverIdent)),
			astbuilder.IfType(
				dst.NewIdent(tempVarIdent),
				overrideInterfaceType,
				runtimeDefaulterIdent,
				customDefaultCall,
				astbuilder.CheckErrorAndReturn(),
			),
			astbuilder.Returns(astbuilder.Nil()),
		),
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
	fn.AddParameter(objectIdent, runtimeObjectTypeExpr)
	fn.AddReturns("error")

	fn.AddComments(fmt.Sprintf("applies defaults to the %s resource", receiver.Name()))
	return fn.DefineFunc(), nil
}
