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

func Test_IPv6ServerFirewallRuleProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPv6ServerFirewallRuleProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPv6ServerFirewallRuleProperties, IPv6ServerFirewallRulePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPv6ServerFirewallRuleProperties runs a test to see if a specific instance of IPv6ServerFirewallRuleProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForIPv6ServerFirewallRuleProperties(subject IPv6ServerFirewallRuleProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPv6ServerFirewallRuleProperties
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

// Generator of IPv6ServerFirewallRuleProperties instances for property testing - lazily instantiated by
// IPv6ServerFirewallRulePropertiesGenerator()
var iPv6ServerFirewallRulePropertiesGenerator gopter.Gen

// IPv6ServerFirewallRulePropertiesGenerator returns a generator of IPv6ServerFirewallRuleProperties instances for property testing.
func IPv6ServerFirewallRulePropertiesGenerator() gopter.Gen {
	if iPv6ServerFirewallRulePropertiesGenerator != nil {
		return iPv6ServerFirewallRulePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPv6ServerFirewallRuleProperties(generators)
	iPv6ServerFirewallRulePropertiesGenerator = gen.Struct(reflect.TypeOf(IPv6ServerFirewallRuleProperties{}), generators)

	return iPv6ServerFirewallRulePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForIPv6ServerFirewallRuleProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPv6ServerFirewallRuleProperties(gens map[string]gopter.Gen) {
	gens["EndIPv6Address"] = gen.PtrOf(gen.AlphaString())
	gens["StartIPv6Address"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersIPV6FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersIPV6FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersIPV6FirewallRule_Spec, ServersIPV6FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersIPV6FirewallRule_Spec runs a test to see if a specific instance of ServersIPV6FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersIPV6FirewallRule_Spec(subject ServersIPV6FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersIPV6FirewallRule_Spec
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

// Generator of ServersIPV6FirewallRule_Spec instances for property testing - lazily instantiated by
// ServersIPV6FirewallRule_SpecGenerator()
var serversIPV6FirewallRule_SpecGenerator gopter.Gen

// ServersIPV6FirewallRule_SpecGenerator returns a generator of ServersIPV6FirewallRule_Spec instances for property testing.
// We first initialize serversIPV6FirewallRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersIPV6FirewallRule_SpecGenerator() gopter.Gen {
	if serversIPV6FirewallRule_SpecGenerator != nil {
		return serversIPV6FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	serversIPV6FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec(generators)
	serversIPV6FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersIPV6FirewallRule_Spec{}), generators)

	return serversIPV6FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersIPV6FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersIPV6FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(IPv6ServerFirewallRulePropertiesGenerator())
}