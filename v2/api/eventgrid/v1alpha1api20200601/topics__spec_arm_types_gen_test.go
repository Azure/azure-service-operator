// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

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

func Test_Topics_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topics_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicsSpecARM, TopicsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicsSpecARM runs a test to see if a specific instance of Topics_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicsSpecARM(subject Topics_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topics_SpecARM
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

// Generator of Topics_SpecARM instances for property testing - lazily instantiated by TopicsSpecARMGenerator()
var topicsSpecARMGenerator gopter.Gen

// TopicsSpecARMGenerator returns a generator of Topics_SpecARM instances for property testing.
func TopicsSpecARMGenerator() gopter.Gen {
	if topicsSpecARMGenerator != nil {
		return topicsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicsSpecARM(generators)
	topicsSpecARMGenerator = gen.Struct(reflect.TypeOf(Topics_SpecARM{}), generators)

	return topicsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForTopicsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
