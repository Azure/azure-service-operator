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
		result.convertUserAssignedIdentitiesCollection,
		result.convertReferenceProperty,
		result.convertSecretProperty,
		result.convertConfigMapProperty,
		result.convertComplexTypeNameProperty)

	result.propertyConversionHandlers = []propertyConversionHandler{
		// Handlers for specific properties come first
		result.namePropertyHandler,
		result.operatorSpecPropertyHandler,
		result.configMapReferencePropertyHandler,
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
		ReceiverType:  astbuilder.PointerTo(builder.receiverTypeExpr),
		Body:          builder.functionBodyStatements(),
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

	decl := astbuilder.ShortDeclaration(
		builder.resultIdent,
		astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(dst.NewIdent(builder.armTypeIdent)).Build()))
	result = append(result, decl)

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
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	if toProp.PropertyName() != "Name" || builder.typeKind != TypeKindSpec {
		return notHandled, nil
	}

	// we do not read from AzureName() but instead use
	// the passed-in 'name' parameter which contains
	// a full name including any owners, etc
	result := astbuilder.QualifiedAssignment(
		dst.NewIdent(builder.resultIdent),
		string(toProp.PropertyName()),
		token.ASSIGN,
		astbuilder.Selector(dst.NewIdent(resolvedParameterString), "Name"))

	return handleWith(result), nil
}

func (builder *convertToARMBuilder) operatorSpecPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	if toProp.PropertyName() != astmodel.OperatorSpecProperty || builder.typeKind != TypeKindSpec {
		return notHandled, nil
	}

	// Do nothing with this property, it exists for the operator only and is not sent to Azure
	return handledWithNOP, nil
}

func (builder *convertToARMBuilder) configMapReferencePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	// This is just an optimization to avoid scanning excess properties collections
	_, isString := astmodel.AsPrimitiveType(toProp.PropertyType())

	// TODO: Do we support slices or maps? Skipped for now
	//isSliceString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.NewArrayType(astmodel.StringType))
	//isMapString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	if !isString {
		return notHandled, nil
	}

	fromProps := fromType.FindAllPropertiesWithTagValue(astmodel.OptionalConfigMapPairTag, string(toProp.PropertyName()))
	if len(fromProps) == 0 {
		return notHandled, nil
	}

	if len(fromProps) != 2 {
		// We expect exactly 2 paired properties
		return notHandled, nil
	}

	// Figure out which property is which type. There should be 1 string and 1 genruntime.ConfigMapReference
	var strProp *astmodel.PropertyDefinition
	var refProp *astmodel.PropertyDefinition
	if propType, ok := astmodel.AsPrimitiveType(fromProps[0].PropertyType()); ok && propType == astmodel.StringType {
		strProp = fromProps[0]
		refProp = fromProps[1]
	} else {
		strProp = fromProps[1]
		refProp = fromProps[0]
	}

	// This is technically more permissive than we would like as it allows collections too, but they won't make it this far because
	// of the FindAllPropertiesWithTagValue above
	optionalType, isOptional := astmodel.AsOptionalType(strProp.PropertyType())
	if !isOptional || !astmodel.TypeEquals(optionalType, astmodel.OptionalStringType) {
		return notHandled, nil
	}
	if !astmodel.TypeEquals(refProp.PropertyType(), astmodel.NewOptionalType(astmodel.ConfigMapReferenceType)) {
		// We expect the other type to be a string
		return notHandled, nil
	}

	strPropSource := astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(strProp.PropertyName()))
	refPropSource := astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(refProp.PropertyName()))

	destination := astbuilder.Selector(dst.NewIdent(builder.resultIdent), string(toProp.PropertyName()))

	strStmts := builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            strPropSource,
			SourceType:        strProp.PropertyType(),
			Destination:       destination,
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(strProp.PropertyName()),
			ConversionContext: nil,
			Locals:            builder.locals,
		},
	)
	refStmts := builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            refPropSource,
			SourceType:        refProp.PropertyType(),
			Destination:       destination,
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(strProp.PropertyName()),
			ConversionContext: nil,
			Locals:            builder.locals,
		},
	)

	return handleWith(
		strStmts,
		refStmts,
	), nil
}

func (builder *convertToARMBuilder) userAssignedIdentitiesPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	if _, ok := astmodel.IsUserAssignedIdentityProperty(toProp); !ok {
		return notHandled, nil
	}

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return notHandled, nil
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

func (builder *convertToARMBuilder) referencePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	// This is just an optimization to avoid scanning excess properties collections
	isString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.StringType)
	isOptionalString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.OptionalStringType)
	isSliceString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.NewArrayType(astmodel.StringType))
	isMapString := astmodel.TypeEquals(toProp.PropertyType(), astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	if !isString && !isOptionalString && !isSliceString && !isMapString {
		return notHandled, nil
	}

	// Find the property which is referring to our toProp in its ARMReferenceTag. If we can't find it, that means
	// there's not one and this handler doesn't apply
	fromProp, foundReference := fromType.FindPropertyWithTagValue(astmodel.ARMReferenceTag, string(toProp.PropertyName()))
	if !foundReference {
		return notHandled, nil
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
//	armObj.X.Y1 = k8sObj.Y1;
//	armObj.X.Y2 = k8sObj.Y2;
//
// in reality each assignment is likely to be another conversion that is specific
// to the type being converted.
func (builder *convertToARMBuilder) flattenedPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

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
		return notHandled, nil
	}

	allDefs := builder.codeGenerationContext.GetAllReachableDefinitions()

	// the toProp shape here must be:
	// 1. maybe a typename, pointing to…
	// 2. maybe optional, wrapping …
	// 3. maybe a typename, pointing to…
	// 4. an object type

	// (1.) resolve the outer typename
	toPropType, err := allDefs.FullyResolve(toProp.PropertyType())
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
		toPropType, err = allDefs.FullyResolve(optType.Element())
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

	return handleWith(result), nil
}

// buildToPropInitializer builds an initializer for a given “to” property
// that assigns it a value if any of the “from” properties are not nil.
//
// Resultant code looks like:
//
//	if (from1 != nil) || (from2 != nil) || … {
//			<resultIdent>.<toProp> = &<toPropTypeName>{}
//	}
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
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return notHandled, nil
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

// convertUserAssignedIdentitiesCollection handles conversion the special UserAssignedIdentities property.
// This function generates code that looks like this:
//
//	result.UserAssignedIdentities = make(map[string]UserAssignedIdentityDetails_ARM, len(identity.UserAssignedIdentities))
//	for _, ident := range identity.UserAssignedIdentities {
//		identARMID, err := resolved.ResolvedReferences.Lookup(ident.Reference)
//		if err != nil {
//			return nil, err
//		}
//		key := identARMID
//		result.UserAssignedIdentities[key] = UserAssignedIdentityDetails_ARM{}
//	}
//	return result, nil
func (builder *convertToARMBuilder) convertUserAssignedIdentitiesCollection(
	conversionBuilder *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, isDestinationMap := params.DestinationType.(*astmodel.MapType)
	if !isDestinationMap {
		return nil, nil
	}

	sourceType, isSourceArray := params.SourceType.(*astmodel.ArrayType)
	if !isSourceArray {
		return nil, nil
	}

	typeName, ok := astmodel.AsTypeName(sourceType.Element())
	if !ok {
		return nil, nil
	}

	if typeName.Name() != astmodel.UserAssignedIdentitiesTypeName {
		return nil, nil
	}

	uaiDef := conversionBuilder.CodeGenerationContext.MustGetDefinition(typeName)

	uaiType, ok := astmodel.AsObjectType(uaiDef.Type())
	if !ok {
		return nil, nil
	}

	// There should be a single "Reference" property
	refProperty, ok := uaiType.Property("Reference")
	if !ok {
		panic(fmt.Sprintf("Found UserAssignedIdentity type without Reference property"))
	}

	locals := params.Locals.Clone()

	itemIdent := locals.CreateLocal("ident")
	keyTypeAst := destinationType.KeyType().AsType(conversionBuilder.CodeGenerationContext)
	valueTypeAst := destinationType.ValueType().AsType(conversionBuilder.CodeGenerationContext)

	makeMapStatement := astbuilder.AssignmentStatement(
		params.Destination,
		token.ASSIGN,
		astbuilder.MakeMapWithCapacity(keyTypeAst, valueTypeAst,
			astbuilder.CallFunc("len", params.Source)))

	key := "key"

	refSelector := astbuilder.Selector(dst.NewIdent(itemIdent), "Reference")

	// Rely on existing conversion handler for ResourceReference type
	conversionStmts := conversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            refSelector,
			SourceType:        refProperty.PropertyType(),
			Destination:       dst.NewIdent(key),
			DestinationType:   destinationType.KeyType(),
			NameHint:          itemIdent,
			ConversionContext: append(params.ConversionContext, destinationType),
			AssignmentHandler: astmodel.AssignmentHandlerDefine,
			Locals:            locals,
		})
	valueBuilder := astbuilder.NewCompositeLiteralBuilder(valueTypeAst).WithoutNewLines()

	conversionStmts = append(
		conversionStmts,
		astbuilder.InsertMap(params.Destination, dst.NewIdent(key), valueBuilder.Build()))

	// Loop over the slice
	loop := astbuilder.IterateOverSlice(
		itemIdent,
		params.Source,
		conversionStmts...)

	return []dst.Stmt{
		makeMapStatement,
		loop}
}

// convertReferenceProperty handles conversion of reference properties.
// This function generates code that looks like this:
//
//	<namehint>ARMID, err := resolved.ResolvedReferences.Lookup(<source>)
//	if err != nil {
//		return nil, err
//	}
//	<destination> = <namehint>ARMID
func (builder *convertToARMBuilder) convertReferenceProperty(
	_ *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	isString := astmodel.TypeEquals(params.DestinationType, astmodel.StringType)
	if !isString {
		return nil, nil
	}

	isReference := astmodel.TypeEquals(params.SourceType, astmodel.ResourceReferenceType)
	if !isReference {
		return nil, nil
	}

	// Don't need to worry about conflicting names here since the property name was unique to begin with
	localVarName := builder.idFactory.CreateLocal(params.NameHint + "ARMID")
	armIDLookup := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(localVarName),
		token.DEFINE,
		astbuilder.CallExpr(
			astbuilder.Selector(dst.NewIdent(resolvedParameterString), "ResolvedReferences"),
			"Lookup",
			params.Source))

	returnIfNotNil := astbuilder.ReturnIfNotNil(dst.NewIdent("err"), astbuilder.Nil(), dst.NewIdent("err"))

	result := params.AssignmentHandlerOrDefault()(params.Destination, dst.NewIdent(localVarName))

	return astbuilder.Statements(armIDLookup, returnIfNotNil, result), nil
}

// convertSecretProperty handles conversion of secret properties.
// This function generates code that looks like this:
//
//	<namehint>Secret, err := resolved.ResolvedSecrets.Lookup(<source>)
//	if err != nil {
//		return nil, errors.Wrap(err, "looking up secret for <source>")
//	}
//	<destination> = <namehint>Secret
func (builder *convertToARMBuilder) convertSecretProperty(
	_ *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	isString := astmodel.TypeEquals(params.DestinationType, astmodel.StringType)
	if !isString {
		return nil, nil
	}

	isSecretReference := astmodel.TypeEquals(params.SourceType, astmodel.SecretReferenceType)
	if !isSecretReference {
		return nil, nil
	}

	errorsPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	localVarName := builder.idFactory.CreateLocal(params.NameHint + "Secret")
	secretLookup := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(localVarName),
		token.DEFINE,
		astbuilder.CallExpr(
			astbuilder.Selector(dst.NewIdent(resolvedParameterString), "ResolvedSecrets"),
			"Lookup",
			params.Source))

	wrappedError := astbuilder.WrapError(
		errorsPackage,
		"err",
		fmt.Sprintf("looking up secret for property %s", params.NameHint))
	returnIfNotNil := astbuilder.ReturnIfNotNil(dst.NewIdent("err"), astbuilder.Nil(), wrappedError)

	result := params.AssignmentHandlerOrDefault()(params.Destination, dst.NewIdent(localVarName))

	return astbuilder.Statements(secretLookup, returnIfNotNil, result), nil
}

// convertConfigMapProperty handles conversion of configMap properties.
// This function generates code that looks like this:
//
//	<namehint>Value, err := resolved.ResolvedConfigMaps.Lookup(<source>)
//	if err != nil {
//		return nil, errors.Wrap(err, "looking up config map value for <source>")
//	}
//	<destination> = <namehint>Value
func (builder *convertToARMBuilder) convertConfigMapProperty(
	_ *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	isString := astmodel.TypeEquals(params.DestinationType, astmodel.StringType)
	if !isString {
		return nil, nil
	}

	isConfigMapReference := astmodel.TypeEquals(params.SourceType, astmodel.ConfigMapReferenceType)
	if !isConfigMapReference {
		return nil, nil
	}

	errorsPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	localVarName := params.Locals.CreateLocal(params.NameHint + "Value")
	configMapLookup := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(localVarName),
		token.DEFINE,
		astbuilder.CallExpr(
			astbuilder.Selector(dst.NewIdent(resolvedParameterString), "ResolvedConfigMaps"),
			"Lookup",
			params.Source))

	wrappedError := astbuilder.WrapError(
		errorsPackage,
		"err",
		fmt.Sprintf("looking up configmap for property %s", params.NameHint))
	returnIfNotNil := astbuilder.ReturnIfNotNil(dst.NewIdent("err"), astbuilder.Nil(), wrappedError)

	result := params.AssignmentHandlerOrDefault()(params.Destination, dst.NewIdent(localVarName))

	return astbuilder.Statements(configMapLookup, returnIfNotNil, result), nil
}

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
//
//	<nameHint>, err := <source>.ToARM(name)
//	if err != nil {
//		return nil, err
//	}
//	<destination> = <nameHint>.(*FooARM)
func (builder *convertToARMBuilder) convertComplexTypeNameProperty(
	conversionBuilder *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(astmodel.TypeName)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(astmodel.TypeName)
	if !ok {
		return nil, nil
	}

	// This is for handling type names that aren't equal
	if astmodel.TypeEquals(sourceType, destinationType) {
		return nil, nil
	}

	var results []dst.Stmt
	propertyLocalVarName := params.Locals.CreateLocal(params.NameHint, astmodel.ARMSuffix)

	// Call ToARM on the property
	results = append(results, callToARMFunction(params.GetSource(), dst.NewIdent(propertyLocalVarName), builder.methodName)...)

	typeAssertExpr := &dst.TypeAssertExpr{
		X:    dst.NewIdent(propertyLocalVarName),
		Type: astbuilder.Dereference(dst.NewIdent(destinationType.Name())),
	}

	if !destinationType.PackageReference.Equals(conversionBuilder.CodeGenerationContext.CurrentPackage()) {
		// needs to be qualified
		packageName, err := conversionBuilder.CodeGenerationContext.GetImportedPackageName(destinationType.PackageReference)
		if err != nil {
			panic(err)
		}

		typeAssertExpr.Type = astbuilder.Dereference(astbuilder.Selector(dst.NewIdent(packageName), destinationType.Name()))
	}

	// TODO: This results in code that isn't very "human-like". Today the contract of most handlers is that they
	// TODO: result in a type which is not a ptr. This is a useful contract as then the caller always knows if they need
	// TODO: to take the address of the inner result (&result) to assign to a ptr field, or not (to assign to a non-ptr field).
	// TODO: Unfortunately, we can't fix this issue by inverting things and making the contract that the type is a ptr type, as
	// TODO: in many cases (primitive types, strings, etc) that doesn't make sense and also results in awkward to read code.
	finalAssignmentExpr := astbuilder.Dereference(typeAssertExpr)

	results = append(results, params.AssignmentHandlerOrDefault()(params.GetDestination(), finalAssignmentExpr))

	return results, nil
}

func callToARMFunction(source dst.Expr, destination dst.Expr, methodName string) []dst.Stmt {
	// Call ToARM on the property
	propertyToARMInvocation := astbuilder.SimpleAssignmentWithErr(
		destination,
		token.DEFINE,
		// Don't use astbuilder.CallExpr() because it flattens dereferences,
		&dst.CallExpr{
			Fun: astbuilder.Selector(source, methodName),
			Args: []dst.Expr{
				dst.NewIdent(resolvedParameterString),
			},
		})

	return astbuilder.Statements(
		propertyToARMInvocation,
		astbuilder.CheckErrorAndReturn(astbuilder.Nil()))
}
