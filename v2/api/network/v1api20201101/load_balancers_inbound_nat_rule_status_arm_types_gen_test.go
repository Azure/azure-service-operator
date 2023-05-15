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

func Test_LoadBalancers_InboundNatRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancers_InboundNatRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS_ARM, LoadBalancers_InboundNatRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS_ARM runs a test to see if a specific instance of LoadBalancers_InboundNatRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS_ARM(subject LoadBalancers_InboundNatRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancers_InboundNatRule_STATUS_ARM
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

// Generator of LoadBalancers_InboundNatRule_STATUS_ARM instances for property testing - lazily instantiated by
// LoadBalancers_InboundNatRule_STATUS_ARMGenerator()
var loadBalancers_InboundNatRule_STATUS_ARMGenerator gopter.Gen

// LoadBalancers_InboundNatRule_STATUS_ARMGenerator returns a generator of LoadBalancers_InboundNatRule_STATUS_ARM instances for property testing.
// We first initialize loadBalancers_InboundNatRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LoadBalancers_InboundNatRule_STATUS_ARMGenerator() gopter.Gen {
	if loadBalancers_InboundNatRule_STATUS_ARMGenerator != nil {
		return loadBalancers_InboundNatRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM(generators)
	loadBalancers_InboundNatRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM(generators)
	loadBalancers_InboundNatRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_STATUS_ARM{}), generators)

	return loadBalancers_InboundNatRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundNatRulePropertiesFormat_STATUS_ARMGenerator())
}

func Test_InboundNatRulePropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundNatRulePropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_ARM, InboundNatRulePropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_ARM runs a test to see if a specific instance of InboundNatRulePropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundNatRulePropertiesFormat_STATUS_ARM(subject InboundNatRulePropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundNatRulePropertiesFormat_STATUS_ARM
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

// Generator of InboundNatRulePropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// InboundNatRulePropertiesFormat_STATUS_ARMGenerator()
var inboundNatRulePropertiesFormat_STATUS_ARMGenerator gopter.Gen

// InboundNatRulePropertiesFormat_STATUS_ARMGenerator returns a generator of InboundNatRulePropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize inboundNatRulePropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func InboundNatRulePropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if inboundNatRulePropertiesFormat_STATUS_ARMGenerator != nil {
		return inboundNatRulePropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM(generators)
	inboundNatRulePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM(generators)
	inboundNatRulePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_STATUS_ARM{}), generators)

	return inboundNatRulePropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BackendPort"] = gen.PtrOf(gen.Int())
	gens["EnableFloatingIP"] = gen.PtrOf(gen.Bool())
	gens["EnableTcpReset"] = gen.PtrOf(gen.Bool())
	gens["FrontendPort"] = gen.PtrOf(gen.Int())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(TransportProtocol_STATUS_All, TransportProtocol_STATUS_Tcp, TransportProtocol_STATUS_Udp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BackendIPConfiguration"] = gen.PtrOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator())
	gens["FrontendIPConfiguration"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
}

func Test_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM, NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM runs a test to see if a specific instance of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM(subject NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM
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

// Generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM instances
// for property testing - lazily instantiated by
// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator()
var networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator gopter.Gen

// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator returns a generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM instances for property testing.
func NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator != nil {
		return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM(generators)
	networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM{}), generators)

	return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS_ARM, SubResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS_ARM runs a test to see if a specific instance of SubResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS_ARM(subject SubResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS_ARM
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

// Generator of SubResource_STATUS_ARM instances for property testing - lazily instantiated by
// SubResource_STATUS_ARMGenerator()
var subResource_STATUS_ARMGenerator gopter.Gen

// SubResource_STATUS_ARMGenerator returns a generator of SubResource_STATUS_ARM instances for property testing.
func SubResource_STATUS_ARMGenerator() gopter.Gen {
	if subResource_STATUS_ARMGenerator != nil {
		return subResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM(generators)
	subResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS_ARM{}), generators)

	return subResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
