// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401storage

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
<<<<<<< HEAD
	gens["Spec"] = StorageAccountsQueueServicesQueue_SpecGenerator()
	gens["Status"] = StorageAccountsQueueServicesQueue_STATUSGenerator()
}

func Test_StorageAccountsQueueServicesQueue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
=======
	gens["Spec"] = StorageAccounts_QueueServices_Queues_SpecGenerator()
	gens["Status"] = StorageQueue_STATUSGenerator()
}

func Test_StorageAccounts_QueueServices_Queues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>> main
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of StorageAccountsQueueServicesQueue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec, StorageAccountsQueueServicesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec runs a test to see if a specific instance of StorageAccountsQueueServicesQueue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue_Spec(subject StorageAccountsQueueServicesQueue_Spec) string {
=======
		"Round trip of StorageAccounts_QueueServices_Queues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec, StorageAccounts_QueueServices_Queues_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec runs a test to see if a specific instance of StorageAccounts_QueueServices_Queues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_QueueServices_Queues_Spec(subject StorageAccounts_QueueServices_Queues_Spec) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<< HEAD
	var actual StorageAccountsQueueServicesQueue_Spec
=======
	var actual StorageAccounts_QueueServices_Queues_Spec
>>>>>>> main
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

<<<<<<< HEAD
// Generator of StorageAccountsQueueServicesQueue_Spec instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueue_SpecGenerator()
var storageAccountsQueueServicesQueue_SpecGenerator gopter.Gen

// StorageAccountsQueueServicesQueue_SpecGenerator returns a generator of StorageAccountsQueueServicesQueue_Spec instances for property testing.
func StorageAccountsQueueServicesQueue_SpecGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueue_SpecGenerator != nil {
		return storageAccountsQueueServicesQueue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(generators)
	storageAccountsQueueServicesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue_Spec{}), generators)

	return storageAccountsQueueServicesQueue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueue_Spec(gens map[string]gopter.Gen) {
=======
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
>>>>>>> main
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_StorageAccountsQueueServicesQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of StorageAccountsQueueServicesQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS, StorageAccountsQueueServicesQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS runs a test to see if a specific instance of StorageAccountsQueueServicesQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue_STATUS(subject StorageAccountsQueueServicesQueue_STATUS) string {
=======
		"Round trip of StorageQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageQueue_STATUS, StorageQueue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageQueue_STATUS runs a test to see if a specific instance of StorageQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageQueue_STATUS(subject StorageQueue_STATUS) string {
>>>>>>> main
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

<<<<<<< HEAD
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
=======
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
>>>>>>> main
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
