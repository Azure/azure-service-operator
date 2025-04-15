/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

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
	hub astmodel.InternalTypeName
	// propertyFunction is a reference to the function we will call to copy properties across
	propertyFunction *PropertyAssignmentFunction
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &ResourceConversionFunction{}

// NewResourceConversionFunction creates a conversion function that populates our hub type from the current instance
// hub is the TypeName of our hub type
// propertyFunction is the function we use to copy properties across
func NewResourceConversionFunction(
	hub astmodel.InternalTypeName,
	propertyFunction *PropertyAssignmentFunction,
	idFactory astmodel.IdentifierFactory,
) *ResourceConversionFunction {
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
		astmodel.ErisReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
	)

	// Include the package required by the parameter of the property assignment function
	// (as we may need to declare a variable of that type)
	result.AddReference(fn.propertyFunction.ParameterType().PackageReference())

	return result
}

func (fn *ResourceConversionFunction) References() astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet(fn.hub, fn.propertyFunction.ParameterType())
	return result
}

func (fn *ResourceConversionFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
) (*dst.FuncDecl, error) {
	// Create a sensible name for our receiver
	receiverName := fn.idFactory.CreateReceiver(receiver.Name())

	// We always use a pointer receiver, so we can modify it
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverTypeExpr,
		Name:          fn.Name(),
	}

	conversionPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeConversion)

	funcDetails.AddParameter("hub", astbuilder.QualifiedTypeName(conversionPackage, "Hub"))
	funcDetails.AddReturns("error")
	funcDetails.AddComments(fn.declarationDocComment(receiver))

	if astmodel.TypeEquals(fn.hub, fn.propertyFunction.ParameterType()) {
		// Not using an intermediate step
		body, err := fn.directConversion(receiverName, codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating direct conversion body")
		}

		funcDetails.Body = body
	} else {
		var body []dst.Stmt
		var err error
		fn.propertyFunction.direction.
			WhenFrom(func() { body, err = fn.indirectConversionFromHub(receiverName, codeGenerationContext) }).
			WhenTo(func() { body, err = fn.indirectConversionToHub(receiverName, codeGenerationContext) })
		if err != nil {
			return nil, eris.Wrap(err, "creating indirect conversion body")
		}

		funcDetails.Body = body
	}

	return funcDetails.DefineFunc(), nil
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
	receiverName string, generationContext *astmodel.CodeGenerationContext,
) ([]dst.Stmt, error) {
	fmtPackage := generationContext.MustGetImportedPackageName(astmodel.FmtReference)

	hubPackage := fn.hub.InternalPackageReference().FolderPath()
	localID := fn.localVariableID()
	localIdent := dst.NewIdent(localID)
	hubIdent := dst.NewIdent("hub")

	hubExpr, err := fn.hub.AsTypeExpr(generationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", fn.hub)
	}

	assignLocal := astbuilder.TypeAssert(
		localIdent,
		hubIdent,
		astbuilder.PointerTo(hubExpr))

	checkAssert := astbuilder.ReturnIfNotOk(
		astbuilder.FormatError(
			fmtPackage,
			fmt.Sprintf("expected %s/%s but received %%T instead", hubPackage, fn.Hub().Name()),
			hubIdent))

	copyAndReturn := astbuilder.Returns(
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), localIdent))
	copyAndReturn.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		assignLocal,
		checkAssert,
		copyAndReturn), nil
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
	receiverName string, generationContext *astmodel.CodeGenerationContext,
) ([]dst.Stmt, error) {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.ErisReference)
	localID := fn.localVariableID()
	errIdent := dst.NewIdent("err")

	intermediateType := fn.propertyFunction.ParameterType()
	intermediateTypeExpr, err := intermediateType.AsTypeExpr(generationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for intermediate type")
	}

	declareLocal := astbuilder.LocalVariableDeclaration(
		localID, intermediateTypeExpr, "// intermediate variable for conversion")
	declareLocal.Decorations().Before = dst.NewLine

	populateLocalFromHub := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallExpr(dst.NewIdent(localID), fn.Name(), dst.NewIdent("hub")))
	populateLocalFromHub.Decs.Before = dst.EmptyLine

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from hub to %s", localID))

	populateReceiverFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), astbuilder.AddrOf(dst.NewIdent(localID))))
	populateReceiverFromLocal.Decs.Before = dst.EmptyLine

	checkForErrorsPopulatingReceiver := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to %s", localID, receiverName))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))
	returnNil.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		declareLocal,
		populateLocalFromHub,
		checkForErrorsPopulatingLocal,
		populateReceiverFromLocal,
		checkForErrorsPopulatingReceiver,
		returnNil), nil
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
	receiverName string, generationContext *astmodel.CodeGenerationContext,
) ([]dst.Stmt, error) {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.ErisReference)
	localID := fn.localVariableID()
	errIdent := dst.NewIdent("err")

	intermediateType := fn.propertyFunction.ParameterType()
	intermediateTypeExpr, err := intermediateType.AsTypeExpr(generationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for intermediate type")
	}

	declareLocal := astbuilder.LocalVariableDeclaration(
		localID, intermediateTypeExpr, "// intermediate variable for conversion")
	declareLocal.Decorations().Before = dst.NewLine

	populateLocalFromReceiver := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunction.Name(), astbuilder.AddrOf(dst.NewIdent(localID))))

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting to %s from %s", localID, receiverName))

	populateHubFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		astbuilder.CallExpr(dst.NewIdent(localID), fn.Name(), dst.NewIdent("hub")))

	checkForErrorsPopulatingHub := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to hub", localID))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))
	returnNil.Decorations().Before = dst.EmptyLine

	return astbuilder.Statements(
		declareLocal,
		populateLocalFromReceiver,
		checkForErrorsPopulatingLocal,
		populateHubFromLocal,
		checkForErrorsPopulatingHub,
		returnNil), nil
}

// localVariableID returns a good identifier to use for a local variable in our function,
// based which direction we are converting
func (fn *ResourceConversionFunction) localVariableID() string {
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
