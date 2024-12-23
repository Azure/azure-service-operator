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

func Test_DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties runs a test to see if a specific instance of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties(subject DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties
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

// Generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties instances for property testing - lazily
// instantiated by DatabaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator()
var databaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator gopter.Gen

// DatabaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator returns a generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties instances for property testing.
func DatabaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator() gopter.Gen {
	if databaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator != nil {
		return databaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties(generators)
	databaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator = gen.Struct(reflect.TypeOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties{}), generators)

	return databaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties(gens map[string]gopter.Gen) {
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Disabled, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled))
	gens["StorageAccountAccessKey"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersDatabasesSecurityAlertPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec, ServersDatabasesSecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec(subject ServersDatabasesSecurityAlertPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy_Spec
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

// Generator of ServersDatabasesSecurityAlertPolicy_Spec instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicy_SpecGenerator()
var serversDatabasesSecurityAlertPolicy_SpecGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicy_SpecGenerator returns a generator of ServersDatabasesSecurityAlertPolicy_Spec instances for property testing.
// We first initialize serversDatabasesSecurityAlertPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesSecurityAlertPolicy_SpecGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicy_SpecGenerator != nil {
		return serversDatabasesSecurityAlertPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	serversDatabasesSecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	serversDatabasesSecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_Spec{}), generators)

	return serversDatabasesSecurityAlertPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyPropertiesGenerator())
}
