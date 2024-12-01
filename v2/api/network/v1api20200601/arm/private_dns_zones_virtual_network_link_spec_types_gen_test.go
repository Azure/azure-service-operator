// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_PrivateDnsZonesVirtualNetworkLink_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec, PrivateDnsZonesVirtualNetworkLink_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_Spec(subject PrivateDnsZonesVirtualNetworkLink_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink_Spec
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

// Generator of PrivateDnsZonesVirtualNetworkLink_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLink_SpecGenerator()
var privateDnsZonesVirtualNetworkLink_SpecGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLink_SpecGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink_Spec instances for property testing.
// We first initialize privateDnsZonesVirtualNetworkLink_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesVirtualNetworkLink_SpecGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLink_SpecGenerator != nil {
		return privateDnsZonesVirtualNetworkLink_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	privateDnsZonesVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(generators)
	privateDnsZonesVirtualNetworkLink_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_Spec{}), generators)

	return privateDnsZonesVirtualNetworkLink_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkLinkPropertiesGenerator())
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
	AddIndependentPropertyGeneratorsForSubResource(generators)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

// AddIndependentPropertyGeneratorsForSubResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkLinkProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkLinkProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkLinkProperties, VirtualNetworkLinkPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkLinkProperties runs a test to see if a specific instance of VirtualNetworkLinkProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkLinkProperties(subject VirtualNetworkLinkProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkLinkProperties
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

// Generator of VirtualNetworkLinkProperties instances for property testing - lazily instantiated by
// VirtualNetworkLinkPropertiesGenerator()
var virtualNetworkLinkPropertiesGenerator gopter.Gen

// VirtualNetworkLinkPropertiesGenerator returns a generator of VirtualNetworkLinkProperties instances for property testing.
// We first initialize virtualNetworkLinkPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkLinkPropertiesGenerator() gopter.Gen {
	if virtualNetworkLinkPropertiesGenerator != nil {
		return virtualNetworkLinkPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	virtualNetworkLinkPropertiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties(generators)
	virtualNetworkLinkPropertiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties{}), generators)

	return virtualNetworkLinkPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties(gens map[string]gopter.Gen) {
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}
