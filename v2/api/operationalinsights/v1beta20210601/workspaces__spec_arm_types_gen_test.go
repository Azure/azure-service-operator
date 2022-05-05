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

func Test_Workspaces_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesSpecARM, WorkspacesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesSpecARM runs a test to see if a specific instance of Workspaces_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesSpecARM(subject Workspaces_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_SpecARM
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

// Generator of Workspaces_SpecARM instances for property testing - lazily instantiated by WorkspacesSpecARMGenerator()
var workspacesSpecARMGenerator gopter.Gen

// WorkspacesSpecARMGenerator returns a generator of Workspaces_SpecARM instances for property testing.
// We first initialize workspacesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacesSpecARMGenerator() gopter.Gen {
	if workspacesSpecARMGenerator != nil {
		return workspacesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesSpecARM(generators)
	workspacesSpecARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesSpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspacesSpecARM(generators)
	workspacesSpecARMGenerator = gen.Struct(reflect.TypeOf(Workspaces_SpecARM{}), generators)

	return workspacesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacesSpecARM(gens map[string]gopter.Gen) {
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspacesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspacePropertiesARMGenerator())
}

func Test_WorkspacePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		WorkspacePropertiesProvisioningStateCanceled,
		WorkspacePropertiesProvisioningStateCreating,
		WorkspacePropertiesProvisioningStateDeleting,
		WorkspacePropertiesProvisioningStateFailed,
		WorkspacePropertiesProvisioningStateProvisioningAccount,
		WorkspacePropertiesProvisioningStateSucceeded,
		WorkspacePropertiesProvisioningStateUpdating))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(WorkspacePropertiesPublicNetworkAccessForIngestionDisabled, WorkspacePropertiesPublicNetworkAccessForIngestionEnabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(WorkspacePropertiesPublicNetworkAccessForQueryDisabled, WorkspacePropertiesPublicNetworkAccessForQueryEnabled))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspacePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacePropertiesARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeaturesARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSkuARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCappingARMGenerator())
}

func Test_WorkspaceCappingARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCappingARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCappingARM, WorkspaceCappingARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCappingARM runs a test to see if a specific instance of WorkspaceCappingARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCappingARM(subject WorkspaceCappingARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCappingARM
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

// Generator of WorkspaceCappingARM instances for property testing - lazily instantiated by
// WorkspaceCappingARMGenerator()
var workspaceCappingARMGenerator gopter.Gen

// WorkspaceCappingARMGenerator returns a generator of WorkspaceCappingARM instances for property testing.
func WorkspaceCappingARMGenerator() gopter.Gen {
	if workspaceCappingARMGenerator != nil {
		return workspaceCappingARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCappingARM(generators)
	workspaceCappingARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCappingARM{}), generators)

	return workspaceCappingARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCappingARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCappingARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
}

func Test_WorkspaceFeaturesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeaturesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeaturesARM, WorkspaceFeaturesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeaturesARM runs a test to see if a specific instance of WorkspaceFeaturesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeaturesARM(subject WorkspaceFeaturesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeaturesARM
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

// Generator of WorkspaceFeaturesARM instances for property testing - lazily instantiated by
// WorkspaceFeaturesARMGenerator()
var workspaceFeaturesARMGenerator gopter.Gen

// WorkspaceFeaturesARMGenerator returns a generator of WorkspaceFeaturesARM instances for property testing.
func WorkspaceFeaturesARMGenerator() gopter.Gen {
	if workspaceFeaturesARMGenerator != nil {
		return workspaceFeaturesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeaturesARM(generators)
	workspaceFeaturesARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeaturesARM{}), generators)

	return workspaceFeaturesARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeaturesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeaturesARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSkuARM, WorkspaceSkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSkuARM runs a test to see if a specific instance of WorkspaceSkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSkuARM(subject WorkspaceSkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSkuARM
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

// Generator of WorkspaceSkuARM instances for property testing - lazily instantiated by WorkspaceSkuARMGenerator()
var workspaceSkuARMGenerator gopter.Gen

// WorkspaceSkuARMGenerator returns a generator of WorkspaceSkuARM instances for property testing.
func WorkspaceSkuARMGenerator() gopter.Gen {
	if workspaceSkuARMGenerator != nil {
		return workspaceSkuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSkuARM(generators)
	workspaceSkuARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSkuARM{}), generators)

	return workspaceSkuARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSkuARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSkuNameCapacityReservation,
		WorkspaceSkuNameFree,
		WorkspaceSkuNameLACluster,
		WorkspaceSkuNamePerGB2018,
		WorkspaceSkuNamePerNode,
		WorkspaceSkuNamePremium,
		WorkspaceSkuNameStandalone,
		WorkspaceSkuNameStandard))
}
