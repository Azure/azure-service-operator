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

func Test_DnsZones_A_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_A_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_A_STATUS_ARM, DnsZones_A_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_A_STATUS_ARM runs a test to see if a specific instance of DnsZones_A_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_A_STATUS_ARM(subject DnsZones_A_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_A_STATUS_ARM
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

// Generator of DnsZones_A_STATUS_ARM instances for property testing - lazily instantiated by
// DnsZones_A_STATUS_ARMGenerator()
var dnsZones_A_STATUS_ARMGenerator gopter.Gen

// DnsZones_A_STATUS_ARMGenerator returns a generator of DnsZones_A_STATUS_ARM instances for property testing.
// We first initialize dnsZones_A_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_A_STATUS_ARMGenerator() gopter.Gen {
	if dnsZones_A_STATUS_ARMGenerator != nil {
		return dnsZones_A_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_STATUS_ARM(generators)
	dnsZones_A_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_A_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsZones_A_STATUS_ARM(generators)
	dnsZones_A_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZones_A_STATUS_ARM{}), generators)

	return dnsZones_A_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_A_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_A_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZones_A_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_A_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_STATUS_ARMGenerator())
}
