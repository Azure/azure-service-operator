// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

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

func Test_Snapshot_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotStatusARM, SnapshotStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotStatusARM runs a test to see if a specific instance of Snapshot_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotStatusARM(subject Snapshot_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_StatusARM
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

// Generator of Snapshot_StatusARM instances for property testing - lazily instantiated by SnapshotStatusARMGenerator()
var snapshotStatusARMGenerator gopter.Gen

// SnapshotStatusARMGenerator returns a generator of Snapshot_StatusARM instances for property testing.
// We first initialize snapshotStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotStatusARMGenerator() gopter.Gen {
	if snapshotStatusARMGenerator != nil {
		return snapshotStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotStatusARM(generators)
	snapshotStatusARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotStatusARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotStatusARM(generators)
	snapshotStatusARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_StatusARM{}), generators)

	return snapshotStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotStatusARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(SnapshotPropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuStatusARMGenerator())
}

func Test_SnapshotProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotPropertiesStatusARM, SnapshotPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotPropertiesStatusARM runs a test to see if a specific instance of SnapshotProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotPropertiesStatusARM(subject SnapshotProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotProperties_StatusARM
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

// Generator of SnapshotProperties_StatusARM instances for property testing - lazily instantiated by
// SnapshotPropertiesStatusARMGenerator()
var snapshotPropertiesStatusARMGenerator gopter.Gen

// SnapshotPropertiesStatusARMGenerator returns a generator of SnapshotProperties_StatusARM instances for property testing.
// We first initialize snapshotPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotPropertiesStatusARMGenerator() gopter.Gen {
	if snapshotPropertiesStatusARMGenerator != nil {
		return snapshotPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotPropertiesStatusARM(generators)
	snapshotPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotPropertiesStatusARM(generators)
	snapshotPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_StatusARM{}), generators)

	return snapshotPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_StatusActiveSAS,
		DiskState_StatusActiveUpload,
		DiskState_StatusAttached,
		DiskState_StatusReadyToUpload,
		DiskState_StatusReserved,
		DiskState_StatusUnattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesStatusHyperVGenerationV1, SnapshotPropertiesStatusHyperVGenerationV2))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_StatusAllowAll, NetworkAccessPolicy_StatusAllowPrivate, NetworkAccessPolicy_StatusDenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesStatusOsTypeLinux, SnapshotPropertiesStatusOsTypeWindows))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataStatusARMGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionStatusARMGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionStatusARMGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanStatusARMGenerator())
}

func Test_SnapshotSku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSkuStatusARM, SnapshotSkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSkuStatusARM runs a test to see if a specific instance of SnapshotSku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSkuStatusARM(subject SnapshotSku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_StatusARM
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

// Generator of SnapshotSku_StatusARM instances for property testing - lazily instantiated by
// SnapshotSkuStatusARMGenerator()
var snapshotSkuStatusARMGenerator gopter.Gen

// SnapshotSkuStatusARMGenerator returns a generator of SnapshotSku_StatusARM instances for property testing.
func SnapshotSkuStatusARMGenerator() gopter.Gen {
	if snapshotSkuStatusARMGenerator != nil {
		return snapshotSkuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSkuStatusARM(generators)
	snapshotSkuStatusARMGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_StatusARM{}), generators)

	return snapshotSkuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSkuStatusNamePremiumLRS, SnapshotSkuStatusNameStandardLRS, SnapshotSkuStatusNameStandardZRS))
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}
