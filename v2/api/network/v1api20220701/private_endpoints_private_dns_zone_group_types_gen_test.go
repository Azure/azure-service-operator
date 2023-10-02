// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"encoding/json"
	v20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/storage"
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

func Test_PrivateEndpointsPrivateDnsZoneGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointsPrivateDnsZoneGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateEndpointsPrivateDnsZoneGroup, PrivateEndpointsPrivateDnsZoneGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateEndpointsPrivateDnsZoneGroup tests if a specific instance of PrivateEndpointsPrivateDnsZoneGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateEndpointsPrivateDnsZoneGroup(subject PrivateEndpointsPrivateDnsZoneGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220701s.PrivateEndpointsPrivateDnsZoneGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateEndpointsPrivateDnsZoneGroup
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

func Test_PrivateEndpointsPrivateDnsZoneGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointsPrivateDnsZoneGroup to PrivateEndpointsPrivateDnsZoneGroup via AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup & AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup, PrivateEndpointsPrivateDnsZoneGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup tests if a specific instance of PrivateEndpointsPrivateDnsZoneGroup can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup(subject PrivateEndpointsPrivateDnsZoneGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.PrivateEndpointsPrivateDnsZoneGroup
	err := copied.AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointsPrivateDnsZoneGroup
	err = actual.AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup(&other)
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

func Test_PrivateEndpointsPrivateDnsZoneGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointsPrivateDnsZoneGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup, PrivateEndpointsPrivateDnsZoneGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup runs a test to see if a specific instance of PrivateEndpointsPrivateDnsZoneGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup(subject PrivateEndpointsPrivateDnsZoneGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointsPrivateDnsZoneGroup
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

// Generator of PrivateEndpointsPrivateDnsZoneGroup instances for property testing - lazily instantiated by
// PrivateEndpointsPrivateDnsZoneGroupGenerator()
var privateEndpointsPrivateDnsZoneGroupGenerator gopter.Gen

// PrivateEndpointsPrivateDnsZoneGroupGenerator returns a generator of PrivateEndpointsPrivateDnsZoneGroup instances for property testing.
func PrivateEndpointsPrivateDnsZoneGroupGenerator() gopter.Gen {
	if privateEndpointsPrivateDnsZoneGroupGenerator != nil {
		return privateEndpointsPrivateDnsZoneGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup(generators)
	privateEndpointsPrivateDnsZoneGroupGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup{}), generators)

	return privateEndpointsPrivateDnsZoneGroupGenerator
}

// AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator()
	gens["Status"] = PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator()
}

func Test_PrivateEndpoints_PrivateDnsZoneGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpoints_PrivateDnsZoneGroup_Spec to PrivateEndpoints_PrivateDnsZoneGroup_Spec via AssignProperties_To_PrivateEndpoints_PrivateDnsZoneGroup_Spec & AssignProperties_From_PrivateEndpoints_PrivateDnsZoneGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec, PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec tests if a specific instance of PrivateEndpoints_PrivateDnsZoneGroup_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec(subject PrivateEndpoints_PrivateDnsZoneGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.PrivateEndpoints_PrivateDnsZoneGroup_Spec
	err := copied.AssignProperties_To_PrivateEndpoints_PrivateDnsZoneGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpoints_PrivateDnsZoneGroup_Spec
	err = actual.AssignProperties_From_PrivateEndpoints_PrivateDnsZoneGroup_Spec(&other)
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

func Test_PrivateEndpoints_PrivateDnsZoneGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpoints_PrivateDnsZoneGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec, PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec runs a test to see if a specific instance of PrivateEndpoints_PrivateDnsZoneGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_Spec(subject PrivateEndpoints_PrivateDnsZoneGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpoints_PrivateDnsZoneGroup_Spec
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

// Generator of PrivateEndpoints_PrivateDnsZoneGroup_Spec instances for property testing - lazily instantiated by
// PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator()
var privateEndpoints_PrivateDnsZoneGroup_SpecGenerator gopter.Gen

// PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator returns a generator of PrivateEndpoints_PrivateDnsZoneGroup_Spec instances for property testing.
// We first initialize privateEndpoints_PrivateDnsZoneGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpoints_PrivateDnsZoneGroup_SpecGenerator() gopter.Gen {
	if privateEndpoints_PrivateDnsZoneGroup_SpecGenerator != nil {
		return privateEndpoints_PrivateDnsZoneGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec(generators)
	privateEndpoints_PrivateDnsZoneGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec(generators)
	privateEndpoints_PrivateDnsZoneGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_Spec{}), generators)

	return privateEndpoints_PrivateDnsZoneGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_Spec(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfigGenerator())
}

func Test_PrivateEndpoints_PrivateDnsZoneGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpoints_PrivateDnsZoneGroup_STATUS to PrivateEndpoints_PrivateDnsZoneGroup_STATUS via AssignProperties_To_PrivateEndpoints_PrivateDnsZoneGroup_STATUS & AssignProperties_From_PrivateEndpoints_PrivateDnsZoneGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS, PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS tests if a specific instance of PrivateEndpoints_PrivateDnsZoneGroup_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(subject PrivateEndpoints_PrivateDnsZoneGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.PrivateEndpoints_PrivateDnsZoneGroup_STATUS
	err := copied.AssignProperties_To_PrivateEndpoints_PrivateDnsZoneGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpoints_PrivateDnsZoneGroup_STATUS
	err = actual.AssignProperties_From_PrivateEndpoints_PrivateDnsZoneGroup_STATUS(&other)
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

func Test_PrivateEndpoints_PrivateDnsZoneGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpoints_PrivateDnsZoneGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS, PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS runs a test to see if a specific instance of PrivateEndpoints_PrivateDnsZoneGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(subject PrivateEndpoints_PrivateDnsZoneGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpoints_PrivateDnsZoneGroup_STATUS
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

// Generator of PrivateEndpoints_PrivateDnsZoneGroup_STATUS instances for property testing - lazily instantiated by
// PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator()
var privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator gopter.Gen

// PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator returns a generator of PrivateEndpoints_PrivateDnsZoneGroup_STATUS instances for property testing.
// We first initialize privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpoints_PrivateDnsZoneGroup_STATUSGenerator() gopter.Gen {
	if privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator != nil {
		return privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(generators)
	privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(generators)
	privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoints_PrivateDnsZoneGroup_STATUS{}), generators)

	return privateEndpoints_PrivateDnsZoneGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointProvisioningState_STATUS_Deleting,
		PrivateEndpointProvisioningState_STATUS_Failed,
		PrivateEndpointProvisioningState_STATUS_Succeeded,
		PrivateEndpointProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpoints_PrivateDnsZoneGroup_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfig_STATUSGenerator())
}

func Test_PrivateDnsZoneConfig_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZoneConfig to PrivateDnsZoneConfig via AssignProperties_To_PrivateDnsZoneConfig & AssignProperties_From_PrivateDnsZoneConfig returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZoneConfig, PrivateDnsZoneConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZoneConfig tests if a specific instance of PrivateDnsZoneConfig can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZoneConfig(subject PrivateDnsZoneConfig) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.PrivateDnsZoneConfig
	err := copied.AssignProperties_To_PrivateDnsZoneConfig(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZoneConfig
	err = actual.AssignProperties_From_PrivateDnsZoneConfig(&other)
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

func Test_PrivateDnsZoneConfig_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneConfig via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneConfig, PrivateDnsZoneConfigGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneConfig runs a test to see if a specific instance of PrivateDnsZoneConfig round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneConfig(subject PrivateDnsZoneConfig) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneConfig
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

// Generator of PrivateDnsZoneConfig instances for property testing - lazily instantiated by
// PrivateDnsZoneConfigGenerator()
var privateDnsZoneConfigGenerator gopter.Gen

// PrivateDnsZoneConfigGenerator returns a generator of PrivateDnsZoneConfig instances for property testing.
func PrivateDnsZoneConfigGenerator() gopter.Gen {
	if privateDnsZoneConfigGenerator != nil {
		return privateDnsZoneConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig(generators)
	privateDnsZoneConfigGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig{}), generators)

	return privateDnsZoneConfigGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateDnsZoneConfig_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZoneConfig_STATUS to PrivateDnsZoneConfig_STATUS via AssignProperties_To_PrivateDnsZoneConfig_STATUS & AssignProperties_From_PrivateDnsZoneConfig_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZoneConfig_STATUS, PrivateDnsZoneConfig_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZoneConfig_STATUS tests if a specific instance of PrivateDnsZoneConfig_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZoneConfig_STATUS(subject PrivateDnsZoneConfig_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.PrivateDnsZoneConfig_STATUS
	err := copied.AssignProperties_To_PrivateDnsZoneConfig_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZoneConfig_STATUS
	err = actual.AssignProperties_From_PrivateDnsZoneConfig_STATUS(&other)
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

func Test_PrivateDnsZoneConfig_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneConfig_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS, PrivateDnsZoneConfig_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS runs a test to see if a specific instance of PrivateDnsZoneConfig_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneConfig_STATUS(subject PrivateDnsZoneConfig_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneConfig_STATUS
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

// Generator of PrivateDnsZoneConfig_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZoneConfig_STATUSGenerator()
var privateDnsZoneConfig_STATUSGenerator gopter.Gen

// PrivateDnsZoneConfig_STATUSGenerator returns a generator of PrivateDnsZoneConfig_STATUS instances for property testing.
// We first initialize privateDnsZoneConfig_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZoneConfig_STATUSGenerator() gopter.Gen {
	if privateDnsZoneConfig_STATUSGenerator != nil {
		return privateDnsZoneConfig_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS(generators)
	privateDnsZoneConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS(generators)
	privateDnsZoneConfig_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_STATUS{}), generators)

	return privateDnsZoneConfig_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_STATUS(gens map[string]gopter.Gen) {
	gens["RecordSets"] = gen.SliceOf(RecordSet_STATUSGenerator())
}

func Test_RecordSet_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RecordSet_STATUS to RecordSet_STATUS via AssignProperties_To_RecordSet_STATUS & AssignProperties_From_RecordSet_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRecordSet_STATUS, RecordSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRecordSet_STATUS tests if a specific instance of RecordSet_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRecordSet_STATUS(subject RecordSet_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220701s.RecordSet_STATUS
	err := copied.AssignProperties_To_RecordSet_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RecordSet_STATUS
	err = actual.AssignProperties_From_RecordSet_STATUS(&other)
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

func Test_RecordSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RecordSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRecordSet_STATUS, RecordSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRecordSet_STATUS runs a test to see if a specific instance of RecordSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRecordSet_STATUS(subject RecordSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RecordSet_STATUS
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

// Generator of RecordSet_STATUS instances for property testing - lazily instantiated by RecordSet_STATUSGenerator()
var recordSet_STATUSGenerator gopter.Gen

// RecordSet_STATUSGenerator returns a generator of RecordSet_STATUS instances for property testing.
func RecordSet_STATUSGenerator() gopter.Gen {
	if recordSet_STATUSGenerator != nil {
		return recordSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSet_STATUS(generators)
	recordSet_STATUSGenerator = gen.Struct(reflect.TypeOf(RecordSet_STATUS{}), generators)

	return recordSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRecordSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRecordSet_STATUS(gens map[string]gopter.Gen) {
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["IpAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointProvisioningState_STATUS_Deleting,
		PrivateEndpointProvisioningState_STATUS_Failed,
		PrivateEndpointProvisioningState_STATUS_Succeeded,
		PrivateEndpointProvisioningState_STATUS_Updating))
	gens["RecordSetName"] = gen.PtrOf(gen.AlphaString())
	gens["RecordType"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}
