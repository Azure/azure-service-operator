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

func Test_Namespaces_Topics_AuthorizationRule_Properties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topics_AuthorizationRule_Properties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topics_AuthorizationRule_Properties_STATUS, Namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topics_AuthorizationRule_Properties_STATUS runs a test to see if a specific instance of Namespaces_Topics_AuthorizationRule_Properties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topics_AuthorizationRule_Properties_STATUS(subject Namespaces_Topics_AuthorizationRule_Properties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topics_AuthorizationRule_Properties_STATUS
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

// Generator of Namespaces_Topics_AuthorizationRule_Properties_STATUS instances for property testing - lazily
// instantiated by Namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator()
var namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator gopter.Gen

// Namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator returns a generator of Namespaces_Topics_AuthorizationRule_Properties_STATUS instances for property testing.
func Namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator() gopter.Gen {
	if namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator != nil {
		return namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_AuthorizationRule_Properties_STATUS(generators)
	namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_AuthorizationRule_Properties_STATUS{}), generators)

	return namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topics_AuthorizationRule_Properties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topics_AuthorizationRule_Properties_STATUS(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(TopicAuthorizationRuleRights_STATUS_Listen, TopicAuthorizationRuleRights_STATUS_Manage, TopicAuthorizationRuleRights_STATUS_Send))
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
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopicAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_Topics_AuthorizationRule_Properties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
