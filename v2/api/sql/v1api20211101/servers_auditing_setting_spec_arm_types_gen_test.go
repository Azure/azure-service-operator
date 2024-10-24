// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_ServerBlobAuditingPolicyProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerBlobAuditingPolicyProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerBlobAuditingPolicyProperties_ARM, ServerBlobAuditingPolicyProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerBlobAuditingPolicyProperties_ARM runs a test to see if a specific instance of ServerBlobAuditingPolicyProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerBlobAuditingPolicyProperties_ARM(subject ServerBlobAuditingPolicyProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerBlobAuditingPolicyProperties_ARM
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

// Generator of ServerBlobAuditingPolicyProperties_ARM instances for property testing - lazily instantiated by
// ServerBlobAuditingPolicyProperties_ARMGenerator()
var serverBlobAuditingPolicyProperties_ARMGenerator gopter.Gen

// ServerBlobAuditingPolicyProperties_ARMGenerator returns a generator of ServerBlobAuditingPolicyProperties_ARM instances for property testing.
func ServerBlobAuditingPolicyProperties_ARMGenerator() gopter.Gen {
	if serverBlobAuditingPolicyProperties_ARMGenerator != nil {
		return serverBlobAuditingPolicyProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_ARM(generators)
	serverBlobAuditingPolicyProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ServerBlobAuditingPolicyProperties_ARM{}), generators)

	return serverBlobAuditingPolicyProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_ARM(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsDevopsAuditEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerBlobAuditingPolicyProperties_State_ARM_Disabled, ServerBlobAuditingPolicyProperties_State_ARM_Enabled))
	gens["StorageAccountAccessKey"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersAuditingSetting_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAuditingSetting_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAuditingSetting_Spec_ARM, ServersAuditingSetting_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAuditingSetting_Spec_ARM runs a test to see if a specific instance of ServersAuditingSetting_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAuditingSetting_Spec_ARM(subject ServersAuditingSetting_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAuditingSetting_Spec_ARM
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

// Generator of ServersAuditingSetting_Spec_ARM instances for property testing - lazily instantiated by
// ServersAuditingSetting_Spec_ARMGenerator()
var serversAuditingSetting_Spec_ARMGenerator gopter.Gen

// ServersAuditingSetting_Spec_ARMGenerator returns a generator of ServersAuditingSetting_Spec_ARM instances for property testing.
// We first initialize serversAuditingSetting_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersAuditingSetting_Spec_ARMGenerator() gopter.Gen {
	if serversAuditingSetting_Spec_ARMGenerator != nil {
		return serversAuditingSetting_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAuditingSetting_Spec_ARM(generators)
	serversAuditingSetting_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ServersAuditingSetting_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAuditingSetting_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServersAuditingSetting_Spec_ARM(generators)
	serversAuditingSetting_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ServersAuditingSetting_Spec_ARM{}), generators)

	return serversAuditingSetting_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServersAuditingSetting_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersAuditingSetting_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersAuditingSetting_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAuditingSetting_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerBlobAuditingPolicyProperties_ARMGenerator())
}
