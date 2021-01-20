/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
)

type convertFromArmBuilder struct {
	conversionBuilder
	typedInputIdent string
	inputIdent      string
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
		typedInputIdent: "typedInput",
		inputIdent:      "armInput",

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
	}

	result.propertyConversionHandlers = []propertyConversionHandler{
		result.namePropertyHandler,
		result.ownerPropertyHandler,
		result.propertiesWithSameNameAndTypeHandler,
		result.propertiesWithSameNameButDifferentTypeHandler(),
	}

	return result
}

func (builder *convertFromArmBuilder) functionDeclaration() *dst.FuncDecl {

	fn := &astbuilder.FuncDetails{
		Name:          builder.methodName,
		ReceiverIdent: builder.receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: builder.receiverTypeExpr,
		},
		Body: builder.functionBodyStatements(),
	}

	fn.AddComments("populates a Kubernetes CRD object from an Azure ARM object")
	fn.AddParameter(
		builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported),
		&dst.SelectorExpr{
			X:   dst.NewIdent(astmodel.GenRuntimePackageName),
			Sel: dst.NewIdent("KnownResourceReference"),
		})

	fn.AddParameter(builder.inputIdent, dst.NewIdent("interface{}"))
	fn.AddReturns("error")
	return fn.DefineFunc()
}

func (builder *convertFromArmBuilder) functionBodyStatements() []dst.Stmt {
	var result []dst.Stmt

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
		&dst.ReturnStmt{
			Results: []dst.Expr{
				dst.NewIdent("nil"),
			},
		})

	return result
}

func (builder *convertFromArmBuilder) assertInputTypeIsArm() []dst.Stmt {
	var result []dst.Stmt

	fmtPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.FmtReference)

	// perform a type assert
	result = append(
		result,
		astbuilder.TypeAssert(
			dst.NewIdent(builder.typedInputIdent),
			dst.NewIdent(builder.inputIdent),
			dst.NewIdent(builder.armTypeIdent)))

	// Check the result of the type assert
	result = append(
		result,
		astbuilder.ReturnIfNotOk(
			astbuilder.FormatError(
				fmtPackage,
				fmt.Sprintf("unexpected type supplied for %s() function. Expected %s, got %%T",
					builder.methodName,
					builder.armTypeIdent),
				dst.NewIdent(builder.inputIdent))))

	return result
}

//////////////////////
// Conversion handlers
//////////////////////

func (builder *convertFromArmBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	if !builder.isSpecType || !toProp.HasName(astmodel.AzureNameProperty) {
		return nil
	}

	// Check to make sure that the ARM object has a "Name" property (which matches our "AzureName")
	fromProp, ok := fromType.Property("Name")
	if !ok {
		panic("ARM resource missing property 'Name'")
	}

	// Invoke SetAzureName(ExtractKubernetesResourceNameFromArmName(this.Name)):
	return []dst.Stmt{
		&dst.ExprStmt{
			X: astbuilder.CallQualifiedFunc(
				builder.receiverIdent,
				"SetAzureName",
				astbuilder.CallQualifiedFunc(
					astmodel.GenRuntimePackageName,
					"ExtractKubernetesResourceNameFromArmName",
					&dst.SelectorExpr{
						X:   dst.NewIdent(builder.typedInputIdent),
						Sel: dst.NewIdent(string(fromProp.PropertyName())),
					}),
			),
		},
	}
}

func (builder *convertFromArmBuilder) ownerPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType) []dst.Stmt {

	if toProp.PropertyName() != builder.idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported) || !builder.isSpecType {
		return nil
	}

	result := astbuilder.SimpleAssignment(
		&dst.SelectorExpr{
			X:   dst.NewIdent(builder.receiverIdent),
			Sel: dst.NewIdent(string(toProp.PropertyName())),
		},
		token.ASSIGN,
		dst.NewIdent(builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)))
	return []dst.Stmt{result}
}

func (builder *convertFromArmBuilder) propertiesWithSameNameAndTypeHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return nil
	}

	// check that we are assigning to the same type or a validated
	// version of the same type
	toType := toProp.PropertyType()
	if toValidated, ok := toType.(astmodel.ValidatedType); ok {
		toType = toValidated.ElementType()
	}

	if !toType.Equals(fromProp.PropertyType()) {
		return nil
	}

	if typeRequiresCopying(toType) {
		// We can't get away with just assigning this field, since
		// it's a reference type. Use the conversion code to copy the
		// elements.

		return builder.fromArmComplexPropertyConversion(
			complexPropertyConversionParameters{
				source: &dst.SelectorExpr{
					X:   dst.NewIdent(builder.typedInputIdent),
					Sel: dst.NewIdent(string(toProp.PropertyName())),
				},
				destination: &dst.SelectorExpr{
					X:   dst.NewIdent(builder.receiverIdent),
					Sel: dst.NewIdent(string(toProp.PropertyName())),
				},
				destinationType:   toType,
				nameHint:          string(toProp.PropertyName()),
				conversionContext: nil,
				assignmentHandler: nil,
				sameTypes:         true,
			},
		)
	}

	result := astbuilder.SimpleAssignment(
		&dst.SelectorExpr{
			X:   dst.NewIdent(builder.receiverIdent),
			Sel: dst.NewIdent(string(fromProp.PropertyName())),
		},
		token.ASSIGN,
		&dst.SelectorExpr{
			X:   dst.NewIdent(builder.typedInputIdent),
			Sel: dst.NewIdent(string(toProp.PropertyName())),
		})

	return []dst.Stmt{result}
}

func (builder *convertFromArmBuilder) propertiesWithSameNameButDifferentTypeHandler() propertyConversionHandler {
	definedErrVar := false

	return func(toProp *astmodel.PropertyDefinition, fromType *astmodel.ObjectType) []dst.Stmt {
		fromProp, ok := fromType.Property(toProp.PropertyName())

		if !ok || toProp.PropertyType().Equals(fromProp.PropertyType()) {
			return nil
		}

		var result []dst.Stmt

		if !definedErrVar {
			result = append(
				result,
				astbuilder.LocalVariableDeclaration("err", dst.NewIdent("error"), ""))
			definedErrVar = true
		}

		complexConversion := builder.fromArmComplexPropertyConversion(
			complexPropertyConversionParameters{
				source: &dst.SelectorExpr{
					X:   dst.NewIdent(builder.typedInputIdent),
					Sel: dst.NewIdent(string(fromProp.PropertyName())),
				},
				destination: &dst.SelectorExpr{
					X:   dst.NewIdent(builder.receiverIdent),
					Sel: dst.NewIdent(string(toProp.PropertyName())),
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
	params complexPropertyConversionParameters) []dst.Stmt {

	switch concrete := params.destinationType.(type) {
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
		return builder.assignPrimitiveType(params)
	case astmodel.ValidatedType:
		// pass through to underlying type
		params.destinationType = concrete.ElementType()
		return builder.fromArmComplexPropertyConversion(params)
	default:
		panic(fmt.Sprintf("don't know how to perform fromArm conversion for type: %s", params.destinationType.String()))
	}
}

// assignPrimitiveType just assigns source to destination directly,
// no conversion needed.
func (builder *convertFromArmBuilder) assignPrimitiveType(
	params complexPropertyConversionParameters) []dst.Stmt {

	return []dst.Stmt{
		params.assignmentHandler(params.Destination(), params.Source()),
	}
}

// convertComplexOptionalProperty handles conversion for optional properties with complex elements
// This function generates code that looks like this:
// 	if <source> != nil {
//		<code for producing result from destinationType.Element()>
//		<destination> = &<result>
//	}
func (builder *convertFromArmBuilder) convertComplexOptionalProperty(
	params complexPropertyConversionParameters) []dst.Stmt {

	destinationType := params.destinationType.(*astmodel.OptionalType)

	tempVarIdent := builder.idFactory.CreateIdentifier(params.nameHint+"Typed", astmodel.NotExported)
	tempVarType := destinationType.Element()

	newSource := &dst.UnaryExpr{
		X:  params.Source(),
		Op: token.MUL,
	}

	innerStatements := builder.fromArmComplexPropertyConversion(
		params.withDestination(dst.NewIdent(tempVarIdent)).
			withDestinationType(tempVarType).
			withAdditionalConversionContext(destinationType).
			withAssignmentHandler(assignmentHandlerDefine).
			withSource(newSource))

	// Tack on the final assignment
	innerStatements = append(
		innerStatements,
		astbuilder.SimpleAssignment(
			params.Destination(),
			token.ASSIGN,
			&dst.UnaryExpr{
				Op: token.AND,
				X:  dst.NewIdent(tempVarIdent),
			}))

	result := &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  params.Source(),
			Op: token.NEQ,
			Y:  dst.NewIdent("nil"),
		},
		Body: &dst.BlockStmt{
			List: innerStatements,
		},
	}

	return []dst.Stmt{result}
}

// convertComplexArrayProperty handles conversion for array properties with complex elements
// This function generates code that looks like this:
// 	for _, item := range <source> {
//		<code for producing result from destinationType.Element()>
//		<destination> = append(<destination>, <result>)
//	}
func (builder *convertFromArmBuilder) convertComplexArrayProperty(
	params complexPropertyConversionParameters) []dst.Stmt {

	var results []dst.Stmt

	itemIdent := "item"
	elemIdent := "elem"

	depth := params.countArraysAndMapsInConversionContext()

	destinationType := params.destinationType.(*astmodel.ArrayType)

	elemType := destinationType.Element()
	actualDestination := params.Destination() // TODO: improve name
	if depth > 0 {
		actualDestination = dst.NewIdent(elemIdent)
		results = append(
			results,
			astbuilder.LocalVariableDeclaration(
				elemIdent,
				destinationType.AsType(builder.codeGenerationContext),
				""))
		elemIdent = fmt.Sprintf("elem%d", depth)
	}

	result := &dst.RangeStmt{
		Key:   dst.NewIdent("_"),
		Value: dst.NewIdent(itemIdent),
		X:     params.Source(),
		Tok:   token.DEFINE,
		Body: &dst.BlockStmt{
			List: builder.fromArmComplexPropertyConversion(
				complexPropertyConversionParameters{
					source:            dst.NewIdent(itemIdent),
					destination:       dst.Clone(actualDestination).(dst.Expr),
					destinationType:   elemType,
					nameHint:          elemIdent,
					conversionContext: append(params.conversionContext, destinationType),
					assignmentHandler: astbuilder.AppendList,
					sameTypes:         params.sameTypes,
				}),
		},
	}
	results = append(results, result)

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.assignmentHandler != nil {
		results = append(results, params.assignmentHandler(params.Destination(), dst.Clone(actualDestination).(dst.Expr)))
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
	params complexPropertyConversionParameters) []dst.Stmt {

	destinationType := params.destinationType.(*astmodel.MapType)

	if _, ok := destinationType.KeyType().(*astmodel.PrimitiveType); !ok {
		panic(fmt.Sprintf("map had non-primitive key type: %v", destinationType.KeyType()))
	}

	depth := params.countArraysAndMapsInConversionContext()

	keyIdent := "key"
	valueIdent := "value"
	elemIdent := "elem"

	actualDestination := params.Destination() // TODO: improve name
	makeMapToken := token.ASSIGN
	if depth > 0 {
		actualDestination = dst.NewIdent(elemIdent)
		elemIdent = fmt.Sprintf("elem%d", depth)
		makeMapToken = token.DEFINE
	}

	handler := func(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
		return astbuilder.InsertMap(lhs, dst.NewIdent(keyIdent), rhs)
	}

	keyTypeAst := destinationType.KeyType().AsType(builder.codeGenerationContext)
	valueTypeAst := destinationType.ValueType().AsType(builder.codeGenerationContext)

	makeMapStatement := astbuilder.SimpleAssignment(
		dst.Clone(actualDestination).(dst.Expr),
		makeMapToken,
		astbuilder.MakeMap(keyTypeAst, valueTypeAst))
	rangeStatement := &dst.RangeStmt{
		Key:   dst.NewIdent(keyIdent),
		Value: dst.NewIdent(valueIdent),
		X:     params.Source(),
		Tok:   token.DEFINE,
		Body: &dst.BlockStmt{
			List: builder.fromArmComplexPropertyConversion(
				complexPropertyConversionParameters{
					source:            dst.NewIdent(valueIdent),
					destination:       dst.Clone(actualDestination).(dst.Expr),
					destinationType:   destinationType.ValueType(),
					nameHint:          elemIdent,
					conversionContext: append(params.conversionContext, destinationType),
					assignmentHandler: handler,
					sameTypes:         params.sameTypes,
				}),
		},
	}

	result := &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  params.Source(),
			Op: token.NEQ,
			Y:  dst.NewIdent("nil"),
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				makeMapStatement,
				rangeStatement,
			},
		},
	}

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.assignmentHandler != nil {
		result.Body.List = append(result.Body.List, params.assignmentHandler(params.Destination(), dst.Clone(actualDestination).(dst.Expr)))
	}

	return []dst.Stmt{result}
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
	params complexPropertyConversionParameters) []dst.Stmt {

	destinationType := params.destinationType.(astmodel.TypeName)
	propertyLocalVar := builder.idFactory.CreateIdentifier(params.nameHint, astmodel.NotExported)
	ownerName := builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)

	newVariable := astbuilder.NewVariable(propertyLocalVar, destinationType.Name())
	if !destinationType.PackageReference.Equals(builder.codeGenerationContext.CurrentPackage()) {
		// struct name has to be qualified
		packageName, err := builder.codeGenerationContext.GetImportedPackageName(destinationType.PackageReference)
		if err != nil {
			panic(err)
		}

		newVariable = astbuilder.NewVariableQualified(
			propertyLocalVar,
			packageName,
			destinationType.Name())
	}

	var results []dst.Stmt
	results = append(results, newVariable)
	results = append(
		results,
		astbuilder.SimpleAssignment(
			dst.NewIdent("err"),
			token.ASSIGN,
			astbuilder.CallQualifiedFunc(
				propertyLocalVar, builder.methodName, dst.NewIdent(ownerName), params.Source())))
	results = append(results, astbuilder.CheckErrorAndReturn())
	if params.assignmentHandler == nil {
		results = append(
			results,
			astbuilder.SimpleAssignment(
				params.Destination(),
				token.ASSIGN,
				dst.NewIdent(propertyLocalVar)))
	} else {
		results = append(
			results,
			params.assignmentHandler(params.Destination(), dst.NewIdent(propertyLocalVar)))
	}

	return results
}
