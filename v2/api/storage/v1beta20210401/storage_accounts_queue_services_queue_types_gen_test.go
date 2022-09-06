// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import (
	"encoding/json"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
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
	var hub v20210401s.StorageAccountsQueueServicesQueue
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

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue tests if a specific instance of StorageAccountsQueueServicesQueue can be assigned to v1beta20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210401s.StorageAccountsQueueServicesQueue
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
	match := cmp.Equal(subject, actual)
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
	gens["Spec"] = StorageAccounts_QueueServices_Queues_SpecGenerator()
	gens["Status"] = StorageQueue_STATUSGenerator()
}

func Test_StorageAccounts_QueueServices_Queues_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_QueueServices_Queues_Spec to StorageAccounts_QueueServices_Queues_Spec via AssignProperties_To_StorageAccounts_QueueServices_Queues_Spec & AssignProperties_From_StorageAccounts_QueueServices_Queues_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_QueueServices_Queues_Spec, StorageAccounts_QueueServices_Queues_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_QueueServices_Queues_Spec tests if a specific instance of StorageAccounts_QueueServices_Queues_Spec can be assigned to v1beta20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_QueueServices_Queues_Spec(subject StorageAccounts_QueueServices_Queues_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210401s.StorageAccounts_QueueServices_Queues_Spec
	err := copied.AssignProperties_To_StorageAccounts_QueueServices_Queues_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_QueueServices_Queues_Spec
	err = actual.AssignProperties_From_StorageAccounts_QueueServices_Queues_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccounts_QueueServices_Queues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_QueueServices_Queues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec, StorageAccounts_QueueServices_Queues_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec runs a test to see if a specific instance of StorageAccounts_QueueServices_Queues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec(subject StorageAccounts_QueueServices_Queues_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_QueueServices_Queues_Spec
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

// Generator of StorageAccounts_QueueServices_Queues_Spec instances for property testing - lazily instantiated by
// StorageAccounts_QueueServices_Queues_SpecGenerator()
var storageAccounts_QueueServices_Queues_SpecGenerator gopter.Gen

// StorageAccounts_QueueServices_Queues_SpecGenerator returns a generator of StorageAccounts_QueueServices_Queues_Spec instances for property testing.
func StorageAccounts_QueueServices_Queues_SpecGenerator() gopter.Gen {
	if storageAccounts_QueueServices_Queues_SpecGenerator != nil {
		return storageAccounts_QueueServices_Queues_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_QueueServices_Queues_Spec(generators)
	storageAccounts_QueueServices_Queues_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueServices_Queues_Spec{}), generators)

	return storageAccounts_QueueServices_Queues_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_QueueServices_Queues_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_QueueServices_Queues_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_StorageQueue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageQueue_STATUS to StorageQueue_STATUS via AssignProperties_To_StorageQueue_STATUS & AssignProperties_From_StorageQueue_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageQueue_STATUS, StorageQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageQueue_STATUS tests if a specific instance of StorageQueue_STATUS can be assigned to v1beta20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageQueue_STATUS(subject StorageQueue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210401s.StorageQueue_STATUS
	err := copied.AssignProperties_To_StorageQueue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageQueue_STATUS
	err = actual.AssignProperties_From_StorageQueue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageQueue_STATUS, StorageQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageQueue_STATUS runs a test to see if a specific instance of StorageQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageQueue_STATUS(subject StorageQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageQueue_STATUS
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

// Generator of StorageQueue_STATUS instances for property testing - lazily instantiated by
// StorageQueue_STATUSGenerator()
var storageQueue_STATUSGenerator gopter.Gen

// StorageQueue_STATUSGenerator returns a generator of StorageQueue_STATUS instances for property testing.
func StorageQueue_STATUSGenerator() gopter.Gen {
	if storageQueue_STATUSGenerator != nil {
		return storageQueue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageQueue_STATUS(generators)
	storageQueue_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageQueue_STATUS{}), generators)

	return storageQueue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageQueue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageQueue_STATUS(gens map[string]gopter.Gen) {
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
