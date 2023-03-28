// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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

func Test_ServersDatabasesSecurityAlertPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesSecurityAlertPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersDatabasesSecurityAlertPolicy, ServersDatabasesSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersDatabasesSecurityAlertPolicy tests if a specific instance of ServersDatabasesSecurityAlertPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersDatabasesSecurityAlertPolicy(subject ServersDatabasesSecurityAlertPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersDatabasesSecurityAlertPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersDatabasesSecurityAlertPolicy
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

func Test_ServersDatabasesSecurityAlertPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesSecurityAlertPolicy to ServersDatabasesSecurityAlertPolicy via AssignProperties_To_ServersDatabasesSecurityAlertPolicy & AssignProperties_From_ServersDatabasesSecurityAlertPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersDatabasesSecurityAlertPolicy, ServersDatabasesSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersDatabasesSecurityAlertPolicy tests if a specific instance of ServersDatabasesSecurityAlertPolicy can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServersDatabasesSecurityAlertPolicy(subject ServersDatabasesSecurityAlertPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersDatabasesSecurityAlertPolicy
	err := copied.AssignProperties_To_ServersDatabasesSecurityAlertPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersDatabasesSecurityAlertPolicy
	err = actual.AssignProperties_From_ServersDatabasesSecurityAlertPolicy(&other)
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

func Test_ServersDatabasesSecurityAlertPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy, ServersDatabasesSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy(subject ServersDatabasesSecurityAlertPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy
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

// Generator of ServersDatabasesSecurityAlertPolicy instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicyGenerator()
var serversDatabasesSecurityAlertPolicyGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicyGenerator returns a generator of ServersDatabasesSecurityAlertPolicy instances for property testing.
func ServersDatabasesSecurityAlertPolicyGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicyGenerator != nil {
		return serversDatabasesSecurityAlertPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy(generators)
	serversDatabasesSecurityAlertPolicyGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy{}), generators)

	return serversDatabasesSecurityAlertPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Databases_SecurityAlertPolicy_SpecGenerator()
	gens["Status"] = Servers_Databases_SecurityAlertPolicy_STATUSGenerator()
}

func Test_Servers_Databases_SecurityAlertPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_SecurityAlertPolicy_Spec to Servers_Databases_SecurityAlertPolicy_Spec via AssignProperties_To_Servers_Databases_SecurityAlertPolicy_Spec & AssignProperties_From_Servers_Databases_SecurityAlertPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_Spec, Servers_Databases_SecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_Spec tests if a specific instance of Servers_Databases_SecurityAlertPolicy_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_Spec(subject Servers_Databases_SecurityAlertPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_SecurityAlertPolicy_Spec
	err := copied.AssignProperties_To_Servers_Databases_SecurityAlertPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_SecurityAlertPolicy_Spec
	err = actual.AssignProperties_From_Servers_Databases_SecurityAlertPolicy_Spec(&other)
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

func Test_Servers_Databases_SecurityAlertPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_SecurityAlertPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec, Servers_Databases_SecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec runs a test to see if a specific instance of Servers_Databases_SecurityAlertPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec(subject Servers_Databases_SecurityAlertPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_SecurityAlertPolicy_Spec
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

// Generator of Servers_Databases_SecurityAlertPolicy_Spec instances for property testing - lazily instantiated by
// Servers_Databases_SecurityAlertPolicy_SpecGenerator()
var servers_Databases_SecurityAlertPolicy_SpecGenerator gopter.Gen

// Servers_Databases_SecurityAlertPolicy_SpecGenerator returns a generator of Servers_Databases_SecurityAlertPolicy_Spec instances for property testing.
func Servers_Databases_SecurityAlertPolicy_SpecGenerator() gopter.Gen {
	if servers_Databases_SecurityAlertPolicy_SpecGenerator != nil {
		return servers_Databases_SecurityAlertPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec(generators)
	servers_Databases_SecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_SecurityAlertPolicy_Spec{}), generators)

	return servers_Databases_SecurityAlertPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Disabled, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled))
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Databases_SecurityAlertPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_SecurityAlertPolicy_STATUS to Servers_Databases_SecurityAlertPolicy_STATUS via AssignProperties_To_Servers_Databases_SecurityAlertPolicy_STATUS & AssignProperties_From_Servers_Databases_SecurityAlertPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_STATUS, Servers_Databases_SecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_STATUS tests if a specific instance of Servers_Databases_SecurityAlertPolicy_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_SecurityAlertPolicy_STATUS(subject Servers_Databases_SecurityAlertPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_SecurityAlertPolicy_STATUS
	err := copied.AssignProperties_To_Servers_Databases_SecurityAlertPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_SecurityAlertPolicy_STATUS
	err = actual.AssignProperties_From_Servers_Databases_SecurityAlertPolicy_STATUS(&other)
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

func Test_Servers_Databases_SecurityAlertPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_SecurityAlertPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_STATUS, Servers_Databases_SecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_STATUS runs a test to see if a specific instance of Servers_Databases_SecurityAlertPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_STATUS(subject Servers_Databases_SecurityAlertPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_SecurityAlertPolicy_STATUS
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

// Generator of Servers_Databases_SecurityAlertPolicy_STATUS instances for property testing - lazily instantiated by
// Servers_Databases_SecurityAlertPolicy_STATUSGenerator()
var servers_Databases_SecurityAlertPolicy_STATUSGenerator gopter.Gen

// Servers_Databases_SecurityAlertPolicy_STATUSGenerator returns a generator of Servers_Databases_SecurityAlertPolicy_STATUS instances for property testing.
// We first initialize servers_Databases_SecurityAlertPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_Databases_SecurityAlertPolicy_STATUSGenerator() gopter.Gen {
	if servers_Databases_SecurityAlertPolicy_STATUSGenerator != nil {
		return servers_Databases_SecurityAlertPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS(generators)
	servers_Databases_SecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_SecurityAlertPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS(generators)
	servers_Databases_SecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_SecurityAlertPolicy_STATUS{}), generators)

	return servers_Databases_SecurityAlertPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Disabled, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Enabled))
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
