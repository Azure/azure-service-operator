// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101preview

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

func Test_NamespacesQueue_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_Spec_ARM, NamespacesQueue_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_Spec_ARM runs a test to see if a specific instance of NamespacesQueue_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_Spec_ARM(subject NamespacesQueue_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue_Spec_ARM
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

// Generator of NamespacesQueue_Spec_ARM instances for property testing - lazily instantiated by
// NamespacesQueue_Spec_ARMGenerator()
var namespacesQueue_Spec_ARMGenerator gopter.Gen

// NamespacesQueue_Spec_ARMGenerator returns a generator of NamespacesQueue_Spec_ARM instances for property testing.
// We first initialize namespacesQueue_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesQueue_Spec_ARMGenerator() gopter.Gen {
	if namespacesQueue_Spec_ARMGenerator != nil {
		return namespacesQueue_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_Spec_ARM(generators)
	namespacesQueue_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesQueue_Spec_ARM(generators)
	namespacesQueue_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_Spec_ARM{}), generators)

	return namespacesQueue_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesQueue_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBQueueProperties_ARMGenerator())
}

func Test_SBQueueProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueueProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueProperties_ARM, SBQueueProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueProperties_ARM runs a test to see if a specific instance of SBQueueProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueProperties_ARM(subject SBQueueProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueueProperties_ARM
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

// Generator of SBQueueProperties_ARM instances for property testing - lazily instantiated by
// SBQueueProperties_ARMGenerator()
var sbQueueProperties_ARMGenerator gopter.Gen

// SBQueueProperties_ARMGenerator returns a generator of SBQueueProperties_ARM instances for property testing.
func SBQueueProperties_ARMGenerator() gopter.Gen {
	if sbQueueProperties_ARMGenerator != nil {
		return sbQueueProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueProperties_ARM(generators)
	sbQueueProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties_ARM{}), generators)

	return sbQueueProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueProperties_ARM(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}
