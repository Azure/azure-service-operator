// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

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

func Test_ServersAuditingSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAuditingSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAuditingSetting, ServersAuditingSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAuditingSetting runs a test to see if a specific instance of ServersAuditingSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAuditingSetting(subject ServersAuditingSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAuditingSetting
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

// Generator of ServersAuditingSetting instances for property testing - lazily instantiated by
// ServersAuditingSettingGenerator()
var serversAuditingSettingGenerator gopter.Gen

// ServersAuditingSettingGenerator returns a generator of ServersAuditingSetting instances for property testing.
func ServersAuditingSettingGenerator() gopter.Gen {
	if serversAuditingSettingGenerator != nil {
		return serversAuditingSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersAuditingSetting(generators)
	serversAuditingSettingGenerator = gen.Struct(reflect.TypeOf(ServersAuditingSetting{}), generators)

	return serversAuditingSettingGenerator
}

// AddRelatedPropertyGeneratorsForServersAuditingSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAuditingSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_AuditingSetting_SpecGenerator()
	gens["Status"] = Servers_AuditingSetting_STATUSGenerator()
}

func Test_Servers_AuditingSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_AuditingSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_AuditingSetting_Spec, Servers_AuditingSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_AuditingSetting_Spec runs a test to see if a specific instance of Servers_AuditingSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_AuditingSetting_Spec(subject Servers_AuditingSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_AuditingSetting_Spec
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

// Generator of Servers_AuditingSetting_Spec instances for property testing - lazily instantiated by
// Servers_AuditingSetting_SpecGenerator()
var servers_AuditingSetting_SpecGenerator gopter.Gen

// Servers_AuditingSetting_SpecGenerator returns a generator of Servers_AuditingSetting_Spec instances for property testing.
func Servers_AuditingSetting_SpecGenerator() gopter.Gen {
	if servers_AuditingSetting_SpecGenerator != nil {
		return servers_AuditingSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AuditingSetting_Spec(generators)
	servers_AuditingSetting_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_AuditingSetting_Spec{}), generators)

	return servers_AuditingSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_AuditingSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_AuditingSetting_Spec(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsDevopsAuditEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_AuditingSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_AuditingSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_AuditingSetting_STATUS, Servers_AuditingSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_AuditingSetting_STATUS runs a test to see if a specific instance of Servers_AuditingSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_AuditingSetting_STATUS(subject Servers_AuditingSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_AuditingSetting_STATUS
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

// Generator of Servers_AuditingSetting_STATUS instances for property testing - lazily instantiated by
// Servers_AuditingSetting_STATUSGenerator()
var servers_AuditingSetting_STATUSGenerator gopter.Gen

// Servers_AuditingSetting_STATUSGenerator returns a generator of Servers_AuditingSetting_STATUS instances for property testing.
func Servers_AuditingSetting_STATUSGenerator() gopter.Gen {
	if servers_AuditingSetting_STATUSGenerator != nil {
		return servers_AuditingSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_AuditingSetting_STATUS(generators)
	servers_AuditingSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_AuditingSetting_STATUS{}), generators)

	return servers_AuditingSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_AuditingSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_AuditingSetting_STATUS(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsDevopsAuditEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
