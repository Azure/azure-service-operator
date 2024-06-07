// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ContainerProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerProperties_ARM, ContainerProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerProperties_ARM runs a test to see if a specific instance of ContainerProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerProperties_ARM(subject ContainerProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ContainerProperties_ARM instances for property testing - lazily instantiated by
// ContainerProperties_ARMGenerator()
var containerProperties_ARMGenerator gopter.Gen

// ContainerProperties_ARMGenerator returns a generator of ContainerProperties_ARM instances for property testing.
// We first initialize containerProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerProperties_ARMGenerator() gopter.Gen {
	if containerProperties_ARMGenerator != nil {
		return containerProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_ARM(generators)
	containerProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForContainerProperties_ARM(generators)
	containerProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_ARM{}), generators)

	return containerProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerProperties_ARM(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3AllSquash"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3RootSquash"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_PublicAccess_Blob, ContainerProperties_PublicAccess_Container, ContainerProperties_PublicAccess_None))
}

// AddRelatedPropertyGeneratorsForContainerProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerProperties_ARM(gens map[string]gopter.Gen) {
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_ARMGenerator())
}

func Test_ImmutableStorageWithVersioning_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_ARM, ImmutableStorageWithVersioning_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_ARM runs a test to see if a specific instance of ImmutableStorageWithVersioning_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_ARM(subject ImmutableStorageWithVersioning_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ImmutableStorageWithVersioning_ARM instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioning_ARMGenerator()
var immutableStorageWithVersioning_ARMGenerator gopter.Gen

// ImmutableStorageWithVersioning_ARMGenerator returns a generator of ImmutableStorageWithVersioning_ARM instances for property testing.
func ImmutableStorageWithVersioning_ARMGenerator() gopter.Gen {
	if immutableStorageWithVersioning_ARMGenerator != nil {
		return immutableStorageWithVersioning_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_ARM(generators)
	immutableStorageWithVersioning_ARMGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_ARM{}), generators)

	return immutableStorageWithVersioning_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_StorageAccounts_BlobServices_Container_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_BlobServices_Container_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec_ARM, StorageAccounts_BlobServices_Container_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec_ARM runs a test to see if a specific instance of StorageAccounts_BlobServices_Container_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec_ARM(subject StorageAccounts_BlobServices_Container_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_BlobServices_Container_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccounts_BlobServices_Container_Spec_ARM instances for property testing - lazily instantiated by
// StorageAccounts_BlobServices_Container_Spec_ARMGenerator()
var storageAccounts_BlobServices_Container_Spec_ARMGenerator gopter.Gen

// StorageAccounts_BlobServices_Container_Spec_ARMGenerator returns a generator of StorageAccounts_BlobServices_Container_Spec_ARM instances for property testing.
// We first initialize storageAccounts_BlobServices_Container_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_BlobServices_Container_Spec_ARMGenerator() gopter.Gen {
	if storageAccounts_BlobServices_Container_Spec_ARMGenerator != nil {
		return storageAccounts_BlobServices_Container_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM(generators)
	storageAccounts_BlobServices_Container_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM(generators)
	storageAccounts_BlobServices_Container_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_Spec_ARM{}), generators)

	return storageAccounts_BlobServices_Container_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerProperties_ARMGenerator())
}
