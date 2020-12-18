/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"github.com/pkg/errors"
)

// These are some magical field names which we're going to use or generate
const (
	AzureNameProperty = "AzureName"
	SetAzureNameFunc  = "SetAzureName"
	OwnerProperty     = "Owner"
)

// AddKubernetesResourceInterfaceImpls adds the required interfaces for
// the resource to be a Kubernetes resource
func AddKubernetesResourceInterfaceImpls(
	resourceName TypeName,
	r *ResourceType,
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

	azureNameProp, ok := spec.Property(AzureNameProperty)
	if !ok {
		return nil, errors.Errorf("resource spec doesn't have %q property", AzureNameProperty)
	}

	getNameFunction, setNameFunction, err := getAzureNameFunctionsForType(&r, spec, azureNameProp.PropertyType(), types)
	if err != nil {
		return nil, err
	}

	getAzureNameProperty := &objectFunction{
		name:      AzureNameProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    getNameFunction,
	}

	getOwnerProperty := &objectFunction{
		name:      OwnerProperty,
		o:         spec,
		idFactory: idFactory,
		asFunc:    ownerFunction,
	}

	r = r.WithInterface(NewInterfaceImplementation(
		MakeTypeName(GenRuntimeReference, "KubernetesResource"),
		getAzureNameProperty,
		getOwnerProperty))

	if setNameFunction != nil {
		// this function applies to Spec not the resource
		// re-fetch the spec ObjectType since the getAzureNameFunctionsForType
		// could have updated it
		spec := r.SpecType()
		spec, err = types.FullyResolve(spec)
		if err != nil {
			return nil, err
		}

		specObj := spec.(*ObjectType)

		r = r.WithSpec(specObj.WithFunction(
			&objectFunction{
				name:      SetAzureNameFunc,
				o:         specObj,
				idFactory: idFactory,
				asFunc:    setNameFunction,
			}))
	}

	if setNameFunction != nil {
		// we also need to generate a Defaulter implementation to default AzureName
		r = r.WithInterface(generateDefaulter(resourceName, spec, idFactory))
	}

	return r, nil
}

var admissionPackageReference PackageReference = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/webhook/admission")
var DefaulterInterfaceName = MakeTypeName(admissionPackageReference, "Defaulter")

func generateDefaulter(resourceName TypeName, spec *ObjectType, idFactory IdentifierFactory) *InterfaceImplementation {
	lpr, _ := resourceName.PackageReference.AsLocalPackage()

	group := lpr.group              // e.g. "microsoft.network.infra.azure.com"
	resource := resourceName.Name() // e.g. "backendaddresspools"
	version := lpr.version          // e.g. "v1"

	group = strings.ToLower(group + GroupSuffix)
	nonPluralResource := strings.ToLower(resource)
	resource = strings.ToLower(resource) + "s" // TODO: this should come from resource?

	// e.g. "mutate-microsoft-network-infra-azure-com-v1-backendaddresspool"
	// note that this must match _exactly_ how controller-runtime generates the path
	// or it will not work!
	path := fmt.Sprintf("/mutate-%s-%s-%s", strings.ReplaceAll(group, ".", "-"), version, nonPluralResource)

	// e.g.  "default.v123.backendaddresspool.infra.azure.com"
	name := fmt.Sprintf("default.%s.%s.%s", version, resource, group)

	annotation := fmt.Sprintf(
		"+kubebuilder:webhook:path=%s,mutating=true,sideEffects=None,"+
			"matchPolicy=Exact,failurePolicy=fail,groups=%s,resources=%s,"+
			"verbs=create;update,versions=%s,name=%s",
		path,
		group,
		resource,
		version,
		name)

	return NewInterfaceImplementation(
		DefaulterInterfaceName,
		&objectFunction{
			name:      "Default",
			o:         spec,
			idFactory: idFactory,
			asFunc:    defaultAzureNameFunction,
		}).WithAnnotation(annotation)
}

// note that this can, as a side-effect, update the resource type
// it is a bit ugly!
func getAzureNameFunctionsForType(r **ResourceType, spec *ObjectType, t Type, types Types) (asFuncType, asFuncType, error) {
	// handle different types of AzureName property
	switch azureNamePropType := t.(type) {
	case ValidatedType:
		if !azureNamePropType.ElementType().Equals(StringType) {
			return nil, nil, errors.Errorf("unable to handle non-string validated types in AzureName property")
		}

		validations := azureNamePropType.Validations().(StringValidations)
		if validations.Pattern != nil {
			if validations.Pattern.String() == "^.*/default$" {
				*r = (*r).WithSpec(spec.WithoutProperty(AzureNameProperty))
				return fixedValueGetAzureNameFunction("default"), nil, nil // no SetAzureName for this case
			} else {
				// ignoring for now:
				//return nil, errors.Errorf("unable to handle pattern in Name property: %s", validations.Pattern.String())
				return getStringAzureNameFunction, setStringAzureNameFunction, nil
			}
		} else {
			// ignoring length validations for now
			//return nil, errors.Errorf("unable to handle validations on Name property …TODO")
			return getStringAzureNameFunction, setStringAzureNameFunction, nil
		}

	case TypeName:
		// resolve property type if it is a typename
		resolvedPropType, err := types.FullyResolve(azureNamePropType)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "unable to resolve type of resource Name property: %s", azureNamePropType.String())
		}

		if t, ok := resolvedPropType.(*EnumType); ok {
			if !t.BaseType().Equals(StringType) {
				return nil, nil, errors.Errorf("unable to handle non-string enum base type in Name property")
			}

			options := t.Options()
			if len(options) == 1 {
				// if there is only one possible value,
				// we make an AzureName function that returns it, and do not
				// provide an AzureName property on the spec
				*r = (*r).WithSpec(spec.WithoutProperty(AzureNameProperty))
				return fixedValueGetAzureNameFunction(options[0].Value), nil, nil // no SetAzureName for this case
			} else {
				// with multiple values, provide an AzureName function that casts from the
				// enum-valued AzureName property:
				return getEnumAzureNameFunction(azureNamePropType), setEnumAzureNameFunction(azureNamePropType), nil
			}
		} else {
			return nil, nil, errors.Errorf("unable to produce AzureName()/SetAzureName() for Name property with type %s", resolvedPropType.String())
		}

	case *PrimitiveType:
		if !azureNamePropType.Equals(StringType) {
			return nil, nil, errors.Errorf("cannot use type %s as type of AzureName property", azureNamePropType.String())
		}

		return getStringAzureNameFunction, setStringAzureNameFunction, nil

	default:
		return nil, nil, errors.Errorf("unsupported type for AzureName property: %s", azureNamePropType.String())
	}
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

// getEnumAzureNameFunction adds an AzureName() function that casts the AzureName property
// with an enum value to a string
func getEnumAzureNameFunction(enumType TypeName) asFuncType {
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
									X: &ast.SelectorExpr{
										X:   ast.NewIdent(receiverIdent),
										Sel: ast.NewIdent("Spec"),
									},
									Sel: ast.NewIdent(AzureNameProperty),
								},
							},
						},
					},
				},
			},
		}

		fn.AddComments(fmt.Sprintf("returns the Azure name of the resource (string representation of %s)", enumType.String()))
		fn.AddReturns("string")
		return fn.DefineFunc()
	}
}

// setEnumAzureNameFunction returns a function that sets the AzureName property to the result of casting
// the argument string to the given enum type
func setEnumAzureNameFunction(enumType TypeName) asFuncType {
	return func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
		receiverIdent := f.idFactory.CreateIdentifier(receiver.Name(), NotExported)
		receiverType := receiver.AsType(codeGenerationContext)

		azureNameProp := &ast.SelectorExpr{
			X:   ast.NewIdent(receiverIdent),
			Sel: ast.NewIdent(AzureNameProperty),
		}

		enumTypeAST := ast.NewIdent(enumType.Name())

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  &ast.StarExpr{X: receiverType},
			Body: []ast.Stmt{
				astbuilder.SimpleAssignment(
					azureNameProp,
					token.ASSIGN,
					&ast.CallExpr{
						// cast from the string value to the enum
						Fun:  enumTypeAST,
						Args: []ast.Expr{ast.NewIdent("azureName")},
					},
				),
			},
		}

		fn.AddComments(fmt.Sprintf("sets the Azure name from the given %s value", enumType.String()))
		fn.AddParameter("azureName", ast.NewIdent("string"))
		return fn.DefineFunc()
	}
}

// fixedValueGetAzureNameFunction adds an AzureName() function that returns a fixed value
func fixedValueGetAzureNameFunction(fixedValue string) asFuncType {

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
	return NewPackageReferenceSet(GenRuntimeReference)
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

// ownerFunction returns a function that returns the owner of the resource
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

// defaultAzureNameFunction returns a function that defaults the AzureName property of the resource spec
// to the Name property of the resource spec
func defaultAzureNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	specSelector := &ast.SelectorExpr{
		X:   ast.NewIdent(receiverIdent),
		Sel: ast.NewIdent("Spec"),
	}

	azureNameProp := &ast.SelectorExpr{
		X:   specSelector,
		Sel: ast.NewIdent(AzureNameProperty),
	}

	nameProp := &ast.SelectorExpr{
		X:   ast.NewIdent(receiverIdent),
		Sel: ast.NewIdent("Name"), // this comes from ObjectMeta
	}

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &ast.StarExpr{
			X: receiverType,
		},
		Body: []ast.Stmt{
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X:  ast.Clone(azureNameProp).(ast.Expr),
					Op: token.EQL,
					Y:  &ast.BasicLit{Kind: token.STRING, Value: "\"\""},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						astbuilder.SimpleAssignment(
							azureNameProp,
							token.ASSIGN,
							nameProp),
					},
				},
			},
		},
	}

	fn.AddComments("defaults the Azure name of the resource to the Kubernetes name")
	return fn.DefineFunc()
}

// setStringAzureNameFunction returns a function that sets the Name property of
// the resource spec to the argument string
func setStringAzureNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &ast.StarExpr{
			X: receiverType,
		},
		Body: []ast.Stmt{
			astbuilder.SimpleAssignment(
				&ast.SelectorExpr{
					X:   ast.NewIdent(receiverIdent),
					Sel: ast.NewIdent(AzureNameProperty),
				},
				token.ASSIGN,
				ast.NewIdent("azureName")),
		},
	}

	fn.AddComments("sets the Azure name of the resource")
	fn.AddParameter("azureName", ast.NewIdent("string"))
	return fn.DefineFunc()
}

// getStringAzureNameFunction returns a function that returns the Name property of the resource spec
func getStringAzureNameFunction(k *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *ast.FuncDecl {
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
