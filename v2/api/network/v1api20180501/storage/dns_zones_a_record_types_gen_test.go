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

func Test_DnsZonesARecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesARecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesARecord, DnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesARecord runs a test to see if a specific instance of DnsZonesARecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesARecord(subject DnsZonesARecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesARecord
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

// Generator of DnsZonesARecord instances for property testing - lazily instantiated by DnsZonesARecordGenerator()
var dnsZonesARecordGenerator gopter.Gen

// DnsZonesARecordGenerator returns a generator of DnsZonesARecord instances for property testing.
func DnsZonesARecordGenerator() gopter.Gen {
	if dnsZonesARecordGenerator != nil {
		return dnsZonesARecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesARecord(generators)
	dnsZonesARecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecord{}), generators)

	return dnsZonesARecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesARecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesARecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZonesARecord_SpecGenerator()
	gens["Status"] = DnsZonesARecord_STATUSGenerator()
}

func Test_DnsZonesARecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesARecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesARecordOperatorSpec, DnsZonesARecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesARecordOperatorSpec runs a test to see if a specific instance of DnsZonesARecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesARecordOperatorSpec(subject DnsZonesARecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesARecordOperatorSpec
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

// Generator of DnsZonesARecordOperatorSpec instances for property testing - lazily instantiated by
// DnsZonesARecordOperatorSpecGenerator()
var dnsZonesARecordOperatorSpecGenerator gopter.Gen

// DnsZonesARecordOperatorSpecGenerator returns a generator of DnsZonesARecordOperatorSpec instances for property testing.
func DnsZonesARecordOperatorSpecGenerator() gopter.Gen {
	if dnsZonesARecordOperatorSpecGenerator != nil {
		return dnsZonesARecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	dnsZonesARecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecordOperatorSpec{}), generators)

	return dnsZonesARecordOperatorSpecGenerator
}

func Test_DnsZonesARecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesARecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesARecord_STATUS, DnsZonesARecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesARecord_STATUS runs a test to see if a specific instance of DnsZonesARecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesARecord_STATUS(subject DnsZonesARecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesARecord_STATUS
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

// Generator of DnsZonesARecord_STATUS instances for property testing - lazily instantiated by
// DnsZonesARecord_STATUSGenerator()
var dnsZonesARecord_STATUSGenerator gopter.Gen

// DnsZonesARecord_STATUSGenerator returns a generator of DnsZonesARecord_STATUS instances for property testing.
// We first initialize dnsZonesARecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesARecord_STATUSGenerator() gopter.Gen {
	if dnsZonesARecord_STATUSGenerator != nil {
		return dnsZonesARecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesARecord_STATUS(generators)
	dnsZonesARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesARecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZonesARecord_STATUS(generators)
	dnsZonesARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecord_STATUS{}), generators)

	return dnsZonesARecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesARecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesARecord_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForDnsZonesARecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesARecord_STATUS(gens map[string]gopter.Gen) {
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

func Test_DnsZonesARecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesARecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesARecord_Spec, DnsZonesARecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesARecord_Spec runs a test to see if a specific instance of DnsZonesARecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesARecord_Spec(subject DnsZonesARecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesARecord_Spec
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

// Generator of DnsZonesARecord_Spec instances for property testing - lazily instantiated by
// DnsZonesARecord_SpecGenerator()
var dnsZonesARecord_SpecGenerator gopter.Gen

// DnsZonesARecord_SpecGenerator returns a generator of DnsZonesARecord_Spec instances for property testing.
// We first initialize dnsZonesARecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesARecord_SpecGenerator() gopter.Gen {
	if dnsZonesARecord_SpecGenerator != nil {
		return dnsZonesARecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesARecord_Spec(generators)
	dnsZonesARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesARecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesARecord_Spec(generators)
	dnsZonesARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesARecord_Spec{}), generators)

	return dnsZonesARecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesARecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesARecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZonesARecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesARecord_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(DnsZonesARecordOperatorSpecGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}
