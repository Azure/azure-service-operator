/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
	"go/token"
)

// HubConversionFunction implements conversions to/from our hub type
// Existing PropertyAssignment functions are used to implement stepwise conversion
type HubConversionFunction struct {
	// name of this conversion function
	name string
	// hub is the TypeName of the canonical hub type, the final target or original source for conversion
	hub astmodel.TypeName
	// direction specifies whether we are converting to the hub type, or from it
	direction Direction
	// propertyFunctionName is the name of the function we call to copy properties across
	propertyFunctionName string
	// intermediateType is the TypeName of an intermediate type we use as part of a multiple step conversion
	// If nil, we are able to convert directly to/from the hub type
	intermediateType *astmodel.TypeName
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &HubConversionFunction{}

// NewConversionToHubFunction creates a conversion function that populates our hub type from the current instance
func NewConversionToHubFunction(
	hub astmodel.TypeName,
	intermediateType astmodel.TypeName,
	propertyFunctionName string,
	idFactory astmodel.IdentifierFactory) *HubConversionFunction {
	result := &HubConversionFunction{
		name:                 "ConvertTo",
		hub:                  hub,
		direction:            ConvertTo,
		propertyFunctionName: propertyFunctionName,
		idFactory:            idFactory,
	}

	if !hub.Equals(intermediateType) {
		result.intermediateType = &intermediateType
	}

	return result
}

// NewConversionFromHubFunction creates a conversion function that populates the current instance from our hub type
func NewConversionFromHubFunction(
	hub astmodel.TypeName,
	intermediateType astmodel.TypeName,
	propertyFunctionName string,
	idFactory astmodel.IdentifierFactory) *HubConversionFunction {
	result := &HubConversionFunction{
		name:                 "ConvertFrom",
		hub:                  hub,
		direction:            ConvertFrom,
		propertyFunctionName: propertyFunctionName,
		idFactory:            idFactory,
	}

	if !hub.Equals(intermediateType) {
		result.intermediateType = &intermediateType
	}

	return result
}

func (fn *HubConversionFunction) Name() string {
	return fn.direction.SelectString("ConvertFrom", "ConvertTo")
}

func (fn *HubConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.ControllerRuntimeConversion,
		astmodel.FmtReference,
		fn.hub.PackageReference)

	if fn.intermediateType != nil {
		result.AddReference(fn.intermediateType.PackageReference)
	}

	return result

}

func (fn *HubConversionFunction) References() astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet(fn.hub)

	if fn.intermediateType != nil {
		result.Add(*fn.intermediateType)
	}

	return result
}

func (fn *HubConversionFunction) AsFunc(
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

	conversionPackage := generationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeConversion)

	funcDetails.AddParameter("hub", astbuilder.QualifiedTypeName(conversionPackage, "Hub"))
	funcDetails.AddReturns("error")
	funcDetails.AddComments(fn.declarationDocComment(receiver))

	if fn.intermediateType == nil {
		funcDetails.Body = fn.DirectConversion(receiverName, generationContext)
	} else {
		fn.direction.
			WhenFrom(func() { funcDetails.Body = fn.IndirectConversionFromHub(receiverName, generationContext) }).
			WhenTo(func() { funcDetails.Body = fn.IndirectConversionToHub(receiverName, generationContext) })
	}

	return funcDetails.DefineFunc()
}

// DirectConversion creates a simple direct conversion between the two types
//
// <local>, ok := <hubAsInterface>.(<actualHubType>)
// if !ok {
//     return errors.Errorf("expected <actualHubType> but received %T instead", <hubAsInterface>)
// }
//
// return receiver.AssignProperties(To|From)<type>(<local>)
//
func (fn *HubConversionFunction) DirectConversion(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	fmtPackage := generationContext.MustGetImportedPackageName(astmodel.FmtReference)

	localId := fn.localVariableId()
	localIdent := dst.NewIdent(localId)
	hubIdent := dst.NewIdent("hub")

	assignLocal := astbuilder.TypeAssert(
		localIdent,
		hubIdent,
		astbuilder.Dereference(fn.hub.AsType(generationContext)))

	checkAssert := astbuilder.ReturnIfNotOk(
		astbuilder.FormatError(
			fmtPackage,
			fmt.Sprintf("expected %s but received %%T instead", fn.hub),
			hubIdent))

	copyAndReturn := astbuilder.Returns(
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunctionName, localIdent))

	return astbuilder.Statements(
		assignLocal,
		checkAssert,
		copyAndReturn)
}

// IndirectConversionFromHub generates a conversion when the type we know about isn't the hub type, but is closer to it
// in our conversion graph
//
// var source <intermediateType>
// err := source.ConvertFrom(<hub>)
// if err != nil {
//     return errors.Wrap(err, "converting from hub to source")
// }
// err = <receiver>.AssignPropertiesFrom<type>(&source)
// if err != nil {
//     return errors.Wrap(err, "converting from source to <type>")
// }
// return nil
//
func (fn *HubConversionFunction) IndirectConversionFromHub(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)
	localId := fn.localVariableId()
	errIdent := dst.NewIdent("err")

	declareLocal := astbuilder.LocalVariableDeclaration(
		localId, fn.intermediateType.AsType(generationContext), "// intermediate variable for conversion")

	populateLocalFromHub := astbuilder.SimpleAssignment(
		errIdent,
		token.DEFINE,
		astbuilder.CallExpr(dst.NewIdent(localId), fn.Name(), dst.NewIdent("hub")))

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from hub to %s", localId))

	populateReceiverFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		token.ASSIGN,
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunctionName, astbuilder.AddrOf(dst.NewIdent(localId))))

	checkForErrorsPopulatingReceiver := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to %s", localId, receiverName))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))

	return astbuilder.Statements(
		declareLocal,
		populateLocalFromHub,
		checkForErrorsPopulatingLocal,
		populateReceiverFromLocal,
		checkForErrorsPopulatingReceiver,
		returnNil)
}

// IndirectConversionToHub generates a conversion when the type we know about isn't the hub type, but is closer to it in
// our conversion graph
//
// var destination <intermediateType>
// err = <receiver>.AssignPropertiesTo<type>(&destination)
// if err != nil {
//     return errors.Wrap(err, "converting to destination from <type>")
// }
// err := destination.ConvertTo(<hub>)
// if err != nil {
//     return errors.Wrap(err, "converting from destination to hub")
// }
// return nil
//
func (fn *HubConversionFunction) IndirectConversionToHub(
	receiverName string, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
	errorsPackage := generationContext.MustGetImportedPackageName(astmodel.GitHubErrorsReference)
	localId := fn.localVariableId()
	errIdent := dst.NewIdent("err")

	declareLocal := astbuilder.LocalVariableDeclaration(
		localId, fn.intermediateType.AsType(generationContext), "// intermediate variable for conversion")

	populateLocalFromReceiver := astbuilder.SimpleAssignment(
		errIdent,
		token.DEFINE,
		astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunctionName, astbuilder.AddrOf(dst.NewIdent(localId))))

	checkForErrorsPopulatingLocal := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting to %s from %s", localId, receiverName))

	populateHubFromLocal := astbuilder.SimpleAssignment(
		errIdent,
		token.ASSIGN,
		astbuilder.CallExpr(dst.NewIdent(localId), fn.Name(), dst.NewIdent("hub")))

	checkForErrorsPopulatingHub := astbuilder.CheckErrorAndWrap(
		errorsPackage,
		fmt.Sprintf("converting from %s to hub", localId))

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))

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
func (fn *HubConversionFunction) localVariableId() string {
	return fn.direction.SelectString("source", "destination")
}

func (fn *HubConversionFunction) declarationDocComment(receiver astmodel.TypeName) string {
	return fn.direction.SelectString(
		fmt.Sprintf("populates our %s from the provided hub %s", receiver.Name(), fn.hub.Name()),
		fmt.Sprintf("populates the provided hub %s from our %s", fn.hub.Name(), receiver.Name()))
}

func (fn *HubConversionFunction) Equals(otherFn astmodel.Function) bool {
	hcf, ok := otherFn.(*HubConversionFunction)
	if !ok {
		return false
	}

	if fn.intermediateType != hcf.intermediateType {
		if fn.intermediateType == nil || hcf.intermediateType == nil {
			return false
		}

		if !fn.intermediateType.Equals(hcf.intermediateType) {
			return false
		}
	}

	return fn.name == hcf.name &&
		fn.direction == hcf.direction &&
		fn.hub.Equals(hcf.hub)
}
