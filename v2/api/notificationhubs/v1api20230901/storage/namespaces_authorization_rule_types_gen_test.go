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

func Test_NamespacesAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule runs a test to see if a specific instance of NamespacesAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule
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

// Generator of NamespacesAuthorizationRule instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleGenerator()
var namespacesAuthorizationRuleGenerator gopter.Gen

// NamespacesAuthorizationRuleGenerator returns a generator of NamespacesAuthorizationRule instances for property testing.
func NamespacesAuthorizationRuleGenerator() gopter.Gen {
	if namespacesAuthorizationRuleGenerator != nil {
		return namespacesAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(generators)
	namespacesAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule{}), generators)

	return namespacesAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesAuthorizationRule_SpecGenerator()
	gens["Status"] = NamespacesAuthorizationRule_STATUSGenerator()
}

func Test_NamespacesAuthorizationRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec, NamespacesAuthorizationRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec runs a test to see if a specific instance of NamespacesAuthorizationRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec(subject NamespacesAuthorizationRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRuleOperatorSpec
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

// Generator of NamespacesAuthorizationRuleOperatorSpec instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleOperatorSpecGenerator()
var namespacesAuthorizationRuleOperatorSpecGenerator gopter.Gen

// NamespacesAuthorizationRuleOperatorSpecGenerator returns a generator of NamespacesAuthorizationRuleOperatorSpec instances for property testing.
func NamespacesAuthorizationRuleOperatorSpecGenerator() gopter.Gen {
	if namespacesAuthorizationRuleOperatorSpecGenerator != nil {
		return namespacesAuthorizationRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	namespacesAuthorizationRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRuleOperatorSpec{}), generators)

	return namespacesAuthorizationRuleOperatorSpecGenerator
}

func Test_NamespacesAuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS, NamespacesAuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS runs a test to see if a specific instance of NamespacesAuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS(subject NamespacesAuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_STATUS
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

// Generator of NamespacesAuthorizationRule_STATUS instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_STATUSGenerator()
var namespacesAuthorizationRule_STATUSGenerator gopter.Gen

// NamespacesAuthorizationRule_STATUSGenerator returns a generator of NamespacesAuthorizationRule_STATUS instances for property testing.
// We first initialize namespacesAuthorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_STATUSGenerator() gopter.Gen {
	if namespacesAuthorizationRule_STATUSGenerator != nil {
		return namespacesAuthorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	namespacesAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	namespacesAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUS{}), generators)

	return namespacesAuthorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedAccessAuthorizationRuleProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesAuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_Spec, NamespacesAuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_Spec runs a test to see if a specific instance of NamespacesAuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_Spec(subject NamespacesAuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_Spec
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

// Generator of NamespacesAuthorizationRule_Spec instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_SpecGenerator()
var namespacesAuthorizationRule_SpecGenerator gopter.Gen

// NamespacesAuthorizationRule_SpecGenerator returns a generator of NamespacesAuthorizationRule_Spec instances for property testing.
// We first initialize namespacesAuthorizationRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_SpecGenerator() gopter.Gen {
	if namespacesAuthorizationRule_SpecGenerator != nil {
		return namespacesAuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	namespacesAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	namespacesAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec{}), generators)

	return namespacesAuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(NamespacesAuthorizationRuleOperatorSpecGenerator())
	gens["Properties"] = gen.PtrOf(SharedAccessAuthorizationRulePropertiesGenerator())
}

func Test_SharedAccessAuthorizationRuleProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedAccessAuthorizationRuleProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties, SharedAccessAuthorizationRulePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties runs a test to see if a specific instance of SharedAccessAuthorizationRuleProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties(subject SharedAccessAuthorizationRuleProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedAccessAuthorizationRuleProperties
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

// Generator of SharedAccessAuthorizationRuleProperties instances for property testing - lazily instantiated by
// SharedAccessAuthorizationRulePropertiesGenerator()
var sharedAccessAuthorizationRulePropertiesGenerator gopter.Gen

// SharedAccessAuthorizationRulePropertiesGenerator returns a generator of SharedAccessAuthorizationRuleProperties instances for property testing.
func SharedAccessAuthorizationRulePropertiesGenerator() gopter.Gen {
	if sharedAccessAuthorizationRulePropertiesGenerator != nil {
		return sharedAccessAuthorizationRulePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties(generators)
	sharedAccessAuthorizationRulePropertiesGenerator = gen.Struct(reflect.TypeOf(SharedAccessAuthorizationRuleProperties{}), generators)

	return sharedAccessAuthorizationRulePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
}

func Test_SharedAccessAuthorizationRuleProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedAccessAuthorizationRuleProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties_STATUS, SharedAccessAuthorizationRuleProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties_STATUS runs a test to see if a specific instance of SharedAccessAuthorizationRuleProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedAccessAuthorizationRuleProperties_STATUS(subject SharedAccessAuthorizationRuleProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedAccessAuthorizationRuleProperties_STATUS
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

// Generator of SharedAccessAuthorizationRuleProperties_STATUS instances for property testing - lazily instantiated by
// SharedAccessAuthorizationRuleProperties_STATUSGenerator()
var sharedAccessAuthorizationRuleProperties_STATUSGenerator gopter.Gen

// SharedAccessAuthorizationRuleProperties_STATUSGenerator returns a generator of SharedAccessAuthorizationRuleProperties_STATUS instances for property testing.
func SharedAccessAuthorizationRuleProperties_STATUSGenerator() gopter.Gen {
	if sharedAccessAuthorizationRuleProperties_STATUSGenerator != nil {
		return sharedAccessAuthorizationRuleProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties_STATUS(generators)
	sharedAccessAuthorizationRuleProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SharedAccessAuthorizationRuleProperties_STATUS{}), generators)

	return sharedAccessAuthorizationRuleProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedAccessAuthorizationRuleProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ClaimType"] = gen.PtrOf(gen.AlphaString())
	gens["ClaimValue"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["ModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["Revision"] = gen.PtrOf(gen.Int())
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
}
