// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

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

func Test_Workspace_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_SpecARM, Workspace_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_SpecARM runs a test to see if a specific instance of Workspace_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_SpecARM(subject Workspace_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_SpecARM
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

// Generator of Workspace_SpecARM instances for property testing - lazily instantiated by Workspace_SpecARMGenerator()
var workspace_SpecARMGenerator gopter.Gen

// Workspace_SpecARMGenerator returns a generator of Workspace_SpecARM instances for property testing.
// We first initialize workspace_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_SpecARMGenerator() gopter.Gen {
	if workspace_SpecARMGenerator != nil {
		return workspace_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_SpecARM(generators)
	workspace_SpecARMGenerator = gen.Struct(reflect.TypeOf(Workspace_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspace_SpecARM(generators)
	workspace_SpecARMGenerator = gen.Struct(reflect.TypeOf(Workspace_SpecARM{}), generators)

	return workspace_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_SpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityARMGenerator())
	gens["Properties"] = gen.PtrOf(WorkspacePropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataARMGenerator())
}

func Test_IdentityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityARM, IdentityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityARM runs a test to see if a specific instance of IdentityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityARM(subject IdentityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityARM
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

// Generator of IdentityARM instances for property testing - lazily instantiated by IdentityARMGenerator()
var identityARMGenerator gopter.Gen

// IdentityARMGenerator returns a generator of IdentityARM instances for property testing.
func IdentityARMGenerator() gopter.Gen {
	if identityARMGenerator != nil {
		return identityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityARM(generators)
	identityARMGenerator = gen.Struct(reflect.TypeOf(IdentityARM{}), generators)

	return identityARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_None,
		Identity_Type_SystemAssigned,
		Identity_Type_SystemAssignedUserAssigned,
		Identity_Type_UserAssigned))
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
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

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemDataARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemDataARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataARM, SystemDataARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataARM runs a test to see if a specific instance of SystemDataARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataARM(subject SystemDataARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemDataARM
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

// Generator of SystemDataARM instances for property testing - lazily instantiated by SystemDataARMGenerator()
var systemDataARMGenerator gopter.Gen

// SystemDataARMGenerator returns a generator of SystemDataARM instances for property testing.
func SystemDataARMGenerator() gopter.Gen {
	if systemDataARMGenerator != nil {
		return systemDataARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataARM(generators)
	systemDataARMGenerator = gen.Struct(reflect.TypeOf(SystemDataARM{}), generators)

	return systemDataARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_Application,
		SystemData_CreatedByType_Key,
		SystemData_CreatedByType_ManagedIdentity,
		SystemData_CreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_Application,
		SystemData_LastModifiedByType_Key,
		SystemData_LastModifiedByType_ManagedIdentity,
		SystemData_LastModifiedByType_User))
}

func Test_WorkspacePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspacePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacePropertiesARM, WorkspacePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacePropertiesARM runs a test to see if a specific instance of WorkspacePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacePropertiesARM(subject WorkspacePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspacePropertiesARM
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

// Generator of WorkspacePropertiesARM instances for property testing - lazily instantiated by
// WorkspacePropertiesARMGenerator()
var workspacePropertiesARMGenerator gopter.Gen

// WorkspacePropertiesARMGenerator returns a generator of WorkspacePropertiesARM instances for property testing.
// We first initialize workspacePropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacePropertiesARMGenerator() gopter.Gen {
	if workspacePropertiesARMGenerator != nil {
		return workspacePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacePropertiesARM(generators)
	workspacePropertiesARMGenerator = gen.Struct(reflect.TypeOf(WorkspacePropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacePropertiesARM(generators)
	AddRelatedPropertyGeneratorsForWorkspacePropertiesARM(generators)
	workspacePropertiesARMGenerator = gen.Struct(reflect.TypeOf(WorkspacePropertiesARM{}), generators)

	return workspacePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacePropertiesARM(gens map[string]gopter.Gen) {
	gens["AllowPublicAccessWhenBehindVnet"] = gen.PtrOf(gen.Bool())
	gens["ApplicationInsights"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerRegistry"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DiscoveryUrl"] = gen.PtrOf(gen.AlphaString())
	gens["FriendlyName"] = gen.PtrOf(gen.AlphaString())
	gens["HbiWorkspace"] = gen.PtrOf(gen.Bool())
	gens["ImageBuildCompute"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVault"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccess_Disabled, WorkspaceProperties_PublicNetworkAccess_Enabled))
	gens["StorageAccount"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspacePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacePropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionPropertyARMGenerator())
	gens["ServiceManagedResourcesSettings"] = gen.PtrOf(ServiceManagedResourcesSettingsARMGenerator())
	gens["SharedPrivateLinkResources"] = gen.SliceOf(SharedPrivateLinkResourceARMGenerator())
}

func Test_EncryptionPropertyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionPropertyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionPropertyARM, EncryptionPropertyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionPropertyARM runs a test to see if a specific instance of EncryptionPropertyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionPropertyARM(subject EncryptionPropertyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionPropertyARM
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

// Generator of EncryptionPropertyARM instances for property testing - lazily instantiated by
// EncryptionPropertyARMGenerator()
var encryptionPropertyARMGenerator gopter.Gen

// EncryptionPropertyARMGenerator returns a generator of EncryptionPropertyARM instances for property testing.
// We first initialize encryptionPropertyARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionPropertyARMGenerator() gopter.Gen {
	if encryptionPropertyARMGenerator != nil {
		return encryptionPropertyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertyARM(generators)
	encryptionPropertyARMGenerator = gen.Struct(reflect.TypeOf(EncryptionPropertyARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionPropertyARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionPropertyARM(generators)
	encryptionPropertyARMGenerator = gen.Struct(reflect.TypeOf(EncryptionPropertyARM{}), generators)

	return encryptionPropertyARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionPropertyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionPropertyARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(EncryptionProperty_Status_Disabled, EncryptionProperty_Status_Enabled))
}

// AddRelatedPropertyGeneratorsForEncryptionPropertyARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionPropertyARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityForCmkARMGenerator())
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesARMGenerator())
}

func Test_ServiceManagedResourcesSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceManagedResourcesSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceManagedResourcesSettingsARM, ServiceManagedResourcesSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceManagedResourcesSettingsARM runs a test to see if a specific instance of ServiceManagedResourcesSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceManagedResourcesSettingsARM(subject ServiceManagedResourcesSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceManagedResourcesSettingsARM
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

// Generator of ServiceManagedResourcesSettingsARM instances for property testing - lazily instantiated by
// ServiceManagedResourcesSettingsARMGenerator()
var serviceManagedResourcesSettingsARMGenerator gopter.Gen

// ServiceManagedResourcesSettingsARMGenerator returns a generator of ServiceManagedResourcesSettingsARM instances for property testing.
func ServiceManagedResourcesSettingsARMGenerator() gopter.Gen {
	if serviceManagedResourcesSettingsARMGenerator != nil {
		return serviceManagedResourcesSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServiceManagedResourcesSettingsARM(generators)
	serviceManagedResourcesSettingsARMGenerator = gen.Struct(reflect.TypeOf(ServiceManagedResourcesSettingsARM{}), generators)

	return serviceManagedResourcesSettingsARMGenerator
}

// AddRelatedPropertyGeneratorsForServiceManagedResourcesSettingsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceManagedResourcesSettingsARM(gens map[string]gopter.Gen) {
	gens["CosmosDb"] = gen.PtrOf(CosmosDbSettingsARMGenerator())
}

func Test_SharedPrivateLinkResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceARM, SharedPrivateLinkResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceARM runs a test to see if a specific instance of SharedPrivateLinkResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceARM(subject SharedPrivateLinkResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceARM
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

// Generator of SharedPrivateLinkResourceARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceARMGenerator()
var sharedPrivateLinkResourceARMGenerator gopter.Gen

// SharedPrivateLinkResourceARMGenerator returns a generator of SharedPrivateLinkResourceARM instances for property testing.
// We first initialize sharedPrivateLinkResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourceARMGenerator() gopter.Gen {
	if sharedPrivateLinkResourceARMGenerator != nil {
		return sharedPrivateLinkResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceARM(generators)
	sharedPrivateLinkResourceARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceARM(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceARM(generators)
	sharedPrivateLinkResourceARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceARM{}), generators)

	return sharedPrivateLinkResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResourceARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedPrivateLinkResourcePropertyARMGenerator())
}

func Test_CosmosDbSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CosmosDbSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCosmosDbSettingsARM, CosmosDbSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCosmosDbSettingsARM runs a test to see if a specific instance of CosmosDbSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCosmosDbSettingsARM(subject CosmosDbSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CosmosDbSettingsARM
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

// Generator of CosmosDbSettingsARM instances for property testing - lazily instantiated by
// CosmosDbSettingsARMGenerator()
var cosmosDbSettingsARMGenerator gopter.Gen

// CosmosDbSettingsARMGenerator returns a generator of CosmosDbSettingsARM instances for property testing.
func CosmosDbSettingsARMGenerator() gopter.Gen {
	if cosmosDbSettingsARMGenerator != nil {
		return cosmosDbSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCosmosDbSettingsARM(generators)
	cosmosDbSettingsARMGenerator = gen.Struct(reflect.TypeOf(CosmosDbSettingsARM{}), generators)

	return cosmosDbSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForCosmosDbSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCosmosDbSettingsARM(gens map[string]gopter.Gen) {
	gens["CollectionsThroughput"] = gen.PtrOf(gen.Int())
}

func Test_IdentityForCmkARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityForCmkARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityForCmkARM, IdentityForCmkARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityForCmkARM runs a test to see if a specific instance of IdentityForCmkARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityForCmkARM(subject IdentityForCmkARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityForCmkARM
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

// Generator of IdentityForCmkARM instances for property testing - lazily instantiated by IdentityForCmkARMGenerator()
var identityForCmkARMGenerator gopter.Gen

// IdentityForCmkARMGenerator returns a generator of IdentityForCmkARM instances for property testing.
func IdentityForCmkARMGenerator() gopter.Gen {
	if identityForCmkARMGenerator != nil {
		return identityForCmkARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityForCmkARM(generators)
	identityForCmkARMGenerator = gen.Struct(reflect.TypeOf(IdentityForCmkARM{}), generators)

	return identityForCmkARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityForCmkARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityForCmkARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesARM, KeyVaultPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesARM runs a test to see if a specific instance of KeyVaultPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesARM(subject KeyVaultPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultPropertiesARM
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

// Generator of KeyVaultPropertiesARM instances for property testing - lazily instantiated by
// KeyVaultPropertiesARMGenerator()
var keyVaultPropertiesARMGenerator gopter.Gen

// KeyVaultPropertiesARMGenerator returns a generator of KeyVaultPropertiesARM instances for property testing.
func KeyVaultPropertiesARMGenerator() gopter.Gen {
	if keyVaultPropertiesARMGenerator != nil {
		return keyVaultPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	return keyVaultPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultArmId"] = gen.PtrOf(gen.AlphaString())
}

func Test_SharedPrivateLinkResourcePropertyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourcePropertyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourcePropertyARM, SharedPrivateLinkResourcePropertyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourcePropertyARM runs a test to see if a specific instance of SharedPrivateLinkResourcePropertyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourcePropertyARM(subject SharedPrivateLinkResourcePropertyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourcePropertyARM
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

// Generator of SharedPrivateLinkResourcePropertyARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResourcePropertyARMGenerator()
var sharedPrivateLinkResourcePropertyARMGenerator gopter.Gen

// SharedPrivateLinkResourcePropertyARMGenerator returns a generator of SharedPrivateLinkResourcePropertyARM instances for property testing.
func SharedPrivateLinkResourcePropertyARMGenerator() gopter.Gen {
	if sharedPrivateLinkResourcePropertyARMGenerator != nil {
		return sharedPrivateLinkResourcePropertyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourcePropertyARM(generators)
	sharedPrivateLinkResourcePropertyARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourcePropertyARM{}), generators)

	return sharedPrivateLinkResourcePropertyARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourcePropertyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourcePropertyARM(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointServiceConnectionStatus_Approved,
		PrivateEndpointServiceConnectionStatus_Disconnected,
		PrivateEndpointServiceConnectionStatus_Pending,
		PrivateEndpointServiceConnectionStatus_Rejected,
		PrivateEndpointServiceConnectionStatus_Timeout))
}
