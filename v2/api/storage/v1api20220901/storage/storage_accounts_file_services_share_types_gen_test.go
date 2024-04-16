// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
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

func Test_StorageAccountsFileServicesShare_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileServicesShare to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsFileServicesShare, StorageAccountsFileServicesShareGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsFileServicesShare tests if a specific instance of StorageAccountsFileServicesShare round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsFileServicesShare(subject StorageAccountsFileServicesShare) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230101s.StorageAccountsFileServicesShare
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsFileServicesShare
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileServicesShare_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileServicesShare to StorageAccountsFileServicesShare via AssignProperties_To_StorageAccountsFileServicesShare & AssignProperties_From_StorageAccountsFileServicesShare returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsFileServicesShare, StorageAccountsFileServicesShareGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsFileServicesShare tests if a specific instance of StorageAccountsFileServicesShare can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsFileServicesShare(subject StorageAccountsFileServicesShare) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.StorageAccountsFileServicesShare
	err := copied.AssignProperties_To_StorageAccountsFileServicesShare(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsFileServicesShare
	err = actual.AssignProperties_From_StorageAccountsFileServicesShare(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileServicesShare_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileServicesShare via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileServicesShare, StorageAccountsFileServicesShareGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileServicesShare runs a test to see if a specific instance of StorageAccountsFileServicesShare round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileServicesShare(subject StorageAccountsFileServicesShare) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileServicesShare
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

// Generator of StorageAccountsFileServicesShare instances for property testing - lazily instantiated by
// StorageAccountsFileServicesShareGenerator()
var storageAccountsFileServicesShareGenerator gopter.Gen

// StorageAccountsFileServicesShareGenerator returns a generator of StorageAccountsFileServicesShare instances for property testing.
func StorageAccountsFileServicesShareGenerator() gopter.Gen {
	if storageAccountsFileServicesShareGenerator != nil {
		return storageAccountsFileServicesShareGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare(generators)
	storageAccountsFileServicesShareGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileServicesShare{}), generators)

	return storageAccountsFileServicesShareGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccounts_FileServices_Share_SpecGenerator()
	gens["Status"] = StorageAccounts_FileServices_Share_STATUSGenerator()
}

func Test_StorageAccounts_FileServices_Share_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_FileServices_Share_Spec to StorageAccounts_FileServices_Share_Spec via AssignProperties_To_StorageAccounts_FileServices_Share_Spec & AssignProperties_From_StorageAccounts_FileServices_Share_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_Spec, StorageAccounts_FileServices_Share_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_Spec tests if a specific instance of StorageAccounts_FileServices_Share_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_Spec(subject StorageAccounts_FileServices_Share_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.StorageAccounts_FileServices_Share_Spec
	err := copied.AssignProperties_To_StorageAccounts_FileServices_Share_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_FileServices_Share_Spec
	err = actual.AssignProperties_From_StorageAccounts_FileServices_Share_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccounts_FileServices_Share_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileServices_Share_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileServices_Share_Spec, StorageAccounts_FileServices_Share_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileServices_Share_Spec runs a test to see if a specific instance of StorageAccounts_FileServices_Share_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileServices_Share_Spec(subject StorageAccounts_FileServices_Share_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileServices_Share_Spec
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

// Generator of StorageAccounts_FileServices_Share_Spec instances for property testing - lazily instantiated by
// StorageAccounts_FileServices_Share_SpecGenerator()
var storageAccounts_FileServices_Share_SpecGenerator gopter.Gen

// StorageAccounts_FileServices_Share_SpecGenerator returns a generator of StorageAccounts_FileServices_Share_Spec instances for property testing.
// We first initialize storageAccounts_FileServices_Share_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_FileServices_Share_SpecGenerator() gopter.Gen {
	if storageAccounts_FileServices_Share_SpecGenerator != nil {
		return storageAccounts_FileServices_Share_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec(generators)
	storageAccounts_FileServices_Share_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileServices_Share_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec(generators)
	storageAccounts_FileServices_Share_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileServices_Share_Spec{}), generators)

	return storageAccounts_FileServices_Share_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec(gens map[string]gopter.Gen) {
	gens["AccessTier"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["EnabledProtocols"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RootSquash"] = gen.PtrOf(gen.AlphaString())
	gens["ShareQuota"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_Spec(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(SignedIdentifierGenerator())
}

func Test_StorageAccounts_FileServices_Share_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_FileServices_Share_STATUS to StorageAccounts_FileServices_Share_STATUS via AssignProperties_To_StorageAccounts_FileServices_Share_STATUS & AssignProperties_From_StorageAccounts_FileServices_Share_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_STATUS, StorageAccounts_FileServices_Share_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_STATUS tests if a specific instance of StorageAccounts_FileServices_Share_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_FileServices_Share_STATUS(subject StorageAccounts_FileServices_Share_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.StorageAccounts_FileServices_Share_STATUS
	err := copied.AssignProperties_To_StorageAccounts_FileServices_Share_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_FileServices_Share_STATUS
	err = actual.AssignProperties_From_StorageAccounts_FileServices_Share_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccounts_FileServices_Share_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_FileServices_Share_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_FileServices_Share_STATUS, StorageAccounts_FileServices_Share_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_FileServices_Share_STATUS runs a test to see if a specific instance of StorageAccounts_FileServices_Share_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_FileServices_Share_STATUS(subject StorageAccounts_FileServices_Share_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_FileServices_Share_STATUS
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

// Generator of StorageAccounts_FileServices_Share_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_FileServices_Share_STATUSGenerator()
var storageAccounts_FileServices_Share_STATUSGenerator gopter.Gen

// StorageAccounts_FileServices_Share_STATUSGenerator returns a generator of StorageAccounts_FileServices_Share_STATUS instances for property testing.
// We first initialize storageAccounts_FileServices_Share_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_FileServices_Share_STATUSGenerator() gopter.Gen {
	if storageAccounts_FileServices_Share_STATUSGenerator != nil {
		return storageAccounts_FileServices_Share_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS(generators)
	storageAccounts_FileServices_Share_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileServices_Share_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS(generators)
	storageAccounts_FileServices_Share_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_FileServices_Share_STATUS{}), generators)

	return storageAccounts_FileServices_Share_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS(gens map[string]gopter.Gen) {
	gens["AccessTier"] = gen.PtrOf(gen.AlphaString())
	gens["AccessTierChangeTime"] = gen.PtrOf(gen.AlphaString())
	gens["AccessTierStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["EnabledProtocols"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseState"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["RootSquash"] = gen.PtrOf(gen.AlphaString())
	gens["ShareQuota"] = gen.PtrOf(gen.Int())
	gens["ShareUsageBytes"] = gen.PtrOf(gen.Int())
	gens["SnapshotTime"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_FileServices_Share_STATUS(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(SignedIdentifier_STATUSGenerator())
}

func Test_SignedIdentifier_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SignedIdentifier to SignedIdentifier via AssignProperties_To_SignedIdentifier & AssignProperties_From_SignedIdentifier returns original",
		prop.ForAll(RunPropertyAssignmentTestForSignedIdentifier, SignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSignedIdentifier tests if a specific instance of SignedIdentifier can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSignedIdentifier(subject SignedIdentifier) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.SignedIdentifier
	err := copied.AssignProperties_To_SignedIdentifier(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SignedIdentifier
	err = actual.AssignProperties_From_SignedIdentifier(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SignedIdentifier_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SignedIdentifier via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSignedIdentifier, SignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSignedIdentifier runs a test to see if a specific instance of SignedIdentifier round trips to JSON and back losslessly
func RunJSONSerializationTestForSignedIdentifier(subject SignedIdentifier) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SignedIdentifier
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

// Generator of SignedIdentifier instances for property testing - lazily instantiated by SignedIdentifierGenerator()
var signedIdentifierGenerator gopter.Gen

// SignedIdentifierGenerator returns a generator of SignedIdentifier instances for property testing.
func SignedIdentifierGenerator() gopter.Gen {
	if signedIdentifierGenerator != nil {
		return signedIdentifierGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSignedIdentifier(generators)
	signedIdentifierGenerator = gen.Struct(reflect.TypeOf(SignedIdentifier{}), generators)

	return signedIdentifierGenerator
}

// AddRelatedPropertyGeneratorsForSignedIdentifier is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSignedIdentifier(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(AccessPolicyGenerator())
}

func Test_SignedIdentifier_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SignedIdentifier_STATUS to SignedIdentifier_STATUS via AssignProperties_To_SignedIdentifier_STATUS & AssignProperties_From_SignedIdentifier_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSignedIdentifier_STATUS, SignedIdentifier_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSignedIdentifier_STATUS tests if a specific instance of SignedIdentifier_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSignedIdentifier_STATUS(subject SignedIdentifier_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.SignedIdentifier_STATUS
	err := copied.AssignProperties_To_SignedIdentifier_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SignedIdentifier_STATUS
	err = actual.AssignProperties_From_SignedIdentifier_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SignedIdentifier_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SignedIdentifier_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSignedIdentifier_STATUS, SignedIdentifier_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSignedIdentifier_STATUS runs a test to see if a specific instance of SignedIdentifier_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSignedIdentifier_STATUS(subject SignedIdentifier_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SignedIdentifier_STATUS
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

// Generator of SignedIdentifier_STATUS instances for property testing - lazily instantiated by
// SignedIdentifier_STATUSGenerator()
var signedIdentifier_STATUSGenerator gopter.Gen

// SignedIdentifier_STATUSGenerator returns a generator of SignedIdentifier_STATUS instances for property testing.
// We first initialize signedIdentifier_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SignedIdentifier_STATUSGenerator() gopter.Gen {
	if signedIdentifier_STATUSGenerator != nil {
		return signedIdentifier_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSignedIdentifier_STATUS(generators)
	signedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(SignedIdentifier_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSignedIdentifier_STATUS(generators)
	AddRelatedPropertyGeneratorsForSignedIdentifier_STATUS(generators)
	signedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(SignedIdentifier_STATUS{}), generators)

	return signedIdentifier_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(AccessPolicy_STATUSGenerator())
}

func Test_AccessPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AccessPolicy to AccessPolicy via AssignProperties_To_AccessPolicy & AssignProperties_From_AccessPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForAccessPolicy, AccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAccessPolicy tests if a specific instance of AccessPolicy can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAccessPolicy(subject AccessPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.AccessPolicy
	err := copied.AssignProperties_To_AccessPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AccessPolicy
	err = actual.AssignProperties_From_AccessPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAccessPolicy, AccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAccessPolicy runs a test to see if a specific instance of AccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForAccessPolicy(subject AccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AccessPolicy
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

// Generator of AccessPolicy instances for property testing - lazily instantiated by AccessPolicyGenerator()
var accessPolicyGenerator gopter.Gen

// AccessPolicyGenerator returns a generator of AccessPolicy instances for property testing.
func AccessPolicyGenerator() gopter.Gen {
	if accessPolicyGenerator != nil {
		return accessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicy(generators)
	accessPolicyGenerator = gen.Struct(reflect.TypeOf(AccessPolicy{}), generators)

	return accessPolicyGenerator
}

// AddIndependentPropertyGeneratorsForAccessPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAccessPolicy(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_AccessPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AccessPolicy_STATUS to AccessPolicy_STATUS via AssignProperties_To_AccessPolicy_STATUS & AssignProperties_From_AccessPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAccessPolicy_STATUS, AccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAccessPolicy_STATUS tests if a specific instance of AccessPolicy_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAccessPolicy_STATUS(subject AccessPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230101s.AccessPolicy_STATUS
	err := copied.AssignProperties_To_AccessPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AccessPolicy_STATUS
	err = actual.AssignProperties_From_AccessPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AccessPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AccessPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAccessPolicy_STATUS, AccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAccessPolicy_STATUS runs a test to see if a specific instance of AccessPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAccessPolicy_STATUS(subject AccessPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AccessPolicy_STATUS
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

// Generator of AccessPolicy_STATUS instances for property testing - lazily instantiated by
// AccessPolicy_STATUSGenerator()
var accessPolicy_STATUSGenerator gopter.Gen

// AccessPolicy_STATUSGenerator returns a generator of AccessPolicy_STATUS instances for property testing.
func AccessPolicy_STATUSGenerator() gopter.Gen {
	if accessPolicy_STATUSGenerator != nil {
		return accessPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicy_STATUS(generators)
	accessPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(AccessPolicy_STATUS{}), generators)

	return accessPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAccessPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAccessPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}
