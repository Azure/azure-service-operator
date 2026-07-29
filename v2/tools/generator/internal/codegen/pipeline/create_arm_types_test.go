/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

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
		test.KnownAsProperty,
	)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name(),
	).SetFlatten(true).MakeTypeOptional()
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
		strip,
	)
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
		test.KnownAsProperty,
	)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name(),
	).SetFlatten(true).MakeTypeOptional()
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
				propertyConfiguration.ReferenceType.Set(config.ReferenceTypeARM)
				return nil
			},
		),
	).To(Succeed())

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
		strip,
	)
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
		test.KnownAsProperty,
	)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name(),
	).SetFlatten(true).MakeTypeOptional()
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
			},
		),
	).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			},
		),
	).
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
		strip,
	)
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
		test.RestrictedNameProperty,
	)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name(),
	).MakeTypeOptional()
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
			},
		),
	).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			},
		),
	).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.RestrictedNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			},
		),
	).
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
		strip,
	)
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
		secretSliceProperty,
	)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name(),
	).MakeTypeOptional()
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
				pc.Secrecy.Set(astmodel.ImportSecretModeRequired)
				return nil
			},
		),
	).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			secretDataProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.Secrecy.Set(astmodel.ImportSecretModeRequired)
				return nil
			},
		),
	).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			secretSliceProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.Secrecy.Set(astmodel.ImportSecretModeRequired)
				return nil
			},
		),
	).
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
		strip,
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateARMTypeConversionsWhenSimplifying_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()

	aliasDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Alias"),
		astmodel.StringType,
	)

	aliasProperty := astmodel.NewPropertyDefinition(
		"Alias",
		"alias",
		aliasDef.Name(),
	).
		WithDescription("Expect alias on CRD type to become string on ARM type")

	qualificationsDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Qualifications"),
		astmodel.NewArrayType(astmodel.StringType),
	)

	qualificationsProperty := astmodel.NewPropertyDefinition(
		"Qualifications",
		"qualifications",
		qualificationsDef.Name(),
	).
		WithDescription("Expect alias of array on CRD type to become array on ARM type")

	codesDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Codes"),
		astmodel.NewMapType(astmodel.StringType, astmodel.StringType),
	)

	codesProperty := astmodel.NewPropertyDefinition(
		"Codes",
		"codes",
		codesDef.Name(),
	).
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
				c.property,
			)

			// Arrange: Create a set of all our definitions
			defs := astmodel.MakeTypeDefinitionSetFromDefinitions(
				c.propertyDef,
				person,
			)

			idFactory := astmodel.NewIdentifierFactory()
			omc := config.NewObjectModelConfiguration()

			// Act: Run the pipeline
			state, err := RunTestPipeline(
				NewState(defs),
				CreateARMTypes(omc, idFactory, logr.Discard()),
				ApplyARMConversionInterface(idFactory, omc),
			)
			g.Expect(err).ToNot(HaveOccurred())

			test.AssertPackagesGenerateExpectedCode(t, state.Definitions(), test.CreateFolderForTest())
		})
	}
}

func TestCreateARMTypes_WithTopLevelOneOf_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// When a spec or status type is itself a one-of type, we have a conflict between the
	// requirement that spec and status types have specific properties defined, and the requirement
	// that all properties on one-of types are pushed down to the leaves.
	//
	// We compromise by permitting a limited number of properties to remain at the top level (at
	// the moment, just Name) but this means we need to proactively push values down when
	// serializing to ARM, and pull those values back up again when deserializing.
	//
	// This test checks that the generated code correctly handles this situation.
	// The object structure created here mirrors that for Kusto/ClusterDatabase

	readWriteKind := astmodel.MakeEnumValue("ReadWrite", `"ReadWrite"`)
	readWriteDatabaseKind := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2022, "ReadWriteDatabaseKind"),
		astmodel.NewEnumType(
			astmodel.StringType,
			readWriteKind,
		),
	)

	readwriteDatabase := test.CreateObjectDefinition(
		test.Pkg2022,
		"ReadWriteDatabase",
		astmodel.NewPropertyDefinition("Kind", "kind", readWriteDatabaseKind.Name()),
		astmodel.NewPropertyDefinition("Location", "location", astmodel.OptionalStringType),
		astmodel.NewPropertyDefinition("Properties", "properties", astmodel.OptionalStringType),
	)

	readOnlyFollowingKind := astmodel.MakeEnumValue("ReadOnlyFollowing", `"ReadOnlyFollowing"`)
	readOnlyFollowingDatabaseKind := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2022, "ReadOnlyFollowingDatabaseKind"),
		astmodel.NewEnumType(
			astmodel.StringType,
			readOnlyFollowingKind,
		),
	)

	readOnlyFollowingDatabase := test.CreateObjectDefinition(
		test.Pkg2022,
		"ReadOnlyFollowingDatabase",
		astmodel.NewPropertyDefinition("Kind", "kind", readOnlyFollowingDatabaseKind.Name()),
		astmodel.NewPropertyDefinition("Location", "location", astmodel.OptionalStringType),
		astmodel.NewPropertyDefinition("Properties", "properties", astmodel.OptionalStringType),
	)

	clusterDatabaseSpec := test.CreateObjectDefinition(
		test.Pkg2022,
		"ClusterDatabase_Spec",
		astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType),
		astmodel.NewPropertyDefinition("ReadOnlyFollowing", "readOnlyFollowing", astmodel.NewOptionalType(readOnlyFollowingDatabase.Name())),
		astmodel.NewPropertyDefinition("ReadWrite", "readWrite", astmodel.NewOptionalType(readwriteDatabase.Name())),
	)

	var err error
	clusterDatabaseSpec, err = clusterDatabaseSpec.ApplyObjectTransformation(
		func(o *astmodel.ObjectType) (astmodel.Type, error) {
			return astmodel.OneOfFlag.ApplyTo(o), nil
		},
	)
	g.Expect(err).NotTo(HaveOccurred())

	clusterDatabaseStatus := test.CreateObjectDefinition(
		test.Pkg2022,
		"ClusterDatabase_Status",
		astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType),
	)

	clusterResource := test.CreateResource(
		test.Pkg2022,
		"ClusterDatabase",
		clusterDatabaseSpec,
		clusterDatabaseStatus,
	)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(
		readwriteDatabase,
		readOnlyFollowingDatabase,
		clusterDatabaseSpec,
		clusterDatabaseStatus,
		clusterResource,
		readWriteDatabaseKind,
		readOnlyFollowingDatabaseKind,
		test.Pkg2020APIVersion,
	)

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
		strip,
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_SimpleResourceRendersSpec tests that a minimal ARM resource with just standard
// properties (name, type, apiVersion) generates correct ARM types and conversions.
func TestCreateARMTypes_SimpleResourceRendersSpec(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := test.CreateSpec(test.Pkg2020, "FakeResource", test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_SimpleResourceComplexProperties tests that an ARM resource with nested object
// references and enum properties generates correct ARM types and conversions.
func TestCreateARMTypes_SimpleResourceComplexProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define the Foo type
	fooNameProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType)
	foo := test.CreateObjectDefinition(test.Pkg2020, "Foo", fooNameProperty)

	// Define color enum
	colorEnum := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Color"),
		astmodel.NewEnumType(
			astmodel.StringType,
			astmodel.MakeEnumValue("blue", `"blue"`),
			astmodel.MakeEnumValue("green", `"green"`),
			astmodel.MakeEnumValue("red", `"red"`),
		),
	)

	// Create spec with complex properties
	optionalFooProp := astmodel.NewPropertyDefinition(
		"OptionalFoo",
		"optionalFoo",
		astmodel.NewOptionalType(foo.Name()),
	)
	requiredFooProp := astmodel.NewPropertyDefinition(
		"Foo",
		"foo",
		foo.Name(),
	).MakeTypeOptional().MakeRequired()
	colorProp := astmodel.NewPropertyDefinition(
		"Color",
		"color",
		astmodel.NewOptionalType(colorEnum.Name()),
	)

	spec := test.CreateSpec(
		test.Pkg2020,
		"FakeResource",
		test.NameProperty,
		optionalFooProp,
		requiredFooProp,
		colorProp,
	)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, foo, colorEnum, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_SimpleResourceArrayProperties tests that an ARM resource with various array
// property types (arrays of objects, arrays of arrays, arrays of maps, arrays of enums) generates
// correct ARM types and conversions.
func TestCreateARMTypes_SimpleResourceArrayProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define the Foo type
	fooNameProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType)
	foo := test.CreateObjectDefinition(test.Pkg2020, "Foo", fooNameProperty)

	// Define Color enum
	colorEnum := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Color"),
		astmodel.NewEnumType(
			astmodel.StringType,
			astmodel.MakeEnumValue("blue", `"blue"`),
			astmodel.MakeEnumValue("green", `"green"`),
			astmodel.MakeEnumValue("red", `"red"`),
		),
	)

	// Array of Foo (required)
	arrayFooProp := astmodel.NewPropertyDefinition(
		"ArrayFoo",
		"arrayFoo",
		astmodel.NewArrayType(foo.Name()),
	).MakeTypeOptional().MakeRequired()
	// Array of arrays of Foo
	arrayOfArraysProp := astmodel.NewPropertyDefinition(
		"ArrayOfArrays",
		"arrayOfArrays",
		astmodel.NewArrayType(astmodel.NewArrayType(foo.Name())),
	)
	// Array of arrays of arrays of Foo
	arrayOfArraysOfArraysProp := astmodel.NewPropertyDefinition(
		"ArrayOfArraysOfArrays",
		"arrayOfArraysOfArrays",
		astmodel.NewArrayType(astmodel.NewArrayType(astmodel.NewArrayType(foo.Name()))),
	)
	// Array of maps of Foo
	arrayOfMapsProp := astmodel.NewPropertyDefinition(
		"ArrayOfMaps",
		"arrayOfMaps",
		astmodel.NewArrayType(astmodel.NewMapType(astmodel.StringType, foo.Name())),
	)
	// Array of enums
	arrayOfEnumsProp := astmodel.NewPropertyDefinition(
		"ArrayOfEnums",
		"arrayOfEnums",
		astmodel.NewArrayType(colorEnum.Name()),
	)

	spec := test.CreateSpec(
		test.Pkg2020,
		"FakeResource",
		test.NameProperty,
		arrayFooProp,
		arrayOfArraysProp,
		arrayOfArraysOfArraysProp,
		arrayOfMapsProp,
		arrayOfEnumsProp,
	)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, foo, colorEnum, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_SimpleResourceMapProperties tests that an ARM resource with various map property
// types (maps of objects, maps of maps, maps of arrays, maps of enums, maps of strings) generates
// correct ARM types and conversions.
func TestCreateARMTypes_SimpleResourceMapProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define the Foo type (note: uses "fooName" in the map properties JSON)
	fooNameProperty := astmodel.NewPropertyDefinition("FooName", "fooName", astmodel.OptionalStringType)
	foo := test.CreateObjectDefinition(test.Pkg2020, "Foo", fooNameProperty)

	// Define Color enum
	colorEnum := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "Color"),
		astmodel.NewEnumType(
			astmodel.StringType,
			astmodel.MakeEnumValue("blue", `"blue"`),
			astmodel.MakeEnumValue("green", `"green"`),
			astmodel.MakeEnumValue("red", `"red"`),
		),
	)

	// Map of Foo (required)
	mapFooProp := astmodel.NewPropertyDefinition(
		"MapFoo",
		"mapFoo",
		astmodel.NewMapType(astmodel.StringType, foo.Name()),
	).MakeTypeOptional().MakeRequired()
	// Map of maps of Foo
	mapOfMapsProp := astmodel.NewPropertyDefinition(
		"MapOfMaps",
		"mapOfMaps",
		astmodel.NewMapType(astmodel.StringType, astmodel.NewMapType(astmodel.StringType, foo.Name())),
	)
	// Map of arrays of Foo
	mapOfArraysProp := astmodel.NewPropertyDefinition(
		"MapOfArrays",
		"mapOfArrays",
		astmodel.NewMapType(astmodel.StringType, astmodel.NewArrayType(foo.Name())),
	)
	// Map of enums
	mapOfEnumsProp := astmodel.NewPropertyDefinition(
		"MapOfEnums",
		"mapOfEnums",
		astmodel.NewMapType(astmodel.StringType, colorEnum.Name()),
	)
	// Map of strings
	mapOfStringsProp := astmodel.NewPropertyDefinition(
		"MapOfStrings",
		"mapOfStrings",
		astmodel.NewMapType(astmodel.StringType, astmodel.StringType),
	)
	// Map of maps of maps of strings
	mapOfMapsOfMapsOfStringsProp := astmodel.NewPropertyDefinition(
		"MapOfMapsOfMapsOfStrings",
		"mapOfMapsOfMapsOfStrings",
		astmodel.NewMapType(astmodel.StringType,
			astmodel.NewMapType(astmodel.StringType,
				astmodel.NewMapType(astmodel.StringType, astmodel.StringType))),
	)
	// Map of maps of maps of Foo
	mapOfMapsOfMapsProp := astmodel.NewPropertyDefinition(
		"MapOfMapsOfMaps",
		"mapOfMapsOfMaps",
		astmodel.NewMapType(astmodel.StringType,
			astmodel.NewMapType(astmodel.StringType,
				astmodel.NewMapType(astmodel.StringType, foo.Name()))),
	)

	spec := test.CreateSpec(
		test.Pkg2020,
		"FakeResource",
		test.NameProperty,
		mapFooProp,
		mapOfMapsProp,
		mapOfArraysProp,
		mapOfEnumsProp,
		mapOfStringsProp,
		mapOfMapsOfMapsOfStringsProp,
		mapOfMapsOfMapsProp,
	)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, foo, colorEnum, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_SimpleResourceJSONFields tests that an ARM resource with raw JSON fields
// (untyped schemas and object-typed schemas) generates correct ARM types and conversions.
func TestCreateARMTypes_SimpleResourceJSONFields(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// mandatoryJson: untyped JSON (empty schema {}), required
	mandatoryJsonProp := astmodel.NewPropertyDefinition(
		"MandatoryJson",
		"mandatoryJson",
		astmodel.AnyType,
	).MakeTypeOptional().MakeRequired()
	// optionalJson: untyped JSON (empty schema {}), optional
	optionalJsonProp := astmodel.NewPropertyDefinition(
		"OptionalJson",
		"optionalJson",
		astmodel.AnyType,
	).MakeTypeOptional()
	// jsonObject: {"type": "object"} → map[string]interface{}, required
	jsonObjectProp := astmodel.NewPropertyDefinition(
		"JsonObject",
		"jsonObject",
		astmodel.NewMapType(astmodel.StringType, astmodel.AnyType),
	).MakeTypeOptional().MakeRequired()

	spec := test.CreateSpec(
		test.Pkg2020,
		"FakeResource",
		test.NameProperty,
		mandatoryJsonProp,
		optionalJsonProp,
		jsonObjectProp,
	)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		ReplaceAnyTypeWithJSON(),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_OneOfResourceConversion tests that a resource with a oneOf discriminated union
// in its properties generates correct ARM types and conversions.
func TestCreateARMTypes_OneOfResourceConversion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Discriminator enums for each variant
	fooDiscrim := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "FooDiscrim"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("foo", `"foo"`)),
	)
	barDiscrim := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "BarDiscrim"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("bar", `"bar"`)),
	)
	bazDiscrim := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "BazDiscrim"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.MakeEnumValue("baz", `"baz"`)),
	)

	// Foo variant: discrim (required) + name
	fooObj := test.CreateObjectDefinition(
		test.Pkg2020,
		"Foo",
		astmodel.NewPropertyDefinition("Discrim", "discrim", astmodel.NewOptionalType(fooDiscrim.Name())).MakeRequired(),
		astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType),
	)

	// Bar variant: discrim (required) + size (required)
	barObj := test.CreateObjectDefinition(
		test.Pkg2020,
		"Bar",
		astmodel.NewPropertyDefinition("Discrim", "discrim", astmodel.NewOptionalType(barDiscrim.Name())).MakeRequired(),
		astmodel.NewPropertyDefinition("Size", "size", astmodel.NewOptionalType(astmodel.IntType)).MakeRequired(),
	)

	// Baz variant: discrim (required) + enabled (required)
	bazObj := test.CreateObjectDefinition(
		test.Pkg2020,
		"Baz",
		astmodel.NewPropertyDefinition("Discrim", "discrim", astmodel.NewOptionalType(bazDiscrim.Name())).MakeRequired(),
		astmodel.NewPropertyDefinition("Enabled", "enabled", astmodel.NewOptionalType(astmodel.BoolType)).MakeRequired(),
	)

	// Properties type with oneOf referencing Foo, Bar, Baz
	propertiesObj := test.CreateObjectDefinition(
		test.Pkg2020,
		"Properties",
		astmodel.NewPropertyDefinition("Bar", "bar", astmodel.NewOptionalType(barObj.Name())),
		astmodel.NewPropertyDefinition("Baz", "baz", astmodel.NewOptionalType(bazObj.Name())),
		astmodel.NewPropertyDefinition("Foo", "foo", astmodel.NewOptionalType(fooObj.Name())),
	)

	// Apply OneOf flag to properties
	var err error
	propertiesObj, err = propertiesObj.ApplyObjectTransformation(
		func(o *astmodel.ObjectType) (astmodel.Type, error) {
			return astmodel.OneOfFlag.ApplyTo(o), nil
		},
	)
	g.Expect(err).NotTo(HaveOccurred())

	// Create spec with properties reference
	propertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		astmodel.NewOptionalType(propertiesObj.Name()),
	)
	spec := test.CreateSpec(test.Pkg2020, "FakeResource", test.NameProperty, propertiesProp)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(
		resource, status, spec, propertiesObj, fooObj, barObj, bazObj,
		fooDiscrim, barDiscrim, bazDiscrim, test.Pkg2020APIVersion,
	)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_IDResourceReference tests that properties containing ARM resource ID strings
// are correctly transformed into ResourceReference types.
func TestCreateARMTypes_IDResourceReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// FakeResourceProperties with various resource reference patterns
	idProp := astmodel.NewPropertyDefinition("Id", "id", astmodel.OptionalStringType).
		WithDescription("A string of the form /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}")
	subnetIdProp := astmodel.NewPropertyDefinition("SubnetId", "subnetId", astmodel.OptionalStringType).
		WithDescription("A string of the form /SUBSCRIPTIONS/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}")
	nsgIdsProp := astmodel.NewPropertyDefinition(
		"NsgIds",
		"nsgIds",
		astmodel.NewArrayType(astmodel.StringType),
	).WithDescription("A collection of NSG IDs of the form /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/networkSecurityGroups/{nsgName}")
	nsgMapProp := astmodel.NewPropertyDefinition(
		"NsgMap",
		"nsgMap",
		astmodel.NewMapType(astmodel.StringType, astmodel.StringType),
	).WithDescription("A map of NSG IDs of the form /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/networkSecurityGroups/{nsgName}")

	specProperties := test.CreateObjectDefinition(
		test.Pkg2020, "FakeResourceProperties", idProp, subnetIdProp, nsgIdsProp, nsgMapProp,
	)

	propertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		astmodel.NewOptionalType(specProperties.Name()),
	)
	spec := test.CreateSpec(test.Pkg2020, "FakeResource", test.NameProperty, propertiesProp)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	for _, prop := range []*astmodel.PropertyDefinition{idProp, subnetIdProp, nsgIdsProp, nsgMapProp} {
		g.Expect(
			omc.ModifyProperty(
				specProperties.Name(),
				prop.PropertyName(),
				func(pc *config.PropertyConfiguration) error {
					pc.ReferenceType.Set(config.ReferenceTypeARM)
					return nil
				},
			),
		).To(Succeed())
	}

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	state, err := RunTestPipeline(
		NewState(defs),
		ApplyCrossResourceReferencesFromConfig(configuration, logr.Discard()),
		TransformCrossResourceReferences(configuration, idFactory),
		CreateARMTypes(omc, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, omc),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_RequiredAndOptionalResourceReferences tests that ARM resource references with
// different requirement levels (required vs optional) generate correct types and conversions.
func TestCreateARMTypes_RequiredAndOptionalResourceReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	requiredVNetProp := astmodel.NewPropertyDefinition("RequiredVNet", "requiredVNet", astmodel.StringType).
		MakeTypeOptional().
		MakeRequired().
		WithDescription("A string of the form /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}")
	optionalVNetProp := astmodel.NewPropertyDefinition("OptionalVNet", "optionalVNet", astmodel.OptionalStringType).
		WithDescription("A string of the form /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}")

	specProperties := test.CreateObjectDefinition(
		test.Pkg2020, "FakeResourceProperties", requiredVNetProp, optionalVNetProp,
	)

	propertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		astmodel.NewOptionalType(specProperties.Name()),
	)
	spec := test.CreateSpec(test.Pkg2020, "FakeResource", test.NameProperty, propertiesProp)
	status := test.CreateStatus(test.Pkg2020, "FakeResource")
	resource := test.CreateARMResource(test.Pkg2020, "FakeResource", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	for _, prop := range []*astmodel.PropertyDefinition{requiredVNetProp, optionalVNetProp} {
		g.Expect(
			omc.ModifyProperty(
				specProperties.Name(),
				prop.PropertyName(),
				func(pc *config.PropertyConfiguration) error {
					pc.ReferenceType.Set(config.ReferenceTypeARM)
					return nil
				},
			),
		).To(Succeed())
	}

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	state, err := RunTestPipeline(
		NewState(defs),
		ApplyCrossResourceReferencesFromConfig(configuration, logr.Discard()),
		TransformCrossResourceReferences(configuration, idFactory),
		CreateARMTypes(omc, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, omc),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

// TestCreateARMTypes_DependentResourceAndOwnership tests that an ARM resource graph with
// hierarchical ownership (A owns B, B owns C and D) generates correct types for each resource.
func TestCreateARMTypes_DependentResourceAndOwnership(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Resource D (child of B, no children of its own)
	specD := test.CreateSpec(test.Pkg2020, "D", test.NameProperty)
	statusD := test.CreateStatus(test.Pkg2020, "D")
	resourceD := test.CreateARMResource(test.Pkg2020, "D", specD, statusD, test.Pkg2020APIVersion)

	// Resource C (child of B, no children of its own)
	specC := test.CreateSpec(test.Pkg2020, "C", test.NameProperty)
	statusC := test.CreateStatus(test.Pkg2020, "C")
	resourceC := test.CreateARMResource(test.Pkg2020, "C", specC, statusC, test.Pkg2020APIVersion)

	// Resource B (child of A, parent of C and D)
	specB := test.CreateSpec(test.Pkg2020, "B", test.NameProperty)
	statusB := test.CreateStatus(test.Pkg2020, "B")
	resourceB := test.CreateARMResource(test.Pkg2020, "B", specB, statusB, test.Pkg2020APIVersion)

	// Resource A (top-level, parent of B)
	specA := test.CreateSpec(test.Pkg2020, "A", test.NameProperty)
	statusA := test.CreateStatus(test.Pkg2020, "A")
	resourceA := test.CreateARMResource(test.Pkg2020, "A", specA, statusA, test.Pkg2020APIVersion)

	// Set up ownership: B is owned by A
	resourceBRT, _ := astmodel.AsResourceType(resourceB.Type())
	resourceB = resourceB.WithType(resourceBRT.WithOwner(resourceA.Name()))

	// C is owned by B
	resourceCRT, _ := astmodel.AsResourceType(resourceC.Type())
	resourceC = resourceC.WithType(resourceCRT.WithOwner(resourceB.Name()))

	// D is owned by B
	resourceDRT, _ := astmodel.AsResourceType(resourceD.Type())
	resourceD = resourceD.WithType(resourceDRT.WithOwner(resourceB.Name()))

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(
		resourceA, specA, statusA,
		resourceB, specB, statusB,
		resourceC, specC, statusC,
		resourceD, specD, statusD,
		test.Pkg2020APIVersion,
	)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewObjectModelConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(cfg, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, cfg),
		SimplifyDefinitions(),
		StripUnreferencedTypeDefinitions(),
	)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}
