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

func Test_FlexibleServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule runs a test to see if a specific instance of FlexibleServersFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule
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

// Generator of FlexibleServersFirewallRule instances for property testing - lazily instantiated by
// FlexibleServersFirewallRuleGenerator()
var flexibleServersFirewallRuleGenerator gopter.Gen

// FlexibleServersFirewallRuleGenerator returns a generator of FlexibleServersFirewallRule instances for property testing.
func FlexibleServersFirewallRuleGenerator() gopter.Gen {
	if flexibleServersFirewallRuleGenerator != nil {
		return flexibleServersFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(generators)
	flexibleServersFirewallRuleGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule{}), generators)

	return flexibleServersFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersFirewallRule_SpecGenerator()
	gens["Status"] = FlexibleServersFirewallRule_STATUSGenerator()
}

func Test_FlexibleServersFirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule_STATUS, FlexibleServersFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule_STATUS runs a test to see if a specific instance of FlexibleServersFirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule_STATUS(subject FlexibleServersFirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule_STATUS
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

// Generator of FlexibleServersFirewallRule_STATUS instances for property testing - lazily instantiated by
// FlexibleServersFirewallRule_STATUSGenerator()
var flexibleServersFirewallRule_STATUSGenerator gopter.Gen

// FlexibleServersFirewallRule_STATUSGenerator returns a generator of FlexibleServersFirewallRule_STATUS instances for property testing.
// We first initialize flexibleServersFirewallRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersFirewallRule_STATUSGenerator() gopter.Gen {
	if flexibleServersFirewallRule_STATUSGenerator != nil {
		return flexibleServersFirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUS(generators)
	flexibleServersFirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUS(generators)
	flexibleServersFirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_STATUS{}), generators)

	return flexibleServersFirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServersFirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule_Spec, FlexibleServersFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule_Spec runs a test to see if a specific instance of FlexibleServersFirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule_Spec(subject FlexibleServersFirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule_Spec
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

// Generator of FlexibleServersFirewallRule_Spec instances for property testing - lazily instantiated by
// FlexibleServersFirewallRule_SpecGenerator()
var flexibleServersFirewallRule_SpecGenerator gopter.Gen

// FlexibleServersFirewallRule_SpecGenerator returns a generator of FlexibleServersFirewallRule_Spec instances for property testing.
func FlexibleServersFirewallRule_SpecGenerator() gopter.Gen {
	if flexibleServersFirewallRule_SpecGenerator != nil {
		return flexibleServersFirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_Spec(generators)
	flexibleServersFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_Spec{}), generators)

	return flexibleServersFirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}
