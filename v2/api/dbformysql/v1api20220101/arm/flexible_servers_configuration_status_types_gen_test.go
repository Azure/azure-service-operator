// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_ConfigurationProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationProperties_STATUS, ConfigurationProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationProperties_STATUS runs a test to see if a specific instance of ConfigurationProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationProperties_STATUS(subject ConfigurationProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationProperties_STATUS
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

// Generator of ConfigurationProperties_STATUS instances for property testing - lazily instantiated by
// ConfigurationProperties_STATUSGenerator()
var configurationProperties_STATUSGenerator gopter.Gen

// ConfigurationProperties_STATUSGenerator returns a generator of ConfigurationProperties_STATUS instances for property testing.
func ConfigurationProperties_STATUSGenerator() gopter.Gen {
	if configurationProperties_STATUSGenerator != nil {
		return configurationProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS(generators)
	configurationProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ConfigurationProperties_STATUS{}), generators)

	return configurationProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["CurrentValue"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DocumentationLink"] = gen.PtrOf(gen.AlphaString())
	gens["IsConfigPendingRestart"] = gen.PtrOf(gen.OneConstOf(ConfigurationProperties_IsConfigPendingRestart_STATUS_False, ConfigurationProperties_IsConfigPendingRestart_STATUS_True))
	gens["IsDynamicConfig"] = gen.PtrOf(gen.OneConstOf(ConfigurationProperties_IsDynamicConfig_STATUS_False, ConfigurationProperties_IsDynamicConfig_STATUS_True))
	gens["IsReadOnly"] = gen.PtrOf(gen.OneConstOf(ConfigurationProperties_IsReadOnly_STATUS_False, ConfigurationProperties_IsReadOnly_STATUS_True))
	gens["Source"] = gen.PtrOf(gen.OneConstOf(ConfigurationProperties_Source_STATUS_SystemDefault, ConfigurationProperties_Source_STATUS_UserOverride))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServersConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration_STATUS, FlexibleServersConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration_STATUS runs a test to see if a specific instance of FlexibleServersConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration_STATUS(subject FlexibleServersConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration_STATUS
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

// Generator of FlexibleServersConfiguration_STATUS instances for property testing - lazily instantiated by
// FlexibleServersConfiguration_STATUSGenerator()
var flexibleServersConfiguration_STATUSGenerator gopter.Gen

// FlexibleServersConfiguration_STATUSGenerator returns a generator of FlexibleServersConfiguration_STATUS instances for property testing.
// We first initialize flexibleServersConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersConfiguration_STATUSGenerator() gopter.Gen {
	if flexibleServersConfiguration_STATUSGenerator != nil {
		return flexibleServersConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	flexibleServersConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	flexibleServersConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS{}), generators)

	return flexibleServersConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConfigurationProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
