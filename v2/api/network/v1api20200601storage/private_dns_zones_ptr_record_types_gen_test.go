// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601storage

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

func Test_PrivateDnsZonesPTRRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesPTRRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesPTRRecord, PrivateDnsZonesPTRRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesPTRRecord runs a test to see if a specific instance of PrivateDnsZonesPTRRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesPTRRecord(subject PrivateDnsZonesPTRRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesPTRRecord
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

// Generator of PrivateDnsZonesPTRRecord instances for property testing - lazily instantiated by
// PrivateDnsZonesPTRRecordGenerator()
var privateDnsZonesPTRRecordGenerator gopter.Gen

// PrivateDnsZonesPTRRecordGenerator returns a generator of PrivateDnsZonesPTRRecord instances for property testing.
func PrivateDnsZonesPTRRecordGenerator() gopter.Gen {
	if privateDnsZonesPTRRecordGenerator != nil {
		return privateDnsZonesPTRRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord(generators)
	privateDnsZonesPTRRecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord{}), generators)

	return privateDnsZonesPTRRecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_PTR_SpecGenerator()
	gens["Status"] = PrivateDnsZones_PTR_STATUSGenerator()
}

func Test_PrivateDnsZones_PTR_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_PTR_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_PTR_Spec, PrivateDnsZones_PTR_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_PTR_Spec runs a test to see if a specific instance of PrivateDnsZones_PTR_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_PTR_Spec(subject PrivateDnsZones_PTR_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_PTR_Spec
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

// Generator of PrivateDnsZones_PTR_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_PTR_SpecGenerator()
var privateDnsZones_PTR_SpecGenerator gopter.Gen

// PrivateDnsZones_PTR_SpecGenerator returns a generator of PrivateDnsZones_PTR_Spec instances for property testing.
// We first initialize privateDnsZones_PTR_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_PTR_SpecGenerator() gopter.Gen {
	if privateDnsZones_PTR_SpecGenerator != nil {
		return privateDnsZones_PTR_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_Spec(generators)
	privateDnsZones_PTR_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_PTR_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_Spec(generators)
	privateDnsZones_PTR_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_PTR_Spec{}), generators)

	return privateDnsZones_PTR_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_PTR_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_PTR_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_PTR_STATUS, PrivateDnsZones_PTR_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_PTR_STATUS runs a test to see if a specific instance of PrivateDnsZones_PTR_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_PTR_STATUS(subject PrivateDnsZones_PTR_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_PTR_STATUS
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

// Generator of PrivateDnsZones_PTR_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_PTR_STATUSGenerator()
var privateDnsZones_PTR_STATUSGenerator gopter.Gen

// PrivateDnsZones_PTR_STATUSGenerator returns a generator of PrivateDnsZones_PTR_STATUS instances for property testing.
// We first initialize privateDnsZones_PTR_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_PTR_STATUSGenerator() gopter.Gen {
	if privateDnsZones_PTR_STATUSGenerator != nil {
		return privateDnsZones_PTR_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_STATUS(generators)
	privateDnsZones_PTR_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_PTR_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_STATUS(generators)
	privateDnsZones_PTR_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_PTR_STATUS{}), generators)

	return privateDnsZones_PTR_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_PTR_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_PTR_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
