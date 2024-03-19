// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	v20240101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101/storage"
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

func Test_NamespacesEventhub_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20240101s.NamespacesEventhub
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhub
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

func Test_NamespacesEventhub_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhub to NamespacesEventhub via AssignProperties_To_NamespacesEventhub & AssignProperties_From_NamespacesEventhub returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhub tests if a specific instance of NamespacesEventhub can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.NamespacesEventhub
	err := copied.AssignProperties_To_NamespacesEventhub(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhub
	err = actual.AssignProperties_From_NamespacesEventhub(&other)
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

func Test_NamespacesEventhub_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhub via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhub, NamespacesEventhubGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhub runs a test to see if a specific instance of NamespacesEventhub round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhub(subject NamespacesEventhub) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhub
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

// Generator of NamespacesEventhub instances for property testing - lazily instantiated by NamespacesEventhubGenerator()
var namespacesEventhubGenerator gopter.Gen

// NamespacesEventhubGenerator returns a generator of NamespacesEventhub instances for property testing.
func NamespacesEventhubGenerator() gopter.Gen {
	if namespacesEventhubGenerator != nil {
		return namespacesEventhubGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhub(generators)
	namespacesEventhubGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhub{}), generators)

	return namespacesEventhubGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhub is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhub(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Eventhub_SpecGenerator()
	gens["Status"] = Namespaces_Eventhub_STATUSGenerator()
}

func Test_Namespaces_Eventhub_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Eventhub_Spec to Namespaces_Eventhub_Spec via AssignProperties_To_Namespaces_Eventhub_Spec & AssignProperties_From_Namespaces_Eventhub_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Eventhub_Spec, Namespaces_Eventhub_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Eventhub_Spec tests if a specific instance of Namespaces_Eventhub_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Eventhub_Spec(subject Namespaces_Eventhub_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.Namespaces_Eventhub_Spec
	err := copied.AssignProperties_To_Namespaces_Eventhub_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Eventhub_Spec
	err = actual.AssignProperties_From_Namespaces_Eventhub_Spec(&other)
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

func Test_Namespaces_Eventhub_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Spec, Namespaces_Eventhub_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Spec runs a test to see if a specific instance of Namespaces_Eventhub_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Spec(subject Namespaces_Eventhub_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Spec
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

// Generator of Namespaces_Eventhub_Spec instances for property testing - lazily instantiated by
// Namespaces_Eventhub_SpecGenerator()
var namespaces_Eventhub_SpecGenerator gopter.Gen

// Namespaces_Eventhub_SpecGenerator returns a generator of Namespaces_Eventhub_Spec instances for property testing.
// We first initialize namespaces_Eventhub_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_SpecGenerator() gopter.Gen {
	if namespaces_Eventhub_SpecGenerator != nil {
		return namespaces_Eventhub_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec(generators)
	namespaces_Eventhub_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec(generators)
	namespaces_Eventhub_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec{}), generators)

	return namespaces_Eventhub_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescriptionGenerator())
}

func Test_Namespaces_Eventhub_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Eventhub_STATUS to Namespaces_Eventhub_STATUS via AssignProperties_To_Namespaces_Eventhub_STATUS & AssignProperties_From_Namespaces_Eventhub_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Eventhub_STATUS, Namespaces_Eventhub_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Eventhub_STATUS tests if a specific instance of Namespaces_Eventhub_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Eventhub_STATUS(subject Namespaces_Eventhub_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.Namespaces_Eventhub_STATUS
	err := copied.AssignProperties_To_Namespaces_Eventhub_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Eventhub_STATUS
	err = actual.AssignProperties_From_Namespaces_Eventhub_STATUS(&other)
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

func Test_Namespaces_Eventhub_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_STATUS, Namespaces_Eventhub_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_STATUS runs a test to see if a specific instance of Namespaces_Eventhub_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_STATUS(subject Namespaces_Eventhub_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_STATUS
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

// Generator of Namespaces_Eventhub_STATUS instances for property testing - lazily instantiated by
// Namespaces_Eventhub_STATUSGenerator()
var namespaces_Eventhub_STATUSGenerator gopter.Gen

// Namespaces_Eventhub_STATUSGenerator returns a generator of Namespaces_Eventhub_STATUS instances for property testing.
// We first initialize namespaces_Eventhub_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_STATUSGenerator() gopter.Gen {
	if namespaces_Eventhub_STATUSGenerator != nil {
		return namespaces_Eventhub_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_STATUS(generators)
	namespaces_Eventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_STATUS(generators)
	namespaces_Eventhub_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_STATUS{}), generators)

	return namespaces_Eventhub_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_STATUS(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescription_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_CaptureDescription_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CaptureDescription to CaptureDescription via AssignProperties_To_CaptureDescription & AssignProperties_From_CaptureDescription returns original",
		prop.ForAll(RunPropertyAssignmentTestForCaptureDescription, CaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCaptureDescription tests if a specific instance of CaptureDescription can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForCaptureDescription(subject CaptureDescription) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.CaptureDescription
	err := copied.AssignProperties_To_CaptureDescription(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CaptureDescription
	err = actual.AssignProperties_From_CaptureDescription(&other)
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

func Test_CaptureDescription_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription, CaptureDescriptionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription runs a test to see if a specific instance of CaptureDescription round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription(subject CaptureDescription) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription
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

// Generator of CaptureDescription instances for property testing - lazily instantiated by CaptureDescriptionGenerator()
var captureDescriptionGenerator gopter.Gen

// CaptureDescriptionGenerator returns a generator of CaptureDescription instances for property testing.
// We first initialize captureDescriptionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescriptionGenerator() gopter.Gen {
	if captureDescriptionGenerator != nil {
		return captureDescriptionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription(generators)
	captureDescriptionGenerator = gen.Struct(reflect.TypeOf(CaptureDescription{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription(generators)
	captureDescriptionGenerator = gen.Struct(reflect.TypeOf(CaptureDescription{}), generators)

	return captureDescriptionGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.AlphaString())
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(DestinationGenerator())
}

func Test_CaptureDescription_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CaptureDescription_STATUS to CaptureDescription_STATUS via AssignProperties_To_CaptureDescription_STATUS & AssignProperties_From_CaptureDescription_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForCaptureDescription_STATUS, CaptureDescription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCaptureDescription_STATUS tests if a specific instance of CaptureDescription_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForCaptureDescription_STATUS(subject CaptureDescription_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.CaptureDescription_STATUS
	err := copied.AssignProperties_To_CaptureDescription_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CaptureDescription_STATUS
	err = actual.AssignProperties_From_CaptureDescription_STATUS(&other)
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

func Test_CaptureDescription_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription_STATUS, CaptureDescription_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription_STATUS runs a test to see if a specific instance of CaptureDescription_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription_STATUS(subject CaptureDescription_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_STATUS
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

// Generator of CaptureDescription_STATUS instances for property testing - lazily instantiated by
// CaptureDescription_STATUSGenerator()
var captureDescription_STATUSGenerator gopter.Gen

// CaptureDescription_STATUSGenerator returns a generator of CaptureDescription_STATUS instances for property testing.
// We first initialize captureDescription_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescription_STATUSGenerator() gopter.Gen {
	if captureDescription_STATUSGenerator != nil {
		return captureDescription_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(generators)
	captureDescription_STATUSGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_STATUS{}), generators)

	return captureDescription_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.AlphaString())
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription_STATUS(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Destination_STATUSGenerator())
}

func Test_Destination_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Destination to Destination via AssignProperties_To_Destination & AssignProperties_From_Destination returns original",
		prop.ForAll(RunPropertyAssignmentTestForDestination, DestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDestination tests if a specific instance of Destination can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDestination(subject Destination) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.Destination
	err := copied.AssignProperties_To_Destination(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Destination
	err = actual.AssignProperties_From_Destination(&other)
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

func Test_Destination_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination, DestinationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination runs a test to see if a specific instance of Destination round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination(subject Destination) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination
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

// Generator of Destination instances for property testing - lazily instantiated by DestinationGenerator()
var destinationGenerator gopter.Gen

// DestinationGenerator returns a generator of Destination instances for property testing.
func DestinationGenerator() gopter.Gen {
	if destinationGenerator != nil {
		return destinationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination(generators)
	destinationGenerator = gen.Struct(reflect.TypeOf(Destination{}), generators)

	return destinationGenerator
}

// AddIndependentPropertyGeneratorsForDestination is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Destination_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Destination_STATUS to Destination_STATUS via AssignProperties_To_Destination_STATUS & AssignProperties_From_Destination_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDestination_STATUS, Destination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDestination_STATUS tests if a specific instance of Destination_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDestination_STATUS(subject Destination_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20240101s.Destination_STATUS
	err := copied.AssignProperties_To_Destination_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Destination_STATUS
	err = actual.AssignProperties_From_Destination_STATUS(&other)
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

func Test_Destination_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_STATUS, Destination_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_STATUS runs a test to see if a specific instance of Destination_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_STATUS(subject Destination_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_STATUS
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

// Generator of Destination_STATUS instances for property testing - lazily instantiated by Destination_STATUSGenerator()
var destination_STATUSGenerator gopter.Gen

// Destination_STATUSGenerator returns a generator of Destination_STATUS instances for property testing.
func Destination_STATUSGenerator() gopter.Gen {
	if destination_STATUSGenerator != nil {
		return destination_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_STATUS(generators)
	destination_STATUSGenerator = gen.Struct(reflect.TypeOf(Destination_STATUS{}), generators)

	return destination_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDestination_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_STATUS(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}
