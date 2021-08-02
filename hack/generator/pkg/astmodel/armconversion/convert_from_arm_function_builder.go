/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

type convertFromARMBuilder struct {
	conversionBuilder
	typedInputIdent       string
	inputIdent            string
	typeConversionBuilder *astmodel.ConversionFunctionBuilder
	locals                *astmodel.KnownLocalsSet
}

func newConvertFromARMFunctionBuilder(
	c *ARMConversionFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *convertFromARMBuilder {

	result := &convertFromARMBuilder{
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
			typeKind:              c.typeKind,
			codeGenerationContext: codeGenerationContext,
		},
		typeConversionBuilder: astmodel.NewConversionFunctionBuilder(c.idFactory, codeGenerationContext),
		locals:                astmodel.NewKnownLocalsSet(c.idFactory),
	}
	// Add the receiver ident into the known locals
	result.locals.Add(result.receiverIdent)

	// It's a bit awkward that there are two levels of "handler" here, but they serve different purposes:
	// The top level propertyConversionHandlers is about determining which properties are involved: given a property on the destination type it
	// determines which property (if any) on the source type will be converted to the destination.
	// The "inner" handler (typeConversionBuilder) is about determining how to convert between two types: given a
	// source type and a destination type, figure out how to make the assignment work. It has no knowledge of broader object strucutre
	// or other properties.
	result.typeConversionBuilder.AddConversionHandlers(result.convertComplexTypeNameProperty)
	result.propertyConversionHandlers = []propertyConversionHandler{
		result.namePropertyHandler,
		result.referencePropertyHandler,
		result.ownerPropertyHandler,
		result.flattenedPropertyHandler,
		result.propertiesWithSameNameHandler,
	}

	return result
}

func (builder *convertFromARMBuilder) functionDeclaration() *dst.FuncDecl {
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
		astbuilder.Selector(dst.NewIdent(astmodel.GenRuntimePackageName), "KnownResourceReference"))

	fn.AddParameter(builder.inputIdent, dst.NewIdent("interface{}"))
	fn.AddReturns("error")
	return fn.DefineFunc()
}

func (builder *convertFromARMBuilder) functionBodyStatements() []dst.Stmt {
	var result []dst.Stmt

	conversionStmts := generateTypeConversionAssignments(
		builder.armType,
		builder.kubeType,
		builder.propertyConversionHandler)

	hasConversions := len(removeEmptyStatements(conversionStmts)) > 0

	assertStmts := builder.assertInputTypeIsARM(hasConversions)

	// perform a type assert and check its results
	result = append(result, assertStmts...)

	if hasConversions {
		result = append(result, conversionStmts...)
	}

	result = append(result, astbuilder.ReturnNoError())

	return result
}

func (builder *convertFromARMBuilder) assertInputTypeIsARM(needsResult bool) []dst.Stmt {
	var result []dst.Stmt

	fmtPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.FmtReference)

	dest := builder.typedInputIdent
	if !needsResult {
		dest = "_" // drop result
	}

	// perform a type assert
	result = append(
		result,
		astbuilder.TypeAssert(
			dst.NewIdent(dest),
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

func (builder *convertFromARMBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	if builder.typeKind != TypeKindSpec || !toProp.HasName(astmodel.AzureNameProperty) {
		return nil
	}

	// Check to make sure that the ARM object has a "Name" property (which matches our "AzureName")
	fromProp, ok := fromType.Property("Name")
	if !ok {
		panic("ARM resource missing property 'Name'")
	}

	// Invoke SetAzureName(ExtractKubernetesResourceNameFromARMName(this.Name)):
	return []dst.Stmt{
		&dst.ExprStmt{
			X: astbuilder.CallQualifiedFunc(
				builder.receiverIdent,
				"SetAzureName",
				astbuilder.CallQualifiedFunc(
					astmodel.GenRuntimePackageName,
					"ExtractKubernetesResourceNameFromARMName",
					astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()))),
			),
		},
	}
}

func (builder *convertFromARMBuilder) referencePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	isResourceReference := toProp.PropertyType().Equals(astmodel.ResourceReferenceTypeName)
	isOptionalResourceReference := toProp.PropertyType().Equals(astmodel.NewOptionalType(astmodel.ResourceReferenceTypeName))

	if !isResourceReference && !isOptionalResourceReference {
		return nil
	}

	// TODO: For now, we are NOT assigning to these. _Status types don't have them and it's unclear what
	// TODO: the fromARM functions do for us on Spec types. We may need them for diffing though. If so we will
	// TODO: need to revisit this and actually assign something
	// Returning an empty statement allows us to "consume" this match and not proceed to the next handler.
	// There is nothing included in the generated code.
	return []dst.Stmt{
		&dst.EmptyStmt{},
	}
}

func (builder *convertFromARMBuilder) ownerPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType) []dst.Stmt {

	if toProp.PropertyName() != builder.idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported) || builder.typeKind != TypeKindSpec {
		return nil
	}

	result := astbuilder.QualifiedAssignment(
		dst.NewIdent(builder.receiverIdent),
		string(toProp.PropertyName()),
		token.ASSIGN,
		dst.NewIdent(builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)))
	return []dst.Stmt{result}
}

// flattenedPropertyHandler generates conversions for properties that
// were flattened out from inside other properties. The code it generates will
// look something like:
//
// If 'X' is a property that was flattened:
//
//   k8sObj.Y1 = armObj.X.Y1;
//   k8sObj.Y2 = armObj.X.Y2;
//
// in reality each assignment is likely to be another conversion that is specific
// to the type being converted.
func (builder *convertFromARMBuilder) flattenedPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	if !toProp.WasFlattened() {
		return nil
	}

	for _, fromProp := range fromType.Properties() {
		if toProp.WasFlattenedFrom(fromProp.PropertyName()) {
			return builder.buildFlattenedAssignment(toProp, fromProp)
		}
	}

	panic(fmt.Sprintf("couldn’t find source ARM property ‘%s’ that k8s property ‘%s’ was flattened from", toProp.FlattenedFrom()[0], toProp.PropertyName()))
}

func (builder *convertFromARMBuilder) buildFlattenedAssignment(toProp *astmodel.PropertyDefinition, fromProp *astmodel.PropertyDefinition) []dst.Stmt {
	if len(toProp.FlattenedFrom()) > 1 {
		// this doesn't appear to happen anywhere in the JSON schemas currently

		var props []string
		for _, ff := range toProp.FlattenedFrom() {
			props = append(props, string(ff))
		}

		panic(fmt.Sprintf("need to implement multiple levels of flattening: property ‘%s’ on %s was flattened from ‘%s’",
			toProp.PropertyName(),
			builder.receiverIdent,
			strings.Join(props, ".")))
	}

	allTypes := builder.codeGenerationContext.GetAllReachableTypes()

	// the from shape here must be:
	// 1. maybe a typename, pointing to…
	// 2. maybe optional, wrapping …
	// 3. maybe a typename, pointing to…
	// 4. an object type

	// (1.) resolve any outer typename
	fromPropType, err := allTypes.FullyResolve(fromProp.PropertyType())
	if err != nil {
		panic(err)
	}

	var fromPropObjType *astmodel.ObjectType
	var objOk bool
	// (2.) resolve any optional type
	generateNilCheck := false
	if fromPropOptType, ok := fromPropType.(*astmodel.OptionalType); ok {
		generateNilCheck = true
		// (3.) resolve any inner typename
		elementType, err := allTypes.FullyResolve(fromPropOptType.Element())
		if err != nil {
			panic(err)
		}

		// (4.) resolve the inner object type
		fromPropObjType, objOk = elementType.(*astmodel.ObjectType)
	} else {
		// (4.) resolve the inner object type
		fromPropObjType, objOk = fromPropType.(*astmodel.ObjectType)
	}

	if !objOk {
		// see pipeline_flatten_properties.go:flattenPropType which will only flatten from (optional) object types
		panic(fmt.Sprintf("property ‘%s’ marked as flattened from non-object type %T, which shouldn’t be possible",
			toProp.PropertyName(),
			fromPropType))
	}

	// *** Now generate the code! ***
	toPropName := toProp.PropertyName()
	nestedProp, ok := fromPropObjType.Property(toPropName)
	if !ok {
		panic("couldn't find source of flattened property")
	}

	// need to make a clone of builder.locals if we are going to nest in an if statement
	locals := builder.locals
	if generateNilCheck {
		locals = locals.Clone()
	}

	stmts := builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()), string(toPropName)),
			SourceType:        nestedProp.PropertyType(),
			Destination:       astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(toPropName)),
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(toProp.PropertyName()),
			ConversionContext: nil,
			AssignmentHandler: nil,
			Locals:            locals,
		})

	// we were unable to generate an inner conversion so we cannot generate the overall conversion
	if len(stmts) == 0 {
		return nil
	}

	if generateNilCheck {
		propToCheck := astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()))
		stmts = []dst.Stmt{
			&dst.IfStmt{
				Cond: astbuilder.NotNil(propToCheck),
				Body: &dst.BlockStmt{List: stmts},
			},
		}
	}

	result := []dst.Stmt{
		&dst.EmptyStmt{
			Decs: dst.EmptyStmtDecorations{
				NodeDecs: dst.NodeDecs{
					End: []string{"// copying flattened property:"},
				},
			},
		},
	}

	return append(result, stmts...)
}

func (builder *convertFromARMBuilder) propertiesWithSameNameHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType) []dst.Stmt {

	fromProp, ok := fromType.Property(toProp.PropertyName())
	if !ok {
		return nil
	}

	return builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:            astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName())),
			SourceType:        fromProp.PropertyType(),
			Destination:       astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(toProp.PropertyName())),
			DestinationType:   toProp.PropertyType(),
			NameHint:          string(toProp.PropertyName()),
			ConversionContext: nil,
			AssignmentHandler: nil,
			Locals:            builder.locals,
		})
}

//////////////////////////////////////////////////////////////////////////////////
// Complex property conversion (for when properties aren't simple primitive types)
//////////////////////////////////////////////////////////////////////////////////

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
//	<nameHint>Converted := <destinationType>{}
//	err = <nameHint>Converted.FromARM(owner, <source>)
//	if err != nil {
//		return err
//	}
//	<destination> = <nameHint>
func (builder *convertFromARMBuilder) convertComplexTypeNameProperty(conversionBuilder *astmodel.ConversionFunctionBuilder, params astmodel.ConversionParameters) []dst.Stmt {
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

	propertyLocalVar := builder.typeConversionBuilder.CreateLocal(params.Locals, "", params.NameHint)
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

	tok := token.ASSIGN
	if !params.Locals.HasName("err") {
		tok = token.DEFINE
		params.Locals.Add("err")
	}

	var results []dst.Stmt
	results = append(results, newVariable)
	results = append(
		results,
		astbuilder.AssignmentStatement(
			dst.NewIdent("err"),
			tok,
			astbuilder.CallQualifiedFunc(
				propertyLocalVar, builder.methodName, dst.NewIdent(ownerName), params.GetSource())))
	results = append(results, astbuilder.CheckErrorAndReturn())
	if params.AssignmentHandler == nil {
		results = append(
			results,
			astbuilder.SimpleAssignment(
				params.GetDestination(),
				dst.NewIdent(propertyLocalVar)))
	} else {
		results = append(
			results,
			params.AssignmentHandler(params.GetDestination(), dst.NewIdent(propertyLocalVar)))
	}

	return results
}
