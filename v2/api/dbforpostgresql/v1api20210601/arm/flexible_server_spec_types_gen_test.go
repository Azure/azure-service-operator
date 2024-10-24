// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_Backup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup, BackupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup runs a test to see if a specific instance of Backup round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup(subject Backup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup
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

// Generator of Backup instances for property testing - lazily instantiated by BackupGenerator()
var backupGenerator gopter.Gen

// BackupGenerator returns a generator of Backup instances for property testing.
func BackupGenerator() gopter.Gen {
	if backupGenerator != nil {
		return backupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup(generators)
	backupGenerator = gen.Struct(reflect.TypeOf(Backup{}), generators)

	return backupGenerator
}

// AddIndependentPropertyGeneratorsForBackup is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(Backup_GeoRedundantBackup_Disabled, Backup_GeoRedundantBackup_Enabled))
}

func Test_FlexibleServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer_Spec, FlexibleServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer_Spec runs a test to see if a specific instance of FlexibleServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer_Spec(subject FlexibleServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer_Spec
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

// Generator of FlexibleServer_Spec instances for property testing - lazily instantiated by
// FlexibleServer_SpecGenerator()
var flexibleServer_SpecGenerator gopter.Gen

// FlexibleServer_SpecGenerator returns a generator of FlexibleServer_Spec instances for property testing.
// We first initialize flexibleServer_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServer_SpecGenerator() gopter.Gen {
	if flexibleServer_SpecGenerator != nil {
		return flexibleServer_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_Spec(generators)
	flexibleServer_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_Spec(generators)
	AddRelatedPropertyGeneratorsForFlexibleServer_Spec(generators)
	flexibleServer_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_Spec{}), generators)

	return flexibleServer_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServer_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServer_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_HighAvailability_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability, HighAvailabilityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability runs a test to see if a specific instance of HighAvailability round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability(subject HighAvailability) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability
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

// Generator of HighAvailability instances for property testing - lazily instantiated by HighAvailabilityGenerator()
var highAvailabilityGenerator gopter.Gen

// HighAvailabilityGenerator returns a generator of HighAvailability instances for property testing.
func HighAvailabilityGenerator() gopter.Gen {
	if highAvailabilityGenerator != nil {
		return highAvailabilityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability(generators)
	highAvailabilityGenerator = gen.Struct(reflect.TypeOf(HighAvailability{}), generators)

	return highAvailabilityGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailability_Mode_Disabled, HighAvailability_Mode_ZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindow_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow, MaintenanceWindowGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow runs a test to see if a specific instance of MaintenanceWindow round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow(subject MaintenanceWindow) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow
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

// Generator of MaintenanceWindow instances for property testing - lazily instantiated by MaintenanceWindowGenerator()
var maintenanceWindowGenerator gopter.Gen

// MaintenanceWindowGenerator returns a generator of MaintenanceWindow instances for property testing.
func MaintenanceWindowGenerator() gopter.Gen {
	if maintenanceWindowGenerator != nil {
		return maintenanceWindowGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow(generators)
	maintenanceWindowGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow{}), generators)

	return maintenanceWindowGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork, NetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork runs a test to see if a specific instance of Network round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork(subject Network) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network
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

// Generator of Network instances for property testing - lazily instantiated by NetworkGenerator()
var networkGenerator gopter.Gen

// NetworkGenerator returns a generator of Network instances for property testing.
func NetworkGenerator() gopter.Gen {
	if networkGenerator != nil {
		return networkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetwork(generators)
	networkGenerator = gen.Struct(reflect.TypeOf(Network{}), generators)

	return networkGenerator
}

// AddIndependentPropertyGeneratorsForNetwork is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetwork(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties, ServerPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties runs a test to see if a specific instance of ServerProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties(subject ServerProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties
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

// Generator of ServerProperties instances for property testing - lazily instantiated by ServerPropertiesGenerator()
var serverPropertiesGenerator gopter.Gen

// ServerPropertiesGenerator returns a generator of ServerProperties instances for property testing.
// We first initialize serverPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesGenerator() gopter.Gen {
	if serverPropertiesGenerator != nil {
		return serverPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties(generators)
	serverPropertiesGenerator = gen.Struct(reflect.TypeOf(ServerProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties(generators)
	AddRelatedPropertyGeneratorsForServerProperties(generators)
	serverPropertiesGenerator = gen.Struct(reflect.TypeOf(ServerProperties{}), generators)

	return serverPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_CreateMode_Create,
		ServerProperties_CreateMode_Default,
		ServerProperties_CreateMode_PointInTimeRestore,
		ServerProperties_CreateMode_Update))
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.OneConstOf(
		ServerVersion_11,
		ServerVersion_12,
		ServerVersion_13,
		ServerVersion_14))
}

// AddRelatedPropertyGeneratorsForServerProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowGenerator())
	gens["Network"] = gen.PtrOf(NetworkGenerator())
	gens["Storage"] = gen.PtrOf(StorageGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_Burstable, Sku_Tier_GeneralPurpose, Sku_Tier_MemoryOptimized))
}

func Test_Storage_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage, StorageGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage runs a test to see if a specific instance of Storage round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage(subject Storage) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage
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

// Generator of Storage instances for property testing - lazily instantiated by StorageGenerator()
var storageGenerator gopter.Gen

// StorageGenerator returns a generator of Storage instances for property testing.
func StorageGenerator() gopter.Gen {
	if storageGenerator != nil {
		return storageGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage(generators)
	storageGenerator = gen.Struct(reflect.TypeOf(Storage{}), generators)

	return storageGenerator
}

// AddIndependentPropertyGeneratorsForStorage is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}
