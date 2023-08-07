/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

var (
	Group    = "lions"
	LionsPkg = test.MakeLocalPackageReference(Group, "20200601")

	// Lion
	LionSpecDef = astmodel.MakeTypeDefinition(
		LionSpecName,
		LionSpecType)
	LionSpecType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionTypeProperty,
		APIVersionProperty,
		LionPropertiesProperty).WithIsResource(true).WithResource(LionName)
	LionStatusDef = astmodel.MakeTypeDefinition(
		LionStatusName,
		LionStatusType)
	LionStatusType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionTypeProperty,
		APIVersionProperty,
		LionPropertiesProperty)

	// LionPride
	LionPrideSpecDef = astmodel.MakeTypeDefinition(
		LionPrideSpecName,
		LionPrideSpecType)
	LionPrideSpecType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionPrideTypeProperty,
		APIVersionProperty,
		LionPridePropertiesProperty).WithIsResource(true).WithResource(LionPrideName)
	LionPrideStatusDef = astmodel.MakeTypeDefinition(
		LionPrideStatusName,
		LionPrideStatusType)
	LionPrideStatusType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionPrideTypeProperty,
		APIVersionProperty,
		LionPridePropertiesProperty)

	// LionCub
	LionCubSpecDef = astmodel.MakeTypeDefinition(
		LionCubSpecName,
		LionCubSpecType)
	LionCubSpecType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionCubTypeProperty,
		APIVersionProperty,
		LionCubPropertiesProperty).WithIsResource(true).WithResource(LionCubName)
	LionCubStatusDef = astmodel.MakeTypeDefinition(
		LionCubStatusName,
		LionCubStatusType)
	LionCubStatusType = astmodel.NewObjectType().WithProperties(
		NameProperty,
		LionCubTypeProperty,
		APIVersionProperty,
		LionCubPropertiesProperty)

	LionResourceType      = astmodel.NewResourceType(LionSpecName, LionStatusName)
	LionPrideResourceType = astmodel.NewResourceType(LionPrideSpecName, LionPrideStatusName)
	LionCubResourceType   = astmodel.NewResourceType(LionCubSpecName, LionCubStatusName)
	Lion                  = astmodel.MakeTypeDefinition(LionName, LionResourceType)
	LionPride             = astmodel.MakeTypeDefinition(LionPrideName, LionPrideResourceType)
	LionCub               = astmodel.MakeTypeDefinition(LionCubName, LionCubResourceType)

	LionRef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionRef"),
		LionSpecType.WithProperty(IDProperty))
	LionCubRef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionCubRef"),
		LionCubSpecType.WithProperty(IDProperty))

	LionPrideTypeEnumDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionPrideType"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("type", "Microsoft.Lions/lionPride")))
	LionTypeEnumDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionType"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("type", "Microsoft.Lions/lion")))
	LionCubTypeEnumDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionCubType"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("type", "Microsoft.Lions/cub")))
	APIVersionTypeEnumDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "APIVersion"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("apiVersion", "2020-06-01")))

	LionPropertiesDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionProperties"),
		astmodel.NewObjectType().WithProperty(RoarVolumeProperty))
	LionPridePropertiesDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionPrideProperties"),
		astmodel.NewObjectType().WithProperties(LionPrideLionsProperty, HuntsProperty))
	LionCubPropertiesDef = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "LionCubProperties"),
		astmodel.NewObjectType().WithProperties(LionCubCutenessProperty))

	// Properties
	NameProperty          = astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType).MakeTypeOptional().MakeRequired()
	LionPrideTypeProperty = astmodel.NewPropertyDefinition("Type", "type", LionPrideTypeEnumDef.Name()).MakeTypeOptional().MakeRequired()
	LionTypeProperty      = astmodel.NewPropertyDefinition("Type", "type", LionTypeEnumDef.Name()).MakeTypeOptional().MakeRequired()
	LionCubTypeProperty   = astmodel.NewPropertyDefinition("Type", "type", LionCubTypeEnumDef.Name()).MakeTypeOptional().MakeRequired()
	APIVersionProperty    = astmodel.NewPropertyDefinition("APIVersion", "apiVersion", APIVersionTypeEnumDef.Name()).MakeTypeOptional().MakeRequired()
	IDProperty            = astmodel.NewPropertyDefinition("Id", "id", astmodel.ARMIDType)

	LionPridePropertiesProperty = astmodel.NewPropertyDefinition("Properties", "properties", LionPridePropertiesDef.Name())
	LionPropertiesProperty      = astmodel.NewPropertyDefinition("Properties", "properties", LionPropertiesDef.Name())
	LionCubPropertiesProperty   = astmodel.NewPropertyDefinition("Properties", "properties", LionCubPropertiesDef.Name())
	LionPrideLionsProperty      = astmodel.NewPropertyDefinition("Lions", "lions", astmodel.NewArrayType(LionRef.Name()))
	LionCubCutenessProperty     = astmodel.NewPropertyDefinition("Cuteness", "cuteness", astmodel.IntType)
	RoarVolumeProperty          = astmodel.NewPropertyDefinition("RoarVolumeDecibels", "roarVolumeDecibels", astmodel.IntType)
	HuntsProperty               = astmodel.NewPropertyDefinition("Hunts", "hunts", astmodel.IntType)
	CubsProperty                = astmodel.NewPropertyDefinition("Cubs", "cubs", astmodel.NewArrayType(LionCubRef.Name()))

	// Names
	LionName            = astmodel.MakeTypeName(LionsPkg, "Lion")
	LionSpecName        = test.MakeSpecName(LionsPkg, "Lion")
	LionStatusName      = test.MakeStatusName(LionsPkg, "Lion")
	LionPrideName       = astmodel.MakeTypeName(LionsPkg, "LionPride")
	LionPrideSpecName   = test.MakeSpecName(LionsPkg, "LionPride")
	LionPrideStatusName = test.MakeStatusName(LionsPkg, "LionPride")
	LionCubName         = astmodel.MakeTypeName(LionsPkg, "LionCub")
	LionCubSpecName     = test.MakeSpecName(LionsPkg, "LionCub")
	LionCubStatusName   = test.MakeStatusName(LionsPkg, "LionCub")
)

func LionDefs() astmodel.TypeDefinitionSet {
	defs := make(astmodel.TypeDefinitionSet)

	defs.AddAll(
		Lion,
		LionSpecDef,
		LionStatusDef,
		LionPropertiesDef,
		LionTypeEnumDef,
		LionPride,
		LionPrideSpecDef,
		LionPrideStatusDef,
		LionPridePropertiesDef,
		LionPrideTypeEnumDef,
		APIVersionTypeEnumDef,
		LionRef)

	return defs
}

func LionCubDefs() astmodel.TypeDefinitionSet {
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(
		LionCub,
		LionCubSpecDef,
		LionCubStatusDef,
		LionCubPropertiesDef,
		LionCubTypeEnumDef,
		LionCubRef)

	return defs
}

func TestGolden_EmbeddedSubresource_IsRemoved(t *testing.T) {
	t.Parallel()

	g := NewWithT(t)

	defs := LionDefs()

	// Add owner
	defs[Lion.Name()] = Lion.WithType(LionResourceType.WithOwner(LionPrideName))

	// TODO: Take this bit and put it into a helper?
	// Define stages to run
	configuration := config.NewConfiguration()
	removeEmbedded := RemoveEmbeddedResources(configuration, logr.Discard())

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		removeEmbedded)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestGolden_EmbeddedResource_IsRemovedRetainsId(t *testing.T) {
	t.Parallel()

	g := NewWithT(t)

	defs := LionDefs()

	// Define stages to run
	configuration := config.NewConfiguration()
	removeEmbedded := RemoveEmbeddedResources(configuration, logr.Discard())

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		removeEmbedded)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestGolden_EmbeddedResourcesWithMultipleEmbeddings_AllEmbeddingsAreRemovedAndRetainOnlyId(t *testing.T) {
	t.Parallel()

	g := NewWithT(t)

	defs := LionDefs()
	defs.AddTypes(LionCubDefs())

	// Add cubs property so that Lion is embedded in LionPride, and LionCub is embedded in Lion
	ot := LionPropertiesDef.Type().(*astmodel.ObjectType)
	defs[LionPropertiesDef.Name()] = LionPropertiesDef.WithType(ot.WithProperty(CubsProperty))

	// Define stages to run
	configuration := config.NewConfiguration()
	removeEmbedded := RemoveEmbeddedResources(configuration, logr.Discard())

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		removeEmbedded)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestGolden_EmbeddedResourceWithCyclesAndResourceLookalikes_RemovesCycles(t *testing.T) {
	t.Parallel()

	g := NewWithT(t)

	defs := LionDefs()

	ot := LionRef.Type().(*astmodel.ObjectType)
	unlabelledRef := LionRef.WithType(ot.WithIsResource(false).ClearResources())
	defs[unlabelledRef.Name()] = unlabelledRef

	// This just trips our resource-lookalike detection
	dummyPropertiesProperty := astmodel.NewPropertyDefinition("Properties", "properties", astmodel.StringType)

	refProperty := astmodel.NewPropertyDefinition("Ref", "ref", unlabelledRef.Name())
	left := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "Left"),
		astmodel.NewObjectType().WithProperties(
			NameProperty,
			dummyPropertiesProperty,
			refProperty))
	right := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(LionsPkg, "Right"),
		astmodel.NewObjectType().WithProperties(
			NameProperty,
			dummyPropertiesProperty,
			refProperty))

	// Lion needs to point to left and right
	leftProperty := astmodel.NewPropertyDefinition("Left", "left", left.Name())
	rightProperty := astmodel.NewPropertyDefinition("Right", "right", right.Name())
	defs[LionPropertiesDef.Name()] = astmodel.MakeTypeDefinition(
		LionPropertiesDef.Name(),
		astmodel.NewObjectType().WithProperties(
			leftProperty,
			rightProperty))

	// Remove the lions property because it adds some confusion/mess to the generated files and doesn't actually
	// target our scenario anyway
	ot = LionPridePropertiesDef.Type().(*astmodel.ObjectType)
	defs[LionPridePropertiesDef.Name()] = LionPridePropertiesDef.WithType(ot.WithoutProperty(LionPrideLionsProperty.PropertyName()).WithProperties(leftProperty, rightProperty))

	defs.AddAll(left, right)

	// Define stages to run
	configuration := config.NewConfiguration()
	removeEmbedded := RemoveEmbeddedResources(configuration, logr.Discard())

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		removeEmbedded)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}
