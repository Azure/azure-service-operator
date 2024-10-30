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

func Test_Backup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup_STATUS, Backup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup_STATUS runs a test to see if a specific instance of Backup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup_STATUS(subject Backup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_STATUS
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

// Generator of Backup_STATUS instances for property testing - lazily instantiated by Backup_STATUSGenerator()
var backup_STATUSGenerator gopter.Gen

// Backup_STATUSGenerator returns a generator of Backup_STATUS instances for property testing.
func Backup_STATUSGenerator() gopter.Gen {
	if backup_STATUSGenerator != nil {
		return backup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup_STATUS(generators)
	backup_STATUSGenerator = gen.Struct(reflect.TypeOf(Backup_STATUS{}), generators)

	return backup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForBackup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup_STATUS(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
}

func Test_DataEncryption_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataEncryption_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataEncryption_STATUS, DataEncryption_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataEncryption_STATUS runs a test to see if a specific instance of DataEncryption_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDataEncryption_STATUS(subject DataEncryption_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataEncryption_STATUS
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

// Generator of DataEncryption_STATUS instances for property testing - lazily instantiated by
// DataEncryption_STATUSGenerator()
var dataEncryption_STATUSGenerator gopter.Gen

// DataEncryption_STATUSGenerator returns a generator of DataEncryption_STATUS instances for property testing.
func DataEncryption_STATUSGenerator() gopter.Gen {
	if dataEncryption_STATUSGenerator != nil {
		return dataEncryption_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataEncryption_STATUS(generators)
	dataEncryption_STATUSGenerator = gen.Struct(reflect.TypeOf(DataEncryption_STATUS{}), generators)

	return dataEncryption_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDataEncryption_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataEncryption_STATUS(gens map[string]gopter.Gen) {
	gens["GeoBackupKeyURI"] = gen.PtrOf(gen.AlphaString())
	gens["GeoBackupUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryKeyURI"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(DataEncryption_Type_STATUS_AzureKeyVault, DataEncryption_Type_STATUS_SystemManaged))
}

func Test_FlexibleServer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer_STATUS, FlexibleServer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer_STATUS runs a test to see if a specific instance of FlexibleServer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer_STATUS(subject FlexibleServer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer_STATUS
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

// Generator of FlexibleServer_STATUS instances for property testing - lazily instantiated by
// FlexibleServer_STATUSGenerator()
var flexibleServer_STATUSGenerator gopter.Gen

// FlexibleServer_STATUSGenerator returns a generator of FlexibleServer_STATUS instances for property testing.
// We first initialize flexibleServer_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServer_STATUSGenerator() gopter.Gen {
	if flexibleServer_STATUSGenerator != nil {
		return flexibleServer_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_STATUS(generators)
	flexibleServer_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServer_STATUS(generators)
	flexibleServer_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_STATUS{}), generators)

	return flexibleServer_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServer_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServer_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer_STATUS(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSGenerator())
	gens["Properties"] = gen.PtrOf(ServerProperties_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_HighAvailability_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability_STATUS, HighAvailability_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability_STATUS runs a test to see if a specific instance of HighAvailability_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability_STATUS(subject HighAvailability_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_STATUS
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

// Generator of HighAvailability_STATUS instances for property testing - lazily instantiated by
// HighAvailability_STATUSGenerator()
var highAvailability_STATUSGenerator gopter.Gen

// HighAvailability_STATUSGenerator returns a generator of HighAvailability_STATUS instances for property testing.
func HighAvailability_STATUSGenerator() gopter.Gen {
	if highAvailability_STATUSGenerator != nil {
		return highAvailability_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability_STATUS(generators)
	highAvailability_STATUSGenerator = gen.Struct(reflect.TypeOf(HighAvailability_STATUS{}), generators)

	return highAvailability_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability_STATUS(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailability_Mode_STATUS_Disabled, HighAvailability_Mode_STATUS_SameZone, HighAvailability_Mode_STATUS_ZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		HighAvailability_State_STATUS_CreatingStandby,
		HighAvailability_State_STATUS_FailingOver,
		HighAvailability_State_STATUS_Healthy,
		HighAvailability_State_STATUS_NotEnabled,
		HighAvailability_State_STATUS_RemovingStandby))
}

func Test_Identity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS, Identity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS runs a test to see if a specific instance of Identity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS(subject Identity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS
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

// Generator of Identity_STATUS instances for property testing - lazily instantiated by Identity_STATUSGenerator()
var identity_STATUSGenerator gopter.Gen

// Identity_STATUSGenerator returns a generator of Identity_STATUS instances for property testing.
func Identity_STATUSGenerator() gopter.Gen {
	if identity_STATUSGenerator != nil {
		return identity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS(generators)
	identity_STATUSGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS{}), generators)

	return identity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_Type_STATUS_UserAssigned))
}

func Test_MaintenanceWindow_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow_STATUS, MaintenanceWindow_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow_STATUS runs a test to see if a specific instance of MaintenanceWindow_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow_STATUS(subject MaintenanceWindow_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_STATUS
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

// Generator of MaintenanceWindow_STATUS instances for property testing - lazily instantiated by
// MaintenanceWindow_STATUSGenerator()
var maintenanceWindow_STATUSGenerator gopter.Gen

// MaintenanceWindow_STATUSGenerator returns a generator of MaintenanceWindow_STATUS instances for property testing.
func MaintenanceWindow_STATUSGenerator() gopter.Gen {
	if maintenanceWindow_STATUSGenerator != nil {
		return maintenanceWindow_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS(generators)
	maintenanceWindow_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_STATUS{}), generators)

	return maintenanceWindow_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork_STATUS, Network_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork_STATUS runs a test to see if a specific instance of Network_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork_STATUS(subject Network_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_STATUS
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

// Generator of Network_STATUS instances for property testing - lazily instantiated by Network_STATUSGenerator()
var network_STATUSGenerator gopter.Gen

// Network_STATUSGenerator returns a generator of Network_STATUS instances for property testing.
func Network_STATUSGenerator() gopter.Gen {
	if network_STATUSGenerator != nil {
		return network_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetwork_STATUS(generators)
	network_STATUSGenerator = gen.Struct(reflect.TypeOf(Network_STATUS{}), generators)

	return network_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNetwork_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetwork_STATUS(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
}

func Test_ServerProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_STATUS, ServerProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_STATUS runs a test to see if a specific instance of ServerProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_STATUS(subject ServerProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_STATUS
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

// Generator of ServerProperties_STATUS instances for property testing - lazily instantiated by
// ServerProperties_STATUSGenerator()
var serverProperties_STATUSGenerator gopter.Gen

// ServerProperties_STATUSGenerator returns a generator of ServerProperties_STATUS instances for property testing.
// We first initialize serverProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_STATUSGenerator() gopter.Gen {
	if serverProperties_STATUSGenerator != nil {
		return serverProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUS(generators)
	serverProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForServerProperties_STATUS(generators)
	serverProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUS{}), generators)

	return serverProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_CreateMode_STATUS_Default,
		ServerProperties_CreateMode_STATUS_GeoRestore,
		ServerProperties_CreateMode_STATUS_PointInTimeRestore,
		ServerProperties_CreateMode_STATUS_Replica))
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicaCapacity"] = gen.PtrOf(gen.Int())
	gens["ReplicationRole"] = gen.PtrOf(gen.OneConstOf(ReplicationRole_STATUS_None, ReplicationRole_STATUS_Replica, ReplicationRole_STATUS_Source))
	gens["RestorePointInTime"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_State_STATUS_Disabled,
		ServerProperties_State_STATUS_Dropping,
		ServerProperties_State_STATUS_Ready,
		ServerProperties_State_STATUS_Starting,
		ServerProperties_State_STATUS_Stopped,
		ServerProperties_State_STATUS_Stopping,
		ServerProperties_State_STATUS_Updating))
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerVersion_STATUS_57, ServerVersion_STATUS_8021))
}

// AddRelatedPropertyGeneratorsForServerProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(Backup_STATUSGenerator())
	gens["DataEncryption"] = gen.PtrOf(DataEncryption_STATUSGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailability_STATUSGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindow_STATUSGenerator())
	gens["Network"] = gen.PtrOf(Network_STATUSGenerator())
	gens["Storage"] = gen.PtrOf(Storage_STATUSGenerator())
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_STATUS_Burstable, Sku_Tier_STATUS_GeneralPurpose, Sku_Tier_STATUS_MemoryOptimized))
}

func Test_Storage_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage_STATUS, Storage_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage_STATUS runs a test to see if a specific instance of Storage_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage_STATUS(subject Storage_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_STATUS
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

// Generator of Storage_STATUS instances for property testing - lazily instantiated by Storage_STATUSGenerator()
var storage_STATUSGenerator gopter.Gen

// Storage_STATUSGenerator returns a generator of Storage_STATUS instances for property testing.
func Storage_STATUSGenerator() gopter.Gen {
	if storage_STATUSGenerator != nil {
		return storage_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage_STATUS(generators)
	storage_STATUSGenerator = gen.Struct(reflect.TypeOf(Storage_STATUS{}), generators)

	return storage_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorage_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage_STATUS(gens map[string]gopter.Gen) {
	gens["AutoGrow"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
	gens["Iops"] = gen.PtrOf(gen.Int())
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
	gens["StorageSku"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}
