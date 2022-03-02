// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

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

func Test_Domains_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domains_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsSpecARM, DomainsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsSpecARM runs a test to see if a specific instance of Domains_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsSpecARM(subject Domains_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domains_SpecARM
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

// Generator of Domains_SpecARM instances for property testing - lazily instantiated by DomainsSpecARMGenerator()
var domainsSpecARMGenerator gopter.Gen

// DomainsSpecARMGenerator returns a generator of Domains_SpecARM instances for property testing.
// We first initialize domainsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsSpecARMGenerator() gopter.Gen {
	if domainsSpecARMGenerator != nil {
		return domainsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsSpecARM(generators)
	domainsSpecARMGenerator = gen.Struct(reflect.TypeOf(Domains_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDomainsSpecARM(generators)
	domainsSpecARMGenerator = gen.Struct(reflect.TypeOf(Domains_SpecARM{}), generators)

	return domainsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainPropertiesARMGenerator())
}

func Test_DomainPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainPropertiesARM, DomainPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainPropertiesARM runs a test to see if a specific instance of DomainPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainPropertiesARM(subject DomainPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainPropertiesARM
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

// Generator of DomainPropertiesARM instances for property testing - lazily instantiated by
//DomainPropertiesARMGenerator()
var domainPropertiesARMGenerator gopter.Gen

// DomainPropertiesARMGenerator returns a generator of DomainPropertiesARM instances for property testing.
// We first initialize domainPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainPropertiesARMGenerator() gopter.Gen {
	if domainPropertiesARMGenerator != nil {
		return domainPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesARM(generators)
	domainPropertiesARMGenerator = gen.Struct(reflect.TypeOf(DomainPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForDomainPropertiesARM(generators)
	domainPropertiesARMGenerator = gen.Struct(reflect.TypeOf(DomainPropertiesARM{}), generators)

	return domainPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainPropertiesARM(gens map[string]gopter.Gen) {
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesInputSchemaCloudEventSchemaV10, DomainPropertiesInputSchemaCustomEventSchema, DomainPropertiesInputSchemaEventGridSchema))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesPublicNetworkAccessDisabled, DomainPropertiesPublicNetworkAccessEnabled))
}

// AddRelatedPropertyGeneratorsForDomainPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainPropertiesARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(JsonInputSchemaMappingARMGenerator())
}

func Test_InboundIpRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRuleARM, InboundIpRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRuleARM runs a test to see if a specific instance of InboundIpRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRuleARM(subject InboundIpRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRuleARM
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

// Generator of InboundIpRuleARM instances for property testing - lazily instantiated by InboundIpRuleARMGenerator()
var inboundIpRuleARMGenerator gopter.Gen

// InboundIpRuleARMGenerator returns a generator of InboundIpRuleARM instances for property testing.
func InboundIpRuleARMGenerator() gopter.Gen {
	if inboundIpRuleARMGenerator != nil {
		return inboundIpRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRuleARM(generators)
	inboundIpRuleARMGenerator = gen.Struct(reflect.TypeOf(InboundIpRuleARM{}), generators)

	return inboundIpRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRuleARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRuleActionAllow))
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_JsonInputSchemaMappingARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonInputSchemaMappingARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonInputSchemaMappingARM, JsonInputSchemaMappingARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonInputSchemaMappingARM runs a test to see if a specific instance of JsonInputSchemaMappingARM round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonInputSchemaMappingARM(subject JsonInputSchemaMappingARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonInputSchemaMappingARM
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

// Generator of JsonInputSchemaMappingARM instances for property testing - lazily instantiated by
//JsonInputSchemaMappingARMGenerator()
var jsonInputSchemaMappingARMGenerator gopter.Gen

// JsonInputSchemaMappingARMGenerator returns a generator of JsonInputSchemaMappingARM instances for property testing.
// We first initialize jsonInputSchemaMappingARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func JsonInputSchemaMappingARMGenerator() gopter.Gen {
	if jsonInputSchemaMappingARMGenerator != nil {
		return jsonInputSchemaMappingARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonInputSchemaMappingARM(generators)
	jsonInputSchemaMappingARMGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMappingARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonInputSchemaMappingARM(generators)
	AddRelatedPropertyGeneratorsForJsonInputSchemaMappingARM(generators)
	jsonInputSchemaMappingARMGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMappingARM{}), generators)

	return jsonInputSchemaMappingARMGenerator
}

// AddIndependentPropertyGeneratorsForJsonInputSchemaMappingARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonInputSchemaMappingARM(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(JsonInputSchemaMappingInputSchemaMappingTypeJson))
}

// AddRelatedPropertyGeneratorsForJsonInputSchemaMappingARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForJsonInputSchemaMappingARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(JsonInputSchemaMappingPropertiesARMGenerator())
}

func Test_JsonInputSchemaMappingPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonInputSchemaMappingPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonInputSchemaMappingPropertiesARM, JsonInputSchemaMappingPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonInputSchemaMappingPropertiesARM runs a test to see if a specific instance of JsonInputSchemaMappingPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonInputSchemaMappingPropertiesARM(subject JsonInputSchemaMappingPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonInputSchemaMappingPropertiesARM
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

// Generator of JsonInputSchemaMappingPropertiesARM instances for property testing - lazily instantiated by
//JsonInputSchemaMappingPropertiesARMGenerator()
var jsonInputSchemaMappingPropertiesARMGenerator gopter.Gen

// JsonInputSchemaMappingPropertiesARMGenerator returns a generator of JsonInputSchemaMappingPropertiesARM instances for property testing.
func JsonInputSchemaMappingPropertiesARMGenerator() gopter.Gen {
	if jsonInputSchemaMappingPropertiesARMGenerator != nil {
		return jsonInputSchemaMappingPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForJsonInputSchemaMappingPropertiesARM(generators)
	jsonInputSchemaMappingPropertiesARMGenerator = gen.Struct(reflect.TypeOf(JsonInputSchemaMappingPropertiesARM{}), generators)

	return jsonInputSchemaMappingPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForJsonInputSchemaMappingPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForJsonInputSchemaMappingPropertiesARM(gens map[string]gopter.Gen) {
	gens["DataVersion"] = gen.PtrOf(JsonFieldWithDefaultARMGenerator())
	gens["EventTime"] = gen.PtrOf(JsonFieldARMGenerator())
	gens["EventType"] = gen.PtrOf(JsonFieldWithDefaultARMGenerator())
	gens["Id"] = gen.PtrOf(JsonFieldARMGenerator())
	gens["Subject"] = gen.PtrOf(JsonFieldWithDefaultARMGenerator())
	gens["Topic"] = gen.PtrOf(JsonFieldARMGenerator())
}

func Test_JsonFieldARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonFieldARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonFieldARM, JsonFieldARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonFieldARM runs a test to see if a specific instance of JsonFieldARM round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonFieldARM(subject JsonFieldARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonFieldARM
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

// Generator of JsonFieldARM instances for property testing - lazily instantiated by JsonFieldARMGenerator()
var jsonFieldARMGenerator gopter.Gen

// JsonFieldARMGenerator returns a generator of JsonFieldARM instances for property testing.
func JsonFieldARMGenerator() gopter.Gen {
	if jsonFieldARMGenerator != nil {
		return jsonFieldARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonFieldARM(generators)
	jsonFieldARMGenerator = gen.Struct(reflect.TypeOf(JsonFieldARM{}), generators)

	return jsonFieldARMGenerator
}

// AddIndependentPropertyGeneratorsForJsonFieldARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonFieldARM(gens map[string]gopter.Gen) {
	gens["SourceField"] = gen.PtrOf(gen.AlphaString())
}

func Test_JsonFieldWithDefaultARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of JsonFieldWithDefaultARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForJsonFieldWithDefaultARM, JsonFieldWithDefaultARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForJsonFieldWithDefaultARM runs a test to see if a specific instance of JsonFieldWithDefaultARM round trips to JSON and back losslessly
func RunJSONSerializationTestForJsonFieldWithDefaultARM(subject JsonFieldWithDefaultARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual JsonFieldWithDefaultARM
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

// Generator of JsonFieldWithDefaultARM instances for property testing - lazily instantiated by
//JsonFieldWithDefaultARMGenerator()
var jsonFieldWithDefaultARMGenerator gopter.Gen

// JsonFieldWithDefaultARMGenerator returns a generator of JsonFieldWithDefaultARM instances for property testing.
func JsonFieldWithDefaultARMGenerator() gopter.Gen {
	if jsonFieldWithDefaultARMGenerator != nil {
		return jsonFieldWithDefaultARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForJsonFieldWithDefaultARM(generators)
	jsonFieldWithDefaultARMGenerator = gen.Struct(reflect.TypeOf(JsonFieldWithDefaultARM{}), generators)

	return jsonFieldWithDefaultARMGenerator
}

// AddIndependentPropertyGeneratorsForJsonFieldWithDefaultARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForJsonFieldWithDefaultARM(gens map[string]gopter.Gen) {
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["SourceField"] = gen.PtrOf(gen.AlphaString())
}
