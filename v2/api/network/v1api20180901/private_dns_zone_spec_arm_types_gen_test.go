// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180901

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

func Test_PrivateDnsZone_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_Spec_ARM, PrivateDnsZone_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_Spec_ARM runs a test to see if a specific instance of PrivateDnsZone_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_Spec_ARM(subject PrivateDnsZone_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_Spec_ARM
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

// Generator of PrivateDnsZone_Spec_ARM instances for property testing - lazily instantiated by
// PrivateDnsZone_Spec_ARMGenerator()
var privateDnsZone_Spec_ARMGenerator gopter.Gen

// PrivateDnsZone_Spec_ARMGenerator returns a generator of PrivateDnsZone_Spec_ARM instances for property testing.
func PrivateDnsZone_Spec_ARMGenerator() gopter.Gen {
	if privateDnsZone_Spec_ARMGenerator != nil {
		return privateDnsZone_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec_ARM(generators)
	privateDnsZone_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_Spec_ARM{}), generators)

	return privateDnsZone_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}
