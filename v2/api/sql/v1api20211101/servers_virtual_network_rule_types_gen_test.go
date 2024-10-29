// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
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

func Test_ServersVirtualNetworkRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersVirtualNetworkRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersVirtualNetworkRule, ServersVirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersVirtualNetworkRule tests if a specific instance of ServersVirtualNetworkRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersVirtualNetworkRule(subject ServersVirtualNetworkRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.ServersVirtualNetworkRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersVirtualNetworkRule
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

func Test_ServersVirtualNetworkRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersVirtualNetworkRule to ServersVirtualNetworkRule via AssignProperties_To_ServersVirtualNetworkRule & AssignProperties_From_ServersVirtualNetworkRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersVirtualNetworkRule, ServersVirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersVirtualNetworkRule tests if a specific instance of ServersVirtualNetworkRule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersVirtualNetworkRule(subject ServersVirtualNetworkRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServersVirtualNetworkRule
	err := copied.AssignProperties_To_ServersVirtualNetworkRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersVirtualNetworkRule
	err = actual.AssignProperties_From_ServersVirtualNetworkRule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersVirtualNetworkRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule, ServersVirtualNetworkRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule runs a test to see if a specific instance of ServersVirtualNetworkRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule(subject ServersVirtualNetworkRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule
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

// Generator of ServersVirtualNetworkRule instances for property testing - lazily instantiated by
// ServersVirtualNetworkRuleGenerator()
var serversVirtualNetworkRuleGenerator gopter.Gen

// ServersVirtualNetworkRuleGenerator returns a generator of ServersVirtualNetworkRule instances for property testing.
func ServersVirtualNetworkRuleGenerator() gopter.Gen {
	if serversVirtualNetworkRuleGenerator != nil {
		return serversVirtualNetworkRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersVirtualNetworkRule(generators)
	serversVirtualNetworkRuleGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule{}), generators)

	return serversVirtualNetworkRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersVirtualNetworkRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersVirtualNetworkRule(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersVirtualNetworkRule_SpecGenerator()
	gens["Status"] = ServersVirtualNetworkRule_STATUSGenerator()
}

func Test_ServersVirtualNetworkRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersVirtualNetworkRule_STATUS to ServersVirtualNetworkRule_STATUS via AssignProperties_To_ServersVirtualNetworkRule_STATUS & AssignProperties_From_ServersVirtualNetworkRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersVirtualNetworkRule_STATUS, ServersVirtualNetworkRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersVirtualNetworkRule_STATUS tests if a specific instance of ServersVirtualNetworkRule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersVirtualNetworkRule_STATUS(subject ServersVirtualNetworkRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServersVirtualNetworkRule_STATUS
	err := copied.AssignProperties_To_ServersVirtualNetworkRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersVirtualNetworkRule_STATUS
	err = actual.AssignProperties_From_ServersVirtualNetworkRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersVirtualNetworkRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule_STATUS, ServersVirtualNetworkRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule_STATUS runs a test to see if a specific instance of ServersVirtualNetworkRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule_STATUS(subject ServersVirtualNetworkRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule_STATUS
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

// Generator of ServersVirtualNetworkRule_STATUS instances for property testing - lazily instantiated by
// ServersVirtualNetworkRule_STATUSGenerator()
var serversVirtualNetworkRule_STATUSGenerator gopter.Gen

// ServersVirtualNetworkRule_STATUSGenerator returns a generator of ServersVirtualNetworkRule_STATUS instances for property testing.
func ServersVirtualNetworkRule_STATUSGenerator() gopter.Gen {
	if serversVirtualNetworkRule_STATUSGenerator != nil {
		return serversVirtualNetworkRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS(generators)
	serversVirtualNetworkRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule_STATUS{}), generators)

	return serversVirtualNetworkRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkRuleProperties_State_STATUS_Deleting,
		VirtualNetworkRuleProperties_State_STATUS_Failed,
		VirtualNetworkRuleProperties_State_STATUS_InProgress,
		VirtualNetworkRuleProperties_State_STATUS_Initializing,
		VirtualNetworkRuleProperties_State_STATUS_Ready,
		VirtualNetworkRuleProperties_State_STATUS_Unknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkSubnetId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersVirtualNetworkRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersVirtualNetworkRule_Spec to ServersVirtualNetworkRule_Spec via AssignProperties_To_ServersVirtualNetworkRule_Spec & AssignProperties_From_ServersVirtualNetworkRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersVirtualNetworkRule_Spec, ServersVirtualNetworkRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersVirtualNetworkRule_Spec tests if a specific instance of ServersVirtualNetworkRule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersVirtualNetworkRule_Spec(subject ServersVirtualNetworkRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.ServersVirtualNetworkRule_Spec
	err := copied.AssignProperties_To_ServersVirtualNetworkRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersVirtualNetworkRule_Spec
	err = actual.AssignProperties_From_ServersVirtualNetworkRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersVirtualNetworkRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersVirtualNetworkRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersVirtualNetworkRule_Spec, ServersVirtualNetworkRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersVirtualNetworkRule_Spec runs a test to see if a specific instance of ServersVirtualNetworkRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersVirtualNetworkRule_Spec(subject ServersVirtualNetworkRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersVirtualNetworkRule_Spec
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

// Generator of ServersVirtualNetworkRule_Spec instances for property testing - lazily instantiated by
// ServersVirtualNetworkRule_SpecGenerator()
var serversVirtualNetworkRule_SpecGenerator gopter.Gen

// ServersVirtualNetworkRule_SpecGenerator returns a generator of ServersVirtualNetworkRule_Spec instances for property testing.
func ServersVirtualNetworkRule_SpecGenerator() gopter.Gen {
	if serversVirtualNetworkRule_SpecGenerator != nil {
		return serversVirtualNetworkRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec(generators)
	serversVirtualNetworkRule_SpecGenerator = gen.Struct(reflect.TypeOf(ServersVirtualNetworkRule_Spec{}), generators)

	return serversVirtualNetworkRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersVirtualNetworkRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
}
