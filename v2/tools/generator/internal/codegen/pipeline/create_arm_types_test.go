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

// We don't need to test everything here because a lot of the common cases are covered in the existing golden
// files tests. We focus on complicated edge cases here, such as flattening. Because of the difficulty in setting these
// things up, we combine testing of the CreateARMTypes and the ApplyARMConversionInterface stages

func TestCreateFlattenedARMType_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()

	cfg := config.NewObjectModelConfiguration()
	createARMTypes := CreateARMTypes(cfg, idFactory, logr.Discard())
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory, cfg)
	flatten := FlattenProperties(logr.Discard())
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState(defs),
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateFlattenedARMTypeWithResourceRef_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(propertyConfiguration *config.PropertyConfiguration) error {
				propertyConfiguration.ARMReference.Set(true)
				return nil
			})).To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	configToARMIDs := ApplyCrossResourceReferencesFromConfig(configuration, logr.Discard())
	crossResourceRefs := TransformCrossResourceReferences(configuration, idFactory)
	createARMTypes := CreateARMTypes(omc, idFactory, logr.Discard())
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory, omc)
	flatten := FlattenProperties(logr.Discard())
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState(defs),
		configToARMIDs,
		crossResourceRefs,
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateFlattenedARMTypeWithConfigMap_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeRequired)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddConfigMaps(configuration)
	createARMTypes := CreateARMTypes(omc, idFactory, logr.Discard())
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory, omc)
	flatten := FlattenProperties(logr.Discard())
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState(defs),
		addConfigMaps,
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateARMTypeWithConfigMap_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		test.RestrictedNameProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeRequired)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.RestrictedNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddConfigMaps(configuration)
	createARMTypes := CreateARMTypes(omc, idFactory, logr.Discard())
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory, omc)
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState(defs),
		addConfigMaps,
		createARMTypes,
		applyARMConversionInterface,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateARMTypeWithSecret_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secretDataProperty := astmodel.NewPropertyDefinition("SecretData", "secretData", astmodel.NewMapType(astmodel.StringType, astmodel.StringType)).
		WithDescription("Secret data")
	secretSliceProperty := astmodel.NewPropertyDefinition("SecretSlice", "secretSlice", astmodel.NewArrayType(astmodel.StringType)).
		WithDescription("Secret data")

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		secretDataProperty,
		secretSliceProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.IsSecret.Set(true)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			secretDataProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.IsSecret.Set(true)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			secretSliceProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.IsSecret.Set(true)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddSecrets(configuration)
	createARMTypes := CreateARMTypes(omc, idFactory, logr.Discard())
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory, omc)
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState(defs),
		addConfigMaps,
		createARMTypes,
		applyARMConversionInterface,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateARMTypeConversionsWhenSimplifying_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()

	aliasDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Alias"),
		astmodel.StringType)

	aliasProperty := astmodel.NewPropertyDefinition(
		"Alias",
		"alias",
		aliasDef.Name()).
		WithDescription("Expect alias on CRD type to become string on ARM type")

	qualificationsDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Qualifications"),
		astmodel.NewArrayType(astmodel.StringType))

	qualificationsProperty := astmodel.NewPropertyDefinition(
		"Qualifications",
		"qualifications",
		qualificationsDef.Name()).
		WithDescription("Expect alias of array on CRD type to become array on ARM type")

	codesDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Codes"),
		astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	codesProperty := astmodel.NewPropertyDefinition(
		"Codes",
		"codes",
		codesDef.Name()).
		WithDescription("Expect alias of map on CRD type to become map on ARM type")

	cases := map[string]struct {
		property    *astmodel.PropertyDefinition
		propertyDef astmodel.TypeDefinition
	}{
		"AliasFlattensToUnderlyingType": {
			property:    aliasProperty,
			propertyDef: aliasDef,
		},
		"AliasOfArrayFlattensToSimpleArray": {
			property:    qualificationsProperty,
			propertyDef: qualificationsDef,
		},
		"AliasOfMapFlattensToSimpleMap": {
			property:    codesProperty,
			propertyDef: codesDef,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Arrange: Create our Person type for ARM conversion
			person := test.CreateObjectDefinition(
				test.Pkg2020,
				"Person",
				test.FullNameProperty,
				c.property)

			// Arrange: Create a set of all our definitions
			defs := astmodel.MakeTypeDefinitionSetFromDefinitions(
				c.propertyDef,
				person)

			idFactory := astmodel.NewIdentifierFactory()
			omc := config.NewObjectModelConfiguration()

			// Act: Run the pipeline
			state, err := RunTestPipeline(
				NewState(defs),
				CreateARMTypes(omc, idFactory, logr.Discard()),
				ApplyARMConversionInterface(idFactory, omc))
			g.Expect(err).ToNot(HaveOccurred())

			test.AssertPackagesGenerateExpectedCode(t, state.Definitions(), test.CreateFolderForTest())
		})
	}
}
