// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120preview

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

func Test_ConfigurationProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationProperties_STATUS_ARM, ConfigurationProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationProperties_STATUS_ARM runs a test to see if a specific instance of ConfigurationProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationProperties_STATUS_ARM(subject ConfigurationProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationProperties_STATUS_ARM
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

// Generator of ConfigurationProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ConfigurationProperties_STATUS_ARMGenerator()
var configurationProperties_STATUS_ARMGenerator gopter.Gen

// ConfigurationProperties_STATUS_ARMGenerator returns a generator of ConfigurationProperties_STATUS_ARM instances for property testing.
func ConfigurationProperties_STATUS_ARMGenerator() gopter.Gen {
	if configurationProperties_STATUS_ARMGenerator != nil {
		return configurationProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS_ARM(generators)
	configurationProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationProperties_STATUS_ARM{}), generators)

	return configurationProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		ConfigurationProperties_DataType_STATUS_ARM_Boolean,
		ConfigurationProperties_DataType_STATUS_ARM_Enumeration,
		ConfigurationProperties_DataType_STATUS_ARM_Integer,
		ConfigurationProperties_DataType_STATUS_ARM_Numeric))
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DocumentationLink"] = gen.PtrOf(gen.AlphaString())
	gens["IsConfigPendingRestart"] = gen.PtrOf(gen.Bool())
	gens["IsDynamicConfig"] = gen.PtrOf(gen.Bool())
	gens["IsReadOnly"] = gen.PtrOf(gen.Bool())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Unit"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServersConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration_STATUS_ARM, FlexibleServersConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration_STATUS_ARM runs a test to see if a specific instance of FlexibleServersConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration_STATUS_ARM(subject FlexibleServersConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration_STATUS_ARM
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

// Generator of FlexibleServersConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// FlexibleServersConfiguration_STATUS_ARMGenerator()
var flexibleServersConfiguration_STATUS_ARMGenerator gopter.Gen

// FlexibleServersConfiguration_STATUS_ARMGenerator returns a generator of FlexibleServersConfiguration_STATUS_ARM instances for property testing.
// We first initialize flexibleServersConfiguration_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if flexibleServersConfiguration_STATUS_ARMGenerator != nil {
		return flexibleServersConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM(generators)
	flexibleServersConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM(generators)
	flexibleServersConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS_ARM{}), generators)

	return flexibleServersConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConfigurationProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}
