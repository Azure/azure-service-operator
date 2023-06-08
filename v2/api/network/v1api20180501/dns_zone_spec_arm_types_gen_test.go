// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

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

func Test_DnsZone_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZone_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZone_Spec_ARM, DnsZone_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZone_Spec_ARM runs a test to see if a specific instance of DnsZone_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZone_Spec_ARM(subject DnsZone_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZone_Spec_ARM
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

// Generator of DnsZone_Spec_ARM instances for property testing - lazily instantiated by DnsZone_Spec_ARMGenerator()
var dnsZone_Spec_ARMGenerator gopter.Gen

// DnsZone_Spec_ARMGenerator returns a generator of DnsZone_Spec_ARM instances for property testing.
// We first initialize dnsZone_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZone_Spec_ARMGenerator() gopter.Gen {
	if dnsZone_Spec_ARMGenerator != nil {
		return dnsZone_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZone_Spec_ARM(generators)
	dnsZone_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZone_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZone_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsZone_Spec_ARM(generators)
	dnsZone_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZone_Spec_ARM{}), generators)

	return dnsZone_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsZone_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZone_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZone_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZone_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ZoneProperties_ARMGenerator())
}

func Test_ZoneProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ZoneProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForZoneProperties_ARM, ZoneProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForZoneProperties_ARM runs a test to see if a specific instance of ZoneProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForZoneProperties_ARM(subject ZoneProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ZoneProperties_ARM
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

// Generator of ZoneProperties_ARM instances for property testing - lazily instantiated by ZoneProperties_ARMGenerator()
var zoneProperties_ARMGenerator gopter.Gen

// ZoneProperties_ARMGenerator returns a generator of ZoneProperties_ARM instances for property testing.
// We first initialize zoneProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ZoneProperties_ARMGenerator() gopter.Gen {
	if zoneProperties_ARMGenerator != nil {
		return zoneProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForZoneProperties_ARM(generators)
	zoneProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ZoneProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForZoneProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForZoneProperties_ARM(generators)
	zoneProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ZoneProperties_ARM{}), generators)

	return zoneProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForZoneProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForZoneProperties_ARM(gens map[string]gopter.Gen) {
	gens["ZoneType"] = gen.PtrOf(gen.OneConstOf(ZoneProperties_ZoneType_Private, ZoneProperties_ZoneType_Public))
}

// AddRelatedPropertyGeneratorsForZoneProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForZoneProperties_ARM(gens map[string]gopter.Gen) {
	gens["RegistrationVirtualNetworks"] = gen.SliceOf(SubResource_ARMGenerator())
	gens["ResolutionVirtualNetworks"] = gen.SliceOf(SubResource_ARMGenerator())
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