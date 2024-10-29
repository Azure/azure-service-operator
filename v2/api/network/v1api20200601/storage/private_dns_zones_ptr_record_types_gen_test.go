// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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
	gens["Spec"] = PrivateDnsZonesPTRRecord_SpecGenerator()
	gens["Status"] = PrivateDnsZonesPTRRecord_STATUSGenerator()
}

func Test_PrivateDnsZonesPTRRecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesPTRRecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS, PrivateDnsZonesPTRRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS runs a test to see if a specific instance of PrivateDnsZonesPTRRecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS(subject PrivateDnsZonesPTRRecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesPTRRecord_STATUS
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

// Generator of PrivateDnsZonesPTRRecord_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesPTRRecord_STATUSGenerator()
var privateDnsZonesPTRRecord_STATUSGenerator gopter.Gen

// PrivateDnsZonesPTRRecord_STATUSGenerator returns a generator of PrivateDnsZonesPTRRecord_STATUS instances for property testing.
// We first initialize privateDnsZonesPTRRecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesPTRRecord_STATUSGenerator() gopter.Gen {
	if privateDnsZonesPTRRecord_STATUSGenerator != nil {
		return privateDnsZonesPTRRecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	privateDnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	privateDnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_STATUS{}), generators)

	return privateDnsZonesPTRRecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}

func Test_PrivateDnsZonesPTRRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesPTRRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesPTRRecord_Spec, PrivateDnsZonesPTRRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesPTRRecord_Spec runs a test to see if a specific instance of PrivateDnsZonesPTRRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesPTRRecord_Spec(subject PrivateDnsZonesPTRRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesPTRRecord_Spec
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

// Generator of PrivateDnsZonesPTRRecord_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesPTRRecord_SpecGenerator()
var privateDnsZonesPTRRecord_SpecGenerator gopter.Gen

// PrivateDnsZonesPTRRecord_SpecGenerator returns a generator of PrivateDnsZonesPTRRecord_Spec instances for property testing.
// We first initialize privateDnsZonesPTRRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesPTRRecord_SpecGenerator() gopter.Gen {
	if privateDnsZonesPTRRecord_SpecGenerator != nil {
		return privateDnsZonesPTRRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec(generators)
	privateDnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec(generators)
	privateDnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_Spec{}), generators)

	return privateDnsZonesPTRRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}
