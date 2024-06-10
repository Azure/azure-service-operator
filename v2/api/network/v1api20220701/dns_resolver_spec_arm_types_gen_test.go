// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_DnsResolverProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolverProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolverProperties_ARM, DnsResolverProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolverProperties_ARM runs a test to see if a specific instance of DnsResolverProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolverProperties_ARM(subject DnsResolverProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolverProperties_ARM
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

// Generator of DnsResolverProperties_ARM instances for property testing - lazily instantiated by
// DnsResolverProperties_ARMGenerator()
var dnsResolverProperties_ARMGenerator gopter.Gen

// DnsResolverProperties_ARMGenerator returns a generator of DnsResolverProperties_ARM instances for property testing.
func DnsResolverProperties_ARMGenerator() gopter.Gen {
	if dnsResolverProperties_ARMGenerator != nil {
		return dnsResolverProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsResolverProperties_ARM(generators)
	dnsResolverProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolverProperties_ARM{}), generators)

	return dnsResolverProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForDnsResolverProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolverProperties_ARM(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(DnsresolverSubResource_ARMGenerator())
}

func Test_DnsResolver_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolver_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolver_Spec_ARM, DnsResolver_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolver_Spec_ARM runs a test to see if a specific instance of DnsResolver_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolver_Spec_ARM(subject DnsResolver_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolver_Spec_ARM
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

// Generator of DnsResolver_Spec_ARM instances for property testing - lazily instantiated by
// DnsResolver_Spec_ARMGenerator()
var dnsResolver_Spec_ARMGenerator gopter.Gen

// DnsResolver_Spec_ARMGenerator returns a generator of DnsResolver_Spec_ARM instances for property testing.
// We first initialize dnsResolver_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolver_Spec_ARMGenerator() gopter.Gen {
	if dnsResolver_Spec_ARMGenerator != nil {
		return dnsResolver_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec_ARM(generators)
	dnsResolver_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsResolver_Spec_ARM(generators)
	dnsResolver_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec_ARM{}), generators)

	return dnsResolver_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolver_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolver_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolver_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolver_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DnsResolverProperties_ARMGenerator())
}
