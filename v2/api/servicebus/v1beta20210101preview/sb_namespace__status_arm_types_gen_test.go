// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

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

func Test_SBNamespace_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBNamespace_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBNamespaceStatusARM, SBNamespaceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBNamespaceStatusARM runs a test to see if a specific instance of SBNamespace_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBNamespaceStatusARM(subject SBNamespace_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBNamespace_StatusARM
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

// Generator of SBNamespace_StatusARM instances for property testing - lazily instantiated by
// SBNamespaceStatusARMGenerator()
var sbNamespaceStatusARMGenerator gopter.Gen

// SBNamespaceStatusARMGenerator returns a generator of SBNamespace_StatusARM instances for property testing.
// We first initialize sbNamespaceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBNamespaceStatusARMGenerator() gopter.Gen {
	if sbNamespaceStatusARMGenerator != nil {
		return sbNamespaceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBNamespaceStatusARM(generators)
	sbNamespaceStatusARMGenerator = gen.Struct(reflect.TypeOf(SBNamespace_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBNamespaceStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBNamespaceStatusARM(generators)
	sbNamespaceStatusARMGenerator = gen.Struct(reflect.TypeOf(SBNamespace_StatusARM{}), generators)

	return sbNamespaceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBNamespaceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBNamespaceStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBNamespaceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBNamespaceStatusARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityStatusARMGenerator())
	gens["Properties"] = gen.PtrOf(SBNamespacePropertiesStatusARMGenerator())
	gens["Sku"] = gen.PtrOf(SBSkuStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_Identity_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityStatusARM, IdentityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityStatusARM runs a test to see if a specific instance of Identity_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityStatusARM(subject Identity_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_StatusARM
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

// Generator of Identity_StatusARM instances for property testing - lazily instantiated by IdentityStatusARMGenerator()
var identityStatusARMGenerator gopter.Gen

// IdentityStatusARMGenerator returns a generator of Identity_StatusARM instances for property testing.
// We first initialize identityStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityStatusARMGenerator() gopter.Gen {
	if identityStatusARMGenerator != nil {
		return identityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	AddRelatedPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	return identityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), DictionaryValueStatusARMGenerator())
}

func Test_SBNamespaceProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBNamespaceProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBNamespacePropertiesStatusARM, SBNamespacePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBNamespacePropertiesStatusARM runs a test to see if a specific instance of SBNamespaceProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBNamespacePropertiesStatusARM(subject SBNamespaceProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBNamespaceProperties_StatusARM
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

// Generator of SBNamespaceProperties_StatusARM instances for property testing - lazily instantiated by
// SBNamespacePropertiesStatusARMGenerator()
var sbNamespacePropertiesStatusARMGenerator gopter.Gen

// SBNamespacePropertiesStatusARMGenerator returns a generator of SBNamespaceProperties_StatusARM instances for property testing.
// We first initialize sbNamespacePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBNamespacePropertiesStatusARMGenerator() gopter.Gen {
	if sbNamespacePropertiesStatusARMGenerator != nil {
		return sbNamespacePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBNamespacePropertiesStatusARM(generators)
	sbNamespacePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBNamespaceProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBNamespacePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForSBNamespacePropertiesStatusARM(generators)
	sbNamespacePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(SBNamespaceProperties_StatusARM{}), generators)

	return sbNamespacePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBNamespacePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBNamespacePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["MetricId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceBusEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForSBNamespacePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBNamespacePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionStatusARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator())
}

func Test_SBSku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBSku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBSkuStatusARM, SBSkuStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBSkuStatusARM runs a test to see if a specific instance of SBSku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBSkuStatusARM(subject SBSku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBSku_StatusARM
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

// Generator of SBSku_StatusARM instances for property testing - lazily instantiated by SBSkuStatusARMGenerator()
var sbSkuStatusARMGenerator gopter.Gen

// SBSkuStatusARMGenerator returns a generator of SBSku_StatusARM instances for property testing.
func SBSkuStatusARMGenerator() gopter.Gen {
	if sbSkuStatusARMGenerator != nil {
		return sbSkuStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBSkuStatusARM(generators)
	sbSkuStatusARMGenerator = gen.Struct(reflect.TypeOf(SBSku_StatusARM{}), generators)

	return sbSkuStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSBSkuStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBSkuStatusARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatusARM, SystemDataStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatusARM runs a test to see if a specific instance of SystemData_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatusARM(subject SystemData_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_StatusARM
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

// Generator of SystemData_StatusARM instances for property testing - lazily instantiated by
// SystemDataStatusARMGenerator()
var systemDataStatusARMGenerator gopter.Gen

// SystemDataStatusARMGenerator returns a generator of SystemData_StatusARM instances for property testing.
func SystemDataStatusARMGenerator() gopter.Gen {
	if systemDataStatusARMGenerator != nil {
		return systemDataStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatusARM(generators)
	systemDataStatusARMGenerator = gen.Struct(reflect.TypeOf(SystemData_StatusARM{}), generators)

	return systemDataStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}

func Test_DictionaryValue_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DictionaryValue_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDictionaryValueStatusARM, DictionaryValueStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDictionaryValueStatusARM runs a test to see if a specific instance of DictionaryValue_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDictionaryValueStatusARM(subject DictionaryValue_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DictionaryValue_StatusARM
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

// Generator of DictionaryValue_StatusARM instances for property testing - lazily instantiated by
// DictionaryValueStatusARMGenerator()
var dictionaryValueStatusARMGenerator gopter.Gen

// DictionaryValueStatusARMGenerator returns a generator of DictionaryValue_StatusARM instances for property testing.
func DictionaryValueStatusARMGenerator() gopter.Gen {
	if dictionaryValueStatusARMGenerator != nil {
		return dictionaryValueStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDictionaryValueStatusARM(generators)
	dictionaryValueStatusARMGenerator = gen.Struct(reflect.TypeOf(DictionaryValue_StatusARM{}), generators)

	return dictionaryValueStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDictionaryValueStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDictionaryValueStatusARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Encryption_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionStatusARM, EncryptionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionStatusARM runs a test to see if a specific instance of Encryption_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionStatusARM(subject Encryption_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_StatusARM
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

// Generator of Encryption_StatusARM instances for property testing - lazily instantiated by
// EncryptionStatusARMGenerator()
var encryptionStatusARMGenerator gopter.Gen

// EncryptionStatusARMGenerator returns a generator of Encryption_StatusARM instances for property testing.
// We first initialize encryptionStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionStatusARMGenerator() gopter.Gen {
	if encryptionStatusARMGenerator != nil {
		return encryptionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionStatusARM(generators)
	encryptionStatusARMGenerator = gen.Struct(reflect.TypeOf(Encryption_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionStatusARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionStatusARM(generators)
	encryptionStatusARMGenerator = gen.Struct(reflect.TypeOf(Encryption_StatusARM{}), generators)

	return encryptionStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionStatusARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.AlphaString())
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionStatusARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultPropertiesStatusARMGenerator())
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM, PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(subject PrivateEndpointConnection_Status_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_Status_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_Status_SubResourceEmbeddedARM instances for property testing.
// We first initialize privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointConnectionStatusSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(generators)
	privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(generators)
	privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionStatusSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_KeyVaultProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesStatusARM, KeyVaultPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesStatusARM runs a test to see if a specific instance of KeyVaultProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesStatusARM(subject KeyVaultProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_StatusARM
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

// Generator of KeyVaultProperties_StatusARM instances for property testing - lazily instantiated by
// KeyVaultPropertiesStatusARMGenerator()
var keyVaultPropertiesStatusARMGenerator gopter.Gen

// KeyVaultPropertiesStatusARMGenerator returns a generator of KeyVaultProperties_StatusARM instances for property testing.
// We first initialize keyVaultPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultPropertiesStatusARMGenerator() gopter.Gen {
	if keyVaultPropertiesStatusARMGenerator != nil {
		return keyVaultPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM(generators)
	keyVaultPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultPropertiesStatusARM(generators)
	keyVaultPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_StatusARM{}), generators)

	return keyVaultPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityPropertiesStatusARMGenerator())
}

func Test_UserAssignedIdentityProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM, UserAssignedIdentityPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM runs a test to see if a specific instance of UserAssignedIdentityProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM(subject UserAssignedIdentityProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_StatusARM
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

// Generator of UserAssignedIdentityProperties_StatusARM instances for property testing - lazily instantiated by
// UserAssignedIdentityPropertiesStatusARMGenerator()
var userAssignedIdentityPropertiesStatusARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesStatusARMGenerator returns a generator of UserAssignedIdentityProperties_StatusARM instances for property testing.
func UserAssignedIdentityPropertiesStatusARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesStatusARMGenerator != nil {
		return userAssignedIdentityPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM(generators)
	userAssignedIdentityPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_StatusARM{}), generators)

	return userAssignedIdentityPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
