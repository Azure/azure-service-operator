// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601storage"
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

func Test_PrivateDnsZonesSRVRecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesSRVRecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesSRVRecord, PrivateDnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesSRVRecord tests if a specific instance of PrivateDnsZonesSRVRecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesSRVRecord(subject PrivateDnsZonesSRVRecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20200601s.PrivateDnsZonesSRVRecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesSRVRecord
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

func Test_PrivateDnsZonesSRVRecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesSRVRecord to PrivateDnsZonesSRVRecord via AssignProperties_To_PrivateDnsZonesSRVRecord & AssignProperties_From_PrivateDnsZonesSRVRecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesSRVRecord, PrivateDnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesSRVRecord tests if a specific instance of PrivateDnsZonesSRVRecord can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesSRVRecord(subject PrivateDnsZonesSRVRecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZonesSRVRecord
	err := copied.AssignProperties_To_PrivateDnsZonesSRVRecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesSRVRecord
	err = actual.AssignProperties_From_PrivateDnsZonesSRVRecord(&other)
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

func Test_PrivateDnsZonesSRVRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesSRVRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesSRVRecord, PrivateDnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesSRVRecord runs a test to see if a specific instance of PrivateDnsZonesSRVRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesSRVRecord(subject PrivateDnsZonesSRVRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesSRVRecord
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

// Generator of PrivateDnsZonesSRVRecord instances for property testing - lazily instantiated by
// PrivateDnsZonesSRVRecordGenerator()
var privateDnsZonesSRVRecordGenerator gopter.Gen

// PrivateDnsZonesSRVRecordGenerator returns a generator of PrivateDnsZonesSRVRecord instances for property testing.
func PrivateDnsZonesSRVRecordGenerator() gopter.Gen {
	if privateDnsZonesSRVRecordGenerator != nil {
		return privateDnsZonesSRVRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesSRVRecord(generators)
	privateDnsZonesSRVRecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesSRVRecord{}), generators)

	return privateDnsZonesSRVRecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesSRVRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesSRVRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_SRV_SpecGenerator()
	gens["Status"] = PrivateDnsZones_SRV_STATUSGenerator()
}

func Test_PrivateDnsZones_SRV_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_SRV_Spec to PrivateDnsZones_SRV_Spec via AssignProperties_To_PrivateDnsZones_SRV_Spec & AssignProperties_From_PrivateDnsZones_SRV_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_SRV_Spec, PrivateDnsZones_SRV_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_SRV_Spec tests if a specific instance of PrivateDnsZones_SRV_Spec can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_SRV_Spec(subject PrivateDnsZones_SRV_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZones_SRV_Spec
	err := copied.AssignProperties_To_PrivateDnsZones_SRV_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_SRV_Spec
	err = actual.AssignProperties_From_PrivateDnsZones_SRV_Spec(&other)
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

func Test_PrivateDnsZones_SRV_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_SRV_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_SRV_Spec, PrivateDnsZones_SRV_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_SRV_Spec runs a test to see if a specific instance of PrivateDnsZones_SRV_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_SRV_Spec(subject PrivateDnsZones_SRV_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_SRV_Spec
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

// Generator of PrivateDnsZones_SRV_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_SRV_SpecGenerator()
var privateDnsZones_SRV_SpecGenerator gopter.Gen

// PrivateDnsZones_SRV_SpecGenerator returns a generator of PrivateDnsZones_SRV_Spec instances for property testing.
// We first initialize privateDnsZones_SRV_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_SRV_SpecGenerator() gopter.Gen {
	if privateDnsZones_SRV_SpecGenerator != nil {
		return privateDnsZones_SRV_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_Spec(generators)
	privateDnsZones_SRV_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_SRV_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_Spec(generators)
	privateDnsZones_SRV_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_SRV_Spec{}), generators)

	return privateDnsZones_SRV_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_SRV_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_SRV_STATUS to PrivateDnsZones_SRV_STATUS via AssignProperties_To_PrivateDnsZones_SRV_STATUS & AssignProperties_From_PrivateDnsZones_SRV_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_SRV_STATUS, PrivateDnsZones_SRV_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_SRV_STATUS tests if a specific instance of PrivateDnsZones_SRV_STATUS can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_SRV_STATUS(subject PrivateDnsZones_SRV_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZones_SRV_STATUS
	err := copied.AssignProperties_To_PrivateDnsZones_SRV_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_SRV_STATUS
	err = actual.AssignProperties_From_PrivateDnsZones_SRV_STATUS(&other)
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

func Test_PrivateDnsZones_SRV_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_SRV_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_SRV_STATUS, PrivateDnsZones_SRV_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_SRV_STATUS runs a test to see if a specific instance of PrivateDnsZones_SRV_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_SRV_STATUS(subject PrivateDnsZones_SRV_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_SRV_STATUS
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

// Generator of PrivateDnsZones_SRV_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_SRV_STATUSGenerator()
var privateDnsZones_SRV_STATUSGenerator gopter.Gen

// PrivateDnsZones_SRV_STATUSGenerator returns a generator of PrivateDnsZones_SRV_STATUS instances for property testing.
// We first initialize privateDnsZones_SRV_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_SRV_STATUSGenerator() gopter.Gen {
	if privateDnsZones_SRV_STATUSGenerator != nil {
		return privateDnsZones_SRV_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_STATUS(generators)
	privateDnsZones_SRV_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_SRV_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_STATUS(generators)
	privateDnsZones_SRV_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_SRV_STATUS{}), generators)

	return privateDnsZones_SRV_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_SRV_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_SRV_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
