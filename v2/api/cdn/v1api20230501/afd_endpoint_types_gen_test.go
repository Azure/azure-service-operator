// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/storage"
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

func Test_AfdEndpoint_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdEndpoint to hub returns original",
		prop.ForAll(RunResourceConversionTestForAfdEndpoint, AfdEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAfdEndpoint tests if a specific instance of AfdEndpoint round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAfdEndpoint(subject AfdEndpoint) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.AfdEndpoint
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual AfdEndpoint
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

func Test_AfdEndpoint_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdEndpoint to AfdEndpoint via AssignProperties_To_AfdEndpoint & AssignProperties_From_AfdEndpoint returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdEndpoint, AfdEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdEndpoint tests if a specific instance of AfdEndpoint can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdEndpoint(subject AfdEndpoint) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdEndpoint
	err := copied.AssignProperties_To_AfdEndpoint(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdEndpoint
	err = actual.AssignProperties_From_AfdEndpoint(&other)
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

func Test_AfdEndpoint_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdEndpoint via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdEndpoint, AfdEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdEndpoint runs a test to see if a specific instance of AfdEndpoint round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdEndpoint(subject AfdEndpoint) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdEndpoint
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

// Generator of AfdEndpoint instances for property testing - lazily instantiated by AfdEndpointGenerator()
var afdEndpointGenerator gopter.Gen

// AfdEndpointGenerator returns a generator of AfdEndpoint instances for property testing.
func AfdEndpointGenerator() gopter.Gen {
	if afdEndpointGenerator != nil {
		return afdEndpointGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAfdEndpoint(generators)
	afdEndpointGenerator = gen.Struct(reflect.TypeOf(AfdEndpoint{}), generators)

	return afdEndpointGenerator
}

// AddRelatedPropertyGeneratorsForAfdEndpoint is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdEndpoint(gens map[string]gopter.Gen) {
	gens["Spec"] = AfdEndpoint_SpecGenerator()
	gens["Status"] = AfdEndpoint_STATUSGenerator()
}

func Test_AfdEndpoint_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdEndpoint_STATUS to AfdEndpoint_STATUS via AssignProperties_To_AfdEndpoint_STATUS & AssignProperties_From_AfdEndpoint_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdEndpoint_STATUS, AfdEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdEndpoint_STATUS tests if a specific instance of AfdEndpoint_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdEndpoint_STATUS(subject AfdEndpoint_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdEndpoint_STATUS
	err := copied.AssignProperties_To_AfdEndpoint_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdEndpoint_STATUS
	err = actual.AssignProperties_From_AfdEndpoint_STATUS(&other)
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

func Test_AfdEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdEndpoint_STATUS, AfdEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdEndpoint_STATUS runs a test to see if a specific instance of AfdEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdEndpoint_STATUS(subject AfdEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdEndpoint_STATUS
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

// Generator of AfdEndpoint_STATUS instances for property testing - lazily instantiated by AfdEndpoint_STATUSGenerator()
var afdEndpoint_STATUSGenerator gopter.Gen

// AfdEndpoint_STATUSGenerator returns a generator of AfdEndpoint_STATUS instances for property testing.
// We first initialize afdEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdEndpoint_STATUSGenerator() gopter.Gen {
	if afdEndpoint_STATUSGenerator != nil {
		return afdEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdEndpoint_STATUS(generators)
	afdEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForAfdEndpoint_STATUS(generators)
	afdEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdEndpoint_STATUS{}), generators)

	return afdEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAfdEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["AutoGeneratedDomainNameLabelScope"] = gen.PtrOf(gen.OneConstOf(
		AutoGeneratedDomainNameLabelScope_STATUS_NoReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_ResourceGroupReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_SubscriptionReuse,
		AutoGeneratedDomainNameLabelScope_STATUS_TenantReuse))
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		AFDEndpointProperties_DeploymentStatus_STATUS_Failed,
		AFDEndpointProperties_DeploymentStatus_STATUS_InProgress,
		AFDEndpointProperties_DeploymentStatus_STATUS_NotStarted,
		AFDEndpointProperties_DeploymentStatus_STATUS_Succeeded))
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDEndpointProperties_EnabledState_STATUS_Disabled, AFDEndpointProperties_EnabledState_STATUS_Enabled))
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		AFDEndpointProperties_ProvisioningState_STATUS_Creating,
		AFDEndpointProperties_ProvisioningState_STATUS_Deleting,
		AFDEndpointProperties_ProvisioningState_STATUS_Failed,
		AFDEndpointProperties_ProvisioningState_STATUS_Succeeded,
		AFDEndpointProperties_ProvisioningState_STATUS_Updating))
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAfdEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_AfdEndpoint_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdEndpoint_Spec to AfdEndpoint_Spec via AssignProperties_To_AfdEndpoint_Spec & AssignProperties_From_AfdEndpoint_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdEndpoint_Spec, AfdEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdEndpoint_Spec tests if a specific instance of AfdEndpoint_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdEndpoint_Spec(subject AfdEndpoint_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdEndpoint_Spec
	err := copied.AssignProperties_To_AfdEndpoint_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdEndpoint_Spec
	err = actual.AssignProperties_From_AfdEndpoint_Spec(&other)
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

func Test_AfdEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdEndpoint_Spec, AfdEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdEndpoint_Spec runs a test to see if a specific instance of AfdEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdEndpoint_Spec(subject AfdEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdEndpoint_Spec
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

// Generator of AfdEndpoint_Spec instances for property testing - lazily instantiated by AfdEndpoint_SpecGenerator()
var afdEndpoint_SpecGenerator gopter.Gen

// AfdEndpoint_SpecGenerator returns a generator of AfdEndpoint_Spec instances for property testing.
func AfdEndpoint_SpecGenerator() gopter.Gen {
	if afdEndpoint_SpecGenerator != nil {
		return afdEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdEndpoint_Spec(generators)
	afdEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(AfdEndpoint_Spec{}), generators)

	return afdEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAfdEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["AutoGeneratedDomainNameLabelScope"] = gen.PtrOf(gen.OneConstOf(
		AutoGeneratedDomainNameLabelScope_NoReuse,
		AutoGeneratedDomainNameLabelScope_ResourceGroupReuse,
		AutoGeneratedDomainNameLabelScope_SubscriptionReuse,
		AutoGeneratedDomainNameLabelScope_TenantReuse))
	gens["AzureName"] = gen.AlphaString()
	gens["EnabledState"] = gen.PtrOf(gen.OneConstOf(AFDEndpointProperties_EnabledState_Disabled, AFDEndpointProperties_EnabledState_Enabled))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}
