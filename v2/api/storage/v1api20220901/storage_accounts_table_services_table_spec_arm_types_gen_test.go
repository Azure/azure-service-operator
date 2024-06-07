// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

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

func Test_StorageAccounts_TableServices_Table_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableServices_Table_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec_ARM, StorageAccounts_TableServices_Table_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec_ARM runs a test to see if a specific instance of StorageAccounts_TableServices_Table_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableServices_Table_Spec_ARM(subject StorageAccounts_TableServices_Table_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableServices_Table_Spec_ARM
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

// Generator of StorageAccounts_TableServices_Table_Spec_ARM instances for property testing - lazily instantiated by
// StorageAccounts_TableServices_Table_Spec_ARMGenerator()
var storageAccounts_TableServices_Table_Spec_ARMGenerator gopter.Gen

// StorageAccounts_TableServices_Table_Spec_ARMGenerator returns a generator of StorageAccounts_TableServices_Table_Spec_ARM instances for property testing.
// We first initialize storageAccounts_TableServices_Table_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_TableServices_Table_Spec_ARMGenerator() gopter.Gen {
	if storageAccounts_TableServices_Table_Spec_ARMGenerator != nil {
		return storageAccounts_TableServices_Table_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM(generators)
	storageAccounts_TableServices_Table_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM(generators)
	storageAccounts_TableServices_Table_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableServices_Table_Spec_ARM{}), generators)

	return storageAccounts_TableServices_Table_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableServices_Table_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TableProperties_ARMGenerator())
}

func Test_TableAccessPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy_ARM, TableAccessPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy_ARM runs a test to see if a specific instance of TableAccessPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy_ARM(subject TableAccessPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy_ARM
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

// Generator of TableAccessPolicy_ARM instances for property testing - lazily instantiated by
// TableAccessPolicy_ARMGenerator()
var tableAccessPolicy_ARMGenerator gopter.Gen

// TableAccessPolicy_ARMGenerator returns a generator of TableAccessPolicy_ARM instances for property testing.
func TableAccessPolicy_ARMGenerator() gopter.Gen {
	if tableAccessPolicy_ARMGenerator != nil {
		return tableAccessPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy_ARM(generators)
	tableAccessPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy_ARM{}), generators)

	return tableAccessPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy_ARM(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableProperties_ARM, TableProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableProperties_ARM runs a test to see if a specific instance of TableProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableProperties_ARM(subject TableProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableProperties_ARM
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

// Generator of TableProperties_ARM instances for property testing - lazily instantiated by
// TableProperties_ARMGenerator()
var tableProperties_ARMGenerator gopter.Gen

// TableProperties_ARMGenerator returns a generator of TableProperties_ARM instances for property testing.
func TableProperties_ARMGenerator() gopter.Gen {
	if tableProperties_ARMGenerator != nil {
		return tableProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTableProperties_ARM(generators)
	tableProperties_ARMGenerator = gen.Struct(reflect.TypeOf(TableProperties_ARM{}), generators)

	return tableProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForTableProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableProperties_ARM(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifier_ARMGenerator())
}

func Test_TableSignedIdentifier_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier_ARM, TableSignedIdentifier_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier_ARM runs a test to see if a specific instance of TableSignedIdentifier_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier_ARM(subject TableSignedIdentifier_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier_ARM
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

// Generator of TableSignedIdentifier_ARM instances for property testing - lazily instantiated by
// TableSignedIdentifier_ARMGenerator()
var tableSignedIdentifier_ARMGenerator gopter.Gen

// TableSignedIdentifier_ARMGenerator returns a generator of TableSignedIdentifier_ARM instances for property testing.
// We first initialize tableSignedIdentifier_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableSignedIdentifier_ARMGenerator() gopter.Gen {
	if tableSignedIdentifier_ARMGenerator != nil {
		return tableSignedIdentifier_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_ARM(generators)
	tableSignedIdentifier_ARMGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_ARM(generators)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier_ARM(generators)
	tableSignedIdentifier_ARMGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_ARM{}), generators)

	return tableSignedIdentifier_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableSignedIdentifier_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableSignedIdentifier_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier_ARM(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicy_ARMGenerator())
}
