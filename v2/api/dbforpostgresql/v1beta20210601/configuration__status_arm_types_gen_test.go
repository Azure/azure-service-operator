// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_Configuration_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationStatusARM, ConfigurationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationStatusARM runs a test to see if a specific instance of Configuration_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationStatusARM(subject Configuration_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_StatusARM
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

// Generator of Configuration_StatusARM instances for property testing - lazily instantiated by
// ConfigurationStatusARMGenerator()
var configurationStatusARMGenerator gopter.Gen

// ConfigurationStatusARMGenerator returns a generator of Configuration_StatusARM instances for property testing.
// We first initialize configurationStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationStatusARMGenerator() gopter.Gen {
	if configurationStatusARMGenerator != nil {
		return configurationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatusARM(generators)
	configurationStatusARMGenerator = gen.Struct(reflect.TypeOf(Configuration_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationStatusARM(generators)
	AddRelatedPropertyGeneratorsForConfigurationStatusARM(generators)
	configurationStatusARMGenerator = gen.Struct(reflect.TypeOf(Configuration_StatusARM{}), generators)

	return configurationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfigurationStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConfigurationPropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_ConfigurationProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationPropertiesStatusARM, ConfigurationPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationPropertiesStatusARM runs a test to see if a specific instance of ConfigurationProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationPropertiesStatusARM(subject ConfigurationProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationProperties_StatusARM
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

// Generator of ConfigurationProperties_StatusARM instances for property testing - lazily instantiated by
// ConfigurationPropertiesStatusARMGenerator()
var configurationPropertiesStatusARMGenerator gopter.Gen

// ConfigurationPropertiesStatusARMGenerator returns a generator of ConfigurationProperties_StatusARM instances for property testing.
func ConfigurationPropertiesStatusARMGenerator() gopter.Gen {
	if configurationPropertiesStatusARMGenerator != nil {
		return configurationPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationPropertiesStatusARM(generators)
	configurationPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationProperties_StatusARM{}), generators)

	return configurationPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		ConfigurationPropertiesStatusDataType_Boolean,
		ConfigurationPropertiesStatusDataType_Enumeration,
		ConfigurationPropertiesStatusDataType_Integer,
		ConfigurationPropertiesStatusDataType_Numeric))
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

func Test_SystemData_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatusARM, SystemDataStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatusARM runs a test to see if a specific instance of SystemData_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatusARM(subject SystemData_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_StatusARM
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

// Generator of SystemData_StatusARM instances for property testing - lazily instantiated by
// SystemDataStatusARMGenerator()
var systemDataStatusARMGenerator gopter.Gen

// SystemDataStatusARMGenerator returns a generator of SystemData_StatusARM instances for property testing.
func SystemDataStatusARMGenerator() gopter.Gen {
	if systemDataStatusARMGenerator != nil {
		return systemDataStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatusARM(generators)
	systemDataStatusARMGenerator = gen.Struct(reflect.TypeOf(SystemData_StatusARM{}), generators)

	return systemDataStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusCreatedByType_Application,
		SystemDataStatusCreatedByType_Key,
		SystemDataStatusCreatedByType_ManagedIdentity,
		SystemDataStatusCreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusLastModifiedByType_Application,
		SystemDataStatusLastModifiedByType_Key,
		SystemDataStatusLastModifiedByType_ManagedIdentity,
		SystemDataStatusLastModifiedByType_User))
}
