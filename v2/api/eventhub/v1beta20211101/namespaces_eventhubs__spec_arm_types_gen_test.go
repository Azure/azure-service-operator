// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_NamespacesEventhubs_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecARM, NamespacesEventhubsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecARM runs a test to see if a specific instance of NamespacesEventhubs_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecARM(subject NamespacesEventhubs_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_SpecARM
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

// Generator of NamespacesEventhubs_SpecARM instances for property testing - lazily instantiated by
// NamespacesEventhubsSpecARMGenerator()
var namespacesEventhubsSpecARMGenerator gopter.Gen

// NamespacesEventhubsSpecARMGenerator returns a generator of NamespacesEventhubs_SpecARM instances for property testing.
// We first initialize namespacesEventhubsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecARMGenerator() gopter.Gen {
	if namespacesEventhubsSpecARMGenerator != nil {
		return namespacesEventhubsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecARM(generators)
	namespacesEventhubsSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecARM(generators)
	namespacesEventhubsSpecARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_SpecARM{}), generators)

	return namespacesEventhubsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamespacesEventhubsSpecPropertiesARMGenerator())
}

func Test_NamespacesEventhubs_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesARM, NamespacesEventhubsSpecPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesARM runs a test to see if a specific instance of NamespacesEventhubs_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesARM(subject NamespacesEventhubs_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec_PropertiesARM
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

// Generator of NamespacesEventhubs_Spec_PropertiesARM instances for property testing - lazily instantiated by
// NamespacesEventhubsSpecPropertiesARMGenerator()
var namespacesEventhubsSpecPropertiesARMGenerator gopter.Gen

// NamespacesEventhubsSpecPropertiesARMGenerator returns a generator of NamespacesEventhubs_Spec_PropertiesARM instances for property testing.
// We first initialize namespacesEventhubsSpecPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecPropertiesARMGenerator() gopter.Gen {
	if namespacesEventhubsSpecPropertiesARMGenerator != nil {
		return namespacesEventhubsSpecPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM(generators)
	namespacesEventhubsSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM(generators)
	namespacesEventhubsSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_PropertiesARM{}), generators)

	return namespacesEventhubsSpecPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator())
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM, NamespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM runs a test to see if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(subject NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM
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

// Generator of NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM instances for property testing - lazily
// instantiated by NamespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator()
var namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator gopter.Gen

// NamespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator returns a generator of NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM instances for property testing.
// We first initialize namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator() gopter.Gen {
	if namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator != nil {
		return namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM{}), generators)

	return namespacesEventhubsSpecPropertiesCaptureDescriptionARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionEncodingAvro, NamespacesEventhubsSpecPropertiesCaptureDescriptionEncodingAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionARM(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator())
}

func Test_NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM, NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM runs a test to see if a specific instance of NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(subject NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM
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

// Generator of NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM instances for property testing -
// lazily instantiated by NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator()
var namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator gopter.Gen

// NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator returns a generator of NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM instances for property testing.
// We first initialize namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator() gopter.Gen {
	if namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator != nil {
		return namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(generators)
	namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM{}), generators)

	return namespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsSpecPropertiesCaptureDescriptionDestinationARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DestinationPropertiesARMGenerator())
}

func Test_DestinationPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DestinationPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestinationPropertiesARM, DestinationPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestinationPropertiesARM runs a test to see if a specific instance of DestinationPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestinationPropertiesARM(subject DestinationPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DestinationPropertiesARM
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

// Generator of DestinationPropertiesARM instances for property testing - lazily instantiated by
// DestinationPropertiesARMGenerator()
var destinationPropertiesARMGenerator gopter.Gen

// DestinationPropertiesARMGenerator returns a generator of DestinationPropertiesARM instances for property testing.
func DestinationPropertiesARMGenerator() gopter.Gen {
	if destinationPropertiesARMGenerator != nil {
		return destinationPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationPropertiesARM(generators)
	destinationPropertiesARMGenerator = gen.Struct(reflect.TypeOf(DestinationPropertiesARM{}), generators)

	return destinationPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDestinationPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestinationPropertiesARM(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}
