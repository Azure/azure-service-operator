// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401preview

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

func Test_SearchService_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchService_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchService_Spec_ARM, SearchService_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchService_Spec_ARM runs a test to see if a specific instance of SearchService_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchService_Spec_ARM(subject SearchService_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchService_Spec_ARM
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

// Generator of SearchService_Spec_ARM instances for property testing - lazily instantiated by
// SearchService_Spec_ARMGenerator()
var searchService_Spec_ARMGenerator gopter.Gen

// SearchService_Spec_ARMGenerator returns a generator of SearchService_Spec_ARM instances for property testing.
// We first initialize searchService_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchService_Spec_ARMGenerator() gopter.Gen {
	if searchService_Spec_ARMGenerator != nil {
		return searchService_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_Spec_ARM(generators)
	searchService_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(SearchService_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForSearchService_Spec_ARM(generators)
	searchService_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(SearchService_Spec_ARM{}), generators)

	return searchService_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSearchService_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchService_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSearchService_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchService_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(SearchServiceProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_Identity_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_ARM, Identity_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_ARM runs a test to see if a specific instance of Identity_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_ARM(subject Identity_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_ARM
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

// Generator of Identity_ARM instances for property testing - lazily instantiated by Identity_ARMGenerator()
var identity_ARMGenerator gopter.Gen

// Identity_ARMGenerator returns a generator of Identity_ARM instances for property testing.
// We first initialize identity_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_ARMGenerator() gopter.Gen {
	if identity_ARMGenerator != nil {
		return identity_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_ARM(generators)
	identity_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_ARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_ARM(generators)
	identity_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_ARM{}), generators)

	return identity_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_None,
		Identity_Type_SystemAssigned,
		Identity_Type_SystemAssignedUserAssigned,
		Identity_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserAssignedIdentityDetails_ARMGenerator())
}

func Test_SearchServiceProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchServiceProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchServiceProperties_ARM, SearchServiceProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchServiceProperties_ARM runs a test to see if a specific instance of SearchServiceProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchServiceProperties_ARM(subject SearchServiceProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchServiceProperties_ARM
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

// Generator of SearchServiceProperties_ARM instances for property testing - lazily instantiated by
// SearchServiceProperties_ARMGenerator()
var searchServiceProperties_ARMGenerator gopter.Gen

// SearchServiceProperties_ARMGenerator returns a generator of SearchServiceProperties_ARM instances for property testing.
// We first initialize searchServiceProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchServiceProperties_ARMGenerator() gopter.Gen {
	if searchServiceProperties_ARMGenerator != nil {
		return searchServiceProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties_ARM(generators)
	searchServiceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForSearchServiceProperties_ARM(generators)
	searchServiceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties_ARM{}), generators)

	return searchServiceProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSearchServiceProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchServiceProperties_ARM(gens map[string]gopter.Gen) {
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["DisabledDataExfiltrationOptions"] = gen.SliceOf(gen.OneConstOf(DisabledDataExfiltrationOption_All))
	gens["HostingMode"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_HostingMode_Default, SearchServiceProperties_HostingMode_HighDensity))
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_PublicNetworkAccess_Disabled, SearchServiceProperties_PublicNetworkAccess_Enabled))
	gens["ReplicaCount"] = gen.PtrOf(gen.Int())
	gens["SemanticSearch"] = gen.PtrOf(gen.OneConstOf(SemanticSearch_Disabled, SemanticSearch_Free, SemanticSearch_Standard))
}

// AddRelatedPropertyGeneratorsForSearchServiceProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchServiceProperties_ARM(gens map[string]gopter.Gen) {
	gens["AuthOptions"] = gen.PtrOf(DataPlaneAuthOptions_ARMGenerator())
	gens["EncryptionWithCmk"] = gen.PtrOf(EncryptionWithCmk_ARMGenerator())
	gens["NetworkRuleSet"] = gen.PtrOf(NetworkRuleSet_ARMGenerator())
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
		Sku_Name_Free,
		Sku_Name_Standard,
		Sku_Name_Standard2,
		Sku_Name_Standard3,
		Sku_Name_Storage_Optimized_L1,
		Sku_Name_Storage_Optimized_L2))
}

func Test_DataPlaneAuthOptions_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAuthOptions_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAuthOptions_ARM, DataPlaneAuthOptions_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAuthOptions_ARM runs a test to see if a specific instance of DataPlaneAuthOptions_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAuthOptions_ARM(subject DataPlaneAuthOptions_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAuthOptions_ARM
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

// Generator of DataPlaneAuthOptions_ARM instances for property testing - lazily instantiated by
// DataPlaneAuthOptions_ARMGenerator()
var dataPlaneAuthOptions_ARMGenerator gopter.Gen

// DataPlaneAuthOptions_ARMGenerator returns a generator of DataPlaneAuthOptions_ARM instances for property testing.
func DataPlaneAuthOptions_ARMGenerator() gopter.Gen {
	if dataPlaneAuthOptions_ARMGenerator != nil {
		return dataPlaneAuthOptions_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_ARM(generators)
	dataPlaneAuthOptions_ARMGenerator = gen.Struct(reflect.TypeOf(DataPlaneAuthOptions_ARM{}), generators)

	return dataPlaneAuthOptions_ARMGenerator
}

// AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDataPlaneAuthOptions_ARM(gens map[string]gopter.Gen) {
	gens["AadOrApiKey"] = gen.PtrOf(DataPlaneAadOrApiKeyAuthOption_ARMGenerator())
}

func Test_EncryptionWithCmk_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionWithCmk_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionWithCmk_ARM, EncryptionWithCmk_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionWithCmk_ARM runs a test to see if a specific instance of EncryptionWithCmk_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionWithCmk_ARM(subject EncryptionWithCmk_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionWithCmk_ARM
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

// Generator of EncryptionWithCmk_ARM instances for property testing - lazily instantiated by
// EncryptionWithCmk_ARMGenerator()
var encryptionWithCmk_ARMGenerator gopter.Gen

// EncryptionWithCmk_ARMGenerator returns a generator of EncryptionWithCmk_ARM instances for property testing.
func EncryptionWithCmk_ARMGenerator() gopter.Gen {
	if encryptionWithCmk_ARMGenerator != nil {
		return encryptionWithCmk_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionWithCmk_ARM(generators)
	encryptionWithCmk_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionWithCmk_ARM{}), generators)

	return encryptionWithCmk_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionWithCmk_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionWithCmk_ARM(gens map[string]gopter.Gen) {
	gens["Enforcement"] = gen.PtrOf(gen.OneConstOf(EncryptionWithCmk_Enforcement_Disabled, EncryptionWithCmk_Enforcement_Enabled, EncryptionWithCmk_Enforcement_Unspecified))
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
	gens["Bypass"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_Bypass_AzurePortal, NetworkRuleSet_Bypass_None))
}

// AddRelatedPropertyGeneratorsForNetworkRuleSet_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSet_ARM(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IpRule_ARMGenerator())
}

func Test_UserAssignedIdentityDetails_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails_ARM, UserAssignedIdentityDetails_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails_ARM runs a test to see if a specific instance of UserAssignedIdentityDetails_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails_ARM(subject UserAssignedIdentityDetails_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails_ARM
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

// Generator of UserAssignedIdentityDetails_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityDetails_ARMGenerator()
var userAssignedIdentityDetails_ARMGenerator gopter.Gen

// UserAssignedIdentityDetails_ARMGenerator returns a generator of UserAssignedIdentityDetails_ARM instances for property testing.
func UserAssignedIdentityDetails_ARMGenerator() gopter.Gen {
	if userAssignedIdentityDetails_ARMGenerator != nil {
		return userAssignedIdentityDetails_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetails_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails_ARM{}), generators)

	return userAssignedIdentityDetails_ARMGenerator
}

func Test_DataPlaneAadOrApiKeyAuthOption_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAadOrApiKeyAuthOption_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_ARM, DataPlaneAadOrApiKeyAuthOption_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_ARM runs a test to see if a specific instance of DataPlaneAadOrApiKeyAuthOption_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption_ARM(subject DataPlaneAadOrApiKeyAuthOption_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAadOrApiKeyAuthOption_ARM
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

// Generator of DataPlaneAadOrApiKeyAuthOption_ARM instances for property testing - lazily instantiated by
// DataPlaneAadOrApiKeyAuthOption_ARMGenerator()
var dataPlaneAadOrApiKeyAuthOption_ARMGenerator gopter.Gen

// DataPlaneAadOrApiKeyAuthOption_ARMGenerator returns a generator of DataPlaneAadOrApiKeyAuthOption_ARM instances for property testing.
func DataPlaneAadOrApiKeyAuthOption_ARMGenerator() gopter.Gen {
	if dataPlaneAadOrApiKeyAuthOption_ARMGenerator != nil {
		return dataPlaneAadOrApiKeyAuthOption_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_ARM(generators)
	dataPlaneAadOrApiKeyAuthOption_ARMGenerator = gen.Struct(reflect.TypeOf(DataPlaneAadOrApiKeyAuthOption_ARM{}), generators)

	return dataPlaneAadOrApiKeyAuthOption_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption_ARM(gens map[string]gopter.Gen) {
	gens["AadAuthFailureMode"] = gen.PtrOf(gen.OneConstOf(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_Http401WithBearerChallenge, DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_Http403))
}

func Test_IpRule_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpRule_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpRule_ARM, IpRule_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpRule_ARM runs a test to see if a specific instance of IpRule_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpRule_ARM(subject IpRule_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpRule_ARM
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

// Generator of IpRule_ARM instances for property testing - lazily instantiated by IpRule_ARMGenerator()
var ipRule_ARMGenerator gopter.Gen

// IpRule_ARMGenerator returns a generator of IpRule_ARM instances for property testing.
func IpRule_ARMGenerator() gopter.Gen {
	if ipRule_ARMGenerator != nil {
		return ipRule_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpRule_ARM(generators)
	ipRule_ARMGenerator = gen.Struct(reflect.TypeOf(IpRule_ARM{}), generators)

	return ipRule_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpRule_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpRule_ARM(gens map[string]gopter.Gen) {
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
