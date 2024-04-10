// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101/storage"
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

func Test_PrivateDnsZonesVirtualNetworkLink_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink, PrivateDnsZonesVirtualNetworkLinkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink(subject PrivateDnsZonesVirtualNetworkLink) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink
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

// Generator of PrivateDnsZonesVirtualNetworkLink instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLinkGenerator()
var privateDnsZonesVirtualNetworkLinkGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLinkGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink instances for property testing.
func PrivateDnsZonesVirtualNetworkLinkGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLinkGenerator != nil {
		return privateDnsZonesVirtualNetworkLinkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink(generators)
	privateDnsZonesVirtualNetworkLinkGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink{}), generators)

	return privateDnsZonesVirtualNetworkLinkGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_VirtualNetworkLink_SpecGenerator()
	gens["Status"] = PrivateDnsZones_VirtualNetworkLink_STATUSGenerator()
}

func Test_PrivateDnsZones_VirtualNetworkLink_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_VirtualNetworkLink_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec, PrivateDnsZones_VirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec runs a test to see if a specific instance of PrivateDnsZones_VirtualNetworkLink_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec(subject PrivateDnsZones_VirtualNetworkLink_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_VirtualNetworkLink_Spec
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

// Generator of PrivateDnsZones_VirtualNetworkLink_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_VirtualNetworkLink_SpecGenerator()
var privateDnsZones_VirtualNetworkLink_SpecGenerator gopter.Gen

// PrivateDnsZones_VirtualNetworkLink_SpecGenerator returns a generator of PrivateDnsZones_VirtualNetworkLink_Spec instances for property testing.
// We first initialize privateDnsZones_VirtualNetworkLink_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_VirtualNetworkLink_SpecGenerator() gopter.Gen {
	if privateDnsZones_VirtualNetworkLink_SpecGenerator != nil {
		return privateDnsZones_VirtualNetworkLink_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec(generators)
	privateDnsZones_VirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec(generators)
	privateDnsZones_VirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_Spec{}), generators)

	return privateDnsZones_VirtualNetworkLink_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}

func Test_PrivateDnsZones_VirtualNetworkLink_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_VirtualNetworkLink_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_STATUS, PrivateDnsZones_VirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_STATUS runs a test to see if a specific instance of PrivateDnsZones_VirtualNetworkLink_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_STATUS(subject PrivateDnsZones_VirtualNetworkLink_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_VirtualNetworkLink_STATUS
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

// Generator of PrivateDnsZones_VirtualNetworkLink_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_VirtualNetworkLink_STATUSGenerator()
var privateDnsZones_VirtualNetworkLink_STATUSGenerator gopter.Gen

// PrivateDnsZones_VirtualNetworkLink_STATUSGenerator returns a generator of PrivateDnsZones_VirtualNetworkLink_STATUS instances for property testing.
// We first initialize privateDnsZones_VirtualNetworkLink_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_VirtualNetworkLink_STATUSGenerator() gopter.Gen {
	if privateDnsZones_VirtualNetworkLink_STATUSGenerator != nil {
		return privateDnsZones_VirtualNetworkLink_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS(generators)
	privateDnsZones_VirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS(generators)
	privateDnsZones_VirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_STATUS{}), generators)

	return privateDnsZones_VirtualNetworkLink_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkLinkState"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_SubResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource to SubResource via AssignProperties_To_SubResource & AssignProperties_From_SubResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource tests if a specific instance of SubResource can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSubResource(subject SubResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SubResource
	err := copied.AssignProperties_To_SubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource
	err = actual.AssignProperties_From_SubResource(&other)
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

func Test_SubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource runs a test to see if a specific instance of SubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource(subject SubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource
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

// Generator of SubResource instances for property testing - lazily instantiated by SubResourceGenerator()
var subResourceGenerator gopter.Gen

// SubResourceGenerator returns a generator of SubResource instances for property testing.
func SubResourceGenerator() gopter.Gen {
	if subResourceGenerator != nil {
		return subResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

func Test_SubResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource_STATUS to SubResource_STATUS via AssignProperties_To_SubResource_STATUS & AssignProperties_From_SubResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource_STATUS tests if a specific instance of SubResource_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SubResource_STATUS
	err := copied.AssignProperties_To_SubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource_STATUS
	err = actual.AssignProperties_From_SubResource_STATUS(&other)
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

func Test_SubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS runs a test to see if a specific instance of SubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS
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

// Generator of SubResource_STATUS instances for property testing - lazily instantiated by SubResource_STATUSGenerator()
var subResource_STATUSGenerator gopter.Gen

// SubResource_STATUSGenerator returns a generator of SubResource_STATUS instances for property testing.
func SubResource_STATUSGenerator() gopter.Gen {
	if subResource_STATUSGenerator != nil {
		return subResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS(generators)
	subResource_STATUSGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS{}), generators)

	return subResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
