// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

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

func Test_VirtualNetworks_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecARM, VirtualNetworksSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecARM runs a test to see if a specific instance of VirtualNetworks_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecARM(subject VirtualNetworks_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_SpecARM
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

// Generator of VirtualNetworks_SpecARM instances for property testing - lazily instantiated by
// VirtualNetworksSpecARMGenerator()
var virtualNetworksSpecARMGenerator gopter.Gen

// VirtualNetworksSpecARMGenerator returns a generator of VirtualNetworks_SpecARM instances for property testing.
// We first initialize virtualNetworksSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecARMGenerator() gopter.Gen {
	if virtualNetworksSpecARMGenerator != nil {
		return virtualNetworksSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecARM(generators)
	virtualNetworksSpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecARM(generators)
	virtualNetworksSpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_SpecARM{}), generators)

	return virtualNetworksSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworksSpecPropertiesARMGenerator())
}

func Test_VirtualNetworks_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecPropertiesARM, VirtualNetworksSpecPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecPropertiesARM runs a test to see if a specific instance of VirtualNetworks_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecPropertiesARM(subject VirtualNetworks_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec_PropertiesARM
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

// Generator of VirtualNetworks_Spec_PropertiesARM instances for property testing - lazily instantiated by
// VirtualNetworksSpecPropertiesARMGenerator()
var virtualNetworksSpecPropertiesARMGenerator gopter.Gen

// VirtualNetworksSpecPropertiesARMGenerator returns a generator of VirtualNetworks_Spec_PropertiesARM instances for property testing.
// We first initialize virtualNetworksSpecPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecPropertiesARMGenerator() gopter.Gen {
	if virtualNetworksSpecPropertiesARMGenerator != nil {
		return virtualNetworksSpecPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesARM(generators)
	virtualNetworksSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesARM(generators)
	virtualNetworksSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_PropertiesARM{}), generators)

	return virtualNetworksSpecPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceARMGenerator())
	gens["Subnets"] = gen.SliceOf(VirtualNetworksSpecPropertiesSubnetsARMGenerator())
	gens["VirtualNetworkPeerings"] = gen.SliceOf(VirtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator())
}

func Test_DhcpOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptionsARM, DhcpOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptionsARM runs a test to see if a specific instance of DhcpOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptionsARM(subject DhcpOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptionsARM
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

// Generator of DhcpOptionsARM instances for property testing - lazily instantiated by DhcpOptionsARMGenerator()
var dhcpOptionsARMGenerator gopter.Gen

// DhcpOptionsARMGenerator returns a generator of DhcpOptionsARM instances for property testing.
func DhcpOptionsARMGenerator() gopter.Gen {
	if dhcpOptionsARMGenerator != nil {
		return dhcpOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptionsARM(generators)
	dhcpOptionsARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptionsARM{}), generators)

	return dhcpOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptionsARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetworks_Spec_Properties_SubnetsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec_Properties_SubnetsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsARM, VirtualNetworksSpecPropertiesSubnetsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsARM runs a test to see if a specific instance of VirtualNetworks_Spec_Properties_SubnetsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsARM(subject VirtualNetworks_Spec_Properties_SubnetsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec_Properties_SubnetsARM
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

// Generator of VirtualNetworks_Spec_Properties_SubnetsARM instances for property testing - lazily instantiated by
// VirtualNetworksSpecPropertiesSubnetsARMGenerator()
var virtualNetworksSpecPropertiesSubnetsARMGenerator gopter.Gen

// VirtualNetworksSpecPropertiesSubnetsARMGenerator returns a generator of VirtualNetworks_Spec_Properties_SubnetsARM instances for property testing.
// We first initialize virtualNetworksSpecPropertiesSubnetsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecPropertiesSubnetsARMGenerator() gopter.Gen {
	if virtualNetworksSpecPropertiesSubnetsARMGenerator != nil {
		return virtualNetworksSpecPropertiesSubnetsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM(generators)
	virtualNetworksSpecPropertiesSubnetsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_SubnetsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM(generators)
	virtualNetworksSpecPropertiesSubnetsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_SubnetsARM{}), generators)

	return virtualNetworksSpecPropertiesSubnetsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator())
}

func Test_VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM, VirtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM runs a test to see if a specific instance of VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(subject VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM
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

// Generator of VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM instances for property testing - lazily
// instantiated by VirtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator()
var virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator gopter.Gen

// VirtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator returns a generator of VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM instances for property testing.
// We first initialize virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator() gopter.Gen {
	if virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator != nil {
		return virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(generators)
	virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(generators)
	virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM{}), generators)

	return virtualNetworksSpecPropertiesVirtualNetworkPeeringsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesVirtualNetworkPeeringsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormatARMGenerator())
}

func Test_VirtualNetworks_Spec_Properties_Subnets_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec_Properties_Subnets_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesARM, VirtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesARM runs a test to see if a specific instance of VirtualNetworks_Spec_Properties_Subnets_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(subject VirtualNetworks_Spec_Properties_Subnets_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec_Properties_Subnets_PropertiesARM
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

// Generator of VirtualNetworks_Spec_Properties_Subnets_PropertiesARM instances for property testing - lazily
// instantiated by VirtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator()
var virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator gopter.Gen

// VirtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator returns a generator of VirtualNetworks_Spec_Properties_Subnets_PropertiesARM instances for property testing.
// We first initialize virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator() gopter.Gen {
	if virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator != nil {
		return virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(generators)
	virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_Subnets_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(generators)
	virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_Subnets_PropertiesARM{}), generators)

	return virtualNetworksSpecPropertiesSubnetsPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["PrivateEndpointNetworkPolicies"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkServiceNetworkPolicies"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesARM(gens map[string]gopter.Gen) {
	gens["Delegations"] = gen.SliceOf(VirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceARMGenerator())
	gens["NatGateway"] = gen.PtrOf(SubResourceARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(SubResourceARMGenerator())
	gens["RouteTable"] = gen.PtrOf(SubResourceARMGenerator())
	gens["ServiceEndpointPolicies"] = gen.SliceOf(SubResourceARMGenerator())
	gens["ServiceEndpoints"] = gen.SliceOf(ServiceEndpointPropertiesFormatARMGenerator())
}

func Test_VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM, VirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM runs a test to see if a specific instance of VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(subject VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM
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

// Generator of VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM instances for property testing -
// lazily instantiated by VirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator()
var virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator gopter.Gen

// VirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator returns a generator of VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM instances for property testing.
// We first initialize virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator() gopter.Gen {
	if virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator != nil {
		return virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(generators)
	virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(generators)
	virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM{}), generators)

	return virtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpecPropertiesSubnetsPropertiesDelegationsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServiceDelegationPropertiesFormatARMGenerator())
}
