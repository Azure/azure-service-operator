// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"encoding/json"
	alpha20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1alpha1api20210601storage"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
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

func Test_FlexibleServersFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersFirewallRule tests if a specific instance of FlexibleServersFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210601s.FlexibleServersFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersFirewallRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersFirewallRule to FlexibleServersFirewallRule via AssignPropertiesToFlexibleServersFirewallRule & AssignPropertiesFromFlexibleServersFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersFirewallRule tests if a specific instance of FlexibleServersFirewallRule can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210601s.FlexibleServersFirewallRule
	err := copied.AssignPropertiesToFlexibleServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersFirewallRule
	err = actual.AssignPropertiesFromFlexibleServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
//FlexibleServersFirewallRuleGenerator()
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
	gens["Spec"] = FlexibleServersFirewallRulesSpecGenerator()
	gens["Status"] = FirewallRuleStatusGenerator()
}

func Test_FirewallRule_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FirewallRule_Status to FirewallRule_Status via AssignPropertiesToFirewallRuleStatus & AssignPropertiesFromFirewallRuleStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForFirewallRuleStatus, FirewallRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFirewallRuleStatus tests if a specific instance of FirewallRule_Status can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFirewallRuleStatus(subject FirewallRule_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210601s.FirewallRule_Status
	err := copied.AssignPropertiesToFirewallRuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FirewallRule_Status
	err = actual.AssignPropertiesFromFirewallRuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FirewallRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleStatus, FirewallRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleStatus runs a test to see if a specific instance of FirewallRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleStatus(subject FirewallRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_Status
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

// Generator of FirewallRule_Status instances for property testing - lazily instantiated by FirewallRuleStatusGenerator()
var firewallRuleStatusGenerator gopter.Gen

// FirewallRuleStatusGenerator returns a generator of FirewallRule_Status instances for property testing.
// We first initialize firewallRuleStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRuleStatusGenerator() gopter.Gen {
	if firewallRuleStatusGenerator != nil {
		return firewallRuleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleStatus(generators)
	firewallRuleStatusGenerator = gen.Struct(reflect.TypeOf(FirewallRule_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleStatus(generators)
	AddRelatedPropertyGeneratorsForFirewallRuleStatus(generators)
	firewallRuleStatusGenerator = gen.Struct(reflect.TypeOf(FirewallRule_Status{}), generators)

	return firewallRuleStatusGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleStatus(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRuleStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRuleStatus(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_FlexibleServersFirewallRules_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersFirewallRules_Spec to FlexibleServersFirewallRules_Spec via AssignPropertiesToFlexibleServersFirewallRulesSpec & AssignPropertiesFromFlexibleServersFirewallRulesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersFirewallRulesSpec, FlexibleServersFirewallRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersFirewallRulesSpec tests if a specific instance of FlexibleServersFirewallRules_Spec can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersFirewallRulesSpec(subject FlexibleServersFirewallRules_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210601s.FlexibleServersFirewallRules_Spec
	err := copied.AssignPropertiesToFlexibleServersFirewallRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersFirewallRules_Spec
	err = actual.AssignPropertiesFromFlexibleServersFirewallRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersFirewallRules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRulesSpec, FlexibleServersFirewallRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRulesSpec runs a test to see if a specific instance of FlexibleServersFirewallRules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRulesSpec(subject FlexibleServersFirewallRules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRules_Spec
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

// Generator of FlexibleServersFirewallRules_Spec instances for property testing - lazily instantiated by
//FlexibleServersFirewallRulesSpecGenerator()
var flexibleServersFirewallRulesSpecGenerator gopter.Gen

// FlexibleServersFirewallRulesSpecGenerator returns a generator of FlexibleServersFirewallRules_Spec instances for property testing.
func FlexibleServersFirewallRulesSpecGenerator() gopter.Gen {
	if flexibleServersFirewallRulesSpecGenerator != nil {
		return flexibleServersFirewallRulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpec(generators)
	flexibleServersFirewallRulesSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRules_Spec{}), generators)

	return flexibleServersFirewallRulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRulesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
