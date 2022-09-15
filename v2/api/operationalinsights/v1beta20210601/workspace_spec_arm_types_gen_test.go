// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_Workspace_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_Spec_ARM, Workspace_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_Spec_ARM runs a test to see if a specific instance of Workspace_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_Spec_ARM(subject Workspace_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_Spec_ARM
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

// Generator of Workspace_Spec_ARM instances for property testing - lazily instantiated by Workspace_Spec_ARMGenerator()
var workspace_Spec_ARMGenerator gopter.Gen

// Workspace_Spec_ARMGenerator returns a generator of Workspace_Spec_ARM instances for property testing.
// We first initialize workspace_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_Spec_ARMGenerator() gopter.Gen {
	if workspace_Spec_ARMGenerator != nil {
		return workspace_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	workspace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	workspace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec_ARM{}), generators)

	return workspace_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_ARMGenerator())
}

func Test_WorkspaceProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties_ARM, WorkspaceProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties_ARM runs a test to see if a specific instance of WorkspaceProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties_ARM(subject WorkspaceProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_ARM
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

// Generator of WorkspaceProperties_ARM instances for property testing - lazily instantiated by
// WorkspaceProperties_ARMGenerator()
var workspaceProperties_ARMGenerator gopter.Gen

// WorkspaceProperties_ARMGenerator returns a generator of WorkspaceProperties_ARM instances for property testing.
// We first initialize workspaceProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceProperties_ARMGenerator() gopter.Gen {
	if workspaceProperties_ARMGenerator != nil {
		return workspaceProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	workspaceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	workspaceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_ARM{}), generators)

	return workspaceProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(gens map[string]gopter.Gen) {
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceProperties_ProvisioningState_Canceled,
		WorkspaceProperties_ProvisioningState_Creating,
		WorkspaceProperties_ProvisioningState_Deleting,
		WorkspaceProperties_ProvisioningState_Failed,
		WorkspaceProperties_ProvisioningState_ProvisioningAccount,
		WorkspaceProperties_ProvisioningState_Succeeded,
		WorkspaceProperties_ProvisioningState_Updating))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccessForIngestion_Disabled, WorkspaceProperties_PublicNetworkAccessForIngestion_Enabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccessForQuery_Disabled, WorkspaceProperties_PublicNetworkAccessForQuery_Enabled))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeatures_ARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSku_ARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCapping_ARMGenerator())
}

func Test_WorkspaceCapping_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCapping_ARM, WorkspaceCapping_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCapping_ARM runs a test to see if a specific instance of WorkspaceCapping_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCapping_ARM(subject WorkspaceCapping_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_ARM
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

// Generator of WorkspaceCapping_ARM instances for property testing - lazily instantiated by
// WorkspaceCapping_ARMGenerator()
var workspaceCapping_ARMGenerator gopter.Gen

// WorkspaceCapping_ARMGenerator returns a generator of WorkspaceCapping_ARM instances for property testing.
func WorkspaceCapping_ARMGenerator() gopter.Gen {
	if workspaceCapping_ARMGenerator != nil {
		return workspaceCapping_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCapping_ARM(generators)
	workspaceCapping_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_ARM{}), generators)

	return workspaceCapping_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCapping_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCapping_ARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
}

func Test_WorkspaceFeatures_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeatures_ARM, WorkspaceFeatures_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeatures_ARM runs a test to see if a specific instance of WorkspaceFeatures_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeatures_ARM(subject WorkspaceFeatures_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_ARM
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

// Generator of WorkspaceFeatures_ARM instances for property testing - lazily instantiated by
// WorkspaceFeatures_ARMGenerator()
var workspaceFeatures_ARMGenerator gopter.Gen

// WorkspaceFeatures_ARMGenerator returns a generator of WorkspaceFeatures_ARM instances for property testing.
func WorkspaceFeatures_ARMGenerator() gopter.Gen {
	if workspaceFeatures_ARMGenerator != nil {
		return workspaceFeatures_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeatures_ARM(generators)
	workspaceFeatures_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_ARM{}), generators)

	return workspaceFeatures_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeatures_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeatures_ARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSku_ARM, WorkspaceSku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSku_ARM runs a test to see if a specific instance of WorkspaceSku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSku_ARM(subject WorkspaceSku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_ARM
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

// Generator of WorkspaceSku_ARM instances for property testing - lazily instantiated by WorkspaceSku_ARMGenerator()
var workspaceSku_ARMGenerator gopter.Gen

// WorkspaceSku_ARMGenerator returns a generator of WorkspaceSku_ARM instances for property testing.
func WorkspaceSku_ARMGenerator() gopter.Gen {
	if workspaceSku_ARMGenerator != nil {
		return workspaceSku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSku_ARM(generators)
	workspaceSku_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_ARM{}), generators)

	return workspaceSku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSku_ARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_Name_CapacityReservation,
		WorkspaceSku_Name_Free,
		WorkspaceSku_Name_LACluster,
		WorkspaceSku_Name_PerGB2018,
		WorkspaceSku_Name_PerNode,
		WorkspaceSku_Name_Premium,
		WorkspaceSku_Name_Standalone,
		WorkspaceSku_Name_Standard))
}
