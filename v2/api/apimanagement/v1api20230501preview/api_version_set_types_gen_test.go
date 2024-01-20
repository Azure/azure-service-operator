// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	v20230501ps "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
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

func Test_ApiVersionSet_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApiVersionSet to hub returns original",
		prop.ForAll(RunResourceConversionTestForApiVersionSet, ApiVersionSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForApiVersionSet tests if a specific instance of ApiVersionSet round trips to the hub storage version and back losslessly
func RunResourceConversionTestForApiVersionSet(subject ApiVersionSet) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.ApiVersionSet
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ApiVersionSet
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

func Test_ApiVersionSet_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApiVersionSet to ApiVersionSet via AssignProperties_To_ApiVersionSet & AssignProperties_From_ApiVersionSet returns original",
		prop.ForAll(RunPropertyAssignmentTestForApiVersionSet, ApiVersionSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApiVersionSet tests if a specific instance of ApiVersionSet can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForApiVersionSet(subject ApiVersionSet) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.ApiVersionSet
	err := copied.AssignProperties_To_ApiVersionSet(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApiVersionSet
	err = actual.AssignProperties_From_ApiVersionSet(&other)
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

func Test_ApiVersionSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiVersionSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiVersionSet, ApiVersionSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiVersionSet runs a test to see if a specific instance of ApiVersionSet round trips to JSON and back losslessly
func RunJSONSerializationTestForApiVersionSet(subject ApiVersionSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiVersionSet
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

// Generator of ApiVersionSet instances for property testing - lazily instantiated by ApiVersionSetGenerator()
var apiVersionSetGenerator gopter.Gen

// ApiVersionSetGenerator returns a generator of ApiVersionSet instances for property testing.
func ApiVersionSetGenerator() gopter.Gen {
	if apiVersionSetGenerator != nil {
		return apiVersionSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForApiVersionSet(generators)
	apiVersionSetGenerator = gen.Struct(reflect.TypeOf(ApiVersionSet{}), generators)

	return apiVersionSetGenerator
}

// AddRelatedPropertyGeneratorsForApiVersionSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApiVersionSet(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_ApiVersionSet_SpecGenerator()
	gens["Status"] = Service_ApiVersionSet_STATUSGenerator()
}

func Test_Service_ApiVersionSet_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_ApiVersionSet_Spec to Service_ApiVersionSet_Spec via AssignProperties_To_Service_ApiVersionSet_Spec & AssignProperties_From_Service_ApiVersionSet_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_ApiVersionSet_Spec, Service_ApiVersionSet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_ApiVersionSet_Spec tests if a specific instance of Service_ApiVersionSet_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForService_ApiVersionSet_Spec(subject Service_ApiVersionSet_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.Service_ApiVersionSet_Spec
	err := copied.AssignProperties_To_Service_ApiVersionSet_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_ApiVersionSet_Spec
	err = actual.AssignProperties_From_Service_ApiVersionSet_Spec(&other)
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

func Test_Service_ApiVersionSet_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_ApiVersionSet_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_ApiVersionSet_Spec, Service_ApiVersionSet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_ApiVersionSet_Spec runs a test to see if a specific instance of Service_ApiVersionSet_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_ApiVersionSet_Spec(subject Service_ApiVersionSet_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_ApiVersionSet_Spec
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

// Generator of Service_ApiVersionSet_Spec instances for property testing - lazily instantiated by
// Service_ApiVersionSet_SpecGenerator()
var service_ApiVersionSet_SpecGenerator gopter.Gen

// Service_ApiVersionSet_SpecGenerator returns a generator of Service_ApiVersionSet_Spec instances for property testing.
func Service_ApiVersionSet_SpecGenerator() gopter.Gen {
	if service_ApiVersionSet_SpecGenerator != nil {
		return service_ApiVersionSet_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec(generators)
	service_ApiVersionSet_SpecGenerator = gen.Struct(reflect.TypeOf(Service_ApiVersionSet_Spec{}), generators)

	return service_ApiVersionSet_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_ApiVersionSet_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionHeaderName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionQueryName"] = gen.PtrOf(gen.AlphaString())
	gens["VersioningScheme"] = gen.PtrOf(gen.OneConstOf(ApiVersionSetContractProperties_VersioningScheme_Header, ApiVersionSetContractProperties_VersioningScheme_Query, ApiVersionSetContractProperties_VersioningScheme_Segment))
}

func Test_Service_ApiVersionSet_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_ApiVersionSet_STATUS to Service_ApiVersionSet_STATUS via AssignProperties_To_Service_ApiVersionSet_STATUS & AssignProperties_From_Service_ApiVersionSet_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_ApiVersionSet_STATUS, Service_ApiVersionSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_ApiVersionSet_STATUS tests if a specific instance of Service_ApiVersionSet_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForService_ApiVersionSet_STATUS(subject Service_ApiVersionSet_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.Service_ApiVersionSet_STATUS
	err := copied.AssignProperties_To_Service_ApiVersionSet_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_ApiVersionSet_STATUS
	err = actual.AssignProperties_From_Service_ApiVersionSet_STATUS(&other)
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

func Test_Service_ApiVersionSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_ApiVersionSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_ApiVersionSet_STATUS, Service_ApiVersionSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_ApiVersionSet_STATUS runs a test to see if a specific instance of Service_ApiVersionSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_ApiVersionSet_STATUS(subject Service_ApiVersionSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_ApiVersionSet_STATUS
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

// Generator of Service_ApiVersionSet_STATUS instances for property testing - lazily instantiated by
// Service_ApiVersionSet_STATUSGenerator()
var service_ApiVersionSet_STATUSGenerator gopter.Gen

// Service_ApiVersionSet_STATUSGenerator returns a generator of Service_ApiVersionSet_STATUS instances for property testing.
func Service_ApiVersionSet_STATUSGenerator() gopter.Gen {
	if service_ApiVersionSet_STATUSGenerator != nil {
		return service_ApiVersionSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_ApiVersionSet_STATUS(generators)
	service_ApiVersionSet_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_ApiVersionSet_STATUS{}), generators)

	return service_ApiVersionSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_ApiVersionSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_ApiVersionSet_STATUS(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VersionHeaderName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionQueryName"] = gen.PtrOf(gen.AlphaString())
	gens["VersioningScheme"] = gen.PtrOf(gen.OneConstOf(ApiVersionSetContractProperties_VersioningScheme_STATUS_Header, ApiVersionSetContractProperties_VersioningScheme_STATUS_Query, ApiVersionSetContractProperties_VersioningScheme_STATUS_Segment))
}
