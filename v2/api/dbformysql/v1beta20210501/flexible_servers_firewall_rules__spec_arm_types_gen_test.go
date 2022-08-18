// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_FlexibleServersFirewallRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRulesSpecARM, FlexibleServersFirewallRulesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRulesSpecARM runs a test to see if a specific instance of FlexibleServersFirewallRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRulesSpecARM(subject FlexibleServersFirewallRules_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRules_SpecARM
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

// Generator of FlexibleServersFirewallRules_SpecARM instances for property testing - lazily instantiated by
// FlexibleServersFirewallRulesSpecARMGenerator()
var flexibleServersFirewallRulesSpecARMGenerator gopter.Gen

// FlexibleServersFirewallRulesSpecARMGenerator returns a generator of FlexibleServersFirewallRules_SpecARM instances for property testing.
// We first initialize flexibleServersFirewallRulesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersFirewallRulesSpecARMGenerator() gopter.Gen {
	if flexibleServersFirewallRulesSpecARMGenerator != nil {
		return flexibleServersFirewallRulesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM(generators)
	flexibleServersFirewallRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM(generators)
	flexibleServersFirewallRulesSpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRules_SpecARM{}), generators)

	return flexibleServersFirewallRulesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRulesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRulePropertiesARMGenerator())
}

func Test_FirewallRulePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRulePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRulePropertiesARM, FirewallRulePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRulePropertiesARM runs a test to see if a specific instance of FirewallRulePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRulePropertiesARM(subject FirewallRulePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRulePropertiesARM
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

// Generator of FirewallRulePropertiesARM instances for property testing - lazily instantiated by
// FirewallRulePropertiesARMGenerator()
var firewallRulePropertiesARMGenerator gopter.Gen

// FirewallRulePropertiesARMGenerator returns a generator of FirewallRulePropertiesARM instances for property testing.
func FirewallRulePropertiesARMGenerator() gopter.Gen {
	if firewallRulePropertiesARMGenerator != nil {
		return firewallRulePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM(generators)
	firewallRulePropertiesARMGenerator = gen.Struct(reflect.TypeOf(FirewallRulePropertiesARM{}), generators)

	return firewallRulePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}
