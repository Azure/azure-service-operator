// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601/storage"
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

func Test_PrivateDnsZonesTXTRecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesTXTRecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesTXTRecord, PrivateDnsZonesTXTRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesTXTRecord tests if a specific instance of PrivateDnsZonesTXTRecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesTXTRecord(subject PrivateDnsZonesTXTRecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200601s.PrivateDnsZonesTXTRecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesTXTRecord
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

func Test_PrivateDnsZonesTXTRecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesTXTRecord to PrivateDnsZonesTXTRecord via AssignProperties_To_PrivateDnsZonesTXTRecord & AssignProperties_From_PrivateDnsZonesTXTRecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesTXTRecord, PrivateDnsZonesTXTRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesTXTRecord tests if a specific instance of PrivateDnsZonesTXTRecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesTXTRecord(subject PrivateDnsZonesTXTRecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZonesTXTRecord
	err := copied.AssignProperties_To_PrivateDnsZonesTXTRecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesTXTRecord
	err = actual.AssignProperties_From_PrivateDnsZonesTXTRecord(&other)
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

func Test_PrivateDnsZonesTXTRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesTXTRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesTXTRecord, PrivateDnsZonesTXTRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesTXTRecord runs a test to see if a specific instance of PrivateDnsZonesTXTRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesTXTRecord(subject PrivateDnsZonesTXTRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesTXTRecord
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

// Generator of PrivateDnsZonesTXTRecord instances for property testing - lazily instantiated by
// PrivateDnsZonesTXTRecordGenerator()
var privateDnsZonesTXTRecordGenerator gopter.Gen

// PrivateDnsZonesTXTRecordGenerator returns a generator of PrivateDnsZonesTXTRecord instances for property testing.
func PrivateDnsZonesTXTRecordGenerator() gopter.Gen {
	if privateDnsZonesTXTRecordGenerator != nil {
		return privateDnsZonesTXTRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesTXTRecord(generators)
	privateDnsZonesTXTRecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesTXTRecord{}), generators)

	return privateDnsZonesTXTRecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesTXTRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesTXTRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_TXT_SpecGenerator()
	gens["Status"] = PrivateDnsZones_TXT_STATUSGenerator()
}

func Test_PrivateDnsZones_TXT_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_TXT_Spec to PrivateDnsZones_TXT_Spec via AssignProperties_To_PrivateDnsZones_TXT_Spec & AssignProperties_From_PrivateDnsZones_TXT_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_TXT_Spec, PrivateDnsZones_TXT_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_TXT_Spec tests if a specific instance of PrivateDnsZones_TXT_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_TXT_Spec(subject PrivateDnsZones_TXT_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZones_TXT_Spec
	err := copied.AssignProperties_To_PrivateDnsZones_TXT_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_TXT_Spec
	err = actual.AssignProperties_From_PrivateDnsZones_TXT_Spec(&other)
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

func Test_PrivateDnsZones_TXT_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_TXT_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_TXT_Spec, PrivateDnsZones_TXT_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_TXT_Spec runs a test to see if a specific instance of PrivateDnsZones_TXT_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_TXT_Spec(subject PrivateDnsZones_TXT_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_TXT_Spec
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

// Generator of PrivateDnsZones_TXT_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_TXT_SpecGenerator()
var privateDnsZones_TXT_SpecGenerator gopter.Gen

// PrivateDnsZones_TXT_SpecGenerator returns a generator of PrivateDnsZones_TXT_Spec instances for property testing.
// We first initialize privateDnsZones_TXT_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_TXT_SpecGenerator() gopter.Gen {
	if privateDnsZones_TXT_SpecGenerator != nil {
		return privateDnsZones_TXT_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_Spec(generators)
	privateDnsZones_TXT_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_TXT_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_Spec(generators)
	privateDnsZones_TXT_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_TXT_Spec{}), generators)

	return privateDnsZones_TXT_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_TXT_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_TXT_STATUS to PrivateDnsZones_TXT_STATUS via AssignProperties_To_PrivateDnsZones_TXT_STATUS & AssignProperties_From_PrivateDnsZones_TXT_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_TXT_STATUS, PrivateDnsZones_TXT_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_TXT_STATUS tests if a specific instance of PrivateDnsZones_TXT_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_TXT_STATUS(subject PrivateDnsZones_TXT_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZones_TXT_STATUS
	err := copied.AssignProperties_To_PrivateDnsZones_TXT_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_TXT_STATUS
	err = actual.AssignProperties_From_PrivateDnsZones_TXT_STATUS(&other)
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

func Test_PrivateDnsZones_TXT_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_TXT_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_TXT_STATUS, PrivateDnsZones_TXT_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_TXT_STATUS runs a test to see if a specific instance of PrivateDnsZones_TXT_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_TXT_STATUS(subject PrivateDnsZones_TXT_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_TXT_STATUS
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

// Generator of PrivateDnsZones_TXT_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_TXT_STATUSGenerator()
var privateDnsZones_TXT_STATUSGenerator gopter.Gen

// PrivateDnsZones_TXT_STATUSGenerator returns a generator of PrivateDnsZones_TXT_STATUS instances for property testing.
// We first initialize privateDnsZones_TXT_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_TXT_STATUSGenerator() gopter.Gen {
	if privateDnsZones_TXT_STATUSGenerator != nil {
		return privateDnsZones_TXT_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_STATUS(generators)
	privateDnsZones_TXT_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_TXT_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_STATUS(generators)
	privateDnsZones_TXT_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_TXT_STATUS{}), generators)

	return privateDnsZones_TXT_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_TXT_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_TXT_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
