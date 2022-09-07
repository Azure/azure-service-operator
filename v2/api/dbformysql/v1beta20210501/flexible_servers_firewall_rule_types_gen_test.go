// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import (
	"encoding/json"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501storage"
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
	parameters.MinSuccessfulTests = 10
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
	var hub v20210501s.FlexibleServersFirewallRule
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
		"Round trip from FlexibleServersFirewallRule to FlexibleServersFirewallRule via AssignProperties_To_FlexibleServersFirewallRule & AssignProperties_From_FlexibleServersFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersFirewallRule tests if a specific instance of FlexibleServersFirewallRule can be assigned to v1beta20210501storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210501s.FlexibleServersFirewallRule
	err := copied.AssignProperties_To_FlexibleServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersFirewallRule
	err = actual.AssignProperties_From_FlexibleServersFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
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
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
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
// FlexibleServersFirewallRuleGenerator()
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
	gens["Spec"] = FlexibleServers_FirewallRule_SpecGenerator()
	gens["Status"] = FirewallRule_STATUSGenerator()
}

func Test_FirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FirewallRule_STATUS to FirewallRule_STATUS via AssignProperties_To_FirewallRule_STATUS & AssignProperties_From_FirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFirewallRule_STATUS, FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFirewallRule_STATUS tests if a specific instance of FirewallRule_STATUS can be assigned to v1beta20210501storage and back losslessly
func RunPropertyAssignmentTestForFirewallRule_STATUS(subject FirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210501s.FirewallRule_STATUS
	err := copied.AssignProperties_To_FirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FirewallRule_STATUS
	err = actual.AssignProperties_From_FirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRule_STATUS, FirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRule_STATUS runs a test to see if a specific instance of FirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRule_STATUS(subject FirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_STATUS
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

// Generator of FirewallRule_STATUS instances for property testing - lazily instantiated by
// FirewallRule_STATUSGenerator()
var firewallRule_STATUSGenerator gopter.Gen

// FirewallRule_STATUSGenerator returns a generator of FirewallRule_STATUS instances for property testing.
// We first initialize firewallRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRule_STATUSGenerator() gopter.Gen {
	if firewallRule_STATUSGenerator != nil {
		return firewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_STATUS(generators)
	firewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForFirewallRule_STATUS(generators)
	firewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUS{}), generators)

	return firewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServers_FirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_FirewallRule_Spec to FlexibleServers_FirewallRule_Spec via AssignProperties_To_FlexibleServers_FirewallRule_Spec & AssignProperties_From_FlexibleServers_FirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec, FlexibleServers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec tests if a specific instance of FlexibleServers_FirewallRule_Spec can be assigned to v1beta20210501storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_FirewallRule_Spec(subject FlexibleServers_FirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210501s.FlexibleServers_FirewallRule_Spec
	err := copied.AssignProperties_To_FlexibleServers_FirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_FirewallRule_Spec
	err = actual.AssignProperties_From_FlexibleServers_FirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServers_FirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_FirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec, FlexibleServers_FirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec runs a test to see if a specific instance of FlexibleServers_FirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_FirewallRule_Spec(subject FlexibleServers_FirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_FirewallRule_Spec
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

// Generator of FlexibleServers_FirewallRule_Spec instances for property testing - lazily instantiated by
// FlexibleServers_FirewallRule_SpecGenerator()
var flexibleServers_FirewallRule_SpecGenerator gopter.Gen

// FlexibleServers_FirewallRule_SpecGenerator returns a generator of FlexibleServers_FirewallRule_Spec instances for property testing.
func FlexibleServers_FirewallRule_SpecGenerator() gopter.Gen {
	if flexibleServers_FirewallRule_SpecGenerator != nil {
		return flexibleServers_FirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec(generators)
	flexibleServers_FirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_FirewallRule_Spec{}), generators)

	return flexibleServers_FirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_FirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
