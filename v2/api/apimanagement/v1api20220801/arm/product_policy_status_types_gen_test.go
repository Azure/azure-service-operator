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

func Test_ProductPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy_STATUS, ProductPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy_STATUS runs a test to see if a specific instance of ProductPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy_STATUS(subject ProductPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy_STATUS
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

// Generator of ProductPolicy_STATUS instances for property testing - lazily instantiated by
// ProductPolicy_STATUSGenerator()
var productPolicy_STATUSGenerator gopter.Gen

// ProductPolicy_STATUSGenerator returns a generator of ProductPolicy_STATUS instances for property testing.
// We first initialize productPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProductPolicy_STATUSGenerator() gopter.Gen {
	if productPolicy_STATUSGenerator != nil {
		return productPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_STATUS(generators)
	productPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForProductPolicy_STATUS(generators)
	productPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_STATUS{}), generators)

	return productPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProductPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProductPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyContractProperties_STATUSGenerator())
}
