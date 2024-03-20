// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

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

func Test_Workspace_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_STATUS_ARM, Workspace_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_STATUS_ARM runs a test to see if a specific instance of Workspace_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_STATUS_ARM(subject Workspace_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_STATUS_ARM
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

// Generator of Workspace_STATUS_ARM instances for property testing - lazily instantiated by
// Workspace_STATUS_ARMGenerator()
var workspace_STATUS_ARMGenerator gopter.Gen

// Workspace_STATUS_ARMGenerator returns a generator of Workspace_STATUS_ARM instances for property testing.
// We first initialize workspace_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_STATUS_ARMGenerator() gopter.Gen {
	if workspace_STATUS_ARMGenerator != nil {
		return workspace_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_STATUS_ARM(generators)
	workspace_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspace_STATUS_ARM(generators)
	workspace_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_STATUS_ARM{}), generators)

	return workspace_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_STATUS_ARMGenerator())
}

func Test_WorkspaceProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties_STATUS_ARM, WorkspaceProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties_STATUS_ARM runs a test to see if a specific instance of WorkspaceProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties_STATUS_ARM(subject WorkspaceProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_STATUS_ARM
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

// Generator of WorkspaceProperties_STATUS_ARM instances for property testing - lazily instantiated by
// WorkspaceProperties_STATUS_ARMGenerator()
var workspaceProperties_STATUS_ARMGenerator gopter.Gen

// WorkspaceProperties_STATUS_ARMGenerator returns a generator of WorkspaceProperties_STATUS_ARM instances for property testing.
// We first initialize workspaceProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceProperties_STATUS_ARMGenerator() gopter.Gen {
	if workspaceProperties_STATUS_ARMGenerator != nil {
		return workspaceProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUS_ARM(generators)
	workspaceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUS_ARM(generators)
	workspaceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_STATUS_ARM{}), generators)

	return workspaceProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["CustomerId"] = gen.PtrOf(gen.AlphaString())
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ModifiedDate"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceProperties_ProvisioningState_STATUS_Canceled,
		WorkspaceProperties_ProvisioningState_STATUS_Creating,
		WorkspaceProperties_ProvisioningState_STATUS_Deleting,
		WorkspaceProperties_ProvisioningState_STATUS_Failed,
		WorkspaceProperties_ProvisioningState_STATUS_ProvisioningAccount,
		WorkspaceProperties_ProvisioningState_STATUS_Succeeded,
		WorkspaceProperties_ProvisioningState_STATUS_Updating))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeatures_STATUS_ARMGenerator())
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResource_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSku_STATUS_ARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCapping_STATUS_ARMGenerator())
}

func Test_PrivateLinkScopedResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM, PrivateLinkScopedResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM runs a test to see if a specific instance of PrivateLinkScopedResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResource_STATUS_ARM(subject PrivateLinkScopedResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_STATUS_ARM
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

// Generator of PrivateLinkScopedResource_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateLinkScopedResource_STATUS_ARMGenerator()
var privateLinkScopedResource_STATUS_ARMGenerator gopter.Gen

// PrivateLinkScopedResource_STATUS_ARMGenerator returns a generator of PrivateLinkScopedResource_STATUS_ARM instances for property testing.
func PrivateLinkScopedResource_STATUS_ARMGenerator() gopter.Gen {
	if privateLinkScopedResource_STATUS_ARMGenerator != nil {
		return privateLinkScopedResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM(generators)
	privateLinkScopedResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_STATUS_ARM{}), generators)

	return privateLinkScopedResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceCapping_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCapping_STATUS_ARM, WorkspaceCapping_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCapping_STATUS_ARM runs a test to see if a specific instance of WorkspaceCapping_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCapping_STATUS_ARM(subject WorkspaceCapping_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_STATUS_ARM
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

// Generator of WorkspaceCapping_STATUS_ARM instances for property testing - lazily instantiated by
// WorkspaceCapping_STATUS_ARMGenerator()
var workspaceCapping_STATUS_ARMGenerator gopter.Gen

// WorkspaceCapping_STATUS_ARMGenerator returns a generator of WorkspaceCapping_STATUS_ARM instances for property testing.
func WorkspaceCapping_STATUS_ARMGenerator() gopter.Gen {
	if workspaceCapping_STATUS_ARMGenerator != nil {
		return workspaceCapping_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUS_ARM(generators)
	workspaceCapping_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_STATUS_ARM{}), generators)

	return workspaceCapping_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
	gens["DataIngestionStatus"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceCapping_DataIngestionStatus_STATUS_ApproachingQuota,
		WorkspaceCapping_DataIngestionStatus_STATUS_ForceOff,
		WorkspaceCapping_DataIngestionStatus_STATUS_ForceOn,
		WorkspaceCapping_DataIngestionStatus_STATUS_OverQuota,
		WorkspaceCapping_DataIngestionStatus_STATUS_RespectQuota,
		WorkspaceCapping_DataIngestionStatus_STATUS_SubscriptionSuspended))
	gens["QuotaNextResetTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceFeatures_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeatures_STATUS_ARM, WorkspaceFeatures_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeatures_STATUS_ARM runs a test to see if a specific instance of WorkspaceFeatures_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeatures_STATUS_ARM(subject WorkspaceFeatures_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_STATUS_ARM
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

// Generator of WorkspaceFeatures_STATUS_ARM instances for property testing - lazily instantiated by
// WorkspaceFeatures_STATUS_ARMGenerator()
var workspaceFeatures_STATUS_ARMGenerator gopter.Gen

// WorkspaceFeatures_STATUS_ARMGenerator returns a generator of WorkspaceFeatures_STATUS_ARM instances for property testing.
func WorkspaceFeatures_STATUS_ARMGenerator() gopter.Gen {
	if workspaceFeatures_STATUS_ARMGenerator != nil {
		return workspaceFeatures_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUS_ARM(generators)
	workspaceFeatures_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_STATUS_ARM{}), generators)

	return workspaceFeatures_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSku_STATUS_ARM, WorkspaceSku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSku_STATUS_ARM runs a test to see if a specific instance of WorkspaceSku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSku_STATUS_ARM(subject WorkspaceSku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_STATUS_ARM
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

// Generator of WorkspaceSku_STATUS_ARM instances for property testing - lazily instantiated by
// WorkspaceSku_STATUS_ARMGenerator()
var workspaceSku_STATUS_ARMGenerator gopter.Gen

// WorkspaceSku_STATUS_ARMGenerator returns a generator of WorkspaceSku_STATUS_ARM instances for property testing.
func WorkspaceSku_STATUS_ARMGenerator() gopter.Gen {
	if workspaceSku_STATUS_ARMGenerator != nil {
		return workspaceSku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSku_STATUS_ARM(generators)
	workspaceSku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_STATUS_ARM{}), generators)

	return workspaceSku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_CapacityReservationLevel_STATUS_100,
		WorkspaceSku_CapacityReservationLevel_STATUS_200,
		WorkspaceSku_CapacityReservationLevel_STATUS_300,
		WorkspaceSku_CapacityReservationLevel_STATUS_400,
		WorkspaceSku_CapacityReservationLevel_STATUS_500,
		WorkspaceSku_CapacityReservationLevel_STATUS_1000,
		WorkspaceSku_CapacityReservationLevel_STATUS_2000,
		WorkspaceSku_CapacityReservationLevel_STATUS_5000))
	gens["LastSkuUpdate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_Name_STATUS_CapacityReservation,
		WorkspaceSku_Name_STATUS_Free,
		WorkspaceSku_Name_STATUS_LACluster,
		WorkspaceSku_Name_STATUS_PerGB2018,
		WorkspaceSku_Name_STATUS_PerNode,
		WorkspaceSku_Name_STATUS_Premium,
		WorkspaceSku_Name_STATUS_Standalone,
		WorkspaceSku_Name_STATUS_Standard))
}
