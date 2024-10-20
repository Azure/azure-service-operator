// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_OutboundFirewallRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OutboundFirewallRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOutboundFirewallRuleProperties_STATUS_ARM, OutboundFirewallRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOutboundFirewallRuleProperties_STATUS_ARM runs a test to see if a specific instance of OutboundFirewallRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForOutboundFirewallRuleProperties_STATUS_ARM(subject OutboundFirewallRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OutboundFirewallRuleProperties_STATUS_ARM
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

// Generator of OutboundFirewallRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// OutboundFirewallRuleProperties_STATUS_ARMGenerator()
var outboundFirewallRuleProperties_STATUS_ARMGenerator gopter.Gen

// OutboundFirewallRuleProperties_STATUS_ARMGenerator returns a generator of OutboundFirewallRuleProperties_STATUS_ARM instances for property testing.
func OutboundFirewallRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if outboundFirewallRuleProperties_STATUS_ARMGenerator != nil {
		return outboundFirewallRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOutboundFirewallRuleProperties_STATUS_ARM(generators)
	outboundFirewallRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(OutboundFirewallRuleProperties_STATUS_ARM{}), generators)

	return outboundFirewallRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForOutboundFirewallRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOutboundFirewallRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersOutboundFirewallRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersOutboundFirewallRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersOutboundFirewallRule_STATUS_ARM, ServersOutboundFirewallRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersOutboundFirewallRule_STATUS_ARM runs a test to see if a specific instance of ServersOutboundFirewallRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServersOutboundFirewallRule_STATUS_ARM(subject ServersOutboundFirewallRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersOutboundFirewallRule_STATUS_ARM
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

// Generator of ServersOutboundFirewallRule_STATUS_ARM instances for property testing - lazily instantiated by
// ServersOutboundFirewallRule_STATUS_ARMGenerator()
var serversOutboundFirewallRule_STATUS_ARMGenerator gopter.Gen

// ServersOutboundFirewallRule_STATUS_ARMGenerator returns a generator of ServersOutboundFirewallRule_STATUS_ARM instances for property testing.
// We first initialize serversOutboundFirewallRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersOutboundFirewallRule_STATUS_ARMGenerator() gopter.Gen {
	if serversOutboundFirewallRule_STATUS_ARMGenerator != nil {
		return serversOutboundFirewallRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM(generators)
	serversOutboundFirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServersOutboundFirewallRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM(generators)
	serversOutboundFirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServersOutboundFirewallRule_STATUS_ARM{}), generators)

	return serversOutboundFirewallRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersOutboundFirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(OutboundFirewallRuleProperties_STATUS_ARMGenerator())
}
