/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
)

// PivotConversionFunction implements a pivot that's used when a conversion reaches a hub type but hasn't yet found the
// innermost required conversion needed.
//
// We use this for Spec types, implementing genruntime.ConvertibleSpec, and for Status types, implementing
// genruntime.ConvertibleStatus.
//
// If the two instances involved in conversion are not in the same "spoke" leading to this "hub" version, we'll pivot to
// the reverse conversion, starting at the far end of that "spoke":
//
// func (s <hubtype>) ConvertFrom(instance <interfaceType>>) error {
//     return instance.ConvertTo(s)
// }
//
type PivotConversionFunction struct {
	// nameFrom is the name for this function when converting FROM a provided instance
	nameFrom string
	// nameTo is the name for this function when converting TO a provided instance
	nameTo string
	// parameterType is common interface type for our parameter
	parameterType astmodel.TypeName
	// direction is the direction of conversion - either FROM the supplied instance, or TO the supplied instance
	// We initialize this from the supplied property conversion function to guarantee consistency that the function
	// we generate is using that function correctly
	direction conversions.Direction
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &PivotConversionFunction{}

// NewSpecPivotConversionFunction creates a pivot conversion function that works to convert between two Spec types
// implementing the interface genruntime.ConvertibleSpec.
// hubType is the TypeName of our hub type
// idFactory is an identifier factory to use for generating local identifiers
func NewSpecPivotConversionFunction(
	direction conversions.Direction,
	idFactory astmodel.IdentifierFactory) *PivotConversionFunction {
	result := &PivotConversionFunction{
		nameFrom:      "ConvertSpecFrom",
		nameTo:        "ConvertSpecTo",
		parameterType: astmodel.ConvertibleSpecInterfaceType,
		direction:     direction,
		idFactory:     idFactory,
	}

	return result
}

// NewStatusPivotConversionFunction creates a pivot conversion function that works to convert between two Status types
// implementing the interface genruntime.ConvertibleStatus.
// hubType is the TypeName of our hub type
// idFactory is an identifier factory to use for generating local identifiers
func NewStatusPivotConversionFunction(
	direction conversions.Direction,
	idFactory astmodel.IdentifierFactory) *PivotConversionFunction {
	result := &PivotConversionFunction{
		nameFrom:      "ConvertStatusFrom",
		nameTo:        "ConvertStatusTo",
		parameterType: astmodel.ConvertibleStatusInterfaceType,
		direction:     direction,
		idFactory:     idFactory,
	}

	return result
}

func (fn *PivotConversionFunction) Name() string {
	return fn.direction.SelectString(fn.nameFrom, fn.nameTo)
}

func (fn *PivotConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		astmodel.GenRuntimeReference,
		fn.parameterType.PackageReference)
}

func (fn *PivotConversionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		fn.parameterType)
}

func (fn *PivotConversionFunction) AsFunc(
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
	funcDetails.Body = fn.bodyForPivot(receiverName, parameterName, generationContext)

	return funcDetails.DefineFunc()
}

// bodyForPivot is used to do the conversion if we hit the hub type without finding the conversion we need
//
// return instance.ConvertTo(s)
//
// Note that the method called is in the *other* *direction*; we restart the conversion at the extreme of the second
// spoke, invoking conversions back towards the hub again.
//
func (fn *PivotConversionFunction) bodyForPivot(
	receiverName string,
	parameterName string,
	generationContext *astmodel.CodeGenerationContext) []dst.Stmt {

	errorsPkg := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)

	fnNameForOtherDirection := fn.direction.SelectString(fn.nameTo, fn.nameFrom)
	parameter := dst.NewIdent(parameterName)
	receiver := dst.NewIdent(receiverName)

	errorMessage := astbuilder.StringLiteralf(
		"attempted conversion between unrelated implementations of %s",
		fn.parameterType)
	recursionCheck := astbuilder.ReturnIfExpr(
		astbuilder.Equal(parameter, receiver),
		astbuilder.CallQualifiedFunc(errorsPkg, "New", errorMessage))
	recursionCheck.Decorations().After = dst.EmptyLine

	callAndReturn := astbuilder.Returns(astbuilder.CallExpr(parameter, fnNameForOtherDirection, receiver))

	return astbuilder.Statements(recursionCheck, callAndReturn)
}

func (fn *PivotConversionFunction) declarationDocComment(receiver astmodel.TypeName, parameter string) string {
	return fn.direction.SelectString(
		fmt.Sprintf("populates our %s from the provided %s", receiver.Name(), parameter),
		fmt.Sprintf("populates the provided %s from our %s", parameter, receiver.Name()))
}

func (fn *PivotConversionFunction) Equals(otherFn astmodel.Function) bool {
	rcf, ok := otherFn.(*ChainedConversionFunction)
	if !ok {
		return false
	}

	return fn.Name() == rcf.Name() &&
		fn.direction == rcf.direction
}
