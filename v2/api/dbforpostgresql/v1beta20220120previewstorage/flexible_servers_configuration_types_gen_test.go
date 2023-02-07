// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220120previewstorage

import (
	"encoding/json"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
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

func Test_FlexibleServersConfiguration_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210601s.FlexibleServersConfiguration
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersConfiguration
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

func Test_FlexibleServersConfiguration_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to FlexibleServersConfiguration via AssignProperties_To_FlexibleServersConfiguration & AssignProperties_From_FlexibleServersConfiguration returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServersConfiguration
	err := copied.AssignProperties_To_FlexibleServersConfiguration(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfiguration
	err = actual.AssignProperties_From_FlexibleServersConfiguration(&other)
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

func Test_FlexibleServersConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration runs a test to see if a specific instance of FlexibleServersConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration
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

// Generator of FlexibleServersConfiguration instances for property testing - lazily instantiated by
// FlexibleServersConfigurationGenerator()
var flexibleServersConfigurationGenerator gopter.Gen

// FlexibleServersConfigurationGenerator returns a generator of FlexibleServersConfiguration instances for property testing.
func FlexibleServersConfigurationGenerator() gopter.Gen {
	if flexibleServersConfigurationGenerator != nil {
		return flexibleServersConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(generators)
	flexibleServersConfigurationGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration{}), generators)

	return flexibleServersConfigurationGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServers_Configuration_SpecGenerator()
	gens["Status"] = FlexibleServers_Configuration_STATUSGenerator()
}

func Test_FlexibleServers_Configuration_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Configuration_Spec to FlexibleServers_Configuration_Spec via AssignProperties_To_FlexibleServers_Configuration_Spec & AssignProperties_From_FlexibleServers_Configuration_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Configuration_Spec, FlexibleServers_Configuration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Configuration_Spec tests if a specific instance of FlexibleServers_Configuration_Spec can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Configuration_Spec(subject FlexibleServers_Configuration_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServers_Configuration_Spec
	err := copied.AssignProperties_To_FlexibleServers_Configuration_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Configuration_Spec
	err = actual.AssignProperties_From_FlexibleServers_Configuration_Spec(&other)
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

func Test_FlexibleServers_Configuration_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Configuration_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Configuration_Spec, FlexibleServers_Configuration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Configuration_Spec runs a test to see if a specific instance of FlexibleServers_Configuration_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Configuration_Spec(subject FlexibleServers_Configuration_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Configuration_Spec
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

// Generator of FlexibleServers_Configuration_Spec instances for property testing - lazily instantiated by
// FlexibleServers_Configuration_SpecGenerator()
var flexibleServers_Configuration_SpecGenerator gopter.Gen

// FlexibleServers_Configuration_SpecGenerator returns a generator of FlexibleServers_Configuration_Spec instances for property testing.
func FlexibleServers_Configuration_SpecGenerator() gopter.Gen {
	if flexibleServers_Configuration_SpecGenerator != nil {
		return flexibleServers_Configuration_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_Spec(generators)
	flexibleServers_Configuration_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Configuration_Spec{}), generators)

	return flexibleServers_Configuration_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServers_Configuration_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Configuration_STATUS to FlexibleServers_Configuration_STATUS via AssignProperties_To_FlexibleServers_Configuration_STATUS & AssignProperties_From_FlexibleServers_Configuration_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Configuration_STATUS, FlexibleServers_Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Configuration_STATUS tests if a specific instance of FlexibleServers_Configuration_STATUS can be assigned to v1beta20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Configuration_STATUS(subject FlexibleServers_Configuration_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210601s.FlexibleServers_Configuration_STATUS
	err := copied.AssignProperties_To_FlexibleServers_Configuration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Configuration_STATUS
	err = actual.AssignProperties_From_FlexibleServers_Configuration_STATUS(&other)
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

func Test_FlexibleServers_Configuration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Configuration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Configuration_STATUS, FlexibleServers_Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Configuration_STATUS runs a test to see if a specific instance of FlexibleServers_Configuration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Configuration_STATUS(subject FlexibleServers_Configuration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Configuration_STATUS
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

// Generator of FlexibleServers_Configuration_STATUS instances for property testing - lazily instantiated by
// FlexibleServers_Configuration_STATUSGenerator()
var flexibleServers_Configuration_STATUSGenerator gopter.Gen

// FlexibleServers_Configuration_STATUSGenerator returns a generator of FlexibleServers_Configuration_STATUS instances for property testing.
// We first initialize flexibleServers_Configuration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Configuration_STATUSGenerator() gopter.Gen {
	if flexibleServers_Configuration_STATUSGenerator != nil {
		return flexibleServers_Configuration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_STATUS(generators)
	flexibleServers_Configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Configuration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Configuration_STATUS(generators)
	flexibleServers_Configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Configuration_STATUS{}), generators)

	return flexibleServers_Configuration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Configuration_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DocumentationLink"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsConfigPendingRestart"] = gen.PtrOf(gen.Bool())
	gens["IsDynamicConfig"] = gen.PtrOf(gen.Bool())
	gens["IsReadOnly"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Unit"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_Configuration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Configuration_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
