// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

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

func Test_StorageAccountsTableServicesTable_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS_ARM, StorageAccountsTableServicesTable_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS_ARM runs a test to see if a specific instance of StorageAccountsTableServicesTable_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS_ARM(subject StorageAccountsTableServicesTable_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable_STATUS_ARM
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

// Generator of StorageAccountsTableServicesTable_STATUS_ARM instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTable_STATUS_ARMGenerator()
var storageAccountsTableServicesTable_STATUS_ARMGenerator gopter.Gen

// StorageAccountsTableServicesTable_STATUS_ARMGenerator returns a generator of StorageAccountsTableServicesTable_STATUS_ARM instances for property testing.
// We first initialize storageAccountsTableServicesTable_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableServicesTable_STATUS_ARMGenerator() gopter.Gen {
	if storageAccountsTableServicesTable_STATUS_ARMGenerator != nil {
		return storageAccountsTableServicesTable_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM(generators)
	storageAccountsTableServicesTable_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM(generators)
	storageAccountsTableServicesTable_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS_ARM{}), generators)

	return storageAccountsTableServicesTable_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TableProperties_STATUS_ARMGenerator())
}

func Test_TableAccessPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy_STATUS_ARM, TableAccessPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy_STATUS_ARM runs a test to see if a specific instance of TableAccessPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy_STATUS_ARM(subject TableAccessPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy_STATUS_ARM
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

// Generator of TableAccessPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// TableAccessPolicy_STATUS_ARMGenerator()
var tableAccessPolicy_STATUS_ARMGenerator gopter.Gen

// TableAccessPolicy_STATUS_ARMGenerator returns a generator of TableAccessPolicy_STATUS_ARM instances for property testing.
func TableAccessPolicy_STATUS_ARMGenerator() gopter.Gen {
	if tableAccessPolicy_STATUS_ARMGenerator != nil {
		return tableAccessPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS_ARM(generators)
	tableAccessPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy_STATUS_ARM{}), generators)

	return tableAccessPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableProperties_STATUS_ARM, TableProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableProperties_STATUS_ARM runs a test to see if a specific instance of TableProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableProperties_STATUS_ARM(subject TableProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableProperties_STATUS_ARM
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

// Generator of TableProperties_STATUS_ARM instances for property testing - lazily instantiated by
// TableProperties_STATUS_ARMGenerator()
var tableProperties_STATUS_ARMGenerator gopter.Gen

// TableProperties_STATUS_ARMGenerator returns a generator of TableProperties_STATUS_ARM instances for property testing.
// We first initialize tableProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableProperties_STATUS_ARMGenerator() gopter.Gen {
	if tableProperties_STATUS_ARMGenerator != nil {
		return tableProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableProperties_STATUS_ARM(generators)
	tableProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForTableProperties_STATUS_ARM(generators)
	tableProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableProperties_STATUS_ARM{}), generators)

	return tableProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["TableName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifier_STATUS_ARMGenerator())
}

func Test_TableSignedIdentifier_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier_STATUS_ARM, TableSignedIdentifier_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier_STATUS_ARM runs a test to see if a specific instance of TableSignedIdentifier_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier_STATUS_ARM(subject TableSignedIdentifier_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier_STATUS_ARM
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

// Generator of TableSignedIdentifier_STATUS_ARM instances for property testing - lazily instantiated by
// TableSignedIdentifier_STATUS_ARMGenerator()
var tableSignedIdentifier_STATUS_ARMGenerator gopter.Gen

// TableSignedIdentifier_STATUS_ARMGenerator returns a generator of TableSignedIdentifier_STATUS_ARM instances for property testing.
// We first initialize tableSignedIdentifier_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableSignedIdentifier_STATUS_ARMGenerator() gopter.Gen {
	if tableSignedIdentifier_STATUS_ARMGenerator != nil {
		return tableSignedIdentifier_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM(generators)
	tableSignedIdentifier_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM(generators)
	tableSignedIdentifier_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS_ARM{}), generators)

	return tableSignedIdentifier_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicy_STATUS_ARMGenerator())
}
