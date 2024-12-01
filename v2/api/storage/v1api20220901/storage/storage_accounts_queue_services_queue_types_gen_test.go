// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
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

func Test_StorageAccountsQueueServicesQueue_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsQueueServicesQueue tests if a specific instance of StorageAccountsQueueServicesQueue round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.StorageAccountsQueueServicesQueue
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsQueueServicesQueue
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

func Test_StorageAccountsQueueServicesQueue_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue to StorageAccountsQueueServicesQueue via AssignProperties_To_StorageAccountsQueueServicesQueue & AssignProperties_From_StorageAccountsQueueServicesQueue returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue tests if a specific instance of StorageAccountsQueueServicesQueue can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccountsQueueServicesQueue
	err := copied.AssignProperties_To_StorageAccountsQueueServicesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueue
	err = actual.AssignProperties_From_StorageAccountsQueueServicesQueue(&other)
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

func Test_StorageAccountsQueueServicesQueue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue runs a test to see if a specific instance of StorageAccountsQueueServicesQueue round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueue
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

// Generator of StorageAccountsQueueServicesQueue instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueueGenerator()
var storageAccountsQueueServicesQueueGenerator gopter.Gen

// StorageAccountsQueueServicesQueueGenerator returns a generator of StorageAccountsQueueServicesQueue instances for property testing.
func StorageAccountsQueueServicesQueueGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueueGenerator != nil {
		return storageAccountsQueueServicesQueueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue(generators)
	storageAccountsQueueServicesQueueGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue{}), generators)

	return storageAccountsQueueServicesQueueGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsQueueServicesQueue_SpecGenerator()
	gens["Status"] = StorageAccountsQueueServicesQueue_STATUSGenerator()
}

func Test_StorageAccountsQueueServicesQueueOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueueOperatorSpec to StorageAccountsQueueServicesQueueOperatorSpec via AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec & AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueueOperatorSpec, StorageAccountsQueueServicesQueueOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueueOperatorSpec tests if a specific instance of StorageAccountsQueueServicesQueueOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueueOperatorSpec(subject StorageAccountsQueueServicesQueueOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccountsQueueServicesQueueOperatorSpec
	err := copied.AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueueOperatorSpec
	err = actual.AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec(&other)
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

func Test_StorageAccountsQueueServicesQueueOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueueOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueueOperatorSpec, StorageAccountsQueueServicesQueueOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueueOperatorSpec runs a test to see if a specific instance of StorageAccountsQueueServicesQueueOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueueOperatorSpec(subject StorageAccountsQueueServicesQueueOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueueOperatorSpec
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

// Generator of StorageAccountsQueueServicesQueueOperatorSpec instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueueOperatorSpecGenerator()
var storageAccountsQueueServicesQueueOperatorSpecGenerator gopter.Gen

// StorageAccountsQueueServicesQueueOperatorSpecGenerator returns a generator of StorageAccountsQueueServicesQueueOperatorSpec instances for property testing.
func StorageAccountsQueueServicesQueueOperatorSpecGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueueOperatorSpecGenerator != nil {
		return storageAccountsQueueServicesQueueOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	storageAccountsQueueServicesQueueOperatorSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueueOperatorSpec{}), generators)

	return storageAccountsQueueServicesQueueOperatorSpecGenerator
}

func Test_StorageAccountsQueueServicesQueue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue_STATUS to StorageAccountsQueueServicesQueue_STATUS via AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS & AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_STATUS, StorageAccountsQueueServicesQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_STATUS tests if a specific instance of StorageAccountsQueueServicesQueue_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_STATUS(subject StorageAccountsQueueServicesQueue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccountsQueueServicesQueue_STATUS
	err := copied.AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueue_STATUS
	err = actual.AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS(&other)
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

func Test_StorageAccountsQueueServicesQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS, StorageAccountsQueueServicesQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS runs a test to see if a specific instance of StorageAccountsQueueServicesQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS(subject StorageAccountsQueueServicesQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueue_STATUS
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

// Generator of StorageAccountsQueueServicesQueue_STATUS instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueue_STATUSGenerator()
var storageAccountsQueueServicesQueue_STATUSGenerator gopter.Gen

// StorageAccountsQueueServicesQueue_STATUSGenerator returns a generator of StorageAccountsQueueServicesQueue_STATUS instances for property testing.
func StorageAccountsQueueServicesQueue_STATUSGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueue_STATUSGenerator != nil {
		return storageAccountsQueueServicesQueue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_STATUS(generators)
	storageAccountsQueueServicesQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_STATUS{}), generators)

	return storageAccountsQueueServicesQueue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_STATUS(gens map[string]gopter.Gen) {
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_StorageAccountsQueueServicesQueue_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue_Spec to StorageAccountsQueueServicesQueue_Spec via AssignProperties_To_StorageAccountsQueueServicesQueue_Spec & AssignProperties_From_StorageAccountsQueueServicesQueue_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_Spec, StorageAccountsQueueServicesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_Spec tests if a specific instance of StorageAccountsQueueServicesQueue_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue_Spec(subject StorageAccountsQueueServicesQueue_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.StorageAccountsQueueServicesQueue_Spec
	err := copied.AssignProperties_To_StorageAccountsQueueServicesQueue_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueue_Spec
	err = actual.AssignProperties_From_StorageAccountsQueueServicesQueue_Spec(&other)
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

func Test_StorageAccountsQueueServicesQueue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec, StorageAccountsQueueServicesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec runs a test to see if a specific instance of StorageAccountsQueueServicesQueue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec(subject StorageAccountsQueueServicesQueue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueue_Spec
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

// Generator of StorageAccountsQueueServicesQueue_Spec instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueue_SpecGenerator()
var storageAccountsQueueServicesQueue_SpecGenerator gopter.Gen

// StorageAccountsQueueServicesQueue_SpecGenerator returns a generator of StorageAccountsQueueServicesQueue_Spec instances for property testing.
// We first initialize storageAccountsQueueServicesQueue_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesQueue_SpecGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueue_SpecGenerator != nil {
		return storageAccountsQueueServicesQueue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(generators)
	storageAccountsQueueServicesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(generators)
	storageAccountsQueueServicesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_Spec{}), generators)

	return storageAccountsQueueServicesQueue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(StorageAccountsQueueServicesQueueOperatorSpecGenerator())
}
