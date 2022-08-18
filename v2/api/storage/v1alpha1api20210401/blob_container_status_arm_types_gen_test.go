// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_BlobContainer_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobContainer_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobContainerSTATUSARM, BlobContainerSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobContainerSTATUSARM runs a test to see if a specific instance of BlobContainer_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobContainerSTATUSARM(subject BlobContainer_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobContainer_STATUSARM
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

// Generator of BlobContainer_STATUSARM instances for property testing - lazily instantiated by
// BlobContainerSTATUSARMGenerator()
var blobContainerSTATUSARMGenerator gopter.Gen

// BlobContainerSTATUSARMGenerator returns a generator of BlobContainer_STATUSARM instances for property testing.
// We first initialize blobContainerSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobContainerSTATUSARMGenerator() gopter.Gen {
	if blobContainerSTATUSARMGenerator != nil {
		return blobContainerSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerSTATUSARM(generators)
	blobContainerSTATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForBlobContainerSTATUSARM(generators)
	blobContainerSTATUSARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_STATUSARM{}), generators)

	return blobContainerSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobContainerSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobContainerSTATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobContainerSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobContainerSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerPropertiesSTATUSARMGenerator())
}

func Test_ContainerProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPropertiesSTATUSARM, ContainerPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPropertiesSTATUSARM runs a test to see if a specific instance of ContainerProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPropertiesSTATUSARM(subject ContainerProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerProperties_STATUSARM
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

// Generator of ContainerProperties_STATUSARM instances for property testing - lazily instantiated by
// ContainerPropertiesSTATUSARMGenerator()
var containerPropertiesSTATUSARMGenerator gopter.Gen

// ContainerPropertiesSTATUSARMGenerator returns a generator of ContainerProperties_STATUSARM instances for property testing.
// We first initialize containerPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerPropertiesSTATUSARMGenerator() gopter.Gen {
	if containerPropertiesSTATUSARMGenerator != nil {
		return containerPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesSTATUSARM(generators)
	containerPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForContainerPropertiesSTATUSARM(generators)
	containerPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUSARM{}), generators)

	return containerPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesSTATUSLeaseDuration_Fixed, ContainerPropertiesSTATUSLeaseDuration_Infinite))
	gens["LeaseState"] = gen.PtrOf(gen.OneConstOf(
		ContainerPropertiesSTATUSLeaseState_Available,
		ContainerPropertiesSTATUSLeaseState_Breaking,
		ContainerPropertiesSTATUSLeaseState_Broken,
		ContainerPropertiesSTATUSLeaseState_Expired,
		ContainerPropertiesSTATUSLeaseState_Leased))
	gens["LeaseStatus"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesSTATUSLeaseStatus_Locked, ContainerPropertiesSTATUSLeaseStatus_Unlocked))
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesSTATUSPublicAccess_Blob, ContainerPropertiesSTATUSPublicAccess_Container, ContainerPropertiesSTATUSPublicAccess_None))
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForContainerPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyPropertiesSTATUSARMGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningSTATUSARMGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldPropertiesSTATUSARMGenerator())
}

func Test_ImmutabilityPolicyProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyPropertiesSTATUSARM, ImmutabilityPolicyPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyPropertiesSTATUSARM runs a test to see if a specific instance of ImmutabilityPolicyProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyPropertiesSTATUSARM(subject ImmutabilityPolicyProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_STATUSARM
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

// Generator of ImmutabilityPolicyProperties_STATUSARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyPropertiesSTATUSARMGenerator()
var immutabilityPolicyPropertiesSTATUSARMGenerator gopter.Gen

// ImmutabilityPolicyPropertiesSTATUSARMGenerator returns a generator of ImmutabilityPolicyProperties_STATUSARM instances for property testing.
// We first initialize immutabilityPolicyPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyPropertiesSTATUSARMGenerator() gopter.Gen {
	if immutabilityPolicyPropertiesSTATUSARMGenerator != nil {
		return immutabilityPolicyPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM(generators)
	immutabilityPolicyPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM(generators)
	immutabilityPolicyPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUSARM{}), generators)

	return immutabilityPolicyPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ImmutabilityPolicyPropertySTATUSARMGenerator())
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryPropertySTATUSARMGenerator())
}

func Test_ImmutableStorageWithVersioning_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioningSTATUSARM, ImmutableStorageWithVersioningSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioningSTATUSARM runs a test to see if a specific instance of ImmutableStorageWithVersioning_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioningSTATUSARM(subject ImmutableStorageWithVersioning_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_STATUSARM
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

// Generator of ImmutableStorageWithVersioning_STATUSARM instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioningSTATUSARMGenerator()
var immutableStorageWithVersioningSTATUSARMGenerator gopter.Gen

// ImmutableStorageWithVersioningSTATUSARMGenerator returns a generator of ImmutableStorageWithVersioning_STATUSARM instances for property testing.
func ImmutableStorageWithVersioningSTATUSARMGenerator() gopter.Gen {
	if immutableStorageWithVersioningSTATUSARMGenerator != nil {
		return immutableStorageWithVersioningSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningSTATUSARM(generators)
	immutableStorageWithVersioningSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_STATUSARM{}), generators)

	return immutableStorageWithVersioningSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningSTATUSARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.OneConstOf(ImmutableStorageWithVersioningSTATUSMigrationState_Completed, ImmutableStorageWithVersioningSTATUSMigrationState_InProgress))
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldPropertiesSTATUSARM, LegalHoldPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldPropertiesSTATUSARM runs a test to see if a specific instance of LegalHoldProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldPropertiesSTATUSARM(subject LegalHoldProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_STATUSARM
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

// Generator of LegalHoldProperties_STATUSARM instances for property testing - lazily instantiated by
// LegalHoldPropertiesSTATUSARMGenerator()
var legalHoldPropertiesSTATUSARMGenerator gopter.Gen

// LegalHoldPropertiesSTATUSARMGenerator returns a generator of LegalHoldProperties_STATUSARM instances for property testing.
// We first initialize legalHoldPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldPropertiesSTATUSARMGenerator() gopter.Gen {
	if legalHoldPropertiesSTATUSARMGenerator != nil {
		return legalHoldPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesSTATUSARM(generators)
	legalHoldPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForLegalHoldPropertiesSTATUSARM(generators)
	legalHoldPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUSARM{}), generators)

	return legalHoldPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Tags"] = gen.SliceOf(TagPropertySTATUSARMGenerator())
}

func Test_ImmutabilityPolicyProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyPropertySTATUSARM, ImmutabilityPolicyPropertySTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyPropertySTATUSARM runs a test to see if a specific instance of ImmutabilityPolicyProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyPropertySTATUSARM(subject ImmutabilityPolicyProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperty_STATUSARM
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

// Generator of ImmutabilityPolicyProperty_STATUSARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyPropertySTATUSARMGenerator()
var immutabilityPolicyPropertySTATUSARMGenerator gopter.Gen

// ImmutabilityPolicyPropertySTATUSARMGenerator returns a generator of ImmutabilityPolicyProperty_STATUSARM instances for property testing.
func ImmutabilityPolicyPropertySTATUSARMGenerator() gopter.Gen {
	if immutabilityPolicyPropertySTATUSARMGenerator != nil {
		return immutabilityPolicyPropertySTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertySTATUSARM(generators)
	immutabilityPolicyPropertySTATUSARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperty_STATUSARM{}), generators)

	return immutabilityPolicyPropertySTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertySTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertySTATUSARM(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilityPolicyPropertySTATUSState_Locked, ImmutabilityPolicyPropertySTATUSState_Unlocked))
}

func Test_TagProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagPropertySTATUSARM, TagPropertySTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagPropertySTATUSARM runs a test to see if a specific instance of TagProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagPropertySTATUSARM(subject TagProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_STATUSARM
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

// Generator of TagProperty_STATUSARM instances for property testing - lazily instantiated by
// TagPropertySTATUSARMGenerator()
var tagPropertySTATUSARMGenerator gopter.Gen

// TagPropertySTATUSARMGenerator returns a generator of TagProperty_STATUSARM instances for property testing.
func TagPropertySTATUSARMGenerator() gopter.Gen {
	if tagPropertySTATUSARMGenerator != nil {
		return tagPropertySTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagPropertySTATUSARM(generators)
	tagPropertySTATUSARMGenerator = gen.Struct(reflect.TypeOf(TagProperty_STATUSARM{}), generators)

	return tagPropertySTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForTagPropertySTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagPropertySTATUSARM(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryPropertySTATUSARM, UpdateHistoryPropertySTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryPropertySTATUSARM runs a test to see if a specific instance of UpdateHistoryProperty_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryPropertySTATUSARM(subject UpdateHistoryProperty_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_STATUSARM
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

// Generator of UpdateHistoryProperty_STATUSARM instances for property testing - lazily instantiated by
// UpdateHistoryPropertySTATUSARMGenerator()
var updateHistoryPropertySTATUSARMGenerator gopter.Gen

// UpdateHistoryPropertySTATUSARMGenerator returns a generator of UpdateHistoryProperty_STATUSARM instances for property testing.
func UpdateHistoryPropertySTATUSARMGenerator() gopter.Gen {
	if updateHistoryPropertySTATUSARMGenerator != nil {
		return updateHistoryPropertySTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryPropertySTATUSARM(generators)
	updateHistoryPropertySTATUSARMGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_STATUSARM{}), generators)

	return updateHistoryPropertySTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryPropertySTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryPropertySTATUSARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.OneConstOf(UpdateHistoryPropertySTATUSUpdate_Extend, UpdateHistoryPropertySTATUSUpdate_Lock, UpdateHistoryPropertySTATUSUpdate_Put))
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}
