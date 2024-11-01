// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_StorageAccountsTableServicesTable_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS, StorageAccountsTableServicesTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS runs a test to see if a specific instance of StorageAccountsTableServicesTable_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable_STATUS(subject StorageAccountsTableServicesTable_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable_STATUS
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

// Generator of StorageAccountsTableServicesTable_STATUS instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTable_STATUSGenerator()
var storageAccountsTableServicesTable_STATUSGenerator gopter.Gen

// StorageAccountsTableServicesTable_STATUSGenerator returns a generator of StorageAccountsTableServicesTable_STATUS instances for property testing.
// We first initialize storageAccountsTableServicesTable_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableServicesTable_STATUSGenerator() gopter.Gen {
	if storageAccountsTableServicesTable_STATUSGenerator != nil {
		return storageAccountsTableServicesTable_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	storageAccountsTableServicesTable_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(generators)
	storageAccountsTableServicesTable_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_STATUS{}), generators)

	return storageAccountsTableServicesTable_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TableProperties_STATUSGenerator())
}

func Test_TableAccessPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy_STATUS, TableAccessPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy_STATUS runs a test to see if a specific instance of TableAccessPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy_STATUS(subject TableAccessPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy_STATUS
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

// Generator of TableAccessPolicy_STATUS instances for property testing - lazily instantiated by
// TableAccessPolicy_STATUSGenerator()
var tableAccessPolicy_STATUSGenerator gopter.Gen

// TableAccessPolicy_STATUSGenerator returns a generator of TableAccessPolicy_STATUS instances for property testing.
func TableAccessPolicy_STATUSGenerator() gopter.Gen {
	if tableAccessPolicy_STATUSGenerator != nil {
		return tableAccessPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS(generators)
	tableAccessPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy_STATUS{}), generators)

	return tableAccessPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableProperties_STATUS, TableProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableProperties_STATUS runs a test to see if a specific instance of TableProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableProperties_STATUS(subject TableProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableProperties_STATUS
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

// Generator of TableProperties_STATUS instances for property testing - lazily instantiated by
// TableProperties_STATUSGenerator()
var tableProperties_STATUSGenerator gopter.Gen

// TableProperties_STATUSGenerator returns a generator of TableProperties_STATUS instances for property testing.
// We first initialize tableProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableProperties_STATUSGenerator() gopter.Gen {
	if tableProperties_STATUSGenerator != nil {
		return tableProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableProperties_STATUS(generators)
	tableProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(TableProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForTableProperties_STATUS(generators)
	tableProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(TableProperties_STATUS{}), generators)

	return tableProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableProperties_STATUS(gens map[string]gopter.Gen) {
	gens["TableName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableProperties_STATUS(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifier_STATUSGenerator())
}

func Test_TableSignedIdentifier_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier_STATUS, TableSignedIdentifier_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier_STATUS runs a test to see if a specific instance of TableSignedIdentifier_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier_STATUS(subject TableSignedIdentifier_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier_STATUS
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

// Generator of TableSignedIdentifier_STATUS instances for property testing - lazily instantiated by
// TableSignedIdentifier_STATUSGenerator()
var tableSignedIdentifier_STATUSGenerator gopter.Gen

// TableSignedIdentifier_STATUSGenerator returns a generator of TableSignedIdentifier_STATUS instances for property testing.
// We first initialize tableSignedIdentifier_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableSignedIdentifier_STATUSGenerator() gopter.Gen {
	if tableSignedIdentifier_STATUSGenerator != nil {
		return tableSignedIdentifier_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	tableSignedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS(generators)
	tableSignedIdentifier_STATUSGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier_STATUS{}), generators)

	return tableSignedIdentifier_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier_STATUS(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicy_STATUSGenerator())
}
