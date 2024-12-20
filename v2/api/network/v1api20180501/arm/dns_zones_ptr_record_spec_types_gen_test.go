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

func Test_DnsZonesPTRRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesPTRRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesPTRRecord_Spec, DnsZonesPTRRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesPTRRecord_Spec runs a test to see if a specific instance of DnsZonesPTRRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesPTRRecord_Spec(subject DnsZonesPTRRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesPTRRecord_Spec
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

// Generator of DnsZonesPTRRecord_Spec instances for property testing - lazily instantiated by
// DnsZonesPTRRecord_SpecGenerator()
var dnsZonesPTRRecord_SpecGenerator gopter.Gen

// DnsZonesPTRRecord_SpecGenerator returns a generator of DnsZonesPTRRecord_Spec instances for property testing.
// We first initialize dnsZonesPTRRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesPTRRecord_SpecGenerator() gopter.Gen {
	if dnsZonesPTRRecord_SpecGenerator != nil {
		return dnsZonesPTRRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	dnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec(generators)
	dnsZonesPTRRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesPTRRecord_Spec{}), generators)

	return dnsZonesPTRRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesPTRRecord_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetPropertiesGenerator())
}
