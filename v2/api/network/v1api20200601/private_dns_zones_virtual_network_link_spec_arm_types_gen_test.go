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

func Test_PrivateDnsZones_VirtualNetworkLink_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_VirtualNetworkLink_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec_ARM, PrivateDnsZones_VirtualNetworkLink_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec_ARM runs a test to see if a specific instance of PrivateDnsZones_VirtualNetworkLink_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(subject PrivateDnsZones_VirtualNetworkLink_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_VirtualNetworkLink_Spec_ARM
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

// Generator of PrivateDnsZones_VirtualNetworkLink_Spec_ARM instances for property testing - lazily instantiated by
// PrivateDnsZones_VirtualNetworkLink_Spec_ARMGenerator()
var privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator gopter.Gen

// PrivateDnsZones_VirtualNetworkLink_Spec_ARMGenerator returns a generator of PrivateDnsZones_VirtualNetworkLink_Spec_ARM instances for property testing.
// We first initialize privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_VirtualNetworkLink_Spec_ARMGenerator() gopter.Gen {
	if privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator != nil {
		return privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(generators)
	privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(generators)
	privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_VirtualNetworkLink_Spec_ARM{}), generators)

	return privateDnsZones_VirtualNetworkLink_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_VirtualNetworkLink_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkLinkProperties_ARMGenerator())
}

func Test_SubResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_ARM, SubResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_ARM runs a test to see if a specific instance of SubResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_ARM(subject SubResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_ARM
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

// Generator of SubResource_ARM instances for property testing - lazily instantiated by SubResource_ARMGenerator()
var subResource_ARMGenerator gopter.Gen

// SubResource_ARMGenerator returns a generator of SubResource_ARM instances for property testing.
func SubResource_ARMGenerator() gopter.Gen {
	if subResource_ARMGenerator != nil {
		return subResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_ARM(generators)
	subResource_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_ARM{}), generators)

	return subResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkLinkProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkLinkProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkLinkProperties_ARM, VirtualNetworkLinkProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkLinkProperties_ARM runs a test to see if a specific instance of VirtualNetworkLinkProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkLinkProperties_ARM(subject VirtualNetworkLinkProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkLinkProperties_ARM
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

// Generator of VirtualNetworkLinkProperties_ARM instances for property testing - lazily instantiated by
// VirtualNetworkLinkProperties_ARMGenerator()
var virtualNetworkLinkProperties_ARMGenerator gopter.Gen

// VirtualNetworkLinkProperties_ARMGenerator returns a generator of VirtualNetworkLinkProperties_ARM instances for property testing.
// We first initialize virtualNetworkLinkProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkLinkProperties_ARMGenerator() gopter.Gen {
	if virtualNetworkLinkProperties_ARMGenerator != nil {
		return virtualNetworkLinkProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_ARM(generators)
	virtualNetworkLinkProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_ARM(generators)
	virtualNetworkLinkProperties_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties_ARM{}), generators)

	return virtualNetworkLinkProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_ARM(gens map[string]gopter.Gen) {
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_ARM(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResource_ARMGenerator())
}
