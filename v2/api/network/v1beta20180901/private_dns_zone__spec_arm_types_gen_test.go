// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901

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

func Test_PrivateDnsZone_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_SpecARM, PrivateDnsZone_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_SpecARM runs a test to see if a specific instance of PrivateDnsZone_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_SpecARM(subject PrivateDnsZone_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_SpecARM
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

// Generator of PrivateDnsZone_SpecARM instances for property testing - lazily instantiated by
// PrivateDnsZone_SpecARMGenerator()
var privateDnsZone_SpecARMGenerator gopter.Gen

// PrivateDnsZone_SpecARMGenerator returns a generator of PrivateDnsZone_SpecARM instances for property testing.
func PrivateDnsZone_SpecARMGenerator() gopter.Gen {
	if privateDnsZone_SpecARMGenerator != nil {
		return privateDnsZone_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_SpecARM(generators)
	privateDnsZone_SpecARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_SpecARM{}), generators)

	return privateDnsZone_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}
