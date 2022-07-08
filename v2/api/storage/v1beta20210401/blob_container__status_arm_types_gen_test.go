// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

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

func Test_BlobContainer_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobContainer_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobContainerStatusARM, BlobContainerStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobContainerStatusARM runs a test to see if a specific instance of BlobContainer_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobContainerStatusARM(subject BlobContainer_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobContainer_StatusARM
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

// Generator of BlobContainer_StatusARM instances for property testing - lazily instantiated by
// BlobContainerStatusARMGenerator()
var blobContainerStatusARMGenerator gopter.Gen

// BlobContainerStatusARMGenerator returns a generator of BlobContainer_StatusARM instances for property testing.
// We first initialize blobContainerStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobContainerStatusARMGenerator() gopter.Gen {
	if blobContainerStatusARMGenerator != nil {
		return blobContainerStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerStatusARM(generators)
	blobContainerStatusARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerStatusARM(generators)
	AddRelatedPropertyGeneratorsForBlobContainerStatusARM(generators)
	blobContainerStatusARMGenerator = gen.Struct(reflect.TypeOf(BlobContainer_StatusARM{}), generators)

	return blobContainerStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobContainerStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobContainerStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobContainerStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobContainerStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerPropertiesStatusARMGenerator())
}

func Test_ContainerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPropertiesStatusARM, ContainerPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPropertiesStatusARM runs a test to see if a specific instance of ContainerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPropertiesStatusARM(subject ContainerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerProperties_StatusARM
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

// Generator of ContainerProperties_StatusARM instances for property testing - lazily instantiated by
// ContainerPropertiesStatusARMGenerator()
var containerPropertiesStatusARMGenerator gopter.Gen

// ContainerPropertiesStatusARMGenerator returns a generator of ContainerProperties_StatusARM instances for property testing.
// We first initialize containerPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerPropertiesStatusARMGenerator() gopter.Gen {
	if containerPropertiesStatusARMGenerator != nil {
		return containerPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesStatusARM(generators)
	containerPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForContainerPropertiesStatusARM(generators)
	containerPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_StatusARM{}), generators)

	return containerPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesStatusLeaseDurationFixed, ContainerPropertiesStatusLeaseDurationInfinite))
	gens["LeaseState"] = gen.PtrOf(gen.OneConstOf(
		ContainerPropertiesStatusLeaseStateAvailable,
		ContainerPropertiesStatusLeaseStateBreaking,
		ContainerPropertiesStatusLeaseStateBroken,
		ContainerPropertiesStatusLeaseStateExpired,
		ContainerPropertiesStatusLeaseStateLeased))
	gens["LeaseStatus"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesStatusLeaseStatusLocked, ContainerPropertiesStatusLeaseStatusUnlocked))
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerPropertiesStatusPublicAccessBlob, ContainerPropertiesStatusPublicAccessContainer, ContainerPropertiesStatusPublicAccessNone))
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForContainerPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyPropertiesStatusARMGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningStatusARMGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldPropertiesStatusARMGenerator())
}

func Test_ImmutabilityPolicyProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyPropertiesStatusARM, ImmutabilityPolicyPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyPropertiesStatusARM runs a test to see if a specific instance of ImmutabilityPolicyProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyPropertiesStatusARM(subject ImmutabilityPolicyProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_StatusARM
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

// Generator of ImmutabilityPolicyProperties_StatusARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyPropertiesStatusARMGenerator()
var immutabilityPolicyPropertiesStatusARMGenerator gopter.Gen

// ImmutabilityPolicyPropertiesStatusARMGenerator returns a generator of ImmutabilityPolicyProperties_StatusARM instances for property testing.
// We first initialize immutabilityPolicyPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyPropertiesStatusARMGenerator() gopter.Gen {
	if immutabilityPolicyPropertiesStatusARMGenerator != nil {
		return immutabilityPolicyPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM(generators)
	immutabilityPolicyPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM(generators)
	immutabilityPolicyPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_StatusARM{}), generators)

	return immutabilityPolicyPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ImmutabilityPolicyPropertyStatusARMGenerator())
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryPropertyStatusARMGenerator())
}

func Test_ImmutableStorageWithVersioning_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioningStatusARM, ImmutableStorageWithVersioningStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioningStatusARM runs a test to see if a specific instance of ImmutableStorageWithVersioning_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioningStatusARM(subject ImmutableStorageWithVersioning_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_StatusARM
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

// Generator of ImmutableStorageWithVersioning_StatusARM instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioningStatusARMGenerator()
var immutableStorageWithVersioningStatusARMGenerator gopter.Gen

// ImmutableStorageWithVersioningStatusARMGenerator returns a generator of ImmutableStorageWithVersioning_StatusARM instances for property testing.
func ImmutableStorageWithVersioningStatusARMGenerator() gopter.Gen {
	if immutableStorageWithVersioningStatusARMGenerator != nil {
		return immutableStorageWithVersioningStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatusARM(generators)
	immutableStorageWithVersioningStatusARMGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_StatusARM{}), generators)

	return immutableStorageWithVersioningStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.OneConstOf(ImmutableStorageWithVersioningStatusMigrationStateCompleted, ImmutableStorageWithVersioningStatusMigrationStateInProgress))
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldPropertiesStatusARM, LegalHoldPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldPropertiesStatusARM runs a test to see if a specific instance of LegalHoldProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldPropertiesStatusARM(subject LegalHoldProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_StatusARM
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

// Generator of LegalHoldProperties_StatusARM instances for property testing - lazily instantiated by
// LegalHoldPropertiesStatusARMGenerator()
var legalHoldPropertiesStatusARMGenerator gopter.Gen

// LegalHoldPropertiesStatusARMGenerator returns a generator of LegalHoldProperties_StatusARM instances for property testing.
// We first initialize legalHoldPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldPropertiesStatusARMGenerator() gopter.Gen {
	if legalHoldPropertiesStatusARMGenerator != nil {
		return legalHoldPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatusARM(generators)
	legalHoldPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatusARM(generators)
	legalHoldPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_StatusARM{}), generators)

	return legalHoldPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Tags"] = gen.SliceOf(TagPropertyStatusARMGenerator())
}

func Test_ImmutabilityPolicyProperty_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperty_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyPropertyStatusARM, ImmutabilityPolicyPropertyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyPropertyStatusARM runs a test to see if a specific instance of ImmutabilityPolicyProperty_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyPropertyStatusARM(subject ImmutabilityPolicyProperty_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperty_StatusARM
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

// Generator of ImmutabilityPolicyProperty_StatusARM instances for property testing - lazily instantiated by
// ImmutabilityPolicyPropertyStatusARMGenerator()
var immutabilityPolicyPropertyStatusARMGenerator gopter.Gen

// ImmutabilityPolicyPropertyStatusARMGenerator returns a generator of ImmutabilityPolicyProperty_StatusARM instances for property testing.
func ImmutabilityPolicyPropertyStatusARMGenerator() gopter.Gen {
	if immutabilityPolicyPropertyStatusARMGenerator != nil {
		return immutabilityPolicyPropertyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertyStatusARM(generators)
	immutabilityPolicyPropertyStatusARMGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperty_StatusARM{}), generators)

	return immutabilityPolicyPropertyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertyStatusARM(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilityPolicyPropertyStatusStateLocked, ImmutabilityPolicyPropertyStatusStateUnlocked))
}

func Test_TagProperty_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagPropertyStatusARM, TagPropertyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagPropertyStatusARM runs a test to see if a specific instance of TagProperty_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagPropertyStatusARM(subject TagProperty_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_StatusARM
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

// Generator of TagProperty_StatusARM instances for property testing - lazily instantiated by
// TagPropertyStatusARMGenerator()
var tagPropertyStatusARMGenerator gopter.Gen

// TagPropertyStatusARMGenerator returns a generator of TagProperty_StatusARM instances for property testing.
func TagPropertyStatusARMGenerator() gopter.Gen {
	if tagPropertyStatusARMGenerator != nil {
		return tagPropertyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagPropertyStatusARM(generators)
	tagPropertyStatusARMGenerator = gen.Struct(reflect.TypeOf(TagProperty_StatusARM{}), generators)

	return tagPropertyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForTagPropertyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagPropertyStatusARM(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryPropertyStatusARM, UpdateHistoryPropertyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryPropertyStatusARM runs a test to see if a specific instance of UpdateHistoryProperty_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryPropertyStatusARM(subject UpdateHistoryProperty_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_StatusARM
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

// Generator of UpdateHistoryProperty_StatusARM instances for property testing - lazily instantiated by
// UpdateHistoryPropertyStatusARMGenerator()
var updateHistoryPropertyStatusARMGenerator gopter.Gen

// UpdateHistoryPropertyStatusARMGenerator returns a generator of UpdateHistoryProperty_StatusARM instances for property testing.
func UpdateHistoryPropertyStatusARMGenerator() gopter.Gen {
	if updateHistoryPropertyStatusARMGenerator != nil {
		return updateHistoryPropertyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatusARM(generators)
	updateHistoryPropertyStatusARMGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_StatusARM{}), generators)

	return updateHistoryPropertyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatusARM(gens map[string]gopter.Gen) {
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.OneConstOf(UpdateHistoryPropertyStatusUpdateExtend, UpdateHistoryPropertyStatusUpdateLock, UpdateHistoryPropertyStatusUpdatePut))
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}
