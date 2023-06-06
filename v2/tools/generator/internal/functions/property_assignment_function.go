/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"github.com/pkg/errors"
	"go/token"
	"sort"

	"github.com/dave/dst"
	"golang.org/x/exp/maps"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
)

// PropertyAssignmentFunction represents a function that assigns all the properties from one resource or object to
// another. Used to perform a single step of the conversions required to/from the hub version, or to convert from
// status to spec.
type PropertyAssignmentFunction struct {
	// Name is the unique name of this function
	name string
	// receiverDefinition is the type on which this function will be hosted
	receiverDefinition astmodel.TypeDefinition
	// otherDefinition is the type we are converting to (or from). This will be a type which is "closer"
	// to the hub storage type, making this a building block of the final conversion.
	otherDefinition astmodel.TypeDefinition
	// conversions is a map of all property conversions we are going to use, keyed by name of the
	// receiver endpoint (which may be a property, function, or property bag item)
	conversions map[string]StoragePropertyConversion
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
	// direction indicates the kind of conversion we are generating
	direction conversions.Direction
	// conversionContext is additional information about the context in which this conversion was made
	conversionContext *conversions.PropertyConversionContext
	// identifier to use for our receiver in generated code
	receiverName string
	// identifier to use for our parameter in generated code
	parameterName string
	// readsFromPropertyBag keeps track of whether we will be reading property values from a property bag
	readsFromPropertyBag bool
	// writesToPropertyBag keeps track of whether we will be writing property values into a property bag
	writesToPropertyBag bool
	// augmentationInterface is the conversion augmentation interface associated with this conversion.
	// If this is nil, there is no augmented conversion associated with this conversion
	augmentationInterface *astmodel.TypeName
	// packageReferences is our set of referenced packages
	packageReferences *astmodel.PackageReferenceSet
	// knownLocals is the set of local variables in the function
	knownLocals *astmodel.KnownLocalsSet
	// sourcePropertyBag is the (optional) property bag property on the source type
	sourcePropertyBag *astmodel.PropertyDefinition
	// destinationPropertyBag is the (optional) property bag property on the destination type
	destinationPropertyBag *astmodel.PropertyDefinition
}

// StoragePropertyConversion represents a function that generates the correct AST to convert a single property value
// Different functions will be used, depending on the types of the properties to be converted.
// source is an expression that returns the source we are converting from (a Resource or other Object)
// destination is an expression that returns the destination we are converting to (again, a Resource or other Object)
// The function returns a sequence of statements to carry out the stated conversion/copy
type StoragePropertyConversion func(
	source dst.Expr,
	destination dst.Expr,
	knownLocals *astmodel.KnownLocalsSet,
	generationContext *astmodel.CodeGenerationContext) ([]dst.Stmt, error)

// Ensure that PropertyAssignmentFunction implements Function
var _ astmodel.Function = &PropertyAssignmentFunction{}

// Name returns the name of this function
func (fn *PropertyAssignmentFunction) Name() string {
	return fn.name
}

// RequiredPackageReferences returns the set of package references required by this function
func (fn *PropertyAssignmentFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return fn.packageReferences
}

// References returns the set of types referenced by this function
func (fn *PropertyAssignmentFunction) References() astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet(fn.ParameterType())
	if fn.augmentationInterface != nil {
		result.Add(*fn.augmentationInterface)
	}

	return result
}

// Equals checks to see if the supplied function is the same as this one
func (fn *PropertyAssignmentFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	if other, ok := f.(*PropertyAssignmentFunction); ok {
		if fn.Name() != other.Name() {
			// Different name means not-equal
			return false
		}

		if len(fn.conversions) != len(other.conversions) {
			// Different count of conversions means not-equal
			return false
		}

		for name := range fn.conversions {
			if _, found := other.conversions[name]; !found {
				// Missing conversion means not-equal
				return false
			}
		}

		return true
	}

	return false
}

// Direction returns this functions direction of conversion
func (fn *PropertyAssignmentFunction) Direction() conversions.Direction {
	return fn.direction
}

// AsFunc renders this function as an AST for serialization to a Go source file
func (fn *PropertyAssignmentFunction) AsFunc(generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	description := fn.direction.SelectString(
		fmt.Sprintf("populates our %s from the provided source %s", receiver.Name(), fn.ParameterType().Name()),
		fmt.Sprintf("populates the provided destination %s from our %s", fn.ParameterType().Name(), receiver.Name()))

	// We always use a pointer receiver, so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: fn.receiverName,
		ReceiverType:  receiverType,
		Name:          fn.Name(),
		Body:          fn.generateBody(fn.receiverName, fn.parameterName, generationContext),
	}

	funcDetails.AddParameter(
		fn.parameterName,
		astbuilder.PointerTo(fn.ParameterType().AsType(generationContext)))

	funcDetails.AddReturns("error")
	funcDetails.AddComments(description)

	return funcDetails.DefineFunc()
}

func (fn *PropertyAssignmentFunction) ParameterType() astmodel.TypeName {
	return fn.otherDefinition.Name()
}

// generateBody returns the statements required for the conversion function
// receiver is an expression for access our receiver type, used to qualify field access
// parameter is an expression for access to our parameter passed to the function, also used for field access
// generationContext is our code generation context, passed to allow resolving of identifiers in other packages
func (fn *PropertyAssignmentFunction) generateBody(
	receiver string,
	parameter string,
	generationContext *astmodel.CodeGenerationContext,
) ([]dst.Stmt, error) {
	// source is the identifier from which we are reading values
	source := fn.direction.SelectString(parameter, receiver)

	// destination is the identifier onto which we write values
	destination := fn.direction.SelectString(receiver, parameter)

	knownLocals := fn.knownLocals.Clone()

	bagPrologue := fn.createPropertyBagPrologue(source, generationContext)
	assignments := fn.generateAssignments(knownLocals, dst.NewIdent(source), dst.NewIdent(destination), generationContext)
	bagEpilogue := fn.propertyBagEpilogue(destination)
	handleOverrideInterface := fn.handleAugmentationInterface(receiver, parameter, knownLocals, generationContext)

	return astbuilder.Statements(
		bagPrologue,
		assignments,
		bagEpilogue,
		handleOverrideInterface,
		astbuilder.ReturnNoError()), nil
}

// createPropertyBagPrologue creates any introductory statements needed to set up our property bag before we start doing
// assignments. We need to handle three cases:
//
//	o If our source has a property bag, we clone it.
//	o If our destination has a property bag (and our source does not), we create a new one.
//	o If neither source nor destination has a property bag, we don't need to do anything.
//
// source is the name of the source to read the property bag from
func (fn *PropertyAssignmentFunction) createPropertyBagPrologue(
	source string,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	// We don't need the prologue if we're not using a property bag at all.
	// So when are we using one?
	// - If we're reading from the property bag
	// - If we're writing to the property bag
	// - If our destination has a property bag that needs initialization
	if !fn.readsFromPropertyBag && !fn.writesToPropertyBag && fn.destinationPropertyBag == nil {
		return nil
	}

	// Don't refactor the local genruntimePkg out to this scope - calling MustGetImportedPackageName() flags the
	// package as referenced, so we must only call that if we are actually going to reference the genruntime package

	var createBag dst.Expr
	var comment string
	genruntimePkg := generationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)

	if fn.sourcePropertyBag != nil {
		createBag = astbuilder.CallQualifiedFunc(
			genruntimePkg,
			"NewPropertyBag",
			astbuilder.Selector(dst.NewIdent(source), string(fn.sourcePropertyBag.PropertyName())))
		comment = "// Clone the existing property bag"
	} else {
		createBag = astbuilder.CallQualifiedFunc(
			genruntimePkg,
			"NewPropertyBag")
		comment = "// Create a new property bag"
	}

	initializeBag := astbuilder.ShortDeclaration(
		fn.conversionContext.PropertyBagName(),
		createBag)
	initializeBag.Decs.Before = dst.NewLine
	astbuilder.AddComment(&initializeBag.Decorations().Start, comment)

	return astbuilder.Statements(initializeBag)
}

// propertyBagEpilogue creates any concluding statements required to handle our property bag after assignments are
// complete.
//
//	o If the destination has a property bag
//	  >  If our bag is empty, we set the destination to nil
//	  >  Otherwise we need to store our current property bag there
//	o Otherwise we do nothing
func (fn *PropertyAssignmentFunction) propertyBagEpilogue(
	destination string,
) []dst.Stmt {
	prop := fn.destinationPropertyBag
	found := prop != nil
	if found {
		bagId := dst.NewIdent(fn.conversionContext.PropertyBagName())
		bagProperty := astbuilder.Selector(dst.NewIdent(destination), string(prop.PropertyName()))

		condition := astbuilder.BinaryExpr(astbuilder.CallFunc("len", bagId), token.GTR, astbuilder.IntLiteral(0))

		storeBag := astbuilder.SimpleAssignment(bagProperty, bagId)
		storeNil := astbuilder.SimpleAssignment(bagProperty, astbuilder.Nil())

		store := astbuilder.SimpleIfElse(
			condition,
			astbuilder.Statements(storeBag),
			astbuilder.Statements(storeNil))
		store.Decs.Before = dst.EmptyLine
		astbuilder.AddComment(&store.Decorations().Start, "// Update the property bag")

		return astbuilder.Statements(store)
	}

	return nil
}

// handleAugmentationInterface handles dealing with the override interface if there is one
// Generates code that looks like:
//
//	var accountAsAny any = account
//	if augmented, ok := accountAsAny.(augmentConversionForBatchAccount); ok {
//	   err := augmentedAccount.AssignPropertiesFrom(source)
//	   if err != nil {
//	       return errors.Wrap(
//	           err,
//	           "calling augmented AssignPropertiesFrom() for conversion from v20210101s.BatchAccount")
//	   }
//	}
func (fn *PropertyAssignmentFunction) handleAugmentationInterface(
	receiver string,
	parameter string,
	knownLocals *astmodel.KnownLocalsSet,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	if fn.augmentationInterface == nil {
		return nil
	}

	overrideInterface := fn.augmentationInterface.AsType(generationContext)
	receiverAsAnyIdent := knownLocals.CreateLocal(receiver + "AsAny")

	sourceAsAny := astbuilder.NewVariableAssignmentWithType(receiverAsAnyIdent, dst.NewIdent("any"), dst.NewIdent(receiver))

	// Clone locals at this point as we're entering an if block
	knownLocals = knownLocals.Clone()

	augmentedReceiverIdent := knownLocals.CreateLocal("augmented" + fn.idFactory.CreateIdentifier(receiver, astmodel.Exported))
	conversionFuncName := fn.Direction().SelectString("AssignPropertiesFrom", "AssignPropertiesTo")
	callAssignOverride := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallQualifiedFunc(augmentedReceiverIdent, conversionFuncName, dst.NewIdent(parameter)))
	returnIfNotNil := astbuilder.ReturnIfNotNil(
		dst.NewIdent("err"),
		astbuilder.WrappedError(
			generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference),
			fmt.Sprintf("calling augmented %s() for conversion", conversionFuncName)))

	ifStmt := astbuilder.IfType(
		dst.NewIdent(receiverAsAnyIdent),
		overrideInterface,
		augmentedReceiverIdent,
		callAssignOverride,
		returnIfNotNil)
	sourceAsAny.Decorations().Before = dst.EmptyLine
	sourceAsAny.Decorations().Start.Prepend(fmt.Sprintf("// Invoke the %s interface (if implemented) to customize the conversion", fn.augmentationInterface.Name()))

	return astbuilder.Statements(
		sourceAsAny,
		ifStmt)
}

// generateAssignments generates a sequence of statements to copy information between the two types
func (fn *PropertyAssignmentFunction) generateAssignments(
	knownLocals *astmodel.KnownLocalsSet,
	source dst.Expr,
	destination dst.Expr,
	generationContext *astmodel.CodeGenerationContext,
) ([]dst.Stmt, error) {
	var result []dst.Stmt

	// Find all the properties for which we have a conversion
	properties := maps.Keys(fn.conversions)

	// Sort the properties into alphabetical order to ensure deterministic generation
	sort.Strings(properties)

	// Accumulate all the statements required for conversions, in alphabetical order
	for _, prop := range properties {
		conversion := fn.conversions[prop]
		block := conversion(source, destination, knownLocals, generationContext)
		if len(block) > 0 {
			firstStatement := block[0]
			firstStatement.Decorations().Before = dst.EmptyLine
			firstStatement.Decorations().Start.Prepend("// " + prop)
			result = append(result, block...)
		}
	}

	return result, nil
}

// sourceType returns the type we are reading information from
// When converting FROM, otherDefinition.Type() is our source
// When converting TO, receiverDefinition.Type() is our source
// Our inverse is destinationType()
func (fn *PropertyAssignmentFunction) sourceType() astmodel.Type {
	return fn.direction.SelectType(fn.otherDefinition.Type(), fn.receiverDefinition.Type())
}

// destinationType returns the type we are writing information from
// When converting FROM, receiverDefinition.Type() is our source
// When converting TO, otherDefinition.Type() is our source
// Our inverse is sourceType()
func (fn *PropertyAssignmentFunction) destinationType() astmodel.Type {
	return fn.direction.SelectType(fn.receiverDefinition.Type(), fn.otherDefinition.Type())
}
