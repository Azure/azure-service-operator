// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

import (
	"encoding/json"
	v20180501s "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501/storage"
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

func Test_DnsZonesCNAMERecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesCNAMERecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForDnsZonesCNAMERecord, DnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDnsZonesCNAMERecord tests if a specific instance of DnsZonesCNAMERecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDnsZonesCNAMERecord(subject DnsZonesCNAMERecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20180501s.DnsZonesCNAMERecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DnsZonesCNAMERecord
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

func Test_DnsZonesCNAMERecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZonesCNAMERecord to DnsZonesCNAMERecord via AssignProperties_To_DnsZonesCNAMERecord & AssignProperties_From_DnsZonesCNAMERecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZonesCNAMERecord, DnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZonesCNAMERecord tests if a specific instance of DnsZonesCNAMERecord can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZonesCNAMERecord(subject DnsZonesCNAMERecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180501s.DnsZonesCNAMERecord
	err := copied.AssignProperties_To_DnsZonesCNAMERecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZonesCNAMERecord
	err = actual.AssignProperties_From_DnsZonesCNAMERecord(&other)
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

func Test_DnsZonesCNAMERecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCNAMERecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCNAMERecord, DnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCNAMERecord runs a test to see if a specific instance of DnsZonesCNAMERecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCNAMERecord(subject DnsZonesCNAMERecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCNAMERecord
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

// Generator of DnsZonesCNAMERecord instances for property testing - lazily instantiated by
// DnsZonesCNAMERecordGenerator()
var dnsZonesCNAMERecordGenerator gopter.Gen

// DnsZonesCNAMERecordGenerator returns a generator of DnsZonesCNAMERecord instances for property testing.
func DnsZonesCNAMERecordGenerator() gopter.Gen {
	if dnsZonesCNAMERecordGenerator != nil {
		return dnsZonesCNAMERecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesCNAMERecord(generators)
	dnsZonesCNAMERecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesCNAMERecord{}), generators)

	return dnsZonesCNAMERecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesCNAMERecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesCNAMERecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZones_CNAME_SpecGenerator()
	gens["Status"] = DnsZones_CNAME_STATUSGenerator()
}

func Test_DnsZones_CNAME_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_CNAME_Spec to DnsZones_CNAME_Spec via AssignProperties_To_DnsZones_CNAME_Spec & AssignProperties_From_DnsZones_CNAME_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_CNAME_Spec, DnsZones_CNAME_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_CNAME_Spec tests if a specific instance of DnsZones_CNAME_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_CNAME_Spec(subject DnsZones_CNAME_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180501s.DnsZones_CNAME_Spec
	err := copied.AssignProperties_To_DnsZones_CNAME_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_CNAME_Spec
	err = actual.AssignProperties_From_DnsZones_CNAME_Spec(&other)
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

func Test_DnsZones_CNAME_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_CNAME_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_CNAME_Spec, DnsZones_CNAME_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_CNAME_Spec runs a test to see if a specific instance of DnsZones_CNAME_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_CNAME_Spec(subject DnsZones_CNAME_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_CNAME_Spec
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

// Generator of DnsZones_CNAME_Spec instances for property testing - lazily instantiated by
// DnsZones_CNAME_SpecGenerator()
var dnsZones_CNAME_SpecGenerator gopter.Gen

// DnsZones_CNAME_SpecGenerator returns a generator of DnsZones_CNAME_Spec instances for property testing.
// We first initialize dnsZones_CNAME_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_CNAME_SpecGenerator() gopter.Gen {
	if dnsZones_CNAME_SpecGenerator != nil {
		return dnsZones_CNAME_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_CNAME_Spec(generators)
	dnsZones_CNAME_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_CNAME_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_CNAME_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZones_CNAME_Spec(generators)
	dnsZones_CNAME_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_CNAME_Spec{}), generators)

	return dnsZones_CNAME_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_CNAME_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_CNAME_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZones_CNAME_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_CNAME_Spec(gens map[string]gopter.Gen) {
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

func Test_DnsZones_CNAME_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DnsZones_CNAME_STATUS to DnsZones_CNAME_STATUS via AssignProperties_To_DnsZones_CNAME_STATUS & AssignProperties_From_DnsZones_CNAME_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDnsZones_CNAME_STATUS, DnsZones_CNAME_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDnsZones_CNAME_STATUS tests if a specific instance of DnsZones_CNAME_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDnsZones_CNAME_STATUS(subject DnsZones_CNAME_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180501s.DnsZones_CNAME_STATUS
	err := copied.AssignProperties_To_DnsZones_CNAME_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DnsZones_CNAME_STATUS
	err = actual.AssignProperties_From_DnsZones_CNAME_STATUS(&other)
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

func Test_DnsZones_CNAME_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_CNAME_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_CNAME_STATUS, DnsZones_CNAME_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_CNAME_STATUS runs a test to see if a specific instance of DnsZones_CNAME_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_CNAME_STATUS(subject DnsZones_CNAME_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_CNAME_STATUS
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

// Generator of DnsZones_CNAME_STATUS instances for property testing - lazily instantiated by
// DnsZones_CNAME_STATUSGenerator()
var dnsZones_CNAME_STATUSGenerator gopter.Gen

// DnsZones_CNAME_STATUSGenerator returns a generator of DnsZones_CNAME_STATUS instances for property testing.
// We first initialize dnsZones_CNAME_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_CNAME_STATUSGenerator() gopter.Gen {
	if dnsZones_CNAME_STATUSGenerator != nil {
		return dnsZones_CNAME_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_CNAME_STATUS(generators)
	dnsZones_CNAME_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_CNAME_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_CNAME_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZones_CNAME_STATUS(generators)
	dnsZones_CNAME_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_CNAME_STATUS{}), generators)

	return dnsZones_CNAME_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_CNAME_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_CNAME_STATUS(gens map[string]gopter.Gen) {
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

// AddRelatedPropertyGeneratorsForDnsZones_CNAME_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_CNAME_STATUS(gens map[string]gopter.Gen) {
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
