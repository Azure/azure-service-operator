/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"sync"

	"github.com/dave/dst"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

type conversionBuilder struct {
	receiverIdent              string
	receiverTypeExpr           dst.Expr
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
	fromType *astmodel.ObjectType) []dst.Stmt {

	for _, conversionHandler := range builder.propertyConversionHandlers {
		stmts := conversionHandler(toProp, fromType)
		if len(stmts) > 0 {
			return stmts
		}
	}

	panic(fmt.Sprintf("No property found for %s in method %s\nFrom: %+v\nTo: %+v", toProp.PropertyName(), builder.methodName, *builder.kubeType, *builder.armType))
}

type propertyConversionHandler = func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []dst.Stmt

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
		panic(fmt.Sprintf("receiver for ARMConversionFunction is not of expected type. TypeName: %v, Type %T", receiver, rt.Type()))
	}

	return receiverType
}

func generateTypeConversionAssignments(
	fromType *astmodel.ObjectType,
	toType *astmodel.ObjectType,
	propertyHandler func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []dst.Stmt) []dst.Stmt {

	var result []dst.Stmt
	for _, toField := range toType.Properties() {
		fieldConversionStmts := propertyHandler(toField, fromType)
		result = append(result, fieldConversionStmts...)
	}

	return result
}

// NewARMTransformerImpl creates a new interface with the specified ARM conversion functions
func NewARMTransformerImpl(
	armTypeName astmodel.TypeName,
	armType *astmodel.ObjectType,
	idFactory astmodel.IdentifierFactory,
	typeKind TypeKind) *astmodel.InterfaceImplementation {

	var convertToARMFunc *ConvertToARMFunction
	if typeKind != StatusType {
		// status type should not have ConvertToARM
		convertToARMFunc = &ConvertToARMFunction{
			ARMConversionFunction: ARMConversionFunction{
				armTypeName: armTypeName,
				armType:     armType,
				idFactory:   idFactory,
				isSpecType:  typeKind == SpecType,
			},
		}
	}

	populateFromARMFunc := &PopulateFromARMFunction{
		ARMConversionFunction: ARMConversionFunction{
			armTypeName: armTypeName,
			armType:     armType,
			idFactory:   idFactory,
			isSpecType:  typeKind == SpecType,
		},
	}

	createEmptyARMValueFunc := CreateEmptyARMValueFunc{idFactory: idFactory, armTypeName: armTypeName}

	if convertToARMFunc != nil {
		// can convert both to and from ARM = the ARMTransformer interface
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.GenRuntimeReference, "ARMTransformer"),
			createEmptyARMValueFunc,
			convertToARMFunc,
			populateFromARMFunc)
	} else {
		// only convert in one direction with the FromARMConverter interface
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.GenRuntimeReference, "FromARMConverter"),
			createEmptyARMValueFunc,
			populateFromARMFunc)
	}
}
