// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601/storage"
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

func Test_PrivateDnsZonesARecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesARecord tests if a specific instance of PrivateDnsZonesARecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.PrivateDnsZonesARecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesARecord
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

func Test_PrivateDnsZonesARecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord to PrivateDnsZonesARecord via AssignProperties_To_PrivateDnsZonesARecord & AssignProperties_From_PrivateDnsZonesARecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesARecord tests if a specific instance of PrivateDnsZonesARecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesARecord
	err := copied.AssignProperties_To_PrivateDnsZonesARecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesARecord
	err = actual.AssignProperties_From_PrivateDnsZonesARecord(&other)
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

func Test_PrivateDnsZonesARecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesARecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesARecord, PrivateDnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesARecord runs a test to see if a specific instance of PrivateDnsZonesARecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesARecord(subject PrivateDnsZonesARecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesARecord
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

// Generator of PrivateDnsZonesARecord instances for property testing - lazily instantiated by
// PrivateDnsZonesARecordGenerator()
var privateDnsZonesARecordGenerator gopter.Gen

// PrivateDnsZonesARecordGenerator returns a generator of PrivateDnsZonesARecord instances for property testing.
func PrivateDnsZonesARecordGenerator() gopter.Gen {
	if privateDnsZonesARecordGenerator != nil {
		return privateDnsZonesARecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord(generators)
	privateDnsZonesARecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord{}), generators)

	return privateDnsZonesARecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZonesARecord_SpecGenerator()
	gens["Status"] = PrivateDnsZonesARecord_STATUSGenerator()
}

func Test_PrivateDnsZonesARecordOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecordOperatorSpec to PrivateDnsZonesARecordOperatorSpec via AssignProperties_To_PrivateDnsZonesARecordOperatorSpec & AssignProperties_From_PrivateDnsZonesARecordOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesARecordOperatorSpec, PrivateDnsZonesARecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesARecordOperatorSpec tests if a specific instance of PrivateDnsZonesARecordOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesARecordOperatorSpec(subject PrivateDnsZonesARecordOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesARecordOperatorSpec
	err := copied.AssignProperties_To_PrivateDnsZonesARecordOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesARecordOperatorSpec
	err = actual.AssignProperties_From_PrivateDnsZonesARecordOperatorSpec(&other)
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

func Test_PrivateDnsZonesARecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesARecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesARecordOperatorSpec, PrivateDnsZonesARecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesARecordOperatorSpec runs a test to see if a specific instance of PrivateDnsZonesARecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesARecordOperatorSpec(subject PrivateDnsZonesARecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesARecordOperatorSpec
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

// Generator of PrivateDnsZonesARecordOperatorSpec instances for property testing - lazily instantiated by
// PrivateDnsZonesARecordOperatorSpecGenerator()
var privateDnsZonesARecordOperatorSpecGenerator gopter.Gen

// PrivateDnsZonesARecordOperatorSpecGenerator returns a generator of PrivateDnsZonesARecordOperatorSpec instances for property testing.
func PrivateDnsZonesARecordOperatorSpecGenerator() gopter.Gen {
	if privateDnsZonesARecordOperatorSpecGenerator != nil {
		return privateDnsZonesARecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	privateDnsZonesARecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecordOperatorSpec{}), generators)

	return privateDnsZonesARecordOperatorSpecGenerator
}

func Test_PrivateDnsZonesARecord_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord_STATUS to PrivateDnsZonesARecord_STATUS via AssignProperties_To_PrivateDnsZonesARecord_STATUS & AssignProperties_From_PrivateDnsZonesARecord_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesARecord_STATUS, PrivateDnsZonesARecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesARecord_STATUS tests if a specific instance of PrivateDnsZonesARecord_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesARecord_STATUS(subject PrivateDnsZonesARecord_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesARecord_STATUS
	err := copied.AssignProperties_To_PrivateDnsZonesARecord_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesARecord_STATUS
	err = actual.AssignProperties_From_PrivateDnsZonesARecord_STATUS(&other)
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

func Test_PrivateDnsZonesARecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesARecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesARecord_STATUS, PrivateDnsZonesARecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesARecord_STATUS runs a test to see if a specific instance of PrivateDnsZonesARecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesARecord_STATUS(subject PrivateDnsZonesARecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesARecord_STATUS
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

// Generator of PrivateDnsZonesARecord_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesARecord_STATUSGenerator()
var privateDnsZonesARecord_STATUSGenerator gopter.Gen

// PrivateDnsZonesARecord_STATUSGenerator returns a generator of PrivateDnsZonesARecord_STATUS instances for property testing.
// We first initialize privateDnsZonesARecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesARecord_STATUSGenerator() gopter.Gen {
	if privateDnsZonesARecord_STATUSGenerator != nil {
		return privateDnsZonesARecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_STATUS(generators)
	privateDnsZonesARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_STATUS(generators)
	privateDnsZonesARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord_STATUS{}), generators)

	return privateDnsZonesARecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}

func Test_PrivateDnsZonesARecord_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesARecord_Spec to PrivateDnsZonesARecord_Spec via AssignProperties_To_PrivateDnsZonesARecord_Spec & AssignProperties_From_PrivateDnsZonesARecord_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesARecord_Spec, PrivateDnsZonesARecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesARecord_Spec tests if a specific instance of PrivateDnsZonesARecord_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesARecord_Spec(subject PrivateDnsZonesARecord_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateDnsZonesARecord_Spec
	err := copied.AssignProperties_To_PrivateDnsZonesARecord_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesARecord_Spec
	err = actual.AssignProperties_From_PrivateDnsZonesARecord_Spec(&other)
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

func Test_PrivateDnsZonesARecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesARecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesARecord_Spec, PrivateDnsZonesARecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesARecord_Spec runs a test to see if a specific instance of PrivateDnsZonesARecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesARecord_Spec(subject PrivateDnsZonesARecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesARecord_Spec
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

// Generator of PrivateDnsZonesARecord_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesARecord_SpecGenerator()
var privateDnsZonesARecord_SpecGenerator gopter.Gen

// PrivateDnsZonesARecord_SpecGenerator returns a generator of PrivateDnsZonesARecord_Spec instances for property testing.
// We first initialize privateDnsZonesARecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesARecord_SpecGenerator() gopter.Gen {
	if privateDnsZonesARecord_SpecGenerator != nil {
		return privateDnsZonesARecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_Spec(generators)
	privateDnsZonesARecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_Spec(generators)
	privateDnsZonesARecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesARecord_Spec{}), generators)

	return privateDnsZonesARecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesARecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesARecord_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(PrivateDnsZonesARecordOperatorSpecGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}
