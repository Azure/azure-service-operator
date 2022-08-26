// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

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

<<<<<<<< HEAD:v2/api/servicebus/v1beta20210101preview/namespaces_queue__spec_arm_types_gen_test.go
func Test_NamespacesQueue_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_Namespaces_Queues_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/servicebus/v1beta20210101preview/namespaces_queues_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/servicebus/v1beta20210101preview/namespaces_queue__spec_arm_types_gen_test.go
		"Round trip of NamespacesQueue_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_SpecARM, NamespacesQueue_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_SpecARM runs a test to see if a specific instance of NamespacesQueue_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_SpecARM(subject NamespacesQueue_SpecARM) string {
========
		"Round trip of Namespaces_Queues_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Queues_SpecARM, Namespaces_Queues_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Queues_SpecARM runs a test to see if a specific instance of Namespaces_Queues_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Queues_SpecARM(subject Namespaces_Queues_SpecARM) string {
>>>>>>>> main:v2/api/servicebus/v1beta20210101preview/namespaces_queues_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/servicebus/v1beta20210101preview/namespaces_queue__spec_arm_types_gen_test.go
	var actual NamespacesQueue_SpecARM
========
	var actual Namespaces_Queues_SpecARM
>>>>>>>> main:v2/api/servicebus/v1beta20210101preview/namespaces_queues_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/servicebus/v1beta20210101preview/namespaces_queue__spec_arm_types_gen_test.go
// Generator of NamespacesQueue_SpecARM instances for property testing - lazily instantiated by
// NamespacesQueue_SpecARMGenerator()
var namespacesQueue_SpecARMGenerator gopter.Gen

// NamespacesQueue_SpecARMGenerator returns a generator of NamespacesQueue_SpecARM instances for property testing.
// We first initialize namespacesQueue_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesQueue_SpecARMGenerator() gopter.Gen {
	if namespacesQueue_SpecARMGenerator != nil {
		return namespacesQueue_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	namespacesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM(generators)
	namespacesQueue_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_SpecARM{}), generators)

	return namespacesQueue_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of Namespaces_Queues_SpecARM instances for property testing - lazily instantiated by
// Namespaces_Queues_SpecARMGenerator()
var namespaces_Queues_SpecARMGenerator gopter.Gen

// Namespaces_Queues_SpecARMGenerator returns a generator of Namespaces_Queues_SpecARM instances for property testing.
// We first initialize namespaces_Queues_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Queues_SpecARMGenerator() gopter.Gen {
	if namespaces_Queues_SpecARMGenerator != nil {
		return namespaces_Queues_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queues_SpecARM(generators)
	namespaces_Queues_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queues_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Queues_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Queues_SpecARM(generators)
	namespaces_Queues_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Queues_SpecARM{}), generators)

	return namespaces_Queues_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Queues_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Queues_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/servicebus/v1beta20210101preview/namespaces_queues_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
}

<<<<<<<< HEAD:v2/api/servicebus/v1beta20210101preview/namespaces_queue__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForNamespaces_Queues_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Queues_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/servicebus/v1beta20210101preview/namespaces_queues_spec_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(SBQueuePropertiesARMGenerator())
}

func Test_SBQueuePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueuePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueuePropertiesARM, SBQueuePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueuePropertiesARM runs a test to see if a specific instance of SBQueuePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueuePropertiesARM(subject SBQueuePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueuePropertiesARM
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

// Generator of SBQueuePropertiesARM instances for property testing - lazily instantiated by
// SBQueuePropertiesARMGenerator()
var sbQueuePropertiesARMGenerator gopter.Gen

// SBQueuePropertiesARMGenerator returns a generator of SBQueuePropertiesARM instances for property testing.
func SBQueuePropertiesARMGenerator() gopter.Gen {
	if sbQueuePropertiesARMGenerator != nil {
		return sbQueuePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueuePropertiesARM(generators)
	sbQueuePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SBQueuePropertiesARM{}), generators)

	return sbQueuePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForSBQueuePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueuePropertiesARM(gens map[string]gopter.Gen) {
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
