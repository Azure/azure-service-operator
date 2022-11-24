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

func Test_Namespace_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespace_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_STATUS_ARM, Namespace_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_STATUS_ARM runs a test to see if a specific instance of Namespace_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_STATUS_ARM(subject Namespace_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_STATUS_ARM
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

// Generator of Namespace_STATUS_ARM instances for property testing - lazily instantiated by
// Namespace_STATUS_ARMGenerator()
var namespace_STATUS_ARMGenerator gopter.Gen

// Namespace_STATUS_ARMGenerator returns a generator of Namespace_STATUS_ARM instances for property testing.
// We first initialize namespace_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_STATUS_ARMGenerator() gopter.Gen {
	if namespace_STATUS_ARMGenerator != nil {
		return namespace_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_STATUS_ARM(generators)
	namespace_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_STATUS_ARM(generators)
	namespace_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_STATUS_ARM{}), generators)

	return namespace_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespace_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(Namespace_Properties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_Identity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS_ARM, Identity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS_ARM runs a test to see if a specific instance of Identity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS_ARM(subject Identity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS_ARM
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

// Generator of Identity_STATUS_ARM instances for property testing - lazily instantiated by
// Identity_STATUS_ARMGenerator()
var identity_STATUS_ARMGenerator gopter.Gen

// Identity_STATUS_ARMGenerator returns a generator of Identity_STATUS_ARM instances for property testing.
// We first initialize identity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_STATUS_ARMGenerator() gopter.Gen {
	if identity_STATUS_ARMGenerator != nil {
		return identity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	return identity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_STATUS_None,
		Identity_Type_STATUS_SystemAssigned,
		Identity_Type_STATUS_SystemAssignedUserAssigned,
		Identity_Type_STATUS_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserAssignedIdentity_STATUS_ARMGenerator())
}

func Test_Namespace_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespace_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_Properties_STATUS_ARM, Namespace_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_Properties_STATUS_ARM runs a test to see if a specific instance of Namespace_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_Properties_STATUS_ARM(subject Namespace_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_Properties_STATUS_ARM
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

// Generator of Namespace_Properties_STATUS_ARM instances for property testing - lazily instantiated by
// Namespace_Properties_STATUS_ARMGenerator()
var namespace_Properties_STATUS_ARMGenerator gopter.Gen

// Namespace_Properties_STATUS_ARMGenerator returns a generator of Namespace_Properties_STATUS_ARM instances for property testing.
// We first initialize namespace_Properties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_Properties_STATUS_ARMGenerator() gopter.Gen {
	if namespace_Properties_STATUS_ARMGenerator != nil {
		return namespace_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Properties_STATUS_ARM(generators)
	namespace_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Properties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Properties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_Properties_STATUS_ARM(generators)
	namespace_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Properties_STATUS_ARM{}), generators)

	return namespace_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForNamespace_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(Encryption_STATUS_ARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_ARMGenerator())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
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

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_STATUS_Basic, Sku_Name_STATUS_Premium, Sku_Name_STATUS_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_STATUS_Basic, Sku_Tier_STATUS_Premium, Sku_Tier_STATUS_Standard))
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}

func Test_Encryption_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryption_STATUS_ARM, Encryption_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryption_STATUS_ARM runs a test to see if a specific instance of Encryption_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryption_STATUS_ARM(subject Encryption_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_STATUS_ARM
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

// Generator of Encryption_STATUS_ARM instances for property testing - lazily instantiated by
// Encryption_STATUS_ARMGenerator()
var encryption_STATUS_ARMGenerator gopter.Gen

// Encryption_STATUS_ARMGenerator returns a generator of Encryption_STATUS_ARM instances for property testing.
// We first initialize encryption_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Encryption_STATUS_ARMGenerator() gopter.Gen {
	if encryption_STATUS_ARMGenerator != nil {
		return encryption_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_STATUS_ARM(generators)
	encryption_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Encryption_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryption_STATUS_ARM(generators)
	encryption_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Encryption_STATUS_ARM{}), generators)

	return encryption_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryption_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryption_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(Encryption_KeySource_STATUS_MicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryption_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryption_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultProperties_STATUS_ARMGenerator())
}

func Test_PrivateEndpointConnection_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM, PrivateEndpointConnection_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM(subject PrivateEndpointConnection_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_ARM
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

// Generator of PrivateEndpointConnection_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUS_ARMGenerator()
var privateEndpointConnection_STATUS_ARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_ARMGenerator returns a generator of PrivateEndpointConnection_STATUS_ARM instances for property testing.
func PrivateEndpointConnection_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_ARMGenerator != nil {
		return privateEndpointConnection_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(generators)
	privateEndpointConnection_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_ARM{}), generators)

	return privateEndpointConnection_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM, UserAssignedIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM runs a test to see if a specific instance of UserAssignedIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS_ARM(subject UserAssignedIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS_ARM
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

// Generator of UserAssignedIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUS_ARMGenerator()
var userAssignedIdentity_STATUS_ARMGenerator gopter.Gen

// UserAssignedIdentity_STATUS_ARMGenerator returns a generator of UserAssignedIdentity_STATUS_ARM instances for property testing.
func UserAssignedIdentity_STATUS_ARMGenerator() gopter.Gen {
	if userAssignedIdentity_STATUS_ARMGenerator != nil {
		return userAssignedIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(generators)
	userAssignedIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS_ARM{}), generators)

	return userAssignedIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM, KeyVaultProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM runs a test to see if a specific instance of KeyVaultProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_STATUS_ARM(subject KeyVaultProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_STATUS_ARM
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

// Generator of KeyVaultProperties_STATUS_ARM instances for property testing - lazily instantiated by
// KeyVaultProperties_STATUS_ARMGenerator()
var keyVaultProperties_STATUS_ARMGenerator gopter.Gen

// KeyVaultProperties_STATUS_ARMGenerator returns a generator of KeyVaultProperties_STATUS_ARM instances for property testing.
// We first initialize keyVaultProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultProperties_STATUS_ARMGenerator() gopter.Gen {
	if keyVaultProperties_STATUS_ARMGenerator != nil {
		return keyVaultProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(generators)
	keyVaultProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(generators)
	keyVaultProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUS_ARM{}), generators)

	return keyVaultProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityProperties_STATUS_ARMGenerator())
}

func Test_UserAssignedIdentityProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM, UserAssignedIdentityProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM runs a test to see if a specific instance of UserAssignedIdentityProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityProperties_STATUS_ARM(subject UserAssignedIdentityProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_STATUS_ARM
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

// Generator of UserAssignedIdentityProperties_STATUS_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityProperties_STATUS_ARMGenerator()
var userAssignedIdentityProperties_STATUS_ARMGenerator gopter.Gen

// UserAssignedIdentityProperties_STATUS_ARMGenerator returns a generator of UserAssignedIdentityProperties_STATUS_ARM instances for property testing.
func UserAssignedIdentityProperties_STATUS_ARMGenerator() gopter.Gen {
	if userAssignedIdentityProperties_STATUS_ARMGenerator != nil {
		return userAssignedIdentityProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM(generators)
	userAssignedIdentityProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_STATUS_ARM{}), generators)

	return userAssignedIdentityProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
