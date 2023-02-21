// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601storage

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
	gens["Spec"] = Servers_Configuration_SpecGenerator()
	gens["Status"] = Servers_Configuration_STATUSGenerator()
}

func Test_Servers_Configuration_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Configuration_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Configuration_Spec, Servers_Configuration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Configuration_Spec runs a test to see if a specific instance of Servers_Configuration_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Configuration_Spec(subject Servers_Configuration_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Configuration_Spec
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

// Generator of Servers_Configuration_Spec instances for property testing - lazily instantiated by
// Servers_Configuration_SpecGenerator()
var servers_Configuration_SpecGenerator gopter.Gen

// Servers_Configuration_SpecGenerator returns a generator of Servers_Configuration_Spec instances for property testing.
func Servers_Configuration_SpecGenerator() gopter.Gen {
	if servers_Configuration_SpecGenerator != nil {
		return servers_Configuration_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Configuration_Spec(generators)
	servers_Configuration_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Configuration_Spec{}), generators)

	return servers_Configuration_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Configuration_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Configuration_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Configuration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Configuration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Configuration_STATUS, Servers_Configuration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Configuration_STATUS runs a test to see if a specific instance of Servers_Configuration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Configuration_STATUS(subject Servers_Configuration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Configuration_STATUS
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

// Generator of Servers_Configuration_STATUS instances for property testing - lazily instantiated by
// Servers_Configuration_STATUSGenerator()
var servers_Configuration_STATUSGenerator gopter.Gen

// Servers_Configuration_STATUSGenerator returns a generator of Servers_Configuration_STATUS instances for property testing.
func Servers_Configuration_STATUSGenerator() gopter.Gen {
	if servers_Configuration_STATUSGenerator != nil {
		return servers_Configuration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Configuration_STATUS(generators)
	servers_Configuration_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Configuration_STATUS{}), generators)

	return servers_Configuration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Configuration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Configuration_STATUS(gens map[string]gopter.Gen) {
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
