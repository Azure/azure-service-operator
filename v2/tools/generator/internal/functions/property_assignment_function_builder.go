/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
)

// PropertyAssignmentFunctionBuilder is a factory used to construct a PropertyAssignmentFunction.
type PropertyAssignmentFunctionBuilder struct {
	// receiverDefinition is the type on which this function will be hosted
	receiverDefinition astmodel.TypeDefinition
	// otherDefinition is the type we are converting to (or from). This will be a type which is "closer"
	// to the hub storage type, making this a building block of the final conversion.
	otherDefinition astmodel.TypeDefinition
	// conversions is a map of all property conversions we are going to use, keyed by name of the
	// receiver endpoint (which may be a property, function, or property bag item)
	conversions map[string]StoragePropertyConversion
	// direction indicates the kind of conversion we are generating
	direction conversions.Direction
	// readsFromPropertyBag keeps track of whether we will be reading property values from a property bag
	readsFromPropertyBag bool
	// writesToPropertyBag keeps track of whether we will be writing property values into a property bag
	writesToPropertyBag bool
	// augmentationInterface is the conversion augmentation interface associated with this conversion.
	// If this is nil, there is no augmented conversion associated with this conversion
	augmentationInterface astmodel.TypeName
	// assignmentSelectors is a list of functions that can be used to select a property to assign to
	assignmentSelectors []assignmentSelector
}

type assignmentSelector struct {
	sequence int                        // lower numbered sequence numbers are executed first
	selector PropertyAssignmentSelector // the callback to invoke
}

// PropertyAssignmentSelector is a function that selects pairs of endpoints to create property assignments
// sourceProperties is the set of readable property endpoints.
// destinationProperties is the set of writable property endpoints.
// assign is a callback function used to assign a value from the source to the destination
// returns an error if there was a problem selecting a property, nil otherwise
type PropertyAssignmentSelector func(
	sourceProperties conversions.ReadableConversionEndpointSet,
	destinationProperties conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	conversionContext *conversions.PropertyConversionContext,
) error

// NewPropertyAssignmentFunctionBuilder creates a new factory for construction of a PropertyAssignmentFunction.
// receiver is the type definition that will be the receiver for this function
// otherDefinition is the type definition to convert TO or FROM
// direction specifies whether we are converting TO or FROM the other definition
func NewPropertyAssignmentFunctionBuilder(
	receiver astmodel.TypeDefinition,
	otherDefinition astmodel.TypeDefinition,
	direction conversions.Direction,
) *PropertyAssignmentFunctionBuilder {
	result := &PropertyAssignmentFunctionBuilder{
		receiverDefinition: receiver,
		otherDefinition:    otherDefinition,
		direction:          direction,
		conversions:        make(map[string]StoragePropertyConversion),
	}

	result.assignmentSelectors = []assignmentSelector{
		{0, result.selectIdenticallyNamedProperties},
		{1, result.selectRenamedProperties},
		{2, result.selectPropertiesWithIdenticalPaths},
		// High sequence numbers to ensure these are executed last
		{100, result.readPropertiesFromPropertyBag},
		{100, result.writePropertiesToPropertyBag},
	}

	return result
}

// UseAugmentationInterface returns the property assignment function with a conversion augmentation interface set
func (builder *PropertyAssignmentFunctionBuilder) UseAugmentationInterface(augmentation astmodel.TypeName) {
	builder.augmentationInterface = augmentation
}

// AddAssignmentSelector adds a new assignment selector to the list of assignment selectors.
// Assignment selectors are executed in the order they are added.
func (builder *PropertyAssignmentFunctionBuilder) AddAssignmentSelector(selector PropertyAssignmentSelector) {
	as := assignmentSelector{
		sequence: len(builder.assignmentSelectors) + 1, // ensure we execute after most existing selectors
		selector: selector,
	}

	builder.assignmentSelectors = append(builder.assignmentSelectors, as)
	slices.SortFunc(
		builder.assignmentSelectors,
		func(i assignmentSelector, j assignmentSelector) int {
			if i.sequence < j.sequence {
				return -1
			} else if i.sequence > j.sequence {
				return 1
			} else {
				return 0
			}
		})
}

// AddSuffixMatchingAssignmentSelector adds a new assignment selector that will match a property with the specified
// sourceSuffix to a property with the specified destinationSuffix.
func (builder *PropertyAssignmentFunctionBuilder) AddSuffixMatchingAssignmentSelector(
	sourceSuffix string,
	destinationSuffix string,
) {
	builder.AddAssignmentSelector(
		builder.createSuffixMatchingAssignmentSelector(sourceSuffix, destinationSuffix))
}

func (builder *PropertyAssignmentFunctionBuilder) Build(
	conversionContext *conversions.PropertyConversionContext,
) (*PropertyAssignmentFunction, error) {
	idFactory := conversionContext.IDFactory()
	knownLocals := astmodel.NewKnownLocalsSet(idFactory)

	// Create the function name
	fnName := conversions.NameOfPropertyAssignmentFunction(
		conversionContext.FunctionBaseName(),
		builder.otherDefinition.Name(),
		builder.direction,
		idFactory)

	// Select names for receiver and parameter
	receiverName := idFactory.CreateReceiver(builder.receiverDefinition.Name().Name())
	parameterName := builder.direction.SelectString("source", "destination")

	// If the two names collide, use a different convention for our parameter name
	if receiverName == parameterName {
		parameterName = builder.direction.SelectString("origin", "target")
	}

	// Flag receiver and parameter names as used
	knownLocals.Add(receiverName, parameterName)

	// Create Endpoints for property conversion
	sourceEndpoints := builder.createReadingEndpoints()
	destinationEndpoints := builder.createWritingEndpoints()

	// Always assign a name for the property bag (see createPropertyBagPrologue to understand why)
	propertyBagName := knownLocals.CreateLocal("propertyBag", "", "Local", "Temp")

	// Package references
	compatPkg := astmodel.MakeCompatPackageReference(builder.receiverDefinition.Name().InternalPackageReference())
	packageReferences := astmodel.NewPackageReferenceSet(
		astmodel.ErisReference,
		astmodel.GenRuntimeReference,
		builder.otherDefinition.Name().PackageReference(),
		compatPkg)

	cc := conversionContext.WithDirection(builder.direction).
		WithPropertyBag(propertyBagName).
		WithPackageReferenceSet(packageReferences)

	// Create conversions
	propertyConversions := make(map[string]StoragePropertyConversion, len(sourceEndpoints))
	err := builder.createConversions(sourceEndpoints, destinationEndpoints, cc, propertyConversions)
	if err != nil {
		parameterType := astmodel.DebugDescription(
			builder.otherDefinition.Name(),
			builder.receiverDefinition.Name().InternalPackageReference(),
		)

		return nil, errors.Wrapf(err, "creating '%s(%s)'", fnName, parameterType)
	}

	result := &PropertyAssignmentFunction{
		name:                   fnName,
		receiverDefinition:     builder.receiverDefinition,
		otherDefinition:        builder.otherDefinition,
		conversions:            propertyConversions,
		idFactory:              idFactory,
		direction:              builder.direction,
		conversionContext:      cc,
		receiverName:           receiverName,
		parameterName:          parameterName,
		knownLocals:            knownLocals,
		packageReferences:      packageReferences,
		augmentationInterface:  builder.augmentationInterface,
		sourcePropertyBag:      builder.findPropertyBagProperty(builder.sourceType()),
		destinationPropertyBag: builder.findPropertyBagProperty(builder.destinationType()),
		readsFromPropertyBag:   builder.readsFromPropertyBag,
		writesToPropertyBag:    builder.writesToPropertyBag,
	}

	return result, nil
}

// createReadingEndpoints creates a ReadableConversionEndpointSet containing all the readable endpoints we need for this
// conversion. If the source has a property bag, we create additional endpoints to match any surplus properties present
// on the DESTINATION type, so we can populate those from the property bag. Returns true if we create any endpoints to
// read from a property bag, false otherwise.
func (builder *PropertyAssignmentFunctionBuilder) createReadingEndpoints() conversions.ReadableConversionEndpointSet {
	sourceEndpoints := conversions.NewReadableConversionEndpointSet()
	sourceEndpoints.CreatePropertyEndpoints(builder.sourceType())
	sourceEndpoints.CreateValueFunctionEndpoints(builder.sourceType())

	return sourceEndpoints
}

// createWritingEndpoints creates a WritableConversionEndpointSet containing all the writable endpoints we need for this
// conversion. If the destination has a property bag, we create additional endpoints to match any surplus properties
// present on the SOURCE type, so we can stash those in the property bag for later use. Returns true if we create any
// endpoints to write into a property bag, false otherwise.
func (builder *PropertyAssignmentFunctionBuilder) createWritingEndpoints() conversions.WritableConversionEndpointSet {
	destinationEndpoints := conversions.NewWritableConversionEndpointSet()
	destinationEndpoints.CreatePropertyEndpoints(builder.destinationType())

	return destinationEndpoints
}

// createConversions iterates through the properties on our receiver type, matching them up with
// our other type and generating conversions where possible.
// sourceEndpoints is a set of endpoints that can be read from.
// destinationEndpoints is a set of endpoints that can be written to.
// conversionContext is the context for the conversion.
// conversions is a map of property names to conversions that will be populated by this function.
func (builder *PropertyAssignmentFunctionBuilder) createConversions(
	sourceEndpoints conversions.ReadableConversionEndpointSet,
	destinationEndpoints conversions.WritableConversionEndpointSet,
	conversionContext *conversions.PropertyConversionContext,
	propertyConversions map[string]StoragePropertyConversion,
) error {
	usedSources := set.Make[string]()
	usedDestinations := set.Make[string]()

	// Assign generates a conversion between a pair of endpoints
	assign := func(
		sourceEndpoint *conversions.ReadableConversionEndpoint,
		destinationEndpoint *conversions.WritableConversionEndpoint,
	) error {
		// Generate a conversion from one endpoint to another
		conv, err := builder.createConversion(sourceEndpoint, destinationEndpoint, conversionContext)
		if err != nil {
			// An error was returned, we abort creating conversions for this object
			return errors.Wrapf(
				err,
				"creating conversion to %s by %s",
				destinationEndpoint,
				sourceEndpoint)
		}

		if conv != nil {
			// A conversion was created, keep it for later, and make sure we don't reuse these endpoints
			propertyConversions[destinationEndpoint.Name()] = conv
			usedSources.Add(sourceEndpoint.Name())
			usedDestinations.Add(destinationEndpoint.Name())
		}

		return nil
	}

	for _, s := range builder.assignmentSelectors {
		err := s.selector(sourceEndpoints, destinationEndpoints, assign, conversionContext)
		if err != nil {
			// Don't need to wrap this error, it's already got context
			return err
		}

		// Remove used source endpoints
		for reader := range usedSources {
			sourceEndpoints.Delete(reader)
		}

		// Remove used destination endpoints
		for writer := range usedDestinations {
			destinationEndpoints.Delete(writer)
		}
	}

	return nil
}

// findPropertyBagProperty looks for a property bag on the specified type and returns it if found, or nil otherwise
// We recognize the property bag by type, so that the name can vary to avoid collisions with other properties if needed.
func (builder *PropertyAssignmentFunctionBuilder) findPropertyBagProperty(instance astmodel.Type) *astmodel.PropertyDefinition {
	if container, ok := astmodel.AsPropertyContainer(instance); ok {
		for _, prop := range container.Properties().Copy() {
			if astmodel.TypeEquals(prop.PropertyType(), astmodel.PropertyBagType) {
				return prop
			}
		}
	}

	return nil
}

// createPropertyConversion tries to create a conversion between the two provided endpoints, using all the available
// conversion functions in sequence order. If no valid conversion can be created an error is returned.
func (builder *PropertyAssignmentFunctionBuilder) createConversion(
	sourceEndpoint *conversions.ReadableConversionEndpoint,
	destinationEndpoint *conversions.WritableConversionEndpoint,
	conversionContext *conversions.PropertyConversionContext,
) (StoragePropertyConversion, error) {
	conversion, err := conversions.CreateTypeConversion(
		sourceEndpoint.Endpoint(),
		destinationEndpoint.Endpoint(),
		conversionContext)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"trying to %s and %s",
			sourceEndpoint, destinationEndpoint)
	}

	return func(
		source dst.Expr,
		destination dst.Expr,
		knownLocals *astmodel.KnownLocalsSet,
		generationContext *astmodel.CodeGenerationContext,
	) ([]dst.Stmt, error) {
		reader := sourceEndpoint.Read(source)
		writer := func(expr dst.Expr) []dst.Stmt {
			return destinationEndpoint.Write(destination, expr)
		}

		stmts, err := conversion(reader, writer, knownLocals, generationContext)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"converting %s to %s",
				sourceEndpoint, destinationEndpoint)
		}

		return stmts, nil
	}, nil
}

// sourceType returns the type we are reading information from
// When converting FROM, otherDefinition.Type() is our source
// When converting TO, receiverDefinition.Type() is our source
// Our inverse is destinationType()
func (builder *PropertyAssignmentFunctionBuilder) sourceType() astmodel.Type {
	return builder.direction.SelectType(builder.otherDefinition.Type(), builder.receiverDefinition.Type())
}

// destinationType returns the type we are writing information from
// When converting FROM, receiverDefinition.Type() is our source
// When converting TO, otherDefinition.Type() is our source
// Our inverse is sourceType()
func (builder *PropertyAssignmentFunctionBuilder) destinationType() astmodel.Type {
	return builder.direction.SelectType(builder.receiverDefinition.Type(), builder.otherDefinition.Type())
}

// selectIdenticallyNamedProperties matches up properties with identical names for conversion
// sourceProperties is a set of endpoints that can be read from.
// destinationProperties is a set of endpoints that can be written to.
// assign is a function that will be called for each matching property, with the source and destination endpoints
// for that property.
// Returns an error if any of the assignments fail.
func (*PropertyAssignmentFunctionBuilder) selectIdenticallyNamedProperties(
	sourceProperties conversions.ReadableConversionEndpointSet,
	destinationProperties conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	ctx *conversions.PropertyConversionContext,
) error {
	for destinationName, destinationEndpoint := range destinationProperties {
		if sourceEndpoint, ok := sourceProperties[destinationName]; ok {
			// We have properties with the same name, we also need them to have the same path.
			// This effectively means they were flattened from the same original property.
			// If they don't, we can't match them up.
			if sourceEndpoint.Endpoint().Path() != destinationEndpoint.Endpoint().Path() {
				continue
			}

			err := assign(sourceEndpoint, destinationEndpoint)
			if err != nil {
				return errors.Wrapf(err, "assigning %s", destinationName)
			}
		}
	}

	return nil
}

// selectPropertiesWithIdenticalPaths matches up properties with paths for conversion.
// This serves to match up properties that were flattened from the original location,
// even if they've ended up with different names.
// sourceProperties is a set of endpoints that can be read from.
// destinationProperties is a set of endpoints that can be written to.
// assign is a function that will be called for each matching property, with the source and destination endpoints
// for that property.
// Returns an error if any of the assignments fail.
func (builder *PropertyAssignmentFunctionBuilder) selectPropertiesWithIdenticalPaths(
	sourceProperties conversions.ReadableConversionEndpointSet,
	destinationProperties conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	ctx *conversions.PropertyConversionContext,
) error {
	// Create a map of source properties, by path
	sourceByPath := make(map[string][]*conversions.ReadableConversionEndpoint)
	for _, source := range sourceProperties {
		path := source.Endpoint().Path()
		if path == "" {
			// No path (so wasn't flattened); skip
			continue
		}

		sourceByPath[path] = append(sourceByPath[path], source)
	}

	for destinationName, destinationEndpoint := range destinationProperties {
		path := destinationEndpoint.Endpoint().Path()
		if path == "" {
			// No path (so wasn't flattened); skip
			continue
		}

		sourceEndpoints, ok := sourceByPath[path]
		if !ok {
			// No match
			continue
		}

		// Look for a unique endpoint that's compatible with the destination
		var sourceEndpoint *conversions.ReadableConversionEndpoint
		for _, src := range sourceEndpoints {
			if !builder.typesCompatible(src.Endpoint().Type(), destinationEndpoint.Endpoint().Type(), ctx) {
				// Types aren't compatible, skip
				continue
			}

			if sourceEndpoint != nil {
				// We've found multiple candidates - we can't handle this
				return errors.Errorf(
					"multiple source properties with path %s are compatible with destination %s, no way to select",
					path,
					destinationName)
			}

			sourceEndpoint = src
		}

		// If we found a compatible endpoint, use it
		if sourceEndpoint != nil {
			err := assign(sourceEndpoint, destinationEndpoint)
			if err != nil {
				return errors.Wrapf(err, "assigning %s", destinationName)
			}
		}
	}

	return nil
}

// selectRenamedProperties matches up properties that were renamed in the later version
// sourceProperties is a set of endpoints that can be read from.
// destinationProperties is a set of endpoints that can be written to.
// assign is a function that will be called for each matching property, with the source and destination endpoints
// for that property.
// Returns an error if any of the assignments fail.
func (builder *PropertyAssignmentFunctionBuilder) selectRenamedProperties(
	sourceProperties conversions.ReadableConversionEndpointSet,
	destinationProperties conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	conversionContext *conversions.PropertyConversionContext,
) error {
	// Create a map of active renames, from the source name to the destination name
	// Depending on the direction of conversion, we need populate this differently
	renames := make(map[string]string)
	if conversionContext.HasDirection(conversions.ConvertTo) {
		// Source type is receiver, use that to look up renames
		for source := range sourceProperties {
			if name, renamed := conversionContext.PropertyRename(
				builder.receiverDefinition.Name(),
				astmodel.PropertyName(source)); renamed {
				renames[source] = name
			}
		}
	} else {
		// Destination type is receiver, use that to look up renames
		for destination := range destinationProperties {
			if name, renamed := conversionContext.PropertyRename(
				builder.receiverDefinition.Name(),
				astmodel.PropertyName(destination)); renamed {
				renames[name] = destination
			}
		}
	}

	// Match up any renamed properties
	for sourceName, sourceEndpoint := range sourceProperties {
		destinationName, renamed := renames[sourceName]
		if !renamed {
			continue
		}

		if destinationEndpoint, ok := destinationProperties[destinationName]; ok {
			err := assign(sourceEndpoint, destinationEndpoint)
			if err != nil {
				return errors.Wrapf(err, "assigning %s", destinationName)
			}
		}
	}

	return nil
}

// readPropertiesFromPropertyBag populates destination properties that don't have a matching source property by reading
// the values from a property bag on the source object.
// sourceEndpoints is a set of endpoints that can be read from.
// destinationEndpoints is a set of endpoints that can be written to.
// assign is a function that will be called for each matching property, with the source and destination endpoints
// for that assignment.
func (builder *PropertyAssignmentFunctionBuilder) readPropertiesFromPropertyBag(
	sourceEndpoints conversions.ReadableConversionEndpointSet,
	destinationEndpoints conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	conversionContext *conversions.PropertyConversionContext,
) error {
	prop := builder.findPropertyBagProperty(builder.sourceType())
	if prop == nil {
		// No property bag on our source type, nothing to read from
		return nil
	}

	// for each destination property that doesn't already have a source endpoint, synthesize an endpoint that reads
	// from the property bag
	for destinationName, destinationEndpoint := range destinationEndpoints {
		if _, ok := sourceEndpoints[destinationName]; ok {
			// already have a source endpoint for this property
			continue
		}

		// Create a new endpoint that reads from the property bag
		typeToRead := builder.findTypeForBag(destinationEndpoint.Endpoint().Type(), conversionContext)
		sourceEndpoint := conversions.NewReadableConversionEndpointReadingPropertyBagMember(destinationName, typeToRead)
		err := assign(sourceEndpoint, destinationEndpoint)
		if err != nil {
			return errors.Wrapf(err, "assigning %s from property bag", destinationName)
		}

		builder.readsFromPropertyBag = true
	}

	return nil
}

// writePropertiesToPropertyBag stores source properties that don't have a matching destination property by writing
// the values to a property bag on the destination object.
// sourceEndpoints is a set of endpoints that can be read from.
// destinationEndpoints is a set of endpoints that can be written to.
// assign is a function that will be called for each matching property, with the source and destination endpoints
// for that assignment.
func (builder *PropertyAssignmentFunctionBuilder) writePropertiesToPropertyBag(
	sourceEndpoints conversions.ReadableConversionEndpointSet,
	destinationEndpoints conversions.WritableConversionEndpointSet,
	assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
	conversionContext *conversions.PropertyConversionContext,
) error {
	prop := builder.findPropertyBagProperty(builder.destinationType())
	if prop == nil {
		// No property bag on our source type, nothing to read from
		return nil
	}

	// for each source property that doesn't already have a destination endpoint, synthesize an endpoint that writes
	// to the property bag
	for sourceName, sourceEndpoint := range sourceEndpoints {
		if _, ok := destinationEndpoints[sourceName]; ok {
			// already have a destination endpoint for this property, nothing to do
			continue
		}

		// Create a new endpoint that writes to the property bag
		typeToWrite := builder.findTypeForBag(sourceEndpoint.Endpoint().Type(), conversionContext)
		destinationEndpoint := conversions.NewWritableConversionEndpointWritingPropertyBagMember(sourceName, typeToWrite)
		err := assign(sourceEndpoint, destinationEndpoint)
		if err != nil {
			return errors.Wrapf(err, "assigning %s to property bag", sourceName)
		}

		builder.writesToPropertyBag = true
	}

	return nil
}

// createSuffixMatchingAssignmentSelector creates an assignment selector that matches source properties with the
// given suffix to destination properties with the given suffix.
func (builder *PropertyAssignmentFunctionBuilder) createSuffixMatchingAssignmentSelector(
	sourceSuffix string,
	destinationSuffix string,
) PropertyAssignmentSelector {
	return func(sourceProperties conversions.ReadableConversionEndpointSet,
		destinationProperties conversions.WritableConversionEndpointSet,
		assign func(reader *conversions.ReadableConversionEndpoint, writer *conversions.WritableConversionEndpoint) error,
		_ *conversions.PropertyConversionContext,
	) error {
		for destinationName, destinationEndpoint := range destinationProperties {
			if !strings.HasSuffix(destinationName, destinationSuffix) {
				continue
			}

			sourceName := strings.TrimSuffix(destinationName, destinationSuffix) + sourceSuffix
			if sourceEndpoint, ok := sourceProperties[sourceName]; ok {
				err := assign(sourceEndpoint, destinationEndpoint)
				if err != nil {
					return errors.Wrapf(err, "assigning %s", destinationName)
				}
			}
		}

		return nil
	}
}

func (builder *PropertyAssignmentFunctionBuilder) findTypeForBag(
	t astmodel.Type,
	conversionContext *conversions.PropertyConversionContext,
) astmodel.Type {
	// If t is optional, look up the underlying type and then wrap
	if opt, ok := astmodel.AsOptionalType(t); ok {
		elem := builder.findTypeForBag(opt.Element(), conversionContext)
		return astmodel.NewOptionalType(elem)
	}

	// If t is an array, look up the underlying type and then wrap
	if arr, ok := astmodel.AsArrayType(t); ok {
		elem := builder.findTypeForBag(arr.Element(), conversionContext)
		return astmodel.NewArrayType(elem)
	}

	// If t is a map, look-up the underlying type of the value and then wrap
	if m, ok := astmodel.AsMapType(t); ok {
		value := builder.findTypeForBag(m.ValueType(), conversionContext)
		return astmodel.NewMapType(m.KeyType(), value)
	}

	// If t is a TypeName, check for the existence of a compatibility type in a subpackge under the receiver
	if tn, ok := astmodel.AsInternalTypeName(t); ok {
		compatPkg := astmodel.MakeCompatPackageReference(
			builder.receiverDefinition.Name().InternalPackageReference())
		compatType := tn.WithPackageReference(compatPkg)
		if conversionContext.Types().Contains(compatType) {
			// Compatibility type exists - use that
			return compatType
		}
	}

	return t
}

// Return true if the types are generally compatible; this is used to disambiguate between properties when
// selecting property pairs for assignment, so it doesn't need to be perfect.
func (*PropertyAssignmentFunctionBuilder) typesCompatible(
	left astmodel.Type,
	right astmodel.Type,
	_ *conversions.PropertyConversionContext,
) bool {
	// if left resolves to an external type name, we're compatible if right is also an external type name
	if _, ok := astmodel.AsExternalTypeName(left); ok {
		_, rightIsExternal := astmodel.AsExternalTypeName(right)
		return rightIsExternal
	}

	// if left resolves to an internal type name, we're compatible if right is also an internal type name
	if _, ok := astmodel.AsInternalTypeName(left); ok {
		_, rightIsInternal := astmodel.AsInternalTypeName(right)
		return rightIsInternal
	}

	// if left resolves to a primitive type, we're compatible if right is also a primitive type
	if _, ok := astmodel.AsPrimitiveType(left); ok {
		_, rightIsPrimitive := astmodel.AsPrimitiveType(right)
		return rightIsPrimitive
	}

	// TODO: Arrays, Maps, and look-through of InternalTypeName for aliases

	return false
}
