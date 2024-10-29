// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

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

func Test_DomainsTopic_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_Spec_ARM, DomainsTopic_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_Spec_ARM runs a test to see if a specific instance of DomainsTopic_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_Spec_ARM(subject DomainsTopic_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_Spec_ARM
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

// Generator of DomainsTopic_Spec_ARM instances for property testing - lazily instantiated by
// DomainsTopic_Spec_ARMGenerator()
var domainsTopic_Spec_ARMGenerator gopter.Gen

// DomainsTopic_Spec_ARMGenerator returns a generator of DomainsTopic_Spec_ARM instances for property testing.
func DomainsTopic_Spec_ARMGenerator() gopter.Gen {
	if domainsTopic_Spec_ARMGenerator != nil {
		return domainsTopic_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_Spec_ARM(generators)
	domainsTopic_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_Spec_ARM{}), generators)

	return domainsTopic_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}
