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

func Test_NamespacesTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic_STATUS, NamespacesTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic_STATUS runs a test to see if a specific instance of NamespacesTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic_STATUS(subject NamespacesTopic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic_STATUS
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

// Generator of NamespacesTopic_STATUS instances for property testing - lazily instantiated by
// NamespacesTopic_STATUSGenerator()
var namespacesTopic_STATUSGenerator gopter.Gen

// NamespacesTopic_STATUSGenerator returns a generator of NamespacesTopic_STATUS instances for property testing.
// We first initialize namespacesTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopic_STATUSGenerator() gopter.Gen {
	if namespacesTopic_STATUSGenerator != nil {
		return namespacesTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	return namespacesTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBTopicProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_SBTopicProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopicProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicProperties_STATUS, SBTopicProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicProperties_STATUS runs a test to see if a specific instance of SBTopicProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicProperties_STATUS(subject SBTopicProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopicProperties_STATUS
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

// Generator of SBTopicProperties_STATUS instances for property testing - lazily instantiated by
// SBTopicProperties_STATUSGenerator()
var sbTopicProperties_STATUSGenerator gopter.Gen

// SBTopicProperties_STATUSGenerator returns a generator of SBTopicProperties_STATUS instances for property testing.
// We first initialize sbTopicProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopicProperties_STATUSGenerator() gopter.Gen {
	if sbTopicProperties_STATUSGenerator != nil {
		return sbTopicProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS(generators)
	sbTopicProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS(generators)
	sbTopicProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_STATUS{}), generators)

	return sbTopicProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
}
