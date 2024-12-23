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

func Test_ProductContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductContractProperties_STATUS, ProductContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductContractProperties_STATUS runs a test to see if a specific instance of ProductContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProductContractProperties_STATUS(subject ProductContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductContractProperties_STATUS
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

// Generator of ProductContractProperties_STATUS instances for property testing - lazily instantiated by
// ProductContractProperties_STATUSGenerator()
var productContractProperties_STATUSGenerator gopter.Gen

// ProductContractProperties_STATUSGenerator returns a generator of ProductContractProperties_STATUS instances for property testing.
func ProductContractProperties_STATUSGenerator() gopter.Gen {
	if productContractProperties_STATUSGenerator != nil {
		return productContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductContractProperties_STATUS(generators)
	productContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ProductContractProperties_STATUS{}), generators)

	return productContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProductContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ApprovalRequired"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ProductContractProperties_State_STATUS_NotPublished, ProductContractProperties_State_STATUS_Published))
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["SubscriptionsLimit"] = gen.PtrOf(gen.Int())
	gens["Terms"] = gen.PtrOf(gen.AlphaString())
}

func Test_Product_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Product_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProduct_STATUS, Product_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProduct_STATUS runs a test to see if a specific instance of Product_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProduct_STATUS(subject Product_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Product_STATUS
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

// Generator of Product_STATUS instances for property testing - lazily instantiated by Product_STATUSGenerator()
var product_STATUSGenerator gopter.Gen

// Product_STATUSGenerator returns a generator of Product_STATUS instances for property testing.
// We first initialize product_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Product_STATUSGenerator() gopter.Gen {
	if product_STATUSGenerator != nil {
		return product_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProduct_STATUS(generators)
	product_STATUSGenerator = gen.Struct(reflect.TypeOf(Product_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProduct_STATUS(generators)
	AddRelatedPropertyGeneratorsForProduct_STATUS(generators)
	product_STATUSGenerator = gen.Struct(reflect.TypeOf(Product_STATUS{}), generators)

	return product_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProduct_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProduct_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProduct_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProduct_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProductContractProperties_STATUSGenerator())
}
