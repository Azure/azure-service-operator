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
	// hubSpec is the TypeName of the canonical hub spec type, the final target or original source for conversion
	hubSpec astmodel.TypeName
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
	// nameFrom is the name for this function when converting FROM a provided instance
	nameFrom string
	// nameTo is the name for this function when converting TO a provided instance
	nameTo string
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &ChainedConversionFunction{}

// NewSpecConversionFunction creates a conversion function that populates our hub spec type from the current instance
// hubSpec is the TypeName of our hub type
// propertyFunction is the function we will call to copy properties across between instances
func NewSpecConversionFunction(
	hubSpec astmodel.TypeName,
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ChainedConversionFunction {
	result := &ChainedConversionFunction{
		hubSpec:                         hubSpec,
		propertyAssignmentFunctionName:  propertyFunction.Name(),
		propertyAssignmentParameterType: propertyFunction.otherDefinition.Name(),
		direction:                       propertyFunction.direction,
		nameFrom:                        "ConvertSpecFrom",
		nameTo:                          "ConvertSpecTo",
		idFactory:                       idFactory,
	}

	return result
}

func (fn *ChainedConversionFunction) Name() string {
	return fn.direction.SelectString(fn.nameFrom, fn.nameTo)
}

func (fn *ChainedConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		astmodel.GenRuntimeReference,
		fn.hubSpec.PackageReference)

	result.AddReference(fn.propertyAssignmentParameterType.PackageReference)

	return result
}

func (fn *ChainedConversionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		fn.hubSpec,
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

	funcDetails.AddParameter(parameterName, astmodel.ConvertibleSpecInterfaceType.AsType(generationContext))
	funcDetails.AddReturns("error")
	funcDetails.AddComments(fn.declarationDocComment(receiver, parameterName))

	if fn.hubSpec.Equals(receiver) {
		// Body on the hub spec pivots the conversion
		funcDetails.Body = fn.bodyForPivot(receiverName, parameterName, generationContext)
	} else {
		// Body on non-hub spec does a conversion
		funcDetails.Body = fn.bodyForConvert(receiverName, parameterName, generationContext)
	}

	return funcDetails.DefineFunc()
}

// bodyForPivot is used to do the conversion if we hit the hub package version without finding the conversion we need
//
// return spec.ConvertTo(s)
//
// Note that the method called is in the *other* *direction*; we restart the conversion at the extreme of the second
// spoke, claling
func (fn *ChainedConversionFunction) bodyForPivot(
	receiverName string, parameterName string, _ *astmodel.CodeGenerationContext) []dst.Stmt {

	fnNameForOtherDirection := fn.direction.SelectString(fn.nameTo, fn.nameFrom)
	parameter := dst.NewIdent(parameterName)

	callAndReturn := astbuilder.Returns(
		astbuilder.CallExpr(dst.NewIdent(receiverName), fnNameForOtherDirection, parameter))

	return astbuilder.Statements(callAndReturn)
}

// bodyForConvert generates a conversion when the type we know about isn't the hub spec type, but is closer to it in our
// conversion graph.
//
// For ConvertFrom, we generate
//
// 	   src, ok := source.(*<intermediateType>)
// 	   if ok {
//         // Populate our spec from source
//         return s.AssignPropertiesFrom(source)
//     }
//
//     // Convert to an intermediate form
//     src = &<intermediateType>{}
//     err := src.ConvertFromSpec(source)
//     if err != nil {
//         return errors.Wrapf(err, "...elided...")
//     }
//
//     // Update our spec from src
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
			fmt.Sprintf("// Populate our spec from %s", parameter),
			fmt.Sprintf("// Populate %s from our spec", parameter)))

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
	//     err := <local>.ConvertFromSpec(<parameter>)
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
	checkForError := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("initial step of conversion in %s()", fn.Name()))
	checkForError.Decorations().After = dst.EmptyLine

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
			fmt.Sprintf("// Update our spec from %s", local),
			fmt.Sprintf("// Update %s from our spec", local)))

	returnErr := astbuilder.Returns(errIdent)

	return astbuilder.Statements(
		typeAssert,
		returnDirectConversion,
		initializeLocal,
		initialStep,
		checkForError,
		finalStep,
		returnErr)
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
		fn.hubSpec.Equals(rcf.hubSpec)
}
