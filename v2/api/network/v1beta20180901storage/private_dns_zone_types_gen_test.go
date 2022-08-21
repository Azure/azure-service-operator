// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901storage

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

func Test_PrivateDnsZone_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone, PrivateDnsZoneGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone runs a test to see if a specific instance of PrivateDnsZone round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone(subject PrivateDnsZone) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone
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

// Generator of PrivateDnsZone instances for property testing - lazily instantiated by PrivateDnsZoneGenerator()
var privateDnsZoneGenerator gopter.Gen

// PrivateDnsZoneGenerator returns a generator of PrivateDnsZone instances for property testing.
func PrivateDnsZoneGenerator() gopter.Gen {
	if privateDnsZoneGenerator != nil {
		return privateDnsZoneGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZone(generators)
	privateDnsZoneGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone{}), generators)

	return privateDnsZoneGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZone is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZone(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZone_SpecGenerator()
	gens["Status"] = PrivateDnsZone_STATUSGenerator()
}

func Test_PrivateDnsZone_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_Spec, PrivateDnsZone_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_Spec runs a test to see if a specific instance of PrivateDnsZone_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_Spec(subject PrivateDnsZone_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_Spec
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

// Generator of PrivateDnsZone_Spec instances for property testing - lazily instantiated by
// PrivateDnsZone_SpecGenerator()
var privateDnsZone_SpecGenerator gopter.Gen

// PrivateDnsZone_SpecGenerator returns a generator of PrivateDnsZone_Spec instances for property testing.
func PrivateDnsZone_SpecGenerator() gopter.Gen {
	if privateDnsZone_SpecGenerator != nil {
		return privateDnsZone_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec(generators)
	privateDnsZone_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_Spec{}), generators)

	return privateDnsZone_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_PrivateDnsZone_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_STATUS, PrivateDnsZone_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_STATUS runs a test to see if a specific instance of PrivateDnsZone_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_STATUS(subject PrivateDnsZone_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_STATUS
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

// Generator of PrivateDnsZone_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZone_STATUSGenerator()
var privateDnsZone_STATUSGenerator gopter.Gen

// PrivateDnsZone_STATUSGenerator returns a generator of PrivateDnsZone_STATUS instances for property testing.
func PrivateDnsZone_STATUSGenerator() gopter.Gen {
	if privateDnsZone_STATUSGenerator != nil {
		return privateDnsZone_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS(generators)
	privateDnsZone_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_STATUS{}), generators)

	return privateDnsZone_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["MaxNumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["NumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}
