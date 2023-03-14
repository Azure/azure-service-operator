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

// ResourceConversionFunction implements conversions to/from our hub type
// Existing PropertyAssignment functions are used to implement stepwise conversion
//
// Direct conversion from the hub type:
//
//	func (<receiver> <receiverType>) Convert<From|To>(hub conversion.Hub) error {
//		source, ok := hub.(*<hub.Type>)
//		if !ok {
//			return fmt.Errorf("expected <hub.Type> but received %T instead", hub)
//		}
//		return <receiver>.AssignProperties<From|To><Type>(source)
//	}
//
// Indirect conversion, multiple steps via an intermediate instance
//
//	func (r <receiver>) Convert<From|To>(hub conversion.Hub) error {
//		var source <vNext>
//		err := source.Convert<From|To>(hub)
//		if err != nil {
//			return errors.Wrap(err, "converting from hub to source")
//		}
//
//		err = <receiver>.AssignProperties<From|To><Type>(&source)
//		if err != nil {
//			return errors.Wrap(err, "converting from source to <type>")
//		}
//
//		return nil
//	}
type ResourceConversionFunction struct {
	// hub is the TypeName of the canonical hub type, the final target or original source for conversion
	hub astmodel.TypeName
	// propertyFunction is a reference to the function we will call to copy properties across
	propertyFunction *PropertyAssignmentFunction
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &ResourceConversionFunction{}

// NewResourceConversionFunction creates a conversion function that populates our hub type from the current instance
// hub is the TypeName of our hub type
// propertyFuntion is the function we use to copy properties across
func NewResourceConversionFunction(
	hub astmodel.TypeName,
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory) *ResourceConversionFunction {
	result := &ResourceConversionFunction{
		hub:              hub,
		propertyFunction: propertyFunction,
		idFactory:        idFactory,
	}

	return result
}

func (fn *ResourceConversionFunction) Name() string {
	return fn.propertyFunction.direction.SelectString("ConvertFrom", "ConvertTo")
}

func (fn *ResourceConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		fn.hub.PackageReference)

	// Include the package required by the parameter of the property assignment function
	result.AddReference(fn.propertyFunction.ParameterType().PackageReference)

	return result
}

func (fn *ResourceConversionFunction) References() astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet(fn.hub, fn.propertyFunction.ParameterType())
	return result
}

func (fn *ResourceConversionFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {

	// Create a sensible name for our receiver
	receiverName := fn.idFactory.CreateReceiver(receiver.Name())

	// We always use a pointer receiver so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverType,
		Name:          fn.Name(),
	}

	conversionPackage := generationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeConversion)

	funcDetails.AddParameter("hub", astbuilder.QualifiedTypeName(conversionPackage, "Hub"))
	funcDetails.AddReturns("error")
	funcDetails.AddComments(fn.declarationDocComment(receiver))

	if astmodel.TypeEquals(fn.hub, fn.propertyFunction.ParameterType()) {
		// Not using an intermediate step
		funcDetails.Body = fn.directConversion(receiverName, generationContext)
	} else {
		fn.propertyFunction.direction.
			WhenFrom(func() { funcDetails.Body = fn.indirectConversionFromHub(receiverName, generationContext) }).
			WhenTo(func() { funcDetails.Body = fn.indirectConversionToHub(receiverName, generationContext) })
	}

	return funcDetails.DefineFunc()
}

// Direction returns this functions direction of conversion
func (fn *ResourceConversionFunction) Direction() conversions.Direction {
	return fn.propertyFunction.Direction()
}

// Hub returns the hub type we convert with
func (fn *ResourceConversionFunction) Hub() astmodel.TypeName {
	return fn.hub
}

// directConversion creates a simple direct conversion between the two types
//
// <local>, ok := <hubAsInterface>.(<actualHubType>)
//
//	if !ok {
//	    return errors.Errorf("expected <actualHubType> but received %T instead", <hubAsInterface>)
//	}
//
// return <receiver>.AssignProperties(To|From)<type>(<local>)
func (fn *ResourceConversionFunction) directConversion(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	fmtPackage := generationContext.MustGetImportedPackageName(astmodel.FmtReference)

	hubGroup, hubVersion := fn.hub.PackageReference.GroupVersion()
	localId := fn.localVariableId()
	localIdent := dst.NewIdent(localId)
	hubIdent := dst.NewIdent("hub")

	assignLocal := astbuilder.TypeAssert(
		localIdent,
		hubIdent,
		astbuilder.PointerTo(fn.hub.AsType(generationContext)))

	checkAssert := astbuilder.ReturnIfNotOk(
		astbuilder.FormatError(
			fmtPackage,
			fmt.Sprintf("expected %s/%s/%s but received %%T instead", hubGroup, hubVersion, fn.Hub().Name()),
			hubIdent))

	copyAndReturn := astbuilder.Returns(
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), localIdent))
	copyAndReturn.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		assignLocal,
		checkAssert,
		copyAndReturn)
}

// indirectConversionFromHub generates a conversion when the type we know about isn't the hub type, but is closer to it
// in our conversion graph
//
// var source <intermediateType>
// err := source.ConvertFrom(<hub>)
//
//	if err != nil {
//	    return errors.Wrap(err, "converting from hub to source")
//	}
//
// err = <receiver>.AssignPropertiesFrom<type>(&source)
//
//	if err != nil {
//	    return errors.Wrap(err, "converting from source to <type>")
//	}
//
// return nil
func (fn *ResourceConversionFunction) indirectConversionFromHub(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)
	localId := fn.localVariableId()
	errIdent := dst.NewIdent("err")

	intermediateType := fn.propertyFunction.ParameterType()

	declareLocal := astbuilder.LocalVariableDeclaration(
		localId, intermediateType.AsType(generationContext), "// intermediate variable for conversion")
	declareLocal.Decorations().Before = dst.NewLine

	populateLocalFromHub := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallExpr(dst.NewIdent(localId), fn.Name(), dst.NewIdent("hub")))
	populateLocalFromHub.Decs.Before = dst.EmptyLine

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from hub to %s", localId))

	populateReceiverFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), astbuilder.AddrOf(dst.NewIdent(localId))))
	populateReceiverFromLocal.Decs.Before = dst.EmptyLine

	checkForErrorsPopulatingReceiver := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to %s", localId, receiverName))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))
	returnNil.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		declareLocal,
		populateLocalFromHub,
		checkForErrorsPopulatingLocal,
		populateReceiverFromLocal,
		checkForErrorsPopulatingReceiver,
		returnNil)
}

// indirectConversionToHub generates a conversion when the type we know about isn't the hub type, but is closer to it in
// our conversion graph
//
// var destination <intermediateType>
// err = <receiver>.AssignPropertiesTo<type>(&destination)
//
//	if err != nil {
//	    return errors.Wrap(err, "converting to destination from <type>")
//	}
//
// err := destination.ConvertTo(<hub>)
//
//	if err != nil {
//	    return errors.Wrap(err, "converting from destination to hub")
//	}
//
// return nil
func (fn *ResourceConversionFunction) indirectConversionToHub(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)
	localId := fn.localVariableId()
	errIdent := dst.NewIdent("err")

	intermediateType := fn.propertyFunction.ParameterType()

	declareLocal := astbuilder.LocalVariableDeclaration(
		localId, intermediateType.AsType(generationContext), "// intermediate variable for conversion")
	declareLocal.Decorations().Before = dst.NewLine

	populateLocalFromReceiver := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), astbuilder.AddrOf(dst.NewIdent(localId))))

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting to %s from %s", localId, receiverName))

	populateHubFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		astbuilder.CallExpr(dst.NewIdent(localId), fn.Name(), dst.NewIdent("hub")))

	checkForErrorsPopulatingHub := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to hub", localId))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))
	returnNil.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		declareLocal,
		populateLocalFromReceiver,
		checkForErrorsPopulatingLocal,
		populateHubFromLocal,
		checkForErrorsPopulatingHub,
		returnNil)
}

// localVariableId returns a good identifier to use for a local variable in our function,
// based which direction we are converting
func (fn *ResourceConversionFunction) localVariableId() string {
	return fn.propertyFunction.direction.SelectString("source", "destination")
}

func (fn *ResourceConversionFunction) declarationDocComment(receiver astmodel.TypeName) string {
	return fn.propertyFunction.direction.SelectString(
		fmt.Sprintf("populates our %s from the provided hub %s", receiver.Name(), fn.hub.Name()),
		fmt.Sprintf("populates the provided hub %s from our %s", fn.hub.Name(), receiver.Name()))
}

func (fn *ResourceConversionFunction) Equals(otherFn astmodel.Function, override astmodel.EqualityOverrides) bool {
	rcf, ok := otherFn.(*ResourceConversionFunction)
	if !ok {
		return false
	}

	if !fn.propertyFunction.Equals(rcf.propertyFunction, override) {
		return false
	}

	return fn.Name() == rcf.Name() &&
		fn.hub.Equals(rcf.hub, override)
}
