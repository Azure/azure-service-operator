// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220501

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

func Test_ConfigurationStore_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationStore_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationStore_STATUS_ARM, ConfigurationStore_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationStore_STATUS_ARM runs a test to see if a specific instance of ConfigurationStore_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationStore_STATUS_ARM(subject ConfigurationStore_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationStore_STATUS_ARM
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

// Generator of ConfigurationStore_STATUS_ARM instances for property testing - lazily instantiated by
// ConfigurationStore_STATUS_ARMGenerator()
var configurationStore_STATUS_ARMGenerator gopter.Gen

// ConfigurationStore_STATUS_ARMGenerator returns a generator of ConfigurationStore_STATUS_ARM instances for property testing.
// We first initialize configurationStore_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationStore_STATUS_ARMGenerator() gopter.Gen {
	if configurationStore_STATUS_ARMGenerator != nil {
		return configurationStore_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStore_STATUS_ARM(generators)
	configurationStore_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationStore_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStore_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForConfigurationStore_STATUS_ARM(generators)
	configurationStore_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationStore_STATUS_ARM{}), generators)

	return configurationStore_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationStore_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationStore_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfigurationStore_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationStore_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(ResourceIdentity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ConfigurationStoreProperties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_ConfigurationStoreProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationStoreProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationStoreProperties_STATUS_ARM, ConfigurationStoreProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationStoreProperties_STATUS_ARM runs a test to see if a specific instance of ConfigurationStoreProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationStoreProperties_STATUS_ARM(subject ConfigurationStoreProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationStoreProperties_STATUS_ARM
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

// Generator of ConfigurationStoreProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ConfigurationStoreProperties_STATUS_ARMGenerator()
var configurationStoreProperties_STATUS_ARMGenerator gopter.Gen

// ConfigurationStoreProperties_STATUS_ARMGenerator returns a generator of ConfigurationStoreProperties_STATUS_ARM instances for property testing.
// We first initialize configurationStoreProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationStoreProperties_STATUS_ARMGenerator() gopter.Gen {
	if configurationStoreProperties_STATUS_ARMGenerator != nil {
		return configurationStoreProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM(generators)
	configurationStoreProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationStoreProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM(generators)
	configurationStoreProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationStoreProperties_STATUS_ARM{}), generators)

	return configurationStoreProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(ConfigurationStoreProperties_CreateMode_STATUS_Default, ConfigurationStoreProperties_CreateMode_STATUS_Recover))
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnablePurgeProtection"] = gen.PtrOf(gen.Bool())
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ConfigurationStoreProperties_ProvisioningState_STATUS_Canceled,
		ConfigurationStoreProperties_ProvisioningState_STATUS_Creating,
		ConfigurationStoreProperties_ProvisioningState_STATUS_Deleting,
		ConfigurationStoreProperties_ProvisioningState_STATUS_Failed,
		ConfigurationStoreProperties_ProvisioningState_STATUS_Succeeded,
		ConfigurationStoreProperties_ProvisioningState_STATUS_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Disabled, ConfigurationStoreProperties_PublicNetworkAccess_STATUS_Enabled))
	gens["SoftDeleteRetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationStoreProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionProperties_STATUS_ARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionReference_STATUS_ARMGenerator())
}

func Test_ResourceIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceIdentity_STATUS_ARM, ResourceIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceIdentity_STATUS_ARM runs a test to see if a specific instance of ResourceIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceIdentity_STATUS_ARM(subject ResourceIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceIdentity_STATUS_ARM
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

// Generator of ResourceIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// ResourceIdentity_STATUS_ARMGenerator()
var resourceIdentity_STATUS_ARMGenerator gopter.Gen

// ResourceIdentity_STATUS_ARMGenerator returns a generator of ResourceIdentity_STATUS_ARM instances for property testing.
// We first initialize resourceIdentity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ResourceIdentity_STATUS_ARMGenerator() gopter.Gen {
	if resourceIdentity_STATUS_ARMGenerator != nil {
		return resourceIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceIdentity_STATUS_ARM(generators)
	resourceIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceIdentity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForResourceIdentity_STATUS_ARM(generators)
	resourceIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceIdentity_STATUS_ARM{}), generators)

	return resourceIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ResourceIdentity_Type_STATUS_None,
		ResourceIdentity_Type_STATUS_SystemAssigned,
		ResourceIdentity_Type_STATUS_SystemAssignedUserAssigned,
		ResourceIdentity_Type_STATUS_UserAssigned))
}

// AddRelatedPropertyGeneratorsForResourceIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForResourceIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserIdentity_STATUS_ARMGenerator())
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

func Test_EncryptionProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperties_STATUS_ARM, EncryptionProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperties_STATUS_ARM runs a test to see if a specific instance of EncryptionProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperties_STATUS_ARM(subject EncryptionProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperties_STATUS_ARM
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

// Generator of EncryptionProperties_STATUS_ARM instances for property testing - lazily instantiated by
// EncryptionProperties_STATUS_ARMGenerator()
var encryptionProperties_STATUS_ARMGenerator gopter.Gen

// EncryptionProperties_STATUS_ARMGenerator returns a generator of EncryptionProperties_STATUS_ARM instances for property testing.
func EncryptionProperties_STATUS_ARMGenerator() gopter.Gen {
	if encryptionProperties_STATUS_ARMGenerator != nil {
		return encryptionProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForEncryptionProperties_STATUS_ARM(generators)
	encryptionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_STATUS_ARM{}), generators)

	return encryptionProperties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForEncryptionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUS_ARMGenerator())
}

func Test_PrivateEndpointConnectionReference_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnectionReference_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionReference_STATUS_ARM, PrivateEndpointConnectionReference_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionReference_STATUS_ARM runs a test to see if a specific instance of PrivateEndpointConnectionReference_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionReference_STATUS_ARM(subject PrivateEndpointConnectionReference_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnectionReference_STATUS_ARM
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

// Generator of PrivateEndpointConnectionReference_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpointConnectionReference_STATUS_ARMGenerator()
var privateEndpointConnectionReference_STATUS_ARMGenerator gopter.Gen

// PrivateEndpointConnectionReference_STATUS_ARMGenerator returns a generator of PrivateEndpointConnectionReference_STATUS_ARM instances for property testing.
func PrivateEndpointConnectionReference_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpointConnectionReference_STATUS_ARMGenerator != nil {
		return privateEndpointConnectionReference_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionReference_STATUS_ARM(generators)
	privateEndpointConnectionReference_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionReference_STATUS_ARM{}), generators)

	return privateEndpointConnectionReference_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionReference_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionReference_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
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

func Test_KeyVaultProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM, KeyVaultProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM runs a test to see if a specific instance of KeyVaultProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM(subject KeyVaultProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_STATUS_ARM
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

// Generator of KeyVaultProperties_STATUS_ARM instances for property testing - lazily instantiated by
// KeyVaultProperties_STATUS_ARMGenerator()
var keyVaultProperties_STATUS_ARMGenerator gopter.Gen

// KeyVaultProperties_STATUS_ARMGenerator returns a generator of KeyVaultProperties_STATUS_ARM instances for property testing.
func KeyVaultProperties_STATUS_ARMGenerator() gopter.Gen {
	if keyVaultProperties_STATUS_ARMGenerator != nil {
		return keyVaultProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(generators)
	keyVaultProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUS_ARM{}), generators)

	return keyVaultProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}
