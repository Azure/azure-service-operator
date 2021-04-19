/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/dave/dst"
	"github.com/pkg/errors"
)

// StorageTypeConversion generates the AST for a given conversion.
// reader is an expression to read the original value.
// writer is a function that accepts an expression for reading a value and creates one or more
// statements to write that value.
// Both of these might be complex expressions, possibly involving indexing into arrays or maps.
type StorageTypeConversion func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt

// StorageTypeConversionFactory represents factory methods that can be used to create StorageTypeConversions
// for a specific pair of types
// source is the endpoint that will be read
// destination is the endpoint that will be written
// ctx contains additional information that may be needed when creating a conversion
type StorageTypeConversionFactory func(
	source *StorageConversionEndpoint,
	destination *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion

// A list of all known type conversion factory methods
var typeConversionFactories []StorageTypeConversionFactory

func init() {
	typeConversionFactories = []StorageTypeConversionFactory{
		// Primitive types and aliases
		assignPrimitiveTypeFromPrimitiveType,
		assignAliasedPrimitiveTypeFromAliasedPrimitiveType,
		// Collection Types
		assignArrayFromArray,
		assignMapFromMap,
		// Enumerations
		assignEnumTypeFromEnumType,
		assignPrimitiveTypeFromEnumType,
		// Complex object types
		assignObjectTypeFromObjectType,
		// Known types
		assignKnownReferenceFromKnownReference,
		// Meta-conversions
		assignFromOptionalType, // Must go before assignToOptionalType so we generate the right zero values
		assignToOptionalType,
		assignToEnumerationType,
		assignFromAliasedPrimitiveType,
		assignToAliasedPrimitiveType,
	}
}

// createTypeConversion tries to create a type conversion between the two provided types, using
// all of the available type conversion functions in priority order to do so.
//
// The method works by considering the conversion requested by sourceEndpoint & destinationEndpoint,
// with recursive calls breaking the conversion down into multiple steps that are then combined.
//
// Example:
//
// createTypeConversion() is called to create a conversion from an optional string to an optional
// Sku, where Sku is a new type based on string:
//
// source *string => destination *Sku
//
// assuming
//     type Sku string
//
// assignFromOptionalType can handle the optionality of sourceEndpoint and makes a recursive call
// to createTypeConversion() with the simpler target:
//
// source string => destination *Sku
//
//     assignToOptionalType can handle the optionality of destinationEndpoint and makes a recursive
//     call to createTypeConversion() with a simpler target:
//
//     source string => destination Sku
//
//         assignToAliasedPrimitiveType can handle the type conversion of string to Sku, and makes
//         a recursive call to createTypeConversion() with a simpler target:
//
//         source string => destination string
//
//             assignPrimitiveTypeFromPrimitiveType can handle primitive values, and generates a
//             conversion that does a simple assignment:
//
//             destination = source
//
//         assignToAliasedPrimitiveType injects the necessary type conversion:
//
//         destination = Sku(source)
//
//     assignToOptionalType injects a local variable and takes it's address
//
//     sku := Sku(source)
//     destination = &sku
//
// finally, assignFromOptionalType injects the check to see if we have a value to assign in the
// first place, assigning a suitable zero value if we don't:
//
// if source != nil {
//     sku := Sku(source)
//     destination := &sku
// } else {
//     destination := ""
// }
//
func createTypeConversion(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) (StorageTypeConversion, error) {
	for _, f := range typeConversionFactories {
		result := f(sourceEndpoint, destinationEndpoint, conversionContext)
		if result != nil {
			return result, nil
		}
	}

	err := errors.Errorf(
		"no conversion found to assign %q from %q",
		destinationEndpoint.name,
		sourceEndpoint.name)

	return nil, err
}

// assignToOptionalType will generate a conversion where the destination is optional, if the
// underlying type of the destination is compatible with the source.
//
// <destination> = &<source>
//
func assignToOptionalType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require destination to be optional
	destinationOptional, destinationIsOptional := AsOptionalType(destinationEndpoint.Type())
	if !destinationIsOptional {
		// Destination is not optional
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	unwrappedEndpoint := destinationEndpoint.WithType(destinationOptional.element)
	conversion, _ := createTypeConversion(sourceEndpoint, unwrappedEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	local := destinationEndpoint.CreateLocal("", "Temp")

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
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

// assignFromOptionalType will handle the case where the source type may be missing (nil)
//
// <original> := <source>
// if <original> != nil {
//    <destination> = *<original>
// } else {
//    <destination> = <zero>
// }
func assignFromOptionalType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be optional
	sourceOptional, sourceIsOptional := AsOptionalType(sourceEndpoint.Type())
	if !sourceIsOptional {
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	unwrappedEndpoint := sourceEndpoint.WithType(sourceOptional.element)
	conversion, _ := createTypeConversion(unwrappedEndpoint, destinationEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	local := sourceEndpoint.CreateLocal("", "Read")

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {

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

		checkForNil := astbuilder.NotEqual(actualReader, dst.NewIdent("nil"))

		// If we have a value, need to convert it to our destination type
		writeActualValue := conversion(
			astbuilder.Dereference(actualReader),
			writer,
			generationContext)

		writeZeroValue := writer(
			destinationEndpoint.Type().AsZero(conversionContext.types, generationContext))

		stmt := &dst.IfStmt{
			Cond: checkForNil,
			Body: astbuilder.StatementBlock(writeActualValue...),
			Else: astbuilder.StatementBlock(writeZeroValue...),
		}

		return astbuilder.Statements(cacheOriginal, stmt)
	}
}

// assignToEnumerationType will generate a conversion where the destination is an enumeration if
// the source is type compatible with the base type of the enumeration
//
// <destination> = <enumeration-cast>(<source>)
//
func assignToEnumerationType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require destination to NOT be optional
	_, dstIsOpt := AsOptionalType(destinationEndpoint.Type())
	if dstIsOpt {
		// Destination is not optional
		return nil
	}

	// Require destination to be an enumeration
	dstName, dstType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	dstEnum, ok := AsEnumType(dstType)
	if !ok {
		return nil
	}

	// Require a conversion between the base type of the enumeration and our source
	dstEp := destinationEndpoint.WithType(dstEnum.baseType)
	conversion, _ := createTypeConversion(sourceEndpoint, dstEp, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
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

// assignPrimitiveTypeFromPrimitiveType will generate a direct assignment if both types have the
// same primitive type and are not optional
//
// <destination> = <source>
//
func assignPrimitiveTypeFromPrimitiveType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	_ *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be a primitive type
	sourcePrimitive, sourceIsPrimitive := AsPrimitiveType(sourceEndpoint.Type())
	if !sourceIsPrimitive {
		return nil
	}

	// Require destination to be a primitive type
	destinationPrimitive, destinationIsPrimitive := AsPrimitiveType(destinationEndpoint.Type())
	if !destinationIsPrimitive {
		return nil
	}

	// Require both properties to have the same primitive type
	if !sourcePrimitive.Equals(destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
		return writer(reader)
	}
}

// assignAliasedPrimitiveTypeFromAliasedPrimitiveType will generate a direct assignment if both
// types have the same underlying primitive type and are not optional
//
// <destination> = <cast>(<source>)
//
func assignAliasedPrimitiveTypeFromAliasedPrimitiveType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be a name that resolves to a primitive type
	_, sourceType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	sourcePrimitive, sourceIsPrimitive := AsPrimitiveType(sourceType)
	if !sourceIsPrimitive {
		return nil
	}

	// Require destination to be a name the resolves to a primitive type
	destinationName, destinationType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	destinationPrimitive, destinationIsPrimitive := AsPrimitiveType(destinationType)
	if !destinationIsPrimitive {
		return nil
	}

	// Require both properties to have the same primitive type
	if !sourcePrimitive.Equals(destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
		return writer(&dst.CallExpr{
			Fun:  destinationName.AsType(generationContext),
			Args: []dst.Expr{reader},
		})
	}
}

// assignFromAliasedPrimitiveType will convert an alias of a primitive type into that primitive
// type as long as it is not optional and we can find a conversion to consume that primitive value
func assignFromAliasedPrimitiveType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require source to be a name that resolves to a primitive type
	_, sourceType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	sourcePrimitive, sourceIsPrimitive := AsPrimitiveType(sourceType)
	if !sourceIsPrimitive {
		return nil
	}

	// Require a conversion for the underlying type
	primitiveEndpoint := sourceEndpoint.WithType(sourcePrimitive)
	conversion, _ := createTypeConversion(primitiveEndpoint, destinationEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {

		actualReader := &dst.CallExpr{
			Fun:  sourcePrimitive.AsType(generationContext),
			Args: []dst.Expr{reader},
		}

		return conversion(actualReader, writer, generationContext)
	}
}

// assignToAliasedPrimitiveType will convert a primitive value into the aliased type as long as it
// is not optional and we can find a conversion to give us the primitive type.
//
// <destination> = <cast>(<source>)
//
func assignToAliasedPrimitiveType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require destination to be a name the resolves to a primitive type
	destinationName, destinationType, ok := conversionContext.ResolveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}
	destinationPrimitive, destinationIsPrimitive := AsPrimitiveType(destinationType)
	if !destinationIsPrimitive {
		return nil
	}

	// Require a conversion for the underlying type
	primitiveEndpoint := sourceEndpoint.WithType(destinationPrimitive)
	conversion, _ := createTypeConversion(sourceEndpoint, primitiveEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {

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
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be an array type
	sourceArray, sourceIsArray := AsArrayType(sourceEndpoint.Type())
	if !sourceIsArray {
		return nil
	}

	// Require destination to be an array type
	destinationArray, destinationIsArray := AsArrayType(destinationEndpoint.Type())
	if !destinationIsArray {
		return nil
	}

	// Require a conversion between the array types
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceArray.element)
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationArray.element)
	conversion, _ := createTypeConversion(unwrappedSourceEndpoint, unwrappedDestinationEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
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
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be a map
	sourceMap, sourceIsMap := AsMapType(sourceEndpoint.Type())
	if !sourceIsMap {
		// Source is not a map
		return nil
	}

	// Require destination to be a map
	destinationMap, destinationIsMap := AsMapType(destinationEndpoint.Type())
	if !destinationIsMap {
		// Destination is not a map
		return nil
	}

	// Require map keys to be identical
	if !sourceMap.key.Equals(destinationMap.key) {
		// Keys are different types
		return nil
	}

	// Require a conversion between the map items
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceMap.value)
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationMap.value)
	conversion, _ := createTypeConversion(unwrappedSourceEndpoint, unwrappedDestinationEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
		// We create three obviously related identifiers to use for the conversion
		itemId := sourceEndpoint.CreateLocal("Value")
		keyId := sourceEndpoint.CreateLocal("Key")
		tempId := sourceEndpoint.CreateLocal("Map")

		declaration := astbuilder.SimpleAssignment(
			dst.NewIdent(tempId),
			token.DEFINE,
			astbuilder.MakeMap(destinationMap.key.AsType(generationContext), destinationMap.value.AsType(generationContext)))

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

// assignEnumTypeFromEnumType will generate a conversion if both types have the same underlying
// primitive type and neither source nor destination is optional
//
// <local> = <baseType>(<source>)
// <destination> = <enum>(<local>)
//
// We don't technically need this one, but it generates nicer code because it bypasses an unnecessary cast.
func assignEnumTypeFromEnumType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be an enumeration
	_, sourceType, sourceFound := conversionContext.ResolveType(sourceEndpoint.Type())
	if !sourceFound {
		return nil
	}
	sourceEnum, sourceIsEnum := AsEnumType(sourceType)
	if !sourceIsEnum {
		return nil
	}

	// Require destination to be an enumeration
	destinationName, destinationType, destinationFound := conversionContext.ResolveType(destinationEndpoint.Type())
	if !destinationFound {
		return nil
	}
	destinationEnum, destinationIsEnum := AsEnumType(destinationType)
	if !destinationIsEnum {
		return nil
	}

	// Require enumerations to have the same base types
	if !sourceEnum.baseType.Equals(destinationEnum.baseType) {
		return nil
	}

	local := destinationEndpoint.CreateLocal("", "As"+destinationName.Name(), "Value")
	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, ctx *CodeGenerationContext) []dst.Stmt {
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

// assignPrimitiveTypeFromEnumType will generate a conversion from an enumeration if the
// destination has the underlying base type of the enumeration and neither source nor destination
// is optional
//
// <local> = <baseType>(<source>)
// <destination> = <enum>(<local>)
//
func assignPrimitiveTypeFromEnumType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, srcOpt := AsOptionalType(sourceEndpoint.Type()); srcOpt {
		return nil
	}

	// Require destination to be non-optional
	if _, dstOpt := AsOptionalType(destinationEndpoint.Type()); dstOpt {
		return nil
	}

	// Require source to be an enumeration
	_, srcType, ok := conversionContext.ResolveType(sourceEndpoint.Type())
	if !ok {
		return nil
	}
	srcEnum, srcIsEnum := AsEnumType(srcType)
	if !srcIsEnum {
		return nil
	}

	// Require destination to be a primitive type
	dstPrimitive, ok := AsPrimitiveType(destinationEndpoint.Type())
	if !ok {
		return nil
	}

	// Require enumeration to have the destination as base type
	if !srcEnum.baseType.Equals(dstPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, ctx *CodeGenerationContext) []dst.Stmt {
		return writer(astbuilder.CallFunc(dstPrimitive.Name(), reader))
	}
}

// assignObjectTypeFromObjectType will generate a conversion if both properties are TypeNames
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
func assignObjectTypeFromObjectType(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	conversionContext *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be an object
	_, sourceType, sourceFound := conversionContext.ResolveType(sourceEndpoint.Type())
	if !sourceFound {
		return nil
	}
	if _, sourceIsObject := AsObjectType(sourceType); !sourceIsObject {
		return nil
	}

	// Require destination to be an object
	destinationName, destinationType, destinationFound := conversionContext.ResolveType(destinationEndpoint.Type())
	if !destinationFound {
		return nil
	}
	_, destinationIsObject := AsObjectType(destinationType)
	if !destinationIsObject {
		return nil
	}

	copyVar := destinationEndpoint.CreateLocal()

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {

		localId := dst.NewIdent(copyVar)
		errLocal := dst.NewIdent("err")

		declaration := astbuilder.LocalVariableDeclaration(copyVar, createTypeDeclaration(destinationName, generationContext), "")

		// If our reader is a dereference, we strip that off (because we need a pointer), else we
		// take the address of it
		var actualReader dst.Expr
		if deref, ok := reader.(*dst.UnaryExpr); ok && deref.Op == token.MUL {
			actualReader = deref.X
		} else {
			actualReader = astbuilder.AddrOf(reader)
		}

		var conversion dst.Stmt
		if destinationName.PackageReference.Equals(generationContext.CurrentPackage()) {
			conversion = astbuilder.SimpleAssignment(
				errLocal,
				token.DEFINE,
				astbuilder.CallExpr(localId, conversionContext.functionName, actualReader))
		} else {
			conversion = astbuilder.SimpleAssignment(
				errLocal,
				token.DEFINE,
				astbuilder.CallExpr(reader, conversionContext.functionName, localId))
		}

		checkForError := astbuilder.ReturnIfNotNil(
			errLocal,
			astbuilder.WrappedErrorf(
				"populating %s from %s, calling %s()",
				destinationEndpoint.name, sourceEndpoint.name, conversionContext.functionName))

		assignment := writer(dst.NewIdent(copyVar))
		return astbuilder.Statements(declaration, conversion, checkForError, assignment)
	}
}

// assignKnownReferenceFromKnownReference will generate a direct assignment if both types are genruntime.KnownResourceReference
//
// <destination> = <source>.Copy()
//
func assignKnownReferenceFromKnownReference(
	sourceEndpoint *StorageConversionEndpoint,
	destinationEndpoint *StorageConversionEndpoint,
	_ *StorageConversionContext) StorageTypeConversion {

	// Require source to be non-optional
	if _, sourceIsOptional := AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require destination to be non-optional
	if _, destinationIsOptional := AsOptionalType(destinationEndpoint.Type()); destinationIsOptional {
		return nil
	}

	// Require source to be a named type
	sourceName, sourceIsName := AsTypeName(sourceEndpoint.Type())
	if !sourceIsName {
		return nil
	}

	// Require destination to be a named type
	destinationName, destinationIsName := AsTypeName(destinationEndpoint.Type())
	if !destinationIsName {
		return nil
	}

	// Require source to be a KnownResourceReference
	if sourceName.Name() != "KnownResourceReference" {
		return nil
	}

	// Require destination to be a KnownResourceReference
	if destinationName.Name() != "KnownResourceReference" {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, generationContext *CodeGenerationContext) []dst.Stmt {
		return writer(astbuilder.CallExpr(reader, "Copy"))
	}
}

func createTypeDeclaration(name TypeName, generationContext *CodeGenerationContext) dst.Expr {
	if name.PackageReference.Equals(generationContext.CurrentPackage()) {
		return dst.NewIdent(name.Name())
	}

	packageName := generationContext.MustGetImportedPackageName(name.PackageReference)
	return astbuilder.Selector(dst.NewIdent(packageName), name.Name())
}
