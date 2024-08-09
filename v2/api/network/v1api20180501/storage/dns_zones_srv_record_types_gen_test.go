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

func Test_DnsZonesSRVRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecord, DnsZonesSRVRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecord runs a test to see if a specific instance of DnsZonesSRVRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecord(subject DnsZonesSRVRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecord
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

// Generator of DnsZonesSRVRecord instances for property testing - lazily instantiated by DnsZonesSRVRecordGenerator()
var dnsZonesSRVRecordGenerator gopter.Gen

// DnsZonesSRVRecordGenerator returns a generator of DnsZonesSRVRecord instances for property testing.
func DnsZonesSRVRecordGenerator() gopter.Gen {
	if dnsZonesSRVRecordGenerator != nil {
		return dnsZonesSRVRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesSRVRecord(generators)
	dnsZonesSRVRecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord{}), generators)

	return dnsZonesSRVRecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesSRVRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesSRVRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZonesSRVRecord_SpecGenerator()
	gens["Status"] = DnsZonesSRVRecord_STATUSGenerator()
}

func Test_DnsZonesSRVRecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecordOperatorSpec, DnsZonesSRVRecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecordOperatorSpec runs a test to see if a specific instance of DnsZonesSRVRecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecordOperatorSpec(subject DnsZonesSRVRecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecordOperatorSpec
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

// Generator of DnsZonesSRVRecordOperatorSpec instances for property testing - lazily instantiated by
// DnsZonesSRVRecordOperatorSpecGenerator()
var dnsZonesSRVRecordOperatorSpecGenerator gopter.Gen

// DnsZonesSRVRecordOperatorSpecGenerator returns a generator of DnsZonesSRVRecordOperatorSpec instances for property testing.
func DnsZonesSRVRecordOperatorSpecGenerator() gopter.Gen {
	if dnsZonesSRVRecordOperatorSpecGenerator != nil {
		return dnsZonesSRVRecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	dnsZonesSRVRecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecordOperatorSpec{}), generators)

	return dnsZonesSRVRecordOperatorSpecGenerator
}

func Test_DnsZonesSRVRecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecord_STATUS, DnsZonesSRVRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecord_STATUS runs a test to see if a specific instance of DnsZonesSRVRecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecord_STATUS(subject DnsZonesSRVRecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecord_STATUS
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

// Generator of DnsZonesSRVRecord_STATUS instances for property testing - lazily instantiated by
// DnsZonesSRVRecord_STATUSGenerator()
var dnsZonesSRVRecord_STATUSGenerator gopter.Gen

// DnsZonesSRVRecord_STATUSGenerator returns a generator of DnsZonesSRVRecord_STATUS instances for property testing.
// We first initialize dnsZonesSRVRecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesSRVRecord_STATUSGenerator() gopter.Gen {
	if dnsZonesSRVRecord_STATUSGenerator != nil {
		return dnsZonesSRVRecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS(generators)
	dnsZonesSRVRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS(generators)
	dnsZonesSRVRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_STATUS{}), generators)

	return dnsZonesSRVRecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecord_STATUSGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecord_STATUSGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_DnsZonesSRVRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecord_Spec, DnsZonesSRVRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecord_Spec runs a test to see if a specific instance of DnsZonesSRVRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecord_Spec(subject DnsZonesSRVRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecord_Spec
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

// Generator of DnsZonesSRVRecord_Spec instances for property testing - lazily instantiated by
// DnsZonesSRVRecord_SpecGenerator()
var dnsZonesSRVRecord_SpecGenerator gopter.Gen

// DnsZonesSRVRecord_SpecGenerator returns a generator of DnsZonesSRVRecord_Spec instances for property testing.
// We first initialize dnsZonesSRVRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesSRVRecord_SpecGenerator() gopter.Gen {
	if dnsZonesSRVRecord_SpecGenerator != nil {
		return dnsZonesSRVRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_Spec(generators)
	dnsZonesSRVRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_Spec(generators)
	dnsZonesSRVRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_Spec{}), generators)

	return dnsZonesSRVRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(DnsZonesSRVRecordOperatorSpecGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}
