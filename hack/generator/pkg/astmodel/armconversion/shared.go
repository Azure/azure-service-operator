/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"
	"sync"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	ast "github.com/dave/dst"
)

type conversionBuilder struct {
	receiverIdent              string
	receiverTypeExpr           ast.Expr
	armTypeIdent               string
	codeGenerationContext      *astmodel.CodeGenerationContext
	idFactory                  astmodel.IdentifierFactory
	isSpecType                 bool
	methodName                 string
	kubeType                   *astmodel.ObjectType
	armType                    *astmodel.ObjectType
	propertyConversionHandlers []propertyConversionHandler
}

type TypeKind int

const (
	OrdinaryType TypeKind = iota
	SpecType
	StatusType
)

func (builder conversionBuilder) propertyConversionHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	for _, conversionHandler := range builder.propertyConversionHandlers {
		stmts := conversionHandler(toProp, fromType)
		if len(stmts) > 0 {
			return stmts
		}
	}

	panic(fmt.Sprintf("No property found for %s in method %s\nFrom: %+v\nTo: %+v", toProp.PropertyName(), builder.methodName, *builder.kubeType, *builder.armType))
}

// deepCopyJSON special cases copying JSON-type fields to call the DeepCopy method.
// It generates code that looks like:
//     <destination> = *<source>.DeepCopy()
func (builder *conversionBuilder) deepCopyJSON(
	params complexPropertyConversionParameters) []ast.Stmt {
	newSource := &ast.UnaryExpr{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   params.Source(),
				Sel: ast.NewIdent("DeepCopy"),
			},
			Args: []ast.Expr{},
		},
		Op: token.MUL,
	}
	assignmentHandler := params.assignmentHandler
	if assignmentHandler == nil {
		assignmentHandler = assignmentHandlerAssign
	}
	return []ast.Stmt{
		assignmentHandler(params.Destination(), newSource),
	}
}

type propertyConversionHandler = func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []ast.Stmt

var once sync.Once
var azureNameProperty *astmodel.PropertyDefinition

func initializeAzureName(idFactory astmodel.IdentifierFactory) {
	azureNameFieldDescription := "The name of the resource in Azure. This is often the same as" +
		" the name of the resource in Kubernetes but it doesn't have to be."
	azureNameProperty = astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(astmodel.AzureNameProperty, astmodel.Exported),
		idFactory.CreateIdentifier(astmodel.AzureNameProperty, astmodel.NotExported),
		astmodel.StringType).WithDescription(azureNameFieldDescription)
}

// GetAzureNameProperty returns the special "AzureName" field
func GetAzureNameProperty(idFactory astmodel.IdentifierFactory) *astmodel.PropertyDefinition {
	once.Do(func() { initializeAzureName(idFactory) })

	return azureNameProperty
}

func getReceiverObjectType(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *astmodel.ObjectType {
	// Determine the type we're operating on
	rt, err := codeGenerationContext.GetImportedDefinition(receiver)
	if err != nil {
		panic(err)
	}

	receiverType, ok := rt.Type().(*astmodel.ObjectType)
	if !ok {
		// Don't expect to have any wrapper types left at this point
		panic(fmt.Sprintf("receiver for ArmConversionFunction is not of expected type. TypeName: %v, Type %T", receiver, rt.Type()))
	}

	return receiverType
}

func generateTypeConversionAssignments(
	fromType *astmodel.ObjectType,
	toType *astmodel.ObjectType,
	propertyHandler func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []ast.Stmt) []ast.Stmt {

	var result []ast.Stmt
	for _, toField := range toType.Properties() {
		fieldConversionStmts := propertyHandler(toField, fromType)
		result = append(result, fieldConversionStmts...)
	}

	return result
}

// NewArmTransformerImpl creates a new interface with the specified ARM conversion functions
func NewArmTransformerImpl(
	armTypeName astmodel.TypeName,
	armType *astmodel.ObjectType,
	idFactory astmodel.IdentifierFactory,
	typeType TypeKind) *astmodel.InterfaceImplementation {

	var convertToArmFunc *ArmConversionFunction
	if typeType != StatusType {
		// status type should not have ConvertToARM
		convertToArmFunc = &ArmConversionFunction{
			name:        "ConvertToArm",
			armTypeName: armTypeName,
			armType:     armType,
			idFactory:   idFactory,
			direction:   ConversionDirectionToArm,
			isSpecType:  typeType == SpecType,
		}
	}

	populateFromArmFunc := &ArmConversionFunction{
		name:        "PopulateFromArm",
		armTypeName: armTypeName,
		armType:     armType,
		idFactory:   idFactory,
		direction:   ConversionDirectionFromArm,
		isSpecType:  typeType == SpecType,
	}

	createEmptyArmValueFunc := CreateEmptyArmValueFunc{idFactory: idFactory, armTypeName: armTypeName}

	if convertToArmFunc != nil {
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.MakeGenRuntimePackageReference(), "ArmTransformer"),
			createEmptyArmValueFunc,
			convertToArmFunc,
			populateFromArmFunc)
	} else {
		// only convert in one direction with the FromArmConverter interface
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.MakeGenRuntimePackageReference(), "FromArmConverter"),
			createEmptyArmValueFunc,
			populateFromArmFunc)
	}
}

type complexPropertyConversionParameters struct {
	source            ast.Expr
	destination       ast.Expr
	destinationType   astmodel.Type
	nameHint          string
	conversionContext []astmodel.Type
	assignmentHandler func(destination, source ast.Expr) ast.Stmt

	// sameTypes indicates that the source and destination types are
	// the same, so no conversion between Arm and non-Arm types is
	// required (although structure copying is).
	sameTypes bool
}

func (params complexPropertyConversionParameters) Source() ast.Expr {
	return ast.Clone(params.source).(ast.Expr)
}

func (params complexPropertyConversionParameters) Destination() ast.Expr {
	return ast.Clone(params.destination).(ast.Expr)
}

func (params complexPropertyConversionParameters) copy() complexPropertyConversionParameters {
	result := params
	result.conversionContext = nil
	result.conversionContext = append(result.conversionContext, params.conversionContext...)

	return result
}

func (params complexPropertyConversionParameters) withAdditionalConversionContext(t astmodel.Type) complexPropertyConversionParameters {
	result := params.copy()
	result.conversionContext = append(result.conversionContext, t)

	return result
}

func (params complexPropertyConversionParameters) withSource(source ast.Expr) complexPropertyConversionParameters {
	result := params.copy()
	result.source = source

	return result
}

func (params complexPropertyConversionParameters) withDestination(destination ast.Expr) complexPropertyConversionParameters {
	result := params.copy()
	result.destination = destination

	return result
}

func (params complexPropertyConversionParameters) withDestinationType(t astmodel.Type) complexPropertyConversionParameters {
	result := params.copy()
	result.destinationType = t

	return result
}

func (params complexPropertyConversionParameters) withAssignmentHandler(
	assignmentHandler func(result ast.Expr, destination ast.Expr) ast.Stmt) complexPropertyConversionParameters {
	result := params.copy()
	result.assignmentHandler = assignmentHandler

	return result
}

// countArraysAndMapsInConversionContext returns the number of arrays/maps which are in the conversion context.
// This is to aid in situations where there are deeply nested conversions (i.e. array of map of maps). In these contexts,
// just using a simple assignment such as "elem := ..." isn't sufficient because elem my already have been defined above by
// an enclosing map/array conversion context. We use the depth to do "elem1 := ..." or "elem7 := ...".
func (params complexPropertyConversionParameters) countArraysAndMapsInConversionContext() int {
	result := 0
	for _, t := range params.conversionContext {
		switch t.(type) {
		case *astmodel.MapType:
			result += 1
		case *astmodel.ArrayType:
			result += 1
		}
	}

	return result
}
