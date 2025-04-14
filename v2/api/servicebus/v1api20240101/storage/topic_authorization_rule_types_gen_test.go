// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_TopicAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicAuthorizationRule, TopicAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicAuthorizationRule runs a test to see if a specific instance of TopicAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicAuthorizationRule(subject TopicAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicAuthorizationRule
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

// Generator of TopicAuthorizationRule instances for property testing - lazily instantiated by
// TopicAuthorizationRuleGenerator()
var topicAuthorizationRuleGenerator gopter.Gen

// TopicAuthorizationRuleGenerator returns a generator of TopicAuthorizationRule instances for property testing.
func TopicAuthorizationRuleGenerator() gopter.Gen {
	if topicAuthorizationRuleGenerator != nil {
		return topicAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTopicAuthorizationRule(generators)
	topicAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRule{}), generators)

	return topicAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForTopicAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = TopicAuthorizationRule_SpecGenerator()
	gens["Status"] = TopicAuthorizationRule_STATUSGenerator()
}

func Test_TopicAuthorizationRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicAuthorizationRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicAuthorizationRuleOperatorSpec, TopicAuthorizationRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicAuthorizationRuleOperatorSpec runs a test to see if a specific instance of TopicAuthorizationRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicAuthorizationRuleOperatorSpec(subject TopicAuthorizationRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicAuthorizationRuleOperatorSpec
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

// Generator of TopicAuthorizationRuleOperatorSpec instances for property testing - lazily instantiated by
// TopicAuthorizationRuleOperatorSpecGenerator()
var topicAuthorizationRuleOperatorSpecGenerator gopter.Gen

// TopicAuthorizationRuleOperatorSpecGenerator returns a generator of TopicAuthorizationRuleOperatorSpec instances for property testing.
func TopicAuthorizationRuleOperatorSpecGenerator() gopter.Gen {
	if topicAuthorizationRuleOperatorSpecGenerator != nil {
		return topicAuthorizationRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	topicAuthorizationRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRuleOperatorSpec{}), generators)

	return topicAuthorizationRuleOperatorSpecGenerator
}

func Test_TopicAuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicAuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicAuthorizationRule_STATUS, TopicAuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicAuthorizationRule_STATUS runs a test to see if a specific instance of TopicAuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicAuthorizationRule_STATUS(subject TopicAuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicAuthorizationRule_STATUS
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

// Generator of TopicAuthorizationRule_STATUS instances for property testing - lazily instantiated by
// TopicAuthorizationRule_STATUSGenerator()
var topicAuthorizationRule_STATUSGenerator gopter.Gen

// TopicAuthorizationRule_STATUSGenerator returns a generator of TopicAuthorizationRule_STATUS instances for property testing.
// We first initialize topicAuthorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TopicAuthorizationRule_STATUSGenerator() gopter.Gen {
	if topicAuthorizationRule_STATUSGenerator != nil {
		return topicAuthorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicAuthorizationRule_STATUS(generators)
	topicAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicAuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForTopicAuthorizationRule_STATUS(generators)
	topicAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRule_STATUS{}), generators)

	return topicAuthorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTopicAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopicAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_TopicAuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicAuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicAuthorizationRule_Spec, TopicAuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicAuthorizationRule_Spec runs a test to see if a specific instance of TopicAuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicAuthorizationRule_Spec(subject TopicAuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicAuthorizationRule_Spec
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

// Generator of TopicAuthorizationRule_Spec instances for property testing - lazily instantiated by
// TopicAuthorizationRule_SpecGenerator()
var topicAuthorizationRule_SpecGenerator gopter.Gen

// TopicAuthorizationRule_SpecGenerator returns a generator of TopicAuthorizationRule_Spec instances for property testing.
// We first initialize topicAuthorizationRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TopicAuthorizationRule_SpecGenerator() gopter.Gen {
	if topicAuthorizationRule_SpecGenerator != nil {
		return topicAuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicAuthorizationRule_Spec(generators)
	topicAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicAuthorizationRule_Spec(generators)
	AddRelatedPropertyGeneratorsForTopicAuthorizationRule_Spec(generators)
	topicAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(TopicAuthorizationRule_Spec{}), generators)

	return topicAuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTopicAuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopicAuthorizationRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(TopicAuthorizationRuleOperatorSpecGenerator())
}
