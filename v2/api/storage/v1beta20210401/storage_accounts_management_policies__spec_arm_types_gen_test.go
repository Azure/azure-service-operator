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

func Test_StorageAccountsManagementPolicies_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsManagementPolicies_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsManagementPoliciesSpecARM, StorageAccountsManagementPoliciesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsManagementPoliciesSpecARM runs a test to see if a specific instance of StorageAccountsManagementPolicies_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsManagementPoliciesSpecARM(subject StorageAccountsManagementPolicies_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsManagementPolicies_SpecARM
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

// Generator of StorageAccountsManagementPolicies_SpecARM instances for property testing - lazily instantiated by
// StorageAccountsManagementPoliciesSpecARMGenerator()
var storageAccountsManagementPoliciesSpecARMGenerator gopter.Gen

// StorageAccountsManagementPoliciesSpecARMGenerator returns a generator of StorageAccountsManagementPolicies_SpecARM instances for property testing.
// We first initialize storageAccountsManagementPoliciesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsManagementPoliciesSpecARMGenerator() gopter.Gen {
	if storageAccountsManagementPoliciesSpecARMGenerator != nil {
		return storageAccountsManagementPoliciesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM(generators)
	storageAccountsManagementPoliciesSpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsManagementPolicies_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM(generators)
	storageAccountsManagementPoliciesSpecARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsManagementPolicies_SpecARM{}), generators)

	return storageAccountsManagementPoliciesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsManagementPoliciesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagementPolicyPropertiesARMGenerator())
}

func Test_ManagementPolicyPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyPropertiesARM, ManagementPolicyPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyPropertiesARM runs a test to see if a specific instance of ManagementPolicyPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyPropertiesARM(subject ManagementPolicyPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyPropertiesARM
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

// Generator of ManagementPolicyPropertiesARM instances for property testing - lazily instantiated by
// ManagementPolicyPropertiesARMGenerator()
var managementPolicyPropertiesARMGenerator gopter.Gen

// ManagementPolicyPropertiesARMGenerator returns a generator of ManagementPolicyPropertiesARM instances for property testing.
func ManagementPolicyPropertiesARMGenerator() gopter.Gen {
	if managementPolicyPropertiesARMGenerator != nil {
		return managementPolicyPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyPropertiesARM(generators)
	managementPolicyPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyPropertiesARM{}), generators)

	return managementPolicyPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyPropertiesARM(gens map[string]gopter.Gen) {
	gens["Policy"] = gen.PtrOf(ManagementPolicySchemaARMGenerator())
}

func Test_ManagementPolicySchemaARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySchemaARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySchemaARM, ManagementPolicySchemaARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySchemaARM runs a test to see if a specific instance of ManagementPolicySchemaARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySchemaARM(subject ManagementPolicySchemaARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySchemaARM
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

// Generator of ManagementPolicySchemaARM instances for property testing - lazily instantiated by
// ManagementPolicySchemaARMGenerator()
var managementPolicySchemaARMGenerator gopter.Gen

// ManagementPolicySchemaARMGenerator returns a generator of ManagementPolicySchemaARM instances for property testing.
func ManagementPolicySchemaARMGenerator() gopter.Gen {
	if managementPolicySchemaARMGenerator != nil {
		return managementPolicySchemaARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySchemaARM(generators)
	managementPolicySchemaARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySchemaARM{}), generators)

	return managementPolicySchemaARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySchemaARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySchemaARM(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(ManagementPolicyRuleARMGenerator())
}

func Test_ManagementPolicyRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyRuleARM, ManagementPolicyRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyRuleARM runs a test to see if a specific instance of ManagementPolicyRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyRuleARM(subject ManagementPolicyRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyRuleARM
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

// Generator of ManagementPolicyRuleARM instances for property testing - lazily instantiated by
// ManagementPolicyRuleARMGenerator()
var managementPolicyRuleARMGenerator gopter.Gen

// ManagementPolicyRuleARMGenerator returns a generator of ManagementPolicyRuleARM instances for property testing.
// We first initialize managementPolicyRuleARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyRuleARMGenerator() gopter.Gen {
	if managementPolicyRuleARMGenerator != nil {
		return managementPolicyRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRuleARM(generators)
	managementPolicyRuleARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRuleARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRuleARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyRuleARM(generators)
	managementPolicyRuleARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRuleARM{}), generators)

	return managementPolicyRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyRuleARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagementPolicyRuleTypeLifecycle))
}

// AddRelatedPropertyGeneratorsForManagementPolicyRuleARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyRuleARM(gens map[string]gopter.Gen) {
	gens["Definition"] = gen.PtrOf(ManagementPolicyDefinitionARMGenerator())
}

func Test_ManagementPolicyDefinitionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyDefinitionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyDefinitionARM, ManagementPolicyDefinitionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyDefinitionARM runs a test to see if a specific instance of ManagementPolicyDefinitionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyDefinitionARM(subject ManagementPolicyDefinitionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyDefinitionARM
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

// Generator of ManagementPolicyDefinitionARM instances for property testing - lazily instantiated by
// ManagementPolicyDefinitionARMGenerator()
var managementPolicyDefinitionARMGenerator gopter.Gen

// ManagementPolicyDefinitionARMGenerator returns a generator of ManagementPolicyDefinitionARM instances for property testing.
func ManagementPolicyDefinitionARMGenerator() gopter.Gen {
	if managementPolicyDefinitionARMGenerator != nil {
		return managementPolicyDefinitionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyDefinitionARM(generators)
	managementPolicyDefinitionARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyDefinitionARM{}), generators)

	return managementPolicyDefinitionARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyDefinitionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyDefinitionARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(ManagementPolicyActionARMGenerator())
	gens["Filters"] = gen.PtrOf(ManagementPolicyFilterARMGenerator())
}

func Test_ManagementPolicyActionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyActionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyActionARM, ManagementPolicyActionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyActionARM runs a test to see if a specific instance of ManagementPolicyActionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyActionARM(subject ManagementPolicyActionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyActionARM
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

// Generator of ManagementPolicyActionARM instances for property testing - lazily instantiated by
// ManagementPolicyActionARMGenerator()
var managementPolicyActionARMGenerator gopter.Gen

// ManagementPolicyActionARMGenerator returns a generator of ManagementPolicyActionARM instances for property testing.
func ManagementPolicyActionARMGenerator() gopter.Gen {
	if managementPolicyActionARMGenerator != nil {
		return managementPolicyActionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyActionARM(generators)
	managementPolicyActionARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyActionARM{}), generators)

	return managementPolicyActionARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyActionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyActionARM(gens map[string]gopter.Gen) {
	gens["BaseBlob"] = gen.PtrOf(ManagementPolicyBaseBlobARMGenerator())
	gens["Snapshot"] = gen.PtrOf(ManagementPolicySnapShotARMGenerator())
	gens["Version"] = gen.PtrOf(ManagementPolicyVersionARMGenerator())
}

func Test_ManagementPolicyFilterARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyFilterARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyFilterARM, ManagementPolicyFilterARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyFilterARM runs a test to see if a specific instance of ManagementPolicyFilterARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyFilterARM(subject ManagementPolicyFilterARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyFilterARM
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

// Generator of ManagementPolicyFilterARM instances for property testing - lazily instantiated by
// ManagementPolicyFilterARMGenerator()
var managementPolicyFilterARMGenerator gopter.Gen

// ManagementPolicyFilterARMGenerator returns a generator of ManagementPolicyFilterARM instances for property testing.
// We first initialize managementPolicyFilterARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyFilterARMGenerator() gopter.Gen {
	if managementPolicyFilterARMGenerator != nil {
		return managementPolicyFilterARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilterARM(generators)
	managementPolicyFilterARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilterARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilterARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyFilterARM(generators)
	managementPolicyFilterARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilterARM{}), generators)

	return managementPolicyFilterARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyFilterARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyFilterARM(gens map[string]gopter.Gen) {
	gens["BlobTypes"] = gen.SliceOf(gen.AlphaString())
	gens["PrefixMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyFilterARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyFilterARM(gens map[string]gopter.Gen) {
	gens["BlobIndexMatch"] = gen.SliceOf(TagFilterARMGenerator())
}

func Test_ManagementPolicyBaseBlobARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyBaseBlobARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyBaseBlobARM, ManagementPolicyBaseBlobARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyBaseBlobARM runs a test to see if a specific instance of ManagementPolicyBaseBlobARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyBaseBlobARM(subject ManagementPolicyBaseBlobARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyBaseBlobARM
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

// Generator of ManagementPolicyBaseBlobARM instances for property testing - lazily instantiated by
// ManagementPolicyBaseBlobARMGenerator()
var managementPolicyBaseBlobARMGenerator gopter.Gen

// ManagementPolicyBaseBlobARMGenerator returns a generator of ManagementPolicyBaseBlobARM instances for property testing.
// We first initialize managementPolicyBaseBlobARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyBaseBlobARMGenerator() gopter.Gen {
	if managementPolicyBaseBlobARMGenerator != nil {
		return managementPolicyBaseBlobARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobARM(generators)
	managementPolicyBaseBlobARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlobARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobARM(generators)
	managementPolicyBaseBlobARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlobARM{}), generators)

	return managementPolicyBaseBlobARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyBaseBlobARM(gens map[string]gopter.Gen) {
	gens["EnableAutoTierToHotFromCool"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyBaseBlobARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterModificationARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterModificationARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterModificationARMGenerator())
}

func Test_ManagementPolicySnapShotARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySnapShotARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySnapShotARM, ManagementPolicySnapShotARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySnapShotARM runs a test to see if a specific instance of ManagementPolicySnapShotARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySnapShotARM(subject ManagementPolicySnapShotARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySnapShotARM
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

// Generator of ManagementPolicySnapShotARM instances for property testing - lazily instantiated by
// ManagementPolicySnapShotARMGenerator()
var managementPolicySnapShotARMGenerator gopter.Gen

// ManagementPolicySnapShotARMGenerator returns a generator of ManagementPolicySnapShotARM instances for property testing.
func ManagementPolicySnapShotARMGenerator() gopter.Gen {
	if managementPolicySnapShotARMGenerator != nil {
		return managementPolicySnapShotARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySnapShotARM(generators)
	managementPolicySnapShotARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySnapShotARM{}), generators)

	return managementPolicySnapShotARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySnapShotARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySnapShotARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreationARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreationARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreationARMGenerator())
}

func Test_ManagementPolicyVersionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyVersionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyVersionARM, ManagementPolicyVersionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyVersionARM runs a test to see if a specific instance of ManagementPolicyVersionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyVersionARM(subject ManagementPolicyVersionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyVersionARM
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

// Generator of ManagementPolicyVersionARM instances for property testing - lazily instantiated by
// ManagementPolicyVersionARMGenerator()
var managementPolicyVersionARMGenerator gopter.Gen

// ManagementPolicyVersionARMGenerator returns a generator of ManagementPolicyVersionARM instances for property testing.
func ManagementPolicyVersionARMGenerator() gopter.Gen {
	if managementPolicyVersionARMGenerator != nil {
		return managementPolicyVersionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyVersionARM(generators)
	managementPolicyVersionARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyVersionARM{}), generators)

	return managementPolicyVersionARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyVersionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyVersionARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreationARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreationARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreationARMGenerator())
}

func Test_TagFilterARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagFilterARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagFilterARM, TagFilterARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagFilterARM runs a test to see if a specific instance of TagFilterARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagFilterARM(subject TagFilterARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagFilterARM
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

// Generator of TagFilterARM instances for property testing - lazily instantiated by TagFilterARMGenerator()
var tagFilterARMGenerator gopter.Gen

// TagFilterARMGenerator returns a generator of TagFilterARM instances for property testing.
func TagFilterARMGenerator() gopter.Gen {
	if tagFilterARMGenerator != nil {
		return tagFilterARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagFilterARM(generators)
	tagFilterARMGenerator = gen.Struct(reflect.TypeOf(TagFilterARM{}), generators)

	return tagFilterARMGenerator
}

// AddIndependentPropertyGeneratorsForTagFilterARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagFilterARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Op"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_DateAfterCreationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterCreationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterCreationARM, DateAfterCreationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterCreationARM runs a test to see if a specific instance of DateAfterCreationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterCreationARM(subject DateAfterCreationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterCreationARM
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

// Generator of DateAfterCreationARM instances for property testing - lazily instantiated by
// DateAfterCreationARMGenerator()
var dateAfterCreationARMGenerator gopter.Gen

// DateAfterCreationARMGenerator returns a generator of DateAfterCreationARM instances for property testing.
func DateAfterCreationARMGenerator() gopter.Gen {
	if dateAfterCreationARMGenerator != nil {
		return dateAfterCreationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterCreationARM(generators)
	dateAfterCreationARMGenerator = gen.Struct(reflect.TypeOf(DateAfterCreationARM{}), generators)

	return dateAfterCreationARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterCreationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterCreationARM(gens map[string]gopter.Gen) {
	gens["DaysAfterCreationGreaterThan"] = gen.PtrOf(gen.Int())
}

func Test_DateAfterModificationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterModificationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterModificationARM, DateAfterModificationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterModificationARM runs a test to see if a specific instance of DateAfterModificationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterModificationARM(subject DateAfterModificationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterModificationARM
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

// Generator of DateAfterModificationARM instances for property testing - lazily instantiated by
// DateAfterModificationARMGenerator()
var dateAfterModificationARMGenerator gopter.Gen

// DateAfterModificationARMGenerator returns a generator of DateAfterModificationARM instances for property testing.
func DateAfterModificationARMGenerator() gopter.Gen {
	if dateAfterModificationARMGenerator != nil {
		return dateAfterModificationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterModificationARM(generators)
	dateAfterModificationARMGenerator = gen.Struct(reflect.TypeOf(DateAfterModificationARM{}), generators)

	return dateAfterModificationARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterModificationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterModificationARM(gens map[string]gopter.Gen) {
	gens["DaysAfterLastAccessTimeGreaterThan"] = gen.PtrOf(gen.Int())
	gens["DaysAfterModificationGreaterThan"] = gen.PtrOf(gen.Int())
}
