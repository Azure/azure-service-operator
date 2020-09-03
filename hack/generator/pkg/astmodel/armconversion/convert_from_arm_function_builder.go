/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"go/ast"
	"go/token"
)

type convertFromArmBuilder struct {
	conversionBuilder
	typedInputIdent *ast.Ident
	inputIdent      *ast.Ident
}

func newConvertFromArmFunctionBuilder(
	c *ArmConversionFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *convertFromArmBuilder {

	result := &convertFromArmBuilder{
		// Note: If we have a property with these names we will have a compilation issue in the generated
		// code. Right now that doesn't seem to be the case anywhere but if it does happen we may need
		// to harden this logic some to choose an unused name.
		typedInputIdent: ast.NewIdent("typedInput"),
		inputIdent:      ast.NewIdent("armInput"),

		conversionBuilder: conversionBuilder{
			methodName:            methodName,
			armType:               c.armType,
			kubeType:              getReceiverObjectType(codeGenerationContext, receiver),
			receiverIdent:         ast.NewIdent(c.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)),
			receiverTypeExpr:      receiver.AsType(codeGenerationContext),
			armTypeIdent:          ast.NewIdent(c.armTypeName.Name()),
			idFactory:             c.idFactory,
			isResource:            c.isResource,
			codeGenerationContext: codeGenerationContext,
		},
	}

	result.propertyConversionHandlers = []propertyConversionHandler{
		result.namePropertyHandler,
		result.ownerPropertyHandler,
		result.propertiesWithSameNameAndTypeHandler,
		result.propertiesWithSameNameButDifferentTypeHandler(),
	}

	return result
}

func (builder *convertFromArmBuilder) functionDeclaration() *ast.FuncDecl {

	return astbuilder.DefineFunc(
		astbuilder.FuncDetails{
			Name:          ast.NewIdent(builder.methodName),
			ReceiverIdent: builder.receiverIdent,
			ReceiverType: &ast.StarExpr{
				X: builder.receiverTypeExpr,
			},
			Comment: "populates a Kubernetes CRD object from an Azure ARM object",
			Params: []*ast.Field{
				{
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent(astmodel.GenRuntimePackageName),
						Sel: ast.NewIdent("KnownResourceReference"),
					},
					Names: []*ast.Ident{
						ast.NewIdent(builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)),
					},
				},
				{
					Type: ast.NewIdent("interface{}"),
					Names: []*ast.Ident{
						builder.inputIdent,
					},
				},
			},
			Returns: []*ast.Field{
				{
					Type: ast.NewIdent("error"),
				},
			},
			Body: builder.functionBodyStatements(),
		})
}

func (builder *convertFromArmBuilder) functionBodyStatements() []ast.Stmt {
	var result []ast.Stmt

	// perform a type assert and check its results
	result = append(result, builder.assertInputTypeIsArm()...)

	// Do all of the assignments for each property
	result = append(
		result,
		generateTypeConversionAssignments(
			builder.armType,
			builder.kubeType,
			builder.propertyConversionHandler)...)

	// Return nil error if we make it to the end
	result = append(
		result,
		&ast.ReturnStmt{
			Results: []ast.Expr{
				ast.NewIdent("nil"),
			},
		})

	return result
}

func (builder *convertFromArmBuilder) assertInputTypeIsArm() []ast.Stmt {
	var result []ast.Stmt

	// perform a type assert
	result = append(
		result,
		astbuilder.TypeAssert(builder.typedInputIdent, builder.inputIdent, builder.armTypeIdent))

	// Check the result of the type assert
	result = append(
		result,
		astbuilder.ReturnIfNotOk(
			astbuilder.FormatError(
				fmt.Sprintf("\"unexpected type supplied for %s() function. Expected %s, got %%T\"",
					builder.methodName,
					builder.armTypeIdent.Name),
				builder.inputIdent)))

	return result
}

//////////////////////
// Conversion handlers
//////////////////////

func (builder *convertFromArmBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	if !toProp.Equals(GetAzureNameProperty(builder.idFactory)) || !builder.isResource {
		return nil
	}

	// Check to make sure that the ARM object has a "Name" property (which matches our "AzureName")
	fromProp, ok := fromType.Property(astmodel.PropertyName("Name"))
	if !ok {
		panic("Arm resource missing property 'Name'")
	}
	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.receiverIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		&ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent(astmodel.GenRuntimePackageName),
				Sel: ast.NewIdent("ExtractKubernetesResourceNameFromArmName"),
			},
			Args: []ast.Expr{
				&ast.SelectorExpr{
					X:   builder.typedInputIdent,
					Sel: ast.NewIdent(string(fromProp.PropertyName())),
				},
			},
		})

	return []ast.Stmt{result}

}

func (builder *convertFromArmBuilder) ownerPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType) []ast.Stmt {

	if toProp.PropertyName() != builder.idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported) || !builder.isResource {
		return nil
	}

	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.receiverIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		ast.NewIdent(builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)))
	return []ast.Stmt{result}
}

func (builder *convertFromArmBuilder) propertiesWithSameNameAndTypeHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())

	if !ok || !toProp.PropertyType().Equals(fromProp.PropertyType()) {
		return nil
	}

	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.receiverIdent,
			Sel: ast.NewIdent(string(fromProp.PropertyName())),
		},
		token.ASSIGN,
		&ast.SelectorExpr{
			X:   builder.typedInputIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		})
	return []ast.Stmt{result}
}

func (builder *convertFromArmBuilder) propertiesWithSameNameButDifferentTypeHandler() propertyConversionHandler {
	definedErrVar := false

	return func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []ast.Stmt {
		fromProp, ok := fromType.Property(toProp.PropertyName())

		if !ok || toProp.PropertyType().Equals(fromProp.PropertyType()) {
			return nil
		}

		var result []ast.Stmt

		if !definedErrVar {
			result = append(
				result,
				astbuilder.SimpleVariableDeclaration(ast.NewIdent("err"), ast.NewIdent("error")))
			definedErrVar = true
		}

		complexConversion := builder.fromArmComplexPropertyConversion(
			complexPropertyConversionParameters{
				source: &ast.SelectorExpr{
					X:   builder.typedInputIdent,
					Sel: ast.NewIdent(string(fromProp.PropertyName())),
				},
				destination: &ast.SelectorExpr{
					X:   builder.receiverIdent,
					Sel: ast.NewIdent(string(toProp.PropertyName())),
				},
				destinationType:   toProp.PropertyType(),
				nameHint:          string(toProp.PropertyName()),
				conversionContext: nil,
				assignmentHandler: nil,
			})

		result = append(result, complexConversion...)
		return result
	}
}

//////////////////////////////////////////////////////////////////////////////////
// Complex property conversion (for when properties aren't simple primitive types)
//////////////////////////////////////////////////////////////////////////////////

func (builder *convertFromArmBuilder) fromArmComplexPropertyConversion(
	params complexPropertyConversionParameters) []ast.Stmt {

	switch params.destinationType.(type) {
	case *astmodel.OptionalType:
		return builder.convertComplexOptionalProperty(params)
	case *astmodel.ArrayType:
		return builder.convertComplexArrayProperty(params)
	case *astmodel.MapType:
		return builder.convertComplexMapProperty(params)
	case astmodel.TypeName:
		return builder.convertComplexTypeNameProperty(params)
	default:
		panic(fmt.Sprintf("don't know how to perform fromArm conversion for type: %T", params.destinationType))
	}
}

// convertComplexOptionalProperty handles conversion for optional properties with complex elements
// This function generates code that looks like this:
// 	if <source> != nil {
//		<code for producing result from destinationType.Element()>
//		<destination> = &<result>
//	}
func (builder *convertFromArmBuilder) convertComplexOptionalProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	handler := func(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
		return astbuilder.SimpleAssignment(
			lhs,
			token.ASSIGN,
			&ast.UnaryExpr{
				Op: token.AND,
				X:  rhs,
			})
	}

	destinationType := params.destinationType.(*astmodel.OptionalType)

	newSource := &ast.UnaryExpr{
		X:  params.source,
		Op: token.MUL,
	}

	innerStatements := builder.fromArmComplexPropertyConversion(
		params.withDestinationType(destinationType.Element()).
			withAdditionalConversionContext(destinationType).
			withAssignmentHandler(handler).
			withSource(newSource))

	result := &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  params.source,
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: innerStatements,
		},
	}

	return []ast.Stmt{result}
}

// convertComplexArrayProperty handles conversion for array properties with complex elements
// This function generates code that looks like this:
// 	for _, item := range <source> {
//		<code for producing result from destinationType.Element()>
//		<destination> = append(<destination>, <result>)
//	}
func (builder *convertFromArmBuilder) convertComplexArrayProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	var results []ast.Stmt

	itemIdent := ast.NewIdent("item")
	elemIdent := ast.NewIdent("elem")

	depth := params.countArraysAndMapsInConversionContext()

	destinationType := params.destinationType.(*astmodel.ArrayType)

	elemType := destinationType.Element()
	actualDestination := params.destination // TODO: improve name
	if depth > 0 {
		actualDestination = elemIdent
		results = append(
			results,
			astbuilder.SimpleVariableDeclaration(
				elemIdent,
				destinationType.AsType(builder.codeGenerationContext)))
		elemIdent = ast.NewIdent(fmt.Sprintf("elem%d", depth))
	}

	result := &ast.RangeStmt{
		Key:   ast.NewIdent("_"),
		Value: itemIdent,
		X:     params.source,
		Tok:   token.DEFINE,
		Body: &ast.BlockStmt{
			List: builder.fromArmComplexPropertyConversion(
				complexPropertyConversionParameters{
					source:            itemIdent,
					destination:       actualDestination,
					destinationType:   elemType,
					nameHint:          elemIdent.Name,
					conversionContext: append(params.conversionContext, destinationType),
					assignmentHandler: astbuilder.AppendList,
				}),
		},
	}
	results = append(results, result)

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.assignmentHandler != nil {
		results = append(results, params.assignmentHandler(params.destination, actualDestination))
	}

	return results
}

// convertComplexMapProperty handles conversion for map properties with complex values.
// This function panics if the map keys are not primitive types.
// This function generates code that looks like this:
// 	if <source> != nil {
//		<destination> = make(map[<destinationType.KeyType()]<destinationType.ValueType()>)
//		for key, value := range <source> {
// 			<code for producing result from destinationType.ValueType()>
//			<destination>[key] = <result>
//		}
//	}
func (builder *convertFromArmBuilder) convertComplexMapProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	destinationType := params.destinationType.(*astmodel.MapType)

	if _, ok := destinationType.KeyType().(*astmodel.PrimitiveType); !ok {
		panic(fmt.Sprintf("map had non-primitive key type: %v", destinationType.KeyType()))
	}

	depth := params.countArraysAndMapsInConversionContext()

	keyIdent := ast.NewIdent("key")
	valueIdent := ast.NewIdent("value")
	elemIdent := ast.NewIdent("elem")

	actualDestination := params.destination // TODO: improve name
	makeMapToken := token.ASSIGN
	if depth > 0 {
		actualDestination = elemIdent
		elemIdent = ast.NewIdent(fmt.Sprintf("elem%d", depth))
		makeMapToken = token.DEFINE
	}

	handler := func(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
		return astbuilder.InsertMap(lhs, keyIdent, rhs)
	}

	keyTypeAst := destinationType.KeyType().AsType(builder.codeGenerationContext)
	valueTypeAst := destinationType.ValueType().AsType(builder.codeGenerationContext)

	makeMapStatement := astbuilder.SimpleAssignment(
		actualDestination,
		makeMapToken,
		astbuilder.MakeMap(keyTypeAst, valueTypeAst))
	rangeStatement := &ast.RangeStmt{
		Key:   keyIdent,
		Value: valueIdent,
		X:     params.source,
		Tok:   token.DEFINE,
		Body: &ast.BlockStmt{
			List: builder.fromArmComplexPropertyConversion(
				complexPropertyConversionParameters{
					source:            valueIdent,
					destination:       actualDestination,
					destinationType:   destinationType.ValueType(),
					nameHint:          elemIdent.Name,
					conversionContext: append(params.conversionContext, destinationType),
					assignmentHandler: handler,
				}),
		},
	}

	result := &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  params.source,
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				makeMapStatement,
				rangeStatement,
			},
		},
	}

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.assignmentHandler != nil {
		result.Body.List = append(result.Body.List, params.assignmentHandler(params.destination, actualDestination))
	}

	return []ast.Stmt{result}
}

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
//	<nameHint> := <destinationType>{}
//	err = <nameHint>.FromArm(owner, <source>)
//	if err != nil {
//		return err
//	}
//	<destination> = <nameHint>
func (builder *convertFromArmBuilder) convertComplexTypeNameProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	var results []ast.Stmt

	destinationType := params.destinationType.(astmodel.TypeName)

	propertyLocalVarName := ast.NewIdent(builder.idFactory.CreateIdentifier(params.nameHint, astmodel.NotExported))

	results = append(results, astbuilder.NewStruct(propertyLocalVarName, ast.NewIdent(destinationType.Name())))
	results = append(
		results,
		astbuilder.SimpleAssignment(
			ast.NewIdent("err"),
			token.ASSIGN,
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   propertyLocalVarName,
					Sel: ast.NewIdent(builder.methodName),
				},
				Args: []ast.Expr{
					ast.NewIdent(builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)),
					params.source,
				},
			}))
	results = append(results, astbuilder.CheckErrorAndReturn())
	if params.assignmentHandler == nil {
		results = append(
			results,
			astbuilder.SimpleAssignment(
				params.destination,
				token.ASSIGN,
				propertyLocalVarName))
	} else {
		results = append(
			results,
			params.assignmentHandler(params.destination, propertyLocalVarName))
	}

	return results
}
