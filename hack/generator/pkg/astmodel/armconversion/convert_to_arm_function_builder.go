/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

const nameParameterString = "name"

type convertToArmBuilder struct {
	conversionBuilder
	resultIdent *ast.Ident
}

func newConvertToArmFunctionBuilder(
	c *ArmConversionFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *convertToArmBuilder {

	result := &convertToArmBuilder{
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
		resultIdent: ast.NewIdent("result"),
	}

	result.propertyConversionHandlers = []propertyConversionHandler{
		result.namePropertyHandler,
		result.typePropertyHandler,
		result.propertiesWithSameNameAndTypeHandler,
		result.propertiesWithSameNameButDifferentTypeHandler,
	}

	return result
}

func (builder *convertToArmBuilder) functionDeclaration() *ast.FuncDecl {
	return astbuilder.DefineFunc(
		astbuilder.FuncDetails{
			Name:          ast.NewIdent(builder.methodName),
			ReceiverIdent: builder.receiverIdent,
			ReceiverType: &ast.StarExpr{
				X: builder.receiverTypeExpr,
			},
			Comment: "converts from a Kubernetes CRD object to an ARM object",
			Params: []*ast.Field{
				{
					Type: ast.NewIdent("string"),
					Names: []*ast.Ident{
						ast.NewIdent(nameParameterString),
					},
				},
			},
			Returns: []*ast.Field{
				{
					Type: ast.NewIdent("interface{}"),
				},
				{
					Type: ast.NewIdent("error"),
				},
			},
			Body: builder.functionBodyStatements(),
		})
}

func (builder *convertToArmBuilder) functionBodyStatements() []ast.Stmt {
	var result []ast.Stmt

	// If we are passed a nil receiver just return nil - this is a bit weird
	// but saves us some nil-checks
	result = append(
		result,
		astbuilder.ReturnIfNil(builder.receiverIdent, ast.NewIdent("nil"), ast.NewIdent("nil")))
	result = append(result, astbuilder.NewStruct(builder.resultIdent, builder.armTypeIdent))

	// Each ARM object property needs to be filled out
	result = append(
		result,
		generateTypeConversionAssignments(
			builder.kubeType,
			builder.armType,
			builder.propertyConversionHandler)...)

	returnStatement := &ast.ReturnStmt{
		Results: []ast.Expr{
			builder.resultIdent,
			ast.NewIdent("nil"),
		},
	}
	result = append(result, returnStatement)

	return result
}

//////////////////////
// Conversion handlers
//////////////////////

func (builder *convertToArmBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	_, ok := fromType.Property(GetAzureNameProperty(builder.idFactory).PropertyName())
	if !ok || toProp.PropertyName() != "Name" || !builder.isResource {
		return nil
	}

	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.resultIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		ast.NewIdent(nameParameterString))
	return []ast.Stmt{result}
}

func (builder *convertToArmBuilder) typePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	if toProp.PropertyName() != "Type" || !builder.isResource {
		return nil
	}

	propertyType := toProp.PropertyType()
	if optionalType, ok := toProp.PropertyType().(*astmodel.OptionalType); ok {
		propertyType = optionalType.Element()
	}

	enumTypeName, ok := propertyType.(astmodel.TypeName)
	if !ok {
		panic(fmt.Sprintf("'Type' property was not an enum, was %s", toProp.PropertyType()))
	}

	def, err := builder.codeGenerationContext.GetImportedDefinition(enumTypeName)
	if err != nil {
		panic(err)
	}

	enumType, ok := def.Type().(*astmodel.EnumType)
	if !ok {
		panic(fmt.Sprintf("Enum %v definition was not of type EnumDefinition", enumTypeName))
	}

	optionId := astmodel.GetEnumValueId(def.Name(), enumType.Options()[0])

	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.resultIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		ast.NewIdent(optionId))

	return []ast.Stmt{result}

}

func (builder *convertToArmBuilder) propertiesWithSameNameAndTypeHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())

	if !ok || !toProp.PropertyType().Equals(fromProp.PropertyType()) {
		return nil
	}

	if typeRequiresCopying(fromProp.PropertyType()) {
		// We can't get away with just assigning this field, since
		// it's a reference type. Use the conversion code to copy the
		// elements.
		source := &ast.SelectorExpr{
			X:   builder.receiverIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		}
		destination := &ast.SelectorExpr{
			X:   builder.resultIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		}
		return builder.toArmComplexPropertyConversion(
			complexPropertyConversionParameters{
				source:            source,
				destination:       destination,
				destinationType:   toProp.PropertyType(),
				nameHint:          string(toProp.PropertyName()),
				conversionContext: nil,
				assignmentHandler: assignmentHandlerAssign,
				sameTypes:         true,
			},
		)
	}

	result := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   builder.resultIdent,
			Sel: ast.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		&ast.SelectorExpr{
			X:   builder.receiverIdent,
			Sel: ast.NewIdent(string(fromProp.PropertyName())),
		})

	return []ast.Stmt{result}
}

func (builder *convertToArmBuilder) propertiesWithSameNameButDifferentTypeHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []ast.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok || toProp.PropertyType().Equals(fromProp.PropertyType()) {
		return nil
	}

	destination := &ast.SelectorExpr{
		X:   builder.resultIdent,
		Sel: ast.NewIdent(string(toProp.PropertyName())),
	}
	source := &ast.SelectorExpr{
		X:   builder.receiverIdent,
		Sel: ast.NewIdent(string(fromProp.PropertyName())),
	}

	return builder.toArmComplexPropertyConversion(
		complexPropertyConversionParameters{
			source:            source,
			destination:       destination,
			destinationType:   toProp.PropertyType(),
			nameHint:          string(toProp.PropertyName()),
			conversionContext: nil,
			assignmentHandler: assignmentHandlerAssign,
		})
}

//////////////////////////////////////////////////////////////////////////////////
// Complex property conversion (for when properties aren't simple primitive types)
//////////////////////////////////////////////////////////////////////////////////

func (builder *convertToArmBuilder) toArmComplexPropertyConversion(
	params complexPropertyConversionParameters) []ast.Stmt {

	switch params.destinationType.(type) {
	case *astmodel.OptionalType:
		return builder.convertComplexOptionalProperty(params)
	case *astmodel.ArrayType:
		return builder.convertComplexArrayProperty(params)
	case *astmodel.MapType:
		return builder.convertComplexMapProperty(params)
	case astmodel.TypeName:
		if params.sameTypes {
			if params.destinationType.Equals(astmodel.JSONType) {
				return builder.deepCopyJSON(params)
			}
			// The only type names we leave alone are enums, which
			// don't need conversion.
			return builder.assignPrimitiveType(params)
		}
		return builder.convertComplexTypeNameProperty(params)
	case *astmodel.PrimitiveType:
		// No conversion needed in this case.
		return builder.assignPrimitiveType(params)
	default:
		panic(fmt.Sprintf("don't know how to perform toArm conversion for type: %T", params.destinationType))
	}
}

func assignmentHandlerDefine(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
	return astbuilder.SimpleAssignment(lhs, token.DEFINE, rhs)
}

func assignmentHandlerAssign(lhs ast.Expr, rhs ast.Expr) ast.Stmt {
	return astbuilder.SimpleAssignment(lhs, token.ASSIGN, rhs)
}

// assignPrimitiveType just assigns source to destination directly,
// no conversion needed.
func (builder *convertToArmBuilder) assignPrimitiveType(
	params complexPropertyConversionParameters) []ast.Stmt {
	return []ast.Stmt{
		params.assignmentHandler(params.destination, params.source),
	}
}

// convertComplexOptionalProperty handles conversion for optional properties with complex elements
// This function generates code that looks like this:
// 	if <source> != nil {
//		<code for producing result from destinationType.Element()>
//		<destination> = &<result>
//	}
func (builder *convertToArmBuilder) convertComplexOptionalProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	destinationType := params.destinationType.(*astmodel.OptionalType)

	tempVarIdent := ast.NewIdent(builder.idFactory.CreateIdentifier(params.nameHint+"Typed", astmodel.NotExported))
	tempVarType := destinationType.Element()

	newSource := &ast.UnaryExpr{
		X:  params.source,
		Op: token.MUL,
	}

	innerStatements := builder.toArmComplexPropertyConversion(
		params.withDestination(tempVarIdent).
			withDestinationType(tempVarType).
			withAdditionalConversionContext(destinationType).
			withAssignmentHandler(assignmentHandlerDefine).
			withSource(newSource))

	// Tack on the final assignment
	innerStatements = append(
		innerStatements,
		astbuilder.SimpleAssignment(
			params.destination,
			token.ASSIGN,
			&ast.UnaryExpr{
				Op: token.AND,
				X:  tempVarIdent,
			}))

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
func (builder *convertToArmBuilder) convertComplexArrayProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	var results []ast.Stmt

	destinationType := params.destinationType.(*astmodel.ArrayType)

	depth := params.countArraysAndMapsInConversionContext()
	typedVarIdent := ast.NewIdent("elemTyped")
	tempVarType := destinationType.Element()
	itemIdent := ast.NewIdent("item")
	elemIdent := ast.NewIdent("elem")

	if depth > 0 {
		results = append(results, astbuilder.SimpleVariableDeclaration(
			typedVarIdent,
			destinationType.AsType(builder.codeGenerationContext)))
		typedVarIdent = ast.NewIdent(fmt.Sprintf("elemTyped%d", depth))
	}

	innerStatements := builder.toArmComplexPropertyConversion(
		complexPropertyConversionParameters{
			source:            itemIdent,
			destination:       typedVarIdent,
			destinationType:   tempVarType,
			nameHint:          elemIdent.Name,
			conversionContext: append(params.conversionContext, destinationType),
			assignmentHandler: assignmentHandlerDefine,
			sameTypes:         params.sameTypes,
		})

	// Append the final statement
	innerStatements = append(innerStatements, astbuilder.AppendList(params.destination, typedVarIdent))

	result := &ast.RangeStmt{
		Key:   ast.NewIdent("_"),
		Value: itemIdent,
		X:     params.source,
		Tok:   token.DEFINE,
		Body: &ast.BlockStmt{
			List: innerStatements,
		},
	}

	results = append(results, result)
	return results
}

// convertComplexMapProperty handles conversion for map properties with complex values.
// This function panics if the map keys are not primitive types.
// This function generates code that looks like this:
//	<destination> = make(map[<destinationType.KeyType()]<destinationType.ValueType()>)
// 	if <source> != nil {
//		for key, value := range <source> {
// 			<code for producing result from destinationType.ValueType()>
//			<destination>[key] = <result>
//		}
//	}
func (builder *convertToArmBuilder) convertComplexMapProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	destinationType := params.destinationType.(*astmodel.MapType)

	if _, ok := destinationType.KeyType().(*astmodel.PrimitiveType); !ok {
		panic(fmt.Sprintf("map had non-primitive key type: %v", destinationType.KeyType()))
	}

	keyIdent := ast.NewIdent("key")
	typedVarIdent := ast.NewIdent("elemTyped")
	valueIdent := ast.NewIdent("value")
	elemIdent := ast.NewIdent("elem")

	depth := params.countArraysAndMapsInConversionContext()
	makeMapToken := token.ASSIGN
	if depth > 0 {
		typedVarIdent = ast.NewIdent(fmt.Sprintf("elemTyped%d", depth))
		makeMapToken = token.DEFINE
	}

	innerStatements := builder.toArmComplexPropertyConversion(
		complexPropertyConversionParameters{
			source:            valueIdent,
			destination:       typedVarIdent,
			destinationType:   destinationType.ValueType(),
			nameHint:          elemIdent.Name,
			conversionContext: append(params.conversionContext, destinationType),
			assignmentHandler: assignmentHandlerDefine,
			sameTypes:         params.sameTypes,
		})

	// Append the final statement
	innerStatements = append(innerStatements, astbuilder.InsertMap(params.destination, keyIdent, typedVarIdent))

	keyTypeAst := destinationType.KeyType().AsType(builder.codeGenerationContext)
	valueTypeAst := destinationType.ValueType().AsType(builder.codeGenerationContext)

	makeMapStatement := astbuilder.SimpleAssignment(
		params.destination,
		makeMapToken,
		astbuilder.MakeMap(keyTypeAst, valueTypeAst))
	rangeStatement := &ast.RangeStmt{
		Key:   keyIdent,
		Value: valueIdent,
		X:     params.source,
		Tok:   token.DEFINE,
		Body: &ast.BlockStmt{
			List: innerStatements,
		},
	}

	return []ast.Stmt{makeMapStatement, rangeStatement}
}

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
// 	<nameHint>, err := <source>.ToArm(name)
//	if err != nil {
//		return nil, err
//	}
//	<destination> = <nameHint>.(FooArm)
func (builder *convertToArmBuilder) convertComplexTypeNameProperty(
	params complexPropertyConversionParameters) []ast.Stmt {

	destinationType := params.destinationType.(astmodel.TypeName)

	var results []ast.Stmt
	propertyLocalVarName := ast.NewIdent(builder.idFactory.CreateIdentifier(params.nameHint, astmodel.NotExported))

	// Call ToArm on the property
	results = append(results, callToArmFunction(params.source, propertyLocalVarName, builder.methodName)...)

	typeAssertExpr := &ast.TypeAssertExpr{
		X:    propertyLocalVarName,
		Type: ast.NewIdent(destinationType.Name()),
	}

	if !destinationType.PackageReference.Equals(builder.codeGenerationContext.CurrentPackage()) {
		// needs to be qualified
		packageName, err := builder.codeGenerationContext.GetImportedPackageName(destinationType.PackageReference)
		if err != nil {
			panic(err)
		}

		typeAssertExpr.Type =
			&ast.SelectorExpr{
				X:   ast.NewIdent(packageName),
				Sel: ast.NewIdent(destinationType.Name()),
			}
	}

	results = append(results, params.assignmentHandler(params.destination, typeAssertExpr))

	return results
}

func callToArmFunction(source ast.Expr, destination ast.Expr, methodName string) []ast.Stmt {
	var results []ast.Stmt

	// Call ToArm on the property
	propertyToArmInvocation := &ast.AssignStmt{
		Lhs: []ast.Expr{
			destination,
			ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   source,
					Sel: ast.NewIdent(methodName),
				},
				Args: []ast.Expr{
					ast.NewIdent(nameParameterString),
				},
			},
		},
	}
	results = append(results, propertyToArmInvocation)
	results = append(results, astbuilder.CheckErrorAndReturn(ast.NewIdent("nil")))

	return results
}

func typeRequiresCopying(theType astmodel.Type) bool {
	switch t := theType.(type) {
	case *astmodel.OptionalType, *astmodel.MapType, *astmodel.ArrayType:
		return true
	case astmodel.TypeName:
		return t == astmodel.JSONType
	}
	return false
}
