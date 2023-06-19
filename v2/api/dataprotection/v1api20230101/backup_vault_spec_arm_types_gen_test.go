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

func Test_BackupVault_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVault_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVault_Spec_ARM, BackupVault_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVault_Spec_ARM runs a test to see if a specific instance of BackupVault_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVault_Spec_ARM(subject BackupVault_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVault_Spec_ARM
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

// Generator of BackupVault_Spec_ARM instances for property testing - lazily instantiated by
// BackupVault_Spec_ARMGenerator()
var backupVault_Spec_ARMGenerator gopter.Gen

// BackupVault_Spec_ARMGenerator returns a generator of BackupVault_Spec_ARM instances for property testing.
// We first initialize backupVault_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackupVault_Spec_ARMGenerator() gopter.Gen {
	if backupVault_Spec_ARMGenerator != nil {
		return backupVault_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_Spec_ARM(generators)
	backupVault_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(BackupVault_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForBackupVault_Spec_ARM(generators)
	backupVault_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(BackupVault_Spec_ARM{}), generators)

	return backupVault_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackupVault_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupVault_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackupVault_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVault_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(DppIdentityDetails_ARMGenerator())
	gens["Properties"] = gen.PtrOf(BackupVaultSpec_ARMGenerator())
}

func Test_BackupVaultSpec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVaultSpec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVaultSpec_ARM, BackupVaultSpec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVaultSpec_ARM runs a test to see if a specific instance of BackupVaultSpec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVaultSpec_ARM(subject BackupVaultSpec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVaultSpec_ARM
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

// Generator of BackupVaultSpec_ARM instances for property testing - lazily instantiated by
// BackupVaultSpec_ARMGenerator()
var backupVaultSpec_ARMGenerator gopter.Gen

// BackupVaultSpec_ARMGenerator returns a generator of BackupVaultSpec_ARM instances for property testing.
func BackupVaultSpec_ARMGenerator() gopter.Gen {
	if backupVaultSpec_ARMGenerator != nil {
		return backupVaultSpec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBackupVaultSpec_ARM(generators)
	backupVaultSpec_ARMGenerator = gen.Struct(reflect.TypeOf(BackupVaultSpec_ARM{}), generators)

	return backupVaultSpec_ARMGenerator
}

// AddRelatedPropertyGeneratorsForBackupVaultSpec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVaultSpec_ARM(gens map[string]gopter.Gen) {
	gens["FeatureSettings"] = gen.PtrOf(FeatureSettings_ARMGenerator())
	gens["MonitoringSettings"] = gen.PtrOf(MonitoringSettings_ARMGenerator())
	gens["SecuritySettings"] = gen.PtrOf(SecuritySettings_ARMGenerator())
	gens["StorageSettings"] = gen.SliceOf(StorageSetting_ARMGenerator())
}

func Test_DppIdentityDetails_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DppIdentityDetails_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDppIdentityDetails_ARM, DppIdentityDetails_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDppIdentityDetails_ARM runs a test to see if a specific instance of DppIdentityDetails_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDppIdentityDetails_ARM(subject DppIdentityDetails_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DppIdentityDetails_ARM
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

// Generator of DppIdentityDetails_ARM instances for property testing - lazily instantiated by
// DppIdentityDetails_ARMGenerator()
var dppIdentityDetails_ARMGenerator gopter.Gen

// DppIdentityDetails_ARMGenerator returns a generator of DppIdentityDetails_ARM instances for property testing.
func DppIdentityDetails_ARMGenerator() gopter.Gen {
	if dppIdentityDetails_ARMGenerator != nil {
		return dppIdentityDetails_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDppIdentityDetails_ARM(generators)
	dppIdentityDetails_ARMGenerator = gen.Struct(reflect.TypeOf(DppIdentityDetails_ARM{}), generators)

	return dppIdentityDetails_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDppIdentityDetails_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDppIdentityDetails_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_FeatureSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FeatureSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFeatureSettings_ARM, FeatureSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFeatureSettings_ARM runs a test to see if a specific instance of FeatureSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFeatureSettings_ARM(subject FeatureSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FeatureSettings_ARM
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

// Generator of FeatureSettings_ARM instances for property testing - lazily instantiated by
// FeatureSettings_ARMGenerator()
var featureSettings_ARMGenerator gopter.Gen

// FeatureSettings_ARMGenerator returns a generator of FeatureSettings_ARM instances for property testing.
func FeatureSettings_ARMGenerator() gopter.Gen {
	if featureSettings_ARMGenerator != nil {
		return featureSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFeatureSettings_ARM(generators)
	featureSettings_ARMGenerator = gen.Struct(reflect.TypeOf(FeatureSettings_ARM{}), generators)

	return featureSettings_ARMGenerator
}

// AddRelatedPropertyGeneratorsForFeatureSettings_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFeatureSettings_ARM(gens map[string]gopter.Gen) {
	gens["CrossSubscriptionRestoreSettings"] = gen.PtrOf(CrossSubscriptionRestoreSettings_ARMGenerator())
}

func Test_MonitoringSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitoringSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitoringSettings_ARM, MonitoringSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitoringSettings_ARM runs a test to see if a specific instance of MonitoringSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitoringSettings_ARM(subject MonitoringSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitoringSettings_ARM
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

// Generator of MonitoringSettings_ARM instances for property testing - lazily instantiated by
// MonitoringSettings_ARMGenerator()
var monitoringSettings_ARMGenerator gopter.Gen

// MonitoringSettings_ARMGenerator returns a generator of MonitoringSettings_ARM instances for property testing.
func MonitoringSettings_ARMGenerator() gopter.Gen {
	if monitoringSettings_ARMGenerator != nil {
		return monitoringSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMonitoringSettings_ARM(generators)
	monitoringSettings_ARMGenerator = gen.Struct(reflect.TypeOf(MonitoringSettings_ARM{}), generators)

	return monitoringSettings_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMonitoringSettings_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMonitoringSettings_ARM(gens map[string]gopter.Gen) {
	gens["AzureMonitorAlertSettings"] = gen.PtrOf(AzureMonitorAlertSettings_ARMGenerator())
}

func Test_SecuritySettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecuritySettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecuritySettings_ARM, SecuritySettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecuritySettings_ARM runs a test to see if a specific instance of SecuritySettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecuritySettings_ARM(subject SecuritySettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecuritySettings_ARM
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

// Generator of SecuritySettings_ARM instances for property testing - lazily instantiated by
// SecuritySettings_ARMGenerator()
var securitySettings_ARMGenerator gopter.Gen

// SecuritySettings_ARMGenerator returns a generator of SecuritySettings_ARM instances for property testing.
func SecuritySettings_ARMGenerator() gopter.Gen {
	if securitySettings_ARMGenerator != nil {
		return securitySettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecuritySettings_ARM(generators)
	securitySettings_ARMGenerator = gen.Struct(reflect.TypeOf(SecuritySettings_ARM{}), generators)

	return securitySettings_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSecuritySettings_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecuritySettings_ARM(gens map[string]gopter.Gen) {
	gens["ImmutabilitySettings"] = gen.PtrOf(ImmutabilitySettings_ARMGenerator())
	gens["SoftDeleteSettings"] = gen.PtrOf(SoftDeleteSettings_ARMGenerator())
}

func Test_StorageSetting_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageSetting_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageSetting_ARM, StorageSetting_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageSetting_ARM runs a test to see if a specific instance of StorageSetting_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageSetting_ARM(subject StorageSetting_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageSetting_ARM
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

// Generator of StorageSetting_ARM instances for property testing - lazily instantiated by StorageSetting_ARMGenerator()
var storageSetting_ARMGenerator gopter.Gen

// StorageSetting_ARMGenerator returns a generator of StorageSetting_ARM instances for property testing.
func StorageSetting_ARMGenerator() gopter.Gen {
	if storageSetting_ARMGenerator != nil {
		return storageSetting_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageSetting_ARM(generators)
	storageSetting_ARMGenerator = gen.Struct(reflect.TypeOf(StorageSetting_ARM{}), generators)

	return storageSetting_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageSetting_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageSetting_ARM(gens map[string]gopter.Gen) {
	gens["DatastoreType"] = gen.PtrOf(gen.OneConstOf(StorageSetting_DatastoreType_ArchiveStore, StorageSetting_DatastoreType_OperationalStore, StorageSetting_DatastoreType_VaultStore))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(StorageSetting_Type_GeoRedundant, StorageSetting_Type_LocallyRedundant, StorageSetting_Type_ZoneRedundant))
}

func Test_AzureMonitorAlertSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureMonitorAlertSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureMonitorAlertSettings_ARM, AzureMonitorAlertSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureMonitorAlertSettings_ARM runs a test to see if a specific instance of AzureMonitorAlertSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureMonitorAlertSettings_ARM(subject AzureMonitorAlertSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureMonitorAlertSettings_ARM
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

// Generator of AzureMonitorAlertSettings_ARM instances for property testing - lazily instantiated by
// AzureMonitorAlertSettings_ARMGenerator()
var azureMonitorAlertSettings_ARMGenerator gopter.Gen

// AzureMonitorAlertSettings_ARMGenerator returns a generator of AzureMonitorAlertSettings_ARM instances for property testing.
func AzureMonitorAlertSettings_ARMGenerator() gopter.Gen {
	if azureMonitorAlertSettings_ARMGenerator != nil {
		return azureMonitorAlertSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_ARM(generators)
	azureMonitorAlertSettings_ARMGenerator = gen.Struct(reflect.TypeOf(AzureMonitorAlertSettings_ARM{}), generators)

	return azureMonitorAlertSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_ARM(gens map[string]gopter.Gen) {
	gens["AlertsForAllJobFailures"] = gen.PtrOf(gen.OneConstOf(AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled, AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled))
}

func Test_CrossSubscriptionRestoreSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CrossSubscriptionRestoreSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCrossSubscriptionRestoreSettings_ARM, CrossSubscriptionRestoreSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCrossSubscriptionRestoreSettings_ARM runs a test to see if a specific instance of CrossSubscriptionRestoreSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCrossSubscriptionRestoreSettings_ARM(subject CrossSubscriptionRestoreSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CrossSubscriptionRestoreSettings_ARM
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

// Generator of CrossSubscriptionRestoreSettings_ARM instances for property testing - lazily instantiated by
// CrossSubscriptionRestoreSettings_ARMGenerator()
var crossSubscriptionRestoreSettings_ARMGenerator gopter.Gen

// CrossSubscriptionRestoreSettings_ARMGenerator returns a generator of CrossSubscriptionRestoreSettings_ARM instances for property testing.
func CrossSubscriptionRestoreSettings_ARMGenerator() gopter.Gen {
	if crossSubscriptionRestoreSettings_ARMGenerator != nil {
		return crossSubscriptionRestoreSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_ARM(generators)
	crossSubscriptionRestoreSettings_ARMGenerator = gen.Struct(reflect.TypeOf(CrossSubscriptionRestoreSettings_ARM{}), generators)

	return crossSubscriptionRestoreSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_ARM(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(CrossSubscriptionRestoreSettings_State_Disabled, CrossSubscriptionRestoreSettings_State_Enabled, CrossSubscriptionRestoreSettings_State_PermanentlyDisabled))
}

func Test_ImmutabilitySettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilitySettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilitySettings_ARM, ImmutabilitySettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilitySettings_ARM runs a test to see if a specific instance of ImmutabilitySettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilitySettings_ARM(subject ImmutabilitySettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilitySettings_ARM
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

// Generator of ImmutabilitySettings_ARM instances for property testing - lazily instantiated by
// ImmutabilitySettings_ARMGenerator()
var immutabilitySettings_ARMGenerator gopter.Gen

// ImmutabilitySettings_ARMGenerator returns a generator of ImmutabilitySettings_ARM instances for property testing.
func ImmutabilitySettings_ARMGenerator() gopter.Gen {
	if immutabilitySettings_ARMGenerator != nil {
		return immutabilitySettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilitySettings_ARM(generators)
	immutabilitySettings_ARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilitySettings_ARM{}), generators)

	return immutabilitySettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilitySettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilitySettings_ARM(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilitySettings_State_Disabled, ImmutabilitySettings_State_Locked, ImmutabilitySettings_State_Unlocked))
}

func Test_SoftDeleteSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SoftDeleteSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSoftDeleteSettings_ARM, SoftDeleteSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSoftDeleteSettings_ARM runs a test to see if a specific instance of SoftDeleteSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSoftDeleteSettings_ARM(subject SoftDeleteSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SoftDeleteSettings_ARM
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

// Generator of SoftDeleteSettings_ARM instances for property testing - lazily instantiated by
// SoftDeleteSettings_ARMGenerator()
var softDeleteSettings_ARMGenerator gopter.Gen

// SoftDeleteSettings_ARMGenerator returns a generator of SoftDeleteSettings_ARM instances for property testing.
func SoftDeleteSettings_ARMGenerator() gopter.Gen {
	if softDeleteSettings_ARMGenerator != nil {
		return softDeleteSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSoftDeleteSettings_ARM(generators)
	softDeleteSettings_ARMGenerator = gen.Struct(reflect.TypeOf(SoftDeleteSettings_ARM{}), generators)

	return softDeleteSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSoftDeleteSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSoftDeleteSettings_ARM(gens map[string]gopter.Gen) {
	gens["RetentionDurationInDays"] = gen.PtrOf(gen.Float64())
	gens["State"] = gen.PtrOf(gen.OneConstOf(SoftDeleteSettings_State_AlwaysOn, SoftDeleteSettings_State_Off, SoftDeleteSettings_State_On))
}