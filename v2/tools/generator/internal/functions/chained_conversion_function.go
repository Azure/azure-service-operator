/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
)

// ChainedConversionFunction implements conversions to/from another type that implements a conversion function
//
// We use this for Spec types, implementing genruntime.ConvertibleSpec, and for Status types, implementing
// genruntime.ConvertibleStatus
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
type ChainedConversionFunction struct {
	// name is the unique name for this function
	name string
	// parameterType is common interface type for our parameter
	parameterType astmodel.TypeName
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

// NewSpecChainedConversionFunction creates a chained conversion function that converts between two Spec types implementing the
// interface genruntime.ConvertibleSpec, using the provided property assignment function as a basis.
// hubType is the TypeName of our hub type
// propertyFunction is the function we will call to copy properties across between specs
// idFactory is an identifier factory to use for generating local identifiers
func NewSpecChainedConversionFunction(
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ChainedConversionFunction {
	result := &ChainedConversionFunction{
		name:                            propertyFunction.direction.SelectString("ConvertSpecFrom", "ConvertSpecTo"),
		parameterType:                   astmodel.ConvertibleSpecInterfaceType,
		propertyAssignmentFunctionName:  propertyFunction.Name(),
		propertyAssignmentParameterType: propertyFunction.ParameterType(),
		direction:                       propertyFunction.direction,
		idFactory:                       idFactory,
	}

	return result
}

// NewStatusChainedConversionFunction creates a chained conversion function that converts between two Status types implementing
// the interface genruntime.ConvertibleStatus, using the provided property assignment function as a basis.
// hubType is the TypeName of our hub type
// propertyFunction is the function we will call to copy properties across between specs
// idFactory is an identifier factory to use for generating local identifiers
func NewStatusChainedConversionFunction(
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ChainedConversionFunction {
	result := &ChainedConversionFunction{
		name:                            propertyFunction.direction.SelectString("ConvertStatusFrom", "ConvertStatusTo"),
		parameterType:                   astmodel.ConvertibleStatusInterfaceType,
		propertyAssignmentFunctionName:  propertyFunction.Name(),
		propertyAssignmentParameterType: propertyFunction.ParameterType(),
		direction:                       propertyFunction.direction,
		idFactory:                       idFactory,
	}

	return result
}

func (fn *ChainedConversionFunction) Name() string {
	return fn.name
}

func (fn *ChainedConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		astmodel.GenRuntimeReference,
		fn.parameterType.PackageReference,
		fn.propertyAssignmentParameterType.PackageReference)
}

func (fn *ChainedConversionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		fn.parameterType,
		fn.propertyAssignmentParameterType)
}

func (fn *ChainedConversionFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {

	// Create a sensible name for our receiver
	receiverName := fn.idFactory.CreateReceiver(receiver.Name())

	// We always use a pointer receiver, so we can modify it
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
	funcDetails.Body = fn.bodyForConvert(receiverName, parameterName, generationContext)

	return funcDetails.DefineFunc()
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
		astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(intermediateType).Build()))
	initializeLocal.Decs.Before = dst.EmptyLine
	astbuilder.AddComment(&initializeLocal.Decs.Start, "// Convert to an intermediate form")

	//
	// Depending on the direction of conversion either
	//
	//     err := <local>.ConvertFrom(<parameter>)
	// or
	//     err := <receiver>.AssignPropertiesTo(<local>)
	//
	initialStep := astbuilder.ShortDeclaration(
		"err",
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

func (fn *ChainedConversionFunction) Equals(otherFn astmodel.Function, override astmodel.EqualityOverrides) bool {
	rcf, ok := otherFn.(*ChainedConversionFunction)
	if !ok {
		return false
	}

	return fn.Name() == rcf.Name() &&
		fn.direction == rcf.direction &&
		fn.propertyAssignmentFunctionName == rcf.propertyAssignmentFunctionName &&
		fn.propertyAssignmentParameterType.Equals(rcf.propertyAssignmentParameterType, override)
}
