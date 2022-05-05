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

func Test_Domain_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainStatusARM, DomainStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainStatusARM runs a test to see if a specific instance of Domain_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainStatusARM(subject Domain_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_StatusARM
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

// Generator of Domain_StatusARM instances for property testing - lazily instantiated by DomainStatusARMGenerator()
var domainStatusARMGenerator gopter.Gen

// DomainStatusARMGenerator returns a generator of Domain_StatusARM instances for property testing.
// We first initialize domainStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainStatusARMGenerator() gopter.Gen {
	if domainStatusARMGenerator != nil {
		return domainStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainStatusARM(generators)
	domainStatusARMGenerator = gen.Struct(reflect.TypeOf(Domain_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainStatusARM(generators)
	AddRelatedPropertyGeneratorsForDomainStatusARM(generators)
	domainStatusARMGenerator = gen.Struct(reflect.TypeOf(Domain_StatusARM{}), generators)

	return domainStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainPropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_DomainProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainPropertiesStatusARM, DomainPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainPropertiesStatusARM runs a test to see if a specific instance of DomainProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainPropertiesStatusARM(subject DomainProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainProperties_StatusARM
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

// Generator of DomainProperties_StatusARM instances for property testing - lazily instantiated by
// DomainPropertiesStatusARMGenerator()
var domainPropertiesStatusARMGenerator gopter.Gen

// DomainPropertiesStatusARMGenerator returns a generator of DomainProperties_StatusARM instances for property testing.
// We first initialize domainPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainPropertiesStatusARMGenerator() gopter.Gen {
	if domainPropertiesStatusARMGenerator != nil {
		return domainPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesStatusARM(generators)
	domainPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForDomainPropertiesStatusARM(generators)
	domainPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_StatusARM{}), generators)

	return domainPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesStatusInputSchemaCloudEventSchemaV10, DomainPropertiesStatusInputSchemaCustomEventSchema, DomainPropertiesStatusInputSchemaEventGridSchema))
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainPropertiesStatusProvisioningStateCanceled,
		DomainPropertiesStatusProvisioningStateCreating,
		DomainPropertiesStatusProvisioningStateDeleting,
		DomainPropertiesStatusProvisioningStateFailed,
		DomainPropertiesStatusProvisioningStateSucceeded,
		DomainPropertiesStatusProvisioningStateUpdating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesStatusPublicNetworkAccessDisabled, DomainPropertiesStatusPublicNetworkAccessEnabled))
}

// AddRelatedPropertyGeneratorsForDomainPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleStatusARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingStatusARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator())
}

func Test_InboundIpRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRuleStatusARM, InboundIpRuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRuleStatusARM runs a test to see if a specific instance of InboundIpRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRuleStatusARM(subject InboundIpRule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule_StatusARM
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

// Generator of InboundIpRule_StatusARM instances for property testing - lazily instantiated by
// InboundIpRuleStatusARMGenerator()
var inboundIpRuleStatusARMGenerator gopter.Gen

// InboundIpRuleStatusARMGenerator returns a generator of InboundIpRule_StatusARM instances for property testing.
func InboundIpRuleStatusARMGenerator() gopter.Gen {
	if inboundIpRuleStatusARMGenerator != nil {
		return inboundIpRuleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRuleStatusARM(generators)
	inboundIpRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(InboundIpRule_StatusARM{}), generators)

	return inboundIpRuleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRuleStatusActionAllow))
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InputSchemaMapping_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMapping_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMappingStatusARM, InputSchemaMappingStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMappingStatusARM runs a test to see if a specific instance of InputSchemaMapping_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMappingStatusARM(subject InputSchemaMapping_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMapping_StatusARM
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

// Generator of InputSchemaMapping_StatusARM instances for property testing - lazily instantiated by
// InputSchemaMappingStatusARMGenerator()
var inputSchemaMappingStatusARMGenerator gopter.Gen

// InputSchemaMappingStatusARMGenerator returns a generator of InputSchemaMapping_StatusARM instances for property testing.
func InputSchemaMappingStatusARMGenerator() gopter.Gen {
	if inputSchemaMappingStatusARMGenerator != nil {
		return inputSchemaMappingStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMappingStatusARM(generators)
	inputSchemaMappingStatusARMGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping_StatusARM{}), generators)

	return inputSchemaMappingStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMappingStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMappingStatusARM(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(InputSchemaMappingStatusInputSchemaMappingTypeJson))
}

func Test_PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM, PrivateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM(subject PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM(generators)
	privateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionStatusDomainSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusDomainSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
