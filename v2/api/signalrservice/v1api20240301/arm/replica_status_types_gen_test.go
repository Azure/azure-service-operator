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

func Test_ReplicaProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReplicaProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplicaProperties_STATUS, ReplicaProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplicaProperties_STATUS runs a test to see if a specific instance of ReplicaProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReplicaProperties_STATUS(subject ReplicaProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReplicaProperties_STATUS
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

// Generator of ReplicaProperties_STATUS instances for property testing - lazily instantiated by
// ReplicaProperties_STATUSGenerator()
var replicaProperties_STATUSGenerator gopter.Gen

// ReplicaProperties_STATUSGenerator returns a generator of ReplicaProperties_STATUS instances for property testing.
func ReplicaProperties_STATUSGenerator() gopter.Gen {
	if replicaProperties_STATUSGenerator != nil {
		return replicaProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplicaProperties_STATUS(generators)
	replicaProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReplicaProperties_STATUS{}), generators)

	return replicaProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReplicaProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReplicaProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Moving,
		ProvisioningState_STATUS_Running,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Unknown,
		ProvisioningState_STATUS_Updating))
	gens["RegionEndpointEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceStopped"] = gen.PtrOf(gen.AlphaString())
}

func Test_Replica_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Replica_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplica_STATUS, Replica_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplica_STATUS runs a test to see if a specific instance of Replica_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReplica_STATUS(subject Replica_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Replica_STATUS
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

// Generator of Replica_STATUS instances for property testing - lazily instantiated by Replica_STATUSGenerator()
var replica_STATUSGenerator gopter.Gen

// Replica_STATUSGenerator returns a generator of Replica_STATUS instances for property testing.
// We first initialize replica_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Replica_STATUSGenerator() gopter.Gen {
	if replica_STATUSGenerator != nil {
		return replica_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_STATUS(generators)
	replica_STATUSGenerator = gen.Struct(reflect.TypeOf(Replica_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_STATUS(generators)
	AddRelatedPropertyGeneratorsForReplica_STATUS(generators)
	replica_STATUSGenerator = gen.Struct(reflect.TypeOf(Replica_STATUS{}), generators)

	return replica_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReplica_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReplica_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReplica_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReplica_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReplicaProperties_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(ResourceSku_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_ResourceSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceSku_STATUS, ResourceSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceSku_STATUS runs a test to see if a specific instance of ResourceSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceSku_STATUS(subject ResourceSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceSku_STATUS
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

// Generator of ResourceSku_STATUS instances for property testing - lazily instantiated by ResourceSku_STATUSGenerator()
var resourceSku_STATUSGenerator gopter.Gen

// ResourceSku_STATUSGenerator returns a generator of ResourceSku_STATUS instances for property testing.
func ResourceSku_STATUSGenerator() gopter.Gen {
	if resourceSku_STATUSGenerator != nil {
		return resourceSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceSku_STATUS(generators)
	resourceSku_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceSku_STATUS{}), generators)

	return resourceSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(
		SignalRSkuTier_STATUS_Basic,
		SignalRSkuTier_STATUS_Free,
		SignalRSkuTier_STATUS_Premium,
		SignalRSkuTier_STATUS_Standard))
}
