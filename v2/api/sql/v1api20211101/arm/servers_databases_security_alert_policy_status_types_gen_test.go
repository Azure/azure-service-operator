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

func Test_DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS runs a test to see if a specific instance of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS(subject DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS
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

// Generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS instances for property testing -
// lazily instantiated by DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator()
var databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator gopter.Gen

// DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator returns a generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS instances for property testing.
func DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator() gopter.Gen {
	if databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator != nil {
		return databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS(generators)
	databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS{}), generators)

	return databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Disabled, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Enabled))
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersDatabasesSecurityAlertPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS, ServersDatabasesSecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS(subject ServersDatabasesSecurityAlertPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy_STATUS
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

// Generator of ServersDatabasesSecurityAlertPolicy_STATUS instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicy_STATUSGenerator()
var serversDatabasesSecurityAlertPolicy_STATUSGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicy_STATUSGenerator returns a generator of ServersDatabasesSecurityAlertPolicy_STATUS instances for property testing.
// We first initialize serversDatabasesSecurityAlertPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesSecurityAlertPolicy_STATUSGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicy_STATUSGenerator != nil {
		return serversDatabasesSecurityAlertPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	serversDatabasesSecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	serversDatabasesSecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_STATUS{}), generators)

	return serversDatabasesSecurityAlertPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
