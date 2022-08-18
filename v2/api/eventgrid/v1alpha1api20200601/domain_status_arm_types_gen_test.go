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

func Test_Domain_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainSTATUSARM, DomainSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainSTATUSARM runs a test to see if a specific instance of Domain_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainSTATUSARM(subject Domain_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_STATUSARM
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

// Generator of Domain_STATUSARM instances for property testing - lazily instantiated by DomainSTATUSARMGenerator()
var domainSTATUSARMGenerator gopter.Gen

// DomainSTATUSARMGenerator returns a generator of Domain_STATUSARM instances for property testing.
// We first initialize domainSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainSTATUSARMGenerator() gopter.Gen {
	if domainSTATUSARMGenerator != nil {
		return domainSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainSTATUSARM(generators)
	domainSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Domain_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForDomainSTATUSARM(generators)
	domainSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Domain_STATUSARM{}), generators)

	return domainSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainPropertiesSTATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSARMGenerator())
}

func Test_DomainProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainPropertiesSTATUSARM, DomainPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainPropertiesSTATUSARM runs a test to see if a specific instance of DomainProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainPropertiesSTATUSARM(subject DomainProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainProperties_STATUSARM
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

// Generator of DomainProperties_STATUSARM instances for property testing - lazily instantiated by
// DomainPropertiesSTATUSARMGenerator()
var domainPropertiesSTATUSARMGenerator gopter.Gen

// DomainPropertiesSTATUSARMGenerator returns a generator of DomainProperties_STATUSARM instances for property testing.
// We first initialize domainPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainPropertiesSTATUSARMGenerator() gopter.Gen {
	if domainPropertiesSTATUSARMGenerator != nil {
		return domainPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesSTATUSARM(generators)
	domainPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForDomainPropertiesSTATUSARM(generators)
	domainPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_STATUSARM{}), generators)

	return domainPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesSTATUSInputSchema_CloudEventSchemaV10, DomainPropertiesSTATUSInputSchema_CustomEventSchema, DomainPropertiesSTATUSInputSchema_EventGridSchema))
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainPropertiesSTATUSProvisioningState_Canceled,
		DomainPropertiesSTATUSProvisioningState_Creating,
		DomainPropertiesSTATUSProvisioningState_Deleting,
		DomainPropertiesSTATUSProvisioningState_Failed,
		DomainPropertiesSTATUSProvisioningState_Succeeded,
		DomainPropertiesSTATUSProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesSTATUSPublicNetworkAccess_Disabled, DomainPropertiesSTATUSPublicNetworkAccess_Enabled))
}

// AddRelatedPropertyGeneratorsForDomainPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleSTATUSARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingSTATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator())
}

func Test_InboundIpRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRuleSTATUSARM, InboundIpRuleSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRuleSTATUSARM runs a test to see if a specific instance of InboundIpRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRuleSTATUSARM(subject InboundIpRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule_STATUSARM
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

// Generator of InboundIpRule_STATUSARM instances for property testing - lazily instantiated by
// InboundIpRuleSTATUSARMGenerator()
var inboundIpRuleSTATUSARMGenerator gopter.Gen

// InboundIpRuleSTATUSARMGenerator returns a generator of InboundIpRule_STATUSARM instances for property testing.
func InboundIpRuleSTATUSARMGenerator() gopter.Gen {
	if inboundIpRuleSTATUSARMGenerator != nil {
		return inboundIpRuleSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRuleSTATUSARM(generators)
	inboundIpRuleSTATUSARMGenerator = gen.Struct(reflect.TypeOf(InboundIpRule_STATUSARM{}), generators)

	return inboundIpRuleSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRuleSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRuleSTATUSARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRuleSTATUSAction_Allow))
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InputSchemaMapping_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMapping_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMappingSTATUSARM, InputSchemaMappingSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMappingSTATUSARM runs a test to see if a specific instance of InputSchemaMapping_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMappingSTATUSARM(subject InputSchemaMapping_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMapping_STATUSARM
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

// Generator of InputSchemaMapping_STATUSARM instances for property testing - lazily instantiated by
// InputSchemaMappingSTATUSARMGenerator()
var inputSchemaMappingSTATUSARMGenerator gopter.Gen

// InputSchemaMappingSTATUSARMGenerator returns a generator of InputSchemaMapping_STATUSARM instances for property testing.
func InputSchemaMappingSTATUSARMGenerator() gopter.Gen {
	if inputSchemaMappingSTATUSARMGenerator != nil {
		return inputSchemaMappingSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMappingSTATUSARM(generators)
	inputSchemaMappingSTATUSARMGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping_STATUSARM{}), generators)

	return inputSchemaMappingSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMappingSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMappingSTATUSARM(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(InputSchemaMappingSTATUSInputSchemaMappingType_Json))
}

func Test_PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM, PrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM(generators)
	privateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionSTATUSDomainSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSDomainSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
