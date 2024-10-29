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

func Test_DnsResolversInboundEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolversInboundEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolversInboundEndpoint_STATUS, DnsResolversInboundEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolversInboundEndpoint_STATUS runs a test to see if a specific instance of DnsResolversInboundEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolversInboundEndpoint_STATUS(subject DnsResolversInboundEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolversInboundEndpoint_STATUS
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

// Generator of DnsResolversInboundEndpoint_STATUS instances for property testing - lazily instantiated by
// DnsResolversInboundEndpoint_STATUSGenerator()
var dnsResolversInboundEndpoint_STATUSGenerator gopter.Gen

// DnsResolversInboundEndpoint_STATUSGenerator returns a generator of DnsResolversInboundEndpoint_STATUS instances for property testing.
// We first initialize dnsResolversInboundEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolversInboundEndpoint_STATUSGenerator() gopter.Gen {
	if dnsResolversInboundEndpoint_STATUSGenerator != nil {
		return dnsResolversInboundEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS(generators)
	dnsResolversInboundEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolversInboundEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS(generators)
	dnsResolversInboundEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsResolversInboundEndpoint_STATUS{}), generators)

	return dnsResolversInboundEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundEndpointProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_InboundEndpointProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundEndpointProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundEndpointProperties_STATUS, InboundEndpointProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundEndpointProperties_STATUS runs a test to see if a specific instance of InboundEndpointProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundEndpointProperties_STATUS(subject InboundEndpointProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundEndpointProperties_STATUS
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

// Generator of InboundEndpointProperties_STATUS instances for property testing - lazily instantiated by
// InboundEndpointProperties_STATUSGenerator()
var inboundEndpointProperties_STATUSGenerator gopter.Gen

// InboundEndpointProperties_STATUSGenerator returns a generator of InboundEndpointProperties_STATUS instances for property testing.
// We first initialize inboundEndpointProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func InboundEndpointProperties_STATUSGenerator() gopter.Gen {
	if inboundEndpointProperties_STATUSGenerator != nil {
		return inboundEndpointProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS(generators)
	inboundEndpointProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS(generators)
	inboundEndpointProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties_STATUS{}), generators)

	return inboundEndpointProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(IpConfiguration_STATUSGenerator())
}

func Test_IpConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpConfiguration_STATUS, IpConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpConfiguration_STATUS runs a test to see if a specific instance of IpConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForIpConfiguration_STATUS(subject IpConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpConfiguration_STATUS
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

// Generator of IpConfiguration_STATUS instances for property testing - lazily instantiated by
// IpConfiguration_STATUSGenerator()
var ipConfiguration_STATUSGenerator gopter.Gen

// IpConfiguration_STATUSGenerator returns a generator of IpConfiguration_STATUS instances for property testing.
// We first initialize ipConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IpConfiguration_STATUSGenerator() gopter.Gen {
	if ipConfiguration_STATUSGenerator != nil {
		return ipConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_STATUS(generators)
	ipConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForIpConfiguration_STATUS(generators)
	ipConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_STATUS{}), generators)

	return ipConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForIpConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIpAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IpConfiguration_PrivateIpAllocationMethod_STATUS_Dynamic, IpConfiguration_PrivateIpAllocationMethod_STATUS_Static))
}

// AddRelatedPropertyGeneratorsForIpConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIpConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResource_STATUSGenerator())
}
