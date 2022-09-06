// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_EHNamespace_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EHNamespace_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEHNamespace_STATUSARM, EHNamespace_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEHNamespace_STATUSARM runs a test to see if a specific instance of EHNamespace_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEHNamespace_STATUSARM(subject EHNamespace_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EHNamespace_STATUSARM
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

// Generator of EHNamespace_STATUSARM instances for property testing - lazily instantiated by
// EHNamespace_STATUSARMGenerator()
var ehNamespace_STATUSARMGenerator gopter.Gen

// EHNamespace_STATUSARMGenerator returns a generator of EHNamespace_STATUSARM instances for property testing.
// We first initialize ehNamespace_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EHNamespace_STATUSARMGenerator() gopter.Gen {
	if ehNamespace_STATUSARMGenerator != nil {
		return ehNamespace_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEHNamespace_STATUSARM(generators)
	ehNamespace_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EHNamespace_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEHNamespace_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForEHNamespace_STATUSARM(generators)
	ehNamespace_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EHNamespace_STATUSARM{}), generators)

	return ehNamespace_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForEHNamespace_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEHNamespace_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEHNamespace_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEHNamespace_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(EHNamespace_STATUS_PropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_EHNamespace_STATUS_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EHNamespace_STATUS_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEHNamespace_STATUS_PropertiesARM, EHNamespace_STATUS_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEHNamespace_STATUS_PropertiesARM runs a test to see if a specific instance of EHNamespace_STATUS_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEHNamespace_STATUS_PropertiesARM(subject EHNamespace_STATUS_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EHNamespace_STATUS_PropertiesARM
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

// Generator of EHNamespace_STATUS_PropertiesARM instances for property testing - lazily instantiated by
// EHNamespace_STATUS_PropertiesARMGenerator()
var ehNamespace_STATUS_PropertiesARMGenerator gopter.Gen

// EHNamespace_STATUS_PropertiesARMGenerator returns a generator of EHNamespace_STATUS_PropertiesARM instances for property testing.
// We first initialize ehNamespace_STATUS_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EHNamespace_STATUS_PropertiesARMGenerator() gopter.Gen {
	if ehNamespace_STATUS_PropertiesARMGenerator != nil {
		return ehNamespace_STATUS_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM(generators)
	ehNamespace_STATUS_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(EHNamespace_STATUS_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM(generators)
	ehNamespace_STATUS_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(EHNamespace_STATUS_PropertiesARM{}), generators)

	return ehNamespace_STATUS_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM(gens map[string]gopter.Gen) {
	gens["AlternateName"] = gen.PtrOf(gen.AlphaString())
	gens["ClusterArmId"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAutoInflateEnabled"] = gen.PtrOf(gen.Bool())
	gens["KafkaEnabled"] = gen.PtrOf(gen.Bool())
	gens["MaximumThroughputUnits"] = gen.PtrOf(gen.Int())
	gens["MetricId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceBusEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEHNamespace_STATUS_PropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(Encryption_STATUSARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator())
}

func Test_Identity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUSARM, Identity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUSARM runs a test to see if a specific instance of Identity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUSARM(subject Identity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUSARM
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

// Generator of Identity_STATUSARM instances for property testing - lazily instantiated by Identity_STATUSARMGenerator()
var identity_STATUSARMGenerator gopter.Gen

// Identity_STATUSARMGenerator returns a generator of Identity_STATUSARM instances for property testing.
// We first initialize identity_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_STATUSARMGenerator() gopter.Gen {
	if identity_STATUSARMGenerator != nil {
		return identity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUSARM(generators)
	identity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_STATUSARM(generators)
	identity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	return identity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_STATUS_Type_None,
		Identity_STATUS_Type_SystemAssigned,
		Identity_STATUS_Type_SystemAssignedUserAssigned,
		Identity_STATUS_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserAssignedIdentity_STATUSARMGenerator())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Name_Basic, Sku_STATUS_Name_Premium, Sku_STATUS_Name_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Tier_Basic, Sku_STATUS_Tier_Premium, Sku_STATUS_Tier_Standard))
}

func Test_Encryption_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryption_STATUSARM, Encryption_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryption_STATUSARM runs a test to see if a specific instance of Encryption_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryption_STATUSARM(subject Encryption_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_STATUSARM
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

// Generator of Encryption_STATUSARM instances for property testing - lazily instantiated by
// Encryption_STATUSARMGenerator()
var encryption_STATUSARMGenerator gopter.Gen

// Encryption_STATUSARMGenerator returns a generator of Encryption_STATUSARM instances for property testing.
// We first initialize encryption_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Encryption_STATUSARMGenerator() gopter.Gen {
	if encryption_STATUSARMGenerator != nil {
		return encryption_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_STATUSARM(generators)
	encryption_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Encryption_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForEncryption_STATUSARM(generators)
	encryption_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Encryption_STATUSARM{}), generators)

	return encryption_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryption_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryption_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(Encryption_STATUS_KeySource_MicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryption_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryption_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultProperties_STATUSARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM, PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(subject PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator()
var privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM instances for property testing.
// We first initialize privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnection_STATUS_SubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnection_STATUS_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_UserAssignedIdentity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUSARM, UserAssignedIdentity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUSARM runs a test to see if a specific instance of UserAssignedIdentity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUSARM(subject UserAssignedIdentity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUSARM
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

// Generator of UserAssignedIdentity_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSARMGenerator()
var userAssignedIdentity_STATUSARMGenerator gopter.Gen

// UserAssignedIdentity_STATUSARMGenerator returns a generator of UserAssignedIdentity_STATUSARM instances for property testing.
func UserAssignedIdentity_STATUSARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSARMGenerator != nil {
		return userAssignedIdentity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(generators)
	userAssignedIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUSARM{}), generators)

	return userAssignedIdentity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_STATUSARM, KeyVaultProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_STATUSARM runs a test to see if a specific instance of KeyVaultProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_STATUSARM(subject KeyVaultProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_STATUSARM
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

// Generator of KeyVaultProperties_STATUSARM instances for property testing - lazily instantiated by
// KeyVaultProperties_STATUSARMGenerator()
var keyVaultProperties_STATUSARMGenerator gopter.Gen

// KeyVaultProperties_STATUSARMGenerator returns a generator of KeyVaultProperties_STATUSARM instances for property testing.
// We first initialize keyVaultProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultProperties_STATUSARMGenerator() gopter.Gen {
	if keyVaultProperties_STATUSARMGenerator != nil {
		return keyVaultProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(generators)
	keyVaultProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUSARM(generators)
	keyVaultProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUSARM{}), generators)

	return keyVaultProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityProperties_STATUSARMGenerator())
}

func Test_UserAssignedIdentityProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM, UserAssignedIdentityProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM runs a test to see if a specific instance of UserAssignedIdentityProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityProperties_STATUSARM(subject UserAssignedIdentityProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_STATUSARM
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

// Generator of UserAssignedIdentityProperties_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentityProperties_STATUSARMGenerator()
var userAssignedIdentityProperties_STATUSARMGenerator gopter.Gen

// UserAssignedIdentityProperties_STATUSARMGenerator returns a generator of UserAssignedIdentityProperties_STATUSARM instances for property testing.
func UserAssignedIdentityProperties_STATUSARMGenerator() gopter.Gen {
	if userAssignedIdentityProperties_STATUSARMGenerator != nil {
		return userAssignedIdentityProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM(generators)
	userAssignedIdentityProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_STATUSARM{}), generators)

	return userAssignedIdentityProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
