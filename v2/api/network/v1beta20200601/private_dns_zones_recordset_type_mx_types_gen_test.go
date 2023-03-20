// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import (
	"encoding/json"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20200601storage"
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

func Test_PrivateDnsZonesRecordsetTypeMX_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesRecordsetTypeMX to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesRecordsetTypeMX, PrivateDnsZonesRecordsetTypeMXGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesRecordsetTypeMX tests if a specific instance of PrivateDnsZonesRecordsetTypeMX round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesRecordsetTypeMX(subject PrivateDnsZonesRecordsetTypeMX) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200601s.PrivateDnsZonesRecordsetTypeMX
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesRecordsetTypeMX
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

func Test_PrivateDnsZonesRecordsetTypeMX_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesRecordsetTypeMX to PrivateDnsZonesRecordsetTypeMX via AssignProperties_To_PrivateDnsZonesRecordsetTypeMX & AssignProperties_From_PrivateDnsZonesRecordsetTypeMX returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesRecordsetTypeMX, PrivateDnsZonesRecordsetTypeMXGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesRecordsetTypeMX tests if a specific instance of PrivateDnsZonesRecordsetTypeMX can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesRecordsetTypeMX(subject PrivateDnsZonesRecordsetTypeMX) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZonesRecordsetTypeMX
	err := copied.AssignProperties_To_PrivateDnsZonesRecordsetTypeMX(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesRecordsetTypeMX
	err = actual.AssignProperties_From_PrivateDnsZonesRecordsetTypeMX(&other)
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

func Test_PrivateDnsZonesRecordsetTypeMX_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesRecordsetTypeMX via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesRecordsetTypeMX, PrivateDnsZonesRecordsetTypeMXGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesRecordsetTypeMX runs a test to see if a specific instance of PrivateDnsZonesRecordsetTypeMX round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesRecordsetTypeMX(subject PrivateDnsZonesRecordsetTypeMX) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesRecordsetTypeMX
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

// Generator of PrivateDnsZonesRecordsetTypeMX instances for property testing - lazily instantiated by
// PrivateDnsZonesRecordsetTypeMXGenerator()
var privateDnsZonesRecordsetTypeMXGenerator gopter.Gen

// PrivateDnsZonesRecordsetTypeMXGenerator returns a generator of PrivateDnsZonesRecordsetTypeMX instances for property testing.
func PrivateDnsZonesRecordsetTypeMXGenerator() gopter.Gen {
	if privateDnsZonesRecordsetTypeMXGenerator != nil {
		return privateDnsZonesRecordsetTypeMXGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesRecordsetTypeMX(generators)
	privateDnsZonesRecordsetTypeMXGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesRecordsetTypeMX{}), generators)

	return privateDnsZonesRecordsetTypeMXGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesRecordsetTypeMX is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesRecordsetTypeMX(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_MX_SpecGenerator()
	gens["Status"] = PrivateDnsZones_MX_STATUSGenerator()
}

func Test_PrivateDnsZones_MX_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_MX_Spec to PrivateDnsZones_MX_Spec via AssignProperties_To_PrivateDnsZones_MX_Spec & AssignProperties_From_PrivateDnsZones_MX_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_MX_Spec, PrivateDnsZones_MX_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_MX_Spec tests if a specific instance of PrivateDnsZones_MX_Spec can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_MX_Spec(subject PrivateDnsZones_MX_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZones_MX_Spec
	err := copied.AssignProperties_To_PrivateDnsZones_MX_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_MX_Spec
	err = actual.AssignProperties_From_PrivateDnsZones_MX_Spec(&other)
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

func Test_PrivateDnsZones_MX_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_MX_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_MX_Spec, PrivateDnsZones_MX_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_MX_Spec runs a test to see if a specific instance of PrivateDnsZones_MX_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_MX_Spec(subject PrivateDnsZones_MX_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_MX_Spec
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

// Generator of PrivateDnsZones_MX_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_MX_SpecGenerator()
var privateDnsZones_MX_SpecGenerator gopter.Gen

// PrivateDnsZones_MX_SpecGenerator returns a generator of PrivateDnsZones_MX_Spec instances for property testing.
// We first initialize privateDnsZones_MX_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_MX_SpecGenerator() gopter.Gen {
	if privateDnsZones_MX_SpecGenerator != nil {
		return privateDnsZones_MX_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_Spec(generators)
	privateDnsZones_MX_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_MX_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_Spec(generators)
	privateDnsZones_MX_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_MX_Spec{}), generators)

	return privateDnsZones_MX_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_MX_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_MX_STATUS to PrivateDnsZones_MX_STATUS via AssignProperties_To_PrivateDnsZones_MX_STATUS & AssignProperties_From_PrivateDnsZones_MX_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_MX_STATUS, PrivateDnsZones_MX_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_MX_STATUS tests if a specific instance of PrivateDnsZones_MX_STATUS can be assigned to v1beta20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_MX_STATUS(subject PrivateDnsZones_MX_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200601s.PrivateDnsZones_MX_STATUS
	err := copied.AssignProperties_To_PrivateDnsZones_MX_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_MX_STATUS
	err = actual.AssignProperties_From_PrivateDnsZones_MX_STATUS(&other)
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

func Test_PrivateDnsZones_MX_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_MX_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_MX_STATUS, PrivateDnsZones_MX_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_MX_STATUS runs a test to see if a specific instance of PrivateDnsZones_MX_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_MX_STATUS(subject PrivateDnsZones_MX_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_MX_STATUS
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

// Generator of PrivateDnsZones_MX_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_MX_STATUSGenerator()
var privateDnsZones_MX_STATUSGenerator gopter.Gen

// PrivateDnsZones_MX_STATUSGenerator returns a generator of PrivateDnsZones_MX_STATUS instances for property testing.
// We first initialize privateDnsZones_MX_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_MX_STATUSGenerator() gopter.Gen {
	if privateDnsZones_MX_STATUSGenerator != nil {
		return privateDnsZones_MX_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_STATUS(generators)
	privateDnsZones_MX_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_MX_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_STATUS(generators)
	privateDnsZones_MX_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_MX_STATUS{}), generators)

	return privateDnsZones_MX_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_MX_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_MX_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
