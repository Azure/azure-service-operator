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

func Test_DnsResolverProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolverProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolverProperties, DnsResolverPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolverProperties runs a test to see if a specific instance of DnsResolverProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolverProperties(subject DnsResolverProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolverProperties
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

// Generator of DnsResolverProperties instances for property testing - lazily instantiated by
// DnsResolverPropertiesGenerator()
var dnsResolverPropertiesGenerator gopter.Gen

// DnsResolverPropertiesGenerator returns a generator of DnsResolverProperties instances for property testing.
func DnsResolverPropertiesGenerator() gopter.Gen {
	if dnsResolverPropertiesGenerator != nil {
		return dnsResolverPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsResolverProperties(generators)
	dnsResolverPropertiesGenerator = gen.Struct(reflect.TypeOf(DnsResolverProperties{}), generators)

	return dnsResolverPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForDnsResolverProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolverProperties(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
}

func Test_DnsResolver_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolver_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolver_Spec, DnsResolver_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolver_Spec runs a test to see if a specific instance of DnsResolver_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolver_Spec(subject DnsResolver_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolver_Spec
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

// Generator of DnsResolver_Spec instances for property testing - lazily instantiated by DnsResolver_SpecGenerator()
var dnsResolver_SpecGenerator gopter.Gen

// DnsResolver_SpecGenerator returns a generator of DnsResolver_Spec instances for property testing.
// We first initialize dnsResolver_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolver_SpecGenerator() gopter.Gen {
	if dnsResolver_SpecGenerator != nil {
		return dnsResolver_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec(generators)
	dnsResolver_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolver_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsResolver_Spec(generators)
	dnsResolver_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolver_Spec{}), generators)

	return dnsResolver_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolver_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolver_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolver_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolver_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DnsResolverPropertiesGenerator())
}
