/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const (
	resolvedParameterString = "resolved"
)

type convertToARMBuilder struct {
	conversionBuilder
	resultIdent           string
	typeConversionBuilder *astmodel.ConversionFunctionBuilder
	locals                *astmodel.KnownLocalsSet
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
			receiverIdent:         c.idFactory.CreateReceiver(receiver.Name()),
			receiverTypeExpr:      receiver.AsType(codeGenerationContext),
			armTypeIdent:          c.armTypeName.Name(),
			idFactory:             c.idFactory,
			typeKind:              c.typeKind,
			codeGenerationContext: codeGenerationContext,
		},
		resultIdent:           "result",
		typeConversionBuilder: astmodel.NewConversionFunctionBuilder(c.idFactory, codeGenerationContext),
		locals:                astmodel.NewKnownLocalsSet(c.idFactory),
	}
	// Add the receiver ident into the known locals
	result.locals.Add(result.receiverIdent)

	// It's a bit awkward that there are two levels of "handler" here, but they serve different purposes:
	// The top level propertyConversionHandlers is about determining which properties are involved: given a property on
	// the destination type it determines which property (if any) on the source type will be converted to the destination.
	// The "inner" handler (typeConversionBuilder) is about determining how to convert between two types: given a
	// source type and a destination type, figure out how to make the assignment work. It has no knowledge of broader
	// object structure or other properties.
	result.typeConversionBuilder.AddConversionHandlers(
		result.convertReferenceProperty,
		result.convertSecretProperty,
		result.convertComplexTypeNameProperty)

	result.propertyConversionHandlers = []propertyConversionHandler{
		// Handlers for specific properties come first
		result.namePropertyHandler,
		// Generic handlers come second
		result.referencePropertyHandler,
		result.flattenedPropertyHandler,
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

	fn.AddParameter(resolvedParameterString, astmodel.ConvertToARMResolvedDetailsType.AsType(builder.codeGenerationContext))
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
		astbuilder.ReturnIfNil(dst.NewIdent(builder.receiverIdent), astbuilder.Nil(), astbuilder.Nil()))
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
			astbuilder.Nil(),
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
	_ *astmodel.ObjectType) ([]dst.Stmt, bool) {

	if toProp.PropertyName() != "Name" || builder.typeKind != TypeKindSpec {
		return nil, false
	}

	// we do not read from AzureName() but instead use
	// the passed-in 'name' parameter which contains
	// a full name including any owners, etc
	result := astbuilder.QualifiedAssignment(
		dst.NewIdent(builder.resultIdent),
		string(toProp.PropertyName()),
		token.ASSIGN,
		astbuilder.Selector(dst.NewIdent(resolvedParameterString), "Name"))

	return []dst.Stmt{result}, true
}

func (builder *convertToARMBuilder) referencePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) ([]dst.Stmt, bool) {

	// This is just an optimization to avoid scanning excess properties collections
	isString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.StringType)
	isOptionalString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.NewOptionalType(astmodel.StringType))
	if !isString && !isOptionalString {
		return nil, false
	}

	// Find the property which is referring to our toProp in its ARMReferenceTag. If we can't find it, that means
	// there's not one and this handler doesn't apply
	fromProp, foundReference := fromType.FindPropertyWithTagValue(astmodel.ARMReferenceTag, string(toProp.PropertyName()))
	if !foundReference {
		return nil, false
	}

	source := &dst.SelectorExpr{
		X:   dst.NewIdent(builder.receiverIdent),
		Sel: dst.NewIdent(string(fromProp.PropertyName())),
	}

	destination := &dst.SelectorExpr{
		X:   dst.NewIdent(builder.resultIdent),
		Sel: dst.NewIdent(string(toProp.PropertyName())),
	}

	return builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            source,
			SourceType:        fromProp.PropertyType(),
			Destination:       destination,
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(fromProp.PropertyName()),
			ConversionContext: nil,
			Locals:            builder.locals,
		},
	), true
}

// flattenedPropertyHandler generates conversions for properties that
// were flattened out from inside other properties. The code it generates will
// look something like:
//
// If 'X' is a property that was flattened:
//
//   armObj.X.Y1 = k8sObj.Y1;
//   armObj.X.Y2 = k8sObj.Y2;
//
// in reality each assignment is likely to be another conversion that is specific
// to the type being converted.
func (builder *convertToARMBuilder) flattenedPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) ([]dst.Stmt, bool) {

	toPropName := toProp.PropertyName()

	// collect any fromProps that were flattened from the to-prop
	var fromProps []*astmodel.PropertyDefinition
	for _, prop := range fromType.Properties().AsSlice() {
		if prop.WasFlattenedFrom(toPropName) {
			fromProps = append(fromProps, prop)
		}
	}

	// there are none to copy; exit
	if len(fromProps) == 0 {
		return nil, false
	}

	allTypes := builder.codeGenerationContext.GetAllReachableTypes()

	// the toProp shape here must be:
	// 1. maybe a typename, pointing to…
	// 2. maybe optional, wrapping …
	// 3. maybe a typename, pointing to…
	// 4. an object type

	// (1.) resolve the outer typename
	toPropType, err := allTypes.FullyResolve(toProp.PropertyType())
	if err != nil {
		panic(err)
	}

	needToInitializeToProp := false // we need to init the target if it is optional
	var toPropTypeName astmodel.TypeName
	// (2.)  resolve any optional type
	if optType, ok := astmodel.AsOptionalType(toPropType); ok {
		needToInitializeToProp = true
		// (3.) resolve any inner typename
		toPropTypeName = optType.Element().(astmodel.TypeName)
		toPropType, err = allTypes.FullyResolve(optType.Element())
		if err != nil {
			panic(err)
		}
	}

	// (4.) we have found the underlying object type
	toPropObjType, _ := astmodel.AsObjectType(toPropType)

	// *** Now generate the code! ***

	// Build the initializer for the to-prop (if needed)
	var result []dst.Stmt
	if needToInitializeToProp {
		result = []dst.Stmt{builder.buildToPropInitializer(fromProps, toPropTypeName, toPropName)}
	}

	// Copy each from-prop into the to-prop
	for _, fromProp := range fromProps {
		// find the corresponding inner property on the to-prop type
		// TODO: If this property is an ARM reference we need a bit of special handling.
		// TODO: See https://github.com/Azure/azure-service-operator/issues/1651 for possible improvements to this.
		toSubPropName := fromProp.FlattenedFrom()[len(fromProp.FlattenedFrom())-1]
		if values, ok := fromProp.Tag(astmodel.ARMReferenceTag); ok {
			toSubPropName = astmodel.PropertyName(values[0])
		}

		toSubProp, ok := toPropObjType.Property(toSubPropName)
		if !ok {
			panic(fmt.Sprintf("unable to find expected property %s inside property %s", fromProp.PropertyName(), toPropName))
		}

		// generate conversion
		stmts := builder.typeConversionBuilder.BuildConversion(
			astmodel.ConversionParameters{
				Source:            astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(fromProp.PropertyName())),
				SourceType:        fromProp.PropertyType(),
				Destination:       astbuilder.Selector(dst.NewIdent(builder.resultIdent), string(toPropName), string(toSubProp.PropertyName())),
				DestinationType:   toSubProp.PropertyType(),
				NameHint:          string(toSubProp.PropertyName()),
				ConversionContext: nil,
				AssignmentHandler: nil,
				Locals:            builder.locals,
			})

		// we were unable to generate an inner conversion, so we cannot generate the overall conversion
		if len(stmts) == 0 {
			return nil, false
		}

		result = append(result, stmts...)
	}

	return result, true
}

// buildToPropInitializer builds an initializer for a given “to” property
// that assigns it a value if any of the “from” properties are not nil.
//
// Resultant code looks like:
// if (from1 != nil) || (from2 != nil) || … {
// 		<resultIdent>.<toProp> = &<toPropTypeName>{}
// }
func (builder *convertToARMBuilder) buildToPropInitializer(
	fromProps []*astmodel.PropertyDefinition,
	toPropTypeName astmodel.TypeName,
	toPropName astmodel.PropertyName) dst.Stmt {

	// build (x != nil, y != nil, …)
	conditions := make([]dst.Expr, 0, len(fromProps))
	for _, prop := range fromProps {
		propSel := astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(prop.PropertyName()))
		conditions = append(conditions, astbuilder.NotNil(propSel))
	}

	// build (x || y || …)
	cond := astbuilder.JoinOr(conditions...)

	literal := astbuilder.NewCompositeLiteralBuilder(toPropTypeName.AsType(builder.codeGenerationContext))

	// build if (conditions…) { target.prop = &TargetType{} }
	return &dst.IfStmt{
		Cond: cond,
		Body: astbuilder.StatementBlock(
			astbuilder.QualifiedAssignment(
				dst.NewIdent(builder.resultIdent),
				string(toPropName),
				token.ASSIGN,
				astbuilder.AddrOf(literal.Build()))),
	}
}

func (builder *convertToARMBuilder) propertiesWithSameNameHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) ([]dst.Stmt, bool) {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return nil, false
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
			Locals:            builder.locals,
		},
	), true
}

// convertReferenceProperty handles conversion of reference properties.
// This function generates code that looks like this:
//	<namehint>ARMID, err := resolved.ResolvedReferences.ARMIDOrErr(<source>)
//	if err != nil {
//		return nil, err
//	}
//	<destination> = <namehint>ARMID
func (builder *convertToARMBuilder) convertReferenceProperty(_ *astmodel.ConversionFunctionBuilder, params astmodel.ConversionParameters) []dst.Stmt {
	isString := astmodel.TypeEquals(params.DestinationType, astmodel.StringType)
	if !isString {
		return nil
	}

	isReference := astmodel.TypeEquals(params.SourceType, astmodel.ResourceReferenceType)
	if !isReference {
		return nil
	}

	// Don't need to worry about conflicting names here since the property name was unique to begin with
	localVarName := builder.idFactory.CreateLocal(params.NameHint + "ARMID")
	armIDLookup := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(localVarName),
		token.DEFINE,
		astbuilder.CallExpr(
			astbuilder.Selector(dst.NewIdent(resolvedParameterString), "ResolvedReferences"),
			"ARMIDOrErr",
			params.Source))

	returnIfNotNil := astbuilder.ReturnIfNotNil(dst.NewIdent("err"), astbuilder.Nil(), dst.NewIdent("err"))

	result := params.AssignmentHandlerOrDefault()(params.Destination, dst.NewIdent(localVarName))

	return []dst.Stmt{armIDLookup, returnIfNotNil, result}
}

// convertSecretProperty handles conversion of secret properties.
// This function generates code that looks like this:
//	<namehint>Secret, err := resolved.ResolvedSecrets.LookupSecret(<source>)
//	if err != nil {
//		return nil, errors.Wrap(err, "looking up secret for <source>")
//	}
//	<destination> = <namehint>Secret
func (builder *convertToARMBuilder) convertSecretProperty(_ *astmodel.ConversionFunctionBuilder, params astmodel.ConversionParameters) []dst.Stmt {
	isString := astmodel.TypeEquals(params.DestinationType, astmodel.StringType)
	if !isString {
		return nil
	}

	isSecretReference := astmodel.TypeEquals(params.SourceType, astmodel.SecretReferenceType)
	if !isSecretReference {
		return nil
	}

	errorsPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	localVarName := builder.idFactory.CreateLocal(params.NameHint + "Secret")
	secretLookup := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(localVarName),
		token.DEFINE,
		astbuilder.CallExpr(
			astbuilder.Selector(dst.NewIdent(resolvedParameterString), "ResolvedSecrets"),
			"LookupSecret",
			params.Source))

	wrappedError := astbuilder.WrapError(
		errorsPackage,
		"err",
		fmt.Sprintf("looking up secret for property %s", params.NameHint))
	returnIfNotNil := astbuilder.ReturnIfNotNil(dst.NewIdent("err"), astbuilder.Nil(), wrappedError)

	result := params.AssignmentHandlerOrDefault()(params.Destination, dst.NewIdent(localVarName))

	return []dst.Stmt{secretLookup, returnIfNotNil, result}
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
	if astmodel.TypeEquals(sourceType, destinationType) {
		return nil
	}

	var results []dst.Stmt
	propertyLocalVarName := params.Locals.CreateLocal(params.NameHint, "ARM")

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
					dst.NewIdent(resolvedParameterString),
				},
			},
		},
	}
	results = append(results, propertyToARMInvocation)
	results = append(results, astbuilder.CheckErrorAndReturn(astbuilder.Nil()))

	return results
}
