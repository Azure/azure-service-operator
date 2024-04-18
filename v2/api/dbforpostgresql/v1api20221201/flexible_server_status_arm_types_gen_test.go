// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221201

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

func Test_FlexibleServer_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer_STATUS_ARM, FlexibleServer_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer_STATUS_ARM runs a test to see if a specific instance of FlexibleServer_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer_STATUS_ARM(subject FlexibleServer_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer_STATUS_ARM
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

// Generator of FlexibleServer_STATUS_ARM instances for property testing - lazily instantiated by
// FlexibleServer_STATUS_ARMGenerator()
var flexibleServer_STATUS_ARMGenerator gopter.Gen

// FlexibleServer_STATUS_ARMGenerator returns a generator of FlexibleServer_STATUS_ARM instances for property testing.
// We first initialize flexibleServer_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServer_STATUS_ARMGenerator() gopter.Gen {
	if flexibleServer_STATUS_ARMGenerator != nil {
		return flexibleServer_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_STATUS_ARM(generators)
	flexibleServer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServer_STATUS_ARM(generators)
	flexibleServer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_STATUS_ARM{}), generators)

	return flexibleServer_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServer_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServer_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ServerProperties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_ServerProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_STATUS_ARM, ServerProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_STATUS_ARM runs a test to see if a specific instance of ServerProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_STATUS_ARM(subject ServerProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_STATUS_ARM
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

// Generator of ServerProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ServerProperties_STATUS_ARMGenerator()
var serverProperties_STATUS_ARMGenerator gopter.Gen

// ServerProperties_STATUS_ARMGenerator returns a generator of ServerProperties_STATUS_ARM instances for property testing.
// We first initialize serverProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_STATUS_ARMGenerator() gopter.Gen {
	if serverProperties_STATUS_ARMGenerator != nil {
		return serverProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUS_ARM(generators)
	serverProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServerProperties_STATUS_ARM(generators)
	serverProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUS_ARM{}), generators)

	return serverProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_CreateMode_STATUS_Create,
		ServerProperties_CreateMode_STATUS_Default,
		ServerProperties_CreateMode_STATUS_GeoRestore,
		ServerProperties_CreateMode_STATUS_PointInTimeRestore,
		ServerProperties_CreateMode_STATUS_Replica,
		ServerProperties_CreateMode_STATUS_Update))
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["MinorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicaCapacity"] = gen.PtrOf(gen.Int())
	gens["ReplicationRole"] = gen.PtrOf(gen.OneConstOf(
		ReplicationRole_STATUS_AsyncReplica,
		ReplicationRole_STATUS_GeoAsyncReplica,
		ReplicationRole_STATUS_None,
		ReplicationRole_STATUS_Primary))
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_State_STATUS_Disabled,
		ServerProperties_State_STATUS_Dropping,
		ServerProperties_State_STATUS_Ready,
		ServerProperties_State_STATUS_Starting,
		ServerProperties_State_STATUS_Stopped,
		ServerProperties_State_STATUS_Stopping,
		ServerProperties_State_STATUS_Updating))
	gens["Version"] = gen.PtrOf(gen.OneConstOf(
		ServerVersion_STATUS_11,
		ServerVersion_STATUS_12,
		ServerVersion_STATUS_13,
		ServerVersion_STATUS_14))
}

// AddRelatedPropertyGeneratorsForServerProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AuthConfig"] = gen.PtrOf(AuthConfig_STATUS_ARMGenerator())
	gens["Backup"] = gen.PtrOf(Backup_STATUS_ARMGenerator())
	gens["DataEncryption"] = gen.PtrOf(DataEncryption_STATUS_ARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailability_STATUS_ARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindow_STATUS_ARMGenerator())
	gens["Network"] = gen.PtrOf(Network_STATUS_ARMGenerator())
	gens["Storage"] = gen.PtrOf(Storage_STATUS_ARMGenerator())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
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

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_STATUS_Burstable, Sku_Tier_STATUS_GeneralPurpose, Sku_Tier_STATUS_MemoryOptimized))
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
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

func Test_UserAssignedIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM, UserAssignedIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM runs a test to see if a specific instance of UserAssignedIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM(subject UserAssignedIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS_ARM
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

// Generator of UserAssignedIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUS_ARMGenerator()
var userAssignedIdentity_STATUS_ARMGenerator gopter.Gen

// UserAssignedIdentity_STATUS_ARMGenerator returns a generator of UserAssignedIdentity_STATUS_ARM instances for property testing.
// We first initialize userAssignedIdentity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentity_STATUS_ARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUS_ARMGenerator != nil {
		return userAssignedIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	userAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	userAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS_ARM{}), generators)

	return userAssignedIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(UserAssignedIdentity_Type_STATUS_None, UserAssignedIdentity_Type_STATUS_UserAssigned))
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserIdentity_STATUS_ARMGenerator())
}

func Test_AuthConfig_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthConfig_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthConfig_STATUS_ARM, AuthConfig_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthConfig_STATUS_ARM runs a test to see if a specific instance of AuthConfig_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthConfig_STATUS_ARM(subject AuthConfig_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthConfig_STATUS_ARM
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

// Generator of AuthConfig_STATUS_ARM instances for property testing - lazily instantiated by
// AuthConfig_STATUS_ARMGenerator()
var authConfig_STATUS_ARMGenerator gopter.Gen

// AuthConfig_STATUS_ARMGenerator returns a generator of AuthConfig_STATUS_ARM instances for property testing.
func AuthConfig_STATUS_ARMGenerator() gopter.Gen {
	if authConfig_STATUS_ARMGenerator != nil {
		return authConfig_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthConfig_STATUS_ARM(generators)
	authConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthConfig_STATUS_ARM{}), generators)

	return authConfig_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthConfig_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthConfig_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ActiveDirectoryAuth"] = gen.PtrOf(gen.OneConstOf(AuthConfig_ActiveDirectoryAuth_STATUS_Disabled, AuthConfig_ActiveDirectoryAuth_STATUS_Enabled))
	gens["PasswordAuth"] = gen.PtrOf(gen.OneConstOf(AuthConfig_PasswordAuth_STATUS_Disabled, AuthConfig_PasswordAuth_STATUS_Enabled))
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Backup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup_STATUS_ARM, Backup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup_STATUS_ARM runs a test to see if a specific instance of Backup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup_STATUS_ARM(subject Backup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_STATUS_ARM
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

// Generator of Backup_STATUS_ARM instances for property testing - lazily instantiated by Backup_STATUS_ARMGenerator()
var backup_STATUS_ARMGenerator gopter.Gen

// Backup_STATUS_ARMGenerator returns a generator of Backup_STATUS_ARM instances for property testing.
func Backup_STATUS_ARMGenerator() gopter.Gen {
	if backup_STATUS_ARMGenerator != nil {
		return backup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup_STATUS_ARM(generators)
	backup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Backup_STATUS_ARM{}), generators)

	return backup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(Backup_GeoRedundantBackup_STATUS_Disabled, Backup_GeoRedundantBackup_STATUS_Enabled))
}

func Test_DataEncryption_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataEncryption_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataEncryption_STATUS_ARM, DataEncryption_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataEncryption_STATUS_ARM runs a test to see if a specific instance of DataEncryption_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataEncryption_STATUS_ARM(subject DataEncryption_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataEncryption_STATUS_ARM
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

// Generator of DataEncryption_STATUS_ARM instances for property testing - lazily instantiated by
// DataEncryption_STATUS_ARMGenerator()
var dataEncryption_STATUS_ARMGenerator gopter.Gen

// DataEncryption_STATUS_ARMGenerator returns a generator of DataEncryption_STATUS_ARM instances for property testing.
func DataEncryption_STATUS_ARMGenerator() gopter.Gen {
	if dataEncryption_STATUS_ARMGenerator != nil {
		return dataEncryption_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataEncryption_STATUS_ARM(generators)
	dataEncryption_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DataEncryption_STATUS_ARM{}), generators)

	return dataEncryption_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDataEncryption_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataEncryption_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrimaryKeyURI"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(DataEncryption_Type_STATUS_AzureKeyVault, DataEncryption_Type_STATUS_SystemManaged))
}

func Test_HighAvailability_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability_STATUS_ARM, HighAvailability_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability_STATUS_ARM runs a test to see if a specific instance of HighAvailability_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability_STATUS_ARM(subject HighAvailability_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_STATUS_ARM
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

// Generator of HighAvailability_STATUS_ARM instances for property testing - lazily instantiated by
// HighAvailability_STATUS_ARMGenerator()
var highAvailability_STATUS_ARMGenerator gopter.Gen

// HighAvailability_STATUS_ARMGenerator returns a generator of HighAvailability_STATUS_ARM instances for property testing.
func HighAvailability_STATUS_ARMGenerator() gopter.Gen {
	if highAvailability_STATUS_ARMGenerator != nil {
		return highAvailability_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability_STATUS_ARM(generators)
	highAvailability_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(HighAvailability_STATUS_ARM{}), generators)

	return highAvailability_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailability_Mode_STATUS_Disabled, HighAvailability_Mode_STATUS_SameZone, HighAvailability_Mode_STATUS_ZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		HighAvailability_State_STATUS_CreatingStandby,
		HighAvailability_State_STATUS_FailingOver,
		HighAvailability_State_STATUS_Healthy,
		HighAvailability_State_STATUS_NotEnabled,
		HighAvailability_State_STATUS_RemovingStandby,
		HighAvailability_State_STATUS_ReplicatingData))
}

func Test_MaintenanceWindow_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow_STATUS_ARM, MaintenanceWindow_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow_STATUS_ARM runs a test to see if a specific instance of MaintenanceWindow_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow_STATUS_ARM(subject MaintenanceWindow_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_STATUS_ARM
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

// Generator of MaintenanceWindow_STATUS_ARM instances for property testing - lazily instantiated by
// MaintenanceWindow_STATUS_ARMGenerator()
var maintenanceWindow_STATUS_ARMGenerator gopter.Gen

// MaintenanceWindow_STATUS_ARMGenerator returns a generator of MaintenanceWindow_STATUS_ARM instances for property testing.
func MaintenanceWindow_STATUS_ARMGenerator() gopter.Gen {
	if maintenanceWindow_STATUS_ARMGenerator != nil {
		return maintenanceWindow_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS_ARM(generators)
	maintenanceWindow_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_STATUS_ARM{}), generators)

	return maintenanceWindow_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork_STATUS_ARM, Network_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork_STATUS_ARM runs a test to see if a specific instance of Network_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork_STATUS_ARM(subject Network_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_STATUS_ARM
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

// Generator of Network_STATUS_ARM instances for property testing - lazily instantiated by Network_STATUS_ARMGenerator()
var network_STATUS_ARMGenerator gopter.Gen

// Network_STATUS_ARMGenerator returns a generator of Network_STATUS_ARM instances for property testing.
func Network_STATUS_ARMGenerator() gopter.Gen {
	if network_STATUS_ARMGenerator != nil {
		return network_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetwork_STATUS_ARM(generators)
	network_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Network_STATUS_ARM{}), generators)

	return network_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetwork_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetwork_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(Network_PublicNetworkAccess_STATUS_Disabled, Network_PublicNetworkAccess_STATUS_Enabled))
}

func Test_Storage_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage_STATUS_ARM, Storage_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage_STATUS_ARM runs a test to see if a specific instance of Storage_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage_STATUS_ARM(subject Storage_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_STATUS_ARM
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

// Generator of Storage_STATUS_ARM instances for property testing - lazily instantiated by Storage_STATUS_ARMGenerator()
var storage_STATUS_ARMGenerator gopter.Gen

// Storage_STATUS_ARMGenerator returns a generator of Storage_STATUS_ARM instances for property testing.
func Storage_STATUS_ARMGenerator() gopter.Gen {
	if storage_STATUS_ARMGenerator != nil {
		return storage_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage_STATUS_ARM(generators)
	storage_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Storage_STATUS_ARM{}), generators)

	return storage_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorage_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}

func Test_UserIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserIdentity_STATUS_ARM, UserIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserIdentity_STATUS_ARM runs a test to see if a specific instance of UserIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserIdentity_STATUS_ARM(subject UserIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserIdentity_STATUS_ARM
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

// Generator of UserIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// UserIdentity_STATUS_ARMGenerator()
var userIdentity_STATUS_ARMGenerator gopter.Gen

// UserIdentity_STATUS_ARMGenerator returns a generator of UserIdentity_STATUS_ARM instances for property testing.
func UserIdentity_STATUS_ARMGenerator() gopter.Gen {
	if userIdentity_STATUS_ARMGenerator != nil {
		return userIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserIdentity_STATUS_ARM(generators)
	userIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserIdentity_STATUS_ARM{}), generators)

	return userIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}
