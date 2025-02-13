// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_ServerBlobAuditingPolicyProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerBlobAuditingPolicyProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerBlobAuditingPolicyProperties_STATUS, ServerBlobAuditingPolicyProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerBlobAuditingPolicyProperties_STATUS runs a test to see if a specific instance of ServerBlobAuditingPolicyProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServerBlobAuditingPolicyProperties_STATUS(subject ServerBlobAuditingPolicyProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerBlobAuditingPolicyProperties_STATUS
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

// Generator of ServerBlobAuditingPolicyProperties_STATUS instances for property testing - lazily instantiated by
// ServerBlobAuditingPolicyProperties_STATUSGenerator()
var serverBlobAuditingPolicyProperties_STATUSGenerator gopter.Gen

// ServerBlobAuditingPolicyProperties_STATUSGenerator returns a generator of ServerBlobAuditingPolicyProperties_STATUS instances for property testing.
func ServerBlobAuditingPolicyProperties_STATUSGenerator() gopter.Gen {
	if serverBlobAuditingPolicyProperties_STATUSGenerator != nil {
		return serverBlobAuditingPolicyProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_STATUS(generators)
	serverBlobAuditingPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerBlobAuditingPolicyProperties_STATUS{}), generators)

	return serverBlobAuditingPolicyProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerBlobAuditingPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsDevopsAuditEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerBlobAuditingPolicyProperties_State_STATUS_Disabled, ServerBlobAuditingPolicyProperties_State_STATUS_Enabled))
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersAuditingSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAuditingSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAuditingSetting_STATUS, ServersAuditingSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAuditingSetting_STATUS runs a test to see if a specific instance of ServersAuditingSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAuditingSetting_STATUS(subject ServersAuditingSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAuditingSetting_STATUS
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

// Generator of ServersAuditingSetting_STATUS instances for property testing - lazily instantiated by
// ServersAuditingSetting_STATUSGenerator()
var serversAuditingSetting_STATUSGenerator gopter.Gen

// ServersAuditingSetting_STATUSGenerator returns a generator of ServersAuditingSetting_STATUS instances for property testing.
// We first initialize serversAuditingSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersAuditingSetting_STATUSGenerator() gopter.Gen {
	if serversAuditingSetting_STATUSGenerator != nil {
		return serversAuditingSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAuditingSetting_STATUS(generators)
	serversAuditingSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersAuditingSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAuditingSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForServersAuditingSetting_STATUS(generators)
	serversAuditingSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersAuditingSetting_STATUS{}), generators)

	return serversAuditingSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersAuditingSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersAuditingSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersAuditingSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAuditingSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerBlobAuditingPolicyProperties_STATUSGenerator())
}
