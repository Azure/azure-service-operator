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

func Test_Workspace_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceStatusARM, WorkspaceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceStatusARM runs a test to see if a specific instance of Workspace_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceStatusARM(subject Workspace_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_StatusARM
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

// Generator of Workspace_StatusARM instances for property testing - lazily instantiated by WorkspaceStatusARMGenerator()
var workspaceStatusARMGenerator gopter.Gen

// WorkspaceStatusARMGenerator returns a generator of Workspace_StatusARM instances for property testing.
// We first initialize workspaceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceStatusARMGenerator() gopter.Gen {
	if workspaceStatusARMGenerator != nil {
		return workspaceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceStatusARM(generators)
	workspaceStatusARMGenerator = gen.Struct(reflect.TypeOf(Workspace_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceStatusARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceStatusARM(generators)
	workspaceStatusARMGenerator = gen.Struct(reflect.TypeOf(Workspace_StatusARM{}), generators)

	return workspaceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspacePropertiesStatusARMGenerator())
}

func Test_WorkspaceProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacePropertiesStatusARM, WorkspacePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacePropertiesStatusARM runs a test to see if a specific instance of WorkspaceProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacePropertiesStatusARM(subject WorkspaceProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_StatusARM
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

// Generator of WorkspaceProperties_StatusARM instances for property testing - lazily instantiated by
// WorkspacePropertiesStatusARMGenerator()
var workspacePropertiesStatusARMGenerator gopter.Gen

// WorkspacePropertiesStatusARMGenerator returns a generator of WorkspaceProperties_StatusARM instances for property testing.
// We first initialize workspacePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacePropertiesStatusARMGenerator() gopter.Gen {
	if workspacePropertiesStatusARMGenerator != nil {
		return workspacePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacePropertiesStatusARM(generators)
	workspacePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForWorkspacePropertiesStatusARM(generators)
	workspacePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_StatusARM{}), generators)

	return workspacePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedDate"] = gen.PtrOf(gen.AlphaString())
	gens["CustomerId"] = gen.PtrOf(gen.AlphaString())
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ModifiedDate"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspacePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeaturesStatusARMGenerator())
	gens["PrivateLinkScopedResources"] = gen.SliceOf(PrivateLinkScopedResourceStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSkuStatusARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCappingStatusARMGenerator())
}

func Test_PrivateLinkScopedResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkScopedResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkScopedResourceStatusARM, PrivateLinkScopedResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkScopedResourceStatusARM runs a test to see if a specific instance of PrivateLinkScopedResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkScopedResourceStatusARM(subject PrivateLinkScopedResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkScopedResource_StatusARM
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

// Generator of PrivateLinkScopedResource_StatusARM instances for property testing - lazily instantiated by
// PrivateLinkScopedResourceStatusARMGenerator()
var privateLinkScopedResourceStatusARMGenerator gopter.Gen

// PrivateLinkScopedResourceStatusARMGenerator returns a generator of PrivateLinkScopedResource_StatusARM instances for property testing.
func PrivateLinkScopedResourceStatusARMGenerator() gopter.Gen {
	if privateLinkScopedResourceStatusARMGenerator != nil {
		return privateLinkScopedResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatusARM(generators)
	privateLinkScopedResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkScopedResource_StatusARM{}), generators)

	return privateLinkScopedResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkScopedResourceStatusARM(gens map[string]gopter.Gen) {
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ScopeId"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceCapping_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCappingStatusARM, WorkspaceCappingStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCappingStatusARM runs a test to see if a specific instance of WorkspaceCapping_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCappingStatusARM(subject WorkspaceCapping_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_StatusARM
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

// Generator of WorkspaceCapping_StatusARM instances for property testing - lazily instantiated by
// WorkspaceCappingStatusARMGenerator()
var workspaceCappingStatusARMGenerator gopter.Gen

// WorkspaceCappingStatusARMGenerator returns a generator of WorkspaceCapping_StatusARM instances for property testing.
func WorkspaceCappingStatusARMGenerator() gopter.Gen {
	if workspaceCappingStatusARMGenerator != nil {
		return workspaceCappingStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCappingStatusARM(generators)
	workspaceCappingStatusARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_StatusARM{}), generators)

	return workspaceCappingStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCappingStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCappingStatusARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
	gens["DataIngestionStatus"] = gen.PtrOf(gen.AlphaString())
	gens["QuotaNextResetTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceFeatures_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeaturesStatusARM, WorkspaceFeaturesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeaturesStatusARM runs a test to see if a specific instance of WorkspaceFeatures_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeaturesStatusARM(subject WorkspaceFeatures_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_StatusARM
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

// Generator of WorkspaceFeatures_StatusARM instances for property testing - lazily instantiated by
// WorkspaceFeaturesStatusARMGenerator()
var workspaceFeaturesStatusARMGenerator gopter.Gen

// WorkspaceFeaturesStatusARMGenerator returns a generator of WorkspaceFeatures_StatusARM instances for property testing.
func WorkspaceFeaturesStatusARMGenerator() gopter.Gen {
	if workspaceFeaturesStatusARMGenerator != nil {
		return workspaceFeaturesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatusARM(generators)
	workspaceFeaturesStatusARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_StatusARM{}), generators)

	return workspaceFeaturesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeaturesStatusARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSkuStatusARM, WorkspaceSkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSkuStatusARM runs a test to see if a specific instance of WorkspaceSku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSkuStatusARM(subject WorkspaceSku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_StatusARM
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

// Generator of WorkspaceSku_StatusARM instances for property testing - lazily instantiated by
// WorkspaceSkuStatusARMGenerator()
var workspaceSkuStatusARMGenerator gopter.Gen

// WorkspaceSkuStatusARMGenerator returns a generator of WorkspaceSku_StatusARM instances for property testing.
func WorkspaceSkuStatusARMGenerator() gopter.Gen {
	if workspaceSkuStatusARMGenerator != nil {
		return workspaceSkuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSkuStatusARM(generators)
	workspaceSkuStatusARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_StatusARM{}), generators)

	return workspaceSkuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSkuStatusARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.Int())
	gens["LastSkuUpdate"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}
