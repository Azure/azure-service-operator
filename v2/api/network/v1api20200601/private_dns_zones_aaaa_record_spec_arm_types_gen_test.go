// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

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

func Test_ARecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ARecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForARecord_ARM, ARecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForARecord_ARM runs a test to see if a specific instance of ARecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForARecord_ARM(subject ARecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ARecord_ARM
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

// Generator of ARecord_ARM instances for property testing - lazily instantiated by ARecord_ARMGenerator()
var aRecord_ARMGenerator gopter.Gen

// ARecord_ARMGenerator returns a generator of ARecord_ARM instances for property testing.
func ARecord_ARMGenerator() gopter.Gen {
	if aRecord_ARMGenerator != nil {
		return aRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForARecord_ARM(generators)
	aRecord_ARMGenerator = gen.Struct(reflect.TypeOf(ARecord_ARM{}), generators)

	return aRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForARecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForARecord_ARM(gens map[string]gopter.Gen) {
	gens["Ipv4Address"] = gen.PtrOf(gen.AlphaString())
}

func Test_AaaaRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AaaaRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAaaaRecord_ARM, AaaaRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAaaaRecord_ARM runs a test to see if a specific instance of AaaaRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAaaaRecord_ARM(subject AaaaRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AaaaRecord_ARM
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

// Generator of AaaaRecord_ARM instances for property testing - lazily instantiated by AaaaRecord_ARMGenerator()
var aaaaRecord_ARMGenerator gopter.Gen

// AaaaRecord_ARMGenerator returns a generator of AaaaRecord_ARM instances for property testing.
func AaaaRecord_ARMGenerator() gopter.Gen {
	if aaaaRecord_ARMGenerator != nil {
		return aaaaRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAaaaRecord_ARM(generators)
	aaaaRecord_ARMGenerator = gen.Struct(reflect.TypeOf(AaaaRecord_ARM{}), generators)

	return aaaaRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAaaaRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAaaaRecord_ARM(gens map[string]gopter.Gen) {
	gens["Ipv6Address"] = gen.PtrOf(gen.AlphaString())
}

func Test_CnameRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CnameRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCnameRecord_ARM, CnameRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCnameRecord_ARM runs a test to see if a specific instance of CnameRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCnameRecord_ARM(subject CnameRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CnameRecord_ARM
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

// Generator of CnameRecord_ARM instances for property testing - lazily instantiated by CnameRecord_ARMGenerator()
var cnameRecord_ARMGenerator gopter.Gen

// CnameRecord_ARMGenerator returns a generator of CnameRecord_ARM instances for property testing.
func CnameRecord_ARMGenerator() gopter.Gen {
	if cnameRecord_ARMGenerator != nil {
		return cnameRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCnameRecord_ARM(generators)
	cnameRecord_ARMGenerator = gen.Struct(reflect.TypeOf(CnameRecord_ARM{}), generators)

	return cnameRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCnameRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCnameRecord_ARM(gens map[string]gopter.Gen) {
	gens["Cname"] = gen.PtrOf(gen.AlphaString())
}

func Test_MxRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MxRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMxRecord_ARM, MxRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMxRecord_ARM runs a test to see if a specific instance of MxRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMxRecord_ARM(subject MxRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MxRecord_ARM
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

// Generator of MxRecord_ARM instances for property testing - lazily instantiated by MxRecord_ARMGenerator()
var mxRecord_ARMGenerator gopter.Gen

// MxRecord_ARMGenerator returns a generator of MxRecord_ARM instances for property testing.
func MxRecord_ARMGenerator() gopter.Gen {
	if mxRecord_ARMGenerator != nil {
		return mxRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMxRecord_ARM(generators)
	mxRecord_ARMGenerator = gen.Struct(reflect.TypeOf(MxRecord_ARM{}), generators)

	return mxRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMxRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMxRecord_ARM(gens map[string]gopter.Gen) {
	gens["Exchange"] = gen.PtrOf(gen.AlphaString())
	gens["Preference"] = gen.PtrOf(gen.Int())
}

func Test_PrivateDnsZonesAAAARecord_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesAAAARecord_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesAAAARecord_Spec_ARM, PrivateDnsZonesAAAARecord_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesAAAARecord_Spec_ARM runs a test to see if a specific instance of PrivateDnsZonesAAAARecord_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesAAAARecord_Spec_ARM(subject PrivateDnsZonesAAAARecord_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesAAAARecord_Spec_ARM
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

// Generator of PrivateDnsZonesAAAARecord_Spec_ARM instances for property testing - lazily instantiated by
// PrivateDnsZonesAAAARecord_Spec_ARMGenerator()
var privateDnsZonesAAAARecord_Spec_ARMGenerator gopter.Gen

// PrivateDnsZonesAAAARecord_Spec_ARMGenerator returns a generator of PrivateDnsZonesAAAARecord_Spec_ARM instances for property testing.
// We first initialize privateDnsZonesAAAARecord_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesAAAARecord_Spec_ARMGenerator() gopter.Gen {
	if privateDnsZonesAAAARecord_Spec_ARMGenerator != nil {
		return privateDnsZonesAAAARecord_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM(generators)
	privateDnsZonesAAAARecord_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesAAAARecord_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM(generators)
	privateDnsZonesAAAARecord_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesAAAARecord_Spec_ARM{}), generators)

	return privateDnsZonesAAAARecord_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesAAAARecord_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_ARMGenerator())
}

func Test_PtrRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PtrRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPtrRecord_ARM, PtrRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPtrRecord_ARM runs a test to see if a specific instance of PtrRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPtrRecord_ARM(subject PtrRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PtrRecord_ARM
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

// Generator of PtrRecord_ARM instances for property testing - lazily instantiated by PtrRecord_ARMGenerator()
var ptrRecord_ARMGenerator gopter.Gen

// PtrRecord_ARMGenerator returns a generator of PtrRecord_ARM instances for property testing.
func PtrRecord_ARMGenerator() gopter.Gen {
	if ptrRecord_ARMGenerator != nil {
		return ptrRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPtrRecord_ARM(generators)
	ptrRecord_ARMGenerator = gen.Struct(reflect.TypeOf(PtrRecord_ARM{}), generators)

	return ptrRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPtrRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPtrRecord_ARM(gens map[string]gopter.Gen) {
	gens["Ptrdname"] = gen.PtrOf(gen.AlphaString())
}

func Test_RecordSetProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RecordSetProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRecordSetProperties_ARM, RecordSetProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRecordSetProperties_ARM runs a test to see if a specific instance of RecordSetProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRecordSetProperties_ARM(subject RecordSetProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RecordSetProperties_ARM
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

// Generator of RecordSetProperties_ARM instances for property testing - lazily instantiated by
// RecordSetProperties_ARMGenerator()
var recordSetProperties_ARMGenerator gopter.Gen

// RecordSetProperties_ARMGenerator returns a generator of RecordSetProperties_ARM instances for property testing.
// We first initialize recordSetProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RecordSetProperties_ARMGenerator() gopter.Gen {
	if recordSetProperties_ARMGenerator != nil {
		return recordSetProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSetProperties_ARM(generators)
	recordSetProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RecordSetProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSetProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForRecordSetProperties_ARM(generators)
	recordSetProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RecordSetProperties_ARM{}), generators)

	return recordSetProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRecordSetProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRecordSetProperties_ARM(gens map[string]gopter.Gen) {
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForRecordSetProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRecordSetProperties_ARM(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_ARMGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_ARMGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_ARMGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_ARMGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_ARMGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_ARMGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_ARMGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_ARMGenerator())
}

func Test_SoaRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SoaRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSoaRecord_ARM, SoaRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSoaRecord_ARM runs a test to see if a specific instance of SoaRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSoaRecord_ARM(subject SoaRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SoaRecord_ARM
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

// Generator of SoaRecord_ARM instances for property testing - lazily instantiated by SoaRecord_ARMGenerator()
var soaRecord_ARMGenerator gopter.Gen

// SoaRecord_ARMGenerator returns a generator of SoaRecord_ARM instances for property testing.
func SoaRecord_ARMGenerator() gopter.Gen {
	if soaRecord_ARMGenerator != nil {
		return soaRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSoaRecord_ARM(generators)
	soaRecord_ARMGenerator = gen.Struct(reflect.TypeOf(SoaRecord_ARM{}), generators)

	return soaRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSoaRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSoaRecord_ARM(gens map[string]gopter.Gen) {
	gens["Email"] = gen.PtrOf(gen.AlphaString())
	gens["ExpireTime"] = gen.PtrOf(gen.Int())
	gens["Host"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTtl"] = gen.PtrOf(gen.Int())
	gens["RefreshTime"] = gen.PtrOf(gen.Int())
	gens["RetryTime"] = gen.PtrOf(gen.Int())
	gens["SerialNumber"] = gen.PtrOf(gen.Int())
}

func Test_SrvRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SrvRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSrvRecord_ARM, SrvRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSrvRecord_ARM runs a test to see if a specific instance of SrvRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSrvRecord_ARM(subject SrvRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SrvRecord_ARM
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

// Generator of SrvRecord_ARM instances for property testing - lazily instantiated by SrvRecord_ARMGenerator()
var srvRecord_ARMGenerator gopter.Gen

// SrvRecord_ARMGenerator returns a generator of SrvRecord_ARM instances for property testing.
func SrvRecord_ARMGenerator() gopter.Gen {
	if srvRecord_ARMGenerator != nil {
		return srvRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSrvRecord_ARM(generators)
	srvRecord_ARMGenerator = gen.Struct(reflect.TypeOf(SrvRecord_ARM{}), generators)

	return srvRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSrvRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSrvRecord_ARM(gens map[string]gopter.Gen) {
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

func Test_TxtRecord_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TxtRecord_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTxtRecord_ARM, TxtRecord_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTxtRecord_ARM runs a test to see if a specific instance of TxtRecord_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTxtRecord_ARM(subject TxtRecord_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TxtRecord_ARM
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

// Generator of TxtRecord_ARM instances for property testing - lazily instantiated by TxtRecord_ARMGenerator()
var txtRecord_ARMGenerator gopter.Gen

// TxtRecord_ARMGenerator returns a generator of TxtRecord_ARM instances for property testing.
func TxtRecord_ARMGenerator() gopter.Gen {
	if txtRecord_ARMGenerator != nil {
		return txtRecord_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTxtRecord_ARM(generators)
	txtRecord_ARMGenerator = gen.Struct(reflect.TypeOf(TxtRecord_ARM{}), generators)

	return txtRecord_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTxtRecord_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTxtRecord_ARM(gens map[string]gopter.Gen) {
	gens["Value"] = gen.SliceOf(gen.AlphaString())
}
