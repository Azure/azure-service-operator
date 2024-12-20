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

func Test_ServersVirtualNetworkRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule, ServersVirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule runs a test to see if a specific instance of ServersVirtualNetworkRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule(subject ServersVirtualNetworkRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule
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

// Generator of ServersVirtualNetworkRule instances for property testing - lazily instantiated by
// ServersVirtualNetworkRuleGenerator()
var serversVirtualNetworkRuleGenerator gopter.Gen

// ServersVirtualNetworkRuleGenerator returns a generator of ServersVirtualNetworkRule instances for property testing.
func ServersVirtualNetworkRuleGenerator() gopter.Gen {
	if serversVirtualNetworkRuleGenerator != nil {
		return serversVirtualNetworkRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersVirtualNetworkRule(generators)
	serversVirtualNetworkRuleGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule{}), generators)

	return serversVirtualNetworkRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersVirtualNetworkRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersVirtualNetworkRule(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersVirtualNetworkRule_SpecGenerator()
	gens["Status"] = ServersVirtualNetworkRule_STATUSGenerator()
}

func Test_ServersVirtualNetworkRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRuleOperatorSpec, ServersVirtualNetworkRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRuleOperatorSpec runs a test to see if a specific instance of ServersVirtualNetworkRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRuleOperatorSpec(subject ServersVirtualNetworkRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRuleOperatorSpec
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

// Generator of ServersVirtualNetworkRuleOperatorSpec instances for property testing - lazily instantiated by
// ServersVirtualNetworkRuleOperatorSpecGenerator()
var serversVirtualNetworkRuleOperatorSpecGenerator gopter.Gen

// ServersVirtualNetworkRuleOperatorSpecGenerator returns a generator of ServersVirtualNetworkRuleOperatorSpec instances for property testing.
func ServersVirtualNetworkRuleOperatorSpecGenerator() gopter.Gen {
	if serversVirtualNetworkRuleOperatorSpecGenerator != nil {
		return serversVirtualNetworkRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	serversVirtualNetworkRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRuleOperatorSpec{}), generators)

	return serversVirtualNetworkRuleOperatorSpecGenerator
}

func Test_ServersVirtualNetworkRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule_STATUS, ServersVirtualNetworkRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule_STATUS runs a test to see if a specific instance of ServersVirtualNetworkRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule_STATUS(subject ServersVirtualNetworkRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule_STATUS
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

// Generator of ServersVirtualNetworkRule_STATUS instances for property testing - lazily instantiated by
// ServersVirtualNetworkRule_STATUSGenerator()
var serversVirtualNetworkRule_STATUSGenerator gopter.Gen

// ServersVirtualNetworkRule_STATUSGenerator returns a generator of ServersVirtualNetworkRule_STATUS instances for property testing.
func ServersVirtualNetworkRule_STATUSGenerator() gopter.Gen {
	if serversVirtualNetworkRule_STATUSGenerator != nil {
		return serversVirtualNetworkRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS(generators)
	serversVirtualNetworkRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule_STATUS{}), generators)

	return serversVirtualNetworkRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkSubnetId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersVirtualNetworkRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule_Spec, ServersVirtualNetworkRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule_Spec runs a test to see if a specific instance of ServersVirtualNetworkRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule_Spec(subject ServersVirtualNetworkRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule_Spec
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

// Generator of ServersVirtualNetworkRule_Spec instances for property testing - lazily instantiated by
// ServersVirtualNetworkRule_SpecGenerator()
var serversVirtualNetworkRule_SpecGenerator gopter.Gen

// ServersVirtualNetworkRule_SpecGenerator returns a generator of ServersVirtualNetworkRule_Spec instances for property testing.
// We first initialize serversVirtualNetworkRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersVirtualNetworkRule_SpecGenerator() gopter.Gen {
	if serversVirtualNetworkRule_SpecGenerator != nil {
		return serversVirtualNetworkRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec(generators)
	serversVirtualNetworkRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec(generators)
	AddRelatedPropertyGeneratorsForServersVirtualNetworkRule_Spec(generators)
	serversVirtualNetworkRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule_Spec{}), generators)

	return serversVirtualNetworkRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersVirtualNetworkRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersVirtualNetworkRule_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServersVirtualNetworkRuleOperatorSpecGenerator())
}
