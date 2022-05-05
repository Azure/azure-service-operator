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

func Test_PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM, PublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM runs a test to see if a specific instance of PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(subject PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM
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

// Generator of PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator()
var publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator gopter.Gen

// PublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator returns a generator of PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing.
// We first initialize publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator() gopter.Gen {
	if publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator != nil {
		return publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	return publicIPAddressStatusPublicIPAddressSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(PublicIPAddressPropertiesFormatStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPAddressSkuStatusARMGenerator())
}

func Test_PublicIPAddressPropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressPropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressPropertiesFormatStatusARM, PublicIPAddressPropertiesFormatStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressPropertiesFormatStatusARM runs a test to see if a specific instance of PublicIPAddressPropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressPropertiesFormatStatusARM(subject PublicIPAddressPropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressPropertiesFormat_StatusARM
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

// Generator of PublicIPAddressPropertiesFormat_StatusARM instances for property testing - lazily instantiated by
// PublicIPAddressPropertiesFormatStatusARMGenerator()
var publicIPAddressPropertiesFormatStatusARMGenerator gopter.Gen

// PublicIPAddressPropertiesFormatStatusARMGenerator returns a generator of PublicIPAddressPropertiesFormat_StatusARM instances for property testing.
// We first initialize publicIPAddressPropertiesFormatStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPAddressPropertiesFormatStatusARMGenerator() gopter.Gen {
	if publicIPAddressPropertiesFormatStatusARMGenerator != nil {
		return publicIPAddressPropertiesFormatStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM(generators)
	publicIPAddressPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM(generators)
	AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM(generators)
	publicIPAddressPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressPropertiesFormat_StatusARM{}), generators)

	return publicIPAddressPropertiesFormatStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["MigrationPhase"] = gen.PtrOf(gen.OneConstOf(
		PublicIPAddressPropertiesFormatStatusMigrationPhaseAbort,
		PublicIPAddressPropertiesFormatStatusMigrationPhaseCommit,
		PublicIPAddressPropertiesFormatStatusMigrationPhaseCommitted,
		PublicIPAddressPropertiesFormatStatusMigrationPhaseNone,
		PublicIPAddressPropertiesFormatStatusMigrationPhasePrepare))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_StatusIPv4, IPVersion_StatusIPv6))
	gens["PublicIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_StatusDynamic, IPAllocationMethod_StatusStatic))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPAddressPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["DdosSettings"] = gen.PtrOf(DdosSettingsStatusARMGenerator())
	gens["DnsSettings"] = gen.PtrOf(PublicIPAddressDnsSettingsStatusARMGenerator())
	gens["IpConfiguration"] = gen.PtrOf(IPConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator())
	gens["IpTags"] = gen.SliceOf(IpTagStatusARMGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator())
	gens["PublicIPPrefix"] = gen.PtrOf(SubResourceStatusARMGenerator())
}

func Test_PublicIPAddressSku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressSku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressSkuStatusARM, PublicIPAddressSkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressSkuStatusARM runs a test to see if a specific instance of PublicIPAddressSku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressSkuStatusARM(subject PublicIPAddressSku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressSku_StatusARM
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

// Generator of PublicIPAddressSku_StatusARM instances for property testing - lazily instantiated by
// PublicIPAddressSkuStatusARMGenerator()
var publicIPAddressSkuStatusARMGenerator gopter.Gen

// PublicIPAddressSkuStatusARMGenerator returns a generator of PublicIPAddressSku_StatusARM instances for property testing.
func PublicIPAddressSkuStatusARMGenerator() gopter.Gen {
	if publicIPAddressSkuStatusARMGenerator != nil {
		return publicIPAddressSkuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressSkuStatusARM(generators)
	publicIPAddressSkuStatusARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressSku_StatusARM{}), generators)

	return publicIPAddressSkuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSkuStatusNameBasic, PublicIPAddressSkuStatusNameStandard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(PublicIPAddressSkuStatusTierGlobal, PublicIPAddressSkuStatusTierRegional))
}

func Test_DdosSettings_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DdosSettings_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDdosSettingsStatusARM, DdosSettingsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDdosSettingsStatusARM runs a test to see if a specific instance of DdosSettings_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDdosSettingsStatusARM(subject DdosSettings_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DdosSettings_StatusARM
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

// Generator of DdosSettings_StatusARM instances for property testing - lazily instantiated by
// DdosSettingsStatusARMGenerator()
var ddosSettingsStatusARMGenerator gopter.Gen

// DdosSettingsStatusARMGenerator returns a generator of DdosSettings_StatusARM instances for property testing.
// We first initialize ddosSettingsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DdosSettingsStatusARMGenerator() gopter.Gen {
	if ddosSettingsStatusARMGenerator != nil {
		return ddosSettingsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettingsStatusARM(generators)
	ddosSettingsStatusARMGenerator = gen.Struct(reflect.TypeOf(DdosSettings_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDdosSettingsStatusARM(generators)
	AddRelatedPropertyGeneratorsForDdosSettingsStatusARM(generators)
	ddosSettingsStatusARMGenerator = gen.Struct(reflect.TypeOf(DdosSettings_StatusARM{}), generators)

	return ddosSettingsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDdosSettingsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDdosSettingsStatusARM(gens map[string]gopter.Gen) {
	gens["ProtectedIP"] = gen.PtrOf(gen.Bool())
	gens["ProtectionCoverage"] = gen.PtrOf(gen.OneConstOf(DdosSettingsStatusProtectionCoverageBasic, DdosSettingsStatusProtectionCoverageStandard))
}

// AddRelatedPropertyGeneratorsForDdosSettingsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDdosSettingsStatusARM(gens map[string]gopter.Gen) {
	gens["DdosCustomPolicy"] = gen.PtrOf(SubResourceStatusARMGenerator())
}

func Test_IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM, IPConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM runs a test to see if a specific instance of IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(subject IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM
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

// Generator of IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by IPConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator()
var ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator gopter.Gen

// IPConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator returns a generator of IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing.
// We first initialize ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IPConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator() gopter.Gen {
	if ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator != nil {
		return ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	return ipConfigurationStatusPublicIPAddressSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIPConfigurationStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(IPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator())
}

func Test_IpTag_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpTag_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpTagStatusARM, IpTagStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpTagStatusARM runs a test to see if a specific instance of IpTag_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpTagStatusARM(subject IpTag_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpTag_StatusARM
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

// Generator of IpTag_StatusARM instances for property testing - lazily instantiated by IpTagStatusARMGenerator()
var ipTagStatusARMGenerator gopter.Gen

// IpTagStatusARMGenerator returns a generator of IpTag_StatusARM instances for property testing.
func IpTagStatusARMGenerator() gopter.Gen {
	if ipTagStatusARMGenerator != nil {
		return ipTagStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpTagStatusARM(generators)
	ipTagStatusARMGenerator = gen.Struct(reflect.TypeOf(IpTag_StatusARM{}), generators)

	return ipTagStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIpTagStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpTagStatusARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM, NatGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM runs a test to see if a specific instance of NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(subject NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM
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

// Generator of NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by NatGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator()
var natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator gopter.Gen

// NatGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator returns a generator of NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing.
// We first initialize natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator() gopter.Gen {
	if natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator != nil {
		return natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	return natGatewayStatusPublicIPAddressSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGatewayStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(NatGatewaySkuStatusARMGenerator())
}

func Test_PublicIPAddressDnsSettings_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPAddressDnsSettings_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPAddressDnsSettingsStatusARM, PublicIPAddressDnsSettingsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPAddressDnsSettingsStatusARM runs a test to see if a specific instance of PublicIPAddressDnsSettings_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPAddressDnsSettingsStatusARM(subject PublicIPAddressDnsSettings_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPAddressDnsSettings_StatusARM
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

// Generator of PublicIPAddressDnsSettings_StatusARM instances for property testing - lazily instantiated by
// PublicIPAddressDnsSettingsStatusARMGenerator()
var publicIPAddressDnsSettingsStatusARMGenerator gopter.Gen

// PublicIPAddressDnsSettingsStatusARMGenerator returns a generator of PublicIPAddressDnsSettings_StatusARM instances for property testing.
func PublicIPAddressDnsSettingsStatusARMGenerator() gopter.Gen {
	if publicIPAddressDnsSettingsStatusARMGenerator != nil {
		return publicIPAddressDnsSettingsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsStatusARM(generators)
	publicIPAddressDnsSettingsStatusARMGenerator = gen.Struct(reflect.TypeOf(PublicIPAddressDnsSettings_StatusARM{}), generators)

	return publicIPAddressDnsSettingsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPAddressDnsSettingsStatusARM(gens map[string]gopter.Gen) {
	gens["DomainNameLabel"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["ReverseFqdn"] = gen.PtrOf(gen.AlphaString())
}

func Test_IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM, IPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM runs a test to see if a specific instance of IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(subject IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM
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

// Generator of IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property
// testing - lazily instantiated by IPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator()
var ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator gopter.Gen

// IPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator returns a generator of IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing.
// We first initialize ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator() gopter.Gen {
	if ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator != nil {
		return ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	return ipConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_StatusDynamic, IPAllocationMethod_StatusStatic))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
}

// AddRelatedPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIPConfigurationPropertiesFormatStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(SubnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator())
}

func Test_NatGatewaySku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySkuStatusARM, NatGatewaySkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySkuStatusARM runs a test to see if a specific instance of NatGatewaySku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySkuStatusARM(subject NatGatewaySku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySku_StatusARM
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

// Generator of NatGatewaySku_StatusARM instances for property testing - lazily instantiated by
// NatGatewaySkuStatusARMGenerator()
var natGatewaySkuStatusARMGenerator gopter.Gen

// NatGatewaySkuStatusARMGenerator returns a generator of NatGatewaySku_StatusARM instances for property testing.
func NatGatewaySkuStatusARMGenerator() gopter.Gen {
	if natGatewaySkuStatusARMGenerator != nil {
		return natGatewaySkuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySkuStatusARM(generators)
	natGatewaySkuStatusARMGenerator = gen.Struct(reflect.TypeOf(NatGatewaySku_StatusARM{}), generators)

	return natGatewaySkuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySkuStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(NatGatewaySkuStatusNameStandard))
}

func Test_Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetStatusPublicIPAddressSubResourceEmbeddedARM, SubnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetStatusPublicIPAddressSubResourceEmbeddedARM runs a test to see if a specific instance of Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetStatusPublicIPAddressSubResourceEmbeddedARM(subject Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM
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

// Generator of Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by SubnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator()
var subnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator gopter.Gen

// SubnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator returns a generator of Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM instances for property testing.
func SubnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator() gopter.Gen {
	if subnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator != nil {
		return subnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetStatusPublicIPAddressSubResourceEmbeddedARM(generators)
	subnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM{}), generators)

	return subnetStatusPublicIPAddressSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnetStatusPublicIPAddressSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetStatusPublicIPAddressSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
