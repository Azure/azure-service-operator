// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

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

func Test_Domains_Topic_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domains_Topic_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomains_Topic_SpecARM, Domains_Topic_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomains_Topic_SpecARM runs a test to see if a specific instance of Domains_Topic_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomains_Topic_SpecARM(subject Domains_Topic_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domains_Topic_SpecARM
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

// Generator of Domains_Topic_SpecARM instances for property testing - lazily instantiated by
// Domains_Topic_SpecARMGenerator()
var domains_Topic_SpecARMGenerator gopter.Gen

// Domains_Topic_SpecARMGenerator returns a generator of Domains_Topic_SpecARM instances for property testing.
func Domains_Topic_SpecARMGenerator() gopter.Gen {
	if domains_Topic_SpecARMGenerator != nil {
		return domains_Topic_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomains_Topic_SpecARM(generators)
	domains_Topic_SpecARMGenerator = gen.Struct(reflect.TypeOf(Domains_Topic_SpecARM{}), generators)

	return domains_Topic_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDomains_Topic_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomains_Topic_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
