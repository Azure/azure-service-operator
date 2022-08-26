// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

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
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_STATUSARMGenerator())
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
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["CustomerId"] = gen.PtrOf(gen.AlphaString())
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ModifiedDate"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
<<<<<<< HEAD
		WorkspaceProperties_ProvisioningState_Canceled_STATUS,
		WorkspaceProperties_ProvisioningState_Creating_STATUS,
		WorkspaceProperties_ProvisioningState_Deleting_STATUS,
		WorkspaceProperties_ProvisioningState_Failed_STATUS,
		WorkspaceProperties_ProvisioningState_ProvisioningAccount_STATUS,
		WorkspaceProperties_ProvisioningState_Succeeded_STATUS,
		WorkspaceProperties_ProvisioningState_Updating_STATUS))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled_STATUS, PublicNetworkAccessType_Enabled_STATUS))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled_STATUS, PublicNetworkAccessType_Enabled_STATUS))
=======
		WorkspaceProperties_STATUS_ProvisioningState_Canceled,
		WorkspaceProperties_STATUS_ProvisioningState_Creating,
		WorkspaceProperties_STATUS_ProvisioningState_Deleting,
		WorkspaceProperties_STATUS_ProvisioningState_Failed,
		WorkspaceProperties_STATUS_ProvisioningState_ProvisioningAccount,
		WorkspaceProperties_STATUS_ProvisioningState_Succeeded,
		WorkspaceProperties_STATUS_ProvisioningState_Updating))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_STATUS_Disabled, PublicNetworkAccessType_STATUS_Enabled))
>>>>>>> main
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeatures_STATUSARMGenerator())
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResource_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSku_STATUSARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCapping_STATUSARMGenerator())
}

func Test_PrivateLinkScopedResource_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResource_STATUSARM, PrivateLinkScopedResource_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResource_STATUSARM runs a test to see if a specific instance of PrivateLinkScopedResource_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResource_STATUSARM(subject PrivateLinkScopedResource_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_STATUSARM
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

// Generator of PrivateLinkScopedResource_STATUSARM instances for property testing - lazily instantiated by
// PrivateLinkScopedResource_STATUSARMGenerator()
var privateLinkScopedResource_STATUSARMGenerator gopter.Gen

// PrivateLinkScopedResource_STATUSARMGenerator returns a generator of PrivateLinkScopedResource_STATUSARM instances for property testing.
func PrivateLinkScopedResource_STATUSARMGenerator() gopter.Gen {
	if privateLinkScopedResource_STATUSARMGenerator != nil {
		return privateLinkScopedResource_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUSARM(generators)
	privateLinkScopedResource_STATUSARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_STATUSARM{}), generators)

	return privateLinkScopedResource_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResource_STATUSARM(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceCapping_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCapping_STATUSARM, WorkspaceCapping_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCapping_STATUSARM runs a test to see if a specific instance of WorkspaceCapping_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCapping_STATUSARM(subject WorkspaceCapping_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_STATUSARM
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

// Generator of WorkspaceCapping_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceCapping_STATUSARMGenerator()
var workspaceCapping_STATUSARMGenerator gopter.Gen

// WorkspaceCapping_STATUSARMGenerator returns a generator of WorkspaceCapping_STATUSARM instances for property testing.
func WorkspaceCapping_STATUSARMGenerator() gopter.Gen {
	if workspaceCapping_STATUSARMGenerator != nil {
		return workspaceCapping_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUSARM(generators)
	workspaceCapping_STATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_STATUSARM{}), generators)

	return workspaceCapping_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCapping_STATUSARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
	gens["DataIngestionStatus"] = gen.PtrOf(gen.OneConstOf(
<<<<<<< HEAD
		WorkspaceCapping_DataIngestionStatus_ApproachingQuota_STATUS,
		WorkspaceCapping_DataIngestionStatus_ForceOff_STATUS,
		WorkspaceCapping_DataIngestionStatus_ForceOn_STATUS,
		WorkspaceCapping_DataIngestionStatus_OverQuota_STATUS,
		WorkspaceCapping_DataIngestionStatus_RespectQuota_STATUS,
		WorkspaceCapping_DataIngestionStatus_SubscriptionSuspended_STATUS))
=======
		WorkspaceCapping_STATUS_DataIngestionStatus_ApproachingQuota,
		WorkspaceCapping_STATUS_DataIngestionStatus_ForceOff,
		WorkspaceCapping_STATUS_DataIngestionStatus_ForceOn,
		WorkspaceCapping_STATUS_DataIngestionStatus_OverQuota,
		WorkspaceCapping_STATUS_DataIngestionStatus_RespectQuota,
		WorkspaceCapping_STATUS_DataIngestionStatus_SubscriptionSuspended))
>>>>>>> main
	gens["QuotaNextResetTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceFeatures_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeatures_STATUSARM, WorkspaceFeatures_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeatures_STATUSARM runs a test to see if a specific instance of WorkspaceFeatures_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeatures_STATUSARM(subject WorkspaceFeatures_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_STATUSARM
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

// Generator of WorkspaceFeatures_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceFeatures_STATUSARMGenerator()
var workspaceFeatures_STATUSARMGenerator gopter.Gen

// WorkspaceFeatures_STATUSARMGenerator returns a generator of WorkspaceFeatures_STATUSARM instances for property testing.
func WorkspaceFeatures_STATUSARMGenerator() gopter.Gen {
	if workspaceFeatures_STATUSARMGenerator != nil {
		return workspaceFeatures_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUSARM(generators)
	workspaceFeatures_STATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_STATUSARM{}), generators)

	return workspaceFeatures_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeatures_STATUSARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSku_STATUSARM, WorkspaceSku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSku_STATUSARM runs a test to see if a specific instance of WorkspaceSku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSku_STATUSARM(subject WorkspaceSku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_STATUSARM
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

// Generator of WorkspaceSku_STATUSARM instances for property testing - lazily instantiated by
// WorkspaceSku_STATUSARMGenerator()
var workspaceSku_STATUSARMGenerator gopter.Gen

// WorkspaceSku_STATUSARMGenerator returns a generator of WorkspaceSku_STATUSARM instances for property testing.
func WorkspaceSku_STATUSARMGenerator() gopter.Gen {
	if workspaceSku_STATUSARMGenerator != nil {
		return workspaceSku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSku_STATUSARM(generators)
	workspaceSku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_STATUSARM{}), generators)

	return workspaceSku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.OneConstOf(
<<<<<<< HEAD
		WorkspaceSku_CapacityReservationLevel_100_STATUS,
		WorkspaceSku_CapacityReservationLevel_1000_STATUS,
		WorkspaceSku_CapacityReservationLevel_200_STATUS,
		WorkspaceSku_CapacityReservationLevel_2000_STATUS,
		WorkspaceSku_CapacityReservationLevel_300_STATUS,
		WorkspaceSku_CapacityReservationLevel_400_STATUS,
		WorkspaceSku_CapacityReservationLevel_500_STATUS,
		WorkspaceSku_CapacityReservationLevel_5000_STATUS))
	gens["LastSkuUpdate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_Name_CapacityReservation_STATUS,
		WorkspaceSku_Name_Free_STATUS,
		WorkspaceSku_Name_LACluster_STATUS,
		WorkspaceSku_Name_PerGB2018_STATUS,
		WorkspaceSku_Name_PerNode_STATUS,
		WorkspaceSku_Name_Premium_STATUS,
		WorkspaceSku_Name_Standalone_STATUS,
		WorkspaceSku_Name_Standard_STATUS))
=======
		WorkspaceSku_STATUS_CapacityReservationLevel_100,
		WorkspaceSku_STATUS_CapacityReservationLevel_1000,
		WorkspaceSku_STATUS_CapacityReservationLevel_200,
		WorkspaceSku_STATUS_CapacityReservationLevel_2000,
		WorkspaceSku_STATUS_CapacityReservationLevel_300,
		WorkspaceSku_STATUS_CapacityReservationLevel_400,
		WorkspaceSku_STATUS_CapacityReservationLevel_500,
		WorkspaceSku_STATUS_CapacityReservationLevel_5000))
	gens["LastSkuUpdate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_STATUS_Name_CapacityReservation,
		WorkspaceSku_STATUS_Name_Free,
		WorkspaceSku_STATUS_Name_LACluster,
		WorkspaceSku_STATUS_Name_PerGB2018,
		WorkspaceSku_STATUS_Name_PerNode,
		WorkspaceSku_STATUS_Name_Premium,
		WorkspaceSku_STATUS_Name_Standalone,
		WorkspaceSku_STATUS_Name_Standard))
>>>>>>> main
}
