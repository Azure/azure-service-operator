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

func Test_StorageAccountsQueueService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueService runs a test to see if a specific instance of StorageAccountsQueueService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueService
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

// Generator of StorageAccountsQueueService instances for property testing - lazily instantiated by
// StorageAccountsQueueServiceGenerator()
var storageAccountsQueueServiceGenerator gopter.Gen

// StorageAccountsQueueServiceGenerator returns a generator of StorageAccountsQueueService instances for property testing.
func StorageAccountsQueueServiceGenerator() gopter.Gen {
	if storageAccountsQueueServiceGenerator != nil {
		return storageAccountsQueueServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueService(generators)
	storageAccountsQueueServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService{}), generators)

	return storageAccountsQueueServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsQueueServicesSpecGenerator()
	gens["Status"] = QueueServicePropertiesSTATUSGenerator()
}

func Test_QueueServiceProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueServiceProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueServicePropertiesSTATUS, QueueServicePropertiesSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueServicePropertiesSTATUS runs a test to see if a specific instance of QueueServiceProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueServicePropertiesSTATUS(subject QueueServiceProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueServiceProperties_STATUS
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

// Generator of QueueServiceProperties_STATUS instances for property testing - lazily instantiated by
// QueueServicePropertiesSTATUSGenerator()
var queueServicePropertiesSTATUSGenerator gopter.Gen

// QueueServicePropertiesSTATUSGenerator returns a generator of QueueServiceProperties_STATUS instances for property testing.
// We first initialize queueServicePropertiesSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func QueueServicePropertiesSTATUSGenerator() gopter.Gen {
	if queueServicePropertiesSTATUSGenerator != nil {
		return queueServicePropertiesSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUS(generators)
	queueServicePropertiesSTATUSGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUS(generators)
	AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUS(generators)
	queueServicePropertiesSTATUSGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_STATUS{}), generators)

	return queueServicePropertiesSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesSTATUSGenerator())
}

func Test_StorageAccountsQueueServices_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServices_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesSpec, StorageAccountsQueueServicesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesSpec runs a test to see if a specific instance of StorageAccountsQueueServices_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesSpec(subject StorageAccountsQueueServices_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServices_Spec
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

// Generator of StorageAccountsQueueServices_Spec instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesSpecGenerator()
var storageAccountsQueueServicesSpecGenerator gopter.Gen

// StorageAccountsQueueServicesSpecGenerator returns a generator of StorageAccountsQueueServices_Spec instances for property testing.
// We first initialize storageAccountsQueueServicesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsQueueServicesSpecGenerator() gopter.Gen {
	if storageAccountsQueueServicesSpecGenerator != nil {
		return storageAccountsQueueServicesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSpec(generators)
	storageAccountsQueueServicesSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServices_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSpec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSpec(generators)
	storageAccountsQueueServicesSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServices_Spec{}), generators)

	return storageAccountsQueueServicesSpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesSpec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
}
