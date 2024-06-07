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

func Test_DnsZonesARecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesARecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsZonesARecord, DnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsZonesARecord tests if a specific instance of DnsZonesARecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsZonesARecord(subject DnsZonesARecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DnsZonesARecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsZonesARecord
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

func Test_DnsZonesARecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesARecord to DnsZonesARecord via AssignProperties_To_DnsZonesARecord & AssignProperties_From_DnsZonesARecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesARecord, DnsZonesARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesARecord tests if a specific instance of DnsZonesARecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesARecord(subject DnsZonesARecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZonesARecord
	err := copied.AssignProperties_To_DnsZonesARecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesARecord
	err = actual.AssignProperties_From_DnsZonesARecord(&other)
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
	gens["Spec"] = DnsZones_A_SpecGenerator()
	gens["Status"] = DnsZones_A_STATUSGenerator()
}

func Test_DnsZones_A_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_A_STATUS to DnsZones_A_STATUS via AssignProperties_To_DnsZones_A_STATUS & AssignProperties_From_DnsZones_A_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_A_STATUS, DnsZones_A_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_A_STATUS tests if a specific instance of DnsZones_A_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_A_STATUS(subject DnsZones_A_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZones_A_STATUS
	err := copied.AssignProperties_To_DnsZones_A_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_A_STATUS
	err = actual.AssignProperties_From_DnsZones_A_STATUS(&other)
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

func Test_DnsZones_A_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_A_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_A_STATUS, DnsZones_A_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_A_STATUS runs a test to see if a specific instance of DnsZones_A_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_A_STATUS(subject DnsZones_A_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_A_STATUS
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

// Generator of DnsZones_A_STATUS instances for property testing - lazily instantiated by DnsZones_A_STATUSGenerator()
var dnsZones_A_STATUSGenerator gopter.Gen

// DnsZones_A_STATUSGenerator returns a generator of DnsZones_A_STATUS instances for property testing.
// We first initialize dnsZones_A_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_A_STATUSGenerator() gopter.Gen {
	if dnsZones_A_STATUSGenerator != nil {
		return dnsZones_A_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_STATUS(generators)
	dnsZones_A_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZones_A_STATUS(generators)
	dnsZones_A_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_STATUS{}), generators)

	return dnsZones_A_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_A_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_A_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForDnsZones_A_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_A_STATUS(gens map[string]gopter.Gen) {
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

func Test_DnsZones_A_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_A_Spec to DnsZones_A_Spec via AssignProperties_To_DnsZones_A_Spec & AssignProperties_From_DnsZones_A_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_A_Spec, DnsZones_A_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_A_Spec tests if a specific instance of DnsZones_A_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_A_Spec(subject DnsZones_A_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DnsZones_A_Spec
	err := copied.AssignProperties_To_DnsZones_A_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_A_Spec
	err = actual.AssignProperties_From_DnsZones_A_Spec(&other)
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

func Test_DnsZones_A_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_A_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_A_Spec, DnsZones_A_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_A_Spec runs a test to see if a specific instance of DnsZones_A_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_A_Spec(subject DnsZones_A_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_A_Spec
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

// Generator of DnsZones_A_Spec instances for property testing - lazily instantiated by DnsZones_A_SpecGenerator()
var dnsZones_A_SpecGenerator gopter.Gen

// DnsZones_A_SpecGenerator returns a generator of DnsZones_A_Spec instances for property testing.
// We first initialize dnsZones_A_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_A_SpecGenerator() gopter.Gen {
	if dnsZones_A_SpecGenerator != nil {
		return dnsZones_A_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_Spec(generators)
	dnsZones_A_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZones_A_Spec(generators)
	dnsZones_A_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_Spec{}), generators)

	return dnsZones_A_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_A_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_A_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZones_A_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_A_Spec(gens map[string]gopter.Gen) {
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
