/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

type conversionBuilder struct {
	receiverIdent              string
	receiverTypeExpr           dst.Expr
	armTypeIdent               string
	codeGenerationContext      *astmodel.CodeGenerationContext
	idFactory                  astmodel.IdentifierFactory
	typeKind                   TypeKind
	methodName                 string
	kubeType                   *astmodel.ObjectType
	armType                    *astmodel.ObjectType
	propertyConversionHandlers []propertyConversionHandler
}

type TypeKind int

const (
	TypeKindOrdinary TypeKind = iota
	TypeKindSpec
	TypeKindStatus
)

func (builder conversionBuilder) propertyConversionHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) ([]dst.Stmt, error) {
	var err error
	for _, conversionHandler := range builder.propertyConversionHandlers {
		var conversion propertyConversionHandlerResult
		conversion, err = conversionHandler(toProp, fromType)
		if err != nil {
			break
		}

		if conversion.matched {
			return conversion.statements, nil
		}
	}

	var kubeDescription strings.Builder
	builder.kubeType.WriteDebugDescription(&kubeDescription, nil)

	var armDescription strings.Builder
	builder.armType.WriteDebugDescription(&armDescription, nil)

	message := fmt.Sprintf(
		"no property found for %q in method %s()\nFrom: %s\nTo: %s",
		toProp.PropertyName(),
		builder.methodName,
		kubeDescription.String(),
		armDescription.String())

	if err != nil {
		return nil, errors.Wrap(err, message)
	}

	return nil, errors.New(message)
}

type propertyConversionHandler = func(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error)

type propertyConversionHandlerResult struct {
	statements []dst.Stmt
	matched    bool
}

// notHandled is a result to use when a handler declines to handle the requested conversion
var notHandled = propertyConversionHandlerResult{
	matched: false,
}

// handledWithNoOp is a result to use when a handler wants to return an empty set of statements for a conversion
var handledWithNoOp = propertyConversionHandlerResult{
	matched: true,
}

// handleWith is a result to use when a handler wants to return a set of statements for a conversion
func handleWith(statements ...any) propertyConversionHandlerResult {
	return propertyConversionHandlerResult{
		statements: astbuilder.Statements(statements...),
		matched:    true,
	}
}

var (
	once              sync.Once
	azureNameProperty *astmodel.PropertyDefinition
)

func initializeAzureName(idFactory astmodel.IdentifierFactory) {
	azureNameFieldDescription := "The name of the resource in Azure. This is often the same as" +
		" the name of the resource in Kubernetes but it doesn't have to be."
	azureNameProperty = astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(astmodel.AzureNameProperty, astmodel.Exported),
		idFactory.CreateStringIdentifier(astmodel.AzureNameProperty, astmodel.NotExported),
		astmodel.StringType).WithDescription(azureNameFieldDescription)
}

// GetAzureNameProperty returns the special "AzureName" field
func GetAzureNameProperty(idFactory astmodel.IdentifierFactory) *astmodel.PropertyDefinition {
	once.Do(func() { initializeAzureName(idFactory) })

	return azureNameProperty
}

func getReceiverObjectType(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *astmodel.ObjectType {
	// Determine the type we're operating on
	rt := codeGenerationContext.MustGetDefinition(receiver)
	receiverType, ok := rt.Type().(*astmodel.ObjectType)
	if !ok {
		// Don't expect to have any wrapper types left at this point
		panic(fmt.Sprintf("receiver for ARMConversionFunction is not of expected type. TypeName: %s, Type %T", receiver, rt.Type()))
	}

	return receiverType
}

func generateTypeConversionAssignments(
	fromType *astmodel.ObjectType,
	toType *astmodel.ObjectType,
	propertyHandler func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) ([]dst.Stmt, error),
) ([]dst.Stmt, error) {
	var result []dst.Stmt
	var errs []error
	for _, toField := range toType.Properties().AsSlice() {
		fieldConversionStmts, err := propertyHandler(toField, fromType)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if len(fieldConversionStmts) > 0 {
			result = append(result, &dst.EmptyStmt{
				Decs: dst.EmptyStmtDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.EmptyLine,
						End:    []string{fmt.Sprintf("// Set property %q:", toField.PropertyName())},
					},
				},
			})
			result = append(result, fieldConversionStmts...)
		} else {
			result = append(result, &dst.EmptyStmt{
				Decs: dst.EmptyStmtDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.EmptyLine,
						End:    []string{fmt.Sprintf("// no assignment for property %q", toField.PropertyName())},
					},
				},
			})
		}
	}

	return result, kerrors.NewAggregate(errs)
}

// NewARMConversionImplementation creates an interface implementation with the specified ARM conversion functions
func NewARMConversionImplementation(
	armTypeName astmodel.TypeName,
	armType *astmodel.ObjectType,
	idFactory astmodel.IdentifierFactory,
	typeKind TypeKind,
) *astmodel.InterfaceImplementation {
	var convertToARMFunc *ConvertToARMFunction
	if typeKind != TypeKindStatus {
		// status type should not have ConvertToARM
		convertToARMFunc = &ConvertToARMFunction{
			ARMConversionFunction: ARMConversionFunction{
				armTypeName: armTypeName,
				armType:     armType,
				idFactory:   idFactory,
				typeKind:    typeKind,
			},
		}
	}

	populateFromARMFunc := &PopulateFromARMFunction{
		ARMConversionFunction: ARMConversionFunction{
			armTypeName: armTypeName,
			armType:     armType,
			idFactory:   idFactory,
			typeKind:    typeKind,
		},
	}

	newEmptyARMValueFunc := functions.NewNewEmptyARMValueFunc(armTypeName, idFactory)

	if convertToARMFunc != nil {
		// can convert both to and from ARM = the ARMTransformer interface
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.GenRuntimeReference, "ARMTransformer"),
			newEmptyARMValueFunc,
			convertToARMFunc,
			populateFromARMFunc)
	} else {
		// can only convert in one direction with the FromARMConverter interface
		return astmodel.NewInterfaceImplementation(
			astmodel.MakeTypeName(astmodel.GenRuntimeReference, "FromARMConverter"),
			newEmptyARMValueFunc,
			populateFromARMFunc)
	}
}

func removeEmptyStatements(stmts []dst.Stmt) []dst.Stmt {
	result := make([]dst.Stmt, 0, len(stmts))
	for _, stmt := range stmts {
		if _, ok := stmt.(*dst.EmptyStmt); ok {
			continue
		}
		result = append(result, stmt)
	}

	return result
}

const (
	ConversionTag        = "conversion"
	NoARMConversionValue = "noarmconversion"
)

func skipNonConvertablePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) ([]dst.Stmt, bool) {
	// If the property has been flagged as not being convertible, skip it
	if toProp.HasTagValue(ConversionTag, NoARMConversionValue) {
		return nil, true
	}

	return nil, false
}
