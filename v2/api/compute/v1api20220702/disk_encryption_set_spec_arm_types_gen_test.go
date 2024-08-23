// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220702

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

func Test_DiskEncryptionSet_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSet_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSet_Spec_ARM, DiskEncryptionSet_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSet_Spec_ARM runs a test to see if a specific instance of DiskEncryptionSet_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSet_Spec_ARM(subject DiskEncryptionSet_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSet_Spec_ARM
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

// Generator of DiskEncryptionSet_Spec_ARM instances for property testing - lazily instantiated by
// DiskEncryptionSet_Spec_ARMGenerator()
var diskEncryptionSet_Spec_ARMGenerator gopter.Gen

// DiskEncryptionSet_Spec_ARMGenerator returns a generator of DiskEncryptionSet_Spec_ARM instances for property testing.
// We first initialize diskEncryptionSet_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskEncryptionSet_Spec_ARMGenerator() gopter.Gen {
	if diskEncryptionSet_Spec_ARMGenerator != nil {
		return diskEncryptionSet_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec_ARM(generators)
	diskEncryptionSet_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec_ARM(generators)
	diskEncryptionSet_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_Spec_ARM{}), generators)

	return diskEncryptionSet_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskEncryptionSet_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskEncryptionSet_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(EncryptionSetIdentity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(EncryptionSetProperties_ARMGenerator())
}

func Test_EncryptionSetIdentity_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity_ARM, EncryptionSetIdentity_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity_ARM runs a test to see if a specific instance of EncryptionSetIdentity_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity_ARM(subject EncryptionSetIdentity_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity_ARM
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

// Generator of EncryptionSetIdentity_ARM instances for property testing - lazily instantiated by
// EncryptionSetIdentity_ARMGenerator()
var encryptionSetIdentity_ARMGenerator gopter.Gen

// EncryptionSetIdentity_ARMGenerator returns a generator of EncryptionSetIdentity_ARM instances for property testing.
// We first initialize encryptionSetIdentity_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetIdentity_ARMGenerator() gopter.Gen {
	if encryptionSetIdentity_ARMGenerator != nil {
		return encryptionSetIdentity_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_ARM(generators)
	encryptionSetIdentity_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetIdentity_ARM(generators)
	encryptionSetIdentity_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_ARM{}), generators)

	return encryptionSetIdentity_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		EncryptionSetIdentity_Type_ARM_None,
		EncryptionSetIdentity_Type_ARM_SystemAssigned,
		EncryptionSetIdentity_Type_ARM_SystemAssignedUserAssigned,
		EncryptionSetIdentity_Type_ARM_UserAssigned))
}

// AddRelatedPropertyGeneratorsForEncryptionSetIdentity_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetIdentity_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetails_ARMGenerator())
}

func Test_EncryptionSetProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetProperties_ARM, EncryptionSetProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetProperties_ARM runs a test to see if a specific instance of EncryptionSetProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetProperties_ARM(subject EncryptionSetProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetProperties_ARM
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

// Generator of EncryptionSetProperties_ARM instances for property testing - lazily instantiated by
// EncryptionSetProperties_ARMGenerator()
var encryptionSetProperties_ARMGenerator gopter.Gen

// EncryptionSetProperties_ARMGenerator returns a generator of EncryptionSetProperties_ARM instances for property testing.
// We first initialize encryptionSetProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetProperties_ARMGenerator() gopter.Gen {
	if encryptionSetProperties_ARMGenerator != nil {
		return encryptionSetProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetProperties_ARM(generators)
	encryptionSetProperties_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetProperties_ARM(generators)
	encryptionSetProperties_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetProperties_ARM{}), generators)

	return encryptionSetProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetProperties_ARM(gens map[string]gopter.Gen) {
	gens["EncryptionType"] = gen.PtrOf(gen.OneConstOf(DiskEncryptionSetType_ARM_ConfidentialVmEncryptedWithCustomerKey, DiskEncryptionSetType_ARM_EncryptionAtRestWithCustomerKey, DiskEncryptionSetType_ARM_EncryptionAtRestWithPlatformAndCustomerKeys))
	gens["FederatedClientId"] = gen.PtrOf(gen.AlphaString())
	gens["RotationToLatestKeyVersionEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionSetProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetProperties_ARM(gens map[string]gopter.Gen) {
	gens["ActiveKey"] = gen.PtrOf(KeyForDiskEncryptionSet_ARMGenerator())
}

func Test_KeyForDiskEncryptionSet_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyForDiskEncryptionSet_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyForDiskEncryptionSet_ARM, KeyForDiskEncryptionSet_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyForDiskEncryptionSet_ARM runs a test to see if a specific instance of KeyForDiskEncryptionSet_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyForDiskEncryptionSet_ARM(subject KeyForDiskEncryptionSet_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyForDiskEncryptionSet_ARM
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

// Generator of KeyForDiskEncryptionSet_ARM instances for property testing - lazily instantiated by
// KeyForDiskEncryptionSet_ARMGenerator()
var keyForDiskEncryptionSet_ARMGenerator gopter.Gen

// KeyForDiskEncryptionSet_ARMGenerator returns a generator of KeyForDiskEncryptionSet_ARM instances for property testing.
// We first initialize keyForDiskEncryptionSet_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyForDiskEncryptionSet_ARMGenerator() gopter.Gen {
	if keyForDiskEncryptionSet_ARMGenerator != nil {
		return keyForDiskEncryptionSet_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_ARM(generators)
	keyForDiskEncryptionSet_ARMGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_ARM(generators)
	keyForDiskEncryptionSet_ARMGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_ARM{}), generators)

	return keyForDiskEncryptionSet_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_ARM(gens map[string]gopter.Gen) {
	gens["KeyUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_ARM(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVault_ARMGenerator())
}

func Test_SourceVault_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SourceVault_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSourceVault_ARM, SourceVault_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSourceVault_ARM runs a test to see if a specific instance of SourceVault_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSourceVault_ARM(subject SourceVault_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SourceVault_ARM
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

// Generator of SourceVault_ARM instances for property testing - lazily instantiated by SourceVault_ARMGenerator()
var sourceVault_ARMGenerator gopter.Gen

// SourceVault_ARMGenerator returns a generator of SourceVault_ARM instances for property testing.
func SourceVault_ARMGenerator() gopter.Gen {
	if sourceVault_ARMGenerator != nil {
		return sourceVault_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSourceVault_ARM(generators)
	sourceVault_ARMGenerator = gen.Struct(reflect.TypeOf(SourceVault_ARM{}), generators)

	return sourceVault_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSourceVault_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSourceVault_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
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
