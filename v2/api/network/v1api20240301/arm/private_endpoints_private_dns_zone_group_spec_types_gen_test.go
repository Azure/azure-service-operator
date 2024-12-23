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
// We first initialize privateDnsZoneConfigGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZoneConfigGenerator() gopter.Gen {
	if privateDnsZoneConfigGenerator != nil {
		return privateDnsZoneConfigGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig(generators)
	privateDnsZoneConfigGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig(generators)
	privateDnsZoneConfigGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig{}), generators)

	return privateDnsZoneConfigGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZonePropertiesFormatGenerator())
}

func Test_PrivateDnsZoneGroupPropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneGroupPropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat, PrivateDnsZoneGroupPropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat runs a test to see if a specific instance of PrivateDnsZoneGroupPropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat(subject PrivateDnsZoneGroupPropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneGroupPropertiesFormat
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

// Generator of PrivateDnsZoneGroupPropertiesFormat instances for property testing - lazily instantiated by
// PrivateDnsZoneGroupPropertiesFormatGenerator()
var privateDnsZoneGroupPropertiesFormatGenerator gopter.Gen

// PrivateDnsZoneGroupPropertiesFormatGenerator returns a generator of PrivateDnsZoneGroupPropertiesFormat instances for property testing.
func PrivateDnsZoneGroupPropertiesFormatGenerator() gopter.Gen {
	if privateDnsZoneGroupPropertiesFormatGenerator != nil {
		return privateDnsZoneGroupPropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat(generators)
	privateDnsZoneGroupPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneGroupPropertiesFormat{}), generators)

	return privateDnsZoneGroupPropertiesFormatGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfigGenerator())
}

func Test_PrivateDnsZonePropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonePropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonePropertiesFormat, PrivateDnsZonePropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonePropertiesFormat runs a test to see if a specific instance of PrivateDnsZonePropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonePropertiesFormat(subject PrivateDnsZonePropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonePropertiesFormat
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

// Generator of PrivateDnsZonePropertiesFormat instances for property testing - lazily instantiated by
// PrivateDnsZonePropertiesFormatGenerator()
var privateDnsZonePropertiesFormatGenerator gopter.Gen

// PrivateDnsZonePropertiesFormatGenerator returns a generator of PrivateDnsZonePropertiesFormat instances for property testing.
func PrivateDnsZonePropertiesFormatGenerator() gopter.Gen {
	if privateDnsZonePropertiesFormatGenerator != nil {
		return privateDnsZonePropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat(generators)
	privateDnsZonePropertiesFormatGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonePropertiesFormat{}), generators)

	return privateDnsZonePropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneId"] = gen.PtrOf(gen.AlphaString())
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
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZoneGroupPropertiesFormatGenerator())
}
