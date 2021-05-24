/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"go/token"
	"sort"
	"strings"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/dave/dst"
	"github.com/pkg/errors"
)

// StorageConversionFunction represents a function that performs single step conversions for storage versions
type StorageConversionFunction struct {
	// name of this conversion function
	name string
	// otherDefinition is the type we are converting to (or from). This will be a type which is "closer"
	// to the hub storage type, making this a building block of the final conversion.
	otherDefinition astmodel.TypeDefinition
	// conversions is a map of all property conversions we are going to use, keyed by name of the
	// receiver property
	conversions map[string]StoragePropertyConversion
	// idFactory is a reference to an identifier factory used for creating Go identifiers
	idFactory astmodel.IdentifierFactory
	// conversionDirection indicates the kind of conversion we are generating
	conversionDirection StorageConversionDirection
	// knownLocals is a cached set of local identifiers that have already been used, to avoid conflicts
	knownLocals *astmodel.KnownLocalsSet
	// conversionContext is additional information about the context in which this conversion was made
	conversionContext *PropertyConversionContext
}

// StoragePropertyConversion represents a function that generates the correct AST to convert a single property value
// Different functions will be used, depending on the types of the properties to be converted.
// source is an expression for the source value that will be read.
// destination is an expression the target value that will be written.
type StoragePropertyConversion func(
	source dst.Expr, destination dst.Expr, generationContext *astmodel.CodeGenerationContext) []dst.Stmt

// StorageConversionDirection specifies the direction of conversion we're implementing with this function
type StorageConversionDirection int

const (
	// ConvertFrom indicates the conversion is from the passed other instance, populating the receiver
	ConvertFrom = StorageConversionDirection(1)
	// ConvertTo indicates the conversion is to the passed other type, populating other
	ConvertTo = StorageConversionDirection(2)
)

// Ensure that StorageConversionFunction implements Function
var _ astmodel.Function = &StorageConversionFunction{}

// NewStorageConversionFromFunction creates a new StorageConversionFunction to convert from the specified source
func NewStorageConversionFromFunction(
	receiver astmodel.TypeDefinition,
	otherDefinition astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	conversionContext *PropertyConversionContext,
) (*StorageConversionFunction, error) {
	result := &StorageConversionFunction{
		otherDefinition:     otherDefinition,
		idFactory:           idFactory,
		conversionDirection: ConvertFrom,
		conversions:         make(map[string]StoragePropertyConversion),
		knownLocals:         astmodel.NewKnownLocalsSet(idFactory),
	}

	version := idFactory.CreateIdentifier(otherDefinition.Name().PackageReference.PackageName(), astmodel.Exported)
	result.name = "ConvertFrom" + version
	result.conversionContext = conversionContext.WithFunctionName(result.name).WithKnownLocals(result.knownLocals)

	err := result.createConversions(receiver, conversionContext.Types())
	if err != nil {
		return nil, errors.Wrapf(err, "creating '%s()'", result.name)
	}

	return result, nil
}

// NewStorageConversionToFunction creates a new StorageConversionFunction to convert to the specified destination
func NewStorageConversionToFunction(
	receiver astmodel.TypeDefinition,
	otherDefinition astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	conversionContext *PropertyConversionContext,
) (*StorageConversionFunction, error) {
	result := &StorageConversionFunction{
		otherDefinition:     otherDefinition,
		idFactory:           idFactory,
		conversionDirection: ConvertTo,
		conversions:         make(map[string]StoragePropertyConversion),
		knownLocals:         astmodel.NewKnownLocalsSet(idFactory),
	}

	version := idFactory.CreateIdentifier(otherDefinition.Name().PackageReference.PackageName(), astmodel.Exported)
	result.name = "ConvertTo" + version
	result.conversionContext = conversionContext.WithFunctionName(result.name).WithKnownLocals(result.knownLocals)

	err := result.createConversions(receiver, conversionContext.Types())
	if err != nil {
		return nil, errors.Wrapf(err, "creating '%s()'", result.name)
	}

	return result, nil
}

// Name returns the name of this function
func (fn *StorageConversionFunction) Name() string {
	return fn.name
}

// RequiredPackageReferences returns the set of package references required by this function
		astmodel.GitHubErrorsReference,
func (fn *StorageConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet(
		astmodel.GitHubErrorsReference,
		fn.otherDefinition.Name().PackageReference)

	return result
}

// References returns the set of types referenced by this function
func (fn *StorageConversionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(fn.otherDefinition.Name())
}

// Equals checks to see if the supplied function is the same as this one
func (fn *StorageConversionFunction) Equals(f astmodel.Function) bool {
	if other, ok := f.(*StorageConversionFunction); ok {
		if fn.name != other.name {
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

// AsFunc renders this function as an AST for serialization to a Go source file
func (fn *StorageConversionFunction) AsFunc(generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {

	var parameterName string
	var description string
	switch fn.conversionDirection {
	case ConvertFrom:
		parameterName = "source"
		description = fmt.Sprintf("populates our %s from the provided source %s", receiver.Name(), fn.otherDefinition.Name().Name())
	case ConvertTo:
		parameterName = "destination"
		description = fmt.Sprintf("populates the provided destination %s from our %s", fn.otherDefinition.Name().Name(), receiver.Name())
	default:
		panic(fmt.Sprintf("unexpected conversion direction %q", fn.conversionDirection))
	}

	// Create a sensible name for our receiver
	receiverName := fn.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)

	// We always use a pointer receiver so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverType,
		Name:          fn.Name(),
		Body:          fn.generateBody(receiverName, parameterName, generationContext),
	}

	parameterPackage := generationContext.MustGetImportedPackageName(fn.otherDefinition.Name().PackageReference)

	funcDetails.AddParameter(
		parameterName,
		&dst.StarExpr{
			X: &dst.SelectorExpr{
				X:   dst.NewIdent(parameterPackage),
				Sel: dst.NewIdent(fn.otherDefinition.Name().Name()),
			},
		})

	funcDetails.AddReturns("error")
	funcDetails.AddComments(description)

	return funcDetails.DefineFunc()
}

// generateBody returns all of the statements required for the conversion function
// receiver is an expression for access our receiver type, used to qualify field access
// parameter is an expression for access to our parameter passed to the function, also used for field access
// generationContext is our code generation context, passed to allow resolving of identifiers in other packages
func (fn *StorageConversionFunction) generateBody(
	receiver string,
	parameter string,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	switch fn.conversionDirection {
	case ConvertFrom:
		return fn.generateDirectConversionFrom(receiver, parameter, generationContext)
	case ConvertTo:
		return fn.generateDirectConversionTo(receiver, parameter, generationContext)
	default:
		panic(fmt.Sprintf("unexpected conversion direction %q", fn.conversionDirection))
	}
}

// generateDirectConversionFrom returns the method body required to directly copy information from
// the parameter instance onto our receiver
func (fn *StorageConversionFunction) generateDirectConversionFrom(
	receiver string,
	parameter string,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	result := fn.generateAssignments(dst.NewIdent(parameter), dst.NewIdent(receiver), generationContext)
	result = append(result, astbuilder.ReturnNoError())
	return result
}

// generateDirectConversionTo returns the method body required to directly copy information from
// our receiver onto the parameter instance
func (fn *StorageConversionFunction) generateDirectConversionTo(
	receiver string,
	parameter string,
	generationContext *astmodel.CodeGenerationContext,
) []dst.Stmt {
	result := fn.generateAssignments(dst.NewIdent(receiver), dst.NewIdent(parameter), generationContext)
	result = append(result, astbuilder.ReturnNoError())
	return result
}

// generateAssignments generates a sequence of statements to copy information between the two types
func (fn *StorageConversionFunction) generateAssignments(
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
func (fn *StorageConversionFunction) createConversions(receiver astmodel.TypeDefinition, types astmodel.Types) error {

	receiverContainer, ok := fn.asPropertyContainer(receiver.Type())
	if !ok {
		var typeDescription strings.Builder
		receiver.Type().WriteDebugDescription(&typeDescription, types)

		return errors.Errorf(
			"expected receiver TypeDefinition %q to be either resource or object type, but found %q",
			receiver.Name().String(),
			typeDescription.String())
	}

	otherContainer, ok := fn.asPropertyContainer(fn.otherDefinition.Type())
	if !ok {
		var typeDescription strings.Builder
		fn.otherDefinition.Type().WriteDebugDescription(&typeDescription, types)

		return errors.Errorf(
			"expected other TypeDefinition %q to be either resource or object type, but found %q",
			fn.otherDefinition.Name().String(),
			typeDescription.String())
	}

	receiverProperties := fn.createPropertyMap(receiverContainer)
	otherProperties := fn.createPropertyMap(otherContainer)

	// Flag receiver name as used
	fn.knownLocals.Add(receiver.Name().Name())

	for _, receiverProperty := range receiverProperties {
		otherProperty, ok := otherProperties[receiverProperty.PropertyName()]
		//TODO: Handle renames
		if ok {
			var conv StoragePropertyConversion
			var err error
			switch fn.conversionDirection {
			case ConvertFrom:
				conv, err = fn.createPropertyConversion(otherProperty, receiverProperty)
			case ConvertTo:
				conv, err = fn.createPropertyConversion(receiverProperty, otherProperty)
			default:
				panic(fmt.Sprintf("unexpected conversion direction %q", fn.conversionDirection))
			}

			if err != nil {
				// An error was returned, we abort creating conversions for this object
				return errors.Wrapf(err, "creating conversion for property %q of %q", receiverProperty.PropertyName(), receiver.Name())
			} else if conv != nil {
				// A conversion was created, keep it for later
				fn.conversions[string(receiverProperty.PropertyName())] = conv
			}
		}
	}

	return nil
}

// asPropertyContainer converts a type into a property container
func (fn *StorageConversionFunction) asPropertyContainer(theType astmodel.Type) (astmodel.PropertyContainer, bool) {
	switch t := theType.(type) {
	case astmodel.PropertyContainer:
		return t, true
	case astmodel.MetaType:
		return fn.asPropertyContainer(t.Unwrap())
	default:
		return nil, false
	}
}

// createPropertyMap extracts the properties from a PropertyContainer and returns them as a map
func (fn *StorageConversionFunction) createPropertyMap(container astmodel.PropertyContainer) map[astmodel.PropertyName]*astmodel.PropertyDefinition {
	result := make(map[astmodel.PropertyName]*astmodel.PropertyDefinition)
	for _, p := range container.Properties() {
		result[p.PropertyName()] = p
	}

	return result
}

// createPropertyConversion tries to create a property conversion between the two provided properties, using all of the
// available conversion functions in priority order to do so. If no valid conversion could be created an error is returned.
func (fn *StorageConversionFunction) createPropertyConversion(
	sourceProperty *astmodel.PropertyDefinition,
	destinationProperty *astmodel.PropertyDefinition) (StoragePropertyConversion, error) {

	sourceEndpoint := NewStorageConversionEndpoint(
		sourceProperty.PropertyType(), string(sourceProperty.PropertyName()), fn.knownLocals)
	destinationEndpoint := NewStorageConversionEndpoint(
		destinationProperty.PropertyType(), string(destinationProperty.PropertyName()), fn.knownLocals)

	conversion, err := CreateTypeConversion(sourceEndpoint, destinationEndpoint, fn.conversionContext)

	if err != nil {
		return nil, errors.Wrapf(
			err,
			"trying to assign %q [%s] by converting from %q [%s]",
			destinationProperty.PropertyName(),
			destinationProperty.PropertyType(),
			sourceProperty.PropertyName(),
			sourceProperty.PropertyType())
	}

	return func(source dst.Expr, destination dst.Expr, generationContext *astmodel.CodeGenerationContext) []dst.Stmt {

		reader := astbuilder.Selector(source, string(sourceProperty.PropertyName()))
		writer := func(expr dst.Expr) []dst.Stmt {
			return []dst.Stmt{
				astbuilder.SimpleAssignment(
					astbuilder.Selector(destination, string(destinationProperty.PropertyName())),
					token.ASSIGN,
					expr),
			}
		}

		return conversion(reader, writer, generationContext)
	}, nil
}
