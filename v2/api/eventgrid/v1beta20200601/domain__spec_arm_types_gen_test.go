// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

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

func Test_Domain_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain_SpecARM, Domain_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain_SpecARM runs a test to see if a specific instance of Domain_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain_SpecARM(subject Domain_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_SpecARM
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

// Generator of Domain_SpecARM instances for property testing - lazily instantiated by Domain_SpecARMGenerator()
var domain_SpecARMGenerator gopter.Gen

// Domain_SpecARMGenerator returns a generator of Domain_SpecARM instances for property testing.
// We first initialize domain_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Domain_SpecARMGenerator() gopter.Gen {
	if domain_SpecARMGenerator != nil {
		return domain_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_SpecARM(generators)
	domain_SpecARMGenerator = gen.Struct(reflect.TypeOf(Domain_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDomain_SpecARM(generators)
	domain_SpecARMGenerator = gen.Struct(reflect.TypeOf(Domain_SpecARM{}), generators)

	return domain_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDomain_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomain_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomain_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainPropertiesARMGenerator())
}

func Test_DomainPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
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
// DomainPropertiesARMGenerator()
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
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainProperties_InputSchema_CloudEventSchemaV1_0, DomainProperties_InputSchema_CustomEventSchema, DomainProperties_InputSchema_EventGridSchema))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainProperties_PublicNetworkAccess_Disabled, DomainProperties_PublicNetworkAccess_Enabled))
}

// AddRelatedPropertyGeneratorsForDomainPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainPropertiesARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingARMGenerator())
}

func Test_InboundIpRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
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
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRule_Action_Allow))
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InputSchemaMappingARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMappingARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMappingARM, InputSchemaMappingARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMappingARM runs a test to see if a specific instance of InputSchemaMappingARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMappingARM(subject InputSchemaMappingARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMappingARM
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

// Generator of InputSchemaMappingARM instances for property testing - lazily instantiated by
// InputSchemaMappingARMGenerator()
var inputSchemaMappingARMGenerator gopter.Gen

// InputSchemaMappingARMGenerator returns a generator of InputSchemaMappingARM instances for property testing.
func InputSchemaMappingARMGenerator() gopter.Gen {
	if inputSchemaMappingARMGenerator != nil {
		return inputSchemaMappingARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMappingARM(generators)
	inputSchemaMappingARMGenerator = gen.Struct(reflect.TypeOf(InputSchemaMappingARM{}), generators)

	return inputSchemaMappingARMGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMappingARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMappingARM(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(InputSchemaMapping_InputSchemaMappingType_Json))
}
