/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

const nameParameterString = "name"
const resolvedReferencesParameterString = "resolvedReferences"

type convertToARMBuilder struct {
	conversionBuilder
	resultIdent           string
	typeConversionBuilder *astmodel.ConversionFunctionBuilder
}

func newConvertToARMFunctionBuilder(
	c *ARMConversionFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *convertToARMBuilder {

	result := &convertToARMBuilder{
		conversionBuilder: conversionBuilder{
			methodName:            methodName,
			armType:               c.armType,
			kubeType:              getReceiverObjectType(codeGenerationContext, receiver),
			receiverIdent:         c.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported),
			receiverTypeExpr:      receiver.AsType(codeGenerationContext),
			armTypeIdent:          c.armTypeName.Name(),
			idFactory:             c.idFactory,
			isSpecType:            c.isSpecType,
			codeGenerationContext: codeGenerationContext,
		},
		resultIdent:           "result",
		typeConversionBuilder: astmodel.NewConversionFunctionBuilder(c.idFactory, codeGenerationContext),
	}

	// It's a bit awkward that there are two levels of "handler" here, but they serve different purposes:
	// The top level propertyConversionHandlers is about determining which properties are involved: given a property on the destination type it
	// determines which property (if any) on the source type will be converted to the destination.
	// The "inner" handler (typeConversionBuilder) is about determining how to convert between two types: given a
	// source type and a destination type, figure out how to make the assignment work. It has no knowledge of broader object strucutre
	// or other properties.
	result.typeConversionBuilder.AddConversionHandlers(result.convertComplexTypeNameProperty)
	result.propertyConversionHandlers = []propertyConversionHandler{
		result.namePropertyHandler,
		result.fixedValuePropertyHandler("Type"),
		result.fixedValuePropertyHandler("ApiVersion"),
		result.propertiesWithSameNameHandler,
	}

	return result
}

func (builder *convertToARMBuilder) functionDeclaration() *dst.FuncDecl {
	fn := &astbuilder.FuncDetails{
		Name:          builder.methodName,
		ReceiverIdent: builder.receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: builder.receiverTypeExpr,
		},
		Body: builder.functionBodyStatements(),
	}

	fn.AddParameter(nameParameterString, dst.NewIdent("string"))
	fn.AddParameter(
		resolvedReferencesParameterString,
		astbuilder.Selector(dst.NewIdent(astmodel.GenRuntimePackageName), "ResolvedReferences"))
	fn.AddReturns("interface{}", "error")
	fn.AddComments("converts from a Kubernetes CRD object to an ARM object")

	return fn.DefineFunc()
}

func (builder *convertToARMBuilder) functionBodyStatements() []dst.Stmt {
	var result []dst.Stmt

	// If we are passed a nil receiver just return nil - this is a bit weird
	// but saves us some nil-checks
	result = append(
		result,
		astbuilder.ReturnIfNil(dst.NewIdent(builder.receiverIdent), dst.NewIdent("nil"), dst.NewIdent("nil")))
	result = append(result, astbuilder.NewVariable(builder.resultIdent, builder.armTypeIdent))

	// Each ARM object property needs to be filled out
	result = append(
		result,
		generateTypeConversionAssignments(
			builder.kubeType,
			builder.armType,
			builder.propertyConversionHandler)...)

	returnStatement := &dst.ReturnStmt{
		Results: []dst.Expr{
			dst.NewIdent(builder.resultIdent),
			dst.NewIdent("nil"),
		},
	}
	result = append(result, returnStatement)

	return result
}

//////////////////////
// Conversion handlers
//////////////////////

func (builder *convertToARMBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	if toProp.PropertyName() != "Name" || !builder.isSpecType {
		return nil
	}

	// we do not read from AzureName() but instead use
	// the passed-in 'name' parameter which contains
	// a full name including any owners, etc
	result := astbuilder.SimpleAssignment(
		astbuilder.Selector(dst.NewIdent(builder.resultIdent), string(toProp.PropertyName())),
		token.ASSIGN,
		dst.NewIdent(nameParameterString))

	return []dst.Stmt{result}
}

func (builder *convertToARMBuilder) fixedValuePropertyHandler(propertyName astmodel.PropertyName) propertyConversionHandler {
	return func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []dst.Stmt {
		if toProp.PropertyName() != propertyName || !builder.isSpecType {
			return nil
		}

		propertyType := toProp.PropertyType()
		if optionalType, ok := toProp.PropertyType().(*astmodel.OptionalType); ok {
			propertyType = optionalType.Element()
		}

		enumTypeName, ok := propertyType.(astmodel.TypeName)
		if !ok {
			panic(fmt.Sprintf("'%s' property was not an enum, was %s", propertyName, toProp.PropertyType()))
		}

		def, err := builder.codeGenerationContext.GetImportedDefinition(enumTypeName)
		if err != nil {
			panic(err)
		}

		enumType, ok := def.Type().(*astmodel.EnumType)
		if !ok {
			panic(fmt.Sprintf("Enum %v definition was not of type EnumDefinition", enumTypeName))
		}

		optionId := astmodel.GetEnumValueId(def.Name().Name(), enumType.Options()[0])

		result := astbuilder.SimpleAssignment(
			astbuilder.Selector(dst.NewIdent(builder.resultIdent), string(toProp.PropertyName())),
			token.ASSIGN,
			dst.NewIdent(optionId))

		return []dst.Stmt{result}
	}
}

func (builder *convertToARMBuilder) propertiesWithSameNameHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return nil
	}

	source := astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(fromProp.PropertyName()))
	destination := astbuilder.Selector(dst.NewIdent(builder.resultIdent), string(toProp.PropertyName()))

	return builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            source,
			SourceType:        fromProp.PropertyType(),
			Destination:       destination,
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(toProp.PropertyName()),
			ConversionContext: nil,
		},
	)
}

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
// 	<nameHint>, err := <source>.ToARM(name)
//	if err != nil {
//		return nil, err
//	}
//	<destination> = <nameHint>.(FooARM)
func (builder *convertToARMBuilder) convertComplexTypeNameProperty(conversionBuilder *astmodel.ConversionFunctionBuilder, params astmodel.ConversionParameters) []dst.Stmt {
	destinationType, ok := params.DestinationType.(astmodel.TypeName)
	if !ok {
		return nil
	}

	sourceType, ok := params.SourceType.(astmodel.TypeName)
	if !ok {
		return nil
	}

	// This is for handling type names that aren't equal
	if sourceType.Equals(destinationType) {
		return nil
	}

	var results []dst.Stmt
	propertyLocalVarName := conversionBuilder.IdFactory.CreateIdentifier(params.NameHint, astmodel.NotExported)

	// Call ToARM on the property
	results = append(results, callToARMFunction(params.GetSource(), dst.NewIdent(propertyLocalVarName), builder.methodName)...)

	typeAssertExpr := &dst.TypeAssertExpr{
		X:    dst.NewIdent(propertyLocalVarName),
		Type: dst.NewIdent(destinationType.Name()),
	}

	if !destinationType.PackageReference.Equals(conversionBuilder.CodeGenerationContext.CurrentPackage()) {
		// needs to be qualified
		packageName, err := conversionBuilder.CodeGenerationContext.GetImportedPackageName(destinationType.PackageReference)
		if err != nil {
			panic(err)
		}

		typeAssertExpr.Type = astbuilder.Selector(dst.NewIdent(packageName), destinationType.Name())
	}

	results = append(results, params.AssignmentHandlerOrDefault()(params.GetDestination(), typeAssertExpr))

	return results
}

func callToARMFunction(source dst.Expr, destination dst.Expr, methodName string) []dst.Stmt {
	var results []dst.Stmt

	// Call ToARM on the property
	propertyToARMInvocation := &dst.AssignStmt{
		Lhs: []dst.Expr{
			destination,
			dst.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: astbuilder.Selector(source, methodName),
				Args: []dst.Expr{
					dst.NewIdent(nameParameterString),
					dst.NewIdent(resolvedReferencesParameterString),
				},
			},
		},
	}
	results = append(results, propertyToARMInvocation)
	results = append(results, astbuilder.CheckErrorAndReturn(dst.NewIdent("nil")))

	return results
}
