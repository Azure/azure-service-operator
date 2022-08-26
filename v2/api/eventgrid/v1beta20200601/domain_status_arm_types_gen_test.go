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

func Test_Domain_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain_STATUSARM, Domain_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain_STATUSARM runs a test to see if a specific instance of Domain_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain_STATUSARM(subject Domain_STATUSARM) string {
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

// Generator of Domain_STATUSARM instances for property testing - lazily instantiated by Domain_STATUSARMGenerator()
var domain_STATUSARMGenerator gopter.Gen

// Domain_STATUSARMGenerator returns a generator of Domain_STATUSARM instances for property testing.
// We first initialize domain_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Domain_STATUSARMGenerator() gopter.Gen {
	if domain_STATUSARMGenerator != nil {
		return domain_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_STATUSARM(generators)
	domain_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Domain_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDomain_STATUSARM(generators)
	domain_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Domain_STATUSARM{}), generators)

	return domain_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomain_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomain_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomain_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DomainProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_DomainProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainProperties_STATUSARM, DomainProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainProperties_STATUSARM runs a test to see if a specific instance of DomainProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainProperties_STATUSARM(subject DomainProperties_STATUSARM) string {
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
// DomainProperties_STATUSARMGenerator()
var domainProperties_STATUSARMGenerator gopter.Gen

// DomainProperties_STATUSARMGenerator returns a generator of DomainProperties_STATUSARM instances for property testing.
// We first initialize domainProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainProperties_STATUSARMGenerator() gopter.Gen {
	if domainProperties_STATUSARMGenerator != nil {
		return domainProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainProperties_STATUSARM(generators)
	domainProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDomainProperties_STATUSARM(generators)
	domainProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DomainProperties_STATUSARM{}), generators)

	return domainProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
<<<<<<< HEAD
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainProperties_InputSchema_CloudEventSchemaV1_0_STATUS, DomainProperties_InputSchema_CustomEventSchema_STATUS, DomainProperties_InputSchema_EventGridSchema_STATUS))
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainProperties_ProvisioningState_Canceled_STATUS,
		DomainProperties_ProvisioningState_Creating_STATUS,
		DomainProperties_ProvisioningState_Deleting_STATUS,
		DomainProperties_ProvisioningState_Failed_STATUS,
		DomainProperties_ProvisioningState_Succeeded_STATUS,
		DomainProperties_ProvisioningState_Updating_STATUS))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainProperties_PublicNetworkAccess_Disabled_STATUS, DomainProperties_PublicNetworkAccess_Enabled_STATUS))
=======
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainProperties_STATUS_InputSchema_CloudEventSchemaV1_0, DomainProperties_STATUS_InputSchema_CustomEventSchema, DomainProperties_STATUS_InputSchema_EventGridSchema))
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainProperties_STATUS_ProvisioningState_Canceled,
		DomainProperties_STATUS_ProvisioningState_Creating,
		DomainProperties_STATUS_ProvisioningState_Deleting,
		DomainProperties_STATUS_ProvisioningState_Failed,
		DomainProperties_STATUS_ProvisioningState_Succeeded,
		DomainProperties_STATUS_ProvisioningState_Updating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainProperties_STATUS_PublicNetworkAccess_Disabled, DomainProperties_STATUS_PublicNetworkAccess_Enabled))
>>>>>>> main
}

// AddRelatedPropertyGeneratorsForDomainProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRule_STATUSARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMapping_STATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator())
<<<<<<< HEAD
}

func Test_SystemData_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUSARM, SystemData_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUSARM runs a test to see if a specific instance of SystemData_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUSARM(subject SystemData_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUSARM
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

// Generator of SystemData_STATUSARM instances for property testing - lazily instantiated by
// SystemData_STATUSARMGenerator()
var systemData_STATUSARMGenerator gopter.Gen

// SystemData_STATUSARMGenerator returns a generator of SystemData_STATUSARM instances for property testing.
func SystemData_STATUSARMGenerator() gopter.Gen {
	if systemData_STATUSARMGenerator != nil {
		return systemData_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUSARM(generators)
	systemData_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUSARM{}), generators)

	return systemData_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUSARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_Application_STATUS,
		SystemData_CreatedByType_Key_STATUS,
		SystemData_CreatedByType_ManagedIdentity_STATUS,
		SystemData_CreatedByType_User_STATUS))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_Application_STATUS,
		SystemData_LastModifiedByType_Key_STATUS,
		SystemData_LastModifiedByType_ManagedIdentity_STATUS,
		SystemData_LastModifiedByType_User_STATUS))
=======
>>>>>>> main
}

func Test_InboundIpRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRule_STATUSARM, InboundIpRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRule_STATUSARM runs a test to see if a specific instance of InboundIpRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRule_STATUSARM(subject InboundIpRule_STATUSARM) string {
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
// InboundIpRule_STATUSARMGenerator()
var inboundIpRule_STATUSARMGenerator gopter.Gen

// InboundIpRule_STATUSARMGenerator returns a generator of InboundIpRule_STATUSARM instances for property testing.
func InboundIpRule_STATUSARMGenerator() gopter.Gen {
	if inboundIpRule_STATUSARMGenerator != nil {
		return inboundIpRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRule_STATUSARM(generators)
	inboundIpRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(InboundIpRule_STATUSARM{}), generators)

	return inboundIpRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRule_STATUSARM(gens map[string]gopter.Gen) {
<<<<<<< HEAD
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRule_Action_Allow_STATUS))
=======
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRule_STATUS_Action_Allow))
>>>>>>> main
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
		prop.ForAll(RunJSONSerializationTestForInputSchemaMapping_STATUSARM, InputSchemaMapping_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMapping_STATUSARM runs a test to see if a specific instance of InputSchemaMapping_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMapping_STATUSARM(subject InputSchemaMapping_STATUSARM) string {
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
// InputSchemaMapping_STATUSARMGenerator()
var inputSchemaMapping_STATUSARMGenerator gopter.Gen

// InputSchemaMapping_STATUSARMGenerator returns a generator of InputSchemaMapping_STATUSARM instances for property testing.
func InputSchemaMapping_STATUSARMGenerator() gopter.Gen {
	if inputSchemaMapping_STATUSARMGenerator != nil {
		return inputSchemaMapping_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMapping_STATUSARM(generators)
	inputSchemaMapping_STATUSARMGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping_STATUSARM{}), generators)

	return inputSchemaMapping_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMapping_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMapping_STATUSARM(gens map[string]gopter.Gen) {
<<<<<<< HEAD
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(InputSchemaMapping_InputSchemaMappingType_Json_STATUS))
=======
	gens["InputSchemaMappingType"] = gen.PtrOf(gen.OneConstOf(InputSchemaMapping_STATUS_InputSchemaMappingType_Json))
>>>>>>> main
}

func Test_PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM, PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM) string {
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
// instantiated by PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator()
var privateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_Domain_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
