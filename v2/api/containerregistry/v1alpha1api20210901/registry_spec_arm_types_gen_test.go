// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

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

func Test_Registry_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Registry_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRegistry_Spec_ARM, Registry_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRegistry_Spec_ARM runs a test to see if a specific instance of Registry_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRegistry_Spec_ARM(subject Registry_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Registry_Spec_ARM
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

// Generator of Registry_Spec_ARM instances for property testing - lazily instantiated by Registry_Spec_ARMGenerator()
var registry_Spec_ARMGenerator gopter.Gen

// Registry_Spec_ARMGenerator returns a generator of Registry_Spec_ARM instances for property testing.
// We first initialize registry_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Registry_Spec_ARMGenerator() gopter.Gen {
	if registry_Spec_ARMGenerator != nil {
		return registry_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistry_Spec_ARM(generators)
	registry_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Registry_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistry_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRegistry_Spec_ARM(generators)
	registry_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Registry_Spec_ARM{}), generators)

	return registry_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRegistry_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRegistry_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRegistry_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRegistry_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityProperties_ARMGenerator())
	gens["Properties"] = gen.PtrOf(RegistryProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_IdentityProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityProperties_ARM, IdentityProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityProperties_ARM runs a test to see if a specific instance of IdentityProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityProperties_ARM(subject IdentityProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityProperties_ARM
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

// Generator of IdentityProperties_ARM instances for property testing - lazily instantiated by
// IdentityProperties_ARMGenerator()
var identityProperties_ARMGenerator gopter.Gen

// IdentityProperties_ARMGenerator returns a generator of IdentityProperties_ARM instances for property testing.
// We first initialize identityProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityProperties_ARMGenerator() gopter.Gen {
	if identityProperties_ARMGenerator != nil {
		return identityProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityProperties_ARM(generators)
	identityProperties_ARMGenerator = gen.Struct(reflect.TypeOf(IdentityProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForIdentityProperties_ARM(generators)
	identityProperties_ARMGenerator = gen.Struct(reflect.TypeOf(IdentityProperties_ARM{}), generators)

	return identityProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityProperties_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		IdentityProperties_Type_None,
		IdentityProperties_Type_SystemAssigned,
		IdentityProperties_Type_SystemAssignedUserAssigned,
		IdentityProperties_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentityProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentityProperties_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserIdentityProperties_ARMGenerator())
}

func Test_RegistryProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RegistryProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRegistryProperties_ARM, RegistryProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRegistryProperties_ARM runs a test to see if a specific instance of RegistryProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRegistryProperties_ARM(subject RegistryProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RegistryProperties_ARM
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

// Generator of RegistryProperties_ARM instances for property testing - lazily instantiated by
// RegistryProperties_ARMGenerator()
var registryProperties_ARMGenerator gopter.Gen

// RegistryProperties_ARMGenerator returns a generator of RegistryProperties_ARM instances for property testing.
// We first initialize registryProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RegistryProperties_ARMGenerator() gopter.Gen {
	if registryProperties_ARMGenerator != nil {
		return registryProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistryProperties_ARM(generators)
	registryProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RegistryProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRegistryProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForRegistryProperties_ARM(generators)
	registryProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RegistryProperties_ARM{}), generators)

	return registryProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRegistryProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRegistryProperties_ARM(gens map[string]gopter.Gen) {
	gens["AdminUserEnabled"] = gen.PtrOf(gen.Bool())
	gens["DataEndpointEnabled"] = gen.PtrOf(gen.Bool())
	gens["NetworkRuleBypassOptions"] = gen.PtrOf(gen.OneConstOf(RegistryProperties_NetworkRuleBypassOptions_AzureServices, RegistryProperties_NetworkRuleBypassOptions_None))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RegistryProperties_PublicNetworkAccess_Disabled, RegistryProperties_PublicNetworkAccess_Enabled))
	gens["ZoneRedundancy"] = gen.PtrOf(gen.OneConstOf(RegistryProperties_ZoneRedundancy_Disabled, RegistryProperties_ZoneRedundancy_Enabled))
}

// AddRelatedPropertyGeneratorsForRegistryProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRegistryProperties_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionProperty_ARMGenerator())
	gens["NetworkRuleSet"] = gen.PtrOf(NetworkRuleSet_ARMGenerator())
	gens["Policies"] = gen.PtrOf(Policies_ARMGenerator())
}

func Test_Sku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_ARM, Sku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_ARM runs a test to see if a specific instance of Sku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_ARM(subject Sku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_ARM
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

// Generator of Sku_ARM instances for property testing - lazily instantiated by Sku_ARMGenerator()
var sku_ARMGenerator gopter.Gen

// Sku_ARMGenerator returns a generator of Sku_ARM instances for property testing.
func Sku_ARMGenerator() gopter.Gen {
	if sku_ARMGenerator != nil {
		return sku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_ARM(generators)
	sku_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_ARM{}), generators)

	return sku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_Basic,
		Sku_Name_Classic,
		Sku_Name_Premium,
		Sku_Name_Standard))
}

func Test_EncryptionProperty_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperty_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperty_ARM, EncryptionProperty_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperty_ARM runs a test to see if a specific instance of EncryptionProperty_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperty_ARM(subject EncryptionProperty_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperty_ARM
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

// Generator of EncryptionProperty_ARM instances for property testing - lazily instantiated by
// EncryptionProperty_ARMGenerator()
var encryptionProperty_ARMGenerator gopter.Gen

// EncryptionProperty_ARMGenerator returns a generator of EncryptionProperty_ARM instances for property testing.
// We first initialize encryptionProperty_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionProperty_ARMGenerator() gopter.Gen {
	if encryptionProperty_ARMGenerator != nil {
		return encryptionProperty_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperty_ARM(generators)
	encryptionProperty_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperty_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperty_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionProperty_ARM(generators)
	encryptionProperty_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperty_ARM{}), generators)

	return encryptionProperty_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionProperty_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionProperty_ARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(EncryptionProperty_Status_Disabled, EncryptionProperty_Status_Enabled))
}

// AddRelatedPropertyGeneratorsForEncryptionProperty_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperty_ARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_ARMGenerator())
}

func Test_NetworkRuleSet_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSet_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSet_ARM, NetworkRuleSet_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSet_ARM runs a test to see if a specific instance of NetworkRuleSet_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSet_ARM(subject NetworkRuleSet_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSet_ARM
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

// Generator of NetworkRuleSet_ARM instances for property testing - lazily instantiated by NetworkRuleSet_ARMGenerator()
var networkRuleSet_ARMGenerator gopter.Gen

// NetworkRuleSet_ARMGenerator returns a generator of NetworkRuleSet_ARM instances for property testing.
// We first initialize networkRuleSet_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkRuleSet_ARMGenerator() gopter.Gen {
	if networkRuleSet_ARMGenerator != nil {
		return networkRuleSet_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSet_ARM(generators)
	networkRuleSet_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSet_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkRuleSet_ARM(generators)
	networkRuleSet_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet_ARM{}), generators)

	return networkRuleSet_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkRuleSet_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkRuleSet_ARM(gens map[string]gopter.Gen) {
	gens["DefaultAction"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_DefaultAction_Allow, NetworkRuleSet_DefaultAction_Deny))
}

// AddRelatedPropertyGeneratorsForNetworkRuleSet_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSet_ARM(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IPRule_ARMGenerator())
}

func Test_Policies_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Policies_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicies_ARM, Policies_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicies_ARM runs a test to see if a specific instance of Policies_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicies_ARM(subject Policies_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Policies_ARM
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

// Generator of Policies_ARM instances for property testing - lazily instantiated by Policies_ARMGenerator()
var policies_ARMGenerator gopter.Gen

// Policies_ARMGenerator returns a generator of Policies_ARM instances for property testing.
func Policies_ARMGenerator() gopter.Gen {
	if policies_ARMGenerator != nil {
		return policies_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPolicies_ARM(generators)
	policies_ARMGenerator = gen.Struct(reflect.TypeOf(Policies_ARM{}), generators)

	return policies_ARMGenerator
}

// AddRelatedPropertyGeneratorsForPolicies_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicies_ARM(gens map[string]gopter.Gen) {
	gens["ExportPolicy"] = gen.PtrOf(ExportPolicy_ARMGenerator())
	gens["QuarantinePolicy"] = gen.PtrOf(QuarantinePolicy_ARMGenerator())
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicy_ARMGenerator())
	gens["TrustPolicy"] = gen.PtrOf(TrustPolicy_ARMGenerator())
}

func Test_UserIdentityProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserIdentityProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserIdentityProperties_ARM, UserIdentityProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserIdentityProperties_ARM runs a test to see if a specific instance of UserIdentityProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserIdentityProperties_ARM(subject UserIdentityProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserIdentityProperties_ARM
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

// Generator of UserIdentityProperties_ARM instances for property testing - lazily instantiated by
// UserIdentityProperties_ARMGenerator()
var userIdentityProperties_ARMGenerator gopter.Gen

// UserIdentityProperties_ARMGenerator returns a generator of UserIdentityProperties_ARM instances for property testing.
func UserIdentityProperties_ARMGenerator() gopter.Gen {
	if userIdentityProperties_ARMGenerator != nil {
		return userIdentityProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserIdentityProperties_ARM(generators)
	userIdentityProperties_ARMGenerator = gen.Struct(reflect.TypeOf(UserIdentityProperties_ARM{}), generators)

	return userIdentityProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserIdentityProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserIdentityProperties_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExportPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExportPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExportPolicy_ARM, ExportPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExportPolicy_ARM runs a test to see if a specific instance of ExportPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExportPolicy_ARM(subject ExportPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExportPolicy_ARM
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

// Generator of ExportPolicy_ARM instances for property testing - lazily instantiated by ExportPolicy_ARMGenerator()
var exportPolicy_ARMGenerator gopter.Gen

// ExportPolicy_ARMGenerator returns a generator of ExportPolicy_ARM instances for property testing.
func ExportPolicy_ARMGenerator() gopter.Gen {
	if exportPolicy_ARMGenerator != nil {
		return exportPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExportPolicy_ARM(generators)
	exportPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(ExportPolicy_ARM{}), generators)

	return exportPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExportPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExportPolicy_ARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(ExportPolicy_Status_Disabled, ExportPolicy_Status_Enabled))
}

func Test_IPRule_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPRule_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPRule_ARM, IPRule_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPRule_ARM runs a test to see if a specific instance of IPRule_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPRule_ARM(subject IPRule_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPRule_ARM
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

// Generator of IPRule_ARM instances for property testing - lazily instantiated by IPRule_ARMGenerator()
var ipRule_ARMGenerator gopter.Gen

// IPRule_ARMGenerator returns a generator of IPRule_ARM instances for property testing.
func IPRule_ARMGenerator() gopter.Gen {
	if ipRule_ARMGenerator != nil {
		return ipRule_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPRule_ARM(generators)
	ipRule_ARMGenerator = gen.Struct(reflect.TypeOf(IPRule_ARM{}), generators)

	return ipRule_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIPRule_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPRule_ARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(IPRule_Action_Allow))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_ARM, KeyVaultProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_ARM runs a test to see if a specific instance of KeyVaultProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_ARM(subject KeyVaultProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_ARM
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

// Generator of KeyVaultProperties_ARM instances for property testing - lazily instantiated by
// KeyVaultProperties_ARMGenerator()
var keyVaultProperties_ARMGenerator gopter.Gen

// KeyVaultProperties_ARMGenerator returns a generator of KeyVaultProperties_ARM instances for property testing.
func KeyVaultProperties_ARMGenerator() gopter.Gen {
	if keyVaultProperties_ARMGenerator != nil {
		return keyVaultProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM(generators)
	keyVaultProperties_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_ARM{}), generators)

	return keyVaultProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}

func Test_QuarantinePolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QuarantinePolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQuarantinePolicy_ARM, QuarantinePolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQuarantinePolicy_ARM runs a test to see if a specific instance of QuarantinePolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQuarantinePolicy_ARM(subject QuarantinePolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QuarantinePolicy_ARM
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

// Generator of QuarantinePolicy_ARM instances for property testing - lazily instantiated by
// QuarantinePolicy_ARMGenerator()
var quarantinePolicy_ARMGenerator gopter.Gen

// QuarantinePolicy_ARMGenerator returns a generator of QuarantinePolicy_ARM instances for property testing.
func QuarantinePolicy_ARMGenerator() gopter.Gen {
	if quarantinePolicy_ARMGenerator != nil {
		return quarantinePolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQuarantinePolicy_ARM(generators)
	quarantinePolicy_ARMGenerator = gen.Struct(reflect.TypeOf(QuarantinePolicy_ARM{}), generators)

	return quarantinePolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForQuarantinePolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQuarantinePolicy_ARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(QuarantinePolicy_Status_Disabled, QuarantinePolicy_Status_Enabled))
}

func Test_RetentionPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetentionPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetentionPolicy_ARM, RetentionPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetentionPolicy_ARM runs a test to see if a specific instance of RetentionPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRetentionPolicy_ARM(subject RetentionPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetentionPolicy_ARM
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

// Generator of RetentionPolicy_ARM instances for property testing - lazily instantiated by
// RetentionPolicy_ARMGenerator()
var retentionPolicy_ARMGenerator gopter.Gen

// RetentionPolicy_ARMGenerator returns a generator of RetentionPolicy_ARM instances for property testing.
func RetentionPolicy_ARMGenerator() gopter.Gen {
	if retentionPolicy_ARMGenerator != nil {
		return retentionPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetentionPolicy_ARM(generators)
	retentionPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(RetentionPolicy_ARM{}), generators)

	return retentionPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRetentionPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetentionPolicy_ARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(RetentionPolicy_Status_Disabled, RetentionPolicy_Status_Enabled))
}

func Test_TrustPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustPolicy_ARM, TrustPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustPolicy_ARM runs a test to see if a specific instance of TrustPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustPolicy_ARM(subject TrustPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustPolicy_ARM
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

// Generator of TrustPolicy_ARM instances for property testing - lazily instantiated by TrustPolicy_ARMGenerator()
var trustPolicy_ARMGenerator gopter.Gen

// TrustPolicy_ARMGenerator returns a generator of TrustPolicy_ARM instances for property testing.
func TrustPolicy_ARMGenerator() gopter.Gen {
	if trustPolicy_ARMGenerator != nil {
		return trustPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustPolicy_ARM(generators)
	trustPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(TrustPolicy_ARM{}), generators)

	return trustPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTrustPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustPolicy_ARM(gens map[string]gopter.Gen) {
	gens["Status"] = gen.PtrOf(gen.OneConstOf(TrustPolicy_Status_Disabled, TrustPolicy_Status_Enabled))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(TrustPolicy_Type_Notary))
}
