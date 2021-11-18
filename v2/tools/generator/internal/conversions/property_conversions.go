/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"go/token"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// PropertyConversion generates the AST for a given property conversion.
// reader is an expression to read the original value.
// writer is a function that accepts an expression for reading a value and creates one or more
// statements to write that value.
// Both of these might be complex expressions, possibly involving indexing into arrays or maps.
type PropertyConversion func(
	reader dst.Expr,
	writer func(dst.Expr) []dst.Stmt,
	knownLocals *astmodel.KnownLocalsSet,
	generationContext *astmodel.CodeGenerationContext) []dst.Stmt

// PropertyConversionFactory represents factory methods that can be used to create a PropertyConversion for a specific
// pair of properties
// source is the property conversion endpoint that will be read
// destination is the property conversion endpoint that will be written
// ctx contains additional information that may be needed when creating the property conversion
//
// Each conversion should be written with lead predicates to make sure that it only fires in the correct circumstances.
// This requires, in particular, that most conversions check for optionality and bag items and exit early when those are
// found.
// Phrased another way, conversions should not rely on the order of listing in propertyConversionFactories in order to
// generate the correct code; any conversion that relies on being "protected" from particular situations by having other
// conversions earlier in the list held by propertyConversionFactories is brittle and likely to generate the incorrect
// code if the order of items in the list is modified.
//
type PropertyConversionFactory func(
	source *TypedConversionEndpoint,
	destination *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion

// A list of all known type conversion factory methods
var propertyConversionFactories []PropertyConversionFactory

func init() {
	propertyConversionFactories = []PropertyConversionFactory{
		// Property bag items
		pullFromBagItem,
		writeToBagItem,
		// Primitive types and aliases
		assignPrimitiveFromPrimitive,
		assignAliasedPrimitiveFromAliasedPrimitive,
		// Handcrafted implementations in genruntime
		assignHandcraftedImplementations,
		// Collection Types
		assignArrayFromArray,
		assignMapFromMap,
		// Enumerations
		assignEnumFromEnum,
		assignPrimitiveFromEnum,
		// Complex object types
		assignObjectFromObject,
		// Known types
		copyKnownType(astmodel.KnownResourceReferenceType, "Copy", returnsValue),
		copyKnownType(astmodel.ResourceReferenceType, "Copy", returnsValue),
		copyKnownType(astmodel.ArbitraryOwnerReference, "Copy", returnsValue),
		copyKnownType(astmodel.ConditionType, "Copy", returnsValue),
		copyKnownType(astmodel.JSONType, "DeepCopy", returnsReference),
		copyKnownType(astmodel.ObjectMetaType, "DeepCopy", returnsReference),
		// Meta-conversions
		assignFromOptional,
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
	err := errors.Errorf(
		"no conversion found to assign %q from %q",
		astmodel.DebugDescription(destinationEndpoint.Type(), conversionContext.Types()),
		astmodel.DebugDescription(sourceEndpoint.Type(), conversionContext.Types()))

	return nil, err
}

// NameOfPropertyAssignmentFunction returns the name of the property assignment function
func NameOfPropertyAssignmentFunction(name astmodel.TypeName, direction Direction, idFactory astmodel.IdentifierFactory) string {
	nameOfOtherType := idFactory.CreateIdentifier(name.Name(), astmodel.Exported)
	return "AssignProperties" + direction.SelectString("From", "To") + nameOfOtherType
}

// writeToBagItem will generate a conversion where the destination is in our property bag
//
// For non-optional sources, the value is directly added
//
// <propertyBag>.Add(<propertyName>, <source>)
//
// For optional sources, the value is only added if non-nil
//
// if <source> != nil {
//     <propertyBag>.Add(<propertyName>, *<source>)
// }
//
func writeToBagItem(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to be a property bag item
	destinationBagItem, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type())
	if !destinationIsBagItem {
		// Destination is not optional
		return nil
	}

	// Work out our source type, and whether it's optional
	actualSourceType := sourceEndpoint.Type()
	sourceOptional, sourceIsOptional := astmodel.AsOptionalType(actualSourceType)
	if sourceIsOptional {
		actualSourceType = sourceOptional.BaseType()
	}

	// Require the item in the bag to be exactly the same type as our source
	// (We don't want to recursively do all the conversions because our property bag item SHOULD always contain
	// exactly the expected type, so no conversion should be required. Plus, our conversions are designed to isolate
	// the source and destination from each other (so that changes to one don't impact the other), but with the
	// property bag everything gets immediately serialized so everything is already nicely isolated.
	if !astmodel.TypeEquals(destinationBagItem.Element(), actualSourceType) {
		return nil
	}

	return func(reader dst.Expr, _ func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		createAddToBag := func(expr dst.Expr) dst.Stmt {
			addToBag := astbuilder.InvokeQualifiedFunc(
				conversionContext.PropertyBagName(),
				"Add",
				astbuilder.StringLiteralf(destinationEndpoint.Name()),
				expr)

			return addToBag
		}

		var writer dst.Stmt
		if sourceIsOptional {
			writer = astbuilder.IfNotNil(
				reader,
				createAddToBag(astbuilder.Dereference(reader)))
		} else {
			writer = createAddToBag(reader)
		}

		return astbuilder.Statements(writer)
	}
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require destination to be optional
	destinationOptional, destinationIsOptional := astmodel.AsOptionalType(destinationEndpoint.Type())
	if !destinationIsOptional {
		// Destination is not optional
		return nil
	}

	// Require source to be non-optional
	// (to ensure that assignFromOptional triggers first when handling option to optional conversion)
	if _, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type()); sourceIsOptional {
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	unwrappedEndpoint := destinationEndpoint.WithType(destinationOptional.Element())
	conversion, _ := CreateTypeConversion(sourceEndpoint, unwrappedEndpoint, conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// Create a writer that uses the address of the passed expression
		// If expr isn't a plain identifier (implying a local variable), we introduce one
		// This both allows us to avoid aliasing and complies with Go language semantics
		addrOfWriter := func(expr dst.Expr) []dst.Stmt {
			if _, ok := expr.(*dst.Ident); ok {
				return writer(astbuilder.AddrOf(expr))
			}

			// Only obtain our local variable name after we know we need it
			// (this avoids reserving the name and not using it, which can interact with other conversions)
			local := knownLocals.CreateSingularLocal(destinationEndpoint.Name(), "", "Temp")

			assignment := astbuilder.ShortDeclaration(local, expr)

			writing := writer(astbuilder.AddrOf(dst.NewIdent(local)))

			return astbuilder.Statements(assignment, writing)
		}

		return conversion(reader, addrOfWriter, knownLocals, generationContext)
	}
}

// pullFromBagItem will populate a property from a property bag
//
// if <propertyBag>.Contains(<sourceName>) {
//     var <value> <destinationType>
//     err := <propertyBag>.Pull(<sourceName>, &<value>)
//     if err != nil {
//         return errors.Wrapf(err, ...)
//     }
//
//     <destination> = <value>
// } else {
//     <destination> = <zero>
// }
//
func pullFromBagItem(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to be a bag item
	sourceBagItem, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type())
	if !sourceIsBagItem {
		return nil
	}

	// Work out our destination type, and whether it's optional
	actualDestinationType := destinationEndpoint.Type()
	destinationOptional, destinationIsOptional := astmodel.AsOptionalType(actualDestinationType)
	if destinationIsOptional {
		actualDestinationType = destinationOptional.BaseType()
	}

	// Require the item in the bag to be exactly the same type as our destination
	// (We don't want to recursively do all the conversions because our property bag item SHOULD always contain
	// exactly the expected type, so no conversion should be required. Plus, our conversions are designed to isolate
	// the source and destination from each other (so that changes to one don't impact the other), but with the
	// property bag everything gets immediately serialized so everything is already nicely isolated.
	if !astmodel.TypeEquals(sourceBagItem.Element(), actualDestinationType) {
		return nil
	}

	errIdent := dst.NewIdent("err")

	return func(_ dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		// our first parameter is an expression to read the value from our original instance, but in this case we're
		// going to read from the property bag, so we're ignoring it.

		local := knownLocals.CreateSingularLocal(sourceEndpoint.Name(), "", "Read")
		errorsPkg := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

		condition := astbuilder.CallQualifiedFunc(
			conversionContext.PropertyBagName(),
			"Contains",
			astbuilder.StringLiteral(sourceEndpoint.Name()))

		declare := astbuilder.NewVariableWithType(
			local,
			sourceBagItem.AsType(generationContext))

		pull := astbuilder.ShortDeclaration(
			"err",
			astbuilder.CallQualifiedFunc(
				conversionContext.PropertyBagName(),
				"Pull",
				astbuilder.StringLiteral(sourceEndpoint.Name()),
				astbuilder.AddrOf(dst.NewIdent(local))))

		returnIfErr := astbuilder.ReturnIfNotNil(
			errIdent,
			astbuilder.WrappedErrorf(
				errorsPkg,
				"pulling '%s' from propertyBag",
				sourceEndpoint.Name()))
		returnIfErr.Decorations().After = dst.EmptyLine

		var reader dst.Expr
		if destinationIsOptional {
			reader = astbuilder.AddrOf(dst.NewIdent(local))
		} else {
			reader = dst.NewIdent(local)
		}

		assignValue := writer(reader)

		assignZero := writer(destinationEndpoint.Type().AsZero(conversionContext.Types(), generationContext))

		ifStatement := astbuilder.SimpleIfElse(
			condition,
			astbuilder.Statements(declare, pull, returnIfErr, assignValue),
			assignZero)

		return astbuilder.Statements(ifStatement)
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
//
// Must trigger before assignToOptional so we generate the right zero values; to enforce this, assignToOptional includes
// a predicate check that the source is NOT optional, allowing this conversion to trigger first.
//
func assignFromOptional(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

	// Require source to be optional
	sourceOptional, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type())
	if !sourceIsOptional {
		return nil
	}

	// Require a conversion between the unwrapped type and our source
	unwrappedEndpoint := sourceEndpoint.WithType(sourceOptional.Element())
	conversion, _ := CreateTypeConversion(
		unwrappedEndpoint,
		destinationEndpoint,
		conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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
			local := knownLocals.CreateSingularLocal(sourceEndpoint.Name(), "", "AsRead")
			cacheOriginal = astbuilder.ShortDeclaration(local, reader)
			actualReader = dst.NewIdent(local)
		}

		checkForNil := astbuilder.AreNotEqual(actualReader, astbuilder.Nil())

		// If we have a value, need to convert it to our destination type
		// We use a cloned knownLocals as the write is within our if statement and we don't want locals to leak
		writeActualValue := conversion(
			astbuilder.Dereference(actualReader),
			writer,
			knownLocals.Clone(),
			generationContext)

		writeZeroValue := writer(
			destinationEndpoint.Type().AsZero(conversionContext.Types(), generationContext))

		stmt := astbuilder.SimpleIfElse(
			checkForNil,
			writeActualValue,
			writeZeroValue)

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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

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

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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
			knownLocals,
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	if !astmodel.TypeEquals(sourcePrimitive, destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	if !astmodel.TypeEquals(sourcePrimitive, destinationPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		actualReader := &dst.CallExpr{
			Fun:  sourcePrimitive.AsType(generationContext),
			Args: []dst.Expr{reader},
		}

		return conversion(actualReader, writer, knownLocals, generationContext)
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

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

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		actualWriter := func(expr dst.Expr) []dst.Stmt {
			castToAlias := &dst.CallExpr{
				Fun:  destinationName.AsType(generationContext),
				Args: []dst.Expr{expr},
			}

			return writer(castToAlias)
		}

		return conversion(reader, actualWriter, knownLocals, generationContext)
	}
}

// handCraftedConversion represents a hand-coded conversion
// this can be used to share code for common conversions (e.g. []string → []string)
type handCraftedConversion struct {
	fromType astmodel.Type
	toType   astmodel.Type

	implPackage astmodel.PackageReference
	implFunc    string
}

var handCraftedConversions = []handCraftedConversion{
	{
		fromType:    astmodel.NewMapType(astmodel.StringType, astmodel.StringType),
		toType:      astmodel.NewMapType(astmodel.StringType, astmodel.StringType),
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "CloneMapOfStringToString",
	},
	{
		fromType:    astmodel.NewArrayType(astmodel.StringType),
		toType:      astmodel.NewArrayType(astmodel.StringType),
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "CloneSliceOfString",
	},
	{
		fromType:    astmodel.NewArrayType(astmodel.ConditionType),
		toType:      astmodel.NewArrayType(astmodel.ConditionType),
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "CloneSliceOfCondition",
	},
	{
		fromType:    astmodel.NewOptionalType(astmodel.IntType),
		toType:      astmodel.NewOptionalType(astmodel.IntType),
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "ClonePointerToInt",
	},
	{
		fromType:    astmodel.NewOptionalType(astmodel.StringType),
		toType:      astmodel.NewOptionalType(astmodel.StringType),
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "ClonePointerToString",
	},
	{
		fromType:    astmodel.NewOptionalType(astmodel.StringType),
		toType:      astmodel.StringType,
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "GetOptionalStringValue",
	},
	{
		fromType:    astmodel.NewOptionalType(astmodel.IntType),
		toType:      astmodel.IntType,
		implPackage: astmodel.GenRuntimeReference,
		implFunc:    "GetOptionalIntValue",
	},
}

func assignHandcraftedImplementations(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

	for _, impl := range handCraftedConversions {
		if astmodel.TypeEquals(sourceEndpoint.Type(), impl.fromType) &&
			astmodel.TypeEquals(destinationEndpoint.Type(), impl.toType) {
			return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, _ *astmodel.KnownLocalsSet, cgc *astmodel.CodeGenerationContext) []dst.Stmt {
				pkg := cgc.MustGetImportedPackageName(impl.implPackage)
				return writer(astbuilder.CallQualifiedFunc(pkg, impl.implFunc, reader))
			}
		}
	}

	return nil
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceArray.Element())
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationArray.Element())
	conversion, _ := CreateTypeConversion(
		unwrappedSourceEndpoint,
		unwrappedDestinationEndpoint,
		conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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
			local := knownLocals.CreateSingularLocal(sourceEndpoint.Name(), "", "Cache")
			cacheOriginal = astbuilder.ShortDeclaration(local, reader)
			actualReader = dst.NewIdent(local)
		}

		checkForNil := astbuilder.AreNotEqual(actualReader, astbuilder.Nil())

		// We create three obviously related identifiers to use for the array conversion
		// The List is created in the current knownLocals scope because we need it after the loop completes.
		// The other two are created in a nested knownLocals scope because they're only needed within the loop; this
		// ensures any other locals needed for the conversion don't leak out into our main scope.
		// These suffixes must not overlap with those used for map conversion. (If these suffixes overlap, the naming
		// becomes difficult to read when converting maps containing slices or vice versa.)
		branchLocals := knownLocals.Clone()
		tempId := branchLocals.CreateSingularLocal(sourceEndpoint.Name(), "List")
		loopLocals := branchLocals.Clone() // Clone after tempId is created so that it's visible within the loop
		itemId := loopLocals.CreateSingularLocal(sourceEndpoint.Name(), "Item")
		indexId := loopLocals.CreateSingularLocal(sourceEndpoint.Name(), "Index")

		declaration := astbuilder.ShortDeclaration(
			tempId,
			astbuilder.MakeList(destinationArray.AsType(generationContext), astbuilder.CallFunc("len", actualReader)))

		writeToElement := func(expr dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					&dst.IndexExpr{
						X:     dst.NewIdent(tempId),
						Index: dst.NewIdent(indexId),
					},
					expr),
			}
		}

		avoidAliasing := astbuilder.ShortDeclaration(itemId, dst.NewIdent(itemId))
		avoidAliasing.Decs.Start.Append("// Shadow the loop variable to avoid aliasing")
		avoidAliasing.Decs.Before = dst.NewLine

		loopBody := astbuilder.Statements(
			avoidAliasing,
			conversion(dst.NewIdent(itemId), writeToElement, loopLocals, generationContext))

		assignValue := writer(dst.NewIdent(tempId))
		loop := astbuilder.IterateOverListWithIndex(indexId, itemId, reader, loopBody...)
		trueBranch := astbuilder.Statements(declaration, loop, assignValue)

		assignZero := writer(astbuilder.Nil())

		return astbuilder.Statements(
			cacheOriginal,
			astbuilder.SimpleIfElse(checkForNil, trueBranch, assignZero))
	}
}

// assignMapFromMap will generate a code fragment to populate an array, assuming the
// underlying types of the two arrays are compatible
//
// if <reader> != nil {
//     <map> := make(map[<key>]<type>)
//     for key, <item> := range <reader> {
//         <map>[<key>] := <item>
//     }
//     <writer> = <map>
// } else {
//     <writer> = <zero>
// }
//
func assignMapFromMap(
	sourceEndpoint *TypedConversionEndpoint,
	destinationEndpoint *TypedConversionEndpoint,
	conversionContext *PropertyConversionContext) PropertyConversion {

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	if !astmodel.TypeEquals(sourceMap.KeyType(), destinationMap.KeyType()) {
		// Keys are different types
		return nil
	}

	// Require a conversion between the map items
	unwrappedSourceEndpoint := sourceEndpoint.WithType(sourceMap.ValueType())
	unwrappedDestinationEndpoint := destinationEndpoint.WithType(destinationMap.ValueType())
	conversion, _ := CreateTypeConversion(
		unwrappedSourceEndpoint,
		unwrappedDestinationEndpoint,
		conversionContext)
	if conversion == nil {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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
			local := knownLocals.CreateSingularLocal(sourceEndpoint.Name(), "", "Cache")
			cacheOriginal = astbuilder.ShortDeclaration(local, reader)
			actualReader = dst.NewIdent(local)
		}

		checkForNil := astbuilder.AreNotEqual(actualReader, astbuilder.Nil())

		// We create three obviously related identifiers to use for the conversion.
		// These are all within the scope of the true branch of our if statement
		// The Map is created in the current knownLocals scope because we need it after the loop completes.
		// The other two are created in a nested knownLocals scope because they're only needed within the loop; this
		// ensures any other locals needed for the conversion don't leak out into our main scope.
		// These suffixes must not overlap with those used for array conversion. (If these suffixes overlap, the naming
		// becomes difficult to read when converting maps containing slices or vice versa.)
		branchLocals := knownLocals.Clone()
		tempId := branchLocals.CreateSingularLocal(sourceEndpoint.Name(), "Map")
		loopLocals := branchLocals.Clone() // Clone after tempId is created so that it's visible within the loop
		itemId := loopLocals.CreateSingularLocal(sourceEndpoint.Name(), "Value")
		keyId := loopLocals.CreateSingularLocal(sourceEndpoint.Name(), "Key")

		declaration := astbuilder.ShortDeclaration(
			tempId,
			astbuilder.MakeMapWithCapacity(
				destinationMap.KeyType().AsType(generationContext),
				destinationMap.ValueType().AsType(generationContext),
				astbuilder.CallFunc("len", actualReader)))

		assignToItem := func(expr dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					&dst.IndexExpr{
						X:     dst.NewIdent(tempId),
						Index: dst.NewIdent(keyId),
					},
					expr),
			}
		}

		avoidAliasing := astbuilder.ShortDeclaration(itemId, dst.NewIdent(itemId))
		avoidAliasing.Decs.Start.Append("// Shadow the loop variable to avoid aliasing")
		avoidAliasing.Decs.Before = dst.NewLine

		loopBody := astbuilder.Statements(
			avoidAliasing,
			conversion(dst.NewIdent(itemId), assignToItem, loopLocals, generationContext))

		loop := astbuilder.IterateOverMapWithValue(keyId, itemId, actualReader, loopBody...)
		assignMap := writer(dst.NewIdent(tempId))
		trueBranch := astbuilder.Statements(declaration, loop, assignMap)

		assignNil := writer(astbuilder.Nil())

		return astbuilder.Statements(
			cacheOriginal,
			astbuilder.SimpleIfElse(checkForNil, trueBranch, assignNil))
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	if !astmodel.TypeEquals(sourceEnum.BaseType(), destinationEnum.BaseType()) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, ctx *astmodel.CodeGenerationContext) []dst.Stmt {
		local := knownLocals.CreateSingularLocal(destinationEndpoint.Name(), "", "As"+destinationName.Name(), "Value")
		result := []dst.Stmt{
			astbuilder.ShortDeclaration(local, astbuilder.CallFunc(destinationName.Name(), reader)),
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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
	if !astmodel.TypeEquals(srcEnum.BaseType(), dstPrimitive) {
		return nil
	}

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, ctx *astmodel.CodeGenerationContext) []dst.Stmt {
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

	// Require destination to not be a bag item
	if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
		return nil
	}

	// Require source to not be a bag item
	if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
		return nil
	}

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

	return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		copyVar := knownLocals.CreateSingularLocal(destinationEndpoint.Name(), "", "Local", "Copy", "Temp")

		// We have to do this at render time in order to ensure the first conversion generated
		// declares 'err', not a later one
		tok := token.ASSIGN
		if knownLocals.TryCreateLocal("err") {
			tok = token.DEFINE
		}

		localId := dst.NewIdent(copyVar)
		errLocal := dst.NewIdent("err")

		errorsPackageName := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

		declaration := astbuilder.LocalVariableDeclaration(copyVar, createTypeDeclaration(destinationName, generationContext), "")

		// If our reader is a dereference, we strip that off (because we need a pointer), else we
		// take the address of it
		var actualReader dst.Expr
		if star, ok := reader.(*dst.StarExpr); ok {
			actualReader = star.X
		} else {
			actualReader = astbuilder.AddrOf(reader)
		}

		functionName := NameOfPropertyAssignmentFunction(sourceName, conversionContext.direction, conversionContext.idFactory)

		var conversion dst.Stmt
		if destinationName.PackageReference.Equals(generationContext.CurrentPackage()) {
			// Destination is our current type
			conversion = astbuilder.AssignmentStatement(
				errLocal,
				tok,
				astbuilder.CallExpr(localId, functionName, actualReader))
		} else {
			// Destination is another type
			conversion = astbuilder.AssignmentStatement(
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
//nolint:deadcode,unused
func assignKnownType(name astmodel.TypeName) func(*TypedConversionEndpoint, *TypedConversionEndpoint, *PropertyConversionContext) PropertyConversion {
	return func(sourceEndpoint *TypedConversionEndpoint, destinationEndpoint *TypedConversionEndpoint, _ *PropertyConversionContext) PropertyConversion {
		// Require destination to not be a bag item
		if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
			return nil
		}

		// Require source to not be a bag item
		if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
			return nil
		}

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
		if !astmodel.TypeEquals(sourceName, name) {
			return nil
		}

		// Require destination to be our specific type
		if !astmodel.TypeEquals(destinationName, name) {
			return nil
		}

		return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
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
		// Require destination to not be a bag item
		if _, destinationIsBagItem := AsPropertyBagMemberType(destinationEndpoint.Type()); destinationIsBagItem {
			return nil
		}

		// Require source to not be a bag item
		if _, sourceIsBagItem := AsPropertyBagMemberType(sourceEndpoint.Type()); sourceIsBagItem {
			return nil
		}

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
		if !astmodel.TypeEquals(sourceName, name) {
			return nil
		}

		// Require destination to be our specific type
		if !astmodel.TypeEquals(destinationName, name) {
			return nil
		}

		return func(reader dst.Expr, writer func(dst.Expr) []dst.Stmt, knownLocals *astmodel.KnownLocalsSet, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
			// If our writer is dereferencing a value, skip that as we don't need to dereference before a method call
			if star, ok := reader.(*dst.StarExpr); ok {
				reader = star.X
			}

			if returnKind == returnsReference {
				// If the copy method returns a ptr, we need to dereference
				// This dereference is always safe because we ensured that both source and destination are always
				// non-optional. The handler assignToOptional() should do the right thing when this happens.
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
