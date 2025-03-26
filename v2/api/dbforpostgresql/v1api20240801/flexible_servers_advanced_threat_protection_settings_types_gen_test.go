// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240801

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801/storage"
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

func Test_FlexibleServersAdvancedThreatProtectionSettings_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdvancedThreatProtectionSettings to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersAdvancedThreatProtectionSettings, FlexibleServersAdvancedThreatProtectionSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersAdvancedThreatProtectionSettings tests if a specific instance of FlexibleServersAdvancedThreatProtectionSettings round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersAdvancedThreatProtectionSettings(subject FlexibleServersAdvancedThreatProtectionSettings) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.FlexibleServersAdvancedThreatProtectionSettings
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersAdvancedThreatProtectionSettings
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

func Test_FlexibleServersAdvancedThreatProtectionSettings_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdvancedThreatProtectionSettings to FlexibleServersAdvancedThreatProtectionSettings via AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings & AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings, FlexibleServersAdvancedThreatProtectionSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings tests if a specific instance of FlexibleServersAdvancedThreatProtectionSettings can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings(subject FlexibleServersAdvancedThreatProtectionSettings) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersAdvancedThreatProtectionSettings
	err := copied.AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersAdvancedThreatProtectionSettings
	err = actual.AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings(&other)
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

func Test_FlexibleServersAdvancedThreatProtectionSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdvancedThreatProtectionSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings, FlexibleServersAdvancedThreatProtectionSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings runs a test to see if a specific instance of FlexibleServersAdvancedThreatProtectionSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings(subject FlexibleServersAdvancedThreatProtectionSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdvancedThreatProtectionSettings
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

// Generator of FlexibleServersAdvancedThreatProtectionSettings instances for property testing - lazily instantiated by
// FlexibleServersAdvancedThreatProtectionSettingsGenerator()
var flexibleServersAdvancedThreatProtectionSettingsGenerator gopter.Gen

// FlexibleServersAdvancedThreatProtectionSettingsGenerator returns a generator of FlexibleServersAdvancedThreatProtectionSettings instances for property testing.
func FlexibleServersAdvancedThreatProtectionSettingsGenerator() gopter.Gen {
	if flexibleServersAdvancedThreatProtectionSettingsGenerator != nil {
		return flexibleServersAdvancedThreatProtectionSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings(generators)
	flexibleServersAdvancedThreatProtectionSettingsGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettings{}), generators)

	return flexibleServersAdvancedThreatProtectionSettingsGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator()
	gens["Status"] = FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator()
}

func Test_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec to FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec via AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec & AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec, FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec tests if a specific instance of FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec(subject FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec
	err := copied.AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec
	err = actual.AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec(&other)
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

func Test_FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec, FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec runs a test to see if a specific instance of FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettingsOperatorSpec(subject FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec
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

// Generator of FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec instances for property testing - lazily
// instantiated by FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator()
var flexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator gopter.Gen

// FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator returns a generator of FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec instances for property testing.
func FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator() gopter.Gen {
	if flexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator != nil {
		return flexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	flexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettingsOperatorSpec{}), generators)

	return flexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator
}

func Test_FlexibleServersAdvancedThreatProtectionSettings_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdvancedThreatProtectionSettings_STATUS to FlexibleServersAdvancedThreatProtectionSettings_STATUS via AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings_STATUS & AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS, FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS tests if a specific instance of FlexibleServersAdvancedThreatProtectionSettings_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS(subject FlexibleServersAdvancedThreatProtectionSettings_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersAdvancedThreatProtectionSettings_STATUS
	err := copied.AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersAdvancedThreatProtectionSettings_STATUS
	err = actual.AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings_STATUS(&other)
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

func Test_FlexibleServersAdvancedThreatProtectionSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdvancedThreatProtectionSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS, FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS runs a test to see if a specific instance of FlexibleServersAdvancedThreatProtectionSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_STATUS(subject FlexibleServersAdvancedThreatProtectionSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdvancedThreatProtectionSettings_STATUS
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

// Generator of FlexibleServersAdvancedThreatProtectionSettings_STATUS instances for property testing - lazily
// instantiated by FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator()
var flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator gopter.Gen

// FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator returns a generator of FlexibleServersAdvancedThreatProtectionSettings_STATUS instances for property testing.
// We first initialize flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdvancedThreatProtectionSettings_STATUSGenerator() gopter.Gen {
	if flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator != nil {
		return flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS(generators)
	flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS(generators)
	flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettings_STATUS{}), generators)

	return flexibleServersAdvancedThreatProtectionSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerThreatProtectionProperties_State_STATUS_Disabled, ServerThreatProtectionProperties_State_STATUS_Enabled))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServersAdvancedThreatProtectionSettings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersAdvancedThreatProtectionSettings_Spec to FlexibleServersAdvancedThreatProtectionSettings_Spec via AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings_Spec & AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_Spec, FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_Spec tests if a specific instance of FlexibleServersAdvancedThreatProtectionSettings_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersAdvancedThreatProtectionSettings_Spec(subject FlexibleServersAdvancedThreatProtectionSettings_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersAdvancedThreatProtectionSettings_Spec
	err := copied.AssignProperties_To_FlexibleServersAdvancedThreatProtectionSettings_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersAdvancedThreatProtectionSettings_Spec
	err = actual.AssignProperties_From_FlexibleServersAdvancedThreatProtectionSettings_Spec(&other)
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

func Test_FlexibleServersAdvancedThreatProtectionSettings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdvancedThreatProtectionSettings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_Spec, FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_Spec runs a test to see if a specific instance of FlexibleServersAdvancedThreatProtectionSettings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdvancedThreatProtectionSettings_Spec(subject FlexibleServersAdvancedThreatProtectionSettings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdvancedThreatProtectionSettings_Spec
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

// Generator of FlexibleServersAdvancedThreatProtectionSettings_Spec instances for property testing - lazily
// instantiated by FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator()
var flexibleServersAdvancedThreatProtectionSettings_SpecGenerator gopter.Gen

// FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator returns a generator of FlexibleServersAdvancedThreatProtectionSettings_Spec instances for property testing.
// We first initialize flexibleServersAdvancedThreatProtectionSettings_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdvancedThreatProtectionSettings_SpecGenerator() gopter.Gen {
	if flexibleServersAdvancedThreatProtectionSettings_SpecGenerator != nil {
		return flexibleServersAdvancedThreatProtectionSettings_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec(generators)
	flexibleServersAdvancedThreatProtectionSettings_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec(generators)
	flexibleServersAdvancedThreatProtectionSettings_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdvancedThreatProtectionSettings_Spec{}), generators)

	return flexibleServersAdvancedThreatProtectionSettings_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(ServerThreatProtectionProperties_State_Disabled, ServerThreatProtectionProperties_State_Enabled))
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdvancedThreatProtectionSettings_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(FlexibleServersAdvancedThreatProtectionSettingsOperatorSpecGenerator())
}
