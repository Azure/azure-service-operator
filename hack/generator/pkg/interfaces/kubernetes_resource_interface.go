/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// AddKubernetesResourceInterfaceImpls adds the required interfaces for
// the resource to be a Kubernetes resource
func AddKubernetesResourceInterfaceImpls(
	resourceName astmodel.TypeName,
	r *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
	types astmodel.Types,
	resourceARMType string) (*astmodel.ResourceType, error) {

	resolvedSpec, err := types.FullyResolve(r.SpecType())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve resource spec type")
	}

	spec, ok := astmodel.AsObjectType(resolvedSpec)
	if !ok {
		return nil, errors.Errorf("resource spec %q did not contain an object", r.SpecType().String())
	}

	// Check the spec first to ensure it looks how we expect
	ownerProperty := idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported)
	_, ok = spec.Property(ownerProperty)
	if !ok {
		return nil, errors.Errorf("resource spec doesn't have %q property", ownerProperty)
	}

	azureNameProp, ok := spec.Property(astmodel.AzureNameProperty)
	if !ok {
		return nil, errors.Errorf("resource spec doesn't have %q property", astmodel.AzureNameProperty)
	}

	getNameFunction, setNameFunction, err := getAzureNameFunctionsForType(&r, spec, azureNameProp.PropertyType(), types)
	if err != nil {
		return nil, err
	}

	getAzureNameProperty := functions.NewObjectFunction(astmodel.AzureNameProperty, idFactory, getNameFunction)
	getAzureNameProperty.AddPackageReference(astmodel.GenRuntimeReference)

	getOwnerProperty := functions.NewObjectFunction(astmodel.OwnerProperty, idFactory, newOwnerFunction(r))
	getOwnerProperty.AddPackageReference(astmodel.GenRuntimeReference)

	getSpecFunction := functions.NewObjectFunction("GetSpec", idFactory, createGetSpecFunction)

	getTypeFunction := functions.NewObjectFunction("Get"+astmodel.TypeProperty, idFactory, newGetTypeFunction(resourceARMType))
	getTypeFunction.AddPackageReference(astmodel.GenRuntimeReference)

	getResourceKindFunction := functions.NewObjectFunction("GetResourceKind", idFactory, newGetResourceKindFunction(r))
	getResourceKindFunction.AddPackageReference(astmodel.GenRuntimeReference)

	fns := []astmodel.Function{
		getAzureNameProperty,
		getOwnerProperty,
		getSpecFunction,
		getTypeFunction,
		getResourceKindFunction,
	}

	if r.StatusType() != nil {
		// Skip Status functions if no status
		getStatusFunction := functions.NewObjectFunction("GetStatus", idFactory, createGetStatusFunction)
		setStatusFunction := functions.NewResourceStatusSetterFunction(r, idFactory)

		fns = append(fns, getStatusFunction, setStatusFunction)
	}

	kubernetesResourceImplementation := astmodel.NewInterfaceImplementation(astmodel.KubernetesResourceType, fns...)

	r = r.WithInterface(kubernetesResourceImplementation)

	if setNameFunction != nil {
		// this function applies to Spec not the resource
		// re-fetch the spec ObjectType since the getAzureNameFunctionsForType
		// could have updated it
		spec := r.SpecType()
		spec, err = types.FullyResolve(spec)
		if err != nil {
			return nil, err
		}

		specObj := spec.(*astmodel.ObjectType)

		setFn := functions.NewObjectFunction(astmodel.SetAzureNameFunc, idFactory, setNameFunction)
		setFn.AddPackageReference(astmodel.GenRuntimeReference)

		r = r.WithSpec(specObj.WithFunction(setFn))
	}

	// Add defaults
	defaulterBuilder := astmodel.NewDefaulterBuilder(resourceName, r, idFactory)
	if setNameFunction != nil {
		defaulterBuilder.AddDefault(astmodel.NewDefaultAzureNameFunction(r, idFactory))
	}
	r = r.WithInterface(defaulterBuilder.ToInterfaceImplementation())

	// Add validations
	validatorBuilder := astmodel.NewValidatorBuilder(resourceName, r, idFactory)
	validatorBuilder.AddValidation(astmodel.ValidationKindCreate, astmodel.NewValidateResourceReferencesFunction(r, idFactory))
	validatorBuilder.AddValidation(astmodel.ValidationKindUpdate, astmodel.NewValidateResourceReferencesFunction(r, idFactory))

	r = r.WithInterface(validatorBuilder.ToInterfaceImplementation())

	return r, nil
}

// note that this can, as a side-effect, update the resource type
// it is a bit ugly!
func getAzureNameFunctionsForType(r **astmodel.ResourceType, spec *astmodel.ObjectType, t astmodel.Type, types astmodel.Types) (functions.ObjectFunctionHandler, functions.ObjectFunctionHandler, error) {
	// handle different types of AzureName property
	switch azureNamePropType := t.(type) {
	case *astmodel.ValidatedType:
		if !astmodel.TypeEquals(azureNamePropType.ElementType(), astmodel.StringType) {
			return nil, nil, errors.Errorf("unable to handle non-string validated types in AzureName property")
		}

		validations := azureNamePropType.Validations().(astmodel.StringValidations)
		if len(validations.Patterns) != 0 {
			if len(validations.Patterns) == 1 &&
				validations.Patterns[0].String() == "^.*/default$" {
				*r = (*r).WithSpec(spec.WithoutProperty(astmodel.AzureNameProperty))
				return fixedValueGetAzureNameFunction("default"), nil, nil // no SetAzureName for this case
			} else {
				// ignoring for now:
				klog.Warningf("unable to handle pattern in Name property: %s", validations.Patterns[0].String())
				return getStringAzureNameFunction, setStringAzureNameFunction, nil
			}
		} else {
			// ignoring length validations for now
			// return nil, errors.Errorf("unable to handle validations on Name property …TODO")
			return getStringAzureNameFunction, setStringAzureNameFunction, nil
		}

	case astmodel.TypeName:
		// resolve property type if it is a typename
		resolvedPropType, err := types.FullyResolve(azureNamePropType)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "unable to resolve type of resource Name property: %s", azureNamePropType.String())
		}

		if t, ok := resolvedPropType.(*astmodel.EnumType); ok {
			if !astmodel.TypeEquals(t.BaseType(), astmodel.StringType) {
				return nil, nil, errors.Errorf("unable to handle non-string enum base type in Name property")
			}

			options := t.Options()
			if len(options) == 1 {
				// if there is only one possible value,
				// we make an AzureName function that returns it, and do not
				// provide an AzureName property on the spec
				*r = (*r).WithSpec(spec.WithoutProperty(astmodel.AzureNameProperty))
				return fixedValueGetAzureNameFunction(options[0].Value), nil, nil // no SetAzureName for this case
			} else {
				// with multiple values, provide an AzureName function that casts from the
				// enum-valued AzureName property:
				return getEnumAzureNameFunction(azureNamePropType), setEnumAzureNameFunction(azureNamePropType), nil
			}
		} else {
			return nil, nil, errors.Errorf("unable to produce AzureName()/SetAzureName() for Name property with type %s", resolvedPropType.String())
		}

	case *astmodel.PrimitiveType:
		if !astmodel.TypeEquals(azureNamePropType, astmodel.StringType) {
			return nil, nil, errors.Errorf("cannot use type %s as type of AzureName property", azureNamePropType.String())
		}

		return getStringAzureNameFunction, setStringAzureNameFunction, nil

	default:
		return nil, nil, errors.Errorf("unsupported type for AzureName property: %s", azureNamePropType.String())
	}
}

// getEnumAzureNameFunction adds an AzureName() function that casts the AzureName property
// with an enum value to a string
func getEnumAzureNameFunction(enumType astmodel.TypeName) functions.ObjectFunctionHandler {
	return func(f *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &dst.StarExpr{X: receiverType},
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.CallFunc("string", astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec", astmodel.AzureNameProperty)))),
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (string representation of %s)", enumType.String()))
		fn.AddReturns("string")
		return fn.DefineFunc()
	}
}

// setEnumAzureNameFunction returns a function that sets the AzureName property to the result of casting
// the argument string to the given enum type
func setEnumAzureNameFunction(enumType astmodel.TypeName) functions.ObjectFunctionHandler {
	return func(f *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		azureNameProp := astbuilder.Selector(dst.NewIdent(receiverIdent), astmodel.AzureNameProperty)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &dst.StarExpr{X: receiverType},
			Body: []dst.Stmt{
				astbuilder.SimpleAssignment(
					azureNameProp,
					astbuilder.CallFunc(enumType.Name(), dst.NewIdent("azureName")),
				),
			},
		}

		fn.AddComments(fmt.Sprintf("sets the Azure name from the given %s value", enumType.String()))
		fn.AddParameter("azureName", dst.NewIdent("string"))
		return fn.DefineFunc()
	}
}

// fixedValueGetAzureNameFunction adds an AzureName() function that returns a fixed value
func fixedValueGetAzureNameFunction(fixedValue string) functions.ObjectFunctionHandler {
	// ensure fixedValue is quoted. This is always the case with enum values we pass,
	// but let's be safe:
	if len(fixedValue) == 0 {
		panic("cannot created fixed value AzureName function with empty fixed value")
	}

	if !(fixedValue[0] == '"' && fixedValue[len(fixedValue)-1] == '"') {
		fixedValue = fmt.Sprintf("%q", fixedValue)
	}

	return func(f *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &dst.StarExpr{X: receiverType},
			Body: []dst.Stmt{
				&dst.ReturnStmt{
					Decs: dst.ReturnStmtDecorations{
						NodeDecs: dst.NodeDecs{
							Before: dst.NewLine,
						},
					},
					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: fixedValue,
						},
					},
				},
			},
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (always %s)", fixedValue))
		fn.AddReturns("string")
		return fn.DefineFunc()
	}
}

// IsKubernetesResourceProperty returns true if the supplied property name is one of our "magical" names
func IsKubernetesResourceProperty(name astmodel.PropertyName) bool {
	return name == astmodel.AzureNameProperty || name == astmodel.OwnerProperty
}

// newOwnerFunction creates the Owner function declaration. This has two possible formats.
// For normal resources:
//	func (<receiver> *<receiver>) Owner() *genruntime.ResourceReference {
//		group, kind := genruntime.LookupOwnerGroupKind(<receiver>.Spec)
//		return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: <receiver>.Namespace, Name: <receiver>.Spec.Owner.Name}
//	}
// For extension resources:
//	func (<receiver> *<receiver>) Owner() *genruntime.ResourceReference {
//		return &genruntime.ResourceReference{Group: <receiver>.Spec.Owner.Group, Kind: <receiver>.Spec.Owner.Kind, name: <receiver>.Spec.Owner.Name}
//	}
func newOwnerFunction(r *astmodel.ResourceType) func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)

		specSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")
		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType: &dst.StarExpr{
				X: receiver.AsType(codeGenerationContext),
			},
			Params: nil,
		}

		fn.AddReturn(astbuilder.Dereference(astmodel.ResourceReferenceType.AsType(codeGenerationContext)))
		fn.AddComments("returns the ResourceReference of the owner, or nil if there is no owner")

		switch r.Kind() {
		case astmodel.ResourceKindNormal:
			fn.AddStatements(
				lookupGroupAndKindStmt("group", "kind", specSelector),
				astbuilder.Returns(createResourceReference(dst.NewIdent("group"), dst.NewIdent("kind"), receiverIdent)))
		case astmodel.ResourceKindExtension:
			owner := astbuilder.Selector(specSelector, astmodel.OwnerProperty)
			group := astbuilder.Selector(owner, "Group")
			kind := astbuilder.Selector(owner, "Kind")

			fn.AddStatements(astbuilder.Returns(createResourceReference(group, kind, receiverIdent)))
		default:
			panic(fmt.Sprintf("unknown resource kind: %s", r.Kind()))
		}

		return fn.DefineFunc()
	}
}

func createGetSpecFunction(
	f *functions.ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {
	receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := astmodel.NewOptionalType(receiver)

	ret := astbuilder.Returns(astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverType.AsType(genContext),
		Name:          "GetSpec",
		Body:          astbuilder.Statements(ret),
	}

	fn.AddReturn(astmodel.ConvertibleSpecInterfaceType.AsType(genContext))
	fn.AddComments("returns the specification of this resource")

	return fn.DefineFunc()
}

func createGetStatusFunction(
	f *functions.ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {
	receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := astmodel.NewOptionalType(receiver)

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverType.AsType(genContext),
		Name:          "GetStatus",
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")))),
	}

	fn.AddReturn(astmodel.ConvertibleStatusInterfaceType.AsType(genContext))
	fn.AddComments("returns the status of this resource")

	return fn.DefineFunc()
}

// newGetTypeFunction returns a function that returns the type of the resource (such as microsoft.compute/disks)
func newGetTypeFunction(armType string) func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := astmodel.NewOptionalType(receiver)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverType.AsType(codeGenerationContext),
			Params:        nil,
			Body: astbuilder.Statements(
				astbuilder.Returns(astbuilder.TextLiteral(armType))),
		}

		fn.AddComments(fmt.Sprintf("returns the ARM Type of the resource. This is always %q", strings.Trim(armType, "\"")))
		fn.AddReturns("string")

		return fn.DefineFunc()
	}
}

// newGetResourceKindFunction creates a function that returns the kind of resource.
// The generated function is a simple getter which will return either genruntime.ResourceKindNormal or genruntime.ResourceKindExtension:
//	func (<receiver> *<receiver>) Kind() genruntime.ResourceKind {
//		return genruntime.ResourceKindNormal
//	}
func newGetResourceKindFunction(r *astmodel.ResourceType) func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
		receiverType := astmodel.NewOptionalType(receiver)

		var resourceKind string
		switch r.Kind() {
		case astmodel.ResourceKindNormal:
			resourceKind = "ResourceKindNormal"
		case astmodel.ResourceKindExtension:
			resourceKind = "ResourceKindExtension"
		default:
			panic(fmt.Sprintf("unknown resource kind %s", r.Kind()))
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverType.AsType(codeGenerationContext),
			Params:        nil,
			Body: astbuilder.Statements(
				astbuilder.Returns(
					astbuilder.Selector(dst.NewIdent(astmodel.GenRuntimeReference.PackageName()), resourceKind))),
		}

		fn.AddComments("returns the kind of the resource")
		fn.AddReturn(astmodel.ResourceKindType.AsType(codeGenerationContext))

		return fn.DefineFunc()
	}
}

func lookupGroupAndKindStmt(
	groupIdent string,
	kindIdent string,
	specSelector *dst.SelectorExpr) *dst.AssignStmt {

	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.NewIdent(groupIdent),
			dst.NewIdent(kindIdent),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent(astmodel.GenRuntimeReference.PackageName()),
					Sel: dst.NewIdent("LookupOwnerGroupKind"),
				},
				Args: []dst.Expr{
					specSelector,
				},
			},
		},
	}
}

func createResourceReference(
	group dst.Expr,
	kind dst.Expr,
	receiverIdent string) dst.Expr {

	specSelector := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Spec"),
	}

	return astbuilder.AddrOf(
		&dst.CompositeLit{
			Type: &dst.SelectorExpr{
				X:   dst.NewIdent(astmodel.GenRuntimeReference.PackageName()),
				Sel: dst.NewIdent("ResourceReference"),
			},
			Elts: []dst.Expr{
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Group"),
					Value: group,
				},
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Kind"),
					Value: kind,
				},
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Namespace"),
					Value: astbuilder.Selector(dst.NewIdent(receiverIdent), "Namespace"),
				},
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Name"),
					Value: astbuilder.Selector(astbuilder.Selector(specSelector, astmodel.OwnerProperty), "Name"),
				},
			},
		})
}

// setStringAzureNameFunction returns a function that sets the Name property of
// the resource spec to the argument string
func setStringAzureNameFunction(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			astbuilder.QualifiedAssignment(
				dst.NewIdent(receiverIdent),
				astmodel.AzureNameProperty,
				token.ASSIGN,
				dst.NewIdent("azureName")),
		},
	}

	fn.AddComments("sets the Azure name of the resource")
	fn.AddParameter("azureName", dst.NewIdent("string"))
	return fn.DefineFunc()
}

// getStringAzureNameFunction returns a function that returns the Name property of the resource spec
func getStringAzureNameFunction(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			&dst.ReturnStmt{
				Decs: dst.ReturnStmtDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
					},
				},
				Results: []dst.Expr{
					&dst.SelectorExpr{
						X: &dst.SelectorExpr{
							X:   dst.NewIdent(receiverIdent),
							Sel: dst.NewIdent("Spec"),
						},
						Sel: dst.NewIdent(astmodel.AzureNameProperty),
					},
				},
			},
		},
	}

	fn.AddComments("returns the Azure name of the resource")
	fn.AddReturns("string")
	return fn.DefineFunc()
}
