/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"go/token"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// PropertyConversion generates the AST for a given property conversion.
// reader is an expression to read the original value.
// writer is a function that accepts an expression for reading a value and creates one or more
// statements to write that value.
// Both of these might be complex expressions, possibly involving indexing into arrays or maps.
type PropertyConversion func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt

// PropertyConversionFactory represents factory methods that can be used to create a PropertyConversion for a specific
// pair of properties
// source is the property conversion endpoint that will be read
// destination is the property conversion endpoint that will be written
// ctx contains additional information that may be needed when creating the property conversion
type PropertyConversionFactory func(
	source *TypedConversionEndpoint,
	destination *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion

// A list of all known type conversion factory methods
var propertyConversionFactories []PropertyConversionFactory

func init() {
	propertyConversionFactories = []PropertyConversionFactory{
		// Primitive types and aliases
		assignPrimitiveFromPrimitive,
		assignAliasedPrimitiveFromAliasedPrimitive,
		// Collection Types
		assignArrayFromArray,
		assignMapFromMap,
		// Enumerations
		assignEnumFromEnum,
		assignPrimitiveFromEnum,
		// Complex object types
		assignObjectFromObject,
		// Known types
		copyKnownType(astmodel.KnownResourceReferenceTypeName, "Copy", returnsValue),
		copyKnownType(astmodel.ResourceReferenceTypeName, "Copy", returnsValue),
		copyKnownType(astmodel.JSONTypeName, "DeepCopy", returnsReference),
		// Meta-conversions
		assignFromOptional, // Must go before assignToOptional so we generate the right zero values
		assignToOptional,
		assignToEnumeration,
		assignFromAliasedPrimitive,
		assignToAliasedPrimitive,
	}
}

// CreateTypeConversion tries to create a type conversion between the two provided types, using
// all of the available type conversion functions in priority order to do so.
//
// The method works by considering the conversion requested by sourceEndpoint & destinationEndpoint,
// with recursive calls breaking the conversion down into multiple steps that are then combined.
//
// Example:
//
// CreateTypeConversion() is called to create a conversion from an optional string to an optional
// Sku, where Sku is a new type based on string:
//
// source *string => destination *Sku
//
// assuming
//     type Sku string
//
// assignFromOptional can handle the optionality of sourceEndpoint and makes a recursive call
// to CreateTypeConversion() with the simpler target:
//
// source string => destination *Sku
//
//     assignToOptional can handle the optionality of destinationEndpoint and makes a recursive
//     call to CreateTypeConversion() with a simpler target:
//
//     source string => destination Sku
//
//         assignToAliasedPrimitive can handle the type conversion of string to Sku, and makes
//         a recursive call to CreateTypeConversion() with a simpler target:
//
//         source string => destination string
//
//             assignPrimitiveFromPrimitive can handle primitive values, and generates a
//             conversion that does a simple assignment:
//
//             destination = source
//
//         assignToAliasedPrimitive injects the necessary type conversion:
//
//         destination = Sku(source)
//
//     assignToOptional injects a local variable and takes it's address
//
//     sku := Sku(source)
//     destination = &sku
//
// finally, assignFromOptional injects the check to see if we have a value to assign in the
// first place, assigning a suitable zero value if we don't:
//
// if source != nil {
//     sku := Sku(source)
//     destination := &sku
// } else {
//     destination := ""
// }
//
// TODO: Make this internal
func CreateTypeConversion(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) (PropertyConversion, error) {
	for _, f := range propertyConversionFactories {
		result := f(sourceEndpoint, destinationEndpoint, conversionContext)
		if result != nil {
			return result, nil
		}
	}

	// No conversion found, we need to generate a useful error message

	var debugDescriptionOfDestination strings.Builder
	destinationEndpoint.Type().WriteDebugDescription(&debugDescriptionOfDestination, conversionContext.Types())

	var debugDescriptionOfSource strings.Builder
	sourceEndpoint.Type().WriteDebugDescription(&debugDescriptionOfSource, conversionContext.Types())

	err := errors.Errorf(
		"no conversion found to assign %q from %q",
		debugDescriptionOfDestination.String(),
		debugDescriptionOfSource.String())

	return nil, err
}

func NameOfPropertyAssignmentFunction(name astmodel.TypeName, direction Direction, idFactory astmodel.IdentifierFactory) string {
	nameOfOtherType := idFactory.CreateIdentifier(name.Name(), astmodel.Exported)
	return direction.SelectString(
		"AssignPropertiesFrom"+nameOfOtherType,
		"AssignPropertiesTo"+nameOfOtherType)
}

// assignToOptional will generate a conversion where the destination is optional, if the
// underlying type of the destination is compatible with the source.
//
// <destination> = &<source>
//
func assignToOptional(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to be optional
	destinationOptional, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type())
	if !destinationIsOptional {
		// Destination is not optional
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	unwrappedEndpoint := destinationEndpoint.WithType(destinationOptional.Element())
	conversion, _ := CreateTypeConversion(sourceEndpoint, unwrappedEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	local := destinationEndpoint.CreateLocal("", "Temp")

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// Create a writer that uses the address of the passed expression
		// If expr isn't a plain identifier (implying a local variable), we introduce one
		// This both allows us to avoid aliasing and complies with Go language semantics
		addrOfWriter := func(expr dst.Expr) []dst.Stmt {
			if _, ok := expr.(*dst.Ident); ok {
				return writer(astbuilder.AddrOf(expr))
			}

			assignment := astbuilder.SimpleAssignment(
				dst.NewIdent(local),
				token.DEFINE,
				expr)

			writing := writer(astbuilder.AddrOf(dst.NewIdent(local)))

			return astbuilder.Statements(assignment, writing)
		}

		return conversion(reader, addrOfWriter, generationContext)
	}
}

// assignFromOptional will handle the case where the source type may be missing (nil)
//
// <original> := <source>
// if <original> != nil {
//    <destination> = *<original>
// } else {
//    <destination> = <zero>
// }
func assignFromOptional(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be optional
	sourceOptional, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type())
	if !sourceIsOptional {
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	// We supply a nested context as any code generated by the conversion will be within the if
	// statement (a nested scope) and we don't want any variables declared in that scope to leak
	// out elsewhere.
	unwrappedEndpoint := sourceEndpoint.WithType(sourceOptional.Element())
	conversion, _ := CreateTypeConversion(
		unwrappedEndpoint,
		destinationEndpoint,
		conversionContext.NestedContext())
	if conversion == nil {
		return nil
	}

	local := sourceEndpoint.CreateLocal("", "Read")

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		var cacheOriginal dst.Stmt
		var actualReader dst.Expr

		// If the value we're reading is a local or a field, it's cheap to read and we can skip
		// using a local (which makes the generated code easier to read). In other cases, we want
		// to cache the value in a local to avoid repeating any expensive conversion.

		switch reader.(type) {
		case *dst.Ident, *dst.SelectorExpr:
			// reading a local variable or a field
			cacheOriginal = nil
			actualReader = reader
		default:
			// Something else, so we cache the original
			cacheOriginal = astbuilder.SimpleAssignment(
				dst.NewIdent(local),
				token.DEFINE,
				reader)
			actualReader = dst.NewIdent(local)
		}

		checkForNil := astbuilder.NotEqual(actualReader, astbuilder.Nil())

		// If we have a value, need to convert it to our destination type
		writeActualValue := conversion(
			astbuilder.Dereference(actualReader),
			writer,
			generationContext)

		writeZeroValue := writer(
			destinationEndpoint.Type().AsZero(conversionContext.Types(), generationContext))

		stmt := &dst.IfStmt{
			Cond: checkForNil,
			Body: astbuilder.StatementBlock(writeActualValue...),
			Else: astbuilder.StatementBlock(writeZeroValue...),
		}

		return astbuilder.Statements(cacheOriginal, stmt)
	}
}

// assignToEnumeration will generate a conversion where the destination is an enumeration if
// the source is type compatible with the base type of the enumeration
//
// <destination> = <enumeration-cast>(<source>)
//
func assignToEnumeration(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to NOT be optional
	_, dstIsOpt := astmodel.AsOptionalType(destinationEndpoint.Type())
	if dstIsOpt {
		// Destination is not optional
		return nil
	}

	// Require destination to be an enumeration
	dstName, dstType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	dstEnum, ok := astmodel.AsEnumType(dstType)
	if !ok {
		return nil
	}

	// Require a conversion between the base type of the enumeration and our source
	dstEp := destinationEndpoint.WithType(dstEnum.BaseType())
	conversion, _ := CreateTypeConversion(sourceEndpoint, dstEp, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		convertingWriter := func(expr dst.Expr) []dst.Stmt {
			cast := &dst.CallExpr{
				Fun:  dstName.AsType(generationContext),
				Args: []dst.Expr{expr},
			}
			return writer(cast)
		}

		return conversion(
			reader,
			convertingWriter,
			generationContext)
	}
}

// assignPrimitiveFromPrimitive will generate a direct assignment if both types have the
// same primitive type and are not optional
//
// <destination> = <source>
//
func assignPrimitiveFromPrimitive(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	_ *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be a primitive type
	sourcePrimitive, sourceIsPrimitive := astmodel.AsPrimitiveType(sourceEndpoint.Type())
	if !sourceIsPrimitive {
		return nil
	}

	// Require destination to be a primitive type
	destinationPrimitive, destinationIsPrimitive := astmodel.AsPrimitiveType(destinationEndpoint.Type())
	if !destinationIsPrimitive {
		return nil
	}

	// Require both properties to have the same primitive type
	if !sourcePrimitive.Equals(destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		return writer(reader)
	}
}

// assignAliasedPrimitiveFromAliasedPrimitive will generate a direct assignment if both
// types have the same underlying primitive type and are not optional
//
// <destination> = <cast>(<source>)
//
func assignAliasedPrimitiveFromAliasedPrimitive(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be a name that resolves to a primitive type
	_, sourceType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	sourcePrimitive, sourceIsPrimitive := astmodel.AsPrimitiveType(sourceType)
	if !sourceIsPrimitive {
		return nil
	}

	// Require destination to be a name the resolves to a primitive type
	destinationName, destinationType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	destinationPrimitive, destinationIsPrimitive := astmodel.AsPrimitiveType(destinationType)
	if !destinationIsPrimitive {
		return nil
	}

	// Require both properties to have the same primitive type
	if !sourcePrimitive.Equals(destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		return writer(&dst.CallExpr{
			Fun:  destinationName.AsType(generationContext),
			Args: []dst.Expr{reader},
		})
	}
}

// assignFromAliasedPrimitive will convert an alias of a primitive type into that primitive
// type as long as it is not optional and we can find a conversion to consume that primitive value
func assignFromAliasedPrimitive(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require source to be a name that resolves to a primitive type
	_, sourceType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	sourcePrimitive, sourceIsPrimitive := astmodel.AsPrimitiveType(sourceType)
	if !sourceIsPrimitive {
		return nil
	}

	// Require a conversion for the underlying type
	primitiveEndpoint := sourceEndpoint.WithType(sourcePrimitive)
	conversion, _ := CreateTypeConversion(primitiveEndpoint, destinationEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		actualReader := &dst.CallExpr{
			Fun:  sourcePrimitive.AsType(generationContext),
			Args: []dst.Expr{reader},
		}

		return conversion(actualReader, writer, generationContext)
	}
}

// assignToAliasedPrimitive will convert a primitive value into the aliased type as long as it
// is not optional and we can find a conversion to give us the primitive type.
//
// <destination> = <cast>(<source>)
//
func assignToAliasedPrimitive(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to be non-optional
	if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require destination to be a name the resolves to a primitive type
	destinationName, destinationType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	destinationPrimitive, destinationIsPrimitive := astmodel.AsPrimitiveType(destinationType)
	if !destinationIsPrimitive {
		return nil
	}

	// Require a conversion for the underlying type
	primitiveEndpoint := sourceEndpoint.WithType(destinationPrimitive)
	conversion, _ := CreateTypeConversion(sourceEndpoint, primitiveEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		actualWriter := func(expr dst.Expr) []dst.Stmt {
			castToAlias := &dst.CallExpr{
				Fun:  destinationName.AsType(generationContext),
				Args: []dst.Expr{expr},
			}

			return writer(castToAlias)
		}

		return conversion(reader, actualWriter, generationContext)
	}
}

// assignArrayFromArray will generate a code fragment to populate an array, assuming the
// underlying types of the two arrays are compatible
//
// <arr> := make([]<type>, len(<reader>))
// for <index>, <value> := range <reader> {
//     // Shadow the loop variable to avoid aliasing
//     <value> := <value>
//     <arr>[<index>] := <value> // Or other conversion as required
// }
// <writer> = <arr>
//
func assignArrayFromArray(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be an array type
	sourceArray, sourceIsArray := astmodel.AsArrayType(sourceEndpoint.Type())
	if !sourceIsArray {
		return nil
	}

	// Require destination to be an array type
	destinationArray, destinationIsArray := astmodel.AsArrayType(destinationEndpoint.Type())
	if !destinationIsArray {
		return nil
	}

	// Require a conversion between the array types
	// We supply a nested context as any code generated by the conversion will be within the loop
	// (a nested scope) and we don't want any variables declared in that scope to leak out elsewhere.
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceArray.Element())
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationArray.Element())
	conversion, _ := CreateTypeConversion(
		unwrappedSourceEndpoint,
		unwrappedDestinationEndpoint,
		conversionContext.NestedContext())
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// We create three obviously related identifiers to use for the conversion
		itemId := sourceEndpoint.CreateLocal("Item")
		indexId := sourceEndpoint.CreateLocal("Index")
		tempId := sourceEndpoint.CreateLocal("List")

		declaration := astbuilder.SimpleAssignment(
			dst.NewIdent(tempId),
			token.DEFINE,
			astbuilder.MakeList(destinationArray.AsType(generationContext), astbuilder.CallFunc("len", reader)))

		writeToElement := func(expr dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					&dst.IndexExpr{
						X:     dst.NewIdent(tempId),
						Index: dst.NewIdent(indexId),
					},
					token.ASSIGN,
					expr),
			}
		}

		avoidAliasing := astbuilder.SimpleAssignment(dst.NewIdent(itemId), token.DEFINE, dst.NewIdent(itemId))
		avoidAliasing.Decs.Start.Append("// Shadow the loop variable to avoid aliasing")
		avoidAliasing.Decs.Before = dst.NewLine

		loopBody := astbuilder.Statements(
			avoidAliasing,
			conversion(dst.NewIdent(itemId), writeToElement, generationContext))

		assign := writer(dst.NewIdent(tempId))
		loop := astbuilder.IterateOverListWithIndex(indexId, itemId, reader, loopBody...)
		return astbuilder.Statements(declaration, loop, assign)
	}
}

// assignMapFromMap will generate a code fragment to populate an array, assuming the
// underlying types of the two arrays are compatible
//
// <map> := make(map[<key>]<type>)
// for key, <item> := range <reader> {
//     <map>[<key>] := <item>
// }
// <writer> = <map>
//
func assignMapFromMap(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be a map
	sourceMap, sourceIsMap := astmodel.AsMapType(sourceEndpoint.Type())
	if !sourceIsMap {
		// Source is not a map
		return nil
	}

	// Require destination to be a map
	destinationMap, destinationIsMap := astmodel.AsMapType(destinationEndpoint.Type())
	if !destinationIsMap {
		// Destination is not a map
		return nil
	}

	// Require map keys to be identical
	if !sourceMap.KeyType().Equals(destinationMap.KeyType()) {
		// Keys are different types
		return nil
	}

	// Require a conversion between the map items
	// We supply a nested context as any code generated by the conversion will be within the loop
	// (a nested scope) and we don't want any variables declared in that scope to leak out elsewhere.
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceMap.ValueType())
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationMap.ValueType())
	conversion, _ := CreateTypeConversion(
		unwrappedSourceEndpoint,
		unwrappedDestinationEndpoint,
		conversionContext.NestedContext())
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// We create three obviously related identifiers to use for the conversion
		itemId := sourceEndpoint.CreateLocal("Value")
		keyId := sourceEndpoint.CreateLocal("Key")
		tempId := sourceEndpoint.CreateLocal("Map")

		declaration := astbuilder.SimpleAssignment(
			dst.NewIdent(tempId),
			token.DEFINE,
			astbuilder.MakeMap(destinationMap.KeyType().AsType(generationContext), destinationMap.ValueType().AsType(generationContext)))

		assignToItem := func(expr dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					&dst.IndexExpr{
						X:     dst.NewIdent(tempId),
						Index: dst.NewIdent(keyId),
					},
					token.ASSIGN,
					expr),
			}
		}

		avoidAliasing := astbuilder.SimpleAssignment(dst.NewIdent(itemId), token.DEFINE, dst.NewIdent(itemId))
		avoidAliasing.Decs.Start.Append("// Shadow the loop variable to avoid aliasing")
		avoidAliasing.Decs.Before = dst.NewLine

		loopBody := astbuilder.Statements(
			avoidAliasing,
			conversion(dst.NewIdent(itemId), assignToItem, generationContext))

		assign := writer(dst.NewIdent(tempId))
		loop := astbuilder.IterateOverMapWithValue(keyId, itemId, reader, loopBody...)
		return astbuilder.Statements(declaration, loop, assign)
	}
}

// assignEnumFromEnum will generate a conversion if both types have the same underlying
// primitive type and neither source nor destination is optional
//
// <local> = <baseType>(<source>)
// <destination> = <enum>(<local>)
//
// We don't technically need this one, but it generates nicer code because it bypasses an unnecessary cast.
func assignEnumFromEnum(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be an enumeration
	_, sourceType, sourceFound := conversionContext.ResolveType(sourceEndpoint.Type())
	if !sourceFound {
		return nil
	}
	sourceEnum, sourceIsEnum := astmodel.AsEnumType(sourceType)
	if !sourceIsEnum {
		return nil
	}

	// Require destination to be an enumeration
	destinationName, destinationType, destinationFound := conversionContext.ResolveType(destinationEndpoint.Type())
	if !destinationFound {
		return nil
	}
	destinationEnum, destinationIsEnum := astmodel.AsEnumType(destinationType)
	if !destinationIsEnum {
		return nil
	}

	// Require enumerations to have the same base types
	if !sourceEnum.BaseType().Equals(destinationEnum.BaseType()) {
		return nil
	}

	local := destinationEndpoint.CreateLocal("", "As"+destinationName.Name(), "Value")
	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, ctx *astmodel.CodeGenerationContext) []dst.Stmt {
		result := []dst.Stmt{
			astbuilder.SimpleAssignment(
				dst.NewIdent(local),
				token.DEFINE,
				astbuilder.CallFunc(destinationName.Name(), reader)),
		}

		result = append(result, writer(dst.NewIdent(local))...)
		return result
	}
}

// assignPrimitiveFromEnum will generate a conversion from an enumeration if the
// destination has the underlying base type of the enumeration and neither source nor destination
// is optional
//
// <local> = <baseType>(<source>)
// <destination> = <enum>(<local>)
//
func assignPrimitiveFromEnum(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, srcOpt := astmodel.AsOptionalType(sourceEndpoint.Type()); srcOpt {
		return nil
	}

	// Require destination to be non-optional
	if _, dstOpt := astmodel.AsOptionalType(destinationEndpoint.Type()); dstOpt {
		return nil
	}

	// Require source to be an enumeration
	_, srcType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	srcEnum, srcIsEnum := astmodel.AsEnumType(srcType)
	if !srcIsEnum {
		return nil
	}

	// Require destination to be a primitive type
	dstPrimitive, ok := astmodel.AsPrimitiveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}

	// Require enumeration to have the destination as base type
	if !srcEnum.BaseType().Equals(dstPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, ctx *astmodel.CodeGenerationContext) []dst.Stmt {
		return writer(astbuilder.CallFunc(dstPrimitive.Name(), reader))
	}
}

// assignObjectFromObject will generate a conversion if both properties are TypeNames
// referencing ObjectType definitions and neither property is optional
//
// For ConvertFrom:
//
// var <local> <destinationType>
// err := <local>.ConvertFrom(<source>)
// if err != nil {
//     return errors.Wrap(err, "while calling <local>.ConvertFrom(<source>)")
// }
// <destination> = <local>
//
// For ConvertTo:
//
// var <local> <destinationType>
// err := <source>.ConvertTo(&<local>)
// if err != nil {
//     return errors.Wrap(err, "while calling <local>.ConvertTo(<source>)")
// }
// <destination> = <local>
//
func assignObjectFromObject(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be the name of an object
	sourceName, sourceType, sourceFound := conversionContext.ResolveType(sourceEndpoint.Type())
	if !sourceFound {
		return nil
	}
	if _, sourceIsObject := astmodel.AsObjectType(sourceType); !sourceIsObject {
		return nil
	}

	// Require destination to be the name of an object
	destinationName, destinationType, destinationFound := conversionContext.ResolveType(destinationEndpoint.Type())
	if !destinationFound {
		return nil
	}
	_, destinationIsObject := astmodel.AsObjectType(destinationType)
	if !destinationIsObject {
		return nil
	}

	copyVar := destinationEndpoint.CreateLocal()

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// We have to do this at render time in order to ensure the first conversion generated
		// declares 'err', not a later one
		tok := token.ASSIGN
		if conversionContext.TryCreateLocal("err") {
			tok = token.DEFINE
		}

		localId := dst.NewIdent(copyVar)
		errLocal := dst.NewIdent("err")

		errorsPackageName := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

		declaration := astbuilder.LocalVariableDeclaration(copyVar, createTypeDeclaration(destinationName, generationContext), "")

		// If our reader is a dereference, we strip that off (because we need a pointer), else we
		// take the address of it
		var actualReader dst.Expr
		if deref, ok := reader.(*dst.UnaryExpr); ok && deref.Op == token.MUL {
			actualReader = deref.X
		} else {
			actualReader = astbuilder.AddrOf(reader)
		}

		functionName := NameOfPropertyAssignmentFunction(sourceName, conversionContext.direction, conversionContext.idFactory)

		var conversion dst.Stmt
		if destinationName.PackageReference.Equals(generationContext.CurrentPackage()) {
			// Destination is our current type
			conversion = astbuilder.SimpleAssignment(
				errLocal,
				tok,
				astbuilder.CallExpr(localId, functionName, actualReader))
		} else {
			// Destination is another type
			conversion = astbuilder.SimpleAssignment(
				errLocal,
				tok,
				astbuilder.CallExpr(reader, functionName, astbuilder.AddrOf(localId)))
		}

		checkForError := astbuilder.ReturnIfNotNil(
			errLocal,
			astbuilder.WrappedErrorf(
				errorsPackageName,
				"populating %s from %s, calling %s()",
				destinationEndpoint.Name(), sourceEndpoint.Name(), functionName))

		assignment := writer(dst.NewIdent(copyVar))
		return astbuilder.Statements(declaration, conversion, checkForError, assignment)
	}
}

// assignKnownType will generate an assignment if both types have the specified TypeName
//
// <destination> = <source>
//
// TODO: Make this internal once referenced ;-)
func AssignKnownType(name astmodel.TypeName) func(*TypedConversionEndpoint, *TypedConversionEndpoint, *PropertyConversionContext) PropertyConversion {
	return func(sourceEndpoint *TypedConversionEndpoint, destinationEndpoint *TypedConversionEndpoint, _ *PropertyConversionContext) PropertyConversion {
		// Require source to be non-optional
		if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
			return nil
		}

		// Require destination to be non-optional
		if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
			return nil
		}

		// Require source to be a named type
		sourceName, sourceIsName := astmodel.AsTypeName(sourceEndpoint.Type())
		if !sourceIsName {
			return nil
		}

		// Require destination to be a named type
		destinationName, destinationIsName := astmodel.AsTypeName(destinationEndpoint.Type())
		if !destinationIsName {
			return nil
		}

		// Require source to be our specific type
		if !sourceName.Equals(name) {
			return nil
		}

		// Require destination to be our specific type
		if !destinationName.Equals(name) {
			return nil
		}

		return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
			return writer(reader)
		}
	}
}

type knownTypeMethodReturn int

const (
	returnsReference = 0
	returnsValue     = 1
)

// copyKnownType will generate an assignment with the results of a call on the specified TypeName
//
// <destination> = <source>.<methodName>()
//
func copyKnownType(name astmodel.TypeName, methodName string, returnKind knownTypeMethodReturn) func(*TypedConversionEndpoint, *TypedConversionEndpoint, *PropertyConversionContext) PropertyConversion {
	return func(sourceEndpoint *TypedConversionEndpoint, destinationEndpoint *TypedConversionEndpoint, _ *PropertyConversionContext) PropertyConversion {
		// Require source to be non-optional
		if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
			return nil
		}

		// Require destination to be non-optional
		if _, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
			return nil
		}

		// Require source to be a named type
		sourceName, sourceIsName := astmodel.AsTypeName(sourceEndpoint.Type())
		if !sourceIsName {
			return nil
		}

		// Require destination to be a named type
		destinationName, destinationIsName := astmodel.AsTypeName(destinationEndpoint.Type())
		if !destinationIsName {
			return nil
		}

		// Require source to be our specific type
		if !sourceName.Equals(name) {
			return nil
		}

		// Require destination to be our specific type
		if !destinationName.Equals(name) {
			return nil
		}

		return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
			// If our writer is dereferencing a value, skip that as we don't need to dereference before a method call
			if unary, ok := astbuilder.AsDereference(reader); ok {
				reader = unary.X
			}

			if returnKind == returnsReference {
				// If the copy method returns a ptr, we need to dereference
				// This dereference is always safe because we ensured that both source and destination are always
				// non optional. The handler assignToOptional() should do the right thing when this happens.
				return writer(astbuilder.Dereference(astbuilder.CallExpr(reader, methodName)))
			}

			return writer(astbuilder.CallExpr(reader, methodName))
		}
	}
}

func createTypeDeclaration(name astmodel.TypeName, generationContext *astmodel.CodeGenerationContext) dst.Expr {
	if name.PackageReference.Equals(generationContext.CurrentPackage()) {
		return dst.NewIdent(name.Name())
	}

	packageName := generationContext.MustGetImportedPackageName(name.PackageReference)
	return astbuilder.Selector(dst.NewIdent(packageName), name.Name())
}
