// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_DataPlaneAadOrApiKeyAuthOption_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAadOrApiKeyAuthOption via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption, DataPlaneAadOrApiKeyAuthOptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption runs a test to see if a specific instance of DataPlaneAadOrApiKeyAuthOption round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAadOrApiKeyAuthOption(subject DataPlaneAadOrApiKeyAuthOption) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAadOrApiKeyAuthOption
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

// Generator of DataPlaneAadOrApiKeyAuthOption instances for property testing - lazily instantiated by
// DataPlaneAadOrApiKeyAuthOptionGenerator()
var dataPlaneAadOrApiKeyAuthOptionGenerator gopter.Gen

// DataPlaneAadOrApiKeyAuthOptionGenerator returns a generator of DataPlaneAadOrApiKeyAuthOption instances for property testing.
func DataPlaneAadOrApiKeyAuthOptionGenerator() gopter.Gen {
	if dataPlaneAadOrApiKeyAuthOptionGenerator != nil {
		return dataPlaneAadOrApiKeyAuthOptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption(generators)
	dataPlaneAadOrApiKeyAuthOptionGenerator = gen.Struct(reflect.TypeOf(DataPlaneAadOrApiKeyAuthOption{}), generators)

	return dataPlaneAadOrApiKeyAuthOptionGenerator
}

// AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataPlaneAadOrApiKeyAuthOption(gens map[string]gopter.Gen) {
	gens["AadAuthFailureMode"] = gen.PtrOf(gen.OneConstOf(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_Http401WithBearerChallenge, DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_Http403))
}

func Test_DataPlaneAuthOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataPlaneAuthOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataPlaneAuthOptions, DataPlaneAuthOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataPlaneAuthOptions runs a test to see if a specific instance of DataPlaneAuthOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForDataPlaneAuthOptions(subject DataPlaneAuthOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataPlaneAuthOptions
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

// Generator of DataPlaneAuthOptions instances for property testing - lazily instantiated by
// DataPlaneAuthOptionsGenerator()
var dataPlaneAuthOptionsGenerator gopter.Gen

// DataPlaneAuthOptionsGenerator returns a generator of DataPlaneAuthOptions instances for property testing.
func DataPlaneAuthOptionsGenerator() gopter.Gen {
	if dataPlaneAuthOptionsGenerator != nil {
		return dataPlaneAuthOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDataPlaneAuthOptions(generators)
	dataPlaneAuthOptionsGenerator = gen.Struct(reflect.TypeOf(DataPlaneAuthOptions{}), generators)

	return dataPlaneAuthOptionsGenerator
}

// AddRelatedPropertyGeneratorsForDataPlaneAuthOptions is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDataPlaneAuthOptions(gens map[string]gopter.Gen) {
	gens["AadOrApiKey"] = gen.PtrOf(DataPlaneAadOrApiKeyAuthOptionGenerator())
}

func Test_EncryptionWithCmk_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionWithCmk via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionWithCmk, EncryptionWithCmkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionWithCmk runs a test to see if a specific instance of EncryptionWithCmk round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionWithCmk(subject EncryptionWithCmk) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionWithCmk
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

// Generator of EncryptionWithCmk instances for property testing - lazily instantiated by EncryptionWithCmkGenerator()
var encryptionWithCmkGenerator gopter.Gen

// EncryptionWithCmkGenerator returns a generator of EncryptionWithCmk instances for property testing.
func EncryptionWithCmkGenerator() gopter.Gen {
	if encryptionWithCmkGenerator != nil {
		return encryptionWithCmkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionWithCmk(generators)
	encryptionWithCmkGenerator = gen.Struct(reflect.TypeOf(EncryptionWithCmk{}), generators)

	return encryptionWithCmkGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionWithCmk is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionWithCmk(gens map[string]gopter.Gen) {
	gens["Enforcement"] = gen.PtrOf(gen.OneConstOf(EncryptionWithCmk_Enforcement_Disabled, EncryptionWithCmk_Enforcement_Enabled, EncryptionWithCmk_Enforcement_Unspecified))
}

func Test_Identity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity, IdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity runs a test to see if a specific instance of Identity round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity(subject Identity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity
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

// Generator of Identity instances for property testing - lazily instantiated by IdentityGenerator()
var identityGenerator gopter.Gen

// IdentityGenerator returns a generator of Identity instances for property testing.
func IdentityGenerator() gopter.Gen {
	if identityGenerator != nil {
		return identityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity(generators)
	identityGenerator = gen.Struct(reflect.TypeOf(Identity{}), generators)

	return identityGenerator
}

// AddIndependentPropertyGeneratorsForIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_Type_None, Identity_Type_SystemAssigned))
}

func Test_IpRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpRule, IpRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpRule runs a test to see if a specific instance of IpRule round trips to JSON and back losslessly
func RunJSONSerializationTestForIpRule(subject IpRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpRule
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

// Generator of IpRule instances for property testing - lazily instantiated by IpRuleGenerator()
var ipRuleGenerator gopter.Gen

// IpRuleGenerator returns a generator of IpRule instances for property testing.
func IpRuleGenerator() gopter.Gen {
	if ipRuleGenerator != nil {
		return ipRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpRule(generators)
	ipRuleGenerator = gen.Struct(reflect.TypeOf(IpRule{}), generators)

	return ipRuleGenerator
}

// AddIndependentPropertyGeneratorsForIpRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpRule(gens map[string]gopter.Gen) {
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkRuleSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSet, NetworkRuleSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSet runs a test to see if a specific instance of NetworkRuleSet round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSet(subject NetworkRuleSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSet
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

// Generator of NetworkRuleSet instances for property testing - lazily instantiated by NetworkRuleSetGenerator()
var networkRuleSetGenerator gopter.Gen

// NetworkRuleSetGenerator returns a generator of NetworkRuleSet instances for property testing.
func NetworkRuleSetGenerator() gopter.Gen {
	if networkRuleSetGenerator != nil {
		return networkRuleSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkRuleSet(generators)
	networkRuleSetGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSet{}), generators)

	return networkRuleSetGenerator
}

// AddRelatedPropertyGeneratorsForNetworkRuleSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSet(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IpRuleGenerator())
}

func Test_SearchServiceProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchServiceProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchServiceProperties, SearchServicePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchServiceProperties runs a test to see if a specific instance of SearchServiceProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchServiceProperties(subject SearchServiceProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchServiceProperties
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

// Generator of SearchServiceProperties instances for property testing - lazily instantiated by
// SearchServicePropertiesGenerator()
var searchServicePropertiesGenerator gopter.Gen

// SearchServicePropertiesGenerator returns a generator of SearchServiceProperties instances for property testing.
// We first initialize searchServicePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchServicePropertiesGenerator() gopter.Gen {
	if searchServicePropertiesGenerator != nil {
		return searchServicePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties(generators)
	searchServicePropertiesGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchServiceProperties(generators)
	AddRelatedPropertyGeneratorsForSearchServiceProperties(generators)
	searchServicePropertiesGenerator = gen.Struct(reflect.TypeOf(SearchServiceProperties{}), generators)

	return searchServicePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSearchServiceProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchServiceProperties(gens map[string]gopter.Gen) {
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["HostingMode"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_HostingMode_Default, SearchServiceProperties_HostingMode_HighDensity))
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(SearchServiceProperties_PublicNetworkAccess_Disabled, SearchServiceProperties_PublicNetworkAccess_Enabled))
	gens["ReplicaCount"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForSearchServiceProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchServiceProperties(gens map[string]gopter.Gen) {
	gens["AuthOptions"] = gen.PtrOf(DataPlaneAuthOptionsGenerator())
	gens["EncryptionWithCmk"] = gen.PtrOf(EncryptionWithCmkGenerator())
	gens["NetworkRuleSet"] = gen.PtrOf(NetworkRuleSetGenerator())
}

func Test_SearchService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SearchService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSearchService_Spec, SearchService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSearchService_Spec runs a test to see if a specific instance of SearchService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSearchService_Spec(subject SearchService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SearchService_Spec
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

// Generator of SearchService_Spec instances for property testing - lazily instantiated by SearchService_SpecGenerator()
var searchService_SpecGenerator gopter.Gen

// SearchService_SpecGenerator returns a generator of SearchService_Spec instances for property testing.
// We first initialize searchService_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SearchService_SpecGenerator() gopter.Gen {
	if searchService_SpecGenerator != nil {
		return searchService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_Spec(generators)
	searchService_SpecGenerator = gen.Struct(reflect.TypeOf(SearchService_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSearchService_Spec(generators)
	AddRelatedPropertyGeneratorsForSearchService_Spec(generators)
	searchService_SpecGenerator = gen.Struct(reflect.TypeOf(SearchService_Spec{}), generators)

	return searchService_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSearchService_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSearchService_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSearchService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSearchService_Spec(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityGenerator())
	gens["Properties"] = gen.PtrOf(SearchServicePropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_Basic,
		Sku_Name_Free,
		Sku_Name_Standard,
		Sku_Name_Standard2,
		Sku_Name_Standard3,
		Sku_Name_Storage_Optimized_L1,
		Sku_Name_Storage_Optimized_L2))
}
