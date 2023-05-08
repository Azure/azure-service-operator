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

func Test_DnsZones_TXT_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_TXT_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_TXT_Spec_ARM, DnsZones_TXT_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_TXT_Spec_ARM runs a test to see if a specific instance of DnsZones_TXT_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_TXT_Spec_ARM(subject DnsZones_TXT_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_TXT_Spec_ARM
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

// Generator of DnsZones_TXT_Spec_ARM instances for property testing - lazily instantiated by
// DnsZones_TXT_Spec_ARMGenerator()
var dnsZones_TXT_Spec_ARMGenerator gopter.Gen

// DnsZones_TXT_Spec_ARMGenerator returns a generator of DnsZones_TXT_Spec_ARM instances for property testing.
// We first initialize dnsZones_TXT_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_TXT_Spec_ARMGenerator() gopter.Gen {
	if dnsZones_TXT_Spec_ARMGenerator != nil {
		return dnsZones_TXT_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec_ARM(generators)
	dnsZones_TXT_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec_ARM(generators)
	dnsZones_TXT_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_Spec_ARM{}), generators)

	return dnsZones_TXT_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_ARMGenerator())
}
