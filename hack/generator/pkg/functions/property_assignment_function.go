/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"go/token"
	"sort"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
)

// PropertyAssignmentFunction represents a function that assigns all the properties from one resource or object to
// another. Performs a single step of the conversions required to/from the hub version.
type PropertyAssignmentFunction struct {
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
	// knownLocals is a cached set of local identifiers that have already been used, to avoid conflicts
	knownLocals *astmodel.KnownLocalsSet
	// conversionContext is additional information about the context in which this conversion was made
	conversionContext *conversions.PropertyConversionContext
	// identifier to use for our receiver in generated code
	receiverName string
	// identifier to use for our parameter in generated code
	parameterName string
	// identifier to use for the property bag local variable in generated code
	propertyBagName string
}

// StoragePropertyConversion represents a function that generates the correct AST to convert a single property value
// Different functions will be used, depending on the types of the properties to be converted.
// source is an expression that returns the source we are converting from (a Resource or other Object)
// destination is an expression that returns the destination we are converting to (again, a Resource or other Object)
// The function returns a sequence of statements to carry out the stated conversion/copy
type StoragePropertyConversion func(
	source dst.Expr, destination dst.Expr, generationContext *astmodel.CodeGenerationContext) []dst.Stmt

// Ensure that PropertyAssignmentFunction implements Function
var _ astmodel.Function = &PropertyAssignmentFunction{}

// NewPropertyAssignmentFunction creates a new PropertyAssignmentFunction to convert with the specified type
// receiver is the type definition that will be the receiver for this function
// otherDefinition is the type definition to convert TO or FROM
// conversionContext is our context for creating a conversion
// direction specifies whether we are converting TO or FROM the other definition
func NewPropertyAssignmentFunction(
	receiver astmodel.TypeDefinition,
	otherDefinition astmodel.TypeDefinition,
	conversionContext *conversions.PropertyConversionContext,
	direction conversions.Direction,
) (*PropertyAssignmentFunction, error) {
	idFactory := conversionContext.IDFactory()

	result := &PropertyAssignmentFunction{
		receiverDefinition: receiver,
		otherDefinition:    otherDefinition,
		idFactory:          idFactory,
		direction:          direction,
		conversions:        make(map[string]StoragePropertyConversion),
		knownLocals:        astmodel.NewKnownLocalsSet(idFactory),
		receiverName:       idFactory.CreateIdentifier(receiver.Name().Name(), astmodel.NotExported),
		parameterName:      direction.SelectString("source", "destination"),
	}

	result.propertyBagName = result.knownLocals.CreateLocal("propertyBag", "", "Local")

	result.conversionContext = conversionContext.WithFunctionName(result.Name()).
		WithKnownLocals(result.knownLocals).
		WithDirection(direction)

	err := result.createConversions(receiver)
	if err != nil {
		return nil, errors.Wrapf(err, "creating '%s()'", result.Name())
	}

	return result, nil
}

// Name returns the name of this function
func (fn *PropertyAssignmentFunction) Name() string {
	return conversions.NameOfPropertyAssignmentFunction(fn.otherDefinition.Name(), fn.direction, fn.idFactory)
}

// RequiredPackageReferences returns the set of package references required by this function
func (fn *PropertyAssignmentFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		astmodel.GenRuntimeReference,
		fn.otherDefinition.Name().PackageReference)

	return result
}

// References returns the set of types referenced by this function
func (fn *PropertyAssignmentFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(fn.otherDefinition.Name())
}

// Equals checks to see if the supplied function is the same as this one
func (fn *PropertyAssignmentFunction) Equals(f astmodel.Function) bool {
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
		fmt.Sprintf("populates our %s from the provided source %s", receiver.Name(), fn.otherDefinition.Name().Name()),
		fmt.Sprintf("populates the provided destination %s from our %s", fn.otherDefinition.Name().Name(), receiver.Name()))

	// We always use a pointer receiver so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: fn.receiverName,
		ReceiverType:  receiverType,
		Name:          fn.Name(),
		Body:          fn.generateBody(fn.receiverName, fn.parameterName, generationContext),
	}

	parameterPackage := generationContext.MustGetImportedPackageName(fn.otherDefinition.Name().PackageReference)

	funcDetails.AddParameter(
		fn.parameterName,
		&dst.StarExpr{
			X: astbuilder.Selector(dst.NewIdent(parameterPackage), fn.otherDefinition.Name().Name()),
		})

	funcDetails.AddReturns("error")
	funcDetails.AddComments(description)

	return funcDetails.DefineFunc()
}

// generateBody returns all of the statements required for the conversion function
// receiver is an expression for access our receiver type, used to qualify field access
// parameter is an expression for access to our parameter passed to the function, also used for field access
// generationContext is our code generation context, passed to allow resolving of identifiers in other packages
func (fn *PropertyAssignmentFunction) generateBody(
	receiver string,
	parameter string,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	source := fn.direction.SelectString(parameter, receiver)
	destination := fn.direction.SelectString(receiver, parameter)

	bagPreamble := fn.propertyBagPrologue(source, generationContext)
	assignments := fn.generateAssignments(dst.NewIdent(source), dst.NewIdent(destination), generationContext)
	bagPostamble := fn.propertyBagEpilogue(destination)

	return astbuilder.Statements(
		bagPreamble,
		assignments,
		bagPostamble,
		astbuilder.ReturnNoError())
}

// propertyBagPrologue creates any introductory statements needed to set up our property bag before we start doing
// assignments. We need to handle three cases:
//   o If our source has a property bag, we clone it.
//   o If our destination has a property bag (and our source does not), we create a new one.
//   o If neither source nor destination has a property bag, we don't need to do anything.
// source is the name of the source to read the property bag from
func (fn *PropertyAssignmentFunction) propertyBagPrologue(
	source string,
	generationContext *astmodel.CodeGenerationContext) []dst.Stmt {

	if srcBag := fn.findPropertyBag(fn.sourceType()); srcBag != nil {
		cloneBag := astbuilder.SimpleAssignment(
			dst.NewIdent(fn.propertyBagName),
			token.DEFINE,
			astbuilder.CallExpr(
				astbuilder.Selector(dst.NewIdent(source), string(srcBag.PropertyName())),
				"Clone"))
		cloneBag.Decs.Before = dst.NewLine
		astbuilder.AddComment(&cloneBag.Decorations().Start, "// Clone the existing property bag")

		return astbuilder.Statements(cloneBag)
	}

	if dstBag := fn.findPropertyBag(fn.destinationType()); dstBag != nil {
		genruntimePkg := generationContext.MustGetImportedPackageName(astmodel.GenRuntimeReference)
		createBag := astbuilder.SimpleAssignment(
			dst.NewIdent(fn.propertyBagName),
			token.DEFINE,
			astbuilder.CallQualifiedFunc(genruntimePkg, "NewPropertyBag"))
		createBag.Decs.Before = dst.NewLine
		astbuilder.AddComment(&createBag.Decorations().Start, "// Create a new property bag")

		return astbuilder.Statements(createBag)
	}

	return nil
}

// propertyBagEpilogue creates any concluding statements required to handle our property bag after assignments are
// complete.
//   o If the destination has a property bag, we need to store our current property bag there
//   o Otherwise we do nothing
func (fn *PropertyAssignmentFunction) propertyBagEpilogue(
	destination string) []dst.Stmt {
	if dstBag := fn.findPropertyBag(fn.destinationType()); dstBag != nil {
		setBag := astbuilder.SimpleAssignment(
			astbuilder.Selector(dst.NewIdent(destination), string(dstBag.PropertyName())),
			token.ASSIGN,
			dst.NewIdent(fn.propertyBagName))
		setBag.Decs.Before = dst.EmptyLine
		astbuilder.AddComment(&setBag.Decorations().Start, "// Update the property bag")

		return astbuilder.Statements(setBag)
	}

	return nil
}

// generateAssignments generates a sequence of statements to copy information between the two types
func (fn *PropertyAssignmentFunction) generateAssignments(
	source dst.Expr,
	destination dst.Expr,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	var result []dst.Stmt

	// Find all the properties for which we have a conversion
	var properties []string
	for p := range fn.conversions {
		properties = append(properties, p)
	}

	// Sort the properties into alphabetical order to ensure deterministic generation
	sort.Slice(properties, func(i, j int) bool {
		return properties[i] < properties[j]
	})

	// Accumulate all the statements required for conversions, in alphabetical order
	for _, prop := range properties {
		conversion := fn.conversions[prop]
		block := conversion(source, destination, generationContext)
		if len(block) > 0 {
			firstStatement := block[0]
			firstStatement.Decorations().Before = dst.EmptyLine
			firstStatement.Decorations().Start.Prepend("// " + prop)
			result = append(result, block...)
		}
	}

	return result
}

// createConversions iterates through the properties on our receiver type, matching them up with
// our other type and generating conversions where possible
func (fn *PropertyAssignmentFunction) createConversions(receiver astmodel.TypeDefinition) error {
	sourceType := fn.sourceType()
	destinationType := fn.destinationType()

	sourceEndpoints := conversions.NewReadableConversionEndpointSet()
	sourceEndpoints.CreatePropertyEndpoints(sourceType, fn.knownLocals)
	sourceEndpoints.CreateValueFunctionEndpoints(sourceType, fn.knownLocals)

	destinationEndpoints := conversions.NewWritableConversionEndpointSet()
	destinationEndpoints.CreatePropertyEndpoints(destinationType, fn.knownLocals)

	// Flag receiver and parameter names as used
	fn.knownLocals.Add(fn.receiverName)
	fn.knownLocals.Add(fn.parameterName)

	for destinationName, destinationEndpoint := range destinationEndpoints {
		sourceEndpoint, ok := sourceEndpoints[destinationName]

		if !ok {
			// TODO: Handle property renames
			continue
		}

		// Generate a conversion from one endpoint to another
		conv, err := fn.createConversion(sourceEndpoint, destinationEndpoint)
		if err != nil {
			// An error was returned, we abort creating conversions for this object
			return errors.Wrapf(
				err,
				"creating conversion to %s by %s",
				destinationEndpoint,
				sourceEndpoint)
		} else if conv != nil {
			// A conversion was created, keep it for later
			fn.conversions[destinationName] = conv
		}
	}

	return nil
}

// createPropertyConversion tries to create a conversion between the two provided endpoints, using all of the
// available conversion functions in priority order to do so. If no valid conversion can be created an error is returned.
func (fn *PropertyAssignmentFunction) createConversion(
	sourceEndpoint conversions.ReadableConversionEndpoint,
	destinationEndpoint conversions.WritableConversionEndpoint) (StoragePropertyConversion, error) {

	conversion, err := conversions.CreateTypeConversion(sourceEndpoint.Endpoint(), destinationEndpoint.Endpoint(), fn.conversionContext)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"trying to %s and %s",
			sourceEndpoint, destinationEndpoint)
	}

	return func(source dst.Expr, destination dst.Expr, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
		reader := sourceEndpoint.Read(source)
		writer := func(expr dst.Expr) []dst.Stmt {
			return destinationEndpoint.Write(destination, expr)
		}

		return conversion(reader, writer, generationContext)
	}, nil
}

// findPropertyBag looks for a property bag on the specified type and returns it if found, or nil otherwise
// We recognize the property bag by type, so that the name can vary to avoid collisions with other properties if needed.
func (fn *PropertyAssignmentFunction) findPropertyBag(instance astmodel.Type) *astmodel.PropertyDefinition {
	if container, ok := astmodel.AsPropertyContainer(instance); ok {
		for _, prop := range container.Properties() {
			if prop.PropertyType().Equals(astmodel.PropertyBagType) {
				return prop
			}
		}
	}

	return nil
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
