// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

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

func Test_FirewallRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM, FirewallRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM runs a test to see if a specific instance of FirewallRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM(subject FirewallRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRuleProperties_STATUS_ARM
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

// Generator of FirewallRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// FirewallRuleProperties_STATUS_ARMGenerator()
var firewallRuleProperties_STATUS_ARMGenerator gopter.Gen

// FirewallRuleProperties_STATUS_ARMGenerator returns a generator of FirewallRuleProperties_STATUS_ARM instances for property testing.
func FirewallRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if firewallRuleProperties_STATUS_ARMGenerator != nil {
		return firewallRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM(generators)
	firewallRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FirewallRuleProperties_STATUS_ARM{}), generators)

	return firewallRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServers_FirewallRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_FirewallRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS_ARM, FlexibleServers_FirewallRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS_ARM runs a test to see if a specific instance of FlexibleServers_FirewallRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_FirewallRule_STATUS_ARM(subject FlexibleServers_FirewallRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_FirewallRule_STATUS_ARM
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

// Generator of FlexibleServers_FirewallRule_STATUS_ARM instances for property testing - lazily instantiated by
// FlexibleServers_FirewallRule_STATUS_ARMGenerator()
var flexibleServers_FirewallRule_STATUS_ARMGenerator gopter.Gen

// FlexibleServers_FirewallRule_STATUS_ARMGenerator returns a generator of FlexibleServers_FirewallRule_STATUS_ARM instances for property testing.
// We first initialize flexibleServers_FirewallRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_FirewallRule_STATUS_ARMGenerator() gopter.Gen {
	if flexibleServers_FirewallRule_STATUS_ARMGenerator != nil {
		return flexibleServers_FirewallRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM(generators)
	flexibleServers_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM(generators)
	flexibleServers_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_STATUS_ARM{}), generators)

	return flexibleServers_FirewallRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRuleProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}
