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

func Test_NetworkSecurityGroup_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_Spec_ARM, NetworkSecurityGroup_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_Spec_ARM runs a test to see if a specific instance of NetworkSecurityGroup_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_Spec_ARM(subject NetworkSecurityGroup_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_Spec_ARM
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

// Generator of NetworkSecurityGroup_Spec_ARM instances for property testing - lazily instantiated by
// NetworkSecurityGroup_Spec_ARMGenerator()
var networkSecurityGroup_Spec_ARMGenerator gopter.Gen

// NetworkSecurityGroup_Spec_ARMGenerator returns a generator of NetworkSecurityGroup_Spec_ARM instances for property testing.
// We first initialize networkSecurityGroup_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_Spec_ARMGenerator() gopter.Gen {
	if networkSecurityGroup_Spec_ARMGenerator != nil {
		return networkSecurityGroup_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM(generators)
	networkSecurityGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM(generators)
	networkSecurityGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Spec_ARM{}), generators)

	return networkSecurityGroup_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormat_ARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_ARM, NetworkSecurityGroupPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_ARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormat_ARM(subject NetworkSecurityGroupPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_ARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_ARM instances for property testing - lazily instantiated by
// NetworkSecurityGroupPropertiesFormat_ARMGenerator()
var networkSecurityGroupPropertiesFormat_ARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormat_ARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_ARM instances for property testing.
func NetworkSecurityGroupPropertiesFormat_ARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormat_ARMGenerator != nil {
		return networkSecurityGroupPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_ARM(generators)
	networkSecurityGroupPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_ARM{}), generators)

	return networkSecurityGroupPropertiesFormat_ARMGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["SecurityRules"] = gen.SliceOf(SecurityRule_ARMGenerator())
}

func Test_SecurityRule_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_ARM, SecurityRule_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_ARM runs a test to see if a specific instance of SecurityRule_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_ARM(subject SecurityRule_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_ARM
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

// Generator of SecurityRule_ARM instances for property testing - lazily instantiated by SecurityRule_ARMGenerator()
var securityRule_ARMGenerator gopter.Gen

// SecurityRule_ARMGenerator returns a generator of SecurityRule_ARM instances for property testing.
// We first initialize securityRule_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRule_ARMGenerator() gopter.Gen {
	if securityRule_ARMGenerator != nil {
		return securityRule_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_ARM(generators)
	securityRule_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRule_ARM(generators)
	securityRule_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_ARM{}), generators)

	return securityRule_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRule_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRule_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator())
}

func Test_SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM, SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM runs a test to see if a specific instance of SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(subject SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM
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

// Generator of SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing
// - lazily instantiated by SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()
var securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator gopter.Gen

// SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator returns a generator of SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing.
// We first initialize securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator != nil {
		return securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	return securityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_Allow, SecurityRuleAccess_Deny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_Inbound, SecurityRuleDirection_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah,
		SecurityRulePropertiesFormat_Protocol_Esp,
		SecurityRulePropertiesFormat_Protocol_Icmp,
		SecurityRulePropertiesFormat_Protocol_Star,
		SecurityRulePropertiesFormat_Protocol_Tcp,
		SecurityRulePropertiesFormat_Protocol_Udp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator())
}

func Test_ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM, ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM runs a test to see if a specific instance of ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM(subject ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM
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

// Generator of ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing
// - lazily instantiated by ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator()
var applicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator gopter.Gen

// ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator returns a generator of ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM instances for property testing.
func ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if applicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator != nil {
		return applicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM(generators)
	applicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM{}), generators)

	return applicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroup_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
