// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_NamespacesEventhubsConsumergroup_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumergroup_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumergroup_SpecARM, NamespacesEventhubsConsumergroup_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumergroup_SpecARM runs a test to see if a specific instance of NamespacesEventhubsConsumergroup_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumergroup_SpecARM(subject NamespacesEventhubsConsumergroup_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumergroup_SpecARM
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

// Generator of NamespacesEventhubsConsumergroup_SpecARM instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumergroup_SpecARMGenerator()
var namespacesEventhubsConsumergroup_SpecARMGenerator gopter.Gen

// NamespacesEventhubsConsumergroup_SpecARMGenerator returns a generator of NamespacesEventhubsConsumergroup_SpecARM instances for property testing.
// We first initialize namespacesEventhubsConsumergroup_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsConsumergroup_SpecARMGenerator() gopter.Gen {
	if namespacesEventhubsConsumergroup_SpecARMGenerator != nil {
		return namespacesEventhubsConsumergroup_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM(generators)
	namespacesEventhubsConsumergroup_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumergroup_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM(generators)
	namespacesEventhubsConsumergroup_SpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumergroup_SpecARM{}), generators)

	return namespacesEventhubsConsumergroup_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumergroup_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator())
}

func Test_NamespacesEventhubsConsumergroup_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumergroup_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumergroup_Spec_PropertiesARM, NamespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumergroup_Spec_PropertiesARM runs a test to see if a specific instance of NamespacesEventhubsConsumergroup_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumergroup_Spec_PropertiesARM(subject NamespacesEventhubsConsumergroup_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumergroup_Spec_PropertiesARM
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

// Generator of NamespacesEventhubsConsumergroup_Spec_PropertiesARM instances for property testing - lazily instantiated
// by NamespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator()
var namespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator gopter.Gen

// NamespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator returns a generator of NamespacesEventhubsConsumergroup_Spec_PropertiesARM instances for property testing.
func NamespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator() gopter.Gen {
	if namespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator != nil {
		return namespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_Spec_PropertiesARM(generators)
	namespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumergroup_Spec_PropertiesARM{}), generators)

	return namespacesEventhubsConsumergroup_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroup_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}
