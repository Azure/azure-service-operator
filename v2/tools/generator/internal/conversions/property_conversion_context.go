/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const (
	AssignPropertiesMethodPrefix = "AssignProperties"
	InitializationMethodPrefix   = "Initialize"
)

// PropertyConversionContext captures additional supporting information that may be needed when a
// storage conversion factory creates a conversion
type PropertyConversionContext struct {
	// definitions is a map of all known type definitions, used to resolve TypeNames to actual definitions
	definitions astmodel.TypeDefinitionSet
	// functionBaseName is the base name of the function we're generating (used to generate the names of function calls)
	functionBaseName string
	// direction is the direction of the conversion we're generating
	direction Direction
	// propertyBagName is the name of the local variable used for a property bag, or "" if we don't have one
	propertyBagName string
	// Configuration containing additional metadata for generating conversions
	configuration *config.ObjectModelConfiguration
	// idFactory is used for generating method names
	idFactory astmodel.IdentifierFactory
	// conversionGraph optionally contains our package conversion graph
	conversionGraph *storage.ConversionGraph
	// additionalReferences is a reference to a shared set of additional package references needed by conversions
	additionalReferences *astmodel.PackageReferenceSet
}

// NewPropertyConversionContext creates a new instance of a PropertyConversionContext
func NewPropertyConversionContext(
	functionBaseName string,
	definitions astmodel.TypeDefinitionSet,
	idFactory astmodel.IdentifierFactory,
) *PropertyConversionContext {
	return &PropertyConversionContext{
		functionBaseName:     functionBaseName,
		definitions:          definitions,
		idFactory:            idFactory,
		propertyBagName:      "",
		additionalReferences: astmodel.NewPackageReferenceSet(),
	}
}

// FunctionBaseName returns the base name of the function we're generating
func (c *PropertyConversionContext) FunctionBaseName() string {
	return c.functionBaseName
}

// Types returns the set of definitions available in this context
func (c *PropertyConversionContext) Types() astmodel.TypeDefinitionSet {
	return c.definitions
}

// IDFactory returns a reference to our identifier factory
func (c *PropertyConversionContext) IDFactory() astmodel.IdentifierFactory {
	return c.idFactory
}

// WithConfiguration returns a new context with the specified configuration included
func (c *PropertyConversionContext) WithConfiguration(configuration *config.ObjectModelConfiguration) *PropertyConversionContext {
	result := c.clone()
	result.configuration = configuration
	return result
}

// WithConversionGraph returns a new context with the specified conversion graph included
func (c *PropertyConversionContext) WithConversionGraph(conversionGraph *storage.ConversionGraph) *PropertyConversionContext {
	result := c.clone()
	result.conversionGraph = conversionGraph
	return result
}

// WithDirection returns a new context with the specified direction
func (c *PropertyConversionContext) WithDirection(dir Direction) *PropertyConversionContext {
	result := c.clone()
	result.direction = dir
	return result
}

// WithPropertyBag returns a new context with the specified property bag name included
func (c *PropertyConversionContext) WithPropertyBag(name string) *PropertyConversionContext {
	result := c.clone()
	result.propertyBagName = name
	return result
}

// WithPackageReferenceSet returns a new context that collects new package references
func (c *PropertyConversionContext) WithPackageReferenceSet(set *astmodel.PackageReferenceSet) *PropertyConversionContext {
	result := c.clone()
	result.additionalReferences = set
	return result
}

// HasDirection returns true if this context has the specified direction.
func (c *PropertyConversionContext) HasDirection(dir Direction) bool {
	return c.direction == dir
}

// ResolveType resolves a type that might be a type name into both the name and the actual
// type it references, returning true iff it was a TypeName that could be resolved
func (c *PropertyConversionContext) ResolveType(t astmodel.Type) (astmodel.InternalTypeName, astmodel.Type, bool) {
	name, ok := astmodel.AsInternalTypeName(t)
	if !ok {
		return astmodel.InternalTypeName{}, nil, false
	}

	actualType, err := c.definitions.FullyResolve(name)
	if err != nil {
		return astmodel.InternalTypeName{}, nil, false
	}

	return name, actualType, true
}

// PropertyBagName returns the name to use for a local property bag variable
func (c *PropertyConversionContext) PropertyBagName() string {
	return c.propertyBagName
}

// TypeRename looks up a type-rename for the specified type, returning the new name and true if found, or empty string
// and false if not. If no configuration is available, acts as though there is no configuration for this rename,
// returning "" and false.
func (c *PropertyConversionContext) TypeRename(name astmodel.InternalTypeName) (string, bool) {
	if c.configuration == nil {
		return "", false
	}

	return c.configuration.TypeNameInNextVersion.Lookup(name)
}

// PropertyRename looks up a property-rename for the specified type, returning the new name and true if found, or empty
// string and false if not. If no configuration is available, acts as though there is no configuration for this rename,
// returning "" and false.
func (c *PropertyConversionContext) PropertyRename(
	container astmodel.InternalTypeName,
	property astmodel.PropertyName,
) (string, bool) {
	if c.configuration == nil {
		return "", false
	}

	return c.configuration.PropertyNameInNextVersion.Lookup(container, property)
}

// FindNextType returns the next type in the storage conversion graph, if any.
// If no conversion graph is available, returns an empty type name and no error.
func (c *PropertyConversionContext) FindNextType(name astmodel.InternalTypeName) (astmodel.InternalTypeName, error) {
	if c.conversionGraph == nil {
		return astmodel.InternalTypeName{}, nil
	}

	return c.conversionGraph.FindNextType(name, c.definitions)
}

// PathExists returns true if a path exists in the conversion graph starting from the specified type name and ending
// at the specified type name. If no conversion graph is available, returns false.
func (c *PropertyConversionContext) PathExists(
	start astmodel.InternalTypeName,
	finish astmodel.InternalTypeName,
) bool {
	if c.conversionGraph == nil {
		return false
	}

	_, found := c.conversionGraph.FindInPath(
		start,
		func(name astmodel.InternalTypeName) bool {
			return name == finish
		})

	return found
}

// FindPivotType returns the type name of the pivot type between the two specified types, if any.
func (c *PropertyConversionContext) FindPivotType(
	start astmodel.InternalTypeName,
	finish astmodel.InternalTypeName,
) (astmodel.InternalTypeName, bool) {
	// Walk the path from 'start' to find all the types visible, these are our candidate results
	candidates := astmodel.NewInternalTypeNameSet()
	c.conversionGraph.FindInPath(
		start,
		func(name astmodel.InternalTypeName) bool {
			candidates.Add(name)
			return false
		})

	// Walk the path from 'finish' to find the first type that's also visible from 'start'
	pivot, found := c.conversionGraph.FindInPath(
		finish,
		func(name astmodel.InternalTypeName) bool {
			return candidates.Contains(name)
		})
	return pivot, found
}

// AddPackageReference adds a new reference that's needed by the given conversion
func (c *PropertyConversionContext) AddPackageReference(ref astmodel.PackageReference) {
	c.additionalReferences.AddReference(ref)
}

// clone returns a new independent copy of this context
func (c *PropertyConversionContext) clone() *PropertyConversionContext {
	return &PropertyConversionContext{
		definitions:          c.definitions,
		functionBaseName:     c.functionBaseName,
		direction:            c.direction,
		propertyBagName:      c.propertyBagName,
		idFactory:            c.idFactory,
		configuration:        c.configuration,
		conversionGraph:      c.conversionGraph,
		additionalReferences: c.additionalReferences,
	}
}

// validateTypeRename is used to validate two types with different names are a properly renamed pair
func (c *PropertyConversionContext) validateTypeRename(
	sourceName astmodel.InternalTypeName,
	destinationName astmodel.InternalTypeName,
) error {
	// Work out which name represents the earlier package release
	// (needed in order to do the lookup as the type rename is configured on the last type *before* the rename.)
	var earlier astmodel.InternalTypeName
	var later astmodel.InternalTypeName
	if c.direction == ConvertTo {
		earlier = sourceName
		later = destinationName
	} else {
		earlier = destinationName
		later = sourceName
	}

	name, ok := c.TypeRename(earlier)
	if !ok {
		// No rename configured, but we can't proceed without one. Return an error - it'll be wrapped with property
		// details by CreateTypeConversion() so we only need the specific details here
		return eris.Errorf(
			"no configuration to rename %s to %s",
			earlier.Name(),
			later.Name())
	}

	if later.Name() != name {
		// Configured rename doesn't match what we found. Return an error - it'll be wrapped with property details
		// by CreateTypeConversion() so we only need the specific details here
		return eris.Errorf(
			"configuration includes rename of %s to %s, but found %s",
			earlier.Name(),
			name,
			later.Name())
	}

	return nil
}
