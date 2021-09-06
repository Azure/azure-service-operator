/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
	"github.com/pkg/errors"

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
	types astmodel.Types) (*astmodel.ResourceType, error) {

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

	getOwnerProperty := functions.NewObjectFunction(astmodel.OwnerProperty, idFactory, ownerFunction)
	getOwnerProperty.AddPackageReference(astmodel.GenRuntimeReference)

	getSpecFunction := functions.NewObjectFunction("GetSpec", idFactory, createGetSpecFunction)
	getStatusFunction := functions.NewObjectFunction("GetStatus", idFactory, createGetStatusFunction)

	kubernetesResourceImplementation := astmodel.NewInterfaceImplementation(
		astmodel.KubernetesResourceType,
		getAzureNameProperty,
		getOwnerProperty,
		getSpecFunction,
		getStatusFunction)

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
		if !azureNamePropType.ElementType().Equals(astmodel.StringType) {
			return nil, nil, errors.Errorf("unable to handle non-string validated types in AzureName property")
		}

		validations := azureNamePropType.Validations().(astmodel.StringValidations)
		if validations.Pattern != nil {
			if validations.Pattern.String() == "^.*/default$" {
				*r = (*r).WithSpec(spec.WithoutProperty(astmodel.AzureNameProperty))
				return fixedValueGetAzureNameFunction("default"), nil, nil // no SetAzureName for this case
			} else {
				// ignoring for now:
				// return nil, errors.Errorf("unable to handle pattern in Name property: %s", validations.Pattern.String())
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
			if !t.BaseType().Equals(astmodel.StringType) {
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
		if !azureNamePropType.Equals(astmodel.StringType) {
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

// ownerFunction returns a function that returns the owner of the resource
func ownerFunction(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiver.AsType(codeGenerationContext),
		},
		Params: nil,
		Returns: []*dst.Field{
			{
				Type: &dst.StarExpr{
					X: &dst.SelectorExpr{
						X:   dst.NewIdent(astmodel.GenRuntimePackageName),
						Sel: dst.NewIdent("ResourceReference"),
					},
				},
			},
		},
		Body: []dst.Stmt{
			lookupGroupAndKindStmt(
				"group",
				"kind",
				&dst.SelectorExpr{
					X:   dst.NewIdent(receiverIdent),
					Sel: dst.NewIdent("Spec"),
				},
			),
			&dst.ReturnStmt{
				Results: []dst.Expr{
					createResourceReference(
						"group",
						"kind",
						receiverIdent,
					),
				},
			},
		},
	}

	fn.AddComments("returns the ResourceReference of the owner, or nil if there is no owner")

	return fn.DefineFunc()
}

func createGetSpecFunction(
	f *functions.ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string) *dst.FuncDecl {
	receiverIdent := f.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)
	receiverType := astmodel.NewOptionalType(receiver)

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverType.AsType(genContext),
		Name:          "GetSpec",
		Body: astbuilder.Statements(
			astbuilder.Returns(
				astbuilder.AddrOf(astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec")))),
	}

	fn.AddReturn(astmodel.ConvertibleSpecInterfaceType.AsType(genContext))
	fn.AddComments("returns the specification of this resource")

	return fn.DefineFunc()
}

func createGetStatusFunction(
	idFactory *functions.ObjectFunction,
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
					X:   dst.NewIdent(astmodel.GenRuntimePackageName),
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
	groupIdent string,
	kindIdent string,
	receiverIdent string) dst.Expr {

	specSelector := &dst.SelectorExpr{
		X:   dst.NewIdent(receiverIdent),
		Sel: dst.NewIdent("Spec"),
	}

	return astbuilder.AddrOf(
		&dst.CompositeLit{
			Type: &dst.SelectorExpr{
				X:   dst.NewIdent(astmodel.GenRuntimePackageName),
				Sel: dst.NewIdent("ResourceReference"),
			},
			Elts: []dst.Expr{
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Group"),
					Value: dst.NewIdent(groupIdent),
				},
				&dst.KeyValueExpr{
					Key:   dst.NewIdent("Kind"),
					Value: dst.NewIdent(kindIdent),
				},
				&dst.KeyValueExpr{
					Key: dst.NewIdent("Namespace"),
					Value: &dst.SelectorExpr{
						X:   dst.NewIdent(receiverIdent),
						Sel: dst.NewIdent("Namespace"),
					},
				},
				&dst.KeyValueExpr{
					Key: dst.NewIdent("Name"),
					Value: &dst.SelectorExpr{
						X: &dst.SelectorExpr{
							X:   specSelector,
							Sel: dst.NewIdent(astmodel.OwnerProperty),
						},
						Sel: dst.NewIdent("Name"),
					},
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
