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

func Test_StorageAccounts_QueueService_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_QueueService_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_QueueService_Spec_ARM, StorageAccounts_QueueService_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_QueueService_Spec_ARM runs a test to see if a specific instance of StorageAccounts_QueueService_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_QueueService_Spec_ARM(subject StorageAccounts_QueueService_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_QueueService_Spec_ARM
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

// Generator of StorageAccounts_QueueService_Spec_ARM instances for property testing - lazily instantiated by
// StorageAccounts_QueueService_Spec_ARMGenerator()
var storageAccounts_QueueService_Spec_ARMGenerator gopter.Gen

// StorageAccounts_QueueService_Spec_ARMGenerator returns a generator of StorageAccounts_QueueService_Spec_ARM instances for property testing.
// We first initialize storageAccounts_QueueService_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_QueueService_Spec_ARMGenerator() gopter.Gen {
	if storageAccounts_QueueService_Spec_ARMGenerator != nil {
		return storageAccounts_QueueService_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM(generators)
	storageAccounts_QueueService_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueService_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM(generators)
	storageAccounts_QueueService_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueService_Spec_ARM{}), generators)

	return storageAccounts_QueueService_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueueServicePropertiesProperties_ARMGenerator())
}

func Test_QueueServicePropertiesProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueServicePropertiesProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueServicePropertiesProperties_ARM, QueueServicePropertiesProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueServicePropertiesProperties_ARM runs a test to see if a specific instance of QueueServicePropertiesProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueServicePropertiesProperties_ARM(subject QueueServicePropertiesProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueServicePropertiesProperties_ARM
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

// Generator of QueueServicePropertiesProperties_ARM instances for property testing - lazily instantiated by
// QueueServicePropertiesProperties_ARMGenerator()
var queueServicePropertiesProperties_ARMGenerator gopter.Gen

// QueueServicePropertiesProperties_ARMGenerator returns a generator of QueueServicePropertiesProperties_ARM instances for property testing.
func QueueServicePropertiesProperties_ARMGenerator() gopter.Gen {
	if queueServicePropertiesProperties_ARMGenerator != nil {
		return queueServicePropertiesProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForQueueServicePropertiesProperties_ARM(generators)
	queueServicePropertiesProperties_ARMGenerator = gen.Struct(reflect.TypeOf(QueueServicePropertiesProperties_ARM{}), generators)

	return queueServicePropertiesProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForQueueServicePropertiesProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForQueueServicePropertiesProperties_ARM(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_ARMGenerator())
}
