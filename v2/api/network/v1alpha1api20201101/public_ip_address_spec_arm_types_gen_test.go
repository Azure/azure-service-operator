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

func Test_PublicIPAddress_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddress_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddress_Spec_ARM, PublicIPAddress_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddress_Spec_ARM runs a test to see if a specific instance of PublicIPAddress_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddress_Spec_ARM(subject PublicIPAddress_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddress_Spec_ARM
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

// Generator of PublicIPAddress_Spec_ARM instances for property testing - lazily instantiated by
// PublicIPAddress_Spec_ARMGenerator()
var publicIPAddress_Spec_ARMGenerator gopter.Gen

// PublicIPAddress_Spec_ARMGenerator returns a generator of PublicIPAddress_Spec_ARM instances for property testing.
// We first initialize publicIPAddress_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddress_Spec_ARMGenerator() gopter.Gen {
	if publicIPAddress_Spec_ARMGenerator != nil {
		return publicIPAddress_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddress_Spec_ARM(generators)
	publicIPAddress_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddress_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddress_Spec_ARM(generators)
	publicIPAddress_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_Spec_ARM{}), generators)

	return publicIPAddress_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddress_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddress_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddress_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddress_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(PublicIPAddressPropertiesFormat_ARMGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPAddressSku_ARMGenerator())
}

func Test_PublicIPAddressPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressPropertiesFormat_ARM, PublicIPAddressPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressPropertiesFormat_ARM runs a test to see if a specific instance of PublicIPAddressPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressPropertiesFormat_ARM(subject PublicIPAddressPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressPropertiesFormat_ARM
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

// Generator of PublicIPAddressPropertiesFormat_ARM instances for property testing - lazily instantiated by
// PublicIPAddressPropertiesFormat_ARMGenerator()
var publicIPAddressPropertiesFormat_ARMGenerator gopter.Gen

// PublicIPAddressPropertiesFormat_ARMGenerator returns a generator of PublicIPAddressPropertiesFormat_ARM instances for property testing.
// We first initialize publicIPAddressPropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressPropertiesFormat_ARMGenerator() gopter.Gen {
	if publicIPAddressPropertiesFormat_ARMGenerator != nil {
		return publicIPAddressPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM(generators)
	publicIPAddressPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM(generators)
	publicIPAddressPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_ARM{}), generators)

	return publicIPAddressPropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_IPv4, IPVersion_IPv6))
	gens["PublicIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_Dynamic, IPAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["DdosSettings"] = gen.PtrOf(DdosSettings_ARMGenerator())
	gens["DnsSettings"] = gen.PtrOf(PublicIPAddressDnsSettings_ARMGenerator())
	gens["IpTags"] = gen.SliceOf(IpTag_ARMGenerator())
	gens["LinkedPublicIPAddress"] = gen.PtrOf(PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator())
	gens["PublicIPPrefix"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["ServicePublicIPAddress"] = gen.PtrOf(PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator())
}

func Test_PublicIPAddressSku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressSku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressSku_ARM, PublicIPAddressSku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressSku_ARM runs a test to see if a specific instance of PublicIPAddressSku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressSku_ARM(subject PublicIPAddressSku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressSku_ARM
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

// Generator of PublicIPAddressSku_ARM instances for property testing - lazily instantiated by
// PublicIPAddressSku_ARMGenerator()
var publicIPAddressSku_ARMGenerator gopter.Gen

// PublicIPAddressSku_ARMGenerator returns a generator of PublicIPAddressSku_ARM instances for property testing.
func PublicIPAddressSku_ARMGenerator() gopter.Gen {
	if publicIPAddressSku_ARMGenerator != nil {
		return publicIPAddressSku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressSku_ARM(generators)
	publicIPAddressSku_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressSku_ARM{}), generators)

	return publicIPAddressSku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSku_Name_Basic, PublicIPAddressSku_Name_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSku_Tier_Global, PublicIPAddressSku_Tier_Regional))
}

func Test_DdosSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DdosSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDdosSettings_ARM, DdosSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDdosSettings_ARM runs a test to see if a specific instance of DdosSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDdosSettings_ARM(subject DdosSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DdosSettings_ARM
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

// Generator of DdosSettings_ARM instances for property testing - lazily instantiated by DdosSettings_ARMGenerator()
var ddosSettings_ARMGenerator gopter.Gen

// DdosSettings_ARMGenerator returns a generator of DdosSettings_ARM instances for property testing.
// We first initialize ddosSettings_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DdosSettings_ARMGenerator() gopter.Gen {
	if ddosSettings_ARMGenerator != nil {
		return ddosSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettings_ARM(generators)
	ddosSettings_ARMGenerator = gen.Struct(reflect.TypeOf(DdosSettings_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettings_ARM(generators)
	AddRelatedPropertyGeneratorsForDdosSettings_ARM(generators)
	ddosSettings_ARMGenerator = gen.Struct(reflect.TypeOf(DdosSettings_ARM{}), generators)

	return ddosSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDdosSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDdosSettings_ARM(gens map[string]gopter.Gen) {
	gens["ProtectedIP"] = gen.PtrOf(gen.Bool())
	gens["ProtectionCoverage"] = gen.PtrOf(gen.OneConstOf(DdosSettings_ProtectionCoverage_Basic, DdosSettings_ProtectionCoverage_Standard))
}

// AddRelatedPropertyGeneratorsForDdosSettings_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDdosSettings_ARM(gens map[string]gopter.Gen) {
	gens["DdosCustomPolicy"] = gen.PtrOf(SubResource_ARMGenerator())
}

func Test_IpTag_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpTag_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpTag_ARM, IpTag_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpTag_ARM runs a test to see if a specific instance of IpTag_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpTag_ARM(subject IpTag_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpTag_ARM
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

// Generator of IpTag_ARM instances for property testing - lazily instantiated by IpTag_ARMGenerator()
var ipTag_ARMGenerator gopter.Gen

// IpTag_ARMGenerator returns a generator of IpTag_ARM instances for property testing.
func IpTag_ARMGenerator() gopter.Gen {
	if ipTag_ARMGenerator != nil {
		return ipTag_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpTag_ARM(generators)
	ipTag_ARMGenerator = gen.Struct(reflect.TypeOf(IpTag_ARM{}), generators)

	return ipTag_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpTag_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpTag_ARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM, NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM runs a test to see if a specific instance of NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM(subject NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM
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

// Generator of NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator()
var natGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator gopter.Gen

// NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator returns a generator of NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM instances for property testing.
func NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if natGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator != nil {
		return natGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM(generators)
	natGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(NatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM{}), generators)

	return natGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySpec_PublicIPAddress_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPAddressDnsSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressDnsSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressDnsSettings_ARM, PublicIPAddressDnsSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressDnsSettings_ARM runs a test to see if a specific instance of PublicIPAddressDnsSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressDnsSettings_ARM(subject PublicIPAddressDnsSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressDnsSettings_ARM
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

// Generator of PublicIPAddressDnsSettings_ARM instances for property testing - lazily instantiated by
// PublicIPAddressDnsSettings_ARMGenerator()
var publicIPAddressDnsSettings_ARMGenerator gopter.Gen

// PublicIPAddressDnsSettings_ARMGenerator returns a generator of PublicIPAddressDnsSettings_ARM instances for property testing.
func PublicIPAddressDnsSettings_ARMGenerator() gopter.Gen {
	if publicIPAddressDnsSettings_ARMGenerator != nil {
		return publicIPAddressDnsSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_ARM(generators)
	publicIPAddressDnsSettings_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressDnsSettings_ARM{}), generators)

	return publicIPAddressDnsSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettings_ARM(gens map[string]gopter.Gen) {
	gens["DomainNameLabel"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["ReverseFqdn"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM, PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM runs a test to see if a specific instance of PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM(subject PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM
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

// Generator of PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM instances for property testing - lazily
// instantiated by PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator()
var publicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator gopter.Gen

// PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator returns a generator of PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM instances for property testing.
func PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if publicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator != nil {
		return publicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM(generators)
	publicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM{}), generators)

	return publicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressSpec_PublicIPAddress_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
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
