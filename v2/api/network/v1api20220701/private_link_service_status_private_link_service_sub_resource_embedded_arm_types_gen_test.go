// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM, PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM runs a test to see if a specific instance of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(subject PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM
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

// Generator of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing -
// lazily instantiated by PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()
var privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator gopter.Gen

// PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator returns a generator of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing.
// We first initialize privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator != nil {
		return privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM{}), generators)

	return privateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(PrivateLinkServiceProperties_STATUS_ARMGenerator())
}

func Test_PrivateLinkServiceProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS_ARM, PrivateLinkServiceProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS_ARM runs a test to see if a specific instance of PrivateLinkServiceProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS_ARM(subject PrivateLinkServiceProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceProperties_STATUS_ARM
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

// Generator of PrivateLinkServiceProperties_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateLinkServiceProperties_STATUS_ARMGenerator()
var privateLinkServiceProperties_STATUS_ARMGenerator gopter.Gen

// PrivateLinkServiceProperties_STATUS_ARMGenerator returns a generator of PrivateLinkServiceProperties_STATUS_ARM instances for property testing.
// We first initialize privateLinkServiceProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceProperties_STATUS_ARMGenerator() gopter.Gen {
	if privateLinkServiceProperties_STATUS_ARMGenerator != nil {
		return privateLinkServiceProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM(generators)
	privateLinkServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM(generators)
	privateLinkServiceProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceProperties_STATUS_ARM{}), generators)

	return privateLinkServiceProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Alias"] = gen.PtrOf(gen.AlphaString())
	gens["EnableProxyProtocol"] = gen.PtrOf(gen.Bool())
	gens["Fqdns"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ApplicationGatewayProvisioningState_STATUS_Deleting,
		ApplicationGatewayProvisioningState_STATUS_Failed,
		ApplicationGatewayProvisioningState_STATUS_Succeeded,
		ApplicationGatewayProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoApproval"] = gen.PtrOf(ResourceSet_STATUS_ARMGenerator())
	gens["IpConfigurations"] = gen.SliceOf(PrivateLinkServiceIpConfiguration_STATUS_ARMGenerator())
	gens["LoadBalancerFrontendIpConfigurations"] = gen.SliceOf(FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_ARMGenerator())
	gens["Visibility"] = gen.PtrOf(ResourceSet_STATUS_ARMGenerator())
}

func Test_FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM, FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM runs a test to see if a specific instance of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(subject FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM
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

// Generator of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing
// - lazily instantiated by FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()
var frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator gopter.Gen

// FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator returns a generator of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing.
func FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator != nil {
		return frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM{}), generators)

	return frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM, NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM runs a test to see if a specific instance of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(subject NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM
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

// Generator of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing -
// lazily instantiated by NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()
var networkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator gopter.Gen

// NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator returns a generator of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing.
func NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if networkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator != nil {
		return networkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	networkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM{}), generators)

	return networkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateEndpointConnection_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM, PrivateEndpointConnection_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM(subject PrivateEndpointConnection_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_ARM
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

// Generator of PrivateEndpointConnection_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUS_ARMGenerator()
var privateEndpointConnection_STATUS_ARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_ARMGenerator returns a generator of PrivateEndpointConnection_STATUS_ARM instances for property testing.
func PrivateEndpointConnection_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_ARMGenerator != nil {
		return privateEndpointConnection_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(generators)
	privateEndpointConnection_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_ARM{}), generators)

	return privateEndpointConnection_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateLinkServiceIpConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceIpConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS_ARM, PrivateLinkServiceIpConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS_ARM runs a test to see if a specific instance of PrivateLinkServiceIpConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS_ARM(subject PrivateLinkServiceIpConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceIpConfiguration_STATUS_ARM
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

// Generator of PrivateLinkServiceIpConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateLinkServiceIpConfiguration_STATUS_ARMGenerator()
var privateLinkServiceIpConfiguration_STATUS_ARMGenerator gopter.Gen

// PrivateLinkServiceIpConfiguration_STATUS_ARMGenerator returns a generator of PrivateLinkServiceIpConfiguration_STATUS_ARM instances for property testing.
// We first initialize privateLinkServiceIpConfiguration_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceIpConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if privateLinkServiceIpConfiguration_STATUS_ARMGenerator != nil {
		return privateLinkServiceIpConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM(generators)
	privateLinkServiceIpConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfiguration_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM(generators)
	privateLinkServiceIpConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfiguration_STATUS_ARM{}), generators)

	return privateLinkServiceIpConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator())
}

func Test_ResourceSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceSet_STATUS_ARM, ResourceSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceSet_STATUS_ARM runs a test to see if a specific instance of ResourceSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceSet_STATUS_ARM(subject ResourceSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceSet_STATUS_ARM
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

// Generator of ResourceSet_STATUS_ARM instances for property testing - lazily instantiated by
// ResourceSet_STATUS_ARMGenerator()
var resourceSet_STATUS_ARMGenerator gopter.Gen

// ResourceSet_STATUS_ARMGenerator returns a generator of ResourceSet_STATUS_ARM instances for property testing.
func ResourceSet_STATUS_ARMGenerator() gopter.Gen {
	if resourceSet_STATUS_ARMGenerator != nil {
		return resourceSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceSet_STATUS_ARM(generators)
	resourceSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceSet_STATUS_ARM{}), generators)

	return resourceSet_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceSet_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Subscriptions"] = gen.SliceOf(gen.AlphaString())
}

func Test_PrivateLinkServiceIpConfigurationProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceIpConfigurationProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM, PrivateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM runs a test to see if a specific instance of PrivateLinkServiceIpConfigurationProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(subject PrivateLinkServiceIpConfigurationProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceIpConfigurationProperties_STATUS_ARM
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

// Generator of PrivateLinkServiceIpConfigurationProperties_STATUS_ARM instances for property testing - lazily
// instantiated by PrivateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator()
var privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator gopter.Gen

// PrivateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator returns a generator of PrivateLinkServiceIpConfigurationProperties_STATUS_ARM instances for property testing.
// We first initialize privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator() gopter.Gen {
	if privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator != nil {
		return privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(generators)
	privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfigurationProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(generators)
	privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfigurationProperties_STATUS_ARM{}), generators)

	return privateLinkServiceIpConfigurationProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Primary"] = gen.PtrOf(gen.Bool())
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_STATUS_IPv4, IPVersion_STATUS_IPv6))
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_STATUS_Dynamic, IPAllocationMethod_STATUS_Static))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ApplicationGatewayProvisioningState_STATUS_Deleting,
		ApplicationGatewayProvisioningState_STATUS_Failed,
		ApplicationGatewayProvisioningState_STATUS_Succeeded,
		ApplicationGatewayProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator())
}

func Test_Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM, Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM runs a test to see if a specific instance of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(subject Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM
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

// Generator of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator()
var subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator gopter.Gen

// Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator returns a generator of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM instances for property testing.
func Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator != nil {
		return subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(generators)
	subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM{}), generators)

	return subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
