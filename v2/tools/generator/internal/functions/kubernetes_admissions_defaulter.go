/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// DefaulterBuilder helps in building an interface implementation for admissions.Defaulter.
type DefaulterBuilder struct {
	resourceName astmodel.TypeName
	resource     *astmodel.ResourceType
	idFactory    astmodel.IdentifierFactory

	defaults []*ResourceFunction
}

// NewDefaulterBuilder creates a new DefaulterBuilder for the given object type.
func NewDefaulterBuilder(resourceName astmodel.TypeName, resource *astmodel.ResourceType, idFactory astmodel.IdentifierFactory) *DefaulterBuilder {
	return &DefaulterBuilder{
		resourceName: resourceName,
		resource:     resource,
		idFactory:    idFactory,
	}
}

// AddDefault adds a default function to the set of default functions to be applied to the given object
func (d *DefaulterBuilder) AddDefault(f *ResourceFunction) {
	if !d.resource.Equals(f.Resource(), astmodel.EqualityOverrides{}) {
		panic("cannot add default function on non-matching object types")
	}
	d.defaults = append(d.defaults, f)
}

// ToInterfaceImplementation creates an InterfaceImplementation that implements the admissions.Defaulter interface.
// This implementation includes calls to all defaults registered with this DefaulterBuilder via the AddDefault function,
// as well as helper functions that allow additional handcrafted defaults to be injected by
// implementing the genruntime.Defaulter interface.
func (d *DefaulterBuilder) ToInterfaceImplementation() *astmodel.InterfaceImplementation {
	group, version := d.resourceName.PackageReference.GroupVersion()

	// e.g. group = "microsoft.network.azure.com"
	// e.g. resource = "backendaddresspools"
	// e.g. version = "v1"

	resource := d.resourceName.Name()

	group = strings.ToLower(group + astmodel.GroupSuffix)
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
			"verbs=create;update,versions=%s,name=%s,admissionReviewVersions=v1beta1",
		path,
		group,
		resource,
		version,
		name)

	funcs := []astmodel.Function{
		NewResourceFunction(
			"Default",
			d.resource,
			d.idFactory,
			d.defaultFunction,
			astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference)),
		NewResourceFunction(
			"defaultImpl",
			d.resource,
			d.idFactory,
			d.localDefault,
			astmodel.NewPackageReferenceSet(astmodel.GenRuntimeReference)),
	}

	// Add the actual individual default functions
	for _, def := range d.defaults {
		funcs = append(funcs, def)
	}

	return astmodel.NewInterfaceImplementation(
		astmodel.DefaulterInterfaceName,
		funcs...).WithAnnotation(annotation)
}

func (d *DefaulterBuilder) localDefault(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	defaults := make([]dst.Stmt, 0, len(d.defaults))
	for _, def := range d.defaults {
		defaults = append(
			defaults,
			&dst.ExprStmt{
				X: astbuilder.CallQualifiedFunc(receiverIdent, def.Name()),
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

func (d *DefaulterBuilder) defaultFunction(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)
	tempVarIdent := "temp"
	runtimeDefaulterIdent := "runtimeDefaulter"

	overrideInterfaceType := astmodel.GenRuntimeDefaulterInterfaceName.AsType(codeGenerationContext)

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
