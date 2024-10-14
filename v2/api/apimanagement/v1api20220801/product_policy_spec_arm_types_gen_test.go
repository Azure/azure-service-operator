// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

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

func Test_ProductPolicy_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy_Spec_ARM, ProductPolicy_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy_Spec_ARM runs a test to see if a specific instance of ProductPolicy_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy_Spec_ARM(subject ProductPolicy_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy_Spec_ARM
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

// Generator of ProductPolicy_Spec_ARM instances for property testing - lazily instantiated by
// ProductPolicy_Spec_ARMGenerator()
var productPolicy_Spec_ARMGenerator gopter.Gen

// ProductPolicy_Spec_ARMGenerator returns a generator of ProductPolicy_Spec_ARM instances for property testing.
// We first initialize productPolicy_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProductPolicy_Spec_ARMGenerator() gopter.Gen {
	if productPolicy_Spec_ARMGenerator != nil {
		return productPolicy_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec_ARM(generators)
	productPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForProductPolicy_Spec_ARM(generators)
	productPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec_ARM{}), generators)

	return productPolicy_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProductPolicy_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProductPolicy_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyContractProperties_ARMGenerator())
}
