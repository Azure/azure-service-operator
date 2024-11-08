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

func Test_PrivateDnsZonesPTRRecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesPTRRecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS, PrivateDnsZonesPTRRecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS runs a test to see if a specific instance of PrivateDnsZonesPTRRecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesPTRRecord_STATUS(subject PrivateDnsZonesPTRRecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesPTRRecord_STATUS
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

// Generator of PrivateDnsZonesPTRRecord_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesPTRRecord_STATUSGenerator()
var privateDnsZonesPTRRecord_STATUSGenerator gopter.Gen

// PrivateDnsZonesPTRRecord_STATUSGenerator returns a generator of PrivateDnsZonesPTRRecord_STATUS instances for property testing.
// We first initialize privateDnsZonesPTRRecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesPTRRecord_STATUSGenerator() gopter.Gen {
	if privateDnsZonesPTRRecord_STATUSGenerator != nil {
		return privateDnsZonesPTRRecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	privateDnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(generators)
	privateDnsZonesPTRRecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesPTRRecord_STATUS{}), generators)

	return privateDnsZonesPTRRecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesPTRRecord_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_STATUSGenerator())
}
