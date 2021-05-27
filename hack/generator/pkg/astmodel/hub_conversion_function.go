package conversions

import (
	"fmt"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
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
	// propertyFunction is the name of the function we call to copy properties across
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
	return fn.name
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

	var localId string
	var declarationComment string
	switch fn.direction {
	case ConvertFrom:
		localId = "source"
		declarationComment = fmt.Sprintf("populates our %s from the provided hub %s", receiver.Name(), fn.hub.Name())
	case ConvertTo:
		localId = "destination"
		declarationComment = fmt.Sprintf("populates the provided hub %s from our %s", fn.hub.Name(), receiver.Name())
	default:
		panic(fmt.Sprintf("unexpected conversion direction %q", fn.direction))
	}

	conversionPackage := generationContext.MustGetImportedPackageName(astmodel.ControllerRuntimeConversion)
	fmtPackage := generationContext.MustGetImportedPackageName(astmodel.FmtReference)

	funcDetails.AddParameter("hub", astbuilder.QualifiedTypeName(conversionPackage, "Hub"))
	funcDetails.AddReturns("error")
	funcDetails.AddComments(declarationComment)

	if fn.intermediateType == nil {
		// Simple case, direct conversion
		local := dst.NewIdent(localId)

		assignLocal := astbuilder.TypeAssert(
			local,
			dst.NewIdent("hub"),
			astbuilder.Dereference(fn.hub.AsType(generationContext)))

		checkAssert := astbuilder.ReturnIfNotOk(
			astbuilder.FormatError(
				fmtPackage,
				fmt.Sprintf("expected %s but received %%T instead", fn.hub),
				dst.NewIdent("hub")))

		copyAndReturn := astbuilder.Returns(
			astbuilder.CallExpr(dst.NewIdent(receiverName), fn.propertyFunctionName, local))

		funcDetails.Body = astbuilder.Statements(assignLocal, checkAssert, copyAndReturn)
	}

	return funcDetails.DefineFunc()
}

// generateBodyForConvertFrom returns all of the statements required for the conversion function
// receiver is an expression for access our receiver type, used to qualify field access
// parameter is an expression for access to our parameter passed to the function, also used for field access
// generationContext is our code generation context, passed to allow resolving of identifiers in other packages
func (fn *HubConversionFunction) generateBodyForConvertFrom(
	receiver string,
	parameter dst.Expr,
	_ *astmodel.CodeGenerationContext,
) []dst.Stmt {
	copyProperties := astbuilder.LocalVariableDeclaration(
		"err",
		astbuilder.CallExpr(dst.NewIdent(receiver), fn.propertyFunctionName, parameter),
		"read properties from the hub type")

	handleError := astbuilder.CheckErrorAndReturn()

	returnNil := astbuilder.Returns(dst.NewIdent("nil"))

	return []dst.Stmt{
		copyProperties,
		handleError,
		returnNil,
	}
}

// generateBodyForConvertTo returns all of the statements required for the conversion function
// receiver is an expression for access our receiver type, used to qualify field access
// parameter is an expression for access to our parameter passed to the function, also used for field access
// generationContext is our code generation context, passed to allow resolving of identifiers in other packages
func (fn *HubConversionFunction) generateBodyForConvertTo(
	receiver string,
	parameter dst.Expr,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	if fn.intermediateType == nil {
		// Direct transformation
		destination := dst.NewIdent("destination")
		assignDestination := astbuilder.TypeAssert(
			destination, parameter, fn.hub.AsType(generationContext))
		copyProperties := astbuilder.LocalVariableDeclaration(
			"err",
			astbuilder.CallExpr(dst.NewIdent(receiver), fn.propertyFunctionName, destination),
			"read properties from the hub type")

		handleError := astbuilder.CheckErrorAndReturn()

		returnNil := astbuilder.Returns(dst.NewIdent("nil"))
		return astbuilder.Statements(assignDestination, copyProperties, handleError, returnNil)
	}

	return astbuilder.Statements()
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
