/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
)

const RepairSkippingPropertiesStageID = "repairSkippingProperties"

// RepairSkippingProperties repairs any properties that skip one or more versions of an object and are reintroduced.
// As described in issue #1776, such properties are a problem because they might create incompatibilities between
// different versions of ASO, or between different versions of a given resource.
//
// To repair these, we need to ensure that objects stored in property bags are always serialized with the same
// shape. For more details, see https://azure.github.io/azure-service-operator/design/adr-2023-09-skipping-properties/
//
// Repair works by scanning for properties that are dropped between versions of a resource. We keep track of all
// these properties, and if a specific property appears more than once (implying there are two or more sequence versions
// of that property) we know a repair is required.
//
// To illustrate, assume we have the following set of objects across versions of our mythical CRM service:
//
// v1: Person(GivenName, FamilyName, Address)
// v2: Person(KnownAs, FullName, FamilyName)
// v3: Person(KnownAs, FullName, FamilyName, Address)
//
// Scanning these types, we find:
//
// FamilyName: present in (v1, v2, v3); no issue
// FullName: present in (v2, v3); no issue
// KnownAs: also present in (v2, v3); no issue
// GivenName: present in (v1); no issue
//
// Address: present in (v1, v3); (skipping v2), repair required.
func RepairSkippingProperties() *Stage {
	stage := NewStage(
		RepairSkippingPropertiesStageID,
		"Repair property bag serialization for properties that skip resource or object versions",
		func(ctx context.Context, state *State) (*State, error) {
			var graph *storage.ConversionGraph
			if g, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo); err != nil {
				return nil, eris.Wrapf(err, "couldn't find conversion graph")
			} else {
				graph = g
			}

			repairer := newSkippingPropertyRepairer(state.Definitions(), graph)

			// Add resources and objects to the graph
			for _, def := range state.Definitions() {
				if t, ok := astmodel.AsPropertyContainer(def.Type()); ok {
					err := repairer.AddProperties(def.Name(), t.Properties().AsSlice()...)
					if err != nil {
						return nil, err
					}
				}
			}

			defs, err := repairer.RepairSkippedProperties()
			if err != nil {
				return nil, err
			}

			return state.WithOverlaidDefinitions(defs), err
		})

	// We have to have a conversion graph to detect skipping properties
	stage.RequiresPrerequisiteStages(CreateConversionGraphStageID)

	// We also require the Conversion Graph to be recreated afterwards,
	// but our stage dependencies can't currently capture that.

	return stage
}

type skippingPropertyRepairer struct {
	links              map[astmodel.PropertyReference]astmodel.PropertyReference // Individual links in chains of related properties
	observedProperties *astmodel.PropertyReferenceSet                            // Set of properties we've observed
	definitions        astmodel.TypeDefinitionSet                                // Set of all known type definitions
	conversionGraph    *storage.ConversionGraph                                  // Graph of conversions between types
	comparer           *structuralComparer                                       // Helper used to compare types for structural equality
}

// newSkippingPropertyRepairer creates a new graph for tracking chains of properties as they evolve through different
// versions of a resource or object.
// definitions is a set of all known types.
// conversionGraph contains every conversion/transition between versions.
func newSkippingPropertyRepairer(
	definitions astmodel.TypeDefinitionSet,
	conversionGraph *storage.ConversionGraph,
) *skippingPropertyRepairer {
	return &skippingPropertyRepairer{
		links:              make(map[astmodel.PropertyReference]astmodel.PropertyReference),
		observedProperties: astmodel.NewPropertyReferenceSet(),
		definitions:        definitions,
		conversionGraph:    conversionGraph,
		comparer:           newStructuralComparer(definitions),
	}
}

// AddProperties adds all the properties from the specified type to the graph.
func (repairer *skippingPropertyRepairer) AddProperties(
	name astmodel.InternalTypeName,
	properties ...*astmodel.PropertyDefinition,
) error {
	var errs []error
	for _, property := range properties {
		if err := repairer.AddProperty(name, property); err != nil {
			errs = append(errs, err)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return eris.Wrapf(err, "adding properties from %s", name)
	}

	return nil
}

// AddProperty adds a single property from a type to the graph, marking it as observed
func (repairer *skippingPropertyRepairer) AddProperty(
	name astmodel.InternalTypeName,
	property *astmodel.PropertyDefinition,
) error {
	ref := astmodel.MakePropertyReference(name, property.PropertyName())
	if err := repairer.establishPropertyChain(ref); err != nil {
		return eris.Wrapf(err, "adding property %s", property.PropertyName())
	}

	repairer.propertyObserved(ref)
	return nil
}

// RepairSkippedProperties scans for properties that skip versions, and injects new types to repair the chain, returning
// a (possibly empty) set of new definitions.
func (repairer *skippingPropertyRepairer) RepairSkippedProperties() (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)
	chains := repairer.findChains().AsSlice()
	var errs []error
	for _, ref := range chains {
		defs, err := repairer.repairChain(ref)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// If the repair added any new types (mostly it won't), include them in the result.
		// Duplicates are allowed as long as they are structurally identical, as the same type
		// may be reached from multiple chains in a given API version.
		err = result.AddTypesAllowDuplicates(defs)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	if len(errs) > 0 {
		return nil, eris.Wrapf(
			kerrors.NewAggregate(errs),
			"failed to repair skipping properties")
	}

	return result, nil
}

// establishPropertyChain ensures that a full property chain exists for the specified property. Any missing links in the
// chain will be created. If the required chain already exists, this is a no-op.
func (repairer *skippingPropertyRepairer) establishPropertyChain(ref astmodel.PropertyReference) error {
	if ref.IsEmpty() || repairer.hasLinkFrom(ref) {
		// Nothing to do
		return nil
	}

	return repairer.createPropertyChain(ref)
}

// createPropertyChain creates a full property chain for the specified property.
// ref is the property reference that specifies the start of our chain.
// It recursively calls establishPropertyChain to avoid creating parts of the chain multiple times.
func (repairer *skippingPropertyRepairer) createPropertyChain(ref astmodel.PropertyReference) error {
	next, err := repairer.conversionGraph.FindNextProperty(ref, repairer.definitions)
	if err != nil {
		return eris.Wrapf(err, "creating property chain link from %s", ref.String())
	}

	repairer.addLink(ref, next)
	return repairer.establishPropertyChain(next)
}

// propertyObserved makes a record that a given property has been observedProperties
func (repairer *skippingPropertyRepairer) propertyObserved(ref astmodel.PropertyReference) {
	repairer.observedProperties.Add(ref)
}

// addLink adds a link between two property references
func (repairer *skippingPropertyRepairer) addLink(ref astmodel.PropertyReference, next astmodel.PropertyReference) {
	repairer.links[ref] = next
}

// findChains finds all the property references that are found only as the start of a chain
func (repairer *skippingPropertyRepairer) findChains() *astmodel.PropertyReferenceSet {
	starts := astmodel.NewPropertyReferenceSet()
	finishes := astmodel.NewPropertyReferenceSet()
	for s, f := range repairer.links {
		starts.Add(s)
		finishes.Add(f)
	}

	return starts.Except(finishes)
}

// hasLinkFrom returns true if we already have a link from the specified property
func (repairer *skippingPropertyRepairer) hasLinkFrom(ref astmodel.PropertyReference) bool {
	_, found := repairer.links[ref]
	return found
}

// repairChain checks for a gap in the specified property chain.
// start is the first property reference in the chain
func (repairer *skippingPropertyRepairer) repairChain(
	start astmodel.PropertyReference,
) (astmodel.TypeDefinitionSet, error) {
	lastObserved, firstMissing := repairer.findBreak(start, repairer.wasPropertyObserved)
	if firstMissing.IsEmpty() {
		// Property was never discontinued
		return nil, nil
	}

	lastMissing, reintroduced := repairer.findBreak(firstMissing, repairer.wasPropertyObserved)
	if reintroduced.IsEmpty() {
		// Property was never reintroduced
		return nil, nil
	}

	// If the properties have the same type, we don't have a break here - so we check the remainder of the chain
	// (This is Ok because the value serialized into the property bag from lastObserved will deserialize into the
	// reintroduced property intact.)
	typesSame := repairer.propertiesHaveStructurallyIdenticalType(lastObserved, reintroduced)
	if typesSame {
		return repairer.repairChain(reintroduced)
	}

	// We've found a skipping property with a different shape. If it's a TypeName we need to repair it.
	// We do this by creating a new type that has the same shape as the missing property, injected just prior to
	// reintroduction.

	lastObservedPropertyType, ok := repairer.lookupPropertyType(lastObserved)
	if !ok {
		// Should never fail, given the way findBreak() works
		panic(fmt.Sprintf("failed to find type for property %s", lastObserved))
	}

	tn, ok := astmodel.AsInternalTypeName(lastObservedPropertyType)
	if !ok {
		// If not a type name, defer to our existing property conversion logic
		// Continue checking the rest of the chain
		return repairer.repairChain(reintroduced)
	}

	// Find the type definition for when the object was last observed
	def := repairer.definitions[tn]

	// Find all the type definitions referenced by this definition (e.g. nested types)
	defs, err := astmodel.FindConnectedDefinitions(repairer.definitions, astmodel.MakeTypeDefinitionSetFromDefinitions(def))
	if err != nil {
		return nil, eris.Wrapf(
			err,
			"failed to find connected definitions from %s",
			tn)
	}

	// Move all of these types into a compatibility package
	compatPkg := astmodel.MakeCompatPackageReference(lastMissing.DeclaringType().InternalPackageReference())
	renamer := astmodel.NewRenamingVisitorFromLambda(
		func(name astmodel.InternalTypeName) astmodel.InternalTypeName {
			return name.WithPackageReference(compatPkg)
		})
	newDefs, err := renamer.RenameAll(defs)
	if err != nil {
		return nil, eris.Wrapf(
			err,
			"failed to rename definitions from %s",
			tn)
	}

	return newDefs, nil
}

// wasPropertyObserved returns true if the property reference has been observedProperties; false otherwise.
func (repairer *skippingPropertyRepairer) wasPropertyObserved(ref astmodel.PropertyReference) bool {
	return repairer.observedProperties.Contains(ref)
}

// findBreak finds a pair of consecutive references where the provided predicate gives a different answer for each.
// ref is the property reference from which to start scanning the chain.
// predicate is a test used to identify the pair of references to return.
// A break is always found at the end of the chain, returning <last>, <empty>.
// If ref is empty, will return <empty>, <empty>
func (repairer *skippingPropertyRepairer) findBreak(
	ref astmodel.PropertyReference,
	predicate func(astmodel.PropertyReference) bool,
) (astmodel.PropertyReference, astmodel.PropertyReference) {
	next := repairer.lookupNext(ref)
	if next.IsEmpty() || predicate(ref) != predicate(next) {
		return ref, next
	}

	return repairer.findBreak(next, predicate)
}

// lookupNext returns the next property in the chain, if any.
// ref is the property reference to look up.
// returns the next property, if found; <empty>> if not.
func (repairer *skippingPropertyRepairer) lookupNext(ref astmodel.PropertyReference) astmodel.PropertyReference {
	if next, ok := repairer.links[ref]; ok {
		return next
	}

	return astmodel.EmptyPropertyReference
}

// propertiesHaveStructurallyIdenticalType returns true if both the passed property-references exist and have the same underlying type
func (repairer *skippingPropertyRepairer) propertiesHaveStructurallyIdenticalType(
	left astmodel.PropertyReference,
	right astmodel.PropertyReference,
) bool {
	leftType, leftOk := repairer.lookupPropertyType(left)
	rightType, rightOk := repairer.lookupPropertyType(right)

	exactlyEqual := leftOk && rightOk && leftType.Equals(rightType, astmodel.EqualityOverrides{})
	if exactlyEqual {
		return true
	}

	// If the types aren't exactly equal, we need to check if they're structurally equal
	return repairer.comparer.areTypesStructurallyEqual(leftType, rightType)
}

// lookupPropertyType accepts a PropertyReference and looks up the actual type of the property, returning true if found,
// or false if not.
func (repairer *skippingPropertyRepairer) lookupPropertyType(ref astmodel.PropertyReference) (astmodel.Type, bool) {
	def, ok := repairer.definitions[ref.DeclaringType()]
	if !ok {
		// Type not found
		return nil, false
	}

	container, ok := astmodel.AsPropertyContainer(def.Type())
	if !ok {
		// Not a property container
		return nil, false
	}

	prop, ok := container.Property(ref.Property())
	if !ok {
		// Not a known property
		return nil, false
	}

	return prop.PropertyType(), true
}

type structuralComparer struct {
	definitions       astmodel.TypeDefinitionSet
	equalityOverrides astmodel.EqualityOverrides
}

func newStructuralComparer(definitions astmodel.TypeDefinitionSet) *structuralComparer {
	result := &structuralComparer{
		definitions: definitions,
	}

	result.equalityOverrides = astmodel.EqualityOverrides{
		InternalTypeName: result.equalTypesReferencedByInternalTypeNames,
		ObjectType:       result.equalObjectTypeStructure,
	}

	return result
}

func (comparer *structuralComparer) areTypesStructurallyEqual(left astmodel.Type, right astmodel.Type) bool {
	return left.Equals(right, comparer.equalityOverrides)
}

func (comparer *structuralComparer) equalObjectTypeStructure(left *astmodel.ObjectType, right *astmodel.ObjectType) bool {
	if left == right {
		return true // short circuit
	}

	// Create a copy of the properties with description removed as we don't care if it matches
	leftProperties := comparer.simplifyProperties(left)
	rightProperties := comparer.simplifyProperties(right)

	return leftProperties.Equals(rightProperties, comparer.equalityOverrides)
}

func (comparer *structuralComparer) equalTypesReferencedByInternalTypeNames(
	left astmodel.InternalTypeName,
	right astmodel.InternalTypeName,
) bool {
	// Look up the definitions referenced by the names and compare them for structural equality
	leftDef, ok := comparer.definitions[left]
	if !ok {
		return false
	}

	rightDef, ok := comparer.definitions[right]
	if !ok {
		return false
	}

	return leftDef.Type().Equals(rightDef.Type(), comparer.equalityOverrides)
}

// simplifyProperties is a helper function used to strip out details of the properties we don't care about
// when evaluating for structural integrity
func (comparer *structuralComparer) simplifyProperties(o *astmodel.ObjectType) astmodel.PropertySet {
	result := astmodel.NewPropertySet()
	o.Properties().ForEach(func(def *astmodel.PropertyDefinition) {
		result.Add(def.WithDescription(""))
	})

	return result
}
