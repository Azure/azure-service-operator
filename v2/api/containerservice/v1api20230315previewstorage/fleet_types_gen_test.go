// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315previewstorage

import (
	"encoding/json"
	v20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202previewstorage"
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

func Test_Fleet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleet, FleetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleet runs a test to see if a specific instance of Fleet round trips to JSON and back losslessly
func RunJSONSerializationTestForFleet(subject Fleet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleet
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

// Generator of Fleet instances for property testing - lazily instantiated by FleetGenerator()
var fleetGenerator gopter.Gen

// FleetGenerator returns a generator of Fleet instances for property testing.
func FleetGenerator() gopter.Gen {
	if fleetGenerator != nil {
		return fleetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFleet(generators)
	fleetGenerator = gen.Struct(reflect.TypeOf(Fleet{}), generators)

	return fleetGenerator
}

// AddRelatedPropertyGeneratorsForFleet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleet(gens map[string]gopter.Gen) {
	gens["Spec"] = Fleet_SpecGenerator()
	gens["Status"] = Fleet_STATUSGenerator()
}

func Test_Fleet_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleet_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleet_Spec, Fleet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleet_Spec runs a test to see if a specific instance of Fleet_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFleet_Spec(subject Fleet_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleet_Spec
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

// Generator of Fleet_Spec instances for property testing - lazily instantiated by Fleet_SpecGenerator()
var fleet_SpecGenerator gopter.Gen

// Fleet_SpecGenerator returns a generator of Fleet_Spec instances for property testing.
// We first initialize fleet_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Fleet_SpecGenerator() gopter.Gen {
	if fleet_SpecGenerator != nil {
		return fleet_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleet_Spec(generators)
	fleet_SpecGenerator = gen.Struct(reflect.TypeOf(Fleet_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleet_Spec(generators)
	AddRelatedPropertyGeneratorsForFleet_Spec(generators)
	fleet_SpecGenerator = gen.Struct(reflect.TypeOf(Fleet_Spec{}), generators)

	return fleet_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFleet_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleet_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFleet_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleet_Spec(gens map[string]gopter.Gen) {
	gens["HubProfile"] = gen.PtrOf(FleetHubProfileGenerator())
}

func Test_Fleet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleet_STATUS, Fleet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleet_STATUS runs a test to see if a specific instance of Fleet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFleet_STATUS(subject Fleet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleet_STATUS
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

// Generator of Fleet_STATUS instances for property testing - lazily instantiated by Fleet_STATUSGenerator()
var fleet_STATUSGenerator gopter.Gen

// Fleet_STATUSGenerator returns a generator of Fleet_STATUS instances for property testing.
// We first initialize fleet_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Fleet_STATUSGenerator() gopter.Gen {
	if fleet_STATUSGenerator != nil {
		return fleet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleet_STATUS(generators)
	fleet_STATUSGenerator = gen.Struct(reflect.TypeOf(Fleet_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleet_STATUS(generators)
	AddRelatedPropertyGeneratorsForFleet_STATUS(generators)
	fleet_STATUSGenerator = gen.Struct(reflect.TypeOf(Fleet_STATUS{}), generators)

	return fleet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFleet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleet_STATUS(gens map[string]gopter.Gen) {
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFleet_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleet_STATUS(gens map[string]gopter.Gen) {
	gens["HubProfile"] = gen.PtrOf(FleetHubProfile_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FleetHubProfile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetHubProfile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetHubProfile, FleetHubProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetHubProfile runs a test to see if a specific instance of FleetHubProfile round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetHubProfile(subject FleetHubProfile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetHubProfile
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

// Generator of FleetHubProfile instances for property testing - lazily instantiated by FleetHubProfileGenerator()
var fleetHubProfileGenerator gopter.Gen

// FleetHubProfileGenerator returns a generator of FleetHubProfile instances for property testing.
func FleetHubProfileGenerator() gopter.Gen {
	if fleetHubProfileGenerator != nil {
		return fleetHubProfileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetHubProfile(generators)
	fleetHubProfileGenerator = gen.Struct(reflect.TypeOf(FleetHubProfile{}), generators)

	return fleetHubProfileGenerator
}

// AddIndependentPropertyGeneratorsForFleetHubProfile is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetHubProfile(gens map[string]gopter.Gen) {
	gens["DnsPrefix"] = gen.PtrOf(gen.AlphaString())
}

func Test_FleetHubProfile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetHubProfile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetHubProfile_STATUS, FleetHubProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetHubProfile_STATUS runs a test to see if a specific instance of FleetHubProfile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetHubProfile_STATUS(subject FleetHubProfile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetHubProfile_STATUS
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

// Generator of FleetHubProfile_STATUS instances for property testing - lazily instantiated by
// FleetHubProfile_STATUSGenerator()
var fleetHubProfile_STATUSGenerator gopter.Gen

// FleetHubProfile_STATUSGenerator returns a generator of FleetHubProfile_STATUS instances for property testing.
func FleetHubProfile_STATUSGenerator() gopter.Gen {
	if fleetHubProfile_STATUSGenerator != nil {
		return fleetHubProfile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleetHubProfile_STATUS(generators)
	fleetHubProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(FleetHubProfile_STATUS{}), generators)

	return fleetHubProfile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFleetHubProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleetHubProfile_STATUS(gens map[string]gopter.Gen) {
	gens["DnsPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["KubernetesVersion"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_STATUS to SystemData_STATUS via AssignProperties_To_SystemData_STATUS & AssignProperties_From_SystemData_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_STATUS tests if a specific instance of SystemData_STATUS can be assigned to v1api20230202previewstorage and back losslessly
func RunPropertyAssignmentTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230202ps.SystemData_STATUS
	err := copied.AssignProperties_To_SystemData_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_STATUS
	err = actual.AssignProperties_From_SystemData_STATUS(&other)
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

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}
