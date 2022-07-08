// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

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

func Test_FirewallRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleStatusARM, FirewallRuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleStatusARM runs a test to see if a specific instance of FirewallRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleStatusARM(subject FirewallRule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_StatusARM
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

// Generator of FirewallRule_StatusARM instances for property testing - lazily instantiated by
// FirewallRuleStatusARMGenerator()
var firewallRuleStatusARMGenerator gopter.Gen

// FirewallRuleStatusARMGenerator returns a generator of FirewallRule_StatusARM instances for property testing.
// We first initialize firewallRuleStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRuleStatusARMGenerator() gopter.Gen {
	if firewallRuleStatusARMGenerator != nil {
		return firewallRuleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleStatusARM(generators)
	firewallRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleStatusARM(generators)
	AddRelatedPropertyGeneratorsForFirewallRuleStatusARM(generators)
	firewallRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_StatusARM{}), generators)

	return firewallRuleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRuleStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRulePropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_FirewallRuleProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRuleProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRulePropertiesStatusARM, FirewallRulePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRulePropertiesStatusARM runs a test to see if a specific instance of FirewallRuleProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRulePropertiesStatusARM(subject FirewallRuleProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRuleProperties_StatusARM
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

// Generator of FirewallRuleProperties_StatusARM instances for property testing - lazily instantiated by
// FirewallRulePropertiesStatusARMGenerator()
var firewallRulePropertiesStatusARMGenerator gopter.Gen

// FirewallRulePropertiesStatusARMGenerator returns a generator of FirewallRuleProperties_StatusARM instances for property testing.
func FirewallRulePropertiesStatusARMGenerator() gopter.Gen {
	if firewallRulePropertiesStatusARMGenerator != nil {
		return firewallRulePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRulePropertiesStatusARM(generators)
	firewallRulePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(FirewallRuleProperties_StatusARM{}), generators)

	return firewallRulePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRulePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRulePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}
