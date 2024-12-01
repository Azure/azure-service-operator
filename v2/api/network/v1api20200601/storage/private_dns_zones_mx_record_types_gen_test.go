// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/storage"
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

func Test_PrivateDnsZonesMXRecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesMXRecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesMXRecord, PrivateDnsZonesMXRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesMXRecord tests if a specific instance of PrivateDnsZonesMXRecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesMXRecord(subject PrivateDnsZonesMXRecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.PrivateDnsZonesMXRecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesMXRecord
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

func Test_PrivateDnsZonesMXRecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesMXRecord to PrivateDnsZonesMXRecord via AssignProperties_To_PrivateDnsZonesMXRecord & AssignProperties_From_PrivateDnsZonesMXRecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesMXRecord, PrivateDnsZonesMXRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesMXRecord tests if a specific instance of PrivateDnsZonesMXRecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesMXRecord(subject PrivateDnsZonesMXRecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesMXRecord
	err := copied.AssignProperties_To_PrivateDnsZonesMXRecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesMXRecord
	err = actual.AssignProperties_From_PrivateDnsZonesMXRecord(&other)
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

func Test_PrivateDnsZonesMXRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesMXRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesMXRecord, PrivateDnsZonesMXRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesMXRecord runs a test to see if a specific instance of PrivateDnsZonesMXRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesMXRecord(subject PrivateDnsZonesMXRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesMXRecord
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

// Generator of PrivateDnsZonesMXRecord instances for property testing - lazily instantiated by
// PrivateDnsZonesMXRecordGenerator()
var privateDnsZonesMXRecordGenerator gopter.Gen

// PrivateDnsZonesMXRecordGenerator returns a generator of PrivateDnsZonesMXRecord instances for property testing.
func PrivateDnsZonesMXRecordGenerator() gopter.Gen {
	if privateDnsZonesMXRecordGenerator != nil {
		return privateDnsZonesMXRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord(generators)
	privateDnsZonesMXRecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecord{}), generators)

	return privateDnsZonesMXRecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZonesMXRecord_SpecGenerator()
	gens["Status"] = PrivateDnsZonesMXRecord_STATUSGenerator()
}

func Test_PrivateDnsZonesMXRecordOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesMXRecordOperatorSpec to PrivateDnsZonesMXRecordOperatorSpec via AssignProperties_To_PrivateDnsZonesMXRecordOperatorSpec & AssignProperties_From_PrivateDnsZonesMXRecordOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesMXRecordOperatorSpec, PrivateDnsZonesMXRecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesMXRecordOperatorSpec tests if a specific instance of PrivateDnsZonesMXRecordOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesMXRecordOperatorSpec(subject PrivateDnsZonesMXRecordOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesMXRecordOperatorSpec
	err := copied.AssignProperties_To_PrivateDnsZonesMXRecordOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesMXRecordOperatorSpec
	err = actual.AssignProperties_From_PrivateDnsZonesMXRecordOperatorSpec(&other)
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

func Test_PrivateDnsZonesMXRecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesMXRecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesMXRecordOperatorSpec, PrivateDnsZonesMXRecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesMXRecordOperatorSpec runs a test to see if a specific instance of PrivateDnsZonesMXRecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesMXRecordOperatorSpec(subject PrivateDnsZonesMXRecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesMXRecordOperatorSpec
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

// Generator of PrivateDnsZonesMXRecordOperatorSpec instances for property testing - lazily instantiated by
// PrivateDnsZonesMXRecordOperatorSpecGenerator()
var privateDnsZonesMXRecordOperatorSpecGenerator gopter.Gen

// PrivateDnsZonesMXRecordOperatorSpecGenerator returns a generator of PrivateDnsZonesMXRecordOperatorSpec instances for property testing.
func PrivateDnsZonesMXRecordOperatorSpecGenerator() gopter.Gen {
	if privateDnsZonesMXRecordOperatorSpecGenerator != nil {
		return privateDnsZonesMXRecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	privateDnsZonesMXRecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecordOperatorSpec{}), generators)

	return privateDnsZonesMXRecordOperatorSpecGenerator
}

func Test_PrivateDnsZonesMXRecord_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesMXRecord_STATUS to PrivateDnsZonesMXRecord_STATUS via AssignProperties_To_PrivateDnsZonesMXRecord_STATUS & AssignProperties_From_PrivateDnsZonesMXRecord_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_STATUS, PrivateDnsZonesMXRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_STATUS tests if a specific instance of PrivateDnsZonesMXRecord_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_STATUS(subject PrivateDnsZonesMXRecord_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesMXRecord_STATUS
	err := copied.AssignProperties_To_PrivateDnsZonesMXRecord_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesMXRecord_STATUS
	err = actual.AssignProperties_From_PrivateDnsZonesMXRecord_STATUS(&other)
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

func Test_PrivateDnsZonesMXRecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesMXRecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesMXRecord_STATUS, PrivateDnsZonesMXRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesMXRecord_STATUS runs a test to see if a specific instance of PrivateDnsZonesMXRecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesMXRecord_STATUS(subject PrivateDnsZonesMXRecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesMXRecord_STATUS
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

// Generator of PrivateDnsZonesMXRecord_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesMXRecord_STATUSGenerator()
var privateDnsZonesMXRecord_STATUSGenerator gopter.Gen

// PrivateDnsZonesMXRecord_STATUSGenerator returns a generator of PrivateDnsZonesMXRecord_STATUS instances for property testing.
// We first initialize privateDnsZonesMXRecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesMXRecord_STATUSGenerator() gopter.Gen {
	if privateDnsZonesMXRecord_STATUSGenerator != nil {
		return privateDnsZonesMXRecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS(generators)
	privateDnsZonesMXRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS(generators)
	privateDnsZonesMXRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecord_STATUS{}), generators)

	return privateDnsZonesMXRecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}

func Test_PrivateDnsZonesMXRecord_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesMXRecord_Spec to PrivateDnsZonesMXRecord_Spec via AssignProperties_To_PrivateDnsZonesMXRecord_Spec & AssignProperties_From_PrivateDnsZonesMXRecord_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_Spec, PrivateDnsZonesMXRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_Spec tests if a specific instance of PrivateDnsZonesMXRecord_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesMXRecord_Spec(subject PrivateDnsZonesMXRecord_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesMXRecord_Spec
	err := copied.AssignProperties_To_PrivateDnsZonesMXRecord_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesMXRecord_Spec
	err = actual.AssignProperties_From_PrivateDnsZonesMXRecord_Spec(&other)
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

func Test_PrivateDnsZonesMXRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesMXRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesMXRecord_Spec, PrivateDnsZonesMXRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesMXRecord_Spec runs a test to see if a specific instance of PrivateDnsZonesMXRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesMXRecord_Spec(subject PrivateDnsZonesMXRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesMXRecord_Spec
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

// Generator of PrivateDnsZonesMXRecord_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesMXRecord_SpecGenerator()
var privateDnsZonesMXRecord_SpecGenerator gopter.Gen

// PrivateDnsZonesMXRecord_SpecGenerator returns a generator of PrivateDnsZonesMXRecord_Spec instances for property testing.
// We first initialize privateDnsZonesMXRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesMXRecord_SpecGenerator() gopter.Gen {
	if privateDnsZonesMXRecord_SpecGenerator != nil {
		return privateDnsZonesMXRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec(generators)
	privateDnsZonesMXRecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec(generators)
	privateDnsZonesMXRecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesMXRecord_Spec{}), generators)

	return privateDnsZonesMXRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesMXRecord_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(PrivateDnsZonesMXRecordOperatorSpecGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}
