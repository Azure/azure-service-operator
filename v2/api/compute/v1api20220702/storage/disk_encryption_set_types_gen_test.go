// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_DiskEncryptionSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSet, DiskEncryptionSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSet runs a test to see if a specific instance of DiskEncryptionSet round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSet(subject DiskEncryptionSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSet
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

// Generator of DiskEncryptionSet instances for property testing - lazily instantiated by DiskEncryptionSetGenerator()
var diskEncryptionSetGenerator gopter.Gen

// DiskEncryptionSetGenerator returns a generator of DiskEncryptionSet instances for property testing.
func DiskEncryptionSetGenerator() gopter.Gen {
	if diskEncryptionSetGenerator != nil {
		return diskEncryptionSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDiskEncryptionSet(generators)
	diskEncryptionSetGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet{}), generators)

	return diskEncryptionSetGenerator
}

// AddRelatedPropertyGeneratorsForDiskEncryptionSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskEncryptionSet(gens map[string]gopter.Gen) {
	gens["Spec"] = DiskEncryptionSet_SpecGenerator()
	gens["Status"] = DiskEncryptionSet_STATUSGenerator()
}

func Test_DiskEncryptionSet_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSet_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSet_Spec, DiskEncryptionSet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSet_Spec runs a test to see if a specific instance of DiskEncryptionSet_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSet_Spec(subject DiskEncryptionSet_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSet_Spec
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

// Generator of DiskEncryptionSet_Spec instances for property testing - lazily instantiated by
// DiskEncryptionSet_SpecGenerator()
var diskEncryptionSet_SpecGenerator gopter.Gen

// DiskEncryptionSet_SpecGenerator returns a generator of DiskEncryptionSet_Spec instances for property testing.
// We first initialize diskEncryptionSet_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskEncryptionSet_SpecGenerator() gopter.Gen {
	if diskEncryptionSet_SpecGenerator != nil {
		return diskEncryptionSet_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec(generators)
	diskEncryptionSet_SpecGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec(generators)
	AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec(generators)
	diskEncryptionSet_SpecGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_Spec{}), generators)

	return diskEncryptionSet_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EncryptionType"] = gen.PtrOf(gen.AlphaString())
	gens["FederatedClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RotationToLatestKeyVersionEnabled"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec(gens map[string]gopter.Gen) {
	gens["ActiveKey"] = gen.PtrOf(KeyForDiskEncryptionSetGenerator())
	gens["Identity"] = gen.PtrOf(EncryptionSetIdentityGenerator())
}

func Test_DiskEncryptionSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSet_STATUS, DiskEncryptionSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSet_STATUS runs a test to see if a specific instance of DiskEncryptionSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSet_STATUS(subject DiskEncryptionSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSet_STATUS
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

// Generator of DiskEncryptionSet_STATUS instances for property testing - lazily instantiated by
// DiskEncryptionSet_STATUSGenerator()
var diskEncryptionSet_STATUSGenerator gopter.Gen

// DiskEncryptionSet_STATUSGenerator returns a generator of DiskEncryptionSet_STATUS instances for property testing.
// We first initialize diskEncryptionSet_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskEncryptionSet_STATUSGenerator() gopter.Gen {
	if diskEncryptionSet_STATUSGenerator != nil {
		return diskEncryptionSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS(generators)
	diskEncryptionSet_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS(generators)
	diskEncryptionSet_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_STATUS{}), generators)

	return diskEncryptionSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS(gens map[string]gopter.Gen) {
	gens["EncryptionType"] = gen.PtrOf(gen.AlphaString())
	gens["FederatedClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LastKeyRotationTimestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RotationToLatestKeyVersionEnabled"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS(gens map[string]gopter.Gen) {
	gens["ActiveKey"] = gen.PtrOf(KeyForDiskEncryptionSet_STATUSGenerator())
	gens["AutoKeyRotationError"] = gen.PtrOf(ApiError_STATUSGenerator())
	gens["Identity"] = gen.PtrOf(EncryptionSetIdentity_STATUSGenerator())
	gens["PreviousKeys"] = gen.SliceOf(KeyForDiskEncryptionSet_STATUSGenerator())
}

func Test_ApiError_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiError_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiError_STATUS, ApiError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiError_STATUS runs a test to see if a specific instance of ApiError_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiError_STATUS(subject ApiError_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiError_STATUS
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

// Generator of ApiError_STATUS instances for property testing - lazily instantiated by ApiError_STATUSGenerator()
var apiError_STATUSGenerator gopter.Gen

// ApiError_STATUSGenerator returns a generator of ApiError_STATUS instances for property testing.
// We first initialize apiError_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApiError_STATUSGenerator() gopter.Gen {
	if apiError_STATUSGenerator != nil {
		return apiError_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiError_STATUS(generators)
	apiError_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiError_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiError_STATUS(generators)
	AddRelatedPropertyGeneratorsForApiError_STATUS(generators)
	apiError_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiError_STATUS{}), generators)

	return apiError_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiError_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiError_STATUS(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApiError_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApiError_STATUS(gens map[string]gopter.Gen) {
	gens["Details"] = gen.SliceOf(ApiErrorBase_STATUSGenerator())
	gens["Innererror"] = gen.PtrOf(InnerError_STATUSGenerator())
}

func Test_EncryptionSetIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity, EncryptionSetIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity runs a test to see if a specific instance of EncryptionSetIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity(subject EncryptionSetIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity
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

// Generator of EncryptionSetIdentity instances for property testing - lazily instantiated by
// EncryptionSetIdentityGenerator()
var encryptionSetIdentityGenerator gopter.Gen

// EncryptionSetIdentityGenerator returns a generator of EncryptionSetIdentity instances for property testing.
// We first initialize encryptionSetIdentityGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetIdentityGenerator() gopter.Gen {
	if encryptionSetIdentityGenerator != nil {
		return encryptionSetIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity(generators)
	encryptionSetIdentityGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetIdentity(generators)
	encryptionSetIdentityGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity{}), generators)

	return encryptionSetIdentityGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionSetIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetIdentity(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.SliceOf(UserAssignedIdentityDetailsGenerator())
}

func Test_EncryptionSetIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity_STATUS, EncryptionSetIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity_STATUS runs a test to see if a specific instance of EncryptionSetIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity_STATUS(subject EncryptionSetIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity_STATUS
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

// Generator of EncryptionSetIdentity_STATUS instances for property testing - lazily instantiated by
// EncryptionSetIdentity_STATUSGenerator()
var encryptionSetIdentity_STATUSGenerator gopter.Gen

// EncryptionSetIdentity_STATUSGenerator returns a generator of EncryptionSetIdentity_STATUS instances for property testing.
// We first initialize encryptionSetIdentity_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetIdentity_STATUSGenerator() gopter.Gen {
	if encryptionSetIdentity_STATUSGenerator != nil {
		return encryptionSetIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS(generators)
	encryptionSetIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS(generators)
	encryptionSetIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_STATUS{}), generators)

	return encryptionSetIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		EncryptionSetIdentity_UserAssignedIdentities_STATUSGenerator())
}

func Test_KeyForDiskEncryptionSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyForDiskEncryptionSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyForDiskEncryptionSet, KeyForDiskEncryptionSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyForDiskEncryptionSet runs a test to see if a specific instance of KeyForDiskEncryptionSet round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyForDiskEncryptionSet(subject KeyForDiskEncryptionSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyForDiskEncryptionSet
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

// Generator of KeyForDiskEncryptionSet instances for property testing - lazily instantiated by
// KeyForDiskEncryptionSetGenerator()
var keyForDiskEncryptionSetGenerator gopter.Gen

// KeyForDiskEncryptionSetGenerator returns a generator of KeyForDiskEncryptionSet instances for property testing.
// We first initialize keyForDiskEncryptionSetGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyForDiskEncryptionSetGenerator() gopter.Gen {
	if keyForDiskEncryptionSetGenerator != nil {
		return keyForDiskEncryptionSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet(generators)
	keyForDiskEncryptionSetGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet(generators)
	AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet(generators)
	keyForDiskEncryptionSetGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet{}), generators)

	return keyForDiskEncryptionSetGenerator
}

// AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet(gens map[string]gopter.Gen) {
	gens["KeyUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVaultGenerator())
}

func Test_KeyForDiskEncryptionSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyForDiskEncryptionSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS, KeyForDiskEncryptionSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS runs a test to see if a specific instance of KeyForDiskEncryptionSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS(subject KeyForDiskEncryptionSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyForDiskEncryptionSet_STATUS
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

// Generator of KeyForDiskEncryptionSet_STATUS instances for property testing - lazily instantiated by
// KeyForDiskEncryptionSet_STATUSGenerator()
var keyForDiskEncryptionSet_STATUSGenerator gopter.Gen

// KeyForDiskEncryptionSet_STATUSGenerator returns a generator of KeyForDiskEncryptionSet_STATUS instances for property testing.
// We first initialize keyForDiskEncryptionSet_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyForDiskEncryptionSet_STATUSGenerator() gopter.Gen {
	if keyForDiskEncryptionSet_STATUSGenerator != nil {
		return keyForDiskEncryptionSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS(generators)
	keyForDiskEncryptionSet_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS(generators)
	AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS(generators)
	keyForDiskEncryptionSet_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_STATUS{}), generators)

	return keyForDiskEncryptionSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS(gens map[string]gopter.Gen) {
	gens["KeyUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVault_STATUSGenerator())
}

func Test_ApiErrorBase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiErrorBase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiErrorBase_STATUS, ApiErrorBase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiErrorBase_STATUS runs a test to see if a specific instance of ApiErrorBase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiErrorBase_STATUS(subject ApiErrorBase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiErrorBase_STATUS
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

// Generator of ApiErrorBase_STATUS instances for property testing - lazily instantiated by
// ApiErrorBase_STATUSGenerator()
var apiErrorBase_STATUSGenerator gopter.Gen

// ApiErrorBase_STATUSGenerator returns a generator of ApiErrorBase_STATUS instances for property testing.
func ApiErrorBase_STATUSGenerator() gopter.Gen {
	if apiErrorBase_STATUSGenerator != nil {
		return apiErrorBase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiErrorBase_STATUS(generators)
	apiErrorBase_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiErrorBase_STATUS{}), generators)

	return apiErrorBase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiErrorBase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiErrorBase_STATUS(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

func Test_EncryptionSetIdentity_UserAssignedIdentities_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity_UserAssignedIdentities_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS, EncryptionSetIdentity_UserAssignedIdentities_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS runs a test to see if a specific instance of EncryptionSetIdentity_UserAssignedIdentities_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS(subject EncryptionSetIdentity_UserAssignedIdentities_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity_UserAssignedIdentities_STATUS
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

// Generator of EncryptionSetIdentity_UserAssignedIdentities_STATUS instances for property testing - lazily instantiated
// by EncryptionSetIdentity_UserAssignedIdentities_STATUSGenerator()
var encryptionSetIdentity_UserAssignedIdentities_STATUSGenerator gopter.Gen

// EncryptionSetIdentity_UserAssignedIdentities_STATUSGenerator returns a generator of EncryptionSetIdentity_UserAssignedIdentities_STATUS instances for property testing.
func EncryptionSetIdentity_UserAssignedIdentities_STATUSGenerator() gopter.Gen {
	if encryptionSetIdentity_UserAssignedIdentities_STATUSGenerator != nil {
		return encryptionSetIdentity_UserAssignedIdentities_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS(generators)
	encryptionSetIdentity_UserAssignedIdentities_STATUSGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_UserAssignedIdentities_STATUS{}), generators)

	return encryptionSetIdentity_UserAssignedIdentities_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_InnerError_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InnerError_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInnerError_STATUS, InnerError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInnerError_STATUS runs a test to see if a specific instance of InnerError_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForInnerError_STATUS(subject InnerError_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InnerError_STATUS
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

// Generator of InnerError_STATUS instances for property testing - lazily instantiated by InnerError_STATUSGenerator()
var innerError_STATUSGenerator gopter.Gen

// InnerError_STATUSGenerator returns a generator of InnerError_STATUS instances for property testing.
func InnerError_STATUSGenerator() gopter.Gen {
	if innerError_STATUSGenerator != nil {
		return innerError_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInnerError_STATUS(generators)
	innerError_STATUSGenerator = gen.Struct(reflect.TypeOf(InnerError_STATUS{}), generators)

	return innerError_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForInnerError_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInnerError_STATUS(gens map[string]gopter.Gen) {
	gens["Errordetail"] = gen.PtrOf(gen.AlphaString())
	gens["Exceptiontype"] = gen.PtrOf(gen.AlphaString())
}

func Test_SourceVault_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SourceVault via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSourceVault, SourceVaultGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSourceVault runs a test to see if a specific instance of SourceVault round trips to JSON and back losslessly
func RunJSONSerializationTestForSourceVault(subject SourceVault) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SourceVault
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

// Generator of SourceVault instances for property testing - lazily instantiated by SourceVaultGenerator()
var sourceVaultGenerator gopter.Gen

// SourceVaultGenerator returns a generator of SourceVault instances for property testing.
func SourceVaultGenerator() gopter.Gen {
	if sourceVaultGenerator != nil {
		return sourceVaultGenerator
	}

	generators := make(map[string]gopter.Gen)
	sourceVaultGenerator = gen.Struct(reflect.TypeOf(SourceVault{}), generators)

	return sourceVaultGenerator
}

func Test_SourceVault_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SourceVault_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSourceVault_STATUS, SourceVault_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSourceVault_STATUS runs a test to see if a specific instance of SourceVault_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSourceVault_STATUS(subject SourceVault_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SourceVault_STATUS
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

// Generator of SourceVault_STATUS instances for property testing - lazily instantiated by SourceVault_STATUSGenerator()
var sourceVault_STATUSGenerator gopter.Gen

// SourceVault_STATUSGenerator returns a generator of SourceVault_STATUS instances for property testing.
func SourceVault_STATUSGenerator() gopter.Gen {
	if sourceVault_STATUSGenerator != nil {
		return sourceVault_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSourceVault_STATUS(generators)
	sourceVault_STATUSGenerator = gen.Struct(reflect.TypeOf(SourceVault_STATUS{}), generators)

	return sourceVault_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSourceVault_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSourceVault_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails, UserAssignedIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails runs a test to see if a specific instance of UserAssignedIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails(subject UserAssignedIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails
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

// Generator of UserAssignedIdentityDetails instances for property testing - lazily instantiated by
// UserAssignedIdentityDetailsGenerator()
var userAssignedIdentityDetailsGenerator gopter.Gen

// UserAssignedIdentityDetailsGenerator returns a generator of UserAssignedIdentityDetails instances for property testing.
func UserAssignedIdentityDetailsGenerator() gopter.Gen {
	if userAssignedIdentityDetailsGenerator != nil {
		return userAssignedIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails{}), generators)

	return userAssignedIdentityDetailsGenerator
}
