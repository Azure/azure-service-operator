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

func Test_ServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersFirewallRule, ServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersFirewallRule runs a test to see if a specific instance of ServersFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersFirewallRule(subject ServersFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersFirewallRule
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

// Generator of ServersFirewallRule instances for property testing - lazily instantiated by
// ServersFirewallRuleGenerator()
var serversFirewallRuleGenerator gopter.Gen

// ServersFirewallRuleGenerator returns a generator of ServersFirewallRule instances for property testing.
func ServersFirewallRuleGenerator() gopter.Gen {
	if serversFirewallRuleGenerator != nil {
		return serversFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersFirewallRule(generators)
	serversFirewallRuleGenerator = gen.Struct(reflect.TypeOf(ServersFirewallRule{}), generators)

	return serversFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersFirewallRule_SpecGenerator()
	gens["Status"] = ServersFirewallRule_STATUSGenerator()
}

func Test_ServersFirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersFirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersFirewallRule_STATUS, ServersFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersFirewallRule_STATUS runs a test to see if a specific instance of ServersFirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersFirewallRule_STATUS(subject ServersFirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersFirewallRule_STATUS
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

// Generator of ServersFirewallRule_STATUS instances for property testing - lazily instantiated by
// ServersFirewallRule_STATUSGenerator()
var serversFirewallRule_STATUSGenerator gopter.Gen

// ServersFirewallRule_STATUSGenerator returns a generator of ServersFirewallRule_STATUS instances for property testing.
func ServersFirewallRule_STATUSGenerator() gopter.Gen {
	if serversFirewallRule_STATUSGenerator != nil {
		return serversFirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersFirewallRule_STATUS(generators)
	serversFirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersFirewallRule_STATUS{}), generators)

	return serversFirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersFirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersFirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersFirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersFirewallRule_Spec, ServersFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersFirewallRule_Spec runs a test to see if a specific instance of ServersFirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersFirewallRule_Spec(subject ServersFirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersFirewallRule_Spec
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

// Generator of ServersFirewallRule_Spec instances for property testing - lazily instantiated by
// ServersFirewallRule_SpecGenerator()
var serversFirewallRule_SpecGenerator gopter.Gen

// ServersFirewallRule_SpecGenerator returns a generator of ServersFirewallRule_Spec instances for property testing.
func ServersFirewallRule_SpecGenerator() gopter.Gen {
	if serversFirewallRule_SpecGenerator != nil {
		return serversFirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersFirewallRule_Spec(generators)
	serversFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersFirewallRule_Spec{}), generators)

	return serversFirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersFirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}
