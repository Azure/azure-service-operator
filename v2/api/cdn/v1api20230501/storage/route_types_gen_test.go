// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_ActivatedResourceReference_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ActivatedResourceReference via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActivatedResourceReference, ActivatedResourceReferenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActivatedResourceReference runs a test to see if a specific instance of ActivatedResourceReference round trips to JSON and back losslessly
func RunJSONSerializationTestForActivatedResourceReference(subject ActivatedResourceReference) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ActivatedResourceReference
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

// Generator of ActivatedResourceReference instances for property testing - lazily instantiated by
// ActivatedResourceReferenceGenerator()
var activatedResourceReferenceGenerator gopter.Gen

// ActivatedResourceReferenceGenerator returns a generator of ActivatedResourceReference instances for property testing.
func ActivatedResourceReferenceGenerator() gopter.Gen {
	if activatedResourceReferenceGenerator != nil {
		return activatedResourceReferenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	activatedResourceReferenceGenerator = gen.Struct(reflect.TypeOf(ActivatedResourceReference{}), generators)

	return activatedResourceReferenceGenerator
}

func Test_ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded, ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded runs a test to see if a specific instance of ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded(subject ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded
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

// Generator of ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded instances for property
// testing - lazily instantiated by
// ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator()
var activatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator gopter.Gen

// ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator returns a generator of ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded instances for property testing.
func ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator() gopter.Gen {
	if activatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator != nil {
		return activatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded(generators)
	activatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded{}), generators)

	return activatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_AfdRouteCacheConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdRouteCacheConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdRouteCacheConfiguration, AfdRouteCacheConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdRouteCacheConfiguration runs a test to see if a specific instance of AfdRouteCacheConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdRouteCacheConfiguration(subject AfdRouteCacheConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdRouteCacheConfiguration
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

// Generator of AfdRouteCacheConfiguration instances for property testing - lazily instantiated by
// AfdRouteCacheConfigurationGenerator()
var afdRouteCacheConfigurationGenerator gopter.Gen

// AfdRouteCacheConfigurationGenerator returns a generator of AfdRouteCacheConfiguration instances for property testing.
// We first initialize afdRouteCacheConfigurationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdRouteCacheConfigurationGenerator() gopter.Gen {
	if afdRouteCacheConfigurationGenerator != nil {
		return afdRouteCacheConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration(generators)
	afdRouteCacheConfigurationGenerator = gen.Struct(reflect.TypeOf(AfdRouteCacheConfiguration{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration(generators)
	AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration(generators)
	afdRouteCacheConfigurationGenerator = gen.Struct(reflect.TypeOf(AfdRouteCacheConfiguration{}), generators)

	return afdRouteCacheConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration(gens map[string]gopter.Gen) {
	gens["QueryParameters"] = gen.PtrOf(gen.AlphaString())
	gens["QueryStringCachingBehavior"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration(gens map[string]gopter.Gen) {
	gens["CompressionSettings"] = gen.PtrOf(CompressionSettingsGenerator())
}

func Test_AfdRouteCacheConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdRouteCacheConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdRouteCacheConfiguration_STATUS, AfdRouteCacheConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdRouteCacheConfiguration_STATUS runs a test to see if a specific instance of AfdRouteCacheConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdRouteCacheConfiguration_STATUS(subject AfdRouteCacheConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdRouteCacheConfiguration_STATUS
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

// Generator of AfdRouteCacheConfiguration_STATUS instances for property testing - lazily instantiated by
// AfdRouteCacheConfiguration_STATUSGenerator()
var afdRouteCacheConfiguration_STATUSGenerator gopter.Gen

// AfdRouteCacheConfiguration_STATUSGenerator returns a generator of AfdRouteCacheConfiguration_STATUS instances for property testing.
// We first initialize afdRouteCacheConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AfdRouteCacheConfiguration_STATUSGenerator() gopter.Gen {
	if afdRouteCacheConfiguration_STATUSGenerator != nil {
		return afdRouteCacheConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS(generators)
	afdRouteCacheConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdRouteCacheConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS(generators)
	afdRouteCacheConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(AfdRouteCacheConfiguration_STATUS{}), generators)

	return afdRouteCacheConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["QueryParameters"] = gen.PtrOf(gen.AlphaString())
	gens["QueryStringCachingBehavior"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdRouteCacheConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["CompressionSettings"] = gen.PtrOf(CompressionSettings_STATUSGenerator())
}

func Test_CompressionSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompressionSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompressionSettings, CompressionSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompressionSettings runs a test to see if a specific instance of CompressionSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForCompressionSettings(subject CompressionSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompressionSettings
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

// Generator of CompressionSettings instances for property testing - lazily instantiated by
// CompressionSettingsGenerator()
var compressionSettingsGenerator gopter.Gen

// CompressionSettingsGenerator returns a generator of CompressionSettings instances for property testing.
func CompressionSettingsGenerator() gopter.Gen {
	if compressionSettingsGenerator != nil {
		return compressionSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompressionSettings(generators)
	compressionSettingsGenerator = gen.Struct(reflect.TypeOf(CompressionSettings{}), generators)

	return compressionSettingsGenerator
}

// AddIndependentPropertyGeneratorsForCompressionSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompressionSettings(gens map[string]gopter.Gen) {
	gens["ContentTypesToCompress"] = gen.SliceOf(gen.AlphaString())
	gens["IsCompressionEnabled"] = gen.PtrOf(gen.Bool())
}

func Test_CompressionSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompressionSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompressionSettings_STATUS, CompressionSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompressionSettings_STATUS runs a test to see if a specific instance of CompressionSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCompressionSettings_STATUS(subject CompressionSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompressionSettings_STATUS
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

// Generator of CompressionSettings_STATUS instances for property testing - lazily instantiated by
// CompressionSettings_STATUSGenerator()
var compressionSettings_STATUSGenerator gopter.Gen

// CompressionSettings_STATUSGenerator returns a generator of CompressionSettings_STATUS instances for property testing.
func CompressionSettings_STATUSGenerator() gopter.Gen {
	if compressionSettings_STATUSGenerator != nil {
		return compressionSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompressionSettings_STATUS(generators)
	compressionSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(CompressionSettings_STATUS{}), generators)

	return compressionSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCompressionSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompressionSettings_STATUS(gens map[string]gopter.Gen) {
	gens["ContentTypesToCompress"] = gen.SliceOf(gen.AlphaString())
	gens["IsCompressionEnabled"] = gen.PtrOf(gen.Bool())
}

func Test_Profiles_AfdEndpoints_Route_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_AfdEndpoints_Route_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_AfdEndpoints_Route_STATUS, Profiles_AfdEndpoints_Route_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_AfdEndpoints_Route_STATUS runs a test to see if a specific instance of Profiles_AfdEndpoints_Route_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_AfdEndpoints_Route_STATUS(subject Profiles_AfdEndpoints_Route_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_AfdEndpoints_Route_STATUS
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

// Generator of Profiles_AfdEndpoints_Route_STATUS instances for property testing - lazily instantiated by
// Profiles_AfdEndpoints_Route_STATUSGenerator()
var profiles_AfdEndpoints_Route_STATUSGenerator gopter.Gen

// Profiles_AfdEndpoints_Route_STATUSGenerator returns a generator of Profiles_AfdEndpoints_Route_STATUS instances for property testing.
// We first initialize profiles_AfdEndpoints_Route_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_AfdEndpoints_Route_STATUSGenerator() gopter.Gen {
	if profiles_AfdEndpoints_Route_STATUSGenerator != nil {
		return profiles_AfdEndpoints_Route_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS(generators)
	profiles_AfdEndpoints_Route_STATUSGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoints_Route_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS(generators)
	AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS(generators)
	profiles_AfdEndpoints_Route_STATUSGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoints_Route_STATUS{}), generators)

	return profiles_AfdEndpoints_Route_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EnabledState"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointName"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardingProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["HttpsRedirect"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LinkToDefaultDomain"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["OriginPath"] = gen.PtrOf(gen.AlphaString())
	gens["PatternsToMatch"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["SupportedProtocols"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_STATUS(gens map[string]gopter.Gen) {
	gens["CacheConfiguration"] = gen.PtrOf(AfdRouteCacheConfiguration_STATUSGenerator())
	gens["CustomDomains"] = gen.SliceOf(ActivatedResourceReference_STATUS_Profiles_AfdEndpoints_Route_SubResourceEmbeddedGenerator())
	gens["OriginGroup"] = gen.PtrOf(ResourceReference_STATUSGenerator())
	gens["RuleSets"] = gen.SliceOf(ResourceReference_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Profiles_AfdEndpoints_Route_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_AfdEndpoints_Route_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_AfdEndpoints_Route_Spec, Profiles_AfdEndpoints_Route_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_AfdEndpoints_Route_Spec runs a test to see if a specific instance of Profiles_AfdEndpoints_Route_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_AfdEndpoints_Route_Spec(subject Profiles_AfdEndpoints_Route_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_AfdEndpoints_Route_Spec
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

// Generator of Profiles_AfdEndpoints_Route_Spec instances for property testing - lazily instantiated by
// Profiles_AfdEndpoints_Route_SpecGenerator()
var profiles_AfdEndpoints_Route_SpecGenerator gopter.Gen

// Profiles_AfdEndpoints_Route_SpecGenerator returns a generator of Profiles_AfdEndpoints_Route_Spec instances for property testing.
// We first initialize profiles_AfdEndpoints_Route_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_AfdEndpoints_Route_SpecGenerator() gopter.Gen {
	if profiles_AfdEndpoints_Route_SpecGenerator != nil {
		return profiles_AfdEndpoints_Route_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec(generators)
	profiles_AfdEndpoints_Route_SpecGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoints_Route_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec(generators)
	AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec(generators)
	profiles_AfdEndpoints_Route_SpecGenerator = gen.Struct(reflect.TypeOf(Profiles_AfdEndpoints_Route_Spec{}), generators)

	return profiles_AfdEndpoints_Route_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnabledState"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardingProtocol"] = gen.PtrOf(gen.AlphaString())
	gens["HttpsRedirect"] = gen.PtrOf(gen.AlphaString())
	gens["LinkToDefaultDomain"] = gen.PtrOf(gen.AlphaString())
	gens["OriginPath"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PatternsToMatch"] = gen.SliceOf(gen.AlphaString())
	gens["SupportedProtocols"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_AfdEndpoints_Route_Spec(gens map[string]gopter.Gen) {
	gens["CacheConfiguration"] = gen.PtrOf(AfdRouteCacheConfigurationGenerator())
	gens["CustomDomains"] = gen.SliceOf(ActivatedResourceReferenceGenerator())
	gens["OriginGroup"] = gen.PtrOf(ResourceReferenceGenerator())
	gens["RuleSets"] = gen.SliceOf(ResourceReferenceGenerator())
}

func Test_Route_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Route via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoute, RouteGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoute runs a test to see if a specific instance of Route round trips to JSON and back losslessly
func RunJSONSerializationTestForRoute(subject Route) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Route
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

// Generator of Route instances for property testing - lazily instantiated by RouteGenerator()
var routeGenerator gopter.Gen

// RouteGenerator returns a generator of Route instances for property testing.
func RouteGenerator() gopter.Gen {
	if routeGenerator != nil {
		return routeGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoute(generators)
	routeGenerator = gen.Struct(reflect.TypeOf(Route{}), generators)

	return routeGenerator
}

// AddRelatedPropertyGeneratorsForRoute is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoute(gens map[string]gopter.Gen) {
	gens["Spec"] = Profiles_AfdEndpoints_Route_SpecGenerator()
	gens["Status"] = Profiles_AfdEndpoints_Route_STATUSGenerator()
}
