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

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM runs a test to see if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM
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

// Generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing
// - lazily instantiated by NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()
var networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator gopter.Gen

// NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator returns a generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing.
// We first initialize networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator != nil {
		return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormat_STATUS_ARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUS_ARM, NetworkSecurityGroupPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(subject NetworkSecurityGroupPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_STATUS_ARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// NetworkSecurityGroupPropertiesFormat_STATUS_ARMGenerator()
var networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormat_STATUS_ARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator != nil {
		return networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(generators)
	networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(generators)
	networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUS_ARM{}), generators)

	return networkSecurityGroupPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRule_STATUS_ARMGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLog_STATUS_ARMGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator())
	gens["Subnets"] = gen.SliceOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator())
}

func Test_FlowLog_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLog_STATUS_ARM, FlowLog_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLog_STATUS_ARM runs a test to see if a specific instance of FlowLog_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLog_STATUS_ARM(subject FlowLog_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_STATUS_ARM
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

// Generator of FlowLog_STATUS_ARM instances for property testing - lazily instantiated by FlowLog_STATUS_ARMGenerator()
var flowLog_STATUS_ARMGenerator gopter.Gen

// FlowLog_STATUS_ARMGenerator returns a generator of FlowLog_STATUS_ARM instances for property testing.
func FlowLog_STATUS_ARMGenerator() gopter.Gen {
	if flowLog_STATUS_ARMGenerator != nil {
		return flowLog_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLog_STATUS_ARM(generators)
	flowLog_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FlowLog_STATUS_ARM{}), generators)

	return flowLog_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlowLog_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLog_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM runs a test to see if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM
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

// Generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing -
// lazily instantiated by NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()
var networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator gopter.Gen

// NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator returns a generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing.
func NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator != nil {
		return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SecurityRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_STATUS_ARM, SecurityRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_STATUS_ARM runs a test to see if a specific instance of SecurityRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_STATUS_ARM(subject SecurityRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_STATUS_ARM
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

// Generator of SecurityRule_STATUS_ARM instances for property testing - lazily instantiated by
// SecurityRule_STATUS_ARMGenerator()
var securityRule_STATUS_ARMGenerator gopter.Gen

// SecurityRule_STATUS_ARMGenerator returns a generator of SecurityRule_STATUS_ARM instances for property testing.
func SecurityRule_STATUS_ARMGenerator() gopter.Gen {
	if securityRule_STATUS_ARMGenerator != nil {
		return securityRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_STATUS_ARM(generators)
	securityRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_STATUS_ARM{}), generators)

	return securityRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM runs a test to see if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM
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

// Generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()
var subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator gopter.Gen

// Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator returns a generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing.
func Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator != nil {
		return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
