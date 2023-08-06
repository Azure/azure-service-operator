/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
)

const DetectSkippingPropertiesStageID = "detectSkippingProperties"

// DetectSkippingProperties detects any properties that skip one or more versions of an object and are reintroduced.
// As described in issue #1776, such properties are a problem because they might create incompatibilities between
// different versions of ASO, or between different versions of a given resource.
//
// We need to design a long term solution to this problem, but in the meantime we just want to detect the problem to
// ensure we don't inadvertently release an update to ASO with this significant hidden flaw.
//
// Detection works by scanning for properties that are dropped between versions of a resource. We keep track of all
// these properties and can raise an error if a specific property appears more than once (implying there are two or more
// sequence versions of that property).
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
// Address: present in (v1, v3); (skipping v2) issue to report, aborting the pipeline.
//
// Additional complexities:
// - We need to handle type renaming between versions.
// - When introduced, we will also need to handle property renaming between versions
func DetectSkippingProperties() *Stage {
	return NewStage(
		DetectSkippingPropertiesStageID,
		"Detect properties that skip resource or object versions",
		func(ctx context.Context, state *State) (*State, error) {
			detector := newSkippingPropertyDetector(state.Definitions(), state.ConversionGraph())

			// Add resources and objects to the graph
			for _, def := range state.Definitions() {
				if t, ok := astmodel.AsPropertyContainer(def.Type()); ok {
					err := detector.AddProperties(def.Name(), t.Properties().AsSlice()...)
					if err != nil {
						return nil, err
					}
				}
			}

			err := detector.CheckForSkippedProperties()

			return state, err
		})
}

type skippingPropertyDetector struct {
	links              map[astmodel.PropertyReference]astmodel.PropertyReference // Individual links in chains of related properties
	observedProperties *astmodel.PropertyReferenceSet                            // Set of properties we've observed
	definitions        astmodel.TypeDefinitionSet                                // Set of all known type definitions
	conversionGraph    *storage.ConversionGraph                                  // Graph of conversions between types
}

// newSkippingPropertyDetector creates a new graph for tracking chains of properties as they evolve through different
// versions of a resource or object.
// definitions is a set of all known types.
// conversionGraph contains every conversion/transition between versions.
func newSkippingPropertyDetector(definitions astmodel.TypeDefinitionSet, conversionGraph *storage.ConversionGraph) *skippingPropertyDetector {
	return &skippingPropertyDetector{
		links:              make(map[astmodel.PropertyReference]astmodel.PropertyReference),
		observedProperties: astmodel.NewPropertyReferenceSet(),
		definitions:        definitions,
		conversionGraph:    conversionGraph,
	}
}

// AddProperties adds all the properties from the specified type to the graph.
func (detector *skippingPropertyDetector) AddProperties(
	name astmodel.TypeName,
	properties ...*astmodel.PropertyDefinition,
) error {
	var errs []error
	for _, property := range properties {
		if err := detector.AddProperty(name, property); err != nil {
			errs = append(errs, err)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return errors.Wrapf(err, "adding properties from %s", name)
	}

	return nil
}

// AddProperty adds a single property from a type to the graph, marking it as observed
func (detector *skippingPropertyDetector) AddProperty(
	name astmodel.TypeName,
	property *astmodel.PropertyDefinition,
) error {
	ref := astmodel.MakePropertyReference(name, property.PropertyName())
	if err := detector.establishPropertyChain(ref); err != nil {
		return errors.Wrapf(err, "adding property %s", property.PropertyName())
	}

	detector.propertyObserved(ref)
	return nil
}

// CheckForSkippedProperties scans for properties that skip versions, and returns an error summarizing the results
func (detector *skippingPropertyDetector) CheckForSkippedProperties() error {
	chains := detector.findChains().AsSlice()
	errs := make([]error, 0, len(chains))
	for _, ref := range chains {
		err := detector.checkChain(ref)
		errs = append(errs, err)
	}

	return kerrors.NewAggregate(errs)
}

// establishPropertyChain ensures that a full property chain exists for the specified property. Any missing links in the
// chain will be created. If the required chain already exists, this is a no-op.
func (detector *skippingPropertyDetector) establishPropertyChain(ref astmodel.PropertyReference) error {
	if ref.IsEmpty() || detector.hasLinkFrom(ref) {
		// Nothing to do
		return nil
	}

	return detector.createPropertyChain(ref)
}

// createPropertyChain creates a full property chain for the specified property.
// ref is the property reference that specifies the start of our chain.
// It recursively calls establishPropertyChain to avoid creating parts of the chain multiple times.
func (detector *skippingPropertyDetector) createPropertyChain(ref astmodel.PropertyReference) error {
	next, err := detector.conversionGraph.FindNextProperty(ref, detector.definitions)
	if err != nil {
		return errors.Wrapf(err, "creating property chain link from %s", ref.String())
	}

	detector.addLink(ref, next)
	return detector.establishPropertyChain(next)
}

// propertyObserved makes a record that a given property has been observedProperties
func (detector *skippingPropertyDetector) propertyObserved(ref astmodel.PropertyReference) {
	detector.observedProperties.Add(ref)
}

// addLink adds a link between two property references
func (detector *skippingPropertyDetector) addLink(ref astmodel.PropertyReference, next astmodel.PropertyReference) {
	detector.links[ref] = next
}

// findChains finds all the property references that are found only as the start of a chain
func (detector *skippingPropertyDetector) findChains() *astmodel.PropertyReferenceSet {
	starts := astmodel.NewPropertyReferenceSet()
	finishes := astmodel.NewPropertyReferenceSet()
	for s, f := range detector.links {
		starts.Add(s)
		finishes.Add(f)
	}

	return starts.Except(finishes)
}

// hasLinkFrom returns true if we already have a link from the specified property
func (detector *skippingPropertyDetector) hasLinkFrom(ref astmodel.PropertyReference) bool {
	_, found := detector.links[ref]
	return found
}

// checkChain checks for a gap in the specified property chain.
// start is the first property reference in the chain
func (detector *skippingPropertyDetector) checkChain(start astmodel.PropertyReference) error {
	lastObserved, firstMissing := detector.findBreak(start, detector.wasPropertyObserved)
	if firstMissing.IsEmpty() {
		// Property was never discontinued
		return nil
	}

	_, reintroduced := detector.findBreak(firstMissing, detector.wasPropertyObserved)
	if reintroduced.IsEmpty() {
		// Property was never reintroduced
		return nil
	}

	// If the properties have the same type, we don't have a break here - so we check the remainder of the chain
	// (This is Ok because the value serialized into the property bag from lastObserved will deserialize into the
	// reintroduced property intact.)
	typesSame, err := detector.propertiesHaveSameType(lastObserved, reintroduced)
	if err != nil {
		return errors.Wrapf(err, "failed to determine if properties %s and %s have the same type", lastObserved, reintroduced)
	}
	if typesSame {
		return detector.checkChain(reintroduced)
	}

	return errors.Errorf(
		"property %s was discontinued but later reintroduced as %s with a different type; "+
			"see https://github.com/Azure/azure-service-operator/issues/1776 for why this is a problem",
		lastObserved,
		reintroduced)
}

// wasPropertyObserved returns true if the property reference has been observedProperties; false otherwise.
func (detector *skippingPropertyDetector) wasPropertyObserved(ref astmodel.PropertyReference) bool {
	return detector.observedProperties.Contains(ref)
}

// findBreak finds a pair of consecutive references where the provided predicate gives a different answer for each.
// ref is the property reference from which to start scanning the chain.
// predicate is a test used to identify the pair of references to return.
// A break is always found at the end of the chain, returning <last>, <empty>.
// If ref is empty, will return <empty>, <empty>
func (detector *skippingPropertyDetector) findBreak(
	ref astmodel.PropertyReference,
	predicate func(astmodel.PropertyReference) bool,
) (astmodel.PropertyReference, astmodel.PropertyReference) {
	next := detector.lookupNext(ref)
	if next.IsEmpty() || predicate(ref) != predicate(next) {
		return ref, next
	}

	return detector.findBreak(next, predicate)
}

// lookupNext returns the next property in the chain, if any.
// ref is the property reference to look up.
// returns the next property, if found; <empty>> if not.
func (detector *skippingPropertyDetector) lookupNext(ref astmodel.PropertyReference) astmodel.PropertyReference {
	if next, ok := detector.links[ref]; ok {
		return next
	}

	return astmodel.EmptyPropertyReference
}

// propertiesHaveSameType returns true if both the passed property references exist and have the same underlying type
func (detector *skippingPropertyDetector) propertiesHaveSameType(
	left astmodel.PropertyReference,
	right astmodel.PropertyReference,
) (bool, error) {
	leftType, leftOk := detector.lookupPropertyType(left)
	rightType, rightOk := detector.lookupPropertyType(right)

	exactlyEqual := leftOk && rightOk && leftType.Equals(rightType, astmodel.EqualityOverrides{})
	if exactlyEqual {
		return true, nil
	}

	equalityOverrides := astmodel.EqualityOverrides{
		TypeName: compareTypeNamesIgnoreVersion,
	}
	equalSameTypeNameDifferentVersion := leftOk && rightOk && leftType.Equals(rightType, equalityOverrides)
	if !equalSameTypeNameDifferentVersion {
		return false, nil
	}

	leftTypeName, leftOk := astmodel.ExtractTypeName(leftType)
	rightTypeName, rightOk := astmodel.ExtractTypeName(rightType)

	// If either isn't a typeName, return false
	// Note that practically speaking actually taking this codepath should be impossible, as to get here the
	// two properties must be: Not exactly equal, but equal if typeName versions are ignored.
	if !leftOk || !rightOk {
		return false, nil
	}

	leftTypeDef, err := detector.definitions.GetDefinition(leftTypeName)
	if err != nil {
		return false, err
	}
	rightTypeDef, err := detector.definitions.GetDefinition(rightTypeName)
	if err != nil {
		return false, err
	}

	// If the types don't match exactly, we have to determine if they match structurally
	leftDefs, err := astmodel.FindConnectedDefinitions(detector.definitions, astmodel.MakeTypeDefinitionSetFromDefinitions(leftTypeDef))
	if err != nil {
		return false, err
	}
	rightDefs, err := astmodel.FindConnectedDefinitions(detector.definitions, astmodel.MakeTypeDefinitionSetFromDefinitions(rightTypeDef))
	if err != nil {
		return false, err
	}

	return areTypeSetsEqual(leftDefs, rightDefs), nil
}

func areTypeSetsEqual(left astmodel.TypeDefinitionSet, right astmodel.TypeDefinitionSet) bool {
	if len(left) != len(right) {
		return false
	}

	packageRefs := set.Make[astmodel.PackageReference]()
	for name := range right {
		packageRefs.Add(name.PackageReference())
	}

	if len(packageRefs) != 1 {
		panic(fmt.Sprintf("expected rhs type names to all be from one package, but were from %d packages instead", len(packageRefs)))
	}

	rightPackageRef := packageRefs.Values()[0]
	equalityOverrides := astmodel.EqualityOverrides{
		TypeName:   compareTypeNamesIgnoreVersion,
		ObjectType: compareObjectTypeStructure,
	}

	for leftName, leftDef := range left {
		rightName := leftName.WithPackageReference(rightPackageRef)
		rightDef := right[rightName]

		equal := leftDef.Type().Equals(rightDef.Type(), equalityOverrides)
		if !equal {
			return false
		}
	}

	return true
}

func compareTypeNamesIgnoreVersion(left astmodel.TypeName, right astmodel.TypeName) bool {
	leftLPR, isLeftLocalRef := left.PackageReference().(astmodel.LocalLikePackageReference)
	rightLPR, isRightLocalRef := right.PackageReference().(astmodel.LocalLikePackageReference)

	// If we're not looking at local references, use the standard equality comparison
	if !isLeftLocalRef || !isRightLocalRef {
		return left.Equals(right, astmodel.EqualityOverrides{})
	}

	// If we are looking at local references, we allow them to differ by api-version and generator-version
	if leftLPR.LocalPathPrefix() != rightLPR.LocalPathPrefix() ||
		leftLPR.Group() != rightLPR.Group() {
		return false
	}

	// The names must be the same
	return left.Name() == right.Name()
}

func compareObjectTypeStructure(left *astmodel.ObjectType, right *astmodel.ObjectType) bool {
	if left == right {
		return true // short circuit
	}

	equalityOverrides := astmodel.EqualityOverrides{
		TypeName: compareTypeNamesIgnoreVersion,
	}

	// Create a copy of the properties with description removed as we don't care if it matches
	leftProperties := astmodel.NewPropertySet()
	rightProperties := astmodel.NewPropertySet(left.Properties().AsSlice()...)
	left.Properties().ForEach(func(def *astmodel.PropertyDefinition) {
		leftProperties.Add(def.WithDescription(""))
	})
	right.Properties().ForEach(func(def *astmodel.PropertyDefinition) {
		rightProperties.Add(def.WithDescription(""))
	})

	if !leftProperties.Equals(rightProperties, equalityOverrides) {
		return false
	}

	return true
}

// lookupPropertyType accepts a PropertyReference and looks up the actual type of the property
func (detector *skippingPropertyDetector) lookupPropertyType(ref astmodel.PropertyReference) (astmodel.Type, bool) {
	def, ok := detector.definitions[ref.DeclaringType()]
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
