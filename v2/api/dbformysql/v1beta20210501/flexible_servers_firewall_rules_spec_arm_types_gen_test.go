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

<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rule__spec_arm_types_gen_test.go
func Test_FlexibleServersFirewallRule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_FlexibleServers_FirewallRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rules_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rule__spec_arm_types_gen_test.go
		"Round trip of FlexibleServersFirewallRule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM, FlexibleServersFirewallRule_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM runs a test to see if a specific instance of FlexibleServersFirewallRule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM(subject FlexibleServersFirewallRule_SpecARM) string {
========
		"Round trip of FlexibleServers_FirewallRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_FirewallRules_SpecARM, FlexibleServers_FirewallRules_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_FirewallRules_SpecARM runs a test to see if a specific instance of FlexibleServers_FirewallRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_FirewallRules_SpecARM(subject FlexibleServers_FirewallRules_SpecARM) string {
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rules_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rule__spec_arm_types_gen_test.go
	var actual FlexibleServersFirewallRule_SpecARM
========
	var actual FlexibleServers_FirewallRules_SpecARM
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rules_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rule__spec_arm_types_gen_test.go
// Generator of FlexibleServersFirewallRule_SpecARM instances for property testing - lazily instantiated by
// FlexibleServersFirewallRule_SpecARMGenerator()
var flexibleServersFirewallRule_SpecARMGenerator gopter.Gen

// FlexibleServersFirewallRule_SpecARMGenerator returns a generator of FlexibleServersFirewallRule_SpecARM instances for property testing.
// We first initialize flexibleServersFirewallRule_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersFirewallRule_SpecARMGenerator() gopter.Gen {
	if flexibleServersFirewallRule_SpecARMGenerator != nil {
		return flexibleServersFirewallRule_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	flexibleServersFirewallRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	flexibleServersFirewallRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_SpecARM{}), generators)

	return flexibleServersFirewallRule_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of FlexibleServers_FirewallRules_SpecARM instances for property testing - lazily instantiated by
// FlexibleServers_FirewallRules_SpecARMGenerator()
var flexibleServers_FirewallRules_SpecARMGenerator gopter.Gen

// FlexibleServers_FirewallRules_SpecARMGenerator returns a generator of FlexibleServers_FirewallRules_SpecARM instances for property testing.
// We first initialize flexibleServers_FirewallRules_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_FirewallRules_SpecARMGenerator() gopter.Gen {
	if flexibleServers_FirewallRules_SpecARMGenerator != nil {
		return flexibleServers_FirewallRules_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM(generators)
	flexibleServers_FirewallRules_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM(generators)
	flexibleServers_FirewallRules_SpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRules_SpecARM{}), generators)

	return flexibleServers_FirewallRules_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rules_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
}

<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rule__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_FirewallRules_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_firewall_rules_spec_arm_types_gen_test.go
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
