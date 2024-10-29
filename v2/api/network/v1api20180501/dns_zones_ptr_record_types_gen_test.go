// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501/storage"
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

func Test_DnsZonesPTRRecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesPTRRecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsZonesPTRRecord, DnsZonesPTRRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsZonesPTRRecord tests if a specific instance of DnsZonesPTRRecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsZonesPTRRecord(subject DnsZonesPTRRecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DnsZonesPTRRecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsZonesPTRRecord
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

func Test_DnsZonesPTRRecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesPTRRecord to DnsZonesPTRRecord via AssignProperties_To_DnsZonesPTRRecord & AssignProperties_From_DnsZonesPTRRecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesPTRRecord, DnsZonesPTRRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesPTRRecord tests if a specific instance of DnsZonesPTRRecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesPTRRecord(subject DnsZonesPTRRecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZonesPTRRecord
	err := copied.AssignProperties_To_DnsZonesPTRRecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesPTRRecord
	err = actual.AssignProperties_From_DnsZonesPTRRecord(&other)
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

func Test_DnsZonesPTRRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesPTRRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesPTRRecord, DnsZonesPTRRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesPTRRecord runs a test to see if a specific instance of DnsZonesPTRRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesPTRRecord(subject DnsZonesPTRRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesPTRRecord
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

// Generator of DnsZonesPTRRecord instances for property testing - lazily instantiated by DnsZonesPTRRecordGenerator()
var dnsZonesPTRRecordGenerator gopter.Gen

// DnsZonesPTRRecordGenerator returns a generator of DnsZonesPTRRecord instances for property testing.
func DnsZonesPTRRecordGenerator() gopter.Gen {
	if dnsZonesPTRRecordGenerator != nil {
		return dnsZonesPTRRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesPTRRecord(generators)
	dnsZonesPTRRecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord{}), generators)

	return dnsZonesPTRRecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesPTRRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesPTRRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZonesPTRRecord_SpecGenerator()
	gens["Status"] = DnsZonesPTRRecord_STATUSGenerator()
}

func Test_DnsZonesPTRRecord_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesPTRRecord_STATUS to DnsZonesPTRRecord_STATUS via AssignProperties_To_DnsZonesPTRRecord_STATUS & AssignProperties_From_DnsZonesPTRRecord_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesPTRRecord_STATUS, DnsZonesPTRRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesPTRRecord_STATUS tests if a specific instance of DnsZonesPTRRecord_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesPTRRecord_STATUS(subject DnsZonesPTRRecord_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZonesPTRRecord_STATUS
	err := copied.AssignProperties_To_DnsZonesPTRRecord_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesPTRRecord_STATUS
	err = actual.AssignProperties_From_DnsZonesPTRRecord_STATUS(&other)
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

func Test_DnsZonesPTRRecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesPTRRecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesPTRRecord_STATUS, DnsZonesPTRRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesPTRRecord_STATUS runs a test to see if a specific instance of DnsZonesPTRRecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesPTRRecord_STATUS(subject DnsZonesPTRRecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesPTRRecord_STATUS
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

// Generator of DnsZonesPTRRecord_STATUS instances for property testing - lazily instantiated by
// DnsZonesPTRRecord_STATUSGenerator()
var dnsZonesPTRRecord_STATUSGenerator gopter.Gen

// DnsZonesPTRRecord_STATUSGenerator returns a generator of DnsZonesPTRRecord_STATUS instances for property testing.
// We first initialize dnsZonesPTRRecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesPTRRecord_STATUSGenerator() gopter.Gen {
	if dnsZonesPTRRecord_STATUSGenerator != nil {
		return dnsZonesPTRRecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_STATUS(generators)
	dnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_STATUS(generators)
	dnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_STATUS{}), generators)

	return dnsZonesPTRRecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
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

func Test_DnsZonesPTRRecord_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesPTRRecord_Spec to DnsZonesPTRRecord_Spec via AssignProperties_To_DnsZonesPTRRecord_Spec & AssignProperties_From_DnsZonesPTRRecord_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesPTRRecord_Spec, DnsZonesPTRRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesPTRRecord_Spec tests if a specific instance of DnsZonesPTRRecord_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesPTRRecord_Spec(subject DnsZonesPTRRecord_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZonesPTRRecord_Spec
	err := copied.AssignProperties_To_DnsZonesPTRRecord_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesPTRRecord_Spec
	err = actual.AssignProperties_From_DnsZonesPTRRecord_Spec(&other)
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

func Test_DnsZonesPTRRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesPTRRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesPTRRecord_Spec, DnsZonesPTRRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesPTRRecord_Spec runs a test to see if a specific instance of DnsZonesPTRRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesPTRRecord_Spec(subject DnsZonesPTRRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesPTRRecord_Spec
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

// Generator of DnsZonesPTRRecord_Spec instances for property testing - lazily instantiated by
// DnsZonesPTRRecord_SpecGenerator()
var dnsZonesPTRRecord_SpecGenerator gopter.Gen

// DnsZonesPTRRecord_SpecGenerator returns a generator of DnsZonesPTRRecord_Spec instances for property testing.
// We first initialize dnsZonesPTRRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesPTRRecord_SpecGenerator() gopter.Gen {
	if dnsZonesPTRRecord_SpecGenerator != nil {
		return dnsZonesPTRRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	dnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	dnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_Spec{}), generators)

	return dnsZonesPTRRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}
