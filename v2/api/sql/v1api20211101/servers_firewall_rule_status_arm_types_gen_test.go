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

func Test_ServerFirewallRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerFirewallRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerFirewallRuleProperties_STATUS_ARM, ServerFirewallRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerFirewallRuleProperties_STATUS_ARM runs a test to see if a specific instance of ServerFirewallRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerFirewallRuleProperties_STATUS_ARM(subject ServerFirewallRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerFirewallRuleProperties_STATUS_ARM
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

// Generator of ServerFirewallRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ServerFirewallRuleProperties_STATUS_ARMGenerator()
var serverFirewallRuleProperties_STATUS_ARMGenerator gopter.Gen

// ServerFirewallRuleProperties_STATUS_ARMGenerator returns a generator of ServerFirewallRuleProperties_STATUS_ARM instances for property testing.
func ServerFirewallRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if serverFirewallRuleProperties_STATUS_ARMGenerator != nil {
		return serverFirewallRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerFirewallRuleProperties_STATUS_ARM(generators)
	serverFirewallRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ServerFirewallRuleProperties_STATUS_ARM{}), generators)

	return serverFirewallRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerFirewallRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerFirewallRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_FirewallRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_FirewallRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_FirewallRule_STATUS_ARM, Servers_FirewallRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_FirewallRule_STATUS_ARM runs a test to see if a specific instance of Servers_FirewallRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_FirewallRule_STATUS_ARM(subject Servers_FirewallRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_FirewallRule_STATUS_ARM
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

// Generator of Servers_FirewallRule_STATUS_ARM instances for property testing - lazily instantiated by
// Servers_FirewallRule_STATUS_ARMGenerator()
var servers_FirewallRule_STATUS_ARMGenerator gopter.Gen

// Servers_FirewallRule_STATUS_ARMGenerator returns a generator of Servers_FirewallRule_STATUS_ARM instances for property testing.
// We first initialize servers_FirewallRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_FirewallRule_STATUS_ARMGenerator() gopter.Gen {
	if servers_FirewallRule_STATUS_ARMGenerator != nil {
		return servers_FirewallRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS_ARM(generators)
	servers_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_FirewallRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_FirewallRule_STATUS_ARM(generators)
	servers_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_FirewallRule_STATUS_ARM{}), generators)

	return servers_FirewallRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerFirewallRuleProperties_STATUS_ARMGenerator())
}
