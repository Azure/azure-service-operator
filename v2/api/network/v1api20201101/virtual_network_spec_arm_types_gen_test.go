// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

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

func Test_ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM, ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM(subject ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing
// - lazily instantiated by ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var applicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
func ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if applicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return applicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	applicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return applicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_DhcpOptions_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions_ARM, DhcpOptions_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions_ARM runs a test to see if a specific instance of DhcpOptions_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions_ARM(subject DhcpOptions_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_ARM
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

// Generator of DhcpOptions_ARM instances for property testing - lazily instantiated by DhcpOptions_ARMGenerator()
var dhcpOptions_ARMGenerator gopter.Gen

// DhcpOptions_ARMGenerator returns a generator of DhcpOptions_ARM instances for property testing.
func DhcpOptions_ARMGenerator() gopter.Gen {
	if dhcpOptions_ARMGenerator != nil {
		return dhcpOptions_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions_ARM(generators)
	dhcpOptions_ARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_ARM{}), generators)

	return dhcpOptions_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions_ARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM, NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM(subject NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var networkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
func NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if networkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return networkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	networkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return networkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM, RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM(subject RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var routeTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
func RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if routeTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return routeTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	routeTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return routeTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM, ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM(subject ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var serviceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
func ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if serviceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return serviceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	serviceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return serviceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM, SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(subject SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
// We first initialize subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return subnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["PrivateEndpointNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Disabled, SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Enabled))
	gens["PrivateLinkServiceNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled, SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Enabled))
}

// AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["ApplicationGatewayIpConfigurations"] = gen.SliceOf(ApplicationGatewayIPConfiguration_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
	gens["Delegations"] = gen.SliceOf(Delegation_ARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["NatGateway"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(NetworkSecurityGroupSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
	gens["RouteTable"] = gen.PtrOf(RouteTableSpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
	gens["ServiceEndpointPolicies"] = gen.SliceOf(ServiceEndpointPolicySpec_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
	gens["ServiceEndpoints"] = gen.SliceOf(ServiceEndpointPropertiesFormat_ARMGenerator())
}

func Test_Subnet_VirtualNetwork_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_VirtualNetwork_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_VirtualNetwork_SubResourceEmbedded_ARM, Subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_VirtualNetwork_SubResourceEmbedded_ARM runs a test to see if a specific instance of Subnet_VirtualNetwork_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(subject Subnet_VirtualNetwork_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_VirtualNetwork_SubResourceEmbedded_ARM
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

// Generator of Subnet_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing - lazily instantiated by
// Subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator()
var subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator gopter.Gen

// Subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator returns a generator of Subnet_VirtualNetwork_SubResourceEmbedded_ARM instances for property testing.
// We first initialize subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator != nil {
		return subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(Subnet_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(generators)
	subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(Subnet_VirtualNetwork_SubResourceEmbedded_ARM{}), generators)

	return subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubnet_VirtualNetwork_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubnetPropertiesFormat_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
}

func Test_VirtualNetworkBgpCommunities_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_ARM, VirtualNetworkBgpCommunities_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_ARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_ARM(subject VirtualNetworkBgpCommunities_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_ARM
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

// Generator of VirtualNetworkBgpCommunities_ARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_ARMGenerator()
var virtualNetworkBgpCommunities_ARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_ARMGenerator returns a generator of VirtualNetworkBgpCommunities_ARM instances for property testing.
func VirtualNetworkBgpCommunities_ARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_ARMGenerator != nil {
		return virtualNetworkBgpCommunities_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_ARM(generators)
	virtualNetworkBgpCommunities_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_ARM{}), generators)

	return virtualNetworkBgpCommunities_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_ARM(gens map[string]gopter.Gen) {
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkPeering_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_ARM, VirtualNetworkPeering_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_ARM runs a test to see if a specific instance of VirtualNetworkPeering_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_ARM(subject VirtualNetworkPeering_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_ARM
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

// Generator of VirtualNetworkPeering_ARM instances for property testing - lazily instantiated by
// VirtualNetworkPeering_ARMGenerator()
var virtualNetworkPeering_ARMGenerator gopter.Gen

// VirtualNetworkPeering_ARMGenerator returns a generator of VirtualNetworkPeering_ARM instances for property testing.
// We first initialize virtualNetworkPeering_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeering_ARMGenerator() gopter.Gen {
	if virtualNetworkPeering_ARMGenerator != nil {
		return virtualNetworkPeering_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_ARM(generators)
	virtualNetworkPeering_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeering_ARM(generators)
	virtualNetworkPeering_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_ARM{}), generators)

	return virtualNetworkPeering_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeering_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeering_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormat_ARMGenerator())
}

func Test_VirtualNetworkPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPropertiesFormat_ARM, VirtualNetworkPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPropertiesFormat_ARM runs a test to see if a specific instance of VirtualNetworkPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPropertiesFormat_ARM(subject VirtualNetworkPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPropertiesFormat_ARM
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

// Generator of VirtualNetworkPropertiesFormat_ARM instances for property testing - lazily instantiated by
// VirtualNetworkPropertiesFormat_ARMGenerator()
var virtualNetworkPropertiesFormat_ARMGenerator gopter.Gen

// VirtualNetworkPropertiesFormat_ARMGenerator returns a generator of VirtualNetworkPropertiesFormat_ARM instances for property testing.
// We first initialize virtualNetworkPropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPropertiesFormat_ARMGenerator() gopter.Gen {
	if virtualNetworkPropertiesFormat_ARMGenerator != nil {
		return virtualNetworkPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM(generators)
	virtualNetworkPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM(generators)
	virtualNetworkPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_ARM{}), generators)

	return virtualNetworkPropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpace_ARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_ARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptions_ARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["Subnets"] = gen.SliceOf(Subnet_VirtualNetwork_SubResourceEmbedded_ARMGenerator())
	gens["VirtualNetworkPeerings"] = gen.SliceOf(VirtualNetworkPeering_ARMGenerator())
}

func Test_VirtualNetwork_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec_ARM, VirtualNetwork_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec_ARM runs a test to see if a specific instance of VirtualNetwork_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec_ARM(subject VirtualNetwork_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec_ARM
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

// Generator of VirtualNetwork_Spec_ARM instances for property testing - lazily instantiated by
// VirtualNetwork_Spec_ARMGenerator()
var virtualNetwork_Spec_ARMGenerator gopter.Gen

// VirtualNetwork_Spec_ARMGenerator returns a generator of VirtualNetwork_Spec_ARM instances for property testing.
// We first initialize virtualNetwork_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_Spec_ARMGenerator() gopter.Gen {
	if virtualNetwork_Spec_ARMGenerator != nil {
		return virtualNetwork_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_ARM(generators)
	virtualNetwork_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_ARM(generators)
	virtualNetwork_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec_ARM{}), generators)

	return virtualNetwork_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworkPropertiesFormat_ARMGenerator())
}
