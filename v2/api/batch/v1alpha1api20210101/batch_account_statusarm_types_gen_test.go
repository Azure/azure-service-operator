// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

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

func Test_BatchAccount_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccount_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccount_STATUSARM, BatchAccount_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccount_STATUSARM runs a test to see if a specific instance of BatchAccount_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccount_STATUSARM(subject BatchAccount_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccount_STATUSARM
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

// Generator of BatchAccount_STATUSARM instances for property testing - lazily instantiated by
// BatchAccount_STATUSARMGenerator()
var batchAccount_STATUSARMGenerator gopter.Gen

// BatchAccount_STATUSARMGenerator returns a generator of BatchAccount_STATUSARM instances for property testing.
// We first initialize batchAccount_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccount_STATUSARMGenerator() gopter.Gen {
	if batchAccount_STATUSARMGenerator != nil {
		return batchAccount_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_STATUSARM(generators)
	batchAccount_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccount_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccount_STATUSARM(generators)
	batchAccount_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccount_STATUSARM{}), generators)

	return batchAccount_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccount_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccount_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBatchAccount_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccount_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(BatchAccountIdentity_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(BatchAccountCreateProperties_STATUSARMGenerator())
}

func Test_BatchAccountCreateProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountCreateProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountCreateProperties_STATUSARM, BatchAccountCreateProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountCreateProperties_STATUSARM runs a test to see if a specific instance of BatchAccountCreateProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountCreateProperties_STATUSARM(subject BatchAccountCreateProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountCreateProperties_STATUSARM
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

// Generator of BatchAccountCreateProperties_STATUSARM instances for property testing - lazily instantiated by
// BatchAccountCreateProperties_STATUSARMGenerator()
var batchAccountCreateProperties_STATUSARMGenerator gopter.Gen

// BatchAccountCreateProperties_STATUSARMGenerator returns a generator of BatchAccountCreateProperties_STATUSARM instances for property testing.
// We first initialize batchAccountCreateProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountCreateProperties_STATUSARMGenerator() gopter.Gen {
	if batchAccountCreateProperties_STATUSARMGenerator != nil {
		return batchAccountCreateProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM(generators)
	batchAccountCreateProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM(generators)
	batchAccountCreateProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountCreateProperties_STATUSARM{}), generators)

	return batchAccountCreateProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["PoolAllocationMode"] = gen.PtrOf(gen.OneConstOf(PoolAllocationMode_BatchService_STATUS, PoolAllocationMode_UserSubscription_STATUS))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_Disabled_STATUS, PublicNetworkAccessType_Enabled_STATUS))
}

// AddRelatedPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountCreateProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AutoStorage"] = gen.PtrOf(AutoStorageBaseProperties_STATUSARMGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionProperties_STATUSARMGenerator())
	gens["KeyVaultReference"] = gen.PtrOf(KeyVaultReference_STATUSARMGenerator())
}

func Test_BatchAccountIdentity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentity_STATUSARM, BatchAccountIdentity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentity_STATUSARM runs a test to see if a specific instance of BatchAccountIdentity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentity_STATUSARM(subject BatchAccountIdentity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_STATUSARM
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

// Generator of BatchAccountIdentity_STATUSARM instances for property testing - lazily instantiated by
// BatchAccountIdentity_STATUSARMGenerator()
var batchAccountIdentity_STATUSARMGenerator gopter.Gen

// BatchAccountIdentity_STATUSARMGenerator returns a generator of BatchAccountIdentity_STATUSARM instances for property testing.
// We first initialize batchAccountIdentity_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountIdentity_STATUSARMGenerator() gopter.Gen {
	if batchAccountIdentity_STATUSARMGenerator != nil {
		return batchAccountIdentity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_STATUSARM(generators)
	batchAccountIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForBatchAccountIdentity_STATUSARM(generators)
	batchAccountIdentity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_STATUSARM{}), generators)

	return batchAccountIdentity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(BatchAccountIdentity_Type_None_STATUS, BatchAccountIdentity_Type_SystemAssigned_STATUS, BatchAccountIdentity_Type_UserAssigned_STATUS))
}

// AddRelatedPropertyGeneratorsForBatchAccountIdentity_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), BatchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator())
}

func Test_AutoStorageBaseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoStorageBaseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoStorageBaseProperties_STATUSARM, AutoStorageBaseProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoStorageBaseProperties_STATUSARM runs a test to see if a specific instance of AutoStorageBaseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoStorageBaseProperties_STATUSARM(subject AutoStorageBaseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoStorageBaseProperties_STATUSARM
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

// Generator of AutoStorageBaseProperties_STATUSARM instances for property testing - lazily instantiated by
// AutoStorageBaseProperties_STATUSARMGenerator()
var autoStorageBaseProperties_STATUSARMGenerator gopter.Gen

// AutoStorageBaseProperties_STATUSARMGenerator returns a generator of AutoStorageBaseProperties_STATUSARM instances for property testing.
func AutoStorageBaseProperties_STATUSARMGenerator() gopter.Gen {
	if autoStorageBaseProperties_STATUSARMGenerator != nil {
		return autoStorageBaseProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_STATUSARM(generators)
	autoStorageBaseProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AutoStorageBaseProperties_STATUSARM{}), generators)

	return autoStorageBaseProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["StorageAccountId"] = gen.PtrOf(gen.AlphaString())
}

func Test_BatchAccountIdentity_UserAssignedIdentities_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_UserAssignedIdentities_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentity_UserAssignedIdentities_STATUSARM, BatchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentity_UserAssignedIdentities_STATUSARM runs a test to see if a specific instance of BatchAccountIdentity_UserAssignedIdentities_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentity_UserAssignedIdentities_STATUSARM(subject BatchAccountIdentity_UserAssignedIdentities_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_UserAssignedIdentities_STATUSARM
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

// Generator of BatchAccountIdentity_UserAssignedIdentities_STATUSARM instances for property testing - lazily
// instantiated by BatchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator()
var batchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator gopter.Gen

// BatchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator returns a generator of BatchAccountIdentity_UserAssignedIdentities_STATUSARM instances for property testing.
func BatchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator() gopter.Gen {
	if batchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator != nil {
		return batchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_UserAssignedIdentities_STATUSARM(generators)
	batchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_UserAssignedIdentities_STATUSARM{}), generators)

	return batchAccountIdentity_UserAssignedIdentities_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentity_UserAssignedIdentities_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentity_UserAssignedIdentities_STATUSARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_EncryptionProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperties_STATUSARM, EncryptionProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperties_STATUSARM runs a test to see if a specific instance of EncryptionProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperties_STATUSARM(subject EncryptionProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperties_STATUSARM
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

// Generator of EncryptionProperties_STATUSARM instances for property testing - lazily instantiated by
// EncryptionProperties_STATUSARMGenerator()
var encryptionProperties_STATUSARMGenerator gopter.Gen

// EncryptionProperties_STATUSARMGenerator returns a generator of EncryptionProperties_STATUSARM instances for property testing.
// We first initialize encryptionProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionProperties_STATUSARMGenerator() gopter.Gen {
	if encryptionProperties_STATUSARMGenerator != nil {
		return encryptionProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties_STATUSARM(generators)
	encryptionProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionProperties_STATUSARM(generators)
	encryptionProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_STATUSARM{}), generators)

	return encryptionProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(EncryptionProperties_KeySource_MicrosoftBatch_STATUS, EncryptionProperties_KeySource_MicrosoftKeyVault_STATUS))
}

// AddRelatedPropertyGeneratorsForEncryptionProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUSARMGenerator())
}

func Test_KeyVaultReference_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultReference_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultReference_STATUSARM, KeyVaultReference_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultReference_STATUSARM runs a test to see if a specific instance of KeyVaultReference_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultReference_STATUSARM(subject KeyVaultReference_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultReference_STATUSARM
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

// Generator of KeyVaultReference_STATUSARM instances for property testing - lazily instantiated by
// KeyVaultReference_STATUSARMGenerator()
var keyVaultReference_STATUSARMGenerator gopter.Gen

// KeyVaultReference_STATUSARMGenerator returns a generator of KeyVaultReference_STATUSARM instances for property testing.
func KeyVaultReference_STATUSARMGenerator() gopter.Gen {
	if keyVaultReference_STATUSARMGenerator != nil {
		return keyVaultReference_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultReference_STATUSARM(generators)
	keyVaultReference_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultReference_STATUSARM{}), generators)

	return keyVaultReference_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultReference_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultReference_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
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
func KeyVaultProperties_STATUSARMGenerator() gopter.Gen {
	if keyVaultProperties_STATUSARMGenerator != nil {
		return keyVaultProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(generators)
	keyVaultProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_STATUSARM{}), generators)

	return keyVaultProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}
