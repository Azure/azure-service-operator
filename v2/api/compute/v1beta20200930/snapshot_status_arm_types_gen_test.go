// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930

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

func Test_Snapshot_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_STATUSARM, Snapshot_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_STATUSARM runs a test to see if a specific instance of Snapshot_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_STATUSARM(subject Snapshot_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_STATUSARM
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

// Generator of Snapshot_STATUSARM instances for property testing - lazily instantiated by Snapshot_STATUSARMGenerator()
var snapshot_STATUSARMGenerator gopter.Gen

// Snapshot_STATUSARMGenerator returns a generator of Snapshot_STATUSARM instances for property testing.
// We first initialize snapshot_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_STATUSARMGenerator() gopter.Gen {
	if snapshot_STATUSARMGenerator != nil {
		return snapshot_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUSARM(generators)
	snapshot_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSnapshot_STATUSARM(generators)
	snapshot_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUSARM{}), generators)

	return snapshot_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_STATUSARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(SnapshotProperties_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSku_STATUSARMGenerator())
}

func Test_SnapshotProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotProperties_STATUSARM, SnapshotProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotProperties_STATUSARM runs a test to see if a specific instance of SnapshotProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotProperties_STATUSARM(subject SnapshotProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotProperties_STATUSARM
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

// Generator of SnapshotProperties_STATUSARM instances for property testing - lazily instantiated by
// SnapshotProperties_STATUSARMGenerator()
var snapshotProperties_STATUSARMGenerator gopter.Gen

// SnapshotProperties_STATUSARMGenerator returns a generator of SnapshotProperties_STATUSARM instances for property testing.
// We first initialize snapshotProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotProperties_STATUSARMGenerator() gopter.Gen {
	if snapshotProperties_STATUSARMGenerator != nil {
		return snapshotProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotProperties_STATUSARM(generators)
	snapshotProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotProperties_STATUSARM(generators)
	snapshotProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SnapshotProperties_STATUSARM{}), generators)

	return snapshotProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_ActiveSAS_STATUS,
		DiskState_ActiveUpload_STATUS,
		DiskState_Attached_STATUS,
		DiskState_ReadyToUpload_STATUS,
		DiskState_Reserved_STATUS,
		DiskState_Unattached_STATUS))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_HyperVGeneration_V1_STATUS, SnapshotProperties_HyperVGeneration_V2_STATUS))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_AllowAll_STATUS, NetworkAccessPolicy_AllowPrivate_STATUS, NetworkAccessPolicy_DenyAll_STATUS))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_OsType_Linux_STATUS, SnapshotProperties_OsType_Windows_STATUS))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationData_STATUSARMGenerator())
	gens["Encryption"] = gen.PtrOf(Encryption_STATUSARMGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollection_STATUSARMGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlan_STATUSARMGenerator())
}

func Test_SnapshotSku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku_STATUSARM, SnapshotSku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku_STATUSARM runs a test to see if a specific instance of SnapshotSku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku_STATUSARM(subject SnapshotSku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_STATUSARM
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

// Generator of SnapshotSku_STATUSARM instances for property testing - lazily instantiated by
// SnapshotSku_STATUSARMGenerator()
var snapshotSku_STATUSARMGenerator gopter.Gen

// SnapshotSku_STATUSARMGenerator returns a generator of SnapshotSku_STATUSARM instances for property testing.
func SnapshotSku_STATUSARMGenerator() gopter.Gen {
	if snapshotSku_STATUSARMGenerator != nil {
		return snapshotSku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku_STATUSARM(generators)
	snapshotSku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_STATUSARM{}), generators)

	return snapshotSku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSku_Name_Premium_LRS_STATUS, SnapshotSku_Name_Standard_LRS_STATUS, SnapshotSku_Name_Standard_ZRS_STATUS))
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}
