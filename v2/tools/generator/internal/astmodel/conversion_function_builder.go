/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// ConversionParameters are parameters for converting between a source type and a destination type.
type ConversionParameters struct {
	Source            dst.Expr
	Destination       dst.Expr
	SourceType        Type
	DestinationType   Type
	NameHint          string
	ConversionContext []Type
	AssignmentHandler func(destination, source dst.Expr) dst.Stmt
	Locals            *KnownLocalsSet
}

// GetSource gets the Source field.
func (params ConversionParameters) GetSource() dst.Expr {
	return dst.Clone(params.Source).(dst.Expr)
}

// GetDestination gets the Destination field.
func (params ConversionParameters) GetDestination() dst.Expr {
	return dst.Clone(params.Destination).(dst.Expr)
}

// WithSource returns a new ConversionParameters with the updated Source.
func (params ConversionParameters) WithSource(source dst.Expr) ConversionParameters {
	result := params.copy()
	result.Source = source

	return result
}

// WithSourceType returns a new ConversionParameters with the updated SourceType.
func (params ConversionParameters) WithSourceType(t Type) ConversionParameters {
	result := params.copy()
	result.SourceType = t

	return result
}

// WithDestination returns a new ConversionParameters with the updated Destination.
func (params ConversionParameters) WithDestination(destination dst.Expr) ConversionParameters {
	result := params.copy()
	result.Destination = destination

	return result
}

// WithDestinationType returns a new ConversionParameters with the updated DestinationType.
func (params ConversionParameters) WithDestinationType(t Type) ConversionParameters {
	result := params.copy()
	result.DestinationType = t

	return result
}

// WithAssignmentHandler returns a new ConversionParameters with the updated AssignmentHandler.
func (params ConversionParameters) WithAssignmentHandler(
	assignmentHandler func(result dst.Expr, destination dst.Expr) dst.Stmt) ConversionParameters {
	result := params.copy()
	result.AssignmentHandler = assignmentHandler

	return result
}

// AssignmentHandlerOrDefault returns the AssignmentHandler or a default assignment handler if AssignmentHandler was nil.
func (params ConversionParameters) AssignmentHandlerOrDefault() func(destination, source dst.Expr) dst.Stmt {
	if params.AssignmentHandler == nil {
		return AssignmentHandlerAssign
	}
	return params.AssignmentHandler
}

// CountArraysAndMapsInConversionContext returns the number of arrays/maps which are in the conversion context.
// This is to aid in situations where there are deeply nested conversions (i.e. array of map of maps). In these contexts,
// temporary variables need to be declared to store intermediate conversion results.
func (params ConversionParameters) CountArraysAndMapsInConversionContext() int {
	result := 0
	for _, t := range params.ConversionContext {
		switch t.(type) {
		case *MapType:
			result += 1
		case *ArrayType:
			result += 1
		}
	}

	return result
}

func (params ConversionParameters) copy() ConversionParameters {
	result := params
	result.ConversionContext = append([]Type(nil), params.ConversionContext...)
	result.Locals = result.Locals.Clone()

	return result
}

type ConversionHandler func(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt

// TODO: There feels like overlap between this and the Storage Conversion Factories? Need further thinking to combine them?
// ConversionFunctionBuilder is used to build a function converting between two similar types.
// It has a set of built-in conversions and can be configured with additional conversions.
type ConversionFunctionBuilder struct {
	// TODO: Better way to let you fuss with this? How can you pick out what I've already put in here to overwrite it?
	conversions []ConversionHandler

	IdFactory             IdentifierFactory
	CodeGenerationContext *CodeGenerationContext
}

// NewConversionFunctionBuilder creates a new ConversionFunctionBuilder with the default conversions already added.
func NewConversionFunctionBuilder(idFactory IdentifierFactory, codeGenerationContext *CodeGenerationContext) *ConversionFunctionBuilder {
	return &ConversionFunctionBuilder{
		IdFactory:             idFactory,
		CodeGenerationContext: codeGenerationContext,
		conversions: []ConversionHandler{
			// Complex wrapper types checked first
			IdentityConvertComplexOptionalProperty,
			IdentityConvertComplexArrayProperty,
			IdentityConvertComplexMapProperty,

			// TODO: a flip function of some kind would be kinda nice (for source vs dest)
			IdentityAssignValidatedTypeDestination,
			IdentityAssignValidatedTypeSource,
			IdentityAssignPrimitiveType,
			AssignToOptional,
			AssignFromOptional,
			IdentityDeepCopyJSON,
			IdentityAssignTypeName,
		},
	}
}

// AddConversionHandlers adds the specified conversion handlers to the end of the conversion list.
func (builder *ConversionFunctionBuilder) AddConversionHandlers(conversionHandlers ...ConversionHandler) {
	builder.conversions = append(builder.conversions, conversionHandlers...)
}

// PrependConversionHandlers adds the specified conversion handlers to the beginning of the conversion list.
func (builder *ConversionFunctionBuilder) PrependConversionHandlers(conversionHandlers ...ConversionHandler) {
	builder.conversions = append(conversionHandlers, builder.conversions...)
}

// BuildConversion creates a conversion between the source and destination defined by params.
func (builder *ConversionFunctionBuilder) BuildConversion(params ConversionParameters) []dst.Stmt {
	for _, conversion := range builder.conversions {
		result := conversion(builder, params)
		if len(result) > 0 {
			return result
		}
	}

	types := builder.CodeGenerationContext.GetAllReachableTypes()
	msg := fmt.Sprintf(
		"don't know how to perform conversion for %s -> %s",
		DebugDescription(params.SourceType, types),
		DebugDescription(params.DestinationType, types))
	panic(msg)
}

// IdentityConvertComplexOptionalProperty handles conversion for optional properties with complex elements
// This function generates code that looks like this:
// 	if <source> != nil {
//		<code for producing result from destinationType.Element()>
//		<destination> = &<result>
//	}
func IdentityConvertComplexOptionalProperty(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	destinationType, ok := params.DestinationType.(*OptionalType)
	if !ok {
		return nil
	}

	sourceType, ok := params.SourceType.(*OptionalType)
	if !ok {
		return nil
	}

	tempVarIdent := params.Locals.CreateLocal(params.NameHint)

	innerStatements := builder.BuildConversion(
		ConversionParameters{
			Source:            astbuilder.Dereference(params.GetSource()),
			SourceType:        sourceType.Element(),
			Destination:       dst.NewIdent(tempVarIdent),
			DestinationType:   destinationType.Element(),
			NameHint:          params.NameHint,
			ConversionContext: append(params.ConversionContext, destinationType),
			AssignmentHandler: AssignmentHandlerDefine,
			Locals:            params.Locals.Clone(),
		})

	// Tack on the final assignment
	innerStatements = append(
		innerStatements,
		astbuilder.SimpleAssignment(
			params.GetDestination(),
			astbuilder.AddrOf(dst.NewIdent(tempVarIdent))))

	result := astbuilder.IfNotNil(
		params.GetSource(),
		innerStatements...)

	return astbuilder.Statements(result)
}

// IdentityConvertComplexArrayProperty handles conversion for array properties with complex elements
// This function generates code that looks like this:
// 	for _, item := range <source> {
//		<code for producing result from destinationType.Element()>
//		<destination> = append(<destination>, <result>)
//	}
func IdentityConvertComplexArrayProperty(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	destinationType, ok := params.DestinationType.(*ArrayType)
	if !ok {
		return nil
	}

	sourceType, ok := params.SourceType.(*ArrayType)
	if !ok {
		return nil
	}

	var results []dst.Stmt

	locals := params.Locals.Clone() // Loop variables are scoped inside the loop
	itemIdent := builder.CreateLocal(locals, "item", params.NameHint)

	depth := params.CountArraysAndMapsInConversionContext()
	destination := params.GetDestination()

	// Check what depth we're at to determine if we need to define an intermediate variable to hold the result
	// or if we'll be able to use the final destination directly.
	if depth > 0 {
		// TODO: The suffix here should maybe be configurable on the function builder?
		innerDestinationIdent := locals.CreateLocal(params.NameHint, "Temp")
		destination = dst.NewIdent(innerDestinationIdent)
		results = append(
			results,
			astbuilder.LocalVariableDeclaration(
				innerDestinationIdent,
				destinationType.AsType(builder.CodeGenerationContext),
				""))
	}

	result := &dst.RangeStmt{
		Key:   dst.NewIdent("_"),
		Value: dst.NewIdent(itemIdent),
		X:     params.GetSource(),
		Tok:   token.DEFINE,
		Body: &dst.BlockStmt{
			List: builder.BuildConversion(
				ConversionParameters{
					Source:            dst.NewIdent(itemIdent),
					SourceType:        sourceType.Element(),
					Destination:       dst.Clone(destination).(dst.Expr),
					DestinationType:   destinationType.Element(),
					NameHint:          itemIdent,
					ConversionContext: append(params.ConversionContext, destinationType),
					AssignmentHandler: astbuilder.AppendSlice,
					Locals:            locals,
				}),
		},
	}
	results = append(results, result)

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.AssignmentHandler != nil {
		results = append(results, params.AssignmentHandler(params.GetDestination(), dst.Clone(destination).(dst.Expr)))
	}

	return results
}

// IdentityConvertComplexMapProperty handles conversion for map properties with complex values.
// This function panics if the map keys are not primitive types.
// This function generates code that looks like this:
// 	if <source> != nil {
//		<destination> = make(map[<destinationType.KeyType()]<destinationType.ValueType()>)
//		for key, value := range <source> {
// 			<code for producing result from destinationType.ValueType()>
//			<destination>[key] = <result>
//		}
//	}
func IdentityConvertComplexMapProperty(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	destinationType, ok := params.DestinationType.(*MapType)
	if !ok {
		return nil
	}

	sourceType, ok := params.SourceType.(*MapType)
	if !ok {
		return nil
	}

	if _, ok := destinationType.KeyType().(*PrimitiveType); !ok {
		msg := fmt.Sprintf(
			"map had non-primitive key type: %s",
			DebugDescription(destinationType.KeyType(), nil))
		panic(msg)
	}

	depth := params.CountArraysAndMapsInConversionContext()

	locals := params.Locals.Clone() // Loop variables are scoped inside the loop

	keyIdent := builder.CreateLocal(locals, "key", params.NameHint)
	valueIdent := builder.CreateLocal(locals, "value", params.NameHint)

	nameHint := valueIdent
	destination := params.GetDestination()
	makeMapToken := token.ASSIGN

	// Check what depth we're at to determine if we need to define an intermediate variable to hold the result
	// or if we'll be able to use the final destination directly.
	if depth > 0 {
		innerDestinationIdent := locals.CreateLocal(params.NameHint, "Temp")
		destination = dst.NewIdent(innerDestinationIdent)
		makeMapToken = token.DEFINE
	}

	handler := func(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
		return astbuilder.InsertMap(lhs, dst.NewIdent(keyIdent), rhs)
	}

	keyTypeAst := destinationType.KeyType().AsType(builder.CodeGenerationContext)
	valueTypeAst := destinationType.ValueType().AsType(builder.CodeGenerationContext)

	makeMapStatement := astbuilder.AssignmentStatement(
		destination,
		makeMapToken,
		astbuilder.MakeMap(keyTypeAst, valueTypeAst))
	rangeStatement := &dst.RangeStmt{
		Key:   dst.NewIdent(keyIdent),
		Value: dst.NewIdent(valueIdent),
		X:     params.GetSource(),
		Tok:   token.DEFINE,
		Body: &dst.BlockStmt{
			List: builder.BuildConversion(
				ConversionParameters{
					Source:            dst.NewIdent(valueIdent),
					SourceType:        sourceType.ValueType(),
					Destination:       dst.Clone(destination).(dst.Expr),
					DestinationType:   destinationType.ValueType(),
					NameHint:          nameHint,
					ConversionContext: append(params.ConversionContext, destinationType),
					AssignmentHandler: handler,
					Locals:            locals,
				}),
		},
	}

	result := astbuilder.IfNotNil(
		params.GetSource(),
		makeMapStatement,
		rangeStatement)

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.AssignmentHandler != nil {
		result.Body.List = append(result.Body.List, params.AssignmentHandler(params.GetDestination(), dst.Clone(destination).(dst.Expr)))
	}

	return []dst.Stmt{result}
}

// IdentityAssignTypeName handles conversion for TypeName's that are the same
// Note that because this handler is dealing with TypeName's and not Optional<TypeName>, it is safe to
// perform a simple assignment rather than a copy.
// This function generates code that looks like this:
//	<destination> <assignmentHandler> <source>
func IdentityAssignTypeName(_ *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	destinationType, ok := params.DestinationType.(TypeName)
	if !ok {
		return nil
	}

	sourceType, ok := params.SourceType.(TypeName)
	if !ok {
		return nil
	}

	// Can only apply basic assignment for typeNames that are the same
	if !TypeEquals(sourceType, destinationType) {
		return nil
	}

	return []dst.Stmt{
		params.AssignmentHandlerOrDefault()(params.GetDestination(), params.GetSource()),
	}
}

// IdentityAssignPrimitiveType just assigns source to destination directly, no conversion needed.
// This function generates code that looks like this:
// <destination> <assignmentHandler> <source>
func IdentityAssignPrimitiveType(_ *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	if _, ok := params.DestinationType.(*PrimitiveType); !ok {
		return nil
	}

	if _, ok := params.SourceType.(*PrimitiveType); !ok {
		return nil
	}

	return []dst.Stmt{
		params.AssignmentHandlerOrDefault()(params.GetDestination(), params.GetSource()),
	}
}

// AssignToOptional assigns address of source to destination.
// This function generates code that looks like this, for simple conversions:
// <destination> <assignmentHandler> &<source>
//
// or:
// <destination>Temp := convert(<source>)
// <destination> <assignmentHandler> &<destination>Temp
func AssignToOptional(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	optDest, ok := params.DestinationType.(*OptionalType)
	if !ok {
		return nil
	}

	if TypeEquals(optDest.Element(), params.SourceType) {
		return []dst.Stmt{
			params.AssignmentHandlerOrDefault()(params.GetDestination(), astbuilder.AddrOf(params.GetSource())),
		}
	}

	// a more complex conversion is needed
	dstType := optDest.Element()
	tmpLocal := builder.CreateLocal(params.Locals, "temp", params.NameHint)

	conversion := builder.BuildConversion(
		ConversionParameters{
			Source:            params.Source,
			SourceType:        params.SourceType,
			Destination:       dst.NewIdent(tmpLocal),
			DestinationType:   dstType,
			NameHint:          tmpLocal,
			ConversionContext: nil,
			AssignmentHandler: nil,
			Locals:            params.Locals,
		})

	if len(conversion) == 0 {
		return nil // unable to build inner conversion
	}

	return astbuilder.Statements(
		astbuilder.LocalVariableDeclaration(tmpLocal, dstType.AsType(builder.CodeGenerationContext), ""),
		conversion,
		params.AssignmentHandlerOrDefault()(params.GetDestination(), astbuilder.AddrOf(dst.NewIdent(tmpLocal))))
}

// AssignFromOptional assigns address of source to destination.
// This function generates code that looks like this, for simple conversions:
// if (<source> != nil) {
//     <destination> <assignmentHandler> *<source>
// }
//
// or:
// if (<source> != nil) {
//     <destination>Temp := convert(*<source>)
//     <destination> <assignmentHandler> <destination>Temp
// }
func AssignFromOptional(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	optSrc, ok := params.SourceType.(*OptionalType)
	if !ok {
		return nil
	}

	if TypeEquals(optSrc.Element(), params.DestinationType) {
		return []dst.Stmt{
			astbuilder.IfNotNil(params.GetSource(),
				params.AssignmentHandlerOrDefault()(params.GetDestination(), astbuilder.Dereference(params.GetSource())),
			),
		}
	}

	// a more complex conversion is needed
	srcType := optSrc.Element()
	tmpLocal := builder.CreateLocal(params.Locals, "temp", params.NameHint)

	conversion := builder.BuildConversion(
		ConversionParameters{
			Source:            astbuilder.Dereference(params.GetSource()),
			SourceType:        srcType,
			Destination:       dst.NewIdent(tmpLocal),
			DestinationType:   params.DestinationType,
			NameHint:          tmpLocal,
			ConversionContext: nil,
			AssignmentHandler: nil,
			Locals:            params.Locals,
		})

	if len(conversion) == 0 {
		return nil // unable to build inner conversion
	}

	var result []dst.Stmt
	result = append(result, astbuilder.LocalVariableDeclaration(tmpLocal, params.DestinationType.AsType(builder.CodeGenerationContext), ""))
	result = append(result, conversion...)
	result = append(result, params.AssignmentHandlerOrDefault()(params.GetDestination(), dst.NewIdent(tmpLocal)))
	return result
}

// IdentityAssignValidatedTypeDestination generates an assignment to the underlying validated type Element
func IdentityAssignValidatedTypeDestination(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	validatedType, ok := params.DestinationType.(*ValidatedType)
	if !ok {
		return nil
	}

	// pass through to underlying type
	params = params.WithDestinationType(validatedType.ElementType())
	return builder.BuildConversion(params)
}

// IdentityAssignValidatedTypeSource generates an assignment to the underlying validated type Element
func IdentityAssignValidatedTypeSource(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	validatedType, ok := params.SourceType.(*ValidatedType)
	if !ok {
		return nil
	}

	// pass through to underlying type
	params = params.WithSourceType(validatedType.ElementType())
	return builder.BuildConversion(params)
}

// IdentityDeepCopyJSON special cases copying JSON-type fields to call the DeepCopy method.
// It generates code that looks like:
//     <destination> = *<source>.DeepCopy()
func IdentityDeepCopyJSON(builder *ConversionFunctionBuilder, params ConversionParameters) []dst.Stmt {
	if !TypeEquals(params.DestinationType, JSONType) {
		return nil
	}

	newSource := astbuilder.Dereference(
		&dst.CallExpr{
			Fun:  astbuilder.Selector(params.GetSource(), "DeepCopy"),
			Args: []dst.Expr{},
		})

	return []dst.Stmt{
		params.AssignmentHandlerOrDefault()(params.GetDestination(), newSource),
	}
}

// AssignmentHandlerDefine is an assignment handler for definitions, using :=
func AssignmentHandlerDefine(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return astbuilder.AssignmentStatement(lhs, token.DEFINE, rhs)
}

// AssignmentHandlerAssign is an assignment handler for standard assignments to existing variables, using =
func AssignmentHandlerAssign(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return astbuilder.SimpleAssignment(lhs, rhs)
}

// CreateLocal creates an unused local variable name.
// Names are chosen according to the following rules:
//   1. If there is no local variable with the <suffix> name, use that.
//   2. If there is a local variable with the <suffix> name, create a variable name <nameHint><suffix>.
// In the case that <nameHint><suffix> is also taken append numbers to the end in standard KnownLocalsSet fashion.
// Note that this function trims numbers on the right hand side of nameHint, so a nameHint of "item1" will get a local
// variable named item<suffix>.
func (builder *ConversionFunctionBuilder) CreateLocal(locals *KnownLocalsSet, suffix string, nameHint string) string {
	ident := suffix
	if locals.HasName(ident) || ident == "" {
		// Trim any trailing numbers so that we don't end up with ident1111
		// Note that this can end up trimming trailing digits on fields that have them (i.e. "loop1" might become "loop").
		// This is ok as it doesn't hurt anything and helps avoid "loop12" (for loop1 number 2).
		trimmedNameHint := strings.TrimRight(nameHint, "0123456789")
		ident = locals.CreateLocal(trimmedNameHint, suffix)
	} else {
		locals.Add(ident)
	}

	return ident
}
