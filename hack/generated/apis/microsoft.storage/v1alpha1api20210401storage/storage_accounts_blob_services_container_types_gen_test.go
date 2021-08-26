// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

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

func Test_StorageAccountsBlobServicesContainer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainer, StorageAccountsBlobServicesContainerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainer runs a test to see if a specific instance of StorageAccountsBlobServicesContainer round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainer(subject StorageAccountsBlobServicesContainer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainer
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

// Generator of StorageAccountsBlobServicesContainer instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesContainerGenerator()
var storageAccountsBlobServicesContainerGenerator gopter.Gen

// StorageAccountsBlobServicesContainerGenerator returns a generator of StorageAccountsBlobServicesContainer instances for property testing.
func StorageAccountsBlobServicesContainerGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainerGenerator != nil {
		return storageAccountsBlobServicesContainerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(generators)
	storageAccountsBlobServicesContainerGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer{}), generators)

	return storageAccountsBlobServicesContainerGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsBlobServicesContainersSpecGenerator()
	gens["Status"] = BlobContainerStatusGenerator()
}

func Test_BlobContainer_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobContainer_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobContainerStatus, BlobContainerStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobContainerStatus runs a test to see if a specific instance of BlobContainer_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobContainerStatus(subject BlobContainer_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobContainer_Status
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

// Generator of BlobContainer_Status instances for property testing - lazily instantiated by
//BlobContainerStatusGenerator()
var blobContainerStatusGenerator gopter.Gen

// BlobContainerStatusGenerator returns a generator of BlobContainer_Status instances for property testing.
// We first initialize blobContainerStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobContainerStatusGenerator() gopter.Gen {
	if blobContainerStatusGenerator != nil {
		return blobContainerStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerStatus(generators)
	blobContainerStatusGenerator = gen.Struct(reflect.TypeOf(BlobContainer_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainerStatus(generators)
	AddRelatedPropertyGeneratorsForBlobContainerStatus(generators)
	blobContainerStatusGenerator = gen.Struct(reflect.TypeOf(BlobContainer_Status{}), generators)

	return blobContainerStatusGenerator
}

// AddIndependentPropertyGeneratorsForBlobContainerStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobContainerStatus(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseState"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobContainerStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobContainerStatus(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyPropertiesStatusGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningStatusGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldPropertiesStatusGenerator())
}

func Test_StorageAccountsBlobServicesContainers_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainers_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainersSpec, StorageAccountsBlobServicesContainersSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainersSpec runs a test to see if a specific instance of StorageAccountsBlobServicesContainers_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainersSpec(subject StorageAccountsBlobServicesContainers_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainers_Spec
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

// Generator of StorageAccountsBlobServicesContainers_Spec instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesContainersSpecGenerator()
var storageAccountsBlobServicesContainersSpecGenerator gopter.Gen

// StorageAccountsBlobServicesContainersSpecGenerator returns a generator of StorageAccountsBlobServicesContainers_Spec instances for property testing.
// We first initialize storageAccountsBlobServicesContainersSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesContainersSpecGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainersSpecGenerator != nil {
		return storageAccountsBlobServicesContainersSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec(generators)
	storageAccountsBlobServicesContainersSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainers_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec(generators)
	storageAccountsBlobServicesContainersSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainers_Spec{}), generators)

	return storageAccountsBlobServicesContainersSpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainersSpec(gens map[string]gopter.Gen) {
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningGenerator())
}

func Test_ImmutabilityPolicyProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyPropertiesStatus, ImmutabilityPolicyPropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyPropertiesStatus runs a test to see if a specific instance of ImmutabilityPolicyProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyPropertiesStatus(subject ImmutabilityPolicyProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_Status
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

// Generator of ImmutabilityPolicyProperties_Status instances for property testing - lazily instantiated by
//ImmutabilityPolicyPropertiesStatusGenerator()
var immutabilityPolicyPropertiesStatusGenerator gopter.Gen

// ImmutabilityPolicyPropertiesStatusGenerator returns a generator of ImmutabilityPolicyProperties_Status instances for property testing.
// We first initialize immutabilityPolicyPropertiesStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyPropertiesStatusGenerator() gopter.Gen {
	if immutabilityPolicyPropertiesStatusGenerator != nil {
		return immutabilityPolicyPropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatus(generators)
	immutabilityPolicyPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatus(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatus(generators)
	immutabilityPolicyPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_Status{}), generators)

	return immutabilityPolicyPropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyPropertiesStatus(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyPropertiesStatus(gens map[string]gopter.Gen) {
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryPropertyStatusGenerator())
}

func Test_ImmutableStorageWithVersioning_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning, ImmutableStorageWithVersioningGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning runs a test to see if a specific instance of ImmutableStorageWithVersioning round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning(subject ImmutableStorageWithVersioning) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning
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

// Generator of ImmutableStorageWithVersioning instances for property testing - lazily instantiated by
//ImmutableStorageWithVersioningGenerator()
var immutableStorageWithVersioningGenerator gopter.Gen

// ImmutableStorageWithVersioningGenerator returns a generator of ImmutableStorageWithVersioning instances for property testing.
func ImmutableStorageWithVersioningGenerator() gopter.Gen {
	if immutableStorageWithVersioningGenerator != nil {
		return immutableStorageWithVersioningGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning(generators)
	immutableStorageWithVersioningGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning{}), generators)

	return immutableStorageWithVersioningGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ImmutableStorageWithVersioning_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioningStatus, ImmutableStorageWithVersioningStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioningStatus runs a test to see if a specific instance of ImmutableStorageWithVersioning_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioningStatus(subject ImmutableStorageWithVersioning_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_Status
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

// Generator of ImmutableStorageWithVersioning_Status instances for property testing - lazily instantiated by
//ImmutableStorageWithVersioningStatusGenerator()
var immutableStorageWithVersioningStatusGenerator gopter.Gen

// ImmutableStorageWithVersioningStatusGenerator returns a generator of ImmutableStorageWithVersioning_Status instances for property testing.
func ImmutableStorageWithVersioningStatusGenerator() gopter.Gen {
	if immutableStorageWithVersioningStatusGenerator != nil {
		return immutableStorageWithVersioningStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatus(generators)
	immutableStorageWithVersioningStatusGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_Status{}), generators)

	return immutableStorageWithVersioningStatusGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioningStatus(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.AlphaString())
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldPropertiesStatus, LegalHoldPropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldPropertiesStatus runs a test to see if a specific instance of LegalHoldProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldPropertiesStatus(subject LegalHoldProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_Status
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

// Generator of LegalHoldProperties_Status instances for property testing - lazily instantiated by
//LegalHoldPropertiesStatusGenerator()
var legalHoldPropertiesStatusGenerator gopter.Gen

// LegalHoldPropertiesStatusGenerator returns a generator of LegalHoldProperties_Status instances for property testing.
// We first initialize legalHoldPropertiesStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldPropertiesStatusGenerator() gopter.Gen {
	if legalHoldPropertiesStatusGenerator != nil {
		return legalHoldPropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatus(generators)
	legalHoldPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatus(generators)
	AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatus(generators)
	legalHoldPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_Status{}), generators)

	return legalHoldPropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldPropertiesStatus(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldPropertiesStatus(gens map[string]gopter.Gen) {
	gens["Tags"] = gen.SliceOf(TagPropertyStatusGenerator())
}

func Test_TagProperty_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagPropertyStatus, TagPropertyStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagPropertyStatus runs a test to see if a specific instance of TagProperty_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForTagPropertyStatus(subject TagProperty_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_Status
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

// Generator of TagProperty_Status instances for property testing - lazily instantiated by TagPropertyStatusGenerator()
var tagPropertyStatusGenerator gopter.Gen

// TagPropertyStatusGenerator returns a generator of TagProperty_Status instances for property testing.
func TagPropertyStatusGenerator() gopter.Gen {
	if tagPropertyStatusGenerator != nil {
		return tagPropertyStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagPropertyStatus(generators)
	tagPropertyStatusGenerator = gen.Struct(reflect.TypeOf(TagProperty_Status{}), generators)

	return tagPropertyStatusGenerator
}

// AddIndependentPropertyGeneratorsForTagPropertyStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagPropertyStatus(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryPropertyStatus, UpdateHistoryPropertyStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryPropertyStatus runs a test to see if a specific instance of UpdateHistoryProperty_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryPropertyStatus(subject UpdateHistoryProperty_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_Status
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

// Generator of UpdateHistoryProperty_Status instances for property testing - lazily instantiated by
//UpdateHistoryPropertyStatusGenerator()
var updateHistoryPropertyStatusGenerator gopter.Gen

// UpdateHistoryPropertyStatusGenerator returns a generator of UpdateHistoryProperty_Status instances for property testing.
func UpdateHistoryPropertyStatusGenerator() gopter.Gen {
	if updateHistoryPropertyStatusGenerator != nil {
		return updateHistoryPropertyStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatus(generators)
	updateHistoryPropertyStatusGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_Status{}), generators)

	return updateHistoryPropertyStatusGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryPropertyStatus(gens map[string]gopter.Gen) {
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}
