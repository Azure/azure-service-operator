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

func Test_DnsResolvers_InboundEndpoint_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolvers_InboundEndpoint_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolvers_InboundEndpoint_Spec_ARM, DnsResolvers_InboundEndpoint_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolvers_InboundEndpoint_Spec_ARM runs a test to see if a specific instance of DnsResolvers_InboundEndpoint_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolvers_InboundEndpoint_Spec_ARM(subject DnsResolvers_InboundEndpoint_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolvers_InboundEndpoint_Spec_ARM
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

// Generator of DnsResolvers_InboundEndpoint_Spec_ARM instances for property testing - lazily instantiated by
// DnsResolvers_InboundEndpoint_Spec_ARMGenerator()
var dnsResolvers_InboundEndpoint_Spec_ARMGenerator gopter.Gen

// DnsResolvers_InboundEndpoint_Spec_ARMGenerator returns a generator of DnsResolvers_InboundEndpoint_Spec_ARM instances for property testing.
// We first initialize dnsResolvers_InboundEndpoint_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolvers_InboundEndpoint_Spec_ARMGenerator() gopter.Gen {
	if dnsResolvers_InboundEndpoint_Spec_ARMGenerator != nil {
		return dnsResolvers_InboundEndpoint_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM(generators)
	dnsResolvers_InboundEndpoint_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolvers_InboundEndpoint_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM(generators)
	dnsResolvers_InboundEndpoint_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolvers_InboundEndpoint_Spec_ARM{}), generators)

	return dnsResolvers_InboundEndpoint_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundEndpointProperties_ARMGenerator())
}

func Test_InboundEndpointProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundEndpointProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundEndpointProperties_ARM, InboundEndpointProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundEndpointProperties_ARM runs a test to see if a specific instance of InboundEndpointProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundEndpointProperties_ARM(subject InboundEndpointProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundEndpointProperties_ARM
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

// Generator of InboundEndpointProperties_ARM instances for property testing - lazily instantiated by
// InboundEndpointProperties_ARMGenerator()
var inboundEndpointProperties_ARMGenerator gopter.Gen

// InboundEndpointProperties_ARMGenerator returns a generator of InboundEndpointProperties_ARM instances for property testing.
func InboundEndpointProperties_ARMGenerator() gopter.Gen {
	if inboundEndpointProperties_ARMGenerator != nil {
		return inboundEndpointProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForInboundEndpointProperties_ARM(generators)
	inboundEndpointProperties_ARMGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties_ARM{}), generators)

	return inboundEndpointProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForInboundEndpointProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundEndpointProperties_ARM(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(IpConfiguration_ARMGenerator())
}

func Test_IpConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpConfiguration_ARM, IpConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpConfiguration_ARM runs a test to see if a specific instance of IpConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpConfiguration_ARM(subject IpConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpConfiguration_ARM
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

// Generator of IpConfiguration_ARM instances for property testing - lazily instantiated by
// IpConfiguration_ARMGenerator()
var ipConfiguration_ARMGenerator gopter.Gen

// IpConfiguration_ARMGenerator returns a generator of IpConfiguration_ARM instances for property testing.
// We first initialize ipConfiguration_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IpConfiguration_ARMGenerator() gopter.Gen {
	if ipConfiguration_ARMGenerator != nil {
		return ipConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_ARM(generators)
	ipConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_ARM(generators)
	AddRelatedPropertyGeneratorsForIpConfiguration_ARM(generators)
	ipConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_ARM{}), generators)

	return ipConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["PrivateIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIpAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IpConfiguration_PrivateIpAllocationMethod_Dynamic, IpConfiguration_PrivateIpAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForIpConfiguration_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIpConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResource_ARMGenerator())
}
