// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

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

func Test_ConfigurationProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationProperties_ARM, ConfigurationProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationProperties_ARM runs a test to see if a specific instance of ConfigurationProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationProperties_ARM(subject ConfigurationProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationProperties_ARM
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

// Generator of ConfigurationProperties_ARM instances for property testing - lazily instantiated by
// ConfigurationProperties_ARMGenerator()
var configurationProperties_ARMGenerator gopter.Gen

// ConfigurationProperties_ARMGenerator returns a generator of ConfigurationProperties_ARM instances for property testing.
func ConfigurationProperties_ARMGenerator() gopter.Gen {
	if configurationProperties_ARMGenerator != nil {
		return configurationProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationProperties_ARM(generators)
	configurationProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationProperties_ARM{}), generators)

	return configurationProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationProperties_ARM(gens map[string]gopter.Gen) {
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Configuration_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Configuration_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Configuration_Spec_ARM, Servers_Configuration_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Configuration_Spec_ARM runs a test to see if a specific instance of Servers_Configuration_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Configuration_Spec_ARM(subject Servers_Configuration_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Configuration_Spec_ARM
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

// Generator of Servers_Configuration_Spec_ARM instances for property testing - lazily instantiated by
// Servers_Configuration_Spec_ARMGenerator()
var servers_Configuration_Spec_ARMGenerator gopter.Gen

// Servers_Configuration_Spec_ARMGenerator returns a generator of Servers_Configuration_Spec_ARM instances for property testing.
// We first initialize servers_Configuration_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_Configuration_Spec_ARMGenerator() gopter.Gen {
	if servers_Configuration_Spec_ARMGenerator != nil {
		return servers_Configuration_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Configuration_Spec_ARM(generators)
	servers_Configuration_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Configuration_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Configuration_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_Configuration_Spec_ARM(generators)
	servers_Configuration_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Configuration_Spec_ARM{}), generators)

	return servers_Configuration_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_Configuration_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Configuration_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServers_Configuration_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_Configuration_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConfigurationProperties_ARMGenerator())
}
