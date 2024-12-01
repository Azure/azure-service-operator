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

func Test_AzureMonitorAlertSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureMonitorAlertSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureMonitorAlertSettings_STATUS, AzureMonitorAlertSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureMonitorAlertSettings_STATUS runs a test to see if a specific instance of AzureMonitorAlertSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureMonitorAlertSettings_STATUS(subject AzureMonitorAlertSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureMonitorAlertSettings_STATUS
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

// Generator of AzureMonitorAlertSettings_STATUS instances for property testing - lazily instantiated by
// AzureMonitorAlertSettings_STATUSGenerator()
var azureMonitorAlertSettings_STATUSGenerator gopter.Gen

// AzureMonitorAlertSettings_STATUSGenerator returns a generator of AzureMonitorAlertSettings_STATUS instances for property testing.
func AzureMonitorAlertSettings_STATUSGenerator() gopter.Gen {
	if azureMonitorAlertSettings_STATUSGenerator != nil {
		return azureMonitorAlertSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_STATUS(generators)
	azureMonitorAlertSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(AzureMonitorAlertSettings_STATUS{}), generators)

	return azureMonitorAlertSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings_STATUS(gens map[string]gopter.Gen) {
	gens["AlertsForAllJobFailures"] = gen.PtrOf(gen.OneConstOf(AzureMonitorAlertSettings_AlertsForAllJobFailures_STATUS_Disabled, AzureMonitorAlertSettings_AlertsForAllJobFailures_STATUS_Enabled))
}

func Test_BackupVaultResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVaultResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVaultResource_STATUS, BackupVaultResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVaultResource_STATUS runs a test to see if a specific instance of BackupVaultResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVaultResource_STATUS(subject BackupVaultResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVaultResource_STATUS
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

// Generator of BackupVaultResource_STATUS instances for property testing - lazily instantiated by
// BackupVaultResource_STATUSGenerator()
var backupVaultResource_STATUSGenerator gopter.Gen

// BackupVaultResource_STATUSGenerator returns a generator of BackupVaultResource_STATUS instances for property testing.
// We first initialize backupVaultResource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackupVaultResource_STATUSGenerator() gopter.Gen {
	if backupVaultResource_STATUSGenerator != nil {
		return backupVaultResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVaultResource_STATUS(generators)
	backupVaultResource_STATUSGenerator = gen.Struct(reflect.TypeOf(BackupVaultResource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVaultResource_STATUS(generators)
	AddRelatedPropertyGeneratorsForBackupVaultResource_STATUS(generators)
	backupVaultResource_STATUSGenerator = gen.Struct(reflect.TypeOf(BackupVaultResource_STATUS{}), generators)

	return backupVaultResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForBackupVaultResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupVaultResource_STATUS(gens map[string]gopter.Gen) {
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackupVaultResource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVaultResource_STATUS(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(DppIdentityDetails_STATUSGenerator())
	gens["Properties"] = gen.PtrOf(BackupVault_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_BackupVault_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVault_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVault_STATUS, BackupVault_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVault_STATUS runs a test to see if a specific instance of BackupVault_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVault_STATUS(subject BackupVault_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVault_STATUS
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

// Generator of BackupVault_STATUS instances for property testing - lazily instantiated by BackupVault_STATUSGenerator()
var backupVault_STATUSGenerator gopter.Gen

// BackupVault_STATUSGenerator returns a generator of BackupVault_STATUS instances for property testing.
// We first initialize backupVault_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackupVault_STATUSGenerator() gopter.Gen {
	if backupVault_STATUSGenerator != nil {
		return backupVault_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_STATUS(generators)
	backupVault_STATUSGenerator = gen.Struct(reflect.TypeOf(BackupVault_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_STATUS(generators)
	AddRelatedPropertyGeneratorsForBackupVault_STATUS(generators)
	backupVault_STATUSGenerator = gen.Struct(reflect.TypeOf(BackupVault_STATUS{}), generators)

	return backupVault_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForBackupVault_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupVault_STATUS(gens map[string]gopter.Gen) {
	gens["IsVaultProtectedByResourceGuard"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		BackupVault_ProvisioningState_STATUS_Failed,
		BackupVault_ProvisioningState_STATUS_Provisioning,
		BackupVault_ProvisioningState_STATUS_Succeeded,
		BackupVault_ProvisioningState_STATUS_Unknown,
		BackupVault_ProvisioningState_STATUS_Updating))
	gens["ResourceMoveState"] = gen.PtrOf(gen.OneConstOf(
		BackupVault_ResourceMoveState_STATUS_CommitFailed,
		BackupVault_ResourceMoveState_STATUS_CommitTimedout,
		BackupVault_ResourceMoveState_STATUS_CriticalFailure,
		BackupVault_ResourceMoveState_STATUS_Failed,
		BackupVault_ResourceMoveState_STATUS_InProgress,
		BackupVault_ResourceMoveState_STATUS_MoveSucceeded,
		BackupVault_ResourceMoveState_STATUS_PartialSuccess,
		BackupVault_ResourceMoveState_STATUS_PrepareFailed,
		BackupVault_ResourceMoveState_STATUS_PrepareTimedout,
		BackupVault_ResourceMoveState_STATUS_Unknown))
}

// AddRelatedPropertyGeneratorsForBackupVault_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVault_STATUS(gens map[string]gopter.Gen) {
	gens["FeatureSettings"] = gen.PtrOf(FeatureSettings_STATUSGenerator())
	gens["MonitoringSettings"] = gen.PtrOf(MonitoringSettings_STATUSGenerator())
	gens["ResourceMoveDetails"] = gen.PtrOf(ResourceMoveDetails_STATUSGenerator())
	gens["SecuritySettings"] = gen.PtrOf(SecuritySettings_STATUSGenerator())
	gens["StorageSettings"] = gen.SliceOf(StorageSetting_STATUSGenerator())
}

func Test_CrossSubscriptionRestoreSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CrossSubscriptionRestoreSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCrossSubscriptionRestoreSettings_STATUS, CrossSubscriptionRestoreSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCrossSubscriptionRestoreSettings_STATUS runs a test to see if a specific instance of CrossSubscriptionRestoreSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCrossSubscriptionRestoreSettings_STATUS(subject CrossSubscriptionRestoreSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CrossSubscriptionRestoreSettings_STATUS
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

// Generator of CrossSubscriptionRestoreSettings_STATUS instances for property testing - lazily instantiated by
// CrossSubscriptionRestoreSettings_STATUSGenerator()
var crossSubscriptionRestoreSettings_STATUSGenerator gopter.Gen

// CrossSubscriptionRestoreSettings_STATUSGenerator returns a generator of CrossSubscriptionRestoreSettings_STATUS instances for property testing.
func CrossSubscriptionRestoreSettings_STATUSGenerator() gopter.Gen {
	if crossSubscriptionRestoreSettings_STATUSGenerator != nil {
		return crossSubscriptionRestoreSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_STATUS(generators)
	crossSubscriptionRestoreSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(CrossSubscriptionRestoreSettings_STATUS{}), generators)

	return crossSubscriptionRestoreSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings_STATUS(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(CrossSubscriptionRestoreSettings_State_STATUS_Disabled, CrossSubscriptionRestoreSettings_State_STATUS_Enabled, CrossSubscriptionRestoreSettings_State_STATUS_PermanentlyDisabled))
}

func Test_DppIdentityDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DppIdentityDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDppIdentityDetails_STATUS, DppIdentityDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDppIdentityDetails_STATUS runs a test to see if a specific instance of DppIdentityDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDppIdentityDetails_STATUS(subject DppIdentityDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DppIdentityDetails_STATUS
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

// Generator of DppIdentityDetails_STATUS instances for property testing - lazily instantiated by
// DppIdentityDetails_STATUSGenerator()
var dppIdentityDetails_STATUSGenerator gopter.Gen

// DppIdentityDetails_STATUSGenerator returns a generator of DppIdentityDetails_STATUS instances for property testing.
func DppIdentityDetails_STATUSGenerator() gopter.Gen {
	if dppIdentityDetails_STATUSGenerator != nil {
		return dppIdentityDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDppIdentityDetails_STATUS(generators)
	dppIdentityDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(DppIdentityDetails_STATUS{}), generators)

	return dppIdentityDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDppIdentityDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDppIdentityDetails_STATUS(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_FeatureSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FeatureSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFeatureSettings_STATUS, FeatureSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFeatureSettings_STATUS runs a test to see if a specific instance of FeatureSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFeatureSettings_STATUS(subject FeatureSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FeatureSettings_STATUS
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

// Generator of FeatureSettings_STATUS instances for property testing - lazily instantiated by
// FeatureSettings_STATUSGenerator()
var featureSettings_STATUSGenerator gopter.Gen

// FeatureSettings_STATUSGenerator returns a generator of FeatureSettings_STATUS instances for property testing.
func FeatureSettings_STATUSGenerator() gopter.Gen {
	if featureSettings_STATUSGenerator != nil {
		return featureSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFeatureSettings_STATUS(generators)
	featureSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(FeatureSettings_STATUS{}), generators)

	return featureSettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForFeatureSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFeatureSettings_STATUS(gens map[string]gopter.Gen) {
	gens["CrossSubscriptionRestoreSettings"] = gen.PtrOf(CrossSubscriptionRestoreSettings_STATUSGenerator())
}

func Test_ImmutabilitySettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilitySettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilitySettings_STATUS, ImmutabilitySettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilitySettings_STATUS runs a test to see if a specific instance of ImmutabilitySettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilitySettings_STATUS(subject ImmutabilitySettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilitySettings_STATUS
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

// Generator of ImmutabilitySettings_STATUS instances for property testing - lazily instantiated by
// ImmutabilitySettings_STATUSGenerator()
var immutabilitySettings_STATUSGenerator gopter.Gen

// ImmutabilitySettings_STATUSGenerator returns a generator of ImmutabilitySettings_STATUS instances for property testing.
func ImmutabilitySettings_STATUSGenerator() gopter.Gen {
	if immutabilitySettings_STATUSGenerator != nil {
		return immutabilitySettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilitySettings_STATUS(generators)
	immutabilitySettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilitySettings_STATUS{}), generators)

	return immutabilitySettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilitySettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilitySettings_STATUS(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilitySettings_State_STATUS_Disabled, ImmutabilitySettings_State_STATUS_Locked, ImmutabilitySettings_State_STATUS_Unlocked))
}

func Test_MonitoringSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitoringSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitoringSettings_STATUS, MonitoringSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitoringSettings_STATUS runs a test to see if a specific instance of MonitoringSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitoringSettings_STATUS(subject MonitoringSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitoringSettings_STATUS
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

// Generator of MonitoringSettings_STATUS instances for property testing - lazily instantiated by
// MonitoringSettings_STATUSGenerator()
var monitoringSettings_STATUSGenerator gopter.Gen

// MonitoringSettings_STATUSGenerator returns a generator of MonitoringSettings_STATUS instances for property testing.
func MonitoringSettings_STATUSGenerator() gopter.Gen {
	if monitoringSettings_STATUSGenerator != nil {
		return monitoringSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMonitoringSettings_STATUS(generators)
	monitoringSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(MonitoringSettings_STATUS{}), generators)

	return monitoringSettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForMonitoringSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMonitoringSettings_STATUS(gens map[string]gopter.Gen) {
	gens["AzureMonitorAlertSettings"] = gen.PtrOf(AzureMonitorAlertSettings_STATUSGenerator())
}

func Test_ResourceMoveDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceMoveDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceMoveDetails_STATUS, ResourceMoveDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceMoveDetails_STATUS runs a test to see if a specific instance of ResourceMoveDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceMoveDetails_STATUS(subject ResourceMoveDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceMoveDetails_STATUS
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

// Generator of ResourceMoveDetails_STATUS instances for property testing - lazily instantiated by
// ResourceMoveDetails_STATUSGenerator()
var resourceMoveDetails_STATUSGenerator gopter.Gen

// ResourceMoveDetails_STATUSGenerator returns a generator of ResourceMoveDetails_STATUS instances for property testing.
func ResourceMoveDetails_STATUSGenerator() gopter.Gen {
	if resourceMoveDetails_STATUSGenerator != nil {
		return resourceMoveDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceMoveDetails_STATUS(generators)
	resourceMoveDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceMoveDetails_STATUS{}), generators)

	return resourceMoveDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceMoveDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceMoveDetails_STATUS(gens map[string]gopter.Gen) {
	gens["CompletionTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["OperationId"] = gen.PtrOf(gen.AlphaString())
	gens["SourceResourcePath"] = gen.PtrOf(gen.AlphaString())
	gens["StartTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["TargetResourcePath"] = gen.PtrOf(gen.AlphaString())
}

func Test_SecuritySettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecuritySettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecuritySettings_STATUS, SecuritySettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecuritySettings_STATUS runs a test to see if a specific instance of SecuritySettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSecuritySettings_STATUS(subject SecuritySettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecuritySettings_STATUS
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

// Generator of SecuritySettings_STATUS instances for property testing - lazily instantiated by
// SecuritySettings_STATUSGenerator()
var securitySettings_STATUSGenerator gopter.Gen

// SecuritySettings_STATUSGenerator returns a generator of SecuritySettings_STATUS instances for property testing.
func SecuritySettings_STATUSGenerator() gopter.Gen {
	if securitySettings_STATUSGenerator != nil {
		return securitySettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecuritySettings_STATUS(generators)
	securitySettings_STATUSGenerator = gen.Struct(reflect.TypeOf(SecuritySettings_STATUS{}), generators)

	return securitySettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForSecuritySettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecuritySettings_STATUS(gens map[string]gopter.Gen) {
	gens["ImmutabilitySettings"] = gen.PtrOf(ImmutabilitySettings_STATUSGenerator())
	gens["SoftDeleteSettings"] = gen.PtrOf(SoftDeleteSettings_STATUSGenerator())
}

func Test_SoftDeleteSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SoftDeleteSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSoftDeleteSettings_STATUS, SoftDeleteSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSoftDeleteSettings_STATUS runs a test to see if a specific instance of SoftDeleteSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSoftDeleteSettings_STATUS(subject SoftDeleteSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SoftDeleteSettings_STATUS
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

// Generator of SoftDeleteSettings_STATUS instances for property testing - lazily instantiated by
// SoftDeleteSettings_STATUSGenerator()
var softDeleteSettings_STATUSGenerator gopter.Gen

// SoftDeleteSettings_STATUSGenerator returns a generator of SoftDeleteSettings_STATUS instances for property testing.
func SoftDeleteSettings_STATUSGenerator() gopter.Gen {
	if softDeleteSettings_STATUSGenerator != nil {
		return softDeleteSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSoftDeleteSettings_STATUS(generators)
	softDeleteSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(SoftDeleteSettings_STATUS{}), generators)

	return softDeleteSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSoftDeleteSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSoftDeleteSettings_STATUS(gens map[string]gopter.Gen) {
	gens["RetentionDurationInDays"] = gen.PtrOf(gen.Float64())
	gens["State"] = gen.PtrOf(gen.OneConstOf(SoftDeleteSettings_State_STATUS_AlwaysOn, SoftDeleteSettings_State_STATUS_Off, SoftDeleteSettings_State_STATUS_On))
}

func Test_StorageSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageSetting_STATUS, StorageSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageSetting_STATUS runs a test to see if a specific instance of StorageSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageSetting_STATUS(subject StorageSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageSetting_STATUS
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

// Generator of StorageSetting_STATUS instances for property testing - lazily instantiated by
// StorageSetting_STATUSGenerator()
var storageSetting_STATUSGenerator gopter.Gen

// StorageSetting_STATUSGenerator returns a generator of StorageSetting_STATUS instances for property testing.
func StorageSetting_STATUSGenerator() gopter.Gen {
	if storageSetting_STATUSGenerator != nil {
		return storageSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageSetting_STATUS(generators)
	storageSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageSetting_STATUS{}), generators)

	return storageSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageSetting_STATUS(gens map[string]gopter.Gen) {
	gens["DatastoreType"] = gen.PtrOf(gen.OneConstOf(StorageSetting_DatastoreType_STATUS_ArchiveStore, StorageSetting_DatastoreType_STATUS_OperationalStore, StorageSetting_DatastoreType_STATUS_VaultStore))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(StorageSetting_Type_STATUS_GeoRedundant, StorageSetting_Type_STATUS_LocallyRedundant, StorageSetting_Type_STATUS_ZoneRedundant))
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
