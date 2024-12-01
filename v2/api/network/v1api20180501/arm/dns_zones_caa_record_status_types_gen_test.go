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

func Test_DnsZonesCAARecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCAARecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCAARecord_STATUS, DnsZonesCAARecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCAARecord_STATUS runs a test to see if a specific instance of DnsZonesCAARecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCAARecord_STATUS(subject DnsZonesCAARecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCAARecord_STATUS
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

// Generator of DnsZonesCAARecord_STATUS instances for property testing - lazily instantiated by
// DnsZonesCAARecord_STATUSGenerator()
var dnsZonesCAARecord_STATUSGenerator gopter.Gen

// DnsZonesCAARecord_STATUSGenerator returns a generator of DnsZonesCAARecord_STATUS instances for property testing.
// We first initialize dnsZonesCAARecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesCAARecord_STATUSGenerator() gopter.Gen {
	if dnsZonesCAARecord_STATUSGenerator != nil {
		return dnsZonesCAARecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	dnsZonesCAARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	dnsZonesCAARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_STATUS{}), generators)

	return dnsZonesCAARecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_STATUSGenerator())
}
