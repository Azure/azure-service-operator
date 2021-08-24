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

func Test_PublicIPAddresses_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddresses_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressesSpecARM, PublicIPAddressesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressesSpecARM runs a test to see if a specific instance of PublicIPAddresses_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressesSpecARM(subject PublicIPAddresses_SpecARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual PublicIPAddresses_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of PublicIPAddresses_SpecARM instances for property testing - lazily
//instantiated by PublicIPAddressesSpecARMGenerator()
var publicIPAddressesSpecARMGenerator gopter.Gen

// PublicIPAddressesSpecARMGenerator returns a generator of PublicIPAddresses_SpecARM instances for property testing.
// We first initialize publicIPAddressesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressesSpecARMGenerator() gopter.Gen {
	if publicIPAddressesSpecARMGenerator != nil {
		return publicIPAddressesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressesSpecARM(generators)
	publicIPAddressesSpecARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddresses_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressesSpecARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressesSpecARM(generators)
	publicIPAddressesSpecARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddresses_SpecARM{}), generators)

	return publicIPAddressesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressesSpecARM(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.OneConstOf(PublicIPAddressesSpecAPIVersion20201101)
	gens["Location"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.OneConstOf(PublicIPAddressesSpecTypeMicrosoftNetworkPublicIPAddresses)
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddressesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressesSpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = PublicIPAddressPropertiesFormatARMGenerator()
	gens["Sku"] = gen.PtrOf(PublicIPAddressSkuARMGenerator())
}

func Test_PublicIPAddressPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressPropertiesFormatARM, PublicIPAddressPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressPropertiesFormatARM runs a test to see if a specific instance of PublicIPAddressPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressPropertiesFormatARM(subject PublicIPAddressPropertiesFormatARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual PublicIPAddressPropertiesFormatARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of PublicIPAddressPropertiesFormatARM instances for property testing -
//lazily instantiated by PublicIPAddressPropertiesFormatARMGenerator()
var publicIPAddressPropertiesFormatARMGenerator gopter.Gen

// PublicIPAddressPropertiesFormatARMGenerator returns a generator of PublicIPAddressPropertiesFormatARM instances for property testing.
// We first initialize publicIPAddressPropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressPropertiesFormatARMGenerator() gopter.Gen {
	if publicIPAddressPropertiesFormatARMGenerator != nil {
		return publicIPAddressPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatARM(generators)
	publicIPAddressPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatARM(generators)
	publicIPAddressPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormatARM{}), generators)

	return publicIPAddressPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressPropertiesFormatPublicIPAddressVersionIPv4, PublicIPAddressPropertiesFormatPublicIPAddressVersionIPv6))
	gens["PublicIPAllocationMethod"] = gen.OneConstOf(PublicIPAddressPropertiesFormatPublicIPAllocationMethodDynamic, PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic)
}

// AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["DdosSettings"] = gen.PtrOf(DdosSettingsARMGenerator())
	gens["DnsSettings"] = gen.PtrOf(PublicIPAddressDnsSettingsARMGenerator())
	gens["IpTags"] = gen.SliceOf(IpTagARMGenerator())
	gens["PublicIPPrefix"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_PublicIPAddressSkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressSkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressSkuARM, PublicIPAddressSkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressSkuARM runs a test to see if a specific instance of PublicIPAddressSkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressSkuARM(subject PublicIPAddressSkuARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual PublicIPAddressSkuARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of PublicIPAddressSkuARM instances for property testing - lazily
//instantiated by PublicIPAddressSkuARMGenerator()
var publicIPAddressSkuARMGenerator gopter.Gen

// PublicIPAddressSkuARMGenerator returns a generator of PublicIPAddressSkuARM instances for property testing.
func PublicIPAddressSkuARMGenerator() gopter.Gen {
	if publicIPAddressSkuARMGenerator != nil {
		return publicIPAddressSkuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressSkuARM(generators)
	publicIPAddressSkuARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressSkuARM{}), generators)

	return publicIPAddressSkuARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressSkuARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSkuNameBasic, PublicIPAddressSkuNameStandard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSkuTierGlobal, PublicIPAddressSkuTierRegional))
}

func Test_DdosSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DdosSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDdosSettingsARM, DdosSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDdosSettingsARM runs a test to see if a specific instance of DdosSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDdosSettingsARM(subject DdosSettingsARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual DdosSettingsARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of DdosSettingsARM instances for property testing - lazily
//instantiated by DdosSettingsARMGenerator()
var ddosSettingsARMGenerator gopter.Gen

// DdosSettingsARMGenerator returns a generator of DdosSettingsARM instances for property testing.
// We first initialize ddosSettingsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DdosSettingsARMGenerator() gopter.Gen {
	if ddosSettingsARMGenerator != nil {
		return ddosSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettingsARM(generators)
	ddosSettingsARMGenerator = gen.Struct(reflect.TypeOf(DdosSettingsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettingsARM(generators)
	AddRelatedPropertyGeneratorsForDdosSettingsARM(generators)
	ddosSettingsARMGenerator = gen.Struct(reflect.TypeOf(DdosSettingsARM{}), generators)

	return ddosSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForDdosSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDdosSettingsARM(gens map[string]gopter.Gen) {
	gens["ProtectedIP"] = gen.PtrOf(gen.Bool())
	gens["ProtectionCoverage"] = gen.PtrOf(gen.OneConstOf(DdosSettingsProtectionCoverageBasic, DdosSettingsProtectionCoverageStandard))
}

// AddRelatedPropertyGeneratorsForDdosSettingsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDdosSettingsARM(gens map[string]gopter.Gen) {
	gens["DdosCustomPolicy"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_IpTagARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpTagARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpTagARM, IpTagARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpTagARM runs a test to see if a specific instance of IpTagARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpTagARM(subject IpTagARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual IpTagARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of IpTagARM instances for property testing - lazily instantiated by
//IpTagARMGenerator()
var ipTagARMGenerator gopter.Gen

// IpTagARMGenerator returns a generator of IpTagARM instances for property testing.
func IpTagARMGenerator() gopter.Gen {
	if ipTagARMGenerator != nil {
		return ipTagARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpTagARM(generators)
	ipTagARMGenerator = gen.Struct(reflect.TypeOf(IpTagARM{}), generators)

	return ipTagARMGenerator
}

// AddIndependentPropertyGeneratorsForIpTagARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpTagARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPAddressDnsSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressDnsSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressDnsSettingsARM, PublicIPAddressDnsSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressDnsSettingsARM runs a test to see if a specific instance of PublicIPAddressDnsSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressDnsSettingsARM(subject PublicIPAddressDnsSettingsARM) string {
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	var actual PublicIPAddressDnsSettingsARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

//Generator of PublicIPAddressDnsSettingsARM instances for property testing -
//lazily instantiated by PublicIPAddressDnsSettingsARMGenerator()
var publicIPAddressDnsSettingsARMGenerator gopter.Gen

// PublicIPAddressDnsSettingsARMGenerator returns a generator of PublicIPAddressDnsSettingsARM instances for property testing.
func PublicIPAddressDnsSettingsARMGenerator() gopter.Gen {
	if publicIPAddressDnsSettingsARMGenerator != nil {
		return publicIPAddressDnsSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsARM(generators)
	publicIPAddressDnsSettingsARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressDnsSettingsARM{}), generators)

	return publicIPAddressDnsSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsARM(gens map[string]gopter.Gen) {
	gens["DomainNameLabel"] = gen.AlphaString()
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["ReverseFqdn"] = gen.PtrOf(gen.AlphaString())
}
