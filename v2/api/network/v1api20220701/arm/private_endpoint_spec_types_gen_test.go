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

func Test_ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded, ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded(subject ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded instances for property testing - lazily
// instantiated by ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator()
var applicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded(generators)
	applicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded{}), generators)

	return applicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExtendedLocation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation, ExtendedLocationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation runs a test to see if a specific instance of ExtendedLocation round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation(subject ExtendedLocation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation
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

// Generator of ExtendedLocation instances for property testing - lazily instantiated by ExtendedLocationGenerator()
var extendedLocationGenerator gopter.Gen

// ExtendedLocationGenerator returns a generator of ExtendedLocation instances for property testing.
func ExtendedLocationGenerator() gopter.Gen {
	if extendedLocationGenerator != nil {
		return extendedLocationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation(generators)
	extendedLocationGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation{}), generators)

	return extendedLocationGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_EdgeZone))
}

func Test_PrivateEndpointIPConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointIPConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointIPConfiguration, PrivateEndpointIPConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointIPConfiguration runs a test to see if a specific instance of PrivateEndpointIPConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointIPConfiguration(subject PrivateEndpointIPConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointIPConfiguration
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

// Generator of PrivateEndpointIPConfiguration instances for property testing - lazily instantiated by
// PrivateEndpointIPConfigurationGenerator()
var privateEndpointIPConfigurationGenerator gopter.Gen

// PrivateEndpointIPConfigurationGenerator returns a generator of PrivateEndpointIPConfiguration instances for property testing.
// We first initialize privateEndpointIPConfigurationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointIPConfigurationGenerator() gopter.Gen {
	if privateEndpointIPConfigurationGenerator != nil {
		return privateEndpointIPConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointIPConfiguration(generators)
	privateEndpointIPConfigurationGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointIPConfiguration{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointIPConfiguration(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointIPConfiguration(generators)
	privateEndpointIPConfigurationGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointIPConfiguration{}), generators)

	return privateEndpointIPConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointIPConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointIPConfiguration(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpointIPConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointIPConfiguration(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateEndpointIPConfigurationPropertiesGenerator())
}

func Test_PrivateEndpointIPConfigurationProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointIPConfigurationProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointIPConfigurationProperties, PrivateEndpointIPConfigurationPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointIPConfigurationProperties runs a test to see if a specific instance of PrivateEndpointIPConfigurationProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointIPConfigurationProperties(subject PrivateEndpointIPConfigurationProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointIPConfigurationProperties
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

// Generator of PrivateEndpointIPConfigurationProperties instances for property testing - lazily instantiated by
// PrivateEndpointIPConfigurationPropertiesGenerator()
var privateEndpointIPConfigurationPropertiesGenerator gopter.Gen

// PrivateEndpointIPConfigurationPropertiesGenerator returns a generator of PrivateEndpointIPConfigurationProperties instances for property testing.
func PrivateEndpointIPConfigurationPropertiesGenerator() gopter.Gen {
	if privateEndpointIPConfigurationPropertiesGenerator != nil {
		return privateEndpointIPConfigurationPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointIPConfigurationProperties(generators)
	privateEndpointIPConfigurationPropertiesGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointIPConfigurationProperties{}), generators)

	return privateEndpointIPConfigurationPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointIPConfigurationProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointIPConfigurationProperties(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["MemberName"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateEndpointProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointProperties, PrivateEndpointPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointProperties runs a test to see if a specific instance of PrivateEndpointProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointProperties(subject PrivateEndpointProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointProperties
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

// Generator of PrivateEndpointProperties instances for property testing - lazily instantiated by
// PrivateEndpointPropertiesGenerator()
var privateEndpointPropertiesGenerator gopter.Gen

// PrivateEndpointPropertiesGenerator returns a generator of PrivateEndpointProperties instances for property testing.
// We first initialize privateEndpointPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointPropertiesGenerator() gopter.Gen {
	if privateEndpointPropertiesGenerator != nil {
		return privateEndpointPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointProperties(generators)
	privateEndpointPropertiesGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointProperties(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointProperties(generators)
	privateEndpointPropertiesGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointProperties{}), generators)

	return privateEndpointPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointProperties(gens map[string]gopter.Gen) {
	gens["CustomNetworkInterfaceName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpointProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointProperties(gens map[string]gopter.Gen) {
	gens["ApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbeddedGenerator())
	gens["IpConfigurations"] = gen.SliceOf(PrivateEndpointIPConfigurationGenerator())
	gens["ManualPrivateLinkServiceConnections"] = gen.SliceOf(PrivateLinkServiceConnectionGenerator())
	gens["PrivateLinkServiceConnections"] = gen.SliceOf(PrivateLinkServiceConnectionGenerator())
	gens["Subnet"] = gen.PtrOf(Subnet_PrivateEndpoint_SubResourceEmbeddedGenerator())
}

func Test_PrivateEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpoint_Spec, PrivateEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpoint_Spec runs a test to see if a specific instance of PrivateEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpoint_Spec(subject PrivateEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpoint_Spec
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

// Generator of PrivateEndpoint_Spec instances for property testing - lazily instantiated by
// PrivateEndpoint_SpecGenerator()
var privateEndpoint_SpecGenerator gopter.Gen

// PrivateEndpoint_SpecGenerator returns a generator of PrivateEndpoint_Spec instances for property testing.
// We first initialize privateEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpoint_SpecGenerator() gopter.Gen {
	if privateEndpoint_SpecGenerator != nil {
		return privateEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoint_Spec(generators)
	privateEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpoint_Spec(generators)
	privateEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateEndpoint_Spec{}), generators)

	return privateEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["Properties"] = gen.PtrOf(PrivateEndpointPropertiesGenerator())
}

func Test_PrivateLinkServiceConnection_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceConnection via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceConnection, PrivateLinkServiceConnectionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceConnection runs a test to see if a specific instance of PrivateLinkServiceConnection round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceConnection(subject PrivateLinkServiceConnection) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceConnection
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

// Generator of PrivateLinkServiceConnection instances for property testing - lazily instantiated by
// PrivateLinkServiceConnectionGenerator()
var privateLinkServiceConnectionGenerator gopter.Gen

// PrivateLinkServiceConnectionGenerator returns a generator of PrivateLinkServiceConnection instances for property testing.
// We first initialize privateLinkServiceConnectionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceConnectionGenerator() gopter.Gen {
	if privateLinkServiceConnectionGenerator != nil {
		return privateLinkServiceConnectionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceConnection(generators)
	privateLinkServiceConnectionGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceConnection{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceConnection(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceConnection(generators)
	privateLinkServiceConnectionGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceConnection{}), generators)

	return privateLinkServiceConnectionGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceConnection is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceConnection(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceConnection is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceConnection(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateLinkServiceConnectionPropertiesGenerator())
}

func Test_PrivateLinkServiceConnectionProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceConnectionProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceConnectionProperties, PrivateLinkServiceConnectionPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceConnectionProperties runs a test to see if a specific instance of PrivateLinkServiceConnectionProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceConnectionProperties(subject PrivateLinkServiceConnectionProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceConnectionProperties
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

// Generator of PrivateLinkServiceConnectionProperties instances for property testing - lazily instantiated by
// PrivateLinkServiceConnectionPropertiesGenerator()
var privateLinkServiceConnectionPropertiesGenerator gopter.Gen

// PrivateLinkServiceConnectionPropertiesGenerator returns a generator of PrivateLinkServiceConnectionProperties instances for property testing.
// We first initialize privateLinkServiceConnectionPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceConnectionPropertiesGenerator() gopter.Gen {
	if privateLinkServiceConnectionPropertiesGenerator != nil {
		return privateLinkServiceConnectionPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionProperties(generators)
	privateLinkServiceConnectionPropertiesGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceConnectionProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionProperties(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceConnectionProperties(generators)
	privateLinkServiceConnectionPropertiesGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceConnectionProperties{}), generators)

	return privateLinkServiceConnectionPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionProperties(gens map[string]gopter.Gen) {
	gens["GroupIds"] = gen.SliceOf(gen.AlphaString())
	gens["PrivateLinkServiceId"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceConnectionProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceConnectionProperties(gens map[string]gopter.Gen) {
	gens["PrivateLinkServiceConnectionState"] = gen.PtrOf(PrivateLinkServiceConnectionStateGenerator())
}

func Test_PrivateLinkServiceConnectionState_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceConnectionState via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceConnectionState, PrivateLinkServiceConnectionStateGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceConnectionState runs a test to see if a specific instance of PrivateLinkServiceConnectionState round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceConnectionState(subject PrivateLinkServiceConnectionState) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceConnectionState
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

// Generator of PrivateLinkServiceConnectionState instances for property testing - lazily instantiated by
// PrivateLinkServiceConnectionStateGenerator()
var privateLinkServiceConnectionStateGenerator gopter.Gen

// PrivateLinkServiceConnectionStateGenerator returns a generator of PrivateLinkServiceConnectionState instances for property testing.
func PrivateLinkServiceConnectionStateGenerator() gopter.Gen {
	if privateLinkServiceConnectionStateGenerator != nil {
		return privateLinkServiceConnectionStateGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionState(generators)
	privateLinkServiceConnectionStateGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceConnectionState{}), generators)

	return privateLinkServiceConnectionStateGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionState is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceConnectionState(gens map[string]gopter.Gen) {
	gens["ActionsRequired"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_PrivateEndpoint_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_PrivateEndpoint_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_PrivateEndpoint_SubResourceEmbedded, Subnet_PrivateEndpoint_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_PrivateEndpoint_SubResourceEmbedded runs a test to see if a specific instance of Subnet_PrivateEndpoint_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_PrivateEndpoint_SubResourceEmbedded(subject Subnet_PrivateEndpoint_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_PrivateEndpoint_SubResourceEmbedded
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

// Generator of Subnet_PrivateEndpoint_SubResourceEmbedded instances for property testing - lazily instantiated by
// Subnet_PrivateEndpoint_SubResourceEmbeddedGenerator()
var subnet_PrivateEndpoint_SubResourceEmbeddedGenerator gopter.Gen

// Subnet_PrivateEndpoint_SubResourceEmbeddedGenerator returns a generator of Subnet_PrivateEndpoint_SubResourceEmbedded instances for property testing.
func Subnet_PrivateEndpoint_SubResourceEmbeddedGenerator() gopter.Gen {
	if subnet_PrivateEndpoint_SubResourceEmbeddedGenerator != nil {
		return subnet_PrivateEndpoint_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_PrivateEndpoint_SubResourceEmbedded(generators)
	subnet_PrivateEndpoint_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(Subnet_PrivateEndpoint_SubResourceEmbedded{}), generators)

	return subnet_PrivateEndpoint_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_PrivateEndpoint_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_PrivateEndpoint_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}