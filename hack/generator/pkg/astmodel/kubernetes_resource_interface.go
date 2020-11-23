/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"github.com/pkg/errors"
)

// These are some magical field names which we're going to use or generate
const (
	AzureNameProperty = "AzureName"
	OwnerProperty     = "Owner"
)

// WithKubernetesResourceInterfaceImpl creates a new interface with the specified ARM conversion functions
func (r *ResourceType) WithKubernetesResourceInterfaceImpl(
	idFactory IdentifierFactory,
	types Types) (*ResourceType, error) {

	resolvedSpec, err := types.FullyResolve(r.SpecType())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve resource spec type")
	}

	spec, err := TypeAsObjectType(resolvedSpec)
	if err != nil {
		return nil, errors.Wrapf(err, "resource spec %q did not contain an object", r.SpecType().String())
	}

	// Check the spec first to ensure it looks how we expect
	ownerProperty := idFactory.CreatePropertyName(OwnerProperty, Exported)
	_, ok := spec.Property(ownerProperty)
	if !ok {
		return nil, errors.Errorf("resource spec doesn't have %q property", ownerProperty)
	}

	nameProp, ok := spec.Property(AzureNameProperty)
	if !ok {
		return nil, errors.Errorf("resource spec doesn't have %q property", AzureNameProperty)
	}

	namePropType := nameProp.PropertyType()
	// resolve property type if it is a typename
	namePropType, err = types.FullyResolve(namePropType)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve type of resource Name property: %s", namePropType.String())
	}

	var nameFunc asFuncType

	// handle different types of Name property
	switch namePropType := namePropType.(type) {
	case *EnumType:
		if !namePropType.BaseType().Equals(StringType) {
			return nil, errors.Errorf("unable to handle non-string enum base type in Name property")
		}

		options := namePropType.Options()
		if len(options) == 1 {
			// if there is only one possible value,
			// we make an AzureName function that returns it, and do not
			// provide an AzureName property on the spec
			r = r.WithSpec(spec.WithoutProperty(AzureNameProperty))
			nameFunc = withFixedValueAzureNameFunction(options[0].Value)
		} else {
			// with multiple values, provide an AzureName function that casts from the
			// enum-valued AzureName property:

			// the property type must be a TypeName pointing to an enum at this
			// point in the pipeline so let's assert that:
			enumType := nameProp.PropertyType().(TypeName)
			nameFunc = withEnumAzureNameFunction(enumType)
		}

	case *PrimitiveType:
		if !namePropType.Equals(StringType) {
			return nil, errors.Errorf("cannot use type %s as type of Name property", namePropType.String())
		}

		nameFunc = azureNameFunction

	default:
		return nil, errors.Errorf("unsupported type for Name property: %s", namePropType.String())
	}

	azureNameFunc := &objectFunction{
		name:      AzureNameProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    nameFunc,
	}

	ownerFunc := &objectFunction{
		name:      OwnerProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    ownerFunction,
	}

	return r.WithInterface(NewInterfaceImplementation(
		MakeTypeName(MakeGenRuntimePackageReference(), "KubernetesResource"),
		ownerFunc,
		azureNameFunc)), nil
}

// objectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type objectFunction struct {
	name      string
	o         *ObjectType
	idFactory IdentifierFactory
	asFunc    asFuncType
}

type asFuncType func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl

// withEnumAzureNameFunction adds an AzureName() function that casts an AzureName property
// with an enum value to a string
func withEnumAzureNameFunction(propType Type) asFuncType {
	return func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
		receiverIdent := f.idFactory.CreateIdentifier(receiver.Name(), NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &ast.StarExpr{X: receiverType},
			Body: []ast.Stmt{
				&ast.ReturnStmt{
					Decs: ast.ReturnStmtDecorations{
						NodeDecs: ast.NodeDecs{
							Before: ast.NewLine,
						},
					},
					Results: []ast.Expr{
						&ast.CallExpr{
							// cast from the enum value to string
							Fun: ast.NewIdent("string"),
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent(receiverIdent),
									Sel: ast.NewIdent(AzureNameProperty),
								},
							},
						},
					},
				},
			},
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (string representation of %s)", propType.String()))
		fn.AddReturns("string")
		return fn.DefineFunc()
	}
}

// withFixedValueAzureNameFunction adds an AzureName() function that returns a fixed value
func withFixedValueAzureNameFunction(fixedValue string) asFuncType {

	// ensure fixedValue is quoted. This is always the case with enum values we pass,
	// but let's be safe:
	if len(fixedValue) == 0 {
		panic("cannot created fixed value AzureName function with empty fixed value")
	}

	if !(fixedValue[0] == '"' && fixedValue[len(fixedValue)-1] == '"') {
		fixedValue = fmt.Sprintf("%q", fixedValue)
	}

	return func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
		receiverIdent := f.idFactory.CreateIdentifier(receiver.Name(), NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &ast.StarExpr{X: receiverType},
			Body: []ast.Stmt{
				&ast.ReturnStmt{
					Decs: ast.ReturnStmtDecorations{
						NodeDecs: ast.NodeDecs{
							Before: ast.NewLine,
						},
					},
					Results: []ast.Expr{
						&ast.BasicLit{
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

var _ Function = &objectFunction{}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *objectFunction) Name() string {
	return k.name
}

func (k *objectFunction) RequiredPackageReferences() *PackageReferenceSet {
	// We only require GenRuntime
	return NewPackageReferenceSet(MakeGenRuntimePackageReference())
}

func (k *objectFunction) References() TypeNameSet {
	return k.o.References()
}

func (k *objectFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *ast.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

func (k *objectFunction) Equals(f Function) bool {
	typedF, ok := f.(*objectFunction)
	if !ok {
		return false
	}

	return k.o.Equals(typedF.o)
}

// IsKubernetesResourceProperty returns true if the supplied property name is one of our "magical" names
func IsKubernetesResourceProperty(name PropertyName) bool {
	return name == AzureNameProperty || name == OwnerProperty
}

func ownerFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {

	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &ast.StarExpr{
			X: receiver.AsType(codeGenerationContext),
		},
		Params: nil,
		Returns: []*ast.Field{
			{
				Type: &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(GenRuntimePackageName),
						Sel: ast.NewIdent("ResourceReference"),
					},
				},
			},
		},
		Body: []ast.Stmt{
			lookupGroupAndKindStmt(
				"group",
				"kind",
				&ast.SelectorExpr{
					X:   ast.NewIdent(receiverIdent),
					Sel: ast.NewIdent("Spec"),
				},
			),
			&ast.ReturnStmt{
				Results: []ast.Expr{
					createResourceReference(
						"group",
						"kind",
						&ast.SelectorExpr{
							X:   ast.NewIdent(receiverIdent),
							Sel: ast.NewIdent("Spec"),
						},
					),
				},
			},
		},
	}

	fn.AddComments("returns the ResourceReference of the owner, or nil if there is no owner")

	return fn.DefineFunc()
}

func lookupGroupAndKindStmt(
	groupIdent string,
	kindIdent string,
	specSelector *ast.SelectorExpr) *ast.AssignStmt {

	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent(groupIdent),
			ast.NewIdent(kindIdent),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent(GenRuntimePackageName),
					Sel: ast.NewIdent("LookupOwnerGroupKind"),
				},
				Args: []ast.Expr{
					specSelector,
				},
			},
		},
	}
}

func createResourceReference(
	groupIdent string,
	kindIdent string,
	specSelector *ast.SelectorExpr) ast.Expr {

	return astbuilder.AddrOf(
		&ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(GenRuntimePackageName),
				Sel: ast.NewIdent("ResourceReference"),
			},
			Elts: []ast.Expr{
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Name"),
					Value: &ast.SelectorExpr{
						X: &ast.SelectorExpr{
							X:   specSelector,
							Sel: ast.NewIdent(OwnerProperty),
						},
						Sel: ast.NewIdent("Name"),
					},
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("Group"),
					Value: ast.NewIdent(groupIdent),
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("Kind"),
					Value: ast.NewIdent(kindIdent),
				},
			},
		})
}

func azureNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &ast.StarExpr{
			X: receiverType,
		},
		Body: []ast.Stmt{
			&ast.ReturnStmt{
				Decs: ast.ReturnStmtDecorations{
					NodeDecs: ast.NodeDecs{
						Before: ast.NewLine,
					},
				},
				Results: []ast.Expr{
					&ast.SelectorExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(receiverIdent),
							Sel: ast.NewIdent("Spec"),
						},
						Sel: ast.NewIdent(AzureNameProperty),
					},
				},
			},
		},
	}

	fn.AddComments("returns the Azure name of the resource")
	fn.AddReturns("string")
	return fn.DefineFunc()
}
