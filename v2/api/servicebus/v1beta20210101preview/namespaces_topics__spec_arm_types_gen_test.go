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

func Test_NamespacesTopics_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopics_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSpecARM, NamespacesTopicsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSpecARM runs a test to see if a specific instance of NamespacesTopics_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSpecARM(subject NamespacesTopics_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopics_SpecARM
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

// Generator of NamespacesTopics_SpecARM instances for property testing - lazily instantiated by
// NamespacesTopicsSpecARMGenerator()
var namespacesTopicsSpecARMGenerator gopter.Gen

// NamespacesTopicsSpecARMGenerator returns a generator of NamespacesTopics_SpecARM instances for property testing.
// We first initialize namespacesTopicsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopicsSpecARMGenerator() gopter.Gen {
	if namespacesTopicsSpecARMGenerator != nil {
		return namespacesTopicsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSpecARM(generators)
	namespacesTopicsSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesTopics_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopicsSpecARM(generators)
	namespacesTopicsSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesTopics_SpecARM{}), generators)

	return namespacesTopicsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopicsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopicsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesTopicsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopicsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBTopicPropertiesARMGenerator())
}

func Test_SBTopicPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopicPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicPropertiesARM, SBTopicPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicPropertiesARM runs a test to see if a specific instance of SBTopicPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicPropertiesARM(subject SBTopicPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopicPropertiesARM
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

// Generator of SBTopicPropertiesARM instances for property testing - lazily instantiated by
// SBTopicPropertiesARMGenerator()
var sbTopicPropertiesARMGenerator gopter.Gen

// SBTopicPropertiesARMGenerator returns a generator of SBTopicPropertiesARM instances for property testing.
func SBTopicPropertiesARMGenerator() gopter.Gen {
	if sbTopicPropertiesARMGenerator != nil {
		return sbTopicPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicPropertiesARM(generators)
	sbTopicPropertiesARMGenerator = gen.Struct(reflect.TypeOf(SBTopicPropertiesARM{}), generators)

	return sbTopicPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicPropertiesARM(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
}
