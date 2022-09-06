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

func Test_Workspace_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_STATUSARM, Workspace_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_STATUSARM runs a test to see if a specific instance of Workspace_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_STATUSARM(subject Workspace_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_STATUSARM
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

// Generator of Workspace_STATUSARM instances for property testing - lazily instantiated by
// Workspace_STATUSARMGenerator()
var workspace_STATUSARMGenerator gopter.Gen

// Workspace_STATUSARMGenerator returns a generator of Workspace_STATUSARM instances for property testing.
// We first initialize workspace_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_STATUSARMGenerator() gopter.Gen {
	if workspace_STATUSARMGenerator != nil {
		return workspace_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_STATUSARM(generators)
	workspace_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Workspace_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForWorkspace_STATUSARM(generators)
	workspace_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Workspace_STATUSARM{}), generators)

	return workspace_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_WorkspaceProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties_STATUSARM, WorkspaceProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties_STATUSARM runs a test to see if a specific instance of WorkspaceProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties_STATUSARM(subject WorkspaceProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_STATUSARM
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

// Generator of WorkspaceProperties_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceProperties_STATUSARMGenerator()
var workspaceProperties_STATUSARMGenerator gopter.Gen

// WorkspaceProperties_STATUSARMGenerator returns a generator of WorkspaceProperties_STATUSARM instances for property testing.
// We first initialize workspaceProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceProperties_STATUSARMGenerator() gopter.Gen {
	if workspaceProperties_STATUSARMGenerator != nil {
		return workspaceProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUSARM(generators)
	workspaceProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUSARM(generators)
	workspaceProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_STATUSARM{}), generators)

	return workspaceProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AllowPublicAccessWhenBehindVnet"] = gen.PtrOf(gen.Bool())
	gens["ApplicationInsights"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerRegistry"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DiscoveryUrl"] = gen.PtrOf(gen.AlphaString())
	gens["FriendlyName"] = gen.PtrOf(gen.AlphaString())
	gens["HbiWorkspace"] = gen.PtrOf(gen.Bool())
	gens["ImageBuildCompute"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVault"] = gen.PtrOf(gen.AlphaString())
	gens["MlFlowTrackingUri"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkCount"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceProperties_STATUS_ProvisioningState_Canceled,
		WorkspaceProperties_STATUS_ProvisioningState_Creating,
		WorkspaceProperties_STATUS_ProvisioningState_Deleting,
		WorkspaceProperties_STATUS_ProvisioningState_Failed,
		WorkspaceProperties_STATUS_ProvisioningState_Succeeded,
		WorkspaceProperties_STATUS_ProvisioningState_Unknown,
		WorkspaceProperties_STATUS_ProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_STATUS_PublicNetworkAccess_Disabled, WorkspaceProperties_STATUS_PublicNetworkAccess_Enabled))
	gens["ServiceProvisionedResourceGroup"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccount"] = gen.PtrOf(gen.AlphaString())
	gens["StorageHnsEnabled"] = gen.PtrOf(gen.Bool())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionProperty_STATUSARMGenerator())
	gens["NotebookInfo"] = gen.PtrOf(NotebookResourceInfo_STATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator())
	gens["ServiceManagedResourcesSettings"] = gen.PtrOf(ServiceManagedResourcesSettings_STATUSARMGenerator())
	gens["SharedPrivateLinkResources"] = gen.SliceOf(SharedPrivateLinkResource_STATUSARMGenerator())
}

func Test_EncryptionProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperty_STATUSARM, EncryptionProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperty_STATUSARM runs a test to see if a specific instance of EncryptionProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperty_STATUSARM(subject EncryptionProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperty_STATUSARM
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

// Generator of EncryptionProperty_STATUSARM instances for property testing - lazily instantiated by
// EncryptionProperty_STATUSARMGenerator()
var encryptionProperty_STATUSARMGenerator gopter.Gen

// EncryptionProperty_STATUSARMGenerator returns a generator of EncryptionProperty_STATUSARM instances for property testing.
// We first initialize encryptionProperty_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionProperty_STATUSARMGenerator() gopter.Gen {
	if encryptionProperty_STATUSARMGenerator != nil {
		return encryptionProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperty_STATUSARM(generators)
	encryptionProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperty_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperty_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionProperty_STATUSARM(generators)
	encryptionProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperty_STATUSARM{}), generators)

	return encryptionProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(EncryptionProperty_STATUS_Status_Disabled, EncryptionProperty_STATUS_Status_Enabled))
}

// AddRelatedPropertyGeneratorsForEncryptionProperty_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityForCmk_STATUSARMGenerator())
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUSARMGenerator())
}

func Test_NotebookResourceInfo_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NotebookResourceInfo_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNotebookResourceInfo_STATUSARM, NotebookResourceInfo_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNotebookResourceInfo_STATUSARM runs a test to see if a specific instance of NotebookResourceInfo_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNotebookResourceInfo_STATUSARM(subject NotebookResourceInfo_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NotebookResourceInfo_STATUSARM
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

// Generator of NotebookResourceInfo_STATUSARM instances for property testing - lazily instantiated by
// NotebookResourceInfo_STATUSARMGenerator()
var notebookResourceInfo_STATUSARMGenerator gopter.Gen

// NotebookResourceInfo_STATUSARMGenerator returns a generator of NotebookResourceInfo_STATUSARM instances for property testing.
// We first initialize notebookResourceInfo_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NotebookResourceInfo_STATUSARMGenerator() gopter.Gen {
	if notebookResourceInfo_STATUSARMGenerator != nil {
		return notebookResourceInfo_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNotebookResourceInfo_STATUSARM(generators)
	notebookResourceInfo_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NotebookResourceInfo_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNotebookResourceInfo_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForNotebookResourceInfo_STATUSARM(generators)
	notebookResourceInfo_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NotebookResourceInfo_STATUSARM{}), generators)

	return notebookResourceInfo_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNotebookResourceInfo_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNotebookResourceInfo_STATUSARM(gens map[string]gopter.Gen) {
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNotebookResourceInfo_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNotebookResourceInfo_STATUSARM(gens map[string]gopter.Gen) {
	gens["NotebookPreparationError"] = gen.PtrOf(NotebookPreparationError_STATUSARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM, PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()
var privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing.
// We first initialize privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_ServiceManagedResourcesSettings_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceManagedResourcesSettings_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceManagedResourcesSettings_STATUSARM, ServiceManagedResourcesSettings_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceManagedResourcesSettings_STATUSARM runs a test to see if a specific instance of ServiceManagedResourcesSettings_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceManagedResourcesSettings_STATUSARM(subject ServiceManagedResourcesSettings_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceManagedResourcesSettings_STATUSARM
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

// Generator of ServiceManagedResourcesSettings_STATUSARM instances for property testing - lazily instantiated by
// ServiceManagedResourcesSettings_STATUSARMGenerator()
var serviceManagedResourcesSettings_STATUSARMGenerator gopter.Gen

// ServiceManagedResourcesSettings_STATUSARMGenerator returns a generator of ServiceManagedResourcesSettings_STATUSARM instances for property testing.
func ServiceManagedResourcesSettings_STATUSARMGenerator() gopter.Gen {
	if serviceManagedResourcesSettings_STATUSARMGenerator != nil {
		return serviceManagedResourcesSettings_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_STATUSARM(generators)
	serviceManagedResourcesSettings_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServiceManagedResourcesSettings_STATUSARM{}), generators)

	return serviceManagedResourcesSettings_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_STATUSARM(gens map[string]gopter.Gen) {
	gens["CosmosDb"] = gen.PtrOf(CosmosDbSettings_STATUSARMGenerator())
}

func Test_SharedPrivateLinkResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResource_STATUSARM, SharedPrivateLinkResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResource_STATUSARM runs a test to see if a specific instance of SharedPrivateLinkResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResource_STATUSARM(subject SharedPrivateLinkResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResource_STATUSARM
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

// Generator of SharedPrivateLinkResource_STATUSARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResource_STATUSARMGenerator()
var sharedPrivateLinkResource_STATUSARMGenerator gopter.Gen

// SharedPrivateLinkResource_STATUSARMGenerator returns a generator of SharedPrivateLinkResource_STATUSARM instances for property testing.
// We first initialize sharedPrivateLinkResource_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResource_STATUSARMGenerator() gopter.Gen {
	if sharedPrivateLinkResource_STATUSARMGenerator != nil {
		return sharedPrivateLinkResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM(generators)
	sharedPrivateLinkResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM(generators)
	sharedPrivateLinkResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource_STATUSARM{}), generators)

	return sharedPrivateLinkResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedPrivateLinkResourceProperty_STATUSARMGenerator())
}

func Test_CosmosDbSettings_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CosmosDbSettings_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCosmosDbSettings_STATUSARM, CosmosDbSettings_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCosmosDbSettings_STATUSARM runs a test to see if a specific instance of CosmosDbSettings_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCosmosDbSettings_STATUSARM(subject CosmosDbSettings_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CosmosDbSettings_STATUSARM
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

// Generator of CosmosDbSettings_STATUSARM instances for property testing - lazily instantiated by
// CosmosDbSettings_STATUSARMGenerator()
var cosmosDbSettings_STATUSARMGenerator gopter.Gen

// CosmosDbSettings_STATUSARMGenerator returns a generator of CosmosDbSettings_STATUSARM instances for property testing.
func CosmosDbSettings_STATUSARMGenerator() gopter.Gen {
	if cosmosDbSettings_STATUSARMGenerator != nil {
		return cosmosDbSettings_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCosmosDbSettings_STATUSARM(generators)
	cosmosDbSettings_STATUSARMGenerator = gen.Struct(reflect.TypeOf(CosmosDbSettings_STATUSARM{}), generators)

	return cosmosDbSettings_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForCosmosDbSettings_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCosmosDbSettings_STATUSARM(gens map[string]gopter.Gen) {
	gens["CollectionsThroughput"] = gen.PtrOf(gen.Int())
}

func Test_IdentityForCmk_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityForCmk_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityForCmk_STATUSARM, IdentityForCmk_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityForCmk_STATUSARM runs a test to see if a specific instance of IdentityForCmk_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityForCmk_STATUSARM(subject IdentityForCmk_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityForCmk_STATUSARM
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

// Generator of IdentityForCmk_STATUSARM instances for property testing - lazily instantiated by
// IdentityForCmk_STATUSARMGenerator()
var identityForCmk_STATUSARMGenerator gopter.Gen

// IdentityForCmk_STATUSARMGenerator returns a generator of IdentityForCmk_STATUSARM instances for property testing.
func IdentityForCmk_STATUSARMGenerator() gopter.Gen {
	if identityForCmk_STATUSARMGenerator != nil {
		return identityForCmk_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityForCmk_STATUSARM(generators)
	identityForCmk_STATUSARMGenerator = gen.Struct(reflect.TypeOf(IdentityForCmk_STATUSARM{}), generators)

	return identityForCmk_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityForCmk_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityForCmk_STATUSARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_STATUSARM, KeyVaultProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_STATUSARM runs a test to see if a specific instance of KeyVaultProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_STATUSARM(subject KeyVaultProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_STATUSARM
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

// Generator of KeyVaultProperties_STATUSARM instances for property testing - lazily instantiated by
// KeyVaultProperties_STATUSARMGenerator()
var keyVaultProperties_STATUSARMGenerator gopter.Gen

// KeyVaultProperties_STATUSARMGenerator returns a generator of KeyVaultProperties_STATUSARM instances for property testing.
func KeyVaultProperties_STATUSARMGenerator() gopter.Gen {
	if keyVaultProperties_STATUSARMGenerator != nil {
		return keyVaultProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(generators)
	keyVaultProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUSARM{}), generators)

	return keyVaultProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultArmId"] = gen.PtrOf(gen.AlphaString())
}

func Test_NotebookPreparationError_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NotebookPreparationError_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNotebookPreparationError_STATUSARM, NotebookPreparationError_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNotebookPreparationError_STATUSARM runs a test to see if a specific instance of NotebookPreparationError_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNotebookPreparationError_STATUSARM(subject NotebookPreparationError_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NotebookPreparationError_STATUSARM
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

// Generator of NotebookPreparationError_STATUSARM instances for property testing - lazily instantiated by
// NotebookPreparationError_STATUSARMGenerator()
var notebookPreparationError_STATUSARMGenerator gopter.Gen

// NotebookPreparationError_STATUSARMGenerator returns a generator of NotebookPreparationError_STATUSARM instances for property testing.
func NotebookPreparationError_STATUSARMGenerator() gopter.Gen {
	if notebookPreparationError_STATUSARMGenerator != nil {
		return notebookPreparationError_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNotebookPreparationError_STATUSARM(generators)
	notebookPreparationError_STATUSARMGenerator = gen.Struct(reflect.TypeOf(NotebookPreparationError_STATUSARM{}), generators)

	return notebookPreparationError_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNotebookPreparationError_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNotebookPreparationError_STATUSARM(gens map[string]gopter.Gen) {
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
	gens["StatusCode"] = gen.PtrOf(gen.Int())
}

func Test_SharedPrivateLinkResourceProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperty_STATUSARM, SharedPrivateLinkResourceProperty_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperty_STATUSARM runs a test to see if a specific instance of SharedPrivateLinkResourceProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperty_STATUSARM(subject SharedPrivateLinkResourceProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperty_STATUSARM
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

// Generator of SharedPrivateLinkResourceProperty_STATUSARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceProperty_STATUSARMGenerator()
var sharedPrivateLinkResourceProperty_STATUSARMGenerator gopter.Gen

// SharedPrivateLinkResourceProperty_STATUSARMGenerator returns a generator of SharedPrivateLinkResourceProperty_STATUSARM instances for property testing.
func SharedPrivateLinkResourceProperty_STATUSARMGenerator() gopter.Gen {
	if sharedPrivateLinkResourceProperty_STATUSARMGenerator != nil {
		return sharedPrivateLinkResourceProperty_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_STATUSARM(generators)
	sharedPrivateLinkResourceProperty_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperty_STATUSARM{}), generators)

	return sharedPrivateLinkResourceProperty_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_STATUSARM(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointServiceConnectionStatus_STATUS_Approved,
		PrivateEndpointServiceConnectionStatus_STATUS_Disconnected,
		PrivateEndpointServiceConnectionStatus_STATUS_Pending,
		PrivateEndpointServiceConnectionStatus_STATUS_Rejected,
		PrivateEndpointServiceConnectionStatus_STATUS_Timeout))
}
