// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetworksSubnet_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksSubnet_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSubnet_SpecARM, VirtualNetworksSubnet_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSubnet_SpecARM runs a test to see if a specific instance of VirtualNetworksSubnet_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSubnet_SpecARM(subject VirtualNetworksSubnet_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksSubnet_SpecARM
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

// Generator of VirtualNetworksSubnet_SpecARM instances for property testing - lazily instantiated by
//VirtualNetworksSubnet_SpecARMGenerator()
var virtualNetworksSubnet_specARMGenerator gopter.Gen

// VirtualNetworksSubnet_SpecARMGenerator returns a generator of VirtualNetworksSubnet_SpecARM instances for property testing.
// We first initialize virtualNetworksSubnet_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSubnet_SpecARMGenerator() gopter.Gen {
	if virtualNetworksSubnet_specARMGenerator != nil {
		return virtualNetworksSubnet_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_SpecARM(generators)
	virtualNetworksSubnet_specARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksSubnet_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_SpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_SpecARM(generators)
	virtualNetworksSubnet_specARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksSubnet_SpecARM{}), generators)

	return virtualNetworksSubnet_specARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubnetPropertiesFormatARMGenerator())
}

func Test_SubnetPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubnetPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetPropertiesFormatARM, SubnetPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetPropertiesFormatARM runs a test to see if a specific instance of SubnetPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetPropertiesFormatARM(subject SubnetPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubnetPropertiesFormatARM
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

// Generator of SubnetPropertiesFormatARM instances for property testing - lazily instantiated by
//SubnetPropertiesFormatARMGenerator()
var subnetPropertiesFormatARMGenerator gopter.Gen

// SubnetPropertiesFormatARMGenerator returns a generator of SubnetPropertiesFormatARM instances for property testing.
// We first initialize subnetPropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubnetPropertiesFormatARMGenerator() gopter.Gen {
	if subnetPropertiesFormatARMGenerator != nil {
		return subnetPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormatARM(generators)
	subnetPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForSubnetPropertiesFormatARM(generators)
	subnetPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormatARM{}), generators)

	return subnetPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnetPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["PrivateEndpointNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormatPrivateEndpointNetworkPoliciesDisabled, SubnetPropertiesFormatPrivateEndpointNetworkPoliciesEnabled))
	gens["PrivateLinkServiceNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormatPrivateLinkServiceNetworkPoliciesDisabled, SubnetPropertiesFormatPrivateLinkServiceNetworkPoliciesEnabled))
}

// AddRelatedPropertyGeneratorsForSubnetPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubnetPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["ApplicationGatewayIpConfigurations"] = gen.SliceOf(ApplicationGatewayIPConfigurationARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceARMGenerator())
	gens["NatGateway"] = gen.PtrOf(SubResourceARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(NetworkSecurityGroupSpecARMGenerator())
	gens["RouteTable"] = gen.PtrOf(RouteTableSpecARMGenerator())
	gens["ServiceEndpointPolicies"] = gen.SliceOf(ServiceEndpointPolicySpecARMGenerator())
	gens["ServiceEndpoints"] = gen.SliceOf(ServiceEndpointPropertiesFormatARMGenerator())
}

func Test_ApplicationGatewayIPConfigurationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewayIPConfigurationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewayIPConfigurationARM, ApplicationGatewayIPConfigurationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewayIPConfigurationARM runs a test to see if a specific instance of ApplicationGatewayIPConfigurationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewayIPConfigurationARM(subject ApplicationGatewayIPConfigurationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewayIPConfigurationARM
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

// Generator of ApplicationGatewayIPConfigurationARM instances for property testing - lazily instantiated by
//ApplicationGatewayIPConfigurationARMGenerator()
var applicationGatewayIPConfigurationARMGenerator gopter.Gen

// ApplicationGatewayIPConfigurationARMGenerator returns a generator of ApplicationGatewayIPConfigurationARM instances for property testing.
// We first initialize applicationGatewayIPConfigurationARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApplicationGatewayIPConfigurationARMGenerator() gopter.Gen {
	if applicationGatewayIPConfigurationARMGenerator != nil {
		return applicationGatewayIPConfigurationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationGatewayIPConfigurationARM(generators)
	applicationGatewayIPConfigurationARMGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewayIPConfigurationARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationGatewayIPConfigurationARM(generators)
	AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationARM(generators)
	applicationGatewayIPConfigurationARMGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewayIPConfigurationARM{}), generators)

	return applicationGatewayIPConfigurationARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationGatewayIPConfigurationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationGatewayIPConfigurationARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApplicationGatewayIPConfigurationPropertiesFormatARMGenerator())
}

func Test_RouteTableSpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTableSpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTableSpecARM, RouteTableSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTableSpecARM runs a test to see if a specific instance of RouteTableSpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTableSpecARM(subject RouteTableSpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTableSpecARM
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

// Generator of RouteTableSpecARM instances for property testing - lazily instantiated by RouteTableSpecARMGenerator()
var routeTableSpecARMGenerator gopter.Gen

// RouteTableSpecARMGenerator returns a generator of RouteTableSpecARM instances for property testing.
// We first initialize routeTableSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTableSpecARMGenerator() gopter.Gen {
	if routeTableSpecARMGenerator != nil {
		return routeTableSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTableSpecARM(generators)
	routeTableSpecARMGenerator = gen.Struct(reflect.TypeOf(RouteTableSpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTableSpecARM(generators)
	AddRelatedPropertyGeneratorsForRouteTableSpecARM(generators)
	routeTableSpecARMGenerator = gen.Struct(reflect.TypeOf(RouteTableSpecARM{}), generators)

	return routeTableSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTableSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTableSpecARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRouteTableSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTableSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RouteTablePropertiesFormatARMGenerator())
}

func Test_ServiceEndpointPolicySpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPolicySpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPolicySpecARM, ServiceEndpointPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPolicySpecARM runs a test to see if a specific instance of ServiceEndpointPolicySpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPolicySpecARM(subject ServiceEndpointPolicySpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPolicySpecARM
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

// Generator of ServiceEndpointPolicySpecARM instances for property testing - lazily instantiated by
//ServiceEndpointPolicySpecARMGenerator()
var serviceEndpointPolicySpecARMGenerator gopter.Gen

// ServiceEndpointPolicySpecARMGenerator returns a generator of ServiceEndpointPolicySpecARM instances for property testing.
// We first initialize serviceEndpointPolicySpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServiceEndpointPolicySpecARMGenerator() gopter.Gen {
	if serviceEndpointPolicySpecARMGenerator != nil {
		return serviceEndpointPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPolicySpecARM(generators)
	serviceEndpointPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicySpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPolicySpecARM(generators)
	AddRelatedPropertyGeneratorsForServiceEndpointPolicySpecARM(generators)
	serviceEndpointPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicySpecARM{}), generators)

	return serviceEndpointPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPolicySpecARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServiceEndpointPolicySpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceEndpointPolicySpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator())
}

func Test_ServiceEndpointPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPropertiesFormatARM, ServiceEndpointPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPropertiesFormatARM runs a test to see if a specific instance of ServiceEndpointPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPropertiesFormatARM(subject ServiceEndpointPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPropertiesFormatARM
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

// Generator of ServiceEndpointPropertiesFormatARM instances for property testing - lazily instantiated by
//ServiceEndpointPropertiesFormatARMGenerator()
var serviceEndpointPropertiesFormatARMGenerator gopter.Gen

// ServiceEndpointPropertiesFormatARMGenerator returns a generator of ServiceEndpointPropertiesFormatARM instances for property testing.
func ServiceEndpointPropertiesFormatARMGenerator() gopter.Gen {
	if serviceEndpointPropertiesFormatARMGenerator != nil {
		return serviceEndpointPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormatARM(generators)
	serviceEndpointPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPropertiesFormatARM{}), generators)

	return serviceEndpointPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["Locations"] = gen.SliceOf(gen.AlphaString())
	gens["Service"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApplicationGatewayIPConfigurationPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewayIPConfigurationPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewayIPConfigurationPropertiesFormatARM, ApplicationGatewayIPConfigurationPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewayIPConfigurationPropertiesFormatARM runs a test to see if a specific instance of ApplicationGatewayIPConfigurationPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewayIPConfigurationPropertiesFormatARM(subject ApplicationGatewayIPConfigurationPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewayIPConfigurationPropertiesFormatARM
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

// Generator of ApplicationGatewayIPConfigurationPropertiesFormatARM instances for property testing - lazily
//instantiated by ApplicationGatewayIPConfigurationPropertiesFormatARMGenerator()
var applicationGatewayIPConfigurationPropertiesFormatARMGenerator gopter.Gen

// ApplicationGatewayIPConfigurationPropertiesFormatARMGenerator returns a generator of ApplicationGatewayIPConfigurationPropertiesFormatARM instances for property testing.
func ApplicationGatewayIPConfigurationPropertiesFormatARMGenerator() gopter.Gen {
	if applicationGatewayIPConfigurationPropertiesFormatARMGenerator != nil {
		return applicationGatewayIPConfigurationPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationPropertiesFormatARM(generators)
	applicationGatewayIPConfigurationPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewayIPConfigurationPropertiesFormatARM{}), generators)

	return applicationGatewayIPConfigurationPropertiesFormatARMGenerator
}

// AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationGatewayIPConfigurationPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_RouteTablePropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablePropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablePropertiesFormatARM, RouteTablePropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablePropertiesFormatARM runs a test to see if a specific instance of RouteTablePropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablePropertiesFormatARM(subject RouteTablePropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablePropertiesFormatARM
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

// Generator of RouteTablePropertiesFormatARM instances for property testing - lazily instantiated by
//RouteTablePropertiesFormatARMGenerator()
var routeTablePropertiesFormatARMGenerator gopter.Gen

// RouteTablePropertiesFormatARMGenerator returns a generator of RouteTablePropertiesFormatARM instances for property testing.
func RouteTablePropertiesFormatARMGenerator() gopter.Gen {
	if routeTablePropertiesFormatARMGenerator != nil {
		return routeTablePropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablePropertiesFormatARM(generators)
	routeTablePropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(RouteTablePropertiesFormatARM{}), generators)

	return routeTablePropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablePropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablePropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
}

func Test_ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM, ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM runs a test to see if a specific instance of ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM(subject ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM
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

// Generator of ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM instances for
//property testing - lazily instantiated by
//ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator()
var serviceEndpointPolicyPropertiesFormat_virtualNetworksSubnet_subResourceEmbeddedARMGenerator gopter.Gen

// ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator returns a generator of ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM instances for property testing.
func ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if serviceEndpointPolicyPropertiesFormat_virtualNetworksSubnet_subResourceEmbeddedARMGenerator != nil {
		return serviceEndpointPolicyPropertiesFormat_virtualNetworksSubnet_subResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM(generators)
	serviceEndpointPolicyPropertiesFormat_virtualNetworksSubnet_subResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM{}), generators)

	return serviceEndpointPolicyPropertiesFormat_virtualNetworksSubnet_subResourceEmbeddedARMGenerator
}

// AddRelatedPropertyGeneratorsForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceEndpointPolicyPropertiesFormat_VirtualNetworksSubnet_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["ServiceEndpointPolicyDefinitions"] = gen.SliceOf(ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator())
}

func Test_ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM, ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM runs a test to see if a specific instance of ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM(subject ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM
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

// Generator of ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM instances for property
//testing - lazily instantiated by ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator()
var serviceEndpointPolicyDefinition_virtualNetworksSubnet_subResourceEmbeddedARMGenerator gopter.Gen

// ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator returns a generator of ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM instances for property testing.
func ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if serviceEndpointPolicyDefinition_virtualNetworksSubnet_subResourceEmbeddedARMGenerator != nil {
		return serviceEndpointPolicyDefinition_virtualNetworksSubnet_subResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM(generators)
	serviceEndpointPolicyDefinition_virtualNetworksSubnet_subResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM{}), generators)

	return serviceEndpointPolicyDefinition_virtualNetworksSubnet_subResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPolicyDefinition_VirtualNetworksSubnet_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
