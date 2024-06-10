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

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM, ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM
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

// Generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()
var applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator gopter.Gen

// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator returns a generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for property testing.
func ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator != nil {
		return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(generators)
	applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM{}), generators)

	return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkSecurityGroups_SecurityRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SecurityRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS_ARM, NetworkSecurityGroups_SecurityRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS_ARM runs a test to see if a specific instance of NetworkSecurityGroups_SecurityRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_STATUS_ARM(subject NetworkSecurityGroups_SecurityRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SecurityRule_STATUS_ARM
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

// Generator of NetworkSecurityGroups_SecurityRule_STATUS_ARM instances for property testing - lazily instantiated by
// NetworkSecurityGroups_SecurityRule_STATUS_ARMGenerator()
var networkSecurityGroups_SecurityRule_STATUS_ARMGenerator gopter.Gen

// NetworkSecurityGroups_SecurityRule_STATUS_ARMGenerator returns a generator of NetworkSecurityGroups_SecurityRule_STATUS_ARM instances for property testing.
// We first initialize networkSecurityGroups_SecurityRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroups_SecurityRule_STATUS_ARMGenerator() gopter.Gen {
	if networkSecurityGroups_SecurityRule_STATUS_ARMGenerator != nil {
		return networkSecurityGroups_SecurityRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM(generators)
	networkSecurityGroups_SecurityRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM(generators)
	networkSecurityGroups_SecurityRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_STATUS_ARM{}), generators)

	return networkSecurityGroups_SecurityRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecurityRulePropertiesFormat_STATUS_ARMGenerator())
}

func Test_SecurityRulePropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRulePropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRulePropertiesFormat_STATUS_ARM, SecurityRulePropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRulePropertiesFormat_STATUS_ARM runs a test to see if a specific instance of SecurityRulePropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRulePropertiesFormat_STATUS_ARM(subject SecurityRulePropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRulePropertiesFormat_STATUS_ARM
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

// Generator of SecurityRulePropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// SecurityRulePropertiesFormat_STATUS_ARMGenerator()
var securityRulePropertiesFormat_STATUS_ARMGenerator gopter.Gen

// SecurityRulePropertiesFormat_STATUS_ARMGenerator returns a generator of SecurityRulePropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize securityRulePropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRulePropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if securityRulePropertiesFormat_STATUS_ARMGenerator != nil {
		return securityRulePropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM(generators)
	securityRulePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM(generators)
	securityRulePropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_STATUS_ARM{}), generators)

	return securityRulePropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_STATUS_Allow, SecurityRuleAccess_STATUS_Deny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_STATUS_Inbound, SecurityRuleDirection_STATUS_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_STATUS_Ah,
		SecurityRulePropertiesFormat_Protocol_STATUS_Esp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Icmp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Star,
		SecurityRulePropertiesFormat_Protocol_STATUS_Tcp,
		SecurityRulePropertiesFormat_Protocol_STATUS_Udp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator())
}
