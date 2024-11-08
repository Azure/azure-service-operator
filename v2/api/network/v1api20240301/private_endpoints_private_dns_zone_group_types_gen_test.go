// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240301

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/storage"
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
	var other storage.PrivateDnsZoneConfig
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
	var other storage.PrivateDnsZoneConfig_STATUS
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
	var hub storage.PrivateEndpointsPrivateDnsZoneGroup
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
	var other storage.PrivateEndpointsPrivateDnsZoneGroup
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
	gens["Spec"] = PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator()
	gens["Status"] = PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator()
}

func Test_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointsPrivateDnsZoneGroupOperatorSpec to PrivateEndpointsPrivateDnsZoneGroupOperatorSpec via AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec & AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec, PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec tests if a specific instance of PrivateEndpointsPrivateDnsZoneGroupOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec(subject PrivateEndpointsPrivateDnsZoneGroupOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateEndpointsPrivateDnsZoneGroupOperatorSpec
	err := copied.AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointsPrivateDnsZoneGroupOperatorSpec
	err = actual.AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec(&other)
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

func Test_PrivateEndpointsPrivateDnsZoneGroupOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointsPrivateDnsZoneGroupOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec, PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec runs a test to see if a specific instance of PrivateEndpointsPrivateDnsZoneGroupOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroupOperatorSpec(subject PrivateEndpointsPrivateDnsZoneGroupOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointsPrivateDnsZoneGroupOperatorSpec
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

// Generator of PrivateEndpointsPrivateDnsZoneGroupOperatorSpec instances for property testing - lazily instantiated by
// PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator()
var privateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator gopter.Gen

// PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator returns a generator of PrivateEndpointsPrivateDnsZoneGroupOperatorSpec instances for property testing.
func PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator() gopter.Gen {
	if privateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator != nil {
		return privateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	privateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroupOperatorSpec{}), generators)

	return privateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator
}

func Test_PrivateEndpointsPrivateDnsZoneGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointsPrivateDnsZoneGroup_STATUS to PrivateEndpointsPrivateDnsZoneGroup_STATUS via AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup_STATUS & AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS, PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS tests if a specific instance of PrivateEndpointsPrivateDnsZoneGroup_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS(subject PrivateEndpointsPrivateDnsZoneGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateEndpointsPrivateDnsZoneGroup_STATUS
	err := copied.AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointsPrivateDnsZoneGroup_STATUS
	err = actual.AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup_STATUS(&other)
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

func Test_PrivateEndpointsPrivateDnsZoneGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointsPrivateDnsZoneGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS, PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS runs a test to see if a specific instance of PrivateEndpointsPrivateDnsZoneGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_STATUS(subject PrivateEndpointsPrivateDnsZoneGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointsPrivateDnsZoneGroup_STATUS
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

// Generator of PrivateEndpointsPrivateDnsZoneGroup_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator()
var privateEndpointsPrivateDnsZoneGroup_STATUSGenerator gopter.Gen

// PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator returns a generator of PrivateEndpointsPrivateDnsZoneGroup_STATUS instances for property testing.
// We first initialize privateEndpointsPrivateDnsZoneGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointsPrivateDnsZoneGroup_STATUSGenerator() gopter.Gen {
	if privateEndpointsPrivateDnsZoneGroup_STATUSGenerator != nil {
		return privateEndpointsPrivateDnsZoneGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS(generators)
	privateEndpointsPrivateDnsZoneGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS(generators)
	privateEndpointsPrivateDnsZoneGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_STATUS{}), generators)

	return privateEndpointsPrivateDnsZoneGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfig_STATUSGenerator())
}

func Test_PrivateEndpointsPrivateDnsZoneGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointsPrivateDnsZoneGroup_Spec to PrivateEndpointsPrivateDnsZoneGroup_Spec via AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup_Spec & AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_Spec, PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_Spec tests if a specific instance of PrivateEndpointsPrivateDnsZoneGroup_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointsPrivateDnsZoneGroup_Spec(subject PrivateEndpointsPrivateDnsZoneGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateEndpointsPrivateDnsZoneGroup_Spec
	err := copied.AssignProperties_To_PrivateEndpointsPrivateDnsZoneGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointsPrivateDnsZoneGroup_Spec
	err = actual.AssignProperties_From_PrivateEndpointsPrivateDnsZoneGroup_Spec(&other)
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

func Test_PrivateEndpointsPrivateDnsZoneGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointsPrivateDnsZoneGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec, PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec runs a test to see if a specific instance of PrivateEndpointsPrivateDnsZoneGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec(subject PrivateEndpointsPrivateDnsZoneGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointsPrivateDnsZoneGroup_Spec
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

// Generator of PrivateEndpointsPrivateDnsZoneGroup_Spec instances for property testing - lazily instantiated by
// PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator()
var privateEndpointsPrivateDnsZoneGroup_SpecGenerator gopter.Gen

// PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator returns a generator of PrivateEndpointsPrivateDnsZoneGroup_Spec instances for property testing.
// We first initialize privateEndpointsPrivateDnsZoneGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointsPrivateDnsZoneGroup_SpecGenerator() gopter.Gen {
	if privateEndpointsPrivateDnsZoneGroup_SpecGenerator != nil {
		return privateEndpointsPrivateDnsZoneGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(generators)
	privateEndpointsPrivateDnsZoneGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(generators)
	privateEndpointsPrivateDnsZoneGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_Spec{}), generators)

	return privateEndpointsPrivateDnsZoneGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(PrivateEndpointsPrivateDnsZoneGroupOperatorSpecGenerator())
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfigGenerator())
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
	var other storage.RecordSet_STATUS
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
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["RecordSetName"] = gen.PtrOf(gen.AlphaString())
	gens["RecordType"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}
