// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"encoding/json"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601storage"
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

func Test_Configuration_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration to hub returns original",
		prop.ForAll(RunResourceConversionTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForConfiguration tests if a specific instance of Configuration round trips to the hub storage version and back losslessly
func RunResourceConversionTestForConfiguration(subject Configuration) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20180601s.Configuration
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Configuration
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

func Test_Configuration_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration to Configuration via AssignProperties_To_Configuration & AssignProperties_From_Configuration returns original",
		prop.ForAll(RunPropertyAssignmentTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConfiguration tests if a specific instance of Configuration can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForConfiguration(subject Configuration) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Configuration
	err := copied.AssignProperties_To_Configuration(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Configuration
	err = actual.AssignProperties_From_Configuration(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Configuration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfiguration runs a test to see if a specific instance of Configuration round trips to JSON and back losslessly
func RunJSONSerializationTestForConfiguration(subject Configuration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration
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

// Generator of Configuration instances for property testing - lazily instantiated by ConfigurationGenerator()
var configurationGenerator gopter.Gen

// ConfigurationGenerator returns a generator of Configuration instances for property testing.
func ConfigurationGenerator() gopter.Gen {
	if configurationGenerator != nil {
		return configurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForConfiguration(generators)
	configurationGenerator = gen.Struct(reflect.TypeOf(Configuration{}), generators)

	return configurationGenerator
}

// AddRelatedPropertyGeneratorsForConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Configurations_SpecGenerator()
	gens["Status"] = Configuration_STATUSGenerator()
}

func Test_Configuration_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Configuration_STATUS to Configuration_STATUS via AssignProperties_To_Configuration_STATUS & AssignProperties_From_Configuration_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForConfiguration_STATUS, Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConfiguration_STATUS tests if a specific instance of Configuration_STATUS can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForConfiguration_STATUS(subject Configuration_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Configuration_STATUS
	err := copied.AssignProperties_To_Configuration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Configuration_STATUS
	err = actual.AssignProperties_From_Configuration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Configuration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfiguration_STATUS, Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfiguration_STATUS runs a test to see if a specific instance of Configuration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConfiguration_STATUS(subject Configuration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_STATUS
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

// Generator of Configuration_STATUS instances for property testing - lazily instantiated by
// Configuration_STATUSGenerator()
var configuration_STATUSGenerator gopter.Gen

// Configuration_STATUSGenerator returns a generator of Configuration_STATUS instances for property testing.
func Configuration_STATUSGenerator() gopter.Gen {
	if configuration_STATUSGenerator != nil {
		return configuration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfiguration_STATUS(generators)
	configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUS{}), generators)

	return configuration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Configurations_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Configurations_Spec to Servers_Configurations_Spec via AssignProperties_To_Servers_Configurations_Spec & AssignProperties_From_Servers_Configurations_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Configurations_Spec, Servers_Configurations_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Configurations_Spec tests if a specific instance of Servers_Configurations_Spec can be assigned to v1beta20180601storage and back losslessly
func RunPropertyAssignmentTestForServers_Configurations_Spec(subject Servers_Configurations_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180601s.Servers_Configurations_Spec
	err := copied.AssignProperties_To_Servers_Configurations_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Configurations_Spec
	err = actual.AssignProperties_From_Servers_Configurations_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Configurations_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Configurations_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Configurations_Spec, Servers_Configurations_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Configurations_Spec runs a test to see if a specific instance of Servers_Configurations_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Configurations_Spec(subject Servers_Configurations_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Configurations_Spec
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

// Generator of Servers_Configurations_Spec instances for property testing - lazily instantiated by
// Servers_Configurations_SpecGenerator()
var servers_Configurations_SpecGenerator gopter.Gen

// Servers_Configurations_SpecGenerator returns a generator of Servers_Configurations_Spec instances for property testing.
func Servers_Configurations_SpecGenerator() gopter.Gen {
	if servers_Configurations_SpecGenerator != nil {
		return servers_Configurations_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Configurations_Spec(generators)
	servers_Configurations_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Configurations_Spec{}), generators)

	return servers_Configurations_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Configurations_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Configurations_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
