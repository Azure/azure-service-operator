// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_Namespaces_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesSpecARM, NamespacesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesSpecARM runs a test to see if a specific instance of Namespaces_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesSpecARM(subject Namespaces_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_SpecARM
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

// Generator of Namespaces_SpecARM instances for property testing - lazily instantiated by NamespacesSpecARMGenerator()
var namespacesSpecARMGenerator gopter.Gen

// NamespacesSpecARMGenerator returns a generator of Namespaces_SpecARM instances for property testing.
// We first initialize namespacesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesSpecARMGenerator() gopter.Gen {
	if namespacesSpecARMGenerator != nil {
		return namespacesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSpecARM(generators)
	namespacesSpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesSpecARM(generators)
	namespacesSpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_SpecARM{}), generators)

	return namespacesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesSpecARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityARMGenerator())
	gens["Properties"] = gen.PtrOf(NamespacesSpecPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
}

func Test_IdentityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityARM, IdentityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityARM runs a test to see if a specific instance of IdentityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityARM(subject IdentityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityARM
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

// Generator of IdentityARM instances for property testing - lazily instantiated by IdentityARMGenerator()
var identityARMGenerator gopter.Gen

// IdentityARMGenerator returns a generator of IdentityARM instances for property testing.
func IdentityARMGenerator() gopter.Gen {
	if identityARMGenerator != nil {
		return identityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityARM(generators)
	identityARMGenerator = gen.Struct(reflect.TypeOf(IdentityARM{}), generators)

	return identityARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned))
}

func Test_Namespaces_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesSpecPropertiesARM, NamespacesSpecPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesSpecPropertiesARM runs a test to see if a specific instance of Namespaces_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesSpecPropertiesARM(subject Namespaces_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Spec_PropertiesARM
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

// Generator of Namespaces_Spec_PropertiesARM instances for property testing - lazily instantiated by
// NamespacesSpecPropertiesARMGenerator()
var namespacesSpecPropertiesARMGenerator gopter.Gen

// NamespacesSpecPropertiesARMGenerator returns a generator of Namespaces_Spec_PropertiesARM instances for property testing.
// We first initialize namespacesSpecPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesSpecPropertiesARMGenerator() gopter.Gen {
	if namespacesSpecPropertiesARMGenerator != nil {
		return namespacesSpecPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSpecPropertiesARM(generators)
	namespacesSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesSpecPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesARM(generators)
	namespacesSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_PropertiesARM{}), generators)

	return namespacesSpecPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesSpecPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["AlternateName"] = gen.PtrOf(gen.AlphaString())
	gens["ClusterArmId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAutoInflateEnabled"] = gen.PtrOf(gen.Bool())
	gens["KafkaEnabled"] = gen.PtrOf(gen.Bool())
	gens["MaximumThroughputUnits"] = gen.PtrOf(gen.Int())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(NamespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator())
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
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

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SkuNameBasic, SkuNamePremium, SkuNameStandard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(SkuTierBasic, SkuTierPremium, SkuTierStandard))
}

func Test_EncryptionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionARM, EncryptionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionARM runs a test to see if a specific instance of EncryptionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionARM(subject EncryptionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionARM
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

// Generator of EncryptionARM instances for property testing - lazily instantiated by EncryptionARMGenerator()
var encryptionARMGenerator gopter.Gen

// EncryptionARMGenerator returns a generator of EncryptionARM instances for property testing.
// We first initialize encryptionARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionARMGenerator() gopter.Gen {
	if encryptionARMGenerator != nil {
		return encryptionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionARM(generators)
	encryptionARMGenerator = gen.Struct(reflect.TypeOf(EncryptionARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionARM(generators)
	encryptionARMGenerator = gen.Struct(reflect.TypeOf(EncryptionARM{}), generators)

	return encryptionARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(EncryptionKeySourceMicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultPropertiesARMGenerator())
}

func Test_Namespaces_Spec_Properties_PrivateEndpointConnectionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesSpecPropertiesPrivateEndpointConnectionsARM, NamespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesSpecPropertiesPrivateEndpointConnectionsARM runs a test to see if a specific instance of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesSpecPropertiesPrivateEndpointConnectionsARM(subject Namespaces_Spec_Properties_PrivateEndpointConnectionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Spec_Properties_PrivateEndpointConnectionsARM
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

// Generator of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM instances for property testing - lazily
// instantiated by NamespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator()
var namespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator gopter.Gen

// NamespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator returns a generator of Namespaces_Spec_Properties_PrivateEndpointConnectionsARM instances for property testing.
func NamespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator() gopter.Gen {
	if namespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator != nil {
		return namespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesPrivateEndpointConnectionsARM(generators)
	namespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Spec_Properties_PrivateEndpointConnectionsARM{}), generators)

	return namespacesSpecPropertiesPrivateEndpointConnectionsARMGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesPrivateEndpointConnectionsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesSpecPropertiesPrivateEndpointConnectionsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateEndpointConnectionPropertiesARMGenerator())
}

func Test_KeyVaultPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultPropertiesARM, KeyVaultPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultPropertiesARM runs a test to see if a specific instance of KeyVaultPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultPropertiesARM(subject KeyVaultPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultPropertiesARM
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

// Generator of KeyVaultPropertiesARM instances for property testing - lazily instantiated by
// KeyVaultPropertiesARMGenerator()
var keyVaultPropertiesARMGenerator gopter.Gen

// KeyVaultPropertiesARMGenerator returns a generator of KeyVaultPropertiesARM instances for property testing.
// We first initialize keyVaultPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultPropertiesARMGenerator() gopter.Gen {
	if keyVaultPropertiesARMGenerator != nil {
		return keyVaultPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM(generators)
	keyVaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultPropertiesARM{}), generators)

	return keyVaultPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityPropertiesARMGenerator())
}

func Test_PrivateEndpointConnectionPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnectionPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM, PrivateEndpointConnectionPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM runs a test to see if a specific instance of PrivateEndpointConnectionPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionPropertiesARM(subject PrivateEndpointConnectionPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnectionPropertiesARM
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

// Generator of PrivateEndpointConnectionPropertiesARM instances for property testing - lazily instantiated by
// PrivateEndpointConnectionPropertiesARMGenerator()
var privateEndpointConnectionPropertiesARMGenerator gopter.Gen

// PrivateEndpointConnectionPropertiesARMGenerator returns a generator of PrivateEndpointConnectionPropertiesARM instances for property testing.
func PrivateEndpointConnectionPropertiesARMGenerator() gopter.Gen {
	if privateEndpointConnectionPropertiesARMGenerator != nil {
		return privateEndpointConnectionPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM(generators)
	privateEndpointConnectionPropertiesARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnectionPropertiesARM{}), generators)

	return privateEndpointConnectionPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointConnectionPropertiesARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpoint"] = gen.PtrOf(PrivateEndpointARMGenerator())
}

func Test_PrivateEndpointARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointARM, PrivateEndpointARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointARM runs a test to see if a specific instance of PrivateEndpointARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointARM(subject PrivateEndpointARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointARM
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

// Generator of PrivateEndpointARM instances for property testing - lazily instantiated by PrivateEndpointARMGenerator()
var privateEndpointARMGenerator gopter.Gen

// PrivateEndpointARMGenerator returns a generator of PrivateEndpointARM instances for property testing.
func PrivateEndpointARMGenerator() gopter.Gen {
	if privateEndpointARMGenerator != nil {
		return privateEndpointARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointARM(generators)
	privateEndpointARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointARM{}), generators)

	return privateEndpointARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesARM, UserAssignedIdentityPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesARM runs a test to see if a specific instance of UserAssignedIdentityPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesARM(subject UserAssignedIdentityPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityPropertiesARM
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

// Generator of UserAssignedIdentityPropertiesARM instances for property testing - lazily instantiated by
// UserAssignedIdentityPropertiesARMGenerator()
var userAssignedIdentityPropertiesARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesARMGenerator returns a generator of UserAssignedIdentityPropertiesARM instances for property testing.
func UserAssignedIdentityPropertiesARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesARMGenerator != nil {
		return userAssignedIdentityPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM(generators)
	userAssignedIdentityPropertiesARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityPropertiesARM{}), generators)

	return userAssignedIdentityPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
