/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

var (
	DefaulterInterfaceName           = MakeTypeName(ControllerRuntimeAdmission, "Defaulter")
	GenRuntimeDefaulterInterfaceName = MakeTypeName(GenRuntimeReference, "Defaulter")
)

// DefaulterBuilder helps in building an interface implementation for admissions.Defaulter.
type DefaulterBuilder struct {
	resourceName TypeName
	resource     *ResourceType
	idFactory    IdentifierFactory

	defaults []*resourceFunction
}

// NewDefaulterBuilder creates a new DefaulterBuilder for the given object type.
func NewDefaulterBuilder(resourceName TypeName, resource *ResourceType, idFactory IdentifierFactory) *DefaulterBuilder {
	return &DefaulterBuilder{
		resourceName: resourceName,
		resource:     resource,
		idFactory:    idFactory,
	}
}

// AddDefault adds an additional default function to the set of default functions to be applied to the given object
func (d *DefaulterBuilder) AddDefault(f *resourceFunction) {
	if !d.resource.Equals(f.resource, EqualityOverrides{}) {
		panic("cannot add default function on non-matching object types")
	}
	d.defaults = append(d.defaults, f)
}

// ToInterfaceImplementation creates an InterfaceImplementation that implements the admissions.Defaulter interface.
// This implementation includes calls to all defaults registered with this DefaulterBuilder via the AddDefault function,
// as well as helper functions that allow additional handcrafted defaults to be injected by
// implementing the genruntime.Defaulter interface.
func (d *DefaulterBuilder) ToInterfaceImplementation() *InterfaceImplementation {
	lpr, ok := d.resourceName.PackageReference.AsLocalPackage()
	if !ok {
		panic(fmt.Sprintf("expected resource name %s to be a local package reference", d.resourceName.String()))
	}

	group := lpr.group                // e.g. "microsoft.network.azure.com"
	resource := d.resourceName.Name() // e.g. "backendaddresspools"
	version := lpr.version            // e.g. "v1"

	group = strings.ToLower(group + GroupSuffix)
	nonPluralResource := strings.ToLower(resource)
	resource = strings.ToLower(d.resourceName.Plural().Name())

	// e.g. "mutate-microsoft-network-azure-com-v1-backendaddresspool"
	// note that this must match _exactly_ how controller-runtime generates the path
	// or it will not work!
	path := fmt.Sprintf("/mutate-%s-%s-%s", strings.ReplaceAll(group, ".", "-"), version, nonPluralResource)

	// e.g.  "default.v123.backendaddresspool.azure.com"
	name := fmt.Sprintf("default.%s.%s.%s", version, resource, group)

	annotation := fmt.Sprintf(
		"+kubebuilder:webhook:path=%s,mutating=true,sideEffects=None,"+
			"matchPolicy=Exact,failurePolicy=fail,groups=%s,resources=%s,"+
			"verbs=create;update,versions=%s,name=%s,admissionReviewVersions=v1",
		path,
		group,
		resource,
		version,
		name)

	funcs := []Function{
		&resourceFunction{
			name:             "Default",
			resource:         d.resource,
			idFactory:        d.idFactory,
			asFunc:           d.defaultFunction,
			requiredPackages: NewPackageReferenceSet(GenRuntimeReference),
		},
		&resourceFunction{
			name:             "defaultImpl",
			resource:         d.resource,
			idFactory:        d.idFactory,
			asFunc:           d.localDefault,
			requiredPackages: NewPackageReferenceSet(GenRuntimeReference),
		},
	}

	// Add the actual individual default functions
	for _, def := range d.defaults {
		funcs = append(funcs, def)
	}

	return NewInterfaceImplementation(
		DefaulterInterfaceName,
		funcs...).WithAnnotation(annotation)
}

func (d *DefaulterBuilder) localDefault(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	var defaults []dst.Stmt
	for _, def := range d.defaults {
		defaults = append(
			defaults,
			&dst.ExprStmt{
				X: astbuilder.CallQualifiedFunc(receiverIdent, def.name),
			})
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: defaults,
	}

	fn.AddComments(fmt.Sprintf("applies the code generated defaults to the %s resource", receiver.Name()))
	return fn.DefineFunc()
}

func (d *DefaulterBuilder) defaultFunction(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)
	tempVarIdent := "temp"
	runtimeDefaulterIdent := "runtimeDefaulter"

	overrideInterfaceType := GenRuntimeDefaulterInterfaceName.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			astbuilder.InvokeQualifiedFunc(receiverIdent, "defaultImpl"), // TODO: This part should maybe be conditional if there are no defaults to define?
			astbuilder.AssignToInterface(tempVarIdent, dst.NewIdent(receiverIdent)),
			astbuilder.IfType(
				dst.NewIdent(tempVarIdent),
				overrideInterfaceType,
				runtimeDefaulterIdent,
				astbuilder.InvokeQualifiedFunc(runtimeDefaulterIdent, "CustomDefault")),
		},
	}

	fn.AddComments(fmt.Sprintf("applies defaults to the %s resource", receiver.Name()))
	return fn.DefineFunc()
}
