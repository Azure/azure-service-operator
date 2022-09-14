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

func Test_NetworkInterface_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_Spec_ARM, NetworkInterface_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_Spec_ARM runs a test to see if a specific instance of NetworkInterface_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_Spec_ARM(subject NetworkInterface_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_Spec_ARM
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

// Generator of NetworkInterface_Spec_ARM instances for property testing - lazily instantiated by
// NetworkInterface_Spec_ARMGenerator()
var networkInterface_Spec_ARMGenerator gopter.Gen

// NetworkInterface_Spec_ARMGenerator returns a generator of NetworkInterface_Spec_ARM instances for property testing.
// We first initialize networkInterface_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterface_Spec_ARMGenerator() gopter.Gen {
	if networkInterface_Spec_ARMGenerator != nil {
		return networkInterface_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Spec_ARM(generators)
	networkInterface_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterface_Spec_ARM(generators)
	networkInterface_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Spec_ARM{}), generators)

	return networkInterface_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterface_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterface_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(NetworkInterface_Properties_Spec_ARMGenerator())
}

func Test_NetworkInterface_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_Properties_Spec_ARM, NetworkInterface_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_Properties_Spec_ARM runs a test to see if a specific instance of NetworkInterface_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_Properties_Spec_ARM(subject NetworkInterface_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_Properties_Spec_ARM
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

// Generator of NetworkInterface_Properties_Spec_ARM instances for property testing - lazily instantiated by
// NetworkInterface_Properties_Spec_ARMGenerator()
var networkInterface_Properties_Spec_ARMGenerator gopter.Gen

// NetworkInterface_Properties_Spec_ARMGenerator returns a generator of NetworkInterface_Properties_Spec_ARM instances for property testing.
// We first initialize networkInterface_Properties_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterface_Properties_Spec_ARMGenerator() gopter.Gen {
	if networkInterface_Properties_Spec_ARMGenerator != nil {
		return networkInterface_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM(generators)
	networkInterface_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Properties_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM(generators)
	networkInterface_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Properties_Spec_ARM{}), generators)

	return networkInterface_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["EnableAcceleratedNetworking"] = gen.PtrOf(gen.Bool())
	gens["EnableIPForwarding"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterface_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["DnsSettings"] = gen.PtrOf(NetworkInterfaceDnsSettings_ARMGenerator())
	gens["IpConfigurations"] = gen.SliceOf(NetworkInterface_Properties_IpConfigurations_Spec_ARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(SubResource_ARMGenerator())
}

func Test_NetworkInterface_Properties_IpConfigurations_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_Properties_IpConfigurations_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_Properties_IpConfigurations_Spec_ARM, NetworkInterface_Properties_IpConfigurations_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_Properties_IpConfigurations_Spec_ARM runs a test to see if a specific instance of NetworkInterface_Properties_IpConfigurations_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_Properties_IpConfigurations_Spec_ARM(subject NetworkInterface_Properties_IpConfigurations_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_Properties_IpConfigurations_Spec_ARM
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

// Generator of NetworkInterface_Properties_IpConfigurations_Spec_ARM instances for property testing - lazily
// instantiated by NetworkInterface_Properties_IpConfigurations_Spec_ARMGenerator()
var networkInterface_Properties_IpConfigurations_Spec_ARMGenerator gopter.Gen

// NetworkInterface_Properties_IpConfigurations_Spec_ARMGenerator returns a generator of NetworkInterface_Properties_IpConfigurations_Spec_ARM instances for property testing.
// We first initialize networkInterface_Properties_IpConfigurations_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterface_Properties_IpConfigurations_Spec_ARMGenerator() gopter.Gen {
	if networkInterface_Properties_IpConfigurations_Spec_ARMGenerator != nil {
		return networkInterface_Properties_IpConfigurations_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM(generators)
	networkInterface_Properties_IpConfigurations_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Properties_IpConfigurations_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM(generators)
	networkInterface_Properties_IpConfigurations_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Properties_IpConfigurations_Spec_ARM{}), generators)

	return networkInterface_Properties_IpConfigurations_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterface_Properties_IpConfigurations_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkInterfaceIPConfigurationPropertiesFormat_ARMGenerator())
}

func Test_NetworkInterfaceDnsSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceDnsSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceDnsSettings_ARM, NetworkInterfaceDnsSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceDnsSettings_ARM runs a test to see if a specific instance of NetworkInterfaceDnsSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceDnsSettings_ARM(subject NetworkInterfaceDnsSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceDnsSettings_ARM
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

// Generator of NetworkInterfaceDnsSettings_ARM instances for property testing - lazily instantiated by
// NetworkInterfaceDnsSettings_ARMGenerator()
var networkInterfaceDnsSettings_ARMGenerator gopter.Gen

// NetworkInterfaceDnsSettings_ARMGenerator returns a generator of NetworkInterfaceDnsSettings_ARM instances for property testing.
func NetworkInterfaceDnsSettings_ARMGenerator() gopter.Gen {
	if networkInterfaceDnsSettings_ARMGenerator != nil {
		return networkInterfaceDnsSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettings_ARM(generators)
	networkInterfaceDnsSettings_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceDnsSettings_ARM{}), generators)

	return networkInterfaceDnsSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettings_ARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
	gens["InternalDnsNameLabel"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_ARM, SubResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_ARM runs a test to see if a specific instance of SubResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_ARM(subject SubResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_ARM
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

// Generator of SubResource_ARM instances for property testing - lazily instantiated by SubResource_ARMGenerator()
var subResource_ARMGenerator gopter.Gen

// SubResource_ARMGenerator returns a generator of SubResource_ARM instances for property testing.
func SubResource_ARMGenerator() gopter.Gen {
	if subResource_ARMGenerator != nil {
		return subResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_ARM(generators)
	subResource_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_ARM{}), generators)

	return subResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterfaceIPConfigurationPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceIPConfigurationPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormat_ARM, NetworkInterfaceIPConfigurationPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormat_ARM runs a test to see if a specific instance of NetworkInterfaceIPConfigurationPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(subject NetworkInterfaceIPConfigurationPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceIPConfigurationPropertiesFormat_ARM
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

// Generator of NetworkInterfaceIPConfigurationPropertiesFormat_ARM instances for property testing - lazily instantiated
// by NetworkInterfaceIPConfigurationPropertiesFormat_ARMGenerator()
var networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator gopter.Gen

// NetworkInterfaceIPConfigurationPropertiesFormat_ARMGenerator returns a generator of NetworkInterfaceIPConfigurationPropertiesFormat_ARM instances for property testing.
// We first initialize networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfaceIPConfigurationPropertiesFormat_ARMGenerator() gopter.Gen {
	if networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator != nil {
		return networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(generators)
	networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfigurationPropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(generators)
	networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfigurationPropertiesFormat_ARM{}), generators)

	return networkInterfaceIPConfigurationPropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["Primary"] = gen.PtrOf(gen.Bool())
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAddressVersion_IPv4, NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAddressVersion_IPv6))
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAllocationMethod_Dynamic, NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["ApplicationGatewayBackendAddressPools"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["ApplicationSecurityGroups"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["LoadBalancerBackendAddressPools"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["LoadBalancerInboundNatRules"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["PublicIPAddress"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["Subnet"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["VirtualNetworkTaps"] = gen.SliceOf(SubResource_ARMGenerator())
}
