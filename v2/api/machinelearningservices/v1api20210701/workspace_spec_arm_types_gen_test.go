// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210701

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

func Test_CosmosDbSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CosmosDbSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCosmosDbSettings_ARM, CosmosDbSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCosmosDbSettings_ARM runs a test to see if a specific instance of CosmosDbSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCosmosDbSettings_ARM(subject CosmosDbSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CosmosDbSettings_ARM
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

// Generator of CosmosDbSettings_ARM instances for property testing - lazily instantiated by
// CosmosDbSettings_ARMGenerator()
var cosmosDbSettings_ARMGenerator gopter.Gen

// CosmosDbSettings_ARMGenerator returns a generator of CosmosDbSettings_ARM instances for property testing.
func CosmosDbSettings_ARMGenerator() gopter.Gen {
	if cosmosDbSettings_ARMGenerator != nil {
		return cosmosDbSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCosmosDbSettings_ARM(generators)
	cosmosDbSettings_ARMGenerator = gen.Struct(reflect.TypeOf(CosmosDbSettings_ARM{}), generators)

	return cosmosDbSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCosmosDbSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCosmosDbSettings_ARM(gens map[string]gopter.Gen) {
	gens["CollectionsThroughput"] = gen.PtrOf(gen.Int())
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
	gens["Identity"] = gen.PtrOf(IdentityForCmk_ARMGenerator())
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_ARMGenerator())
}

func Test_IdentityForCmk_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityForCmk_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityForCmk_ARM, IdentityForCmk_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityForCmk_ARM runs a test to see if a specific instance of IdentityForCmk_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityForCmk_ARM(subject IdentityForCmk_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityForCmk_ARM
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

// Generator of IdentityForCmk_ARM instances for property testing - lazily instantiated by IdentityForCmk_ARMGenerator()
var identityForCmk_ARMGenerator gopter.Gen

// IdentityForCmk_ARMGenerator returns a generator of IdentityForCmk_ARM instances for property testing.
func IdentityForCmk_ARMGenerator() gopter.Gen {
	if identityForCmk_ARMGenerator != nil {
		return identityForCmk_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityForCmk_ARM(generators)
	identityForCmk_ARMGenerator = gen.Struct(reflect.TypeOf(IdentityForCmk_ARM{}), generators)

	return identityForCmk_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityForCmk_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityForCmk_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
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
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetails_ARMGenerator())
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
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultArmId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceManagedResourcesSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceManagedResourcesSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceManagedResourcesSettings_ARM, ServiceManagedResourcesSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceManagedResourcesSettings_ARM runs a test to see if a specific instance of ServiceManagedResourcesSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceManagedResourcesSettings_ARM(subject ServiceManagedResourcesSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceManagedResourcesSettings_ARM
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

// Generator of ServiceManagedResourcesSettings_ARM instances for property testing - lazily instantiated by
// ServiceManagedResourcesSettings_ARMGenerator()
var serviceManagedResourcesSettings_ARMGenerator gopter.Gen

// ServiceManagedResourcesSettings_ARMGenerator returns a generator of ServiceManagedResourcesSettings_ARM instances for property testing.
func ServiceManagedResourcesSettings_ARMGenerator() gopter.Gen {
	if serviceManagedResourcesSettings_ARMGenerator != nil {
		return serviceManagedResourcesSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_ARM(generators)
	serviceManagedResourcesSettings_ARMGenerator = gen.Struct(reflect.TypeOf(ServiceManagedResourcesSettings_ARM{}), generators)

	return serviceManagedResourcesSettings_ARMGenerator
}

// AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings_ARM(gens map[string]gopter.Gen) {
	gens["CosmosDb"] = gen.PtrOf(CosmosDbSettings_ARMGenerator())
}

func Test_SharedPrivateLinkResourceProperty_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperty_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperty_ARM, SharedPrivateLinkResourceProperty_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperty_ARM runs a test to see if a specific instance of SharedPrivateLinkResourceProperty_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperty_ARM(subject SharedPrivateLinkResourceProperty_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperty_ARM
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

// Generator of SharedPrivateLinkResourceProperty_ARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceProperty_ARMGenerator()
var sharedPrivateLinkResourceProperty_ARMGenerator gopter.Gen

// SharedPrivateLinkResourceProperty_ARMGenerator returns a generator of SharedPrivateLinkResourceProperty_ARM instances for property testing.
func SharedPrivateLinkResourceProperty_ARMGenerator() gopter.Gen {
	if sharedPrivateLinkResourceProperty_ARMGenerator != nil {
		return sharedPrivateLinkResourceProperty_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_ARM(generators)
	sharedPrivateLinkResourceProperty_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperty_ARM{}), generators)

	return sharedPrivateLinkResourceProperty_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty_ARM(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointServiceConnectionStatus_Approved,
		PrivateEndpointServiceConnectionStatus_Disconnected,
		PrivateEndpointServiceConnectionStatus_Pending,
		PrivateEndpointServiceConnectionStatus_Rejected,
		PrivateEndpointServiceConnectionStatus_Timeout))
}

func Test_SharedPrivateLinkResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResource_ARM, SharedPrivateLinkResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResource_ARM runs a test to see if a specific instance of SharedPrivateLinkResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResource_ARM(subject SharedPrivateLinkResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResource_ARM
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

// Generator of SharedPrivateLinkResource_ARM instances for property testing - lazily instantiated by
// SharedPrivateLinkResource_ARMGenerator()
var sharedPrivateLinkResource_ARMGenerator gopter.Gen

// SharedPrivateLinkResource_ARMGenerator returns a generator of SharedPrivateLinkResource_ARM instances for property testing.
// We first initialize sharedPrivateLinkResource_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResource_ARMGenerator() gopter.Gen {
	if sharedPrivateLinkResource_ARMGenerator != nil {
		return sharedPrivateLinkResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_ARM(generators)
	sharedPrivateLinkResource_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_ARM(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_ARM(generators)
	sharedPrivateLinkResource_ARMGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource_ARM{}), generators)

	return sharedPrivateLinkResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResource_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResource_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedPrivateLinkResourceProperty_ARMGenerator())
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
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_ARM, SystemData_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_ARM runs a test to see if a specific instance of SystemData_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_ARM(subject SystemData_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_ARM
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

// Generator of SystemData_ARM instances for property testing - lazily instantiated by SystemData_ARMGenerator()
var systemData_ARMGenerator gopter.Gen

// SystemData_ARMGenerator returns a generator of SystemData_ARM instances for property testing.
func SystemData_ARMGenerator() gopter.Gen {
	if systemData_ARMGenerator != nil {
		return systemData_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_ARM(generators)
	systemData_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_ARM{}), generators)

	return systemData_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_Application,
		SystemData_CreatedByType_Key,
		SystemData_CreatedByType_ManagedIdentity,
		SystemData_CreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_Application,
		SystemData_LastModifiedByType_Key,
		SystemData_LastModifiedByType_ManagedIdentity,
		SystemData_LastModifiedByType_User))
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

func Test_WorkspaceProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties_ARM, WorkspaceProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties_ARM runs a test to see if a specific instance of WorkspaceProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties_ARM(subject WorkspaceProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_ARM
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

// Generator of WorkspaceProperties_ARM instances for property testing - lazily instantiated by
// WorkspaceProperties_ARMGenerator()
var workspaceProperties_ARMGenerator gopter.Gen

// WorkspaceProperties_ARMGenerator returns a generator of WorkspaceProperties_ARM instances for property testing.
// We first initialize workspaceProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceProperties_ARMGenerator() gopter.Gen {
	if workspaceProperties_ARMGenerator != nil {
		return workspaceProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	workspaceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM(generators)
	workspaceProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_ARM{}), generators)

	return workspaceProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties_ARM(gens map[string]gopter.Gen) {
	gens["AllowPublicAccessWhenBehindVnet"] = gen.PtrOf(gen.Bool())
	gens["ApplicationInsights"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerRegistry"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DiscoveryUrl"] = gen.PtrOf(gen.AlphaString())
	gens["FriendlyName"] = gen.PtrOf(gen.AlphaString())
	gens["HbiWorkspace"] = gen.PtrOf(gen.Bool())
	gens["ImageBuildCompute"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVault"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccess_Disabled, WorkspaceProperties_PublicNetworkAccess_Enabled))
	gens["StorageAccount"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionProperty_ARMGenerator())
	gens["ServiceManagedResourcesSettings"] = gen.PtrOf(ServiceManagedResourcesSettings_ARMGenerator())
	gens["SharedPrivateLinkResources"] = gen.SliceOf(SharedPrivateLinkResource_ARMGenerator())
}

func Test_Workspace_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_Spec_ARM, Workspace_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_Spec_ARM runs a test to see if a specific instance of Workspace_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_Spec_ARM(subject Workspace_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_Spec_ARM
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

// Generator of Workspace_Spec_ARM instances for property testing - lazily instantiated by Workspace_Spec_ARMGenerator()
var workspace_Spec_ARMGenerator gopter.Gen

// Workspace_Spec_ARMGenerator returns a generator of Workspace_Spec_ARM instances for property testing.
// We first initialize workspace_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_Spec_ARMGenerator() gopter.Gen {
	if workspace_Spec_ARMGenerator != nil {
		return workspace_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	workspace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM(generators)
	workspace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec_ARM{}), generators)

	return workspace_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_ARMGenerator())
}
