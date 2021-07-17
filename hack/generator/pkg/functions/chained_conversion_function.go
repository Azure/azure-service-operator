/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
)

// ChainedConversionFunction implements conversions to/from another type that implements a conversion function
//
// We use this for Spec types, implementing genruntime.ConvertibleSpec.
//
// Existing PropertyAssignment functions are used to implement stepwise conversion.
//
// For most types, we check to see if the type we're passed is one we can convert directly. If it is, we use the
// preexisting AssignProperties*() method. Otherwise, we chain to that type to do the conversion to an intermediate
// instance and then convert using that.
//
// func (r <receiver>) ConvertFrom(instance <interfaceType>) error {
// 	   source, ok := instance.(*<otherType>)
// 	   if !ok {
//         // Need indirect conversion
//         source = &<otherType>{}
//         source.ConvertFrom(instance)
//     }
//
//     return r.AssignPropertiesFrom(source)
// }
//
// If the two instances involved in conversion are not in the same "spoke" leading to our "hub" version, we'll reach
// the hub type and then pivot to the reverse conversion. The implementation on the hub instance does this pivot:
//
// func (s <hubtype>) ConvertFrom(instance <interfaceType>>) error {
//     return instance.ConvertTo(s)
// }
//
type ChainedConversionFunction struct {
	// nameFrom is the name for this function when converting FROM a provided instance
	nameFrom string
	// nameTo is the name for this function when converting TO a provided instance
	nameTo string
	// parameterType is common interface type for our parameter
	parameterType astmodel.TypeName
	// hubType is the TypeName of the hub type of our conversions, the central type of the hub-and-spoke model
	hubType astmodel.TypeName
	// direction is the direction of conversion - either FROM the supplied instance, or TO the supplied instance
	// We initialize this from the supplied property conversion function to guarantee consistency that the function
	// we generate is using that function correctly
	direction conversions.Direction
	// propertyAssignmentFunctionName is the name of a function we can call to copy properties from one instance to
	// another. We initialize this from the supplied PropertyAssignmentFunction
	propertyAssignmentFunctionName string
	// propertyAssignmentParameterType is the type of the parameter we need to pass to propertyAssignmentFunctionName
	// when we generate calls. We initialize this from the supplied PropertyAssignmentFunction
	propertyAssignmentParameterType astmodel.TypeName
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &ChainedConversionFunction{}

// NewSpecConversionFunction creates a chained conversion function that converts between two Spec types implementing the
// interface genruntime.ConvertibleSpec.
// hubType is the TypeName of our hub type
// propertyFunction is the function we will call to copy properties across between specs
// idFactory is an identifier factory to use for generating local identifiers
func NewSpecConversionFunction(
	hubType astmodel.TypeName,
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ChainedConversionFunction {
	result := &ChainedConversionFunction{
		nameFrom:                        "ConvertSpecFrom",
		nameTo:                          "ConvertSpecTo",
		parameterType:                   astmodel.ConvertibleSpecInterfaceType,
		hubType:                         hubType,
		propertyAssignmentFunctionName:  propertyFunction.Name(),
		propertyAssignmentParameterType: propertyFunction.otherDefinition.Name(),
		direction:                       propertyFunction.direction,
		idFactory:                       idFactory,
	}

	return result
}

// NewStatusConversionFunction creates a chained conversion function that converts between two Status types implementing
// the interface genruntime.ConvertibleStatus.
// hubType is the TypeName of our hub type
// propertyFunction is the function we will call to copy properties across between specs
// idFactory is an identifier factory to use for generating local identifiers
func NewStatusConversionFunction(
	hubType astmodel.TypeName,
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ChainedConversionFunction {
	result := &ChainedConversionFunction{
		nameFrom:                        "ConvertStatusFrom",
		nameTo:                          "ConvertStatusTo",
		parameterType:                   astmodel.ConvertibleStatusInterfaceType,
		hubType:                         hubType,
		propertyAssignmentFunctionName:  propertyFunction.Name(),
		propertyAssignmentParameterType: propertyFunction.otherDefinition.Name(),
		direction:                       propertyFunction.direction,
		idFactory:                       idFactory,
	}

	return result
}


func (fn *ChainedConversionFunction) Name() string {
	return fn.direction.SelectString(fn.nameFrom, fn.nameTo)
}

func (fn *ChainedConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		astmodel.GenRuntimeReference,
		fn.parameterType.PackageReference,
		fn.hubType.PackageReference,
		fn.propertyAssignmentParameterType.PackageReference)
}

func (fn *ChainedConversionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		fn.parameterType,
		fn.hubType,
		fn.propertyAssignmentParameterType)
}

func (fn *ChainedConversionFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {

	// Create a sensible name for our receiver
	receiverName := fn.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)

	// We always use a pointer receiver so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverType,
		Name:          fn.Name(),
	}

	parameterName := fn.direction.SelectString("source", "destination")
	funcDetails.AddParameter(parameterName, fn.parameterType.AsType(generationContext))

	funcDetails.AddReturns("error")
	funcDetails.AddComments(fn.declarationDocComment(receiver, parameterName))

	if fn.hubType.Equals(receiver) {
		// Body on the hub type pivots the conversion
		funcDetails.Body = fn.bodyForPivot(receiverName, parameterName)
	} else {
		// Body on non-hub type does one step of a conversion
		funcDetails.Body = fn.bodyForConvert(receiverName, parameterName, generationContext)
	}

	return funcDetails.DefineFunc()
}

// bodyForPivot is used to do the conversion if we hit the hub type without finding the conversion we need
//
// return instance.ConvertTo(s)
//
// Note that the method called is in the *other* *direction*; we restart the conversion at the extreme of the second
// spoke, invoking conversions back towards the hub again.
//
// TODO: If invoked with two unrelated instances (that don't share a common hub type), this will currently go into an
// infinite call loop, at least until it encounteres a stack overflow. Need to change this to panic if called while
// already active in order to provide a diagnostic.
func (fn *ChainedConversionFunction) bodyForPivot(receiverName string, parameterName string) []dst.Stmt {

	fnNameForOtherDirection := fn.direction.SelectString(fn.nameTo, fn.nameFrom)
	parameter := dst.NewIdent(parameterName)

	callAndReturn := astbuilder.Returns(
		astbuilder.CallExpr(dst.NewIdent(receiverName), fnNameForOtherDirection, parameter))

	return astbuilder.Statements(callAndReturn)
}

// bodyForConvert generates a conversion when the type we know about isn't the hub type, but is closer to it in our
// conversion graph.
//
// For ConvertFrom, we generate
//
// 	   src, ok := source.(*<intermediateType>)
// 	   if ok {
//         // Populate our instance from source
//         return s.AssignPropertiesFrom(source)
//     }
//
//     // Convert to an intermediate form
//     src = &<intermediateType>{}
//     err := src.ConvertFrom(source)
//     if err != nil {
//         return errors.Wrapf(err, "...elided...")
//     }
//
//     // Update our instance from src
//     return s.AssignPropertiesFrom(src)
//
// For ConvertTo, we have essentially the same structure, but two-step conversion is done in the other order.
//
func (fn *ChainedConversionFunction) bodyForConvert(
	receiverName string, parameterName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {

	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	receiver := dst.NewIdent(receiverName)
	parameter := dst.NewIdent(parameterName)
	local := dst.NewIdent(fn.localVariableId())
	errIdent := dst.NewIdent("err")

	intermediateType := fn.propertyAssignmentParameterType.AsType(generationContext)

	// <local>, ok := <parameter>.(<intermediateType>)
	typeAssert := astbuilder.TypeAssert(local, parameter, astbuilder.Dereference(intermediateType))

	// return <receiver>.AssignProperties(From|To)(<local>)
	directConversion := astbuilder.Returns(
		astbuilder.CallExpr(receiver, fn.propertyAssignmentFunctionName, local))
	astbuilder.AddComment(
		&directConversion.Decorations().Start,
		fn.direction.SelectString(
			fmt.Sprintf("// Populate our instance from %s", parameter),
			fmt.Sprintf("// Populate %s from our instance", parameter)))

	// if ok { ...elided... }
	returnDirectConversion := astbuilder.IfOk(
		directConversion)

	// <local> = &<intermediateType>{}
	initializeLocal := astbuilder.SimpleAssignment(
		local,
		token.ASSIGN,
		astbuilder.AddrOf(astbuilder.NewCompositeLiteralDetails(intermediateType).Build()))
	initializeLocal.Decs.Before = dst.EmptyLine
	astbuilder.AddComment(&initializeLocal.Decs.Start, "// Convert to an intermediate form")

	//
	// Depending on the direction of conversion either
	//
	//     err := <local>.ConvertFrom(<parameter>)
	// or
	//     err := <receiver>.AssignPropertiesTo(<local>)
	//
	initialStep := astbuilder.SimpleAssignment(
		errIdent,
		token.DEFINE,
		fn.direction.SelectExpr(
			astbuilder.CallExpr(local, fn.Name(), parameter),
			astbuilder.CallExpr(receiver, fn.propertyAssignmentFunctionName, local)))

	// if err != nil { ...elided...}
	checkInitialStepForError := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("initial step of conversion in %s()", fn.Name()))
	checkInitialStepForError.Decorations().After = dst.EmptyLine

	//
	// Depending on the direction of conversion, either
	//
	//     err = <receiver>.AssignPropertiesFrom(<local>)
	// or
	//     err = <local>.ConvertTo(<parameter>)
	//
	finalStep := astbuilder.SimpleAssignment(
		errIdent,
		token.ASSIGN,
		fn.direction.SelectExpr(
			astbuilder.CallExpr(receiver, fn.propertyAssignmentFunctionName, local),
			astbuilder.CallExpr(local, fn.Name(), parameter)))
	astbuilder.AddComment(
		&finalStep.Decorations().Start,
		fn.direction.SelectString(
			fmt.Sprintf("// Update our instance from %s", local),
			fmt.Sprintf("// Update %s from our instance", local)))

	// if err != nil { ...elided...}
	checkFinalStepForError := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("final step of conversion in %s()", fn.Name()))
	checkFinalStepForError.Decorations().After = dst.EmptyLine


	returnNil := astbuilder.Returns(astbuilder.Nil())

	return astbuilder.Statements(
		typeAssert,
		returnDirectConversion,
		initializeLocal,
		initialStep,
		checkInitialStepForError,
		finalStep,
		checkFinalStepForError,
		returnNil)
}

// localVariableId returns a good identifier to use for a local variable in our function,
// based which direction we are converting
func (fn *ChainedConversionFunction) localVariableId() string {
	return fn.direction.SelectString("src", "dst")
}

func (fn *ChainedConversionFunction) declarationDocComment(receiver astmodel.TypeName, parameter string) string {
	return fn.direction.SelectString(
		fmt.Sprintf("populates our %s from the provided %s", receiver.Name(), parameter),
		fmt.Sprintf("populates the provided %s from our %s", parameter, receiver.Name()))
}

func (fn *ChainedConversionFunction) Equals(otherFn astmodel.Function) bool {
	rcf, ok := otherFn.(*ChainedConversionFunction)
	if !ok {
		return false
	}

	return fn.Name() == rcf.Name() &&
		fn.direction == rcf.direction &&
		fn.propertyAssignmentFunctionName == rcf.propertyAssignmentFunctionName &&
		fn.propertyAssignmentParameterType.Equals(rcf.propertyAssignmentParameterType) &&
		fn.hubType.Equals(rcf.hubType)
}
